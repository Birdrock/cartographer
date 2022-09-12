// Code generated by counterfeiter. DO NOT EDIT.
package repositoryfakes

import (
	"sync"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type FakeRESTMapper struct {
	KindForStub        func(schema.GroupVersionResource) (schema.GroupVersionKind, error)
	kindForMutex       sync.RWMutex
	kindForArgsForCall []struct {
		arg1 schema.GroupVersionResource
	}
	kindForReturns struct {
		result1 schema.GroupVersionKind
		result2 error
	}
	kindForReturnsOnCall map[int]struct {
		result1 schema.GroupVersionKind
		result2 error
	}
	KindsForStub        func(schema.GroupVersionResource) ([]schema.GroupVersionKind, error)
	kindsForMutex       sync.RWMutex
	kindsForArgsForCall []struct {
		arg1 schema.GroupVersionResource
	}
	kindsForReturns struct {
		result1 []schema.GroupVersionKind
		result2 error
	}
	kindsForReturnsOnCall map[int]struct {
		result1 []schema.GroupVersionKind
		result2 error
	}
	RESTMappingStub        func(schema.GroupKind, ...string) (*meta.RESTMapping, error)
	rESTMappingMutex       sync.RWMutex
	rESTMappingArgsForCall []struct {
		arg1 schema.GroupKind
		arg2 []string
	}
	rESTMappingReturns struct {
		result1 *meta.RESTMapping
		result2 error
	}
	rESTMappingReturnsOnCall map[int]struct {
		result1 *meta.RESTMapping
		result2 error
	}
	RESTMappingsStub        func(schema.GroupKind, ...string) ([]*meta.RESTMapping, error)
	rESTMappingsMutex       sync.RWMutex
	rESTMappingsArgsForCall []struct {
		arg1 schema.GroupKind
		arg2 []string
	}
	rESTMappingsReturns struct {
		result1 []*meta.RESTMapping
		result2 error
	}
	rESTMappingsReturnsOnCall map[int]struct {
		result1 []*meta.RESTMapping
		result2 error
	}
	ResourceForStub        func(schema.GroupVersionResource) (schema.GroupVersionResource, error)
	resourceForMutex       sync.RWMutex
	resourceForArgsForCall []struct {
		arg1 schema.GroupVersionResource
	}
	resourceForReturns struct {
		result1 schema.GroupVersionResource
		result2 error
	}
	resourceForReturnsOnCall map[int]struct {
		result1 schema.GroupVersionResource
		result2 error
	}
	ResourceSingularizerStub        func(string) (string, error)
	resourceSingularizerMutex       sync.RWMutex
	resourceSingularizerArgsForCall []struct {
		arg1 string
	}
	resourceSingularizerReturns struct {
		result1 string
		result2 error
	}
	resourceSingularizerReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ResourcesForStub        func(schema.GroupVersionResource) ([]schema.GroupVersionResource, error)
	resourcesForMutex       sync.RWMutex
	resourcesForArgsForCall []struct {
		arg1 schema.GroupVersionResource
	}
	resourcesForReturns struct {
		result1 []schema.GroupVersionResource
		result2 error
	}
	resourcesForReturnsOnCall map[int]struct {
		result1 []schema.GroupVersionResource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRESTMapper) KindFor(arg1 schema.GroupVersionResource) (schema.GroupVersionKind, error) {
	fake.kindForMutex.Lock()
	ret, specificReturn := fake.kindForReturnsOnCall[len(fake.kindForArgsForCall)]
	fake.kindForArgsForCall = append(fake.kindForArgsForCall, struct {
		arg1 schema.GroupVersionResource
	}{arg1})
	stub := fake.KindForStub
	fakeReturns := fake.kindForReturns
	fake.recordInvocation("KindFor", []interface{}{arg1})
	fake.kindForMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRESTMapper) KindForCallCount() int {
	fake.kindForMutex.RLock()
	defer fake.kindForMutex.RUnlock()
	return len(fake.kindForArgsForCall)
}

func (fake *FakeRESTMapper) KindForCalls(stub func(schema.GroupVersionResource) (schema.GroupVersionKind, error)) {
	fake.kindForMutex.Lock()
	defer fake.kindForMutex.Unlock()
	fake.KindForStub = stub
}

func (fake *FakeRESTMapper) KindForArgsForCall(i int) schema.GroupVersionResource {
	fake.kindForMutex.RLock()
	defer fake.kindForMutex.RUnlock()
	argsForCall := fake.kindForArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRESTMapper) KindForReturns(result1 schema.GroupVersionKind, result2 error) {
	fake.kindForMutex.Lock()
	defer fake.kindForMutex.Unlock()
	fake.KindForStub = nil
	fake.kindForReturns = struct {
		result1 schema.GroupVersionKind
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) KindForReturnsOnCall(i int, result1 schema.GroupVersionKind, result2 error) {
	fake.kindForMutex.Lock()
	defer fake.kindForMutex.Unlock()
	fake.KindForStub = nil
	if fake.kindForReturnsOnCall == nil {
		fake.kindForReturnsOnCall = make(map[int]struct {
			result1 schema.GroupVersionKind
			result2 error
		})
	}
	fake.kindForReturnsOnCall[i] = struct {
		result1 schema.GroupVersionKind
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) KindsFor(arg1 schema.GroupVersionResource) ([]schema.GroupVersionKind, error) {
	fake.kindsForMutex.Lock()
	ret, specificReturn := fake.kindsForReturnsOnCall[len(fake.kindsForArgsForCall)]
	fake.kindsForArgsForCall = append(fake.kindsForArgsForCall, struct {
		arg1 schema.GroupVersionResource
	}{arg1})
	stub := fake.KindsForStub
	fakeReturns := fake.kindsForReturns
	fake.recordInvocation("KindsFor", []interface{}{arg1})
	fake.kindsForMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRESTMapper) KindsForCallCount() int {
	fake.kindsForMutex.RLock()
	defer fake.kindsForMutex.RUnlock()
	return len(fake.kindsForArgsForCall)
}

func (fake *FakeRESTMapper) KindsForCalls(stub func(schema.GroupVersionResource) ([]schema.GroupVersionKind, error)) {
	fake.kindsForMutex.Lock()
	defer fake.kindsForMutex.Unlock()
	fake.KindsForStub = stub
}

func (fake *FakeRESTMapper) KindsForArgsForCall(i int) schema.GroupVersionResource {
	fake.kindsForMutex.RLock()
	defer fake.kindsForMutex.RUnlock()
	argsForCall := fake.kindsForArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRESTMapper) KindsForReturns(result1 []schema.GroupVersionKind, result2 error) {
	fake.kindsForMutex.Lock()
	defer fake.kindsForMutex.Unlock()
	fake.KindsForStub = nil
	fake.kindsForReturns = struct {
		result1 []schema.GroupVersionKind
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) KindsForReturnsOnCall(i int, result1 []schema.GroupVersionKind, result2 error) {
	fake.kindsForMutex.Lock()
	defer fake.kindsForMutex.Unlock()
	fake.KindsForStub = nil
	if fake.kindsForReturnsOnCall == nil {
		fake.kindsForReturnsOnCall = make(map[int]struct {
			result1 []schema.GroupVersionKind
			result2 error
		})
	}
	fake.kindsForReturnsOnCall[i] = struct {
		result1 []schema.GroupVersionKind
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) RESTMapping(arg1 schema.GroupKind, arg2 ...string) (*meta.RESTMapping, error) {
	fake.rESTMappingMutex.Lock()
	ret, specificReturn := fake.rESTMappingReturnsOnCall[len(fake.rESTMappingArgsForCall)]
	fake.rESTMappingArgsForCall = append(fake.rESTMappingArgsForCall, struct {
		arg1 schema.GroupKind
		arg2 []string
	}{arg1, arg2})
	stub := fake.RESTMappingStub
	fakeReturns := fake.rESTMappingReturns
	fake.recordInvocation("RESTMapping", []interface{}{arg1, arg2})
	fake.rESTMappingMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRESTMapper) RESTMappingCallCount() int {
	fake.rESTMappingMutex.RLock()
	defer fake.rESTMappingMutex.RUnlock()
	return len(fake.rESTMappingArgsForCall)
}

func (fake *FakeRESTMapper) RESTMappingCalls(stub func(schema.GroupKind, ...string) (*meta.RESTMapping, error)) {
	fake.rESTMappingMutex.Lock()
	defer fake.rESTMappingMutex.Unlock()
	fake.RESTMappingStub = stub
}

func (fake *FakeRESTMapper) RESTMappingArgsForCall(i int) (schema.GroupKind, []string) {
	fake.rESTMappingMutex.RLock()
	defer fake.rESTMappingMutex.RUnlock()
	argsForCall := fake.rESTMappingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRESTMapper) RESTMappingReturns(result1 *meta.RESTMapping, result2 error) {
	fake.rESTMappingMutex.Lock()
	defer fake.rESTMappingMutex.Unlock()
	fake.RESTMappingStub = nil
	fake.rESTMappingReturns = struct {
		result1 *meta.RESTMapping
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) RESTMappingReturnsOnCall(i int, result1 *meta.RESTMapping, result2 error) {
	fake.rESTMappingMutex.Lock()
	defer fake.rESTMappingMutex.Unlock()
	fake.RESTMappingStub = nil
	if fake.rESTMappingReturnsOnCall == nil {
		fake.rESTMappingReturnsOnCall = make(map[int]struct {
			result1 *meta.RESTMapping
			result2 error
		})
	}
	fake.rESTMappingReturnsOnCall[i] = struct {
		result1 *meta.RESTMapping
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) RESTMappings(arg1 schema.GroupKind, arg2 ...string) ([]*meta.RESTMapping, error) {
	fake.rESTMappingsMutex.Lock()
	ret, specificReturn := fake.rESTMappingsReturnsOnCall[len(fake.rESTMappingsArgsForCall)]
	fake.rESTMappingsArgsForCall = append(fake.rESTMappingsArgsForCall, struct {
		arg1 schema.GroupKind
		arg2 []string
	}{arg1, arg2})
	stub := fake.RESTMappingsStub
	fakeReturns := fake.rESTMappingsReturns
	fake.recordInvocation("RESTMappings", []interface{}{arg1, arg2})
	fake.rESTMappingsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRESTMapper) RESTMappingsCallCount() int {
	fake.rESTMappingsMutex.RLock()
	defer fake.rESTMappingsMutex.RUnlock()
	return len(fake.rESTMappingsArgsForCall)
}

func (fake *FakeRESTMapper) RESTMappingsCalls(stub func(schema.GroupKind, ...string) ([]*meta.RESTMapping, error)) {
	fake.rESTMappingsMutex.Lock()
	defer fake.rESTMappingsMutex.Unlock()
	fake.RESTMappingsStub = stub
}

func (fake *FakeRESTMapper) RESTMappingsArgsForCall(i int) (schema.GroupKind, []string) {
	fake.rESTMappingsMutex.RLock()
	defer fake.rESTMappingsMutex.RUnlock()
	argsForCall := fake.rESTMappingsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRESTMapper) RESTMappingsReturns(result1 []*meta.RESTMapping, result2 error) {
	fake.rESTMappingsMutex.Lock()
	defer fake.rESTMappingsMutex.Unlock()
	fake.RESTMappingsStub = nil
	fake.rESTMappingsReturns = struct {
		result1 []*meta.RESTMapping
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) RESTMappingsReturnsOnCall(i int, result1 []*meta.RESTMapping, result2 error) {
	fake.rESTMappingsMutex.Lock()
	defer fake.rESTMappingsMutex.Unlock()
	fake.RESTMappingsStub = nil
	if fake.rESTMappingsReturnsOnCall == nil {
		fake.rESTMappingsReturnsOnCall = make(map[int]struct {
			result1 []*meta.RESTMapping
			result2 error
		})
	}
	fake.rESTMappingsReturnsOnCall[i] = struct {
		result1 []*meta.RESTMapping
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) ResourceFor(arg1 schema.GroupVersionResource) (schema.GroupVersionResource, error) {
	fake.resourceForMutex.Lock()
	ret, specificReturn := fake.resourceForReturnsOnCall[len(fake.resourceForArgsForCall)]
	fake.resourceForArgsForCall = append(fake.resourceForArgsForCall, struct {
		arg1 schema.GroupVersionResource
	}{arg1})
	stub := fake.ResourceForStub
	fakeReturns := fake.resourceForReturns
	fake.recordInvocation("ResourceFor", []interface{}{arg1})
	fake.resourceForMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRESTMapper) ResourceForCallCount() int {
	fake.resourceForMutex.RLock()
	defer fake.resourceForMutex.RUnlock()
	return len(fake.resourceForArgsForCall)
}

func (fake *FakeRESTMapper) ResourceForCalls(stub func(schema.GroupVersionResource) (schema.GroupVersionResource, error)) {
	fake.resourceForMutex.Lock()
	defer fake.resourceForMutex.Unlock()
	fake.ResourceForStub = stub
}

func (fake *FakeRESTMapper) ResourceForArgsForCall(i int) schema.GroupVersionResource {
	fake.resourceForMutex.RLock()
	defer fake.resourceForMutex.RUnlock()
	argsForCall := fake.resourceForArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRESTMapper) ResourceForReturns(result1 schema.GroupVersionResource, result2 error) {
	fake.resourceForMutex.Lock()
	defer fake.resourceForMutex.Unlock()
	fake.ResourceForStub = nil
	fake.resourceForReturns = struct {
		result1 schema.GroupVersionResource
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) ResourceForReturnsOnCall(i int, result1 schema.GroupVersionResource, result2 error) {
	fake.resourceForMutex.Lock()
	defer fake.resourceForMutex.Unlock()
	fake.ResourceForStub = nil
	if fake.resourceForReturnsOnCall == nil {
		fake.resourceForReturnsOnCall = make(map[int]struct {
			result1 schema.GroupVersionResource
			result2 error
		})
	}
	fake.resourceForReturnsOnCall[i] = struct {
		result1 schema.GroupVersionResource
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) ResourceSingularizer(arg1 string) (string, error) {
	fake.resourceSingularizerMutex.Lock()
	ret, specificReturn := fake.resourceSingularizerReturnsOnCall[len(fake.resourceSingularizerArgsForCall)]
	fake.resourceSingularizerArgsForCall = append(fake.resourceSingularizerArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ResourceSingularizerStub
	fakeReturns := fake.resourceSingularizerReturns
	fake.recordInvocation("ResourceSingularizer", []interface{}{arg1})
	fake.resourceSingularizerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRESTMapper) ResourceSingularizerCallCount() int {
	fake.resourceSingularizerMutex.RLock()
	defer fake.resourceSingularizerMutex.RUnlock()
	return len(fake.resourceSingularizerArgsForCall)
}

func (fake *FakeRESTMapper) ResourceSingularizerCalls(stub func(string) (string, error)) {
	fake.resourceSingularizerMutex.Lock()
	defer fake.resourceSingularizerMutex.Unlock()
	fake.ResourceSingularizerStub = stub
}

func (fake *FakeRESTMapper) ResourceSingularizerArgsForCall(i int) string {
	fake.resourceSingularizerMutex.RLock()
	defer fake.resourceSingularizerMutex.RUnlock()
	argsForCall := fake.resourceSingularizerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRESTMapper) ResourceSingularizerReturns(result1 string, result2 error) {
	fake.resourceSingularizerMutex.Lock()
	defer fake.resourceSingularizerMutex.Unlock()
	fake.ResourceSingularizerStub = nil
	fake.resourceSingularizerReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) ResourceSingularizerReturnsOnCall(i int, result1 string, result2 error) {
	fake.resourceSingularizerMutex.Lock()
	defer fake.resourceSingularizerMutex.Unlock()
	fake.ResourceSingularizerStub = nil
	if fake.resourceSingularizerReturnsOnCall == nil {
		fake.resourceSingularizerReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.resourceSingularizerReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) ResourcesFor(arg1 schema.GroupVersionResource) ([]schema.GroupVersionResource, error) {
	fake.resourcesForMutex.Lock()
	ret, specificReturn := fake.resourcesForReturnsOnCall[len(fake.resourcesForArgsForCall)]
	fake.resourcesForArgsForCall = append(fake.resourcesForArgsForCall, struct {
		arg1 schema.GroupVersionResource
	}{arg1})
	stub := fake.ResourcesForStub
	fakeReturns := fake.resourcesForReturns
	fake.recordInvocation("ResourcesFor", []interface{}{arg1})
	fake.resourcesForMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRESTMapper) ResourcesForCallCount() int {
	fake.resourcesForMutex.RLock()
	defer fake.resourcesForMutex.RUnlock()
	return len(fake.resourcesForArgsForCall)
}

func (fake *FakeRESTMapper) ResourcesForCalls(stub func(schema.GroupVersionResource) ([]schema.GroupVersionResource, error)) {
	fake.resourcesForMutex.Lock()
	defer fake.resourcesForMutex.Unlock()
	fake.ResourcesForStub = stub
}

func (fake *FakeRESTMapper) ResourcesForArgsForCall(i int) schema.GroupVersionResource {
	fake.resourcesForMutex.RLock()
	defer fake.resourcesForMutex.RUnlock()
	argsForCall := fake.resourcesForArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRESTMapper) ResourcesForReturns(result1 []schema.GroupVersionResource, result2 error) {
	fake.resourcesForMutex.Lock()
	defer fake.resourcesForMutex.Unlock()
	fake.ResourcesForStub = nil
	fake.resourcesForReturns = struct {
		result1 []schema.GroupVersionResource
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) ResourcesForReturnsOnCall(i int, result1 []schema.GroupVersionResource, result2 error) {
	fake.resourcesForMutex.Lock()
	defer fake.resourcesForMutex.Unlock()
	fake.ResourcesForStub = nil
	if fake.resourcesForReturnsOnCall == nil {
		fake.resourcesForReturnsOnCall = make(map[int]struct {
			result1 []schema.GroupVersionResource
			result2 error
		})
	}
	fake.resourcesForReturnsOnCall[i] = struct {
		result1 []schema.GroupVersionResource
		result2 error
	}{result1, result2}
}

func (fake *FakeRESTMapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.kindForMutex.RLock()
	defer fake.kindForMutex.RUnlock()
	fake.kindsForMutex.RLock()
	defer fake.kindsForMutex.RUnlock()
	fake.rESTMappingMutex.RLock()
	defer fake.rESTMappingMutex.RUnlock()
	fake.rESTMappingsMutex.RLock()
	defer fake.rESTMappingsMutex.RUnlock()
	fake.resourceForMutex.RLock()
	defer fake.resourceForMutex.RUnlock()
	fake.resourceSingularizerMutex.RLock()
	defer fake.resourceSingularizerMutex.RUnlock()
	fake.resourcesForMutex.RLock()
	defer fake.resourcesForMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRESTMapper) recordInvocation(key string, args []interface{}) {
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

var _ meta.RESTMapper = new(FakeRESTMapper)
