// This file was generated by counterfeiter
package workerfakes

import (
	"sync"

	"github.com/concourse/atc/worker"
	"github.com/pivotal-golang/lager"
)

type FakeVolumeClient struct {
	FindVolumeStub        func(lager.Logger, worker.VolumeSpec) (worker.Volume, bool, error)
	findVolumeMutex       sync.RWMutex
	findVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
	}
	findVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	CreateVolumeStub        func(logger lager.Logger, vs worker.VolumeSpec, teamID int) (worker.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		logger lager.Logger
		vs     worker.VolumeSpec
		teamID int
	}
	createVolumeReturns struct {
		result1 worker.Volume
		result2 error
	}
	ListVolumesStub        func(lager.Logger, worker.VolumeProperties) ([]worker.Volume, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeProperties
	}
	listVolumesReturns struct {
		result1 []worker.Volume
		result2 error
	}
	LookupVolumeStub        func(lager.Logger, string) (worker.Volume, bool, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeClient) FindVolume(arg1 lager.Logger, arg2 worker.VolumeSpec) (worker.Volume, bool, error) {
	fake.findVolumeMutex.Lock()
	fake.findVolumeArgsForCall = append(fake.findVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
	}{arg1, arg2})
	fake.recordInvocation("FindVolume", []interface{}{arg1, arg2})
	fake.findVolumeMutex.Unlock()
	if fake.FindVolumeStub != nil {
		return fake.FindVolumeStub(arg1, arg2)
	} else {
		return fake.findVolumeReturns.result1, fake.findVolumeReturns.result2, fake.findVolumeReturns.result3
	}
}

func (fake *FakeVolumeClient) FindVolumeCallCount() int {
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	return len(fake.findVolumeArgsForCall)
}

func (fake *FakeVolumeClient) FindVolumeArgsForCall(i int) (lager.Logger, worker.VolumeSpec) {
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	return fake.findVolumeArgsForCall[i].arg1, fake.findVolumeArgsForCall[i].arg2
}

func (fake *FakeVolumeClient) FindVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.FindVolumeStub = nil
	fake.findVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) CreateVolume(logger lager.Logger, vs worker.VolumeSpec, teamID int) (worker.Volume, error) {
	fake.createVolumeMutex.Lock()
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		logger lager.Logger
		vs     worker.VolumeSpec
		teamID int
	}{logger, vs, teamID})
	fake.recordInvocation("CreateVolume", []interface{}{logger, vs, teamID})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(logger, vs, teamID)
	} else {
		return fake.createVolumeReturns.result1, fake.createVolumeReturns.result2
	}
}

func (fake *FakeVolumeClient) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeVolumeClient) CreateVolumeArgsForCall(i int) (lager.Logger, worker.VolumeSpec, int) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return fake.createVolumeArgsForCall[i].logger, fake.createVolumeArgsForCall[i].vs, fake.createVolumeArgsForCall[i].teamID
}

func (fake *FakeVolumeClient) CreateVolumeReturns(result1 worker.Volume, result2 error) {
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) ListVolumes(arg1 lager.Logger, arg2 worker.VolumeProperties) ([]worker.Volume, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeProperties
	}{arg1, arg2})
	fake.recordInvocation("ListVolumes", []interface{}{arg1, arg2})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub(arg1, arg2)
	} else {
		return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
	}
}

func (fake *FakeVolumeClient) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeVolumeClient) ListVolumesArgsForCall(i int) (lager.Logger, worker.VolumeProperties) {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return fake.listVolumesArgsForCall[i].arg1, fake.listVolumesArgsForCall[i].arg2
}

func (fake *FakeVolumeClient) ListVolumesReturns(result1 []worker.Volume, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) LookupVolume(arg1 lager.Logger, arg2 string) (worker.Volume, bool, error) {
	fake.lookupVolumeMutex.Lock()
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupVolume", []interface{}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(arg1, arg2)
	} else {
		return fake.lookupVolumeReturns.result1, fake.lookupVolumeReturns.result2, fake.lookupVolumeReturns.result3
	}
}

func (fake *FakeVolumeClient) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeVolumeClient) LookupVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.lookupVolumeArgsForCall[i].arg1, fake.lookupVolumeArgsForCall[i].arg2
}

func (fake *FakeVolumeClient) LookupVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVolumeClient) recordInvocation(key string, args []interface{}) {
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

var _ worker.VolumeClient = new(FakeVolumeClient)
