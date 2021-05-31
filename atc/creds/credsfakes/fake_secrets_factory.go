// Code generated by counterfeiter. DO NOT EDIT.
package credsfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/creds"
)

type FakeSecretsFactory struct {
	NewSecretsStub        func() creds.Secrets
	newSecretsMutex       sync.RWMutex
	newSecretsArgsForCall []struct {
	}
	newSecretsReturns struct {
		result1 creds.Secrets
	}
	newSecretsReturnsOnCall map[int]struct {
		result1 creds.Secrets
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecretsFactory) NewSecrets() creds.Secrets {
	fake.newSecretsMutex.Lock()
	ret, specificReturn := fake.newSecretsReturnsOnCall[len(fake.newSecretsArgsForCall)]
	fake.newSecretsArgsForCall = append(fake.newSecretsArgsForCall, struct {
	}{})
	stub := fake.NewSecretsStub
	fakeReturns := fake.newSecretsReturns
	fake.recordInvocation("NewSecrets", []interface{}{})
	fake.newSecretsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSecretsFactory) NewSecretsCallCount() int {
	fake.newSecretsMutex.RLock()
	defer fake.newSecretsMutex.RUnlock()
	return len(fake.newSecretsArgsForCall)
}

func (fake *FakeSecretsFactory) NewSecretsCalls(stub func() creds.Secrets) {
	fake.newSecretsMutex.Lock()
	defer fake.newSecretsMutex.Unlock()
	fake.NewSecretsStub = stub
}

func (fake *FakeSecretsFactory) NewSecretsReturns(result1 creds.Secrets) {
	fake.newSecretsMutex.Lock()
	defer fake.newSecretsMutex.Unlock()
	fake.NewSecretsStub = nil
	fake.newSecretsReturns = struct {
		result1 creds.Secrets
	}{result1}
}

func (fake *FakeSecretsFactory) NewSecretsReturnsOnCall(i int, result1 creds.Secrets) {
	fake.newSecretsMutex.Lock()
	defer fake.newSecretsMutex.Unlock()
	fake.NewSecretsStub = nil
	if fake.newSecretsReturnsOnCall == nil {
		fake.newSecretsReturnsOnCall = make(map[int]struct {
			result1 creds.Secrets
		})
	}
	fake.newSecretsReturnsOnCall[i] = struct {
		result1 creds.Secrets
	}{result1}
}

func (fake *FakeSecretsFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newSecretsMutex.RLock()
	defer fake.newSecretsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecretsFactory) recordInvocation(key string, args []interface{}) {
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

var _ creds.SecretsFactory = new(FakeSecretsFactory)