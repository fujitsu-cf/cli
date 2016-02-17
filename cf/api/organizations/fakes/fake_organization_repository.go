// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/fujitsu-cf/cli/cf/api/organizations"
	"github.com/fujitsu-cf/cli/cf/models"
)

type FakeOrganizationRepository struct {
	ListOrgsStub        func(limit int) ([]models.Organization, error)
	listOrgsMutex       sync.RWMutex
	listOrgsArgsForCall []struct {
		limit int
	}
	listOrgsReturns struct {
		result1 []models.Organization
		result2 error
	}
	GetManyOrgsByGuidStub        func(orgGuids []string) (orgs []models.Organization, apiErr error)
	getManyOrgsByGuidMutex       sync.RWMutex
	getManyOrgsByGuidArgsForCall []struct {
		orgGuids []string
	}
	getManyOrgsByGuidReturns struct {
		result1 []models.Organization
		result2 error
	}
	FindByNameStub        func(name string) (org models.Organization, apiErr error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		name string
	}
	findByNameReturns struct {
		result1 models.Organization
		result2 error
	}
	CreateStub        func(org models.Organization) (apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		org models.Organization
	}
	createReturns struct {
		result1 error
	}
	RenameStub        func(orgGuid string, name string) (apiErr error)
	renameMutex       sync.RWMutex
	renameArgsForCall []struct {
		orgGuid string
		name    string
	}
	renameReturns struct {
		result1 error
	}
	DeleteStub        func(orgGuid string) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		orgGuid string
	}
	deleteReturns struct {
		result1 error
	}
	SharePrivateDomainStub        func(orgGuid string, domainGuid string) (apiErr error)
	sharePrivateDomainMutex       sync.RWMutex
	sharePrivateDomainArgsForCall []struct {
		orgGuid    string
		domainGuid string
	}
	sharePrivateDomainReturns struct {
		result1 error
	}
	UnsharePrivateDomainStub        func(orgGuid string, domainGuid string) (apiErr error)
	unsharePrivateDomainMutex       sync.RWMutex
	unsharePrivateDomainArgsForCall []struct {
		orgGuid    string
		domainGuid string
	}
	unsharePrivateDomainReturns struct {
		result1 error
	}
}

func (fake *FakeOrganizationRepository) ListOrgs(limit int) ([]models.Organization, error) {
	fake.listOrgsMutex.Lock()
	fake.listOrgsArgsForCall = append(fake.listOrgsArgsForCall, struct {
		limit int
	}{limit})
	fake.listOrgsMutex.Unlock()
	if fake.ListOrgsStub != nil {
		return fake.ListOrgsStub(limit)
	} else {
		return fake.listOrgsReturns.result1, fake.listOrgsReturns.result2
	}
}

func (fake *FakeOrganizationRepository) ListOrgsCallCount() int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return len(fake.listOrgsArgsForCall)
}

func (fake *FakeOrganizationRepository) ListOrgsArgsForCall(i int) int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return fake.listOrgsArgsForCall[i].limit
}

func (fake *FakeOrganizationRepository) ListOrgsReturns(result1 []models.Organization, result2 error) {
	fake.ListOrgsStub = nil
	fake.listOrgsReturns = struct {
		result1 []models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGuid(orgGuids []string) (orgs []models.Organization, apiErr error) {
	fake.getManyOrgsByGuidMutex.Lock()
	fake.getManyOrgsByGuidArgsForCall = append(fake.getManyOrgsByGuidArgsForCall, struct {
		orgGuids []string
	}{orgGuids})
	fake.getManyOrgsByGuidMutex.Unlock()
	if fake.GetManyOrgsByGuidStub != nil {
		return fake.GetManyOrgsByGuidStub(orgGuids)
	} else {
		return fake.getManyOrgsByGuidReturns.result1, fake.getManyOrgsByGuidReturns.result2
	}
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGuidCallCount() int {
	fake.getManyOrgsByGuidMutex.RLock()
	defer fake.getManyOrgsByGuidMutex.RUnlock()
	return len(fake.getManyOrgsByGuidArgsForCall)
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGuidArgsForCall(i int) []string {
	fake.getManyOrgsByGuidMutex.RLock()
	defer fake.getManyOrgsByGuidMutex.RUnlock()
	return fake.getManyOrgsByGuidArgsForCall[i].orgGuids
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGuidReturns(result1 []models.Organization, result2 error) {
	fake.GetManyOrgsByGuidStub = nil
	fake.getManyOrgsByGuidReturns = struct {
		result1 []models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationRepository) FindByName(name string) (org models.Organization, apiErr error) {
	fake.findByNameMutex.Lock()
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		name string
	}{name})
	fake.findByNameMutex.Unlock()
	if fake.FindByNameStub != nil {
		return fake.FindByNameStub(name)
	} else {
		return fake.findByNameReturns.result1, fake.findByNameReturns.result2
	}
}

func (fake *FakeOrganizationRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeOrganizationRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return fake.findByNameArgsForCall[i].name
}

func (fake *FakeOrganizationRepository) FindByNameReturns(result1 models.Organization, result2 error) {
	fake.FindByNameStub = nil
	fake.findByNameReturns = struct {
		result1 models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationRepository) Create(org models.Organization) (apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		org models.Organization
	}{org})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(org)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeOrganizationRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeOrganizationRepository) CreateArgsForCall(i int) models.Organization {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].org
}

func (fake *FakeOrganizationRepository) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) Rename(orgGuid string, name string) (apiErr error) {
	fake.renameMutex.Lock()
	fake.renameArgsForCall = append(fake.renameArgsForCall, struct {
		orgGuid string
		name    string
	}{orgGuid, name})
	fake.renameMutex.Unlock()
	if fake.RenameStub != nil {
		return fake.RenameStub(orgGuid, name)
	} else {
		return fake.renameReturns.result1
	}
}

func (fake *FakeOrganizationRepository) RenameCallCount() int {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	return len(fake.renameArgsForCall)
}

func (fake *FakeOrganizationRepository) RenameArgsForCall(i int) (string, string) {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	return fake.renameArgsForCall[i].orgGuid, fake.renameArgsForCall[i].name
}

func (fake *FakeOrganizationRepository) RenameReturns(result1 error) {
	fake.RenameStub = nil
	fake.renameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) Delete(orgGuid string) (apiErr error) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		orgGuid string
	}{orgGuid})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(orgGuid)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeOrganizationRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeOrganizationRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].orgGuid
}

func (fake *FakeOrganizationRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) SharePrivateDomain(orgGuid string, domainGuid string) (apiErr error) {
	fake.sharePrivateDomainMutex.Lock()
	fake.sharePrivateDomainArgsForCall = append(fake.sharePrivateDomainArgsForCall, struct {
		orgGuid    string
		domainGuid string
	}{orgGuid, domainGuid})
	fake.sharePrivateDomainMutex.Unlock()
	if fake.SharePrivateDomainStub != nil {
		return fake.SharePrivateDomainStub(orgGuid, domainGuid)
	} else {
		return fake.sharePrivateDomainReturns.result1
	}
}

func (fake *FakeOrganizationRepository) SharePrivateDomainCallCount() int {
	fake.sharePrivateDomainMutex.RLock()
	defer fake.sharePrivateDomainMutex.RUnlock()
	return len(fake.sharePrivateDomainArgsForCall)
}

func (fake *FakeOrganizationRepository) SharePrivateDomainArgsForCall(i int) (string, string) {
	fake.sharePrivateDomainMutex.RLock()
	defer fake.sharePrivateDomainMutex.RUnlock()
	return fake.sharePrivateDomainArgsForCall[i].orgGuid, fake.sharePrivateDomainArgsForCall[i].domainGuid
}

func (fake *FakeOrganizationRepository) SharePrivateDomainReturns(result1 error) {
	fake.SharePrivateDomainStub = nil
	fake.sharePrivateDomainReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomain(orgGuid string, domainGuid string) (apiErr error) {
	fake.unsharePrivateDomainMutex.Lock()
	fake.unsharePrivateDomainArgsForCall = append(fake.unsharePrivateDomainArgsForCall, struct {
		orgGuid    string
		domainGuid string
	}{orgGuid, domainGuid})
	fake.unsharePrivateDomainMutex.Unlock()
	if fake.UnsharePrivateDomainStub != nil {
		return fake.UnsharePrivateDomainStub(orgGuid, domainGuid)
	} else {
		return fake.unsharePrivateDomainReturns.result1
	}
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomainCallCount() int {
	fake.unsharePrivateDomainMutex.RLock()
	defer fake.unsharePrivateDomainMutex.RUnlock()
	return len(fake.unsharePrivateDomainArgsForCall)
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomainArgsForCall(i int) (string, string) {
	fake.unsharePrivateDomainMutex.RLock()
	defer fake.unsharePrivateDomainMutex.RUnlock()
	return fake.unsharePrivateDomainArgsForCall[i].orgGuid, fake.unsharePrivateDomainArgsForCall[i].domainGuid
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomainReturns(result1 error) {
	fake.UnsharePrivateDomainStub = nil
	fake.unsharePrivateDomainReturns = struct {
		result1 error
	}{result1}
}

var _ organizations.OrganizationRepository = new(FakeOrganizationRepository)
