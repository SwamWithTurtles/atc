// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/scheduler/factory"
)

type FakeBuildFactory struct {
	CreateStub        func(atc.JobConfig, atc.ResourceConfigs, atc.ResourceTypes, []db.BuildInput) (atc.Plan, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 atc.JobConfig
		arg2 atc.ResourceConfigs
		arg3 atc.ResourceTypes
		arg4 []db.BuildInput
	}
	createReturns struct {
		result1 atc.Plan
		result2 error
	}
}

func (fake *FakeBuildFactory) Create(arg1 atc.JobConfig, arg2 atc.ResourceConfigs, arg3 atc.ResourceTypes, arg4 []db.BuildInput) (atc.Plan, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 atc.JobConfig
		arg2 atc.ResourceConfigs
		arg3 atc.ResourceTypes
		arg4 []db.BuildInput
	}{arg1, arg2, arg3, arg4})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeBuildFactory) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBuildFactory) CreateArgsForCall(i int) (atc.JobConfig, atc.ResourceConfigs, atc.ResourceTypes, []db.BuildInput) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2, fake.createArgsForCall[i].arg3, fake.createArgsForCall[i].arg4
}

func (fake *FakeBuildFactory) CreateReturns(result1 atc.Plan, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 atc.Plan
		result2 error
	}{result1, result2}
}

var _ factory.BuildFactory = new(FakeBuildFactory)
