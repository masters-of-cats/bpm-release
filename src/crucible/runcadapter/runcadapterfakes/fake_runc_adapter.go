// This file was generated by counterfeiter
package runcadapterfakes

import (
	"crucible/config"
	"crucible/runcadapter"
	"sync"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeRuncAdapter struct {
	BuildSpecStub        func(jobName string, jobConfig *config.CrucibleConfig) (specs.Spec, error)
	buildSpecMutex       sync.RWMutex
	buildSpecArgsForCall []struct {
		jobName   string
		jobConfig *config.CrucibleConfig
	}
	buildSpecReturns struct {
		result1 specs.Spec
		result2 error
	}
	buildSpecReturnsOnCall map[int]struct {
		result1 specs.Spec
		result2 error
	}
	BuildBundleStub        func(bundlesRoot, jobName string, jobSpec specs.Spec) (string, error)
	buildBundleMutex       sync.RWMutex
	buildBundleArgsForCall []struct {
		bundlesRoot string
		jobName     string
		jobSpec     specs.Spec
	}
	buildBundleReturns struct {
		result1 string
		result2 error
	}
	buildBundleReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RunContainerStub        func(bundlePath, jobName string) error
	runContainerMutex       sync.RWMutex
	runContainerArgsForCall []struct {
		bundlePath string
		jobName    string
	}
	runContainerReturns struct {
		result1 error
	}
	runContainerReturnsOnCall map[int]struct {
		result1 error
	}
	StopContainerStub        func(jobName string) error
	stopContainerMutex       sync.RWMutex
	stopContainerArgsForCall []struct {
		jobName string
	}
	stopContainerReturns struct {
		result1 error
	}
	stopContainerReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyBundleStub        func(bundlesRoot, jobName string) error
	destroyBundleMutex       sync.RWMutex
	destroyBundleArgsForCall []struct {
		bundlesRoot string
		jobName     string
	}
	destroyBundleReturns struct {
		result1 error
	}
	destroyBundleReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRuncAdapter) BuildSpec(jobName string, jobConfig *config.CrucibleConfig) (specs.Spec, error) {
	fake.buildSpecMutex.Lock()
	ret, specificReturn := fake.buildSpecReturnsOnCall[len(fake.buildSpecArgsForCall)]
	fake.buildSpecArgsForCall = append(fake.buildSpecArgsForCall, struct {
		jobName   string
		jobConfig *config.CrucibleConfig
	}{jobName, jobConfig})
	fake.recordInvocation("BuildSpec", []interface{}{jobName, jobConfig})
	fake.buildSpecMutex.Unlock()
	if fake.BuildSpecStub != nil {
		return fake.BuildSpecStub(jobName, jobConfig)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.buildSpecReturns.result1, fake.buildSpecReturns.result2
}

func (fake *FakeRuncAdapter) BuildSpecCallCount() int {
	fake.buildSpecMutex.RLock()
	defer fake.buildSpecMutex.RUnlock()
	return len(fake.buildSpecArgsForCall)
}

func (fake *FakeRuncAdapter) BuildSpecArgsForCall(i int) (string, *config.CrucibleConfig) {
	fake.buildSpecMutex.RLock()
	defer fake.buildSpecMutex.RUnlock()
	return fake.buildSpecArgsForCall[i].jobName, fake.buildSpecArgsForCall[i].jobConfig
}

func (fake *FakeRuncAdapter) BuildSpecReturns(result1 specs.Spec, result2 error) {
	fake.BuildSpecStub = nil
	fake.buildSpecReturns = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncAdapter) BuildSpecReturnsOnCall(i int, result1 specs.Spec, result2 error) {
	fake.BuildSpecStub = nil
	if fake.buildSpecReturnsOnCall == nil {
		fake.buildSpecReturnsOnCall = make(map[int]struct {
			result1 specs.Spec
			result2 error
		})
	}
	fake.buildSpecReturnsOnCall[i] = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncAdapter) BuildBundle(bundlesRoot string, jobName string, jobSpec specs.Spec) (string, error) {
	fake.buildBundleMutex.Lock()
	ret, specificReturn := fake.buildBundleReturnsOnCall[len(fake.buildBundleArgsForCall)]
	fake.buildBundleArgsForCall = append(fake.buildBundleArgsForCall, struct {
		bundlesRoot string
		jobName     string
		jobSpec     specs.Spec
	}{bundlesRoot, jobName, jobSpec})
	fake.recordInvocation("BuildBundle", []interface{}{bundlesRoot, jobName, jobSpec})
	fake.buildBundleMutex.Unlock()
	if fake.BuildBundleStub != nil {
		return fake.BuildBundleStub(bundlesRoot, jobName, jobSpec)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.buildBundleReturns.result1, fake.buildBundleReturns.result2
}

func (fake *FakeRuncAdapter) BuildBundleCallCount() int {
	fake.buildBundleMutex.RLock()
	defer fake.buildBundleMutex.RUnlock()
	return len(fake.buildBundleArgsForCall)
}

func (fake *FakeRuncAdapter) BuildBundleArgsForCall(i int) (string, string, specs.Spec) {
	fake.buildBundleMutex.RLock()
	defer fake.buildBundleMutex.RUnlock()
	return fake.buildBundleArgsForCall[i].bundlesRoot, fake.buildBundleArgsForCall[i].jobName, fake.buildBundleArgsForCall[i].jobSpec
}

func (fake *FakeRuncAdapter) BuildBundleReturns(result1 string, result2 error) {
	fake.BuildBundleStub = nil
	fake.buildBundleReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncAdapter) BuildBundleReturnsOnCall(i int, result1 string, result2 error) {
	fake.BuildBundleStub = nil
	if fake.buildBundleReturnsOnCall == nil {
		fake.buildBundleReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.buildBundleReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncAdapter) RunContainer(bundlePath string, jobName string) error {
	fake.runContainerMutex.Lock()
	ret, specificReturn := fake.runContainerReturnsOnCall[len(fake.runContainerArgsForCall)]
	fake.runContainerArgsForCall = append(fake.runContainerArgsForCall, struct {
		bundlePath string
		jobName    string
	}{bundlePath, jobName})
	fake.recordInvocation("RunContainer", []interface{}{bundlePath, jobName})
	fake.runContainerMutex.Unlock()
	if fake.RunContainerStub != nil {
		return fake.RunContainerStub(bundlePath, jobName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runContainerReturns.result1
}

func (fake *FakeRuncAdapter) RunContainerCallCount() int {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return len(fake.runContainerArgsForCall)
}

func (fake *FakeRuncAdapter) RunContainerArgsForCall(i int) (string, string) {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return fake.runContainerArgsForCall[i].bundlePath, fake.runContainerArgsForCall[i].jobName
}

func (fake *FakeRuncAdapter) RunContainerReturns(result1 error) {
	fake.RunContainerStub = nil
	fake.runContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) RunContainerReturnsOnCall(i int, result1 error) {
	fake.RunContainerStub = nil
	if fake.runContainerReturnsOnCall == nil {
		fake.runContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) StopContainer(jobName string) error {
	fake.stopContainerMutex.Lock()
	ret, specificReturn := fake.stopContainerReturnsOnCall[len(fake.stopContainerArgsForCall)]
	fake.stopContainerArgsForCall = append(fake.stopContainerArgsForCall, struct {
		jobName string
	}{jobName})
	fake.recordInvocation("StopContainer", []interface{}{jobName})
	fake.stopContainerMutex.Unlock()
	if fake.StopContainerStub != nil {
		return fake.StopContainerStub(jobName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stopContainerReturns.result1
}

func (fake *FakeRuncAdapter) StopContainerCallCount() int {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return len(fake.stopContainerArgsForCall)
}

func (fake *FakeRuncAdapter) StopContainerArgsForCall(i int) string {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return fake.stopContainerArgsForCall[i].jobName
}

func (fake *FakeRuncAdapter) StopContainerReturns(result1 error) {
	fake.StopContainerStub = nil
	fake.stopContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) StopContainerReturnsOnCall(i int, result1 error) {
	fake.StopContainerStub = nil
	if fake.stopContainerReturnsOnCall == nil {
		fake.stopContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) DestroyBundle(bundlesRoot string, jobName string) error {
	fake.destroyBundleMutex.Lock()
	ret, specificReturn := fake.destroyBundleReturnsOnCall[len(fake.destroyBundleArgsForCall)]
	fake.destroyBundleArgsForCall = append(fake.destroyBundleArgsForCall, struct {
		bundlesRoot string
		jobName     string
	}{bundlesRoot, jobName})
	fake.recordInvocation("DestroyBundle", []interface{}{bundlesRoot, jobName})
	fake.destroyBundleMutex.Unlock()
	if fake.DestroyBundleStub != nil {
		return fake.DestroyBundleStub(bundlesRoot, jobName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyBundleReturns.result1
}

func (fake *FakeRuncAdapter) DestroyBundleCallCount() int {
	fake.destroyBundleMutex.RLock()
	defer fake.destroyBundleMutex.RUnlock()
	return len(fake.destroyBundleArgsForCall)
}

func (fake *FakeRuncAdapter) DestroyBundleArgsForCall(i int) (string, string) {
	fake.destroyBundleMutex.RLock()
	defer fake.destroyBundleMutex.RUnlock()
	return fake.destroyBundleArgsForCall[i].bundlesRoot, fake.destroyBundleArgsForCall[i].jobName
}

func (fake *FakeRuncAdapter) DestroyBundleReturns(result1 error) {
	fake.DestroyBundleStub = nil
	fake.destroyBundleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) DestroyBundleReturnsOnCall(i int, result1 error) {
	fake.DestroyBundleStub = nil
	if fake.destroyBundleReturnsOnCall == nil {
		fake.destroyBundleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyBundleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildSpecMutex.RLock()
	defer fake.buildSpecMutex.RUnlock()
	fake.buildBundleMutex.RLock()
	defer fake.buildBundleMutex.RUnlock()
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	fake.destroyBundleMutex.RLock()
	defer fake.destroyBundleMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRuncAdapter) recordInvocation(key string, args []interface{}) {
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

var _ runcadapter.RuncAdapter = new(FakeRuncAdapter)