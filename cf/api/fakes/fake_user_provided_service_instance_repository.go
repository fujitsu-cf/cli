// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/fujitsu-cf/cli/cf/api"
	"github.com/fujitsu-cf/cli/cf/models"
)

type FakeUserProvidedServiceInstanceRepository struct {
	CreateStub        func(name, drainUrl string, routeServiceUrl string, params map[string]interface{}) (apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		name            string
		drainUrl        string
		routeServiceUrl string
		params          map[string]interface{}
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(serviceInstanceFields models.ServiceInstanceFields) (apiErr error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		serviceInstanceFields models.ServiceInstanceFields
	}
	updateReturns struct {
		result1 error
	}
	GetSummariesStub        func() (models.UserProvidedServiceSummary, error)
	getSummariesMutex       sync.RWMutex
	getSummariesArgsForCall []struct{}
	getSummariesReturns     struct {
		result1 models.UserProvidedServiceSummary
		result2 error
	}
}

func (fake *FakeUserProvidedServiceInstanceRepository) Create(name string, drainUrl string, routeServiceUrl string, params map[string]interface{}) (apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		name            string
		drainUrl        string
		routeServiceUrl string
		params          map[string]interface{}
	}{name, drainUrl, routeServiceUrl, params})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(name, drainUrl, routeServiceUrl, params)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeUserProvidedServiceInstanceRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeUserProvidedServiceInstanceRepository) CreateArgsForCall(i int) (string, string, string, map[string]interface{}) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].name, fake.createArgsForCall[i].drainUrl, fake.createArgsForCall[i].routeServiceUrl, fake.createArgsForCall[i].params
}

func (fake *FakeUserProvidedServiceInstanceRepository) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserProvidedServiceInstanceRepository) Update(serviceInstanceFields models.ServiceInstanceFields) (apiErr error) {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		serviceInstanceFields models.ServiceInstanceFields
	}{serviceInstanceFields})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(serviceInstanceFields)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeUserProvidedServiceInstanceRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeUserProvidedServiceInstanceRepository) UpdateArgsForCall(i int) models.ServiceInstanceFields {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].serviceInstanceFields
}

func (fake *FakeUserProvidedServiceInstanceRepository) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserProvidedServiceInstanceRepository) GetSummaries() (models.UserProvidedServiceSummary, error) {
	fake.getSummariesMutex.Lock()
	fake.getSummariesArgsForCall = append(fake.getSummariesArgsForCall, struct{}{})
	fake.getSummariesMutex.Unlock()
	if fake.GetSummariesStub != nil {
		return fake.GetSummariesStub()
	} else {
		return fake.getSummariesReturns.result1, fake.getSummariesReturns.result2
	}
}

func (fake *FakeUserProvidedServiceInstanceRepository) GetSummariesCallCount() int {
	fake.getSummariesMutex.RLock()
	defer fake.getSummariesMutex.RUnlock()
	return len(fake.getSummariesArgsForCall)
}

func (fake *FakeUserProvidedServiceInstanceRepository) GetSummariesReturns(result1 models.UserProvidedServiceSummary, result2 error) {
	fake.GetSummariesStub = nil
	fake.getSummariesReturns = struct {
		result1 models.UserProvidedServiceSummary
		result2 error
	}{result1, result2}
}

var _ api.UserProvidedServiceInstanceRepository = new(FakeUserProvidedServiceInstanceRepository)
