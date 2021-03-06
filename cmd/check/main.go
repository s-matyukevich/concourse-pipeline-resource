package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/robdimsdale/concourse-pipeline-resource/check"
	"github.com/robdimsdale/concourse-pipeline-resource/concourse"
	"github.com/robdimsdale/concourse-pipeline-resource/concourse/api"
	"github.com/robdimsdale/concourse-pipeline-resource/logger"
	"github.com/robdimsdale/concourse-pipeline-resource/validator"
	"github.com/robdimsdale/sanitizer"
)

const (
	atcExternalURLEnvKey = "ATC_EXTERNAL_URL"
)

var (
	// version is deliberately left uninitialized so it can be set at compile-time
	version string

	l *logger.Logger
)

func main() {
	if version == "" {
		version = "dev"
	}

	var input concourse.CheckRequest

	logFile, err := ioutil.TempFile("", "concourse-pipeline-resource-check.log")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(logFile, "Concourse Pipeline Resource version: %s\n", version)

	fmt.Fprintf(os.Stderr, "Logging to %s\n", logFile.Name())

	err = json.NewDecoder(os.Stdin).Decode(&input)
	if err != nil {
		fmt.Fprintf(logFile, "Exiting with error: %v\n", err)
		log.Fatalln(err)
	}

	sanitized := concourse.SanitizedSource(input.Source)
	sanitizer := sanitizer.NewSanitizer(sanitized, logFile)

	l = logger.NewLogger(sanitizer)

	err = validator.ValidateCheck(input)
	if err != nil {
		l.Debugf("Exiting with error: %v\n", err)
		log.Fatalln(err)
	}

	if input.Source.Target == "" {
		input.Source.Target = os.Getenv(atcExternalURLEnvKey)
	}

	insecure := false
	if input.Source.Insecure != "" {
		var err error
		insecure, err = strconv.ParseBool(input.Source.Insecure)
		if err != nil {
			log.Fatalln("Invalid value for insecure: %v", input.Source.Insecure)
		}
	}

	teamClients := make(map[string]*http.Client)
	for _, t := range input.Source.Teams {
		teamName := t.Name

		if teamClients[teamName] != nil {
			continue
		}

		token, err := api.LoginWithBasicAuth(
			input.Source.Target,
			t.Name,
			t.Username,
			t.Password,
			insecure,
		)
		if err != nil {
			l.Debugf("Exiting with error: %v\n", err)
			log.Fatalln(err)
		}

		httpClient := api.OAuthHTTPClient(token, insecure)
		teamClients[teamName] = httpClient
	}

	apiClient := api.NewClient(input.Source.Target, teamClients)

	checkCommand := check.NewCheckCommand(version, l, logFile.Name(), apiClient)
	response, err := checkCommand.Run(input)
	if err != nil {
		l.Debugf("Exiting with error: %v\n", err)
		log.Fatalln(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		l.Debugf("Exiting with error: %v\n", err)
		log.Fatalln(err)
	}
}
