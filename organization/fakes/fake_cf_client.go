// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	organization "github.com/pivotalservices/cf-mgmt/organization"
)

type FakeCFClient struct {
	CreateDomainStub        func(string, string) (*cfclient.Domain, error)
	createDomainMutex       sync.RWMutex
	createDomainArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createDomainReturns struct {
		result1 *cfclient.Domain
		result2 error
	}
	createDomainReturnsOnCall map[int]struct {
		result1 *cfclient.Domain
		result2 error
	}
	CreateOrgStub        func(cfclient.OrgRequest) (cfclient.Org, error)
	createOrgMutex       sync.RWMutex
	createOrgArgsForCall []struct {
		arg1 cfclient.OrgRequest
	}
	createOrgReturns struct {
		result1 cfclient.Org
		result2 error
	}
	createOrgReturnsOnCall map[int]struct {
		result1 cfclient.Org
		result2 error
	}
	DeleteDomainStub        func(string) error
	deleteDomainMutex       sync.RWMutex
	deleteDomainArgsForCall []struct {
		arg1 string
	}
	deleteDomainReturns struct {
		result1 error
	}
	deleteDomainReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteOrgStub        func(string, bool, bool) error
	deleteOrgMutex       sync.RWMutex
	deleteOrgArgsForCall []struct {
		arg1 string
		arg2 bool
		arg3 bool
	}
	deleteOrgReturns struct {
		result1 error
	}
	deleteOrgReturnsOnCall map[int]struct {
		result1 error
	}
	GetOrgByGuidStub        func(string) (cfclient.Org, error)
	getOrgByGuidMutex       sync.RWMutex
	getOrgByGuidArgsForCall []struct {
		arg1 string
	}
	getOrgByGuidReturns struct {
		result1 cfclient.Org
		result2 error
	}
	getOrgByGuidReturnsOnCall map[int]struct {
		result1 cfclient.Org
		result2 error
	}
	ListDomainsStub        func() ([]cfclient.Domain, error)
	listDomainsMutex       sync.RWMutex
	listDomainsArgsForCall []struct {
	}
	listDomainsReturns struct {
		result1 []cfclient.Domain
		result2 error
	}
	listDomainsReturnsOnCall map[int]struct {
		result1 []cfclient.Domain
		result2 error
	}
	ListOrgPrivateDomainsStub        func(string) ([]cfclient.Domain, error)
	listOrgPrivateDomainsMutex       sync.RWMutex
	listOrgPrivateDomainsArgsForCall []struct {
		arg1 string
	}
	listOrgPrivateDomainsReturns struct {
		result1 []cfclient.Domain
		result2 error
	}
	listOrgPrivateDomainsReturnsOnCall map[int]struct {
		result1 []cfclient.Domain
		result2 error
	}
	ListOrgsStub        func() ([]cfclient.Org, error)
	listOrgsMutex       sync.RWMutex
	listOrgsArgsForCall []struct {
	}
	listOrgsReturns struct {
		result1 []cfclient.Org
		result2 error
	}
	listOrgsReturnsOnCall map[int]struct {
		result1 []cfclient.Org
		result2 error
	}
	ShareOrgPrivateDomainStub        func(string, string) (*cfclient.Domain, error)
	shareOrgPrivateDomainMutex       sync.RWMutex
	shareOrgPrivateDomainArgsForCall []struct {
		arg1 string
		arg2 string
	}
	shareOrgPrivateDomainReturns struct {
		result1 *cfclient.Domain
		result2 error
	}
	shareOrgPrivateDomainReturnsOnCall map[int]struct {
		result1 *cfclient.Domain
		result2 error
	}
	SupportsMetadataAPIStub        func() (bool, error)
	supportsMetadataAPIMutex       sync.RWMutex
	supportsMetadataAPIArgsForCall []struct {
	}
	supportsMetadataAPIReturns struct {
		result1 bool
		result2 error
	}
	supportsMetadataAPIReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	UnshareOrgPrivateDomainStub        func(string, string) error
	unshareOrgPrivateDomainMutex       sync.RWMutex
	unshareOrgPrivateDomainArgsForCall []struct {
		arg1 string
		arg2 string
	}
	unshareOrgPrivateDomainReturns struct {
		result1 error
	}
	unshareOrgPrivateDomainReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateOrgStub        func(string, cfclient.OrgRequest) (cfclient.Org, error)
	updateOrgMutex       sync.RWMutex
	updateOrgArgsForCall []struct {
		arg1 string
		arg2 cfclient.OrgRequest
	}
	updateOrgReturns struct {
		result1 cfclient.Org
		result2 error
	}
	updateOrgReturnsOnCall map[int]struct {
		result1 cfclient.Org
		result2 error
	}
	UpdateOrgMetadataStub        func(string, cfclient.Metadata) error
	updateOrgMetadataMutex       sync.RWMutex
	updateOrgMetadataArgsForCall []struct {
		arg1 string
		arg2 cfclient.Metadata
	}
	updateOrgMetadataReturns struct {
		result1 error
	}
	updateOrgMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFClient) CreateDomain(arg1 string, arg2 string) (*cfclient.Domain, error) {
	fake.createDomainMutex.Lock()
	ret, specificReturn := fake.createDomainReturnsOnCall[len(fake.createDomainArgsForCall)]
	fake.createDomainArgsForCall = append(fake.createDomainArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateDomain", []interface{}{arg1, arg2})
	fake.createDomainMutex.Unlock()
	if fake.CreateDomainStub != nil {
		return fake.CreateDomainStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createDomainReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) CreateDomainCallCount() int {
	fake.createDomainMutex.RLock()
	defer fake.createDomainMutex.RUnlock()
	return len(fake.createDomainArgsForCall)
}

func (fake *FakeCFClient) CreateDomainCalls(stub func(string, string) (*cfclient.Domain, error)) {
	fake.createDomainMutex.Lock()
	defer fake.createDomainMutex.Unlock()
	fake.CreateDomainStub = stub
}

func (fake *FakeCFClient) CreateDomainArgsForCall(i int) (string, string) {
	fake.createDomainMutex.RLock()
	defer fake.createDomainMutex.RUnlock()
	argsForCall := fake.createDomainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFClient) CreateDomainReturns(result1 *cfclient.Domain, result2 error) {
	fake.createDomainMutex.Lock()
	defer fake.createDomainMutex.Unlock()
	fake.CreateDomainStub = nil
	fake.createDomainReturns = struct {
		result1 *cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) CreateDomainReturnsOnCall(i int, result1 *cfclient.Domain, result2 error) {
	fake.createDomainMutex.Lock()
	defer fake.createDomainMutex.Unlock()
	fake.CreateDomainStub = nil
	if fake.createDomainReturnsOnCall == nil {
		fake.createDomainReturnsOnCall = make(map[int]struct {
			result1 *cfclient.Domain
			result2 error
		})
	}
	fake.createDomainReturnsOnCall[i] = struct {
		result1 *cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) CreateOrg(arg1 cfclient.OrgRequest) (cfclient.Org, error) {
	fake.createOrgMutex.Lock()
	ret, specificReturn := fake.createOrgReturnsOnCall[len(fake.createOrgArgsForCall)]
	fake.createOrgArgsForCall = append(fake.createOrgArgsForCall, struct {
		arg1 cfclient.OrgRequest
	}{arg1})
	fake.recordInvocation("CreateOrg", []interface{}{arg1})
	fake.createOrgMutex.Unlock()
	if fake.CreateOrgStub != nil {
		return fake.CreateOrgStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createOrgReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) CreateOrgCallCount() int {
	fake.createOrgMutex.RLock()
	defer fake.createOrgMutex.RUnlock()
	return len(fake.createOrgArgsForCall)
}

func (fake *FakeCFClient) CreateOrgCalls(stub func(cfclient.OrgRequest) (cfclient.Org, error)) {
	fake.createOrgMutex.Lock()
	defer fake.createOrgMutex.Unlock()
	fake.CreateOrgStub = stub
}

func (fake *FakeCFClient) CreateOrgArgsForCall(i int) cfclient.OrgRequest {
	fake.createOrgMutex.RLock()
	defer fake.createOrgMutex.RUnlock()
	argsForCall := fake.createOrgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) CreateOrgReturns(result1 cfclient.Org, result2 error) {
	fake.createOrgMutex.Lock()
	defer fake.createOrgMutex.Unlock()
	fake.CreateOrgStub = nil
	fake.createOrgReturns = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) CreateOrgReturnsOnCall(i int, result1 cfclient.Org, result2 error) {
	fake.createOrgMutex.Lock()
	defer fake.createOrgMutex.Unlock()
	fake.CreateOrgStub = nil
	if fake.createOrgReturnsOnCall == nil {
		fake.createOrgReturnsOnCall = make(map[int]struct {
			result1 cfclient.Org
			result2 error
		})
	}
	fake.createOrgReturnsOnCall[i] = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) DeleteDomain(arg1 string) error {
	fake.deleteDomainMutex.Lock()
	ret, specificReturn := fake.deleteDomainReturnsOnCall[len(fake.deleteDomainArgsForCall)]
	fake.deleteDomainArgsForCall = append(fake.deleteDomainArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteDomain", []interface{}{arg1})
	fake.deleteDomainMutex.Unlock()
	if fake.DeleteDomainStub != nil {
		return fake.DeleteDomainStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteDomainReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) DeleteDomainCallCount() int {
	fake.deleteDomainMutex.RLock()
	defer fake.deleteDomainMutex.RUnlock()
	return len(fake.deleteDomainArgsForCall)
}

func (fake *FakeCFClient) DeleteDomainCalls(stub func(string) error) {
	fake.deleteDomainMutex.Lock()
	defer fake.deleteDomainMutex.Unlock()
	fake.DeleteDomainStub = stub
}

func (fake *FakeCFClient) DeleteDomainArgsForCall(i int) string {
	fake.deleteDomainMutex.RLock()
	defer fake.deleteDomainMutex.RUnlock()
	argsForCall := fake.deleteDomainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) DeleteDomainReturns(result1 error) {
	fake.deleteDomainMutex.Lock()
	defer fake.deleteDomainMutex.Unlock()
	fake.DeleteDomainStub = nil
	fake.deleteDomainReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) DeleteDomainReturnsOnCall(i int, result1 error) {
	fake.deleteDomainMutex.Lock()
	defer fake.deleteDomainMutex.Unlock()
	fake.DeleteDomainStub = nil
	if fake.deleteDomainReturnsOnCall == nil {
		fake.deleteDomainReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteDomainReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) DeleteOrg(arg1 string, arg2 bool, arg3 bool) error {
	fake.deleteOrgMutex.Lock()
	ret, specificReturn := fake.deleteOrgReturnsOnCall[len(fake.deleteOrgArgsForCall)]
	fake.deleteOrgArgsForCall = append(fake.deleteOrgArgsForCall, struct {
		arg1 string
		arg2 bool
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteOrg", []interface{}{arg1, arg2, arg3})
	fake.deleteOrgMutex.Unlock()
	if fake.DeleteOrgStub != nil {
		return fake.DeleteOrgStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteOrgReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) DeleteOrgCallCount() int {
	fake.deleteOrgMutex.RLock()
	defer fake.deleteOrgMutex.RUnlock()
	return len(fake.deleteOrgArgsForCall)
}

func (fake *FakeCFClient) DeleteOrgCalls(stub func(string, bool, bool) error) {
	fake.deleteOrgMutex.Lock()
	defer fake.deleteOrgMutex.Unlock()
	fake.DeleteOrgStub = stub
}

func (fake *FakeCFClient) DeleteOrgArgsForCall(i int) (string, bool, bool) {
	fake.deleteOrgMutex.RLock()
	defer fake.deleteOrgMutex.RUnlock()
	argsForCall := fake.deleteOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCFClient) DeleteOrgReturns(result1 error) {
	fake.deleteOrgMutex.Lock()
	defer fake.deleteOrgMutex.Unlock()
	fake.DeleteOrgStub = nil
	fake.deleteOrgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) DeleteOrgReturnsOnCall(i int, result1 error) {
	fake.deleteOrgMutex.Lock()
	defer fake.deleteOrgMutex.Unlock()
	fake.DeleteOrgStub = nil
	if fake.deleteOrgReturnsOnCall == nil {
		fake.deleteOrgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteOrgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) GetOrgByGuid(arg1 string) (cfclient.Org, error) {
	fake.getOrgByGuidMutex.Lock()
	ret, specificReturn := fake.getOrgByGuidReturnsOnCall[len(fake.getOrgByGuidArgsForCall)]
	fake.getOrgByGuidArgsForCall = append(fake.getOrgByGuidArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrgByGuid", []interface{}{arg1})
	fake.getOrgByGuidMutex.Unlock()
	if fake.GetOrgByGuidStub != nil {
		return fake.GetOrgByGuidStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getOrgByGuidReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) GetOrgByGuidCallCount() int {
	fake.getOrgByGuidMutex.RLock()
	defer fake.getOrgByGuidMutex.RUnlock()
	return len(fake.getOrgByGuidArgsForCall)
}

func (fake *FakeCFClient) GetOrgByGuidCalls(stub func(string) (cfclient.Org, error)) {
	fake.getOrgByGuidMutex.Lock()
	defer fake.getOrgByGuidMutex.Unlock()
	fake.GetOrgByGuidStub = stub
}

func (fake *FakeCFClient) GetOrgByGuidArgsForCall(i int) string {
	fake.getOrgByGuidMutex.RLock()
	defer fake.getOrgByGuidMutex.RUnlock()
	argsForCall := fake.getOrgByGuidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) GetOrgByGuidReturns(result1 cfclient.Org, result2 error) {
	fake.getOrgByGuidMutex.Lock()
	defer fake.getOrgByGuidMutex.Unlock()
	fake.GetOrgByGuidStub = nil
	fake.getOrgByGuidReturns = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) GetOrgByGuidReturnsOnCall(i int, result1 cfclient.Org, result2 error) {
	fake.getOrgByGuidMutex.Lock()
	defer fake.getOrgByGuidMutex.Unlock()
	fake.GetOrgByGuidStub = nil
	if fake.getOrgByGuidReturnsOnCall == nil {
		fake.getOrgByGuidReturnsOnCall = make(map[int]struct {
			result1 cfclient.Org
			result2 error
		})
	}
	fake.getOrgByGuidReturnsOnCall[i] = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListDomains() ([]cfclient.Domain, error) {
	fake.listDomainsMutex.Lock()
	ret, specificReturn := fake.listDomainsReturnsOnCall[len(fake.listDomainsArgsForCall)]
	fake.listDomainsArgsForCall = append(fake.listDomainsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListDomains", []interface{}{})
	fake.listDomainsMutex.Unlock()
	if fake.ListDomainsStub != nil {
		return fake.ListDomainsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listDomainsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) ListDomainsCallCount() int {
	fake.listDomainsMutex.RLock()
	defer fake.listDomainsMutex.RUnlock()
	return len(fake.listDomainsArgsForCall)
}

func (fake *FakeCFClient) ListDomainsCalls(stub func() ([]cfclient.Domain, error)) {
	fake.listDomainsMutex.Lock()
	defer fake.listDomainsMutex.Unlock()
	fake.ListDomainsStub = stub
}

func (fake *FakeCFClient) ListDomainsReturns(result1 []cfclient.Domain, result2 error) {
	fake.listDomainsMutex.Lock()
	defer fake.listDomainsMutex.Unlock()
	fake.ListDomainsStub = nil
	fake.listDomainsReturns = struct {
		result1 []cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListDomainsReturnsOnCall(i int, result1 []cfclient.Domain, result2 error) {
	fake.listDomainsMutex.Lock()
	defer fake.listDomainsMutex.Unlock()
	fake.ListDomainsStub = nil
	if fake.listDomainsReturnsOnCall == nil {
		fake.listDomainsReturnsOnCall = make(map[int]struct {
			result1 []cfclient.Domain
			result2 error
		})
	}
	fake.listDomainsReturnsOnCall[i] = struct {
		result1 []cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListOrgPrivateDomains(arg1 string) ([]cfclient.Domain, error) {
	fake.listOrgPrivateDomainsMutex.Lock()
	ret, specificReturn := fake.listOrgPrivateDomainsReturnsOnCall[len(fake.listOrgPrivateDomainsArgsForCall)]
	fake.listOrgPrivateDomainsArgsForCall = append(fake.listOrgPrivateDomainsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListOrgPrivateDomains", []interface{}{arg1})
	fake.listOrgPrivateDomainsMutex.Unlock()
	if fake.ListOrgPrivateDomainsStub != nil {
		return fake.ListOrgPrivateDomainsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listOrgPrivateDomainsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) ListOrgPrivateDomainsCallCount() int {
	fake.listOrgPrivateDomainsMutex.RLock()
	defer fake.listOrgPrivateDomainsMutex.RUnlock()
	return len(fake.listOrgPrivateDomainsArgsForCall)
}

func (fake *FakeCFClient) ListOrgPrivateDomainsCalls(stub func(string) ([]cfclient.Domain, error)) {
	fake.listOrgPrivateDomainsMutex.Lock()
	defer fake.listOrgPrivateDomainsMutex.Unlock()
	fake.ListOrgPrivateDomainsStub = stub
}

func (fake *FakeCFClient) ListOrgPrivateDomainsArgsForCall(i int) string {
	fake.listOrgPrivateDomainsMutex.RLock()
	defer fake.listOrgPrivateDomainsMutex.RUnlock()
	argsForCall := fake.listOrgPrivateDomainsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) ListOrgPrivateDomainsReturns(result1 []cfclient.Domain, result2 error) {
	fake.listOrgPrivateDomainsMutex.Lock()
	defer fake.listOrgPrivateDomainsMutex.Unlock()
	fake.ListOrgPrivateDomainsStub = nil
	fake.listOrgPrivateDomainsReturns = struct {
		result1 []cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListOrgPrivateDomainsReturnsOnCall(i int, result1 []cfclient.Domain, result2 error) {
	fake.listOrgPrivateDomainsMutex.Lock()
	defer fake.listOrgPrivateDomainsMutex.Unlock()
	fake.ListOrgPrivateDomainsStub = nil
	if fake.listOrgPrivateDomainsReturnsOnCall == nil {
		fake.listOrgPrivateDomainsReturnsOnCall = make(map[int]struct {
			result1 []cfclient.Domain
			result2 error
		})
	}
	fake.listOrgPrivateDomainsReturnsOnCall[i] = struct {
		result1 []cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListOrgs() ([]cfclient.Org, error) {
	fake.listOrgsMutex.Lock()
	ret, specificReturn := fake.listOrgsReturnsOnCall[len(fake.listOrgsArgsForCall)]
	fake.listOrgsArgsForCall = append(fake.listOrgsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListOrgs", []interface{}{})
	fake.listOrgsMutex.Unlock()
	if fake.ListOrgsStub != nil {
		return fake.ListOrgsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listOrgsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) ListOrgsCallCount() int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return len(fake.listOrgsArgsForCall)
}

func (fake *FakeCFClient) ListOrgsCalls(stub func() ([]cfclient.Org, error)) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = stub
}

func (fake *FakeCFClient) ListOrgsReturns(result1 []cfclient.Org, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	fake.listOrgsReturns = struct {
		result1 []cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListOrgsReturnsOnCall(i int, result1 []cfclient.Org, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	if fake.listOrgsReturnsOnCall == nil {
		fake.listOrgsReturnsOnCall = make(map[int]struct {
			result1 []cfclient.Org
			result2 error
		})
	}
	fake.listOrgsReturnsOnCall[i] = struct {
		result1 []cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ShareOrgPrivateDomain(arg1 string, arg2 string) (*cfclient.Domain, error) {
	fake.shareOrgPrivateDomainMutex.Lock()
	ret, specificReturn := fake.shareOrgPrivateDomainReturnsOnCall[len(fake.shareOrgPrivateDomainArgsForCall)]
	fake.shareOrgPrivateDomainArgsForCall = append(fake.shareOrgPrivateDomainArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ShareOrgPrivateDomain", []interface{}{arg1, arg2})
	fake.shareOrgPrivateDomainMutex.Unlock()
	if fake.ShareOrgPrivateDomainStub != nil {
		return fake.ShareOrgPrivateDomainStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.shareOrgPrivateDomainReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) ShareOrgPrivateDomainCallCount() int {
	fake.shareOrgPrivateDomainMutex.RLock()
	defer fake.shareOrgPrivateDomainMutex.RUnlock()
	return len(fake.shareOrgPrivateDomainArgsForCall)
}

func (fake *FakeCFClient) ShareOrgPrivateDomainCalls(stub func(string, string) (*cfclient.Domain, error)) {
	fake.shareOrgPrivateDomainMutex.Lock()
	defer fake.shareOrgPrivateDomainMutex.Unlock()
	fake.ShareOrgPrivateDomainStub = stub
}

func (fake *FakeCFClient) ShareOrgPrivateDomainArgsForCall(i int) (string, string) {
	fake.shareOrgPrivateDomainMutex.RLock()
	defer fake.shareOrgPrivateDomainMutex.RUnlock()
	argsForCall := fake.shareOrgPrivateDomainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFClient) ShareOrgPrivateDomainReturns(result1 *cfclient.Domain, result2 error) {
	fake.shareOrgPrivateDomainMutex.Lock()
	defer fake.shareOrgPrivateDomainMutex.Unlock()
	fake.ShareOrgPrivateDomainStub = nil
	fake.shareOrgPrivateDomainReturns = struct {
		result1 *cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ShareOrgPrivateDomainReturnsOnCall(i int, result1 *cfclient.Domain, result2 error) {
	fake.shareOrgPrivateDomainMutex.Lock()
	defer fake.shareOrgPrivateDomainMutex.Unlock()
	fake.ShareOrgPrivateDomainStub = nil
	if fake.shareOrgPrivateDomainReturnsOnCall == nil {
		fake.shareOrgPrivateDomainReturnsOnCall = make(map[int]struct {
			result1 *cfclient.Domain
			result2 error
		})
	}
	fake.shareOrgPrivateDomainReturnsOnCall[i] = struct {
		result1 *cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) SupportsMetadataAPI() (bool, error) {
	fake.supportsMetadataAPIMutex.Lock()
	ret, specificReturn := fake.supportsMetadataAPIReturnsOnCall[len(fake.supportsMetadataAPIArgsForCall)]
	fake.supportsMetadataAPIArgsForCall = append(fake.supportsMetadataAPIArgsForCall, struct {
	}{})
	fake.recordInvocation("SupportsMetadataAPI", []interface{}{})
	fake.supportsMetadataAPIMutex.Unlock()
	if fake.SupportsMetadataAPIStub != nil {
		return fake.SupportsMetadataAPIStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.supportsMetadataAPIReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) SupportsMetadataAPICallCount() int {
	fake.supportsMetadataAPIMutex.RLock()
	defer fake.supportsMetadataAPIMutex.RUnlock()
	return len(fake.supportsMetadataAPIArgsForCall)
}

func (fake *FakeCFClient) SupportsMetadataAPICalls(stub func() (bool, error)) {
	fake.supportsMetadataAPIMutex.Lock()
	defer fake.supportsMetadataAPIMutex.Unlock()
	fake.SupportsMetadataAPIStub = stub
}

func (fake *FakeCFClient) SupportsMetadataAPIReturns(result1 bool, result2 error) {
	fake.supportsMetadataAPIMutex.Lock()
	defer fake.supportsMetadataAPIMutex.Unlock()
	fake.SupportsMetadataAPIStub = nil
	fake.supportsMetadataAPIReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) SupportsMetadataAPIReturnsOnCall(i int, result1 bool, result2 error) {
	fake.supportsMetadataAPIMutex.Lock()
	defer fake.supportsMetadataAPIMutex.Unlock()
	fake.SupportsMetadataAPIStub = nil
	if fake.supportsMetadataAPIReturnsOnCall == nil {
		fake.supportsMetadataAPIReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.supportsMetadataAPIReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) UnshareOrgPrivateDomain(arg1 string, arg2 string) error {
	fake.unshareOrgPrivateDomainMutex.Lock()
	ret, specificReturn := fake.unshareOrgPrivateDomainReturnsOnCall[len(fake.unshareOrgPrivateDomainArgsForCall)]
	fake.unshareOrgPrivateDomainArgsForCall = append(fake.unshareOrgPrivateDomainArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UnshareOrgPrivateDomain", []interface{}{arg1, arg2})
	fake.unshareOrgPrivateDomainMutex.Unlock()
	if fake.UnshareOrgPrivateDomainStub != nil {
		return fake.UnshareOrgPrivateDomainStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unshareOrgPrivateDomainReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) UnshareOrgPrivateDomainCallCount() int {
	fake.unshareOrgPrivateDomainMutex.RLock()
	defer fake.unshareOrgPrivateDomainMutex.RUnlock()
	return len(fake.unshareOrgPrivateDomainArgsForCall)
}

func (fake *FakeCFClient) UnshareOrgPrivateDomainCalls(stub func(string, string) error) {
	fake.unshareOrgPrivateDomainMutex.Lock()
	defer fake.unshareOrgPrivateDomainMutex.Unlock()
	fake.UnshareOrgPrivateDomainStub = stub
}

func (fake *FakeCFClient) UnshareOrgPrivateDomainArgsForCall(i int) (string, string) {
	fake.unshareOrgPrivateDomainMutex.RLock()
	defer fake.unshareOrgPrivateDomainMutex.RUnlock()
	argsForCall := fake.unshareOrgPrivateDomainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFClient) UnshareOrgPrivateDomainReturns(result1 error) {
	fake.unshareOrgPrivateDomainMutex.Lock()
	defer fake.unshareOrgPrivateDomainMutex.Unlock()
	fake.UnshareOrgPrivateDomainStub = nil
	fake.unshareOrgPrivateDomainReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UnshareOrgPrivateDomainReturnsOnCall(i int, result1 error) {
	fake.unshareOrgPrivateDomainMutex.Lock()
	defer fake.unshareOrgPrivateDomainMutex.Unlock()
	fake.UnshareOrgPrivateDomainStub = nil
	if fake.unshareOrgPrivateDomainReturnsOnCall == nil {
		fake.unshareOrgPrivateDomainReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unshareOrgPrivateDomainReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UpdateOrg(arg1 string, arg2 cfclient.OrgRequest) (cfclient.Org, error) {
	fake.updateOrgMutex.Lock()
	ret, specificReturn := fake.updateOrgReturnsOnCall[len(fake.updateOrgArgsForCall)]
	fake.updateOrgArgsForCall = append(fake.updateOrgArgsForCall, struct {
		arg1 string
		arg2 cfclient.OrgRequest
	}{arg1, arg2})
	fake.recordInvocation("UpdateOrg", []interface{}{arg1, arg2})
	fake.updateOrgMutex.Unlock()
	if fake.UpdateOrgStub != nil {
		return fake.UpdateOrgStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateOrgReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) UpdateOrgCallCount() int {
	fake.updateOrgMutex.RLock()
	defer fake.updateOrgMutex.RUnlock()
	return len(fake.updateOrgArgsForCall)
}

func (fake *FakeCFClient) UpdateOrgCalls(stub func(string, cfclient.OrgRequest) (cfclient.Org, error)) {
	fake.updateOrgMutex.Lock()
	defer fake.updateOrgMutex.Unlock()
	fake.UpdateOrgStub = stub
}

func (fake *FakeCFClient) UpdateOrgArgsForCall(i int) (string, cfclient.OrgRequest) {
	fake.updateOrgMutex.RLock()
	defer fake.updateOrgMutex.RUnlock()
	argsForCall := fake.updateOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFClient) UpdateOrgReturns(result1 cfclient.Org, result2 error) {
	fake.updateOrgMutex.Lock()
	defer fake.updateOrgMutex.Unlock()
	fake.UpdateOrgStub = nil
	fake.updateOrgReturns = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) UpdateOrgReturnsOnCall(i int, result1 cfclient.Org, result2 error) {
	fake.updateOrgMutex.Lock()
	defer fake.updateOrgMutex.Unlock()
	fake.UpdateOrgStub = nil
	if fake.updateOrgReturnsOnCall == nil {
		fake.updateOrgReturnsOnCall = make(map[int]struct {
			result1 cfclient.Org
			result2 error
		})
	}
	fake.updateOrgReturnsOnCall[i] = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) UpdateOrgMetadata(arg1 string, arg2 cfclient.Metadata) error {
	fake.updateOrgMetadataMutex.Lock()
	ret, specificReturn := fake.updateOrgMetadataReturnsOnCall[len(fake.updateOrgMetadataArgsForCall)]
	fake.updateOrgMetadataArgsForCall = append(fake.updateOrgMetadataArgsForCall, struct {
		arg1 string
		arg2 cfclient.Metadata
	}{arg1, arg2})
	fake.recordInvocation("UpdateOrgMetadata", []interface{}{arg1, arg2})
	fake.updateOrgMetadataMutex.Unlock()
	if fake.UpdateOrgMetadataStub != nil {
		return fake.UpdateOrgMetadataStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateOrgMetadataReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) UpdateOrgMetadataCallCount() int {
	fake.updateOrgMetadataMutex.RLock()
	defer fake.updateOrgMetadataMutex.RUnlock()
	return len(fake.updateOrgMetadataArgsForCall)
}

func (fake *FakeCFClient) UpdateOrgMetadataCalls(stub func(string, cfclient.Metadata) error) {
	fake.updateOrgMetadataMutex.Lock()
	defer fake.updateOrgMetadataMutex.Unlock()
	fake.UpdateOrgMetadataStub = stub
}

func (fake *FakeCFClient) UpdateOrgMetadataArgsForCall(i int) (string, cfclient.Metadata) {
	fake.updateOrgMetadataMutex.RLock()
	defer fake.updateOrgMetadataMutex.RUnlock()
	argsForCall := fake.updateOrgMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFClient) UpdateOrgMetadataReturns(result1 error) {
	fake.updateOrgMetadataMutex.Lock()
	defer fake.updateOrgMetadataMutex.Unlock()
	fake.UpdateOrgMetadataStub = nil
	fake.updateOrgMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UpdateOrgMetadataReturnsOnCall(i int, result1 error) {
	fake.updateOrgMetadataMutex.Lock()
	defer fake.updateOrgMetadataMutex.Unlock()
	fake.UpdateOrgMetadataStub = nil
	if fake.updateOrgMetadataReturnsOnCall == nil {
		fake.updateOrgMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateOrgMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createDomainMutex.RLock()
	defer fake.createDomainMutex.RUnlock()
	fake.createOrgMutex.RLock()
	defer fake.createOrgMutex.RUnlock()
	fake.deleteDomainMutex.RLock()
	defer fake.deleteDomainMutex.RUnlock()
	fake.deleteOrgMutex.RLock()
	defer fake.deleteOrgMutex.RUnlock()
	fake.getOrgByGuidMutex.RLock()
	defer fake.getOrgByGuidMutex.RUnlock()
	fake.listDomainsMutex.RLock()
	defer fake.listDomainsMutex.RUnlock()
	fake.listOrgPrivateDomainsMutex.RLock()
	defer fake.listOrgPrivateDomainsMutex.RUnlock()
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	fake.shareOrgPrivateDomainMutex.RLock()
	defer fake.shareOrgPrivateDomainMutex.RUnlock()
	fake.supportsMetadataAPIMutex.RLock()
	defer fake.supportsMetadataAPIMutex.RUnlock()
	fake.unshareOrgPrivateDomainMutex.RLock()
	defer fake.unshareOrgPrivateDomainMutex.RUnlock()
	fake.updateOrgMutex.RLock()
	defer fake.updateOrgMutex.RUnlock()
	fake.updateOrgMetadataMutex.RLock()
	defer fake.updateOrgMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFClient) recordInvocation(key string, args []interface{}) {
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

var _ organization.CFClient = new(FakeCFClient)
