// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/fujitsu-cf/cli/cf/terminal"
)

type FakeOutputCapture struct {
	SetOutputBucketStub        func(*[]string)
	setOutputBucketMutex       sync.RWMutex
	setOutputBucketArgsForCall []struct {
		arg1 *[]string
	}
}

func (fake *FakeOutputCapture) SetOutputBucket(arg1 *[]string) {
	fake.setOutputBucketMutex.Lock()
	fake.setOutputBucketArgsForCall = append(fake.setOutputBucketArgsForCall, struct {
		arg1 *[]string
	}{arg1})
	fake.setOutputBucketMutex.Unlock()
	if fake.SetOutputBucketStub != nil {
		fake.SetOutputBucketStub(arg1)
	}
}

func (fake *FakeOutputCapture) SetOutputBucketCallCount() int {
	fake.setOutputBucketMutex.RLock()
	defer fake.setOutputBucketMutex.RUnlock()
	return len(fake.setOutputBucketArgsForCall)
}

func (fake *FakeOutputCapture) SetOutputBucketArgsForCall(i int) *[]string {
	fake.setOutputBucketMutex.RLock()
	defer fake.setOutputBucketMutex.RUnlock()
	return fake.setOutputBucketArgsForCall[i].arg1
}

var _ terminal.OutputCapture = new(FakeOutputCapture)
