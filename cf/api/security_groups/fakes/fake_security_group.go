// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/fujitsu-cf/cli/cf/models"

	. "github.com/fujitsu-cf/cli/cf/api/security_groups"
)

type FakeSecurityGroupRepo struct {
	CreateStub        func(name string, rules []map[string]interface{}) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 []map[string]interface{}
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(guid string, rules []map[string]interface{}) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 string
		arg2 []map[string]interface{}
	}
	updateReturns struct {
		result1 error
	}
	ReadStub        func(string) (models.SecurityGroup, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 string
	}
	readReturns struct {
		result1 models.SecurityGroup
		result2 error
	}
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	FindAllStub        func() ([]models.SecurityGroup, error)
	findAllMutex       sync.RWMutex
	findAllArgsForCall []struct{}
	findAllReturns     struct {
		result1 []models.SecurityGroup
		result2 error
	}
}

func (fake *FakeSecurityGroupRepo) Create(arg1 string, arg2 []map[string]interface{}) error {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 []map[string]interface{}
	}{arg1, arg2})
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeSecurityGroupRepo) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSecurityGroupRepo) CreateArgsForCall(i int) (string, []map[string]interface{}) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2
}

func (fake *FakeSecurityGroupRepo) CreateReturns(result1 error) {
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupRepo) Update(arg1 string, arg2 []map[string]interface{}) error {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 string
		arg2 []map[string]interface{}
	}{arg1, arg2})
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeSecurityGroupRepo) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeSecurityGroupRepo) UpdateArgsForCall(i int) (string, []map[string]interface{}) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1, fake.updateArgsForCall[i].arg2
}

func (fake *FakeSecurityGroupRepo) UpdateReturns(result1 error) {
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupRepo) Read(arg1 string) (models.SecurityGroup, error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeSecurityGroupRepo) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeSecurityGroupRepo) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].arg1
}

func (fake *FakeSecurityGroupRepo) ReadReturns(result1 models.SecurityGroup, result2 error) {
	fake.readReturns = struct {
		result1 models.SecurityGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeSecurityGroupRepo) Delete(arg1 string) error {
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

func (fake *FakeSecurityGroupRepo) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSecurityGroupRepo) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1
}

func (fake *FakeSecurityGroupRepo) DeleteReturns(result1 error) {
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupRepo) FindAll() ([]models.SecurityGroup, error) {
	fake.findAllMutex.Lock()
	defer fake.findAllMutex.Unlock()
	fake.findAllArgsForCall = append(fake.findAllArgsForCall, struct{}{})
	if fake.FindAllStub != nil {
		return fake.FindAllStub()
	} else {
		return fake.findAllReturns.result1, fake.findAllReturns.result2
	}
}

func (fake *FakeSecurityGroupRepo) FindAllCallCount() int {
	fake.findAllMutex.RLock()
	defer fake.findAllMutex.RUnlock()
	return len(fake.findAllArgsForCall)
}

func (fake *FakeSecurityGroupRepo) FindAllReturns(result1 []models.SecurityGroup, result2 error) {
	fake.findAllReturns = struct {
		result1 []models.SecurityGroup
		result2 error
	}{result1, result2}
}

var _ SecurityGroupRepo = new(FakeSecurityGroupRepo)
