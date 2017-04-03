// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/pivotalservices/cf-mgmt/uaac (interfaces: Manager)

package mock_uaac

import (
	gomock "github.com/golang/mock/gomock"
	uaac "github.com/pivotalservices/cf-mgmt/uaac"
)

// Mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *_MockManagerRecorder
}

// Recorder for MockManager (not exported)
type _MockManagerRecorder struct {
	mock *MockManager
}

func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &_MockManagerRecorder{mock}
	return mock
}

func (_m *MockManager) EXPECT() *_MockManagerRecorder {
	return _m.recorder
}

func (_m *MockManager) CreateExternalUser(_param0 string, _param1 string, _param2 string, _param3 string) error {
	ret := _m.ctrl.Call(_m, "CreateExternalUser", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) CreateExternalUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateExternalUser", arg0, arg1, arg2, arg3)
}

func (_m *MockManager) ListUsers() (map[string]string, error) {
	ret := _m.ctrl.Call(_m, "ListUsers")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) ListUsers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListUsers")
}

func (_m *MockManager) UsersByID() (map[string]uaac.User, error) {
	ret := _m.ctrl.Call(_m, "UsersByID")
	ret0, _ := ret[0].(map[string]uaac.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) UsersByID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UsersByID")
}
