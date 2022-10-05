// Code generated by MockGen. DO NOT EDIT.
// Source: ./get_by_name.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pokemon "github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

// MockGetByName is a mock of GetByName interface.
type MockGetByName struct {
	ctrl     *gomock.Controller
	recorder *MockGetByNameMockRecorder
}

// MockGetByNameMockRecorder is the mock recorder for MockGetByName.
type MockGetByNameMockRecorder struct {
	mock *MockGetByName
}

// NewMockGetByName creates a new mock instance.
func NewMockGetByName(ctrl *gomock.Controller) *MockGetByName {
	mock := &MockGetByName{ctrl: ctrl}
	mock.recorder = &MockGetByNameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetByName) EXPECT() *MockGetByNameMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockGetByName) Get(ctx context.Context, name string) (*pokemon.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, name)
	ret0, _ := ret[0].(*pokemon.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGetByNameMockRecorder) Get(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetByName)(nil).Get), ctx, name)
}
