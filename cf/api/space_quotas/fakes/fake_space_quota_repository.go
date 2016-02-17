// This file was generated by counterfeiter
package fakes

import (
	. "github.com/fujitsu-cf/cli/cf/api/space_quotas"
	"github.com/fujitsu-cf/cli/cf/models"
	"sync"
)

type FakeSpaceQuotaRepository struct {
	FindByNameStub        func(name string) (quota models.SpaceQuota, apiErr error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		arg1 string
	}
	findByNameReturns struct {
		result1 models.SpaceQuota
		result2 error
	}
	FindByOrgStub        func(guid string) (quota []models.SpaceQuota, apiErr error)
	findByOrgMutex       sync.RWMutex
	findByOrgArgsForCall []struct {
		arg1 string
	}
	findByOrgReturns struct {
		result1 []models.SpaceQuota
		result2 error
	}
	FindByGuidStub        func(guid string) (quota models.SpaceQuota, apiErr error)
	findByGuidMutex       sync.RWMutex
	findByGuidArgsForCall []struct {
		arg1 string
	}
	findByGuidReturns struct {
		result1 models.SpaceQuota
		result2 error
	}
	AssociateSpaceWithQuotaStub        func(spaceGuid string, quotaGuid string) error
	associateSpaceWithQuotaMutex       sync.RWMutex
	associateSpaceWithQuotaArgsForCall []struct {
		arg1 string
		arg2 string
	}
	associateSpaceWithQuotaReturns struct {
		result1 error
	}
	UnassignQuotaFromSpaceStub        func(spaceGuid string, quotaGuid string) error
	unassignQuotaFromSpaceMutex       sync.RWMutex
	unassignQuotaFromSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	unassignQuotaFromSpaceReturns struct {
		result1 error
	}
	CreateStub        func(quota models.SpaceQuota) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 models.SpaceQuota
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(quota models.SpaceQuota) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 models.SpaceQuota
	}
	updateReturns struct {
		result1 error
	}
	DeleteStub        func(quotaGuid string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeSpaceQuotaRepository) FindByName(arg1 string) (quota models.SpaceQuota, apiErr error) {
	fake.findByNameMutex.Lock()
	defer fake.findByNameMutex.Unlock()
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.FindByNameStub != nil {
		return fake.FindByNameStub(arg1)
	} else {
		return fake.findByNameReturns.result1, fake.findByNameReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return fake.findByNameArgsForCall[i].arg1
}

func (fake *FakeSpaceQuotaRepository) FindByNameReturns(result1 models.SpaceQuota, result2 error) {
	fake.findByNameReturns = struct {
		result1 models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) FindByOrg(arg1 string) (quota []models.SpaceQuota, apiErr error) {
	fake.findByOrgMutex.Lock()
	defer fake.findByOrgMutex.Unlock()
	fake.findByOrgArgsForCall = append(fake.findByOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.FindByOrgStub != nil {
		return fake.FindByOrgStub(arg1)
	} else {
		return fake.findByOrgReturns.result1, fake.findByOrgReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByOrgCallCount() int {
	fake.findByOrgMutex.RLock()
	defer fake.findByOrgMutex.RUnlock()
	return len(fake.findByOrgArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByOrgArgsForCall(i int) string {
	fake.findByOrgMutex.RLock()
	defer fake.findByOrgMutex.RUnlock()
	return fake.findByOrgArgsForCall[i].arg1
}

func (fake *FakeSpaceQuotaRepository) FindByOrgReturns(result1 []models.SpaceQuota, result2 error) {
	fake.findByOrgReturns = struct {
		result1 []models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) FindByGuid(arg1 string) (quota models.SpaceQuota, apiErr error) {
	fake.findByGuidMutex.Lock()
	defer fake.findByGuidMutex.Unlock()
	fake.findByGuidArgsForCall = append(fake.findByGuidArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.FindByGuidStub != nil {
		return fake.FindByGuidStub(arg1)
	} else {
		return fake.findByGuidReturns.result1, fake.findByGuidReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByGuidCallCount() int {
	fake.findByGuidMutex.RLock()
	defer fake.findByGuidMutex.RUnlock()
	return len(fake.findByGuidArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByGuidArgsForCall(i int) string {
	fake.findByGuidMutex.RLock()
	defer fake.findByGuidMutex.RUnlock()
	return fake.findByGuidArgsForCall[i].arg1
}

func (fake *FakeSpaceQuotaRepository) FindByGuidReturns(result1 models.SpaceQuota, result2 error) {
	fake.findByGuidReturns = struct {
		result1 models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuota(arg1 string, arg2 string) error {
	fake.associateSpaceWithQuotaMutex.Lock()
	defer fake.associateSpaceWithQuotaMutex.Unlock()
	fake.associateSpaceWithQuotaArgsForCall = append(fake.associateSpaceWithQuotaArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	if fake.AssociateSpaceWithQuotaStub != nil {
		return fake.AssociateSpaceWithQuotaStub(arg1, arg2)
	} else {
		return fake.associateSpaceWithQuotaReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuotaCallCount() int {
	fake.associateSpaceWithQuotaMutex.RLock()
	defer fake.associateSpaceWithQuotaMutex.RUnlock()
	return len(fake.associateSpaceWithQuotaArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuotaArgsForCall(i int) (string, string) {
	fake.associateSpaceWithQuotaMutex.RLock()
	defer fake.associateSpaceWithQuotaMutex.RUnlock()
	return fake.associateSpaceWithQuotaArgsForCall[i].arg1, fake.associateSpaceWithQuotaArgsForCall[i].arg2
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuotaReturns(result1 error) {
	fake.associateSpaceWithQuotaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpace(arg1 string, arg2 string) error {
	fake.unassignQuotaFromSpaceMutex.Lock()
	defer fake.unassignQuotaFromSpaceMutex.Unlock()
	fake.unassignQuotaFromSpaceArgsForCall = append(fake.unassignQuotaFromSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	if fake.UnassignQuotaFromSpaceStub != nil {
		return fake.UnassignQuotaFromSpaceStub(arg1, arg2)
	} else {
		return fake.unassignQuotaFromSpaceReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpaceCallCount() int {
	fake.unassignQuotaFromSpaceMutex.RLock()
	defer fake.unassignQuotaFromSpaceMutex.RUnlock()
	return len(fake.unassignQuotaFromSpaceArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpaceArgsForCall(i int) (string, string) {
	fake.unassignQuotaFromSpaceMutex.RLock()
	defer fake.unassignQuotaFromSpaceMutex.RUnlock()
	return fake.unassignQuotaFromSpaceArgsForCall[i].arg1, fake.unassignQuotaFromSpaceArgsForCall[i].arg2
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpaceReturns(result1 error) {
	fake.unassignQuotaFromSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Create(arg1 models.SpaceQuota) error {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 models.SpaceQuota
	}{arg1})
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) CreateArgsForCall(i int) models.SpaceQuota {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *FakeSpaceQuotaRepository) CreateReturns(result1 error) {
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Update(arg1 models.SpaceQuota) error {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 models.SpaceQuota
	}{arg1})
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) UpdateArgsForCall(i int) models.SpaceQuota {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1
}

func (fake *FakeSpaceQuotaRepository) UpdateReturns(result1 error) {
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1
}

func (fake *FakeSpaceQuotaRepository) DeleteReturns(result1 error) {
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ SpaceQuotaRepository = new(FakeSpaceQuotaRepository)
