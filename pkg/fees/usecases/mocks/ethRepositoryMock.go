// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rsmarincu/glassnode/pkg/fees/usecases (interfaces: ETHRepository)

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repository "github.com/rsmarincu/glassnode/pkg/fees/repository"
)

// MockETHRepository is a mock of ETHRepository interface.
type MockETHRepository struct {
	ctrl     *gomock.Controller
	recorder *MockETHRepositoryMockRecorder
}

// MockETHRepositoryMockRecorder is the mock recorder for MockETHRepository.
type MockETHRepositoryMockRecorder struct {
	mock *MockETHRepository
}

// NewMockETHRepository creates a new mock instance.
func NewMockETHRepository(ctrl *gomock.Controller) *MockETHRepository {
	mock := &MockETHRepository{ctrl: ctrl}
	mock.recorder = &MockETHRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockETHRepository) EXPECT() *MockETHRepositoryMockRecorder {
	return m.recorder
}

// QueryEOATransactions mocks base method.
func (m *MockETHRepository) QueryEOATransactions(arg0 context.Context) ([]*repository.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryEOATransactions", arg0)
	ret0, _ := ret[0].([]*repository.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryEOATransactions indicates an expected call of QueryEOATransactions.
func (mr *MockETHRepositoryMockRecorder) QueryEOATransactions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryEOATransactions", reflect.TypeOf((*MockETHRepository)(nil).QueryEOATransactions), arg0)
}
