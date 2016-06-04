// This file was generated by counterfeiter
package fakes

import (
	"io"
	"os"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/worker"
	"github.com/pivotal-golang/lager"
)

type FakeImageFetcher struct {
	FetchImageStub        func(logger lager.Logger, imageResource atc.ImageResource, cancel <-chan os.Signal, containerID worker.Identifier, containerMetadata worker.Metadata, delegate worker.ImageFetchingDelegate, workerClient worker.Client, tags atc.Tags, resourceTypes atc.ResourceTypes, privileged bool) (worker.Volume, io.ReadCloser, atc.Version, error)
	fetchImageMutex       sync.RWMutex
	fetchImageArgsForCall []struct {
		logger            lager.Logger
		imageResource     atc.ImageResource
		cancel            <-chan os.Signal
		containerID       worker.Identifier
		containerMetadata worker.Metadata
		delegate          worker.ImageFetchingDelegate
		workerClient      worker.Client
		tags              atc.Tags
		resourceTypes     atc.ResourceTypes
		privileged        bool
	}
	fetchImageReturns struct {
		result1 worker.Volume
		result2 io.ReadCloser
		result3 atc.Version
		result4 error
	}
}

func (fake *FakeImageFetcher) FetchImage(logger lager.Logger, imageResource atc.ImageResource, cancel <-chan os.Signal, containerID worker.Identifier, containerMetadata worker.Metadata, delegate worker.ImageFetchingDelegate, workerClient worker.Client, tags atc.Tags, resourceTypes atc.ResourceTypes, privileged bool) (worker.Volume, io.ReadCloser, atc.Version, error) {
	fake.fetchImageMutex.Lock()
	fake.fetchImageArgsForCall = append(fake.fetchImageArgsForCall, struct {
		logger            lager.Logger
		imageResource     atc.ImageResource
		cancel            <-chan os.Signal
		containerID       worker.Identifier
		containerMetadata worker.Metadata
		delegate          worker.ImageFetchingDelegate
		workerClient      worker.Client
		tags              atc.Tags
		resourceTypes     atc.ResourceTypes
		privileged        bool
	}{logger, imageResource, cancel, containerID, containerMetadata, delegate, workerClient, tags, resourceTypes, privileged})
	fake.fetchImageMutex.Unlock()
	if fake.FetchImageStub != nil {
		return fake.FetchImageStub(logger, imageResource, cancel, containerID, containerMetadata, delegate, workerClient, tags, resourceTypes, privileged)
	} else {
		return fake.fetchImageReturns.result1, fake.fetchImageReturns.result2, fake.fetchImageReturns.result3, fake.fetchImageReturns.result4
	}
}

func (fake *FakeImageFetcher) FetchImageCallCount() int {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	return len(fake.fetchImageArgsForCall)
}

func (fake *FakeImageFetcher) FetchImageArgsForCall(i int) (lager.Logger, atc.ImageResource, <-chan os.Signal, worker.Identifier, worker.Metadata, worker.ImageFetchingDelegate, worker.Client, atc.Tags, atc.ResourceTypes, bool) {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	return fake.fetchImageArgsForCall[i].logger, fake.fetchImageArgsForCall[i].imageResource, fake.fetchImageArgsForCall[i].cancel, fake.fetchImageArgsForCall[i].containerID, fake.fetchImageArgsForCall[i].containerMetadata, fake.fetchImageArgsForCall[i].delegate, fake.fetchImageArgsForCall[i].workerClient, fake.fetchImageArgsForCall[i].tags, fake.fetchImageArgsForCall[i].resourceTypes, fake.fetchImageArgsForCall[i].privileged
}

func (fake *FakeImageFetcher) FetchImageReturns(result1 worker.Volume, result2 io.ReadCloser, result3 atc.Version, result4 error) {
	fake.FetchImageStub = nil
	fake.fetchImageReturns = struct {
		result1 worker.Volume
		result2 io.ReadCloser
		result3 atc.Version
		result4 error
	}{result1, result2, result3, result4}
}

var _ worker.ImageFetcher = new(FakeImageFetcher)