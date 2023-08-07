// Code generated by MockGen. DO NOT EDIT.
// Source: go-task/internal/auths/usecase (interfaces: Authusecase)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	domain "go-task/domain/auths/request"
	domain0 "go-task/domain/auths/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockAuthusecase is a mock of Authusecase interface.
type MockAuthusecase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthusecaseMockRecorder
}

// MockAuthusecaseMockRecorder is the mock recorder for MockAuthusecase.
type MockAuthusecaseMockRecorder struct {
	mock *MockAuthusecase
}

// NewMockAuthusecase creates a new mock instance.
func NewMockAuthusecase(ctrl *gomock.Controller) *MockAuthusecase {
	mock := &MockAuthusecase{ctrl: ctrl}
	mock.recorder = &MockAuthusecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthusecase) EXPECT() *MockAuthusecaseMockRecorder {
	return m.recorder
}

// GetUserByEmail mocks base method.
func (m *MockAuthusecase) GetUserByEmail(arg0 string) (domain0.RegisterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(domain0.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockAuthusecaseMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockAuthusecase)(nil).GetUserByEmail), arg0)
}

// GetUserByID mocks base method.
func (m *MockAuthusecase) GetUserByID(arg0 uuid.UUID) (domain0.RegisterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0)
	ret0, _ := ret[0].(domain0.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockAuthusecaseMockRecorder) GetUserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockAuthusecase)(nil).GetUserByID), arg0)
}

// LoginUser mocks base method.
func (m *MockAuthusecase) LoginUser(arg0 domain.LoginModel) (domain0.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", arg0)
	ret0, _ := ret[0].(domain0.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockAuthusecaseMockRecorder) LoginUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockAuthusecase)(nil).LoginUser), arg0)
}

// Register mocks base method.
func (m *MockAuthusecase) Register(arg0 domain.RegisterModel) (domain0.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(domain0.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthusecaseMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthusecase)(nil).Register), arg0)
}
