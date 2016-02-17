// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/fujitsu-cf/cli/cf/actors/plan_builder"
	"github.com/fujitsu-cf/cli/cf/models"
)

type FakePlanBuilder struct {
	AttachOrgsToPlansStub        func([]models.ServicePlanFields) ([]models.ServicePlanFields, error)
	attachOrgsToPlansMutex       sync.RWMutex
	attachOrgsToPlansArgsForCall []struct {
		arg1 []models.ServicePlanFields
	}
	attachOrgsToPlansReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
	AttachOrgToPlansStub        func([]models.ServicePlanFields, string) ([]models.ServicePlanFields, error)
	attachOrgToPlansMutex       sync.RWMutex
	attachOrgToPlansArgsForCall []struct {
		arg1 []models.ServicePlanFields
		arg2 string
	}
	attachOrgToPlansReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
	GetPlansForServiceForOrgStub        func(string, string) ([]models.ServicePlanFields, error)
	getPlansForServiceForOrgMutex       sync.RWMutex
	getPlansForServiceForOrgArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getPlansForServiceForOrgReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
	GetPlansForServiceWithOrgsStub        func(string) ([]models.ServicePlanFields, error)
	getPlansForServiceWithOrgsMutex       sync.RWMutex
	getPlansForServiceWithOrgsArgsForCall []struct {
		arg1 string
	}
	getPlansForServiceWithOrgsReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
	GetPlansForManyServicesWithOrgsStub        func([]string) ([]models.ServicePlanFields, error)
	getPlansForManyServicesWithOrgsMutex       sync.RWMutex
	getPlansForManyServicesWithOrgsArgsForCall []struct {
		arg1 []string
	}
	getPlansForManyServicesWithOrgsReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
	GetPlansForServiceStub        func(string) ([]models.ServicePlanFields, error)
	getPlansForServiceMutex       sync.RWMutex
	getPlansForServiceArgsForCall []struct {
		arg1 string
	}
	getPlansForServiceReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
	GetPlansVisibleToOrgStub        func(string) ([]models.ServicePlanFields, error)
	getPlansVisibleToOrgMutex       sync.RWMutex
	getPlansVisibleToOrgArgsForCall []struct {
		arg1 string
	}
	getPlansVisibleToOrgReturns struct {
		result1 []models.ServicePlanFields
		result2 error
	}
}

func (fake *FakePlanBuilder) AttachOrgsToPlans(arg1 []models.ServicePlanFields) ([]models.ServicePlanFields, error) {
	fake.attachOrgsToPlansMutex.Lock()
	fake.attachOrgsToPlansArgsForCall = append(fake.attachOrgsToPlansArgsForCall, struct {
		arg1 []models.ServicePlanFields
	}{arg1})
	fake.attachOrgsToPlansMutex.Unlock()
	if fake.AttachOrgsToPlansStub != nil {
		return fake.AttachOrgsToPlansStub(arg1)
	} else {
		return fake.attachOrgsToPlansReturns.result1, fake.attachOrgsToPlansReturns.result2
	}
}

func (fake *FakePlanBuilder) AttachOrgsToPlansCallCount() int {
	fake.attachOrgsToPlansMutex.RLock()
	defer fake.attachOrgsToPlansMutex.RUnlock()
	return len(fake.attachOrgsToPlansArgsForCall)
}

func (fake *FakePlanBuilder) AttachOrgsToPlansArgsForCall(i int) []models.ServicePlanFields {
	fake.attachOrgsToPlansMutex.RLock()
	defer fake.attachOrgsToPlansMutex.RUnlock()
	return fake.attachOrgsToPlansArgsForCall[i].arg1
}

func (fake *FakePlanBuilder) AttachOrgsToPlansReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.AttachOrgsToPlansStub = nil
	fake.attachOrgsToPlansReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

func (fake *FakePlanBuilder) AttachOrgToPlans(arg1 []models.ServicePlanFields, arg2 string) ([]models.ServicePlanFields, error) {
	fake.attachOrgToPlansMutex.Lock()
	fake.attachOrgToPlansArgsForCall = append(fake.attachOrgToPlansArgsForCall, struct {
		arg1 []models.ServicePlanFields
		arg2 string
	}{arg1, arg2})
	fake.attachOrgToPlansMutex.Unlock()
	if fake.AttachOrgToPlansStub != nil {
		return fake.AttachOrgToPlansStub(arg1, arg2)
	} else {
		return fake.attachOrgToPlansReturns.result1, fake.attachOrgToPlansReturns.result2
	}
}

func (fake *FakePlanBuilder) AttachOrgToPlansCallCount() int {
	fake.attachOrgToPlansMutex.RLock()
	defer fake.attachOrgToPlansMutex.RUnlock()
	return len(fake.attachOrgToPlansArgsForCall)
}

func (fake *FakePlanBuilder) AttachOrgToPlansArgsForCall(i int) ([]models.ServicePlanFields, string) {
	fake.attachOrgToPlansMutex.RLock()
	defer fake.attachOrgToPlansMutex.RUnlock()
	return fake.attachOrgToPlansArgsForCall[i].arg1, fake.attachOrgToPlansArgsForCall[i].arg2
}

func (fake *FakePlanBuilder) AttachOrgToPlansReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.AttachOrgToPlansStub = nil
	fake.attachOrgToPlansReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

func (fake *FakePlanBuilder) GetPlansForServiceForOrg(arg1 string, arg2 string) ([]models.ServicePlanFields, error) {
	fake.getPlansForServiceForOrgMutex.Lock()
	fake.getPlansForServiceForOrgArgsForCall = append(fake.getPlansForServiceForOrgArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.getPlansForServiceForOrgMutex.Unlock()
	if fake.GetPlansForServiceForOrgStub != nil {
		return fake.GetPlansForServiceForOrgStub(arg1, arg2)
	} else {
		return fake.getPlansForServiceForOrgReturns.result1, fake.getPlansForServiceForOrgReturns.result2
	}
}

func (fake *FakePlanBuilder) GetPlansForServiceForOrgCallCount() int {
	fake.getPlansForServiceForOrgMutex.RLock()
	defer fake.getPlansForServiceForOrgMutex.RUnlock()
	return len(fake.getPlansForServiceForOrgArgsForCall)
}

func (fake *FakePlanBuilder) GetPlansForServiceForOrgArgsForCall(i int) (string, string) {
	fake.getPlansForServiceForOrgMutex.RLock()
	defer fake.getPlansForServiceForOrgMutex.RUnlock()
	return fake.getPlansForServiceForOrgArgsForCall[i].arg1, fake.getPlansForServiceForOrgArgsForCall[i].arg2
}

func (fake *FakePlanBuilder) GetPlansForServiceForOrgReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.GetPlansForServiceForOrgStub = nil
	fake.getPlansForServiceForOrgReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

func (fake *FakePlanBuilder) GetPlansForServiceWithOrgs(arg1 string) ([]models.ServicePlanFields, error) {
	fake.getPlansForServiceWithOrgsMutex.Lock()
	fake.getPlansForServiceWithOrgsArgsForCall = append(fake.getPlansForServiceWithOrgsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getPlansForServiceWithOrgsMutex.Unlock()
	if fake.GetPlansForServiceWithOrgsStub != nil {
		return fake.GetPlansForServiceWithOrgsStub(arg1)
	} else {
		return fake.getPlansForServiceWithOrgsReturns.result1, fake.getPlansForServiceWithOrgsReturns.result2
	}
}

func (fake *FakePlanBuilder) GetPlansForServiceWithOrgsCallCount() int {
	fake.getPlansForServiceWithOrgsMutex.RLock()
	defer fake.getPlansForServiceWithOrgsMutex.RUnlock()
	return len(fake.getPlansForServiceWithOrgsArgsForCall)
}

func (fake *FakePlanBuilder) GetPlansForServiceWithOrgsArgsForCall(i int) string {
	fake.getPlansForServiceWithOrgsMutex.RLock()
	defer fake.getPlansForServiceWithOrgsMutex.RUnlock()
	return fake.getPlansForServiceWithOrgsArgsForCall[i].arg1
}

func (fake *FakePlanBuilder) GetPlansForServiceWithOrgsReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.GetPlansForServiceWithOrgsStub = nil
	fake.getPlansForServiceWithOrgsReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

func (fake *FakePlanBuilder) GetPlansForManyServicesWithOrgs(arg1 []string) ([]models.ServicePlanFields, error) {
	fake.getPlansForManyServicesWithOrgsMutex.Lock()
	fake.getPlansForManyServicesWithOrgsArgsForCall = append(fake.getPlansForManyServicesWithOrgsArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.getPlansForManyServicesWithOrgsMutex.Unlock()
	if fake.GetPlansForManyServicesWithOrgsStub != nil {
		return fake.GetPlansForManyServicesWithOrgsStub(arg1)
	} else {
		return fake.getPlansForManyServicesWithOrgsReturns.result1, fake.getPlansForManyServicesWithOrgsReturns.result2
	}
}

func (fake *FakePlanBuilder) GetPlansForManyServicesWithOrgsCallCount() int {
	fake.getPlansForManyServicesWithOrgsMutex.RLock()
	defer fake.getPlansForManyServicesWithOrgsMutex.RUnlock()
	return len(fake.getPlansForManyServicesWithOrgsArgsForCall)
}

func (fake *FakePlanBuilder) GetPlansForManyServicesWithOrgsArgsForCall(i int) []string {
	fake.getPlansForManyServicesWithOrgsMutex.RLock()
	defer fake.getPlansForManyServicesWithOrgsMutex.RUnlock()
	return fake.getPlansForManyServicesWithOrgsArgsForCall[i].arg1
}

func (fake *FakePlanBuilder) GetPlansForManyServicesWithOrgsReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.GetPlansForManyServicesWithOrgsStub = nil
	fake.getPlansForManyServicesWithOrgsReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

func (fake *FakePlanBuilder) GetPlansForService(arg1 string) ([]models.ServicePlanFields, error) {
	fake.getPlansForServiceMutex.Lock()
	fake.getPlansForServiceArgsForCall = append(fake.getPlansForServiceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getPlansForServiceMutex.Unlock()
	if fake.GetPlansForServiceStub != nil {
		return fake.GetPlansForServiceStub(arg1)
	} else {
		return fake.getPlansForServiceReturns.result1, fake.getPlansForServiceReturns.result2
	}
}

func (fake *FakePlanBuilder) GetPlansForServiceCallCount() int {
	fake.getPlansForServiceMutex.RLock()
	defer fake.getPlansForServiceMutex.RUnlock()
	return len(fake.getPlansForServiceArgsForCall)
}

func (fake *FakePlanBuilder) GetPlansForServiceArgsForCall(i int) string {
	fake.getPlansForServiceMutex.RLock()
	defer fake.getPlansForServiceMutex.RUnlock()
	return fake.getPlansForServiceArgsForCall[i].arg1
}

func (fake *FakePlanBuilder) GetPlansForServiceReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.GetPlansForServiceStub = nil
	fake.getPlansForServiceReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

func (fake *FakePlanBuilder) GetPlansVisibleToOrg(arg1 string) ([]models.ServicePlanFields, error) {
	fake.getPlansVisibleToOrgMutex.Lock()
	fake.getPlansVisibleToOrgArgsForCall = append(fake.getPlansVisibleToOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getPlansVisibleToOrgMutex.Unlock()
	if fake.GetPlansVisibleToOrgStub != nil {
		return fake.GetPlansVisibleToOrgStub(arg1)
	} else {
		return fake.getPlansVisibleToOrgReturns.result1, fake.getPlansVisibleToOrgReturns.result2
	}
}

func (fake *FakePlanBuilder) GetPlansVisibleToOrgCallCount() int {
	fake.getPlansVisibleToOrgMutex.RLock()
	defer fake.getPlansVisibleToOrgMutex.RUnlock()
	return len(fake.getPlansVisibleToOrgArgsForCall)
}

func (fake *FakePlanBuilder) GetPlansVisibleToOrgArgsForCall(i int) string {
	fake.getPlansVisibleToOrgMutex.RLock()
	defer fake.getPlansVisibleToOrgMutex.RUnlock()
	return fake.getPlansVisibleToOrgArgsForCall[i].arg1
}

func (fake *FakePlanBuilder) GetPlansVisibleToOrgReturns(result1 []models.ServicePlanFields, result2 error) {
	fake.GetPlansVisibleToOrgStub = nil
	fake.getPlansVisibleToOrgReturns = struct {
		result1 []models.ServicePlanFields
		result2 error
	}{result1, result2}
}

var _ plan_builder.PlanBuilder = new(FakePlanBuilder)
