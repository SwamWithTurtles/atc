// This file was generated by counterfeiter
package resourcefakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeResourceFactory struct {
	NewPutResourceStub        func(logger lager.Logger, signals <-chan os.Signal, buildID int, planID atc.PlanID, metadata dbng.ContainerMetadata, containerSpec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate) (resource.Resource, error)
	newPutResourceMutex       sync.RWMutex
	newPutResourceArgsForCall []struct {
		logger                lager.Logger
		signals               <-chan os.Signal
		buildID               int
		planID                atc.PlanID
		metadata              dbng.ContainerMetadata
		containerSpec         worker.ContainerSpec
		resourceTypes         atc.VersionedResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
	}
	newPutResourceReturns struct {
		result1 resource.Resource
		result2 error
	}
	newPutResourceReturnsOnCall map[int]struct {
		result1 resource.Resource
		result2 error
	}
	NewCheckResourceStub        func(logger lager.Logger, signals <-chan os.Signal, resourceUser dbng.ResourceUser, resourceType string, resourceSource atc.Source, metadata dbng.ContainerMetadata, resourceSpec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate) (resource.Resource, error)
	newCheckResourceMutex       sync.RWMutex
	newCheckResourceArgsForCall []struct {
		logger                lager.Logger
		signals               <-chan os.Signal
		resourceUser          dbng.ResourceUser
		resourceType          string
		resourceSource        atc.Source
		metadata              dbng.ContainerMetadata
		resourceSpec          worker.ContainerSpec
		resourceTypes         atc.VersionedResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
	}
	newCheckResourceReturns struct {
		result1 resource.Resource
		result2 error
	}
	newCheckResourceReturnsOnCall map[int]struct {
		result1 resource.Resource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceFactory) NewPutResource(logger lager.Logger, signals <-chan os.Signal, buildID int, planID atc.PlanID, metadata dbng.ContainerMetadata, containerSpec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate) (resource.Resource, error) {
	fake.newPutResourceMutex.Lock()
	ret, specificReturn := fake.newPutResourceReturnsOnCall[len(fake.newPutResourceArgsForCall)]
	fake.newPutResourceArgsForCall = append(fake.newPutResourceArgsForCall, struct {
		logger                lager.Logger
		signals               <-chan os.Signal
		buildID               int
		planID                atc.PlanID
		metadata              dbng.ContainerMetadata
		containerSpec         worker.ContainerSpec
		resourceTypes         atc.VersionedResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
	}{logger, signals, buildID, planID, metadata, containerSpec, resourceTypes, imageFetchingDelegate})
	fake.recordInvocation("NewPutResource", []interface{}{logger, signals, buildID, planID, metadata, containerSpec, resourceTypes, imageFetchingDelegate})
	fake.newPutResourceMutex.Unlock()
	if fake.NewPutResourceStub != nil {
		return fake.NewPutResourceStub(logger, signals, buildID, planID, metadata, containerSpec, resourceTypes, imageFetchingDelegate)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newPutResourceReturns.result1, fake.newPutResourceReturns.result2
}

func (fake *FakeResourceFactory) NewPutResourceCallCount() int {
	fake.newPutResourceMutex.RLock()
	defer fake.newPutResourceMutex.RUnlock()
	return len(fake.newPutResourceArgsForCall)
}

func (fake *FakeResourceFactory) NewPutResourceArgsForCall(i int) (lager.Logger, <-chan os.Signal, int, atc.PlanID, dbng.ContainerMetadata, worker.ContainerSpec, atc.VersionedResourceTypes, worker.ImageFetchingDelegate) {
	fake.newPutResourceMutex.RLock()
	defer fake.newPutResourceMutex.RUnlock()
	return fake.newPutResourceArgsForCall[i].logger, fake.newPutResourceArgsForCall[i].signals, fake.newPutResourceArgsForCall[i].buildID, fake.newPutResourceArgsForCall[i].planID, fake.newPutResourceArgsForCall[i].metadata, fake.newPutResourceArgsForCall[i].containerSpec, fake.newPutResourceArgsForCall[i].resourceTypes, fake.newPutResourceArgsForCall[i].imageFetchingDelegate
}

func (fake *FakeResourceFactory) NewPutResourceReturns(result1 resource.Resource, result2 error) {
	fake.NewPutResourceStub = nil
	fake.newPutResourceReturns = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) NewPutResourceReturnsOnCall(i int, result1 resource.Resource, result2 error) {
	fake.NewPutResourceStub = nil
	if fake.newPutResourceReturnsOnCall == nil {
		fake.newPutResourceReturnsOnCall = make(map[int]struct {
			result1 resource.Resource
			result2 error
		})
	}
	fake.newPutResourceReturnsOnCall[i] = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) NewCheckResource(logger lager.Logger, signals <-chan os.Signal, resourceUser dbng.ResourceUser, resourceType string, resourceSource atc.Source, metadata dbng.ContainerMetadata, resourceSpec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes, imageFetchingDelegate worker.ImageFetchingDelegate) (resource.Resource, error) {
	fake.newCheckResourceMutex.Lock()
	ret, specificReturn := fake.newCheckResourceReturnsOnCall[len(fake.newCheckResourceArgsForCall)]
	fake.newCheckResourceArgsForCall = append(fake.newCheckResourceArgsForCall, struct {
		logger                lager.Logger
		signals               <-chan os.Signal
		resourceUser          dbng.ResourceUser
		resourceType          string
		resourceSource        atc.Source
		metadata              dbng.ContainerMetadata
		resourceSpec          worker.ContainerSpec
		resourceTypes         atc.VersionedResourceTypes
		imageFetchingDelegate worker.ImageFetchingDelegate
	}{logger, signals, resourceUser, resourceType, resourceSource, metadata, resourceSpec, resourceTypes, imageFetchingDelegate})
	fake.recordInvocation("NewCheckResource", []interface{}{logger, signals, resourceUser, resourceType, resourceSource, metadata, resourceSpec, resourceTypes, imageFetchingDelegate})
	fake.newCheckResourceMutex.Unlock()
	if fake.NewCheckResourceStub != nil {
		return fake.NewCheckResourceStub(logger, signals, resourceUser, resourceType, resourceSource, metadata, resourceSpec, resourceTypes, imageFetchingDelegate)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newCheckResourceReturns.result1, fake.newCheckResourceReturns.result2
}

func (fake *FakeResourceFactory) NewCheckResourceCallCount() int {
	fake.newCheckResourceMutex.RLock()
	defer fake.newCheckResourceMutex.RUnlock()
	return len(fake.newCheckResourceArgsForCall)
}

func (fake *FakeResourceFactory) NewCheckResourceArgsForCall(i int) (lager.Logger, <-chan os.Signal, dbng.ResourceUser, string, atc.Source, dbng.ContainerMetadata, worker.ContainerSpec, atc.VersionedResourceTypes, worker.ImageFetchingDelegate) {
	fake.newCheckResourceMutex.RLock()
	defer fake.newCheckResourceMutex.RUnlock()
	return fake.newCheckResourceArgsForCall[i].logger, fake.newCheckResourceArgsForCall[i].signals, fake.newCheckResourceArgsForCall[i].resourceUser, fake.newCheckResourceArgsForCall[i].resourceType, fake.newCheckResourceArgsForCall[i].resourceSource, fake.newCheckResourceArgsForCall[i].metadata, fake.newCheckResourceArgsForCall[i].resourceSpec, fake.newCheckResourceArgsForCall[i].resourceTypes, fake.newCheckResourceArgsForCall[i].imageFetchingDelegate
}

func (fake *FakeResourceFactory) NewCheckResourceReturns(result1 resource.Resource, result2 error) {
	fake.NewCheckResourceStub = nil
	fake.newCheckResourceReturns = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) NewCheckResourceReturnsOnCall(i int, result1 resource.Resource, result2 error) {
	fake.NewCheckResourceStub = nil
	if fake.newCheckResourceReturnsOnCall == nil {
		fake.newCheckResourceReturnsOnCall = make(map[int]struct {
			result1 resource.Resource
			result2 error
		})
	}
	fake.newCheckResourceReturnsOnCall[i] = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newPutResourceMutex.RLock()
	defer fake.newPutResourceMutex.RUnlock()
	fake.newCheckResourceMutex.RLock()
	defer fake.newCheckResourceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResourceFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ resource.ResourceFactory = new(FakeResourceFactory)