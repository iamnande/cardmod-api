// Code generated by MockGen. DO NOT EDIT.
// Source: calculation.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockCalculation is a mock of Calculation interface.
type MockCalculation struct {
	ctrl     *gomock.Controller
	recorder *MockCalculationMockRecorder
}

// MockCalculationMockRecorder is the mock recorder for MockCalculation.
type MockCalculationMockRecorder struct {
	mock *MockCalculation
}

// NewMockCalculation creates a new mock instance.
func NewMockCalculation(ctrl *gomock.Controller) *MockCalculation {
	mock := &MockCalculation{ctrl: ctrl}
	mock.recorder = &MockCalculationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculation) EXPECT() *MockCalculationMockRecorder {
	return m.recorder
}

// CardID mocks base method.
func (m *MockCalculation) CardID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CardID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// CardID indicates an expected call of CardID.
func (mr *MockCalculationMockRecorder) CardID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CardID", reflect.TypeOf((*MockCalculation)(nil).CardID))
}

// CardRatio mocks base method.
func (m *MockCalculation) CardRatio() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CardRatio")
	ret0, _ := ret[0].(int32)
	return ret0
}

// CardRatio indicates an expected call of CardRatio.
func (mr *MockCalculationMockRecorder) CardRatio() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CardRatio", reflect.TypeOf((*MockCalculation)(nil).CardRatio))
}

// ID mocks base method.
func (m *MockCalculation) ID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockCalculationMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockCalculation)(nil).ID))
}

// MagicID mocks base method.
func (m *MockCalculation) MagicID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MagicID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// MagicID indicates an expected call of MagicID.
func (mr *MockCalculationMockRecorder) MagicID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MagicID", reflect.TypeOf((*MockCalculation)(nil).MagicID))
}

// MagicRatio mocks base method.
func (m *MockCalculation) MagicRatio() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MagicRatio")
	ret0, _ := ret[0].(int32)
	return ret0
}

// MagicRatio indicates an expected call of MagicRatio.
func (mr *MockCalculationMockRecorder) MagicRatio() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MagicRatio", reflect.TypeOf((*MockCalculation)(nil).MagicRatio))
}
