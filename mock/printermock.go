// Code generated by MockGen. DO NOT EDIT.
// Source: printer/printer.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	issue "github.com/wata727/tflint/issue"
	reflect "reflect"
)

// MockPrinterIF is a mock of PrinterIF interface
type MockPrinterIF struct {
	ctrl     *gomock.Controller
	recorder *MockPrinterIFMockRecorder
}

// MockPrinterIFMockRecorder is the mock recorder for MockPrinterIF
type MockPrinterIFMockRecorder struct {
	mock *MockPrinterIF
}

// NewMockPrinterIF creates a new mock instance
func NewMockPrinterIF(ctrl *gomock.Controller) *MockPrinterIF {
	mock := &MockPrinterIF{ctrl: ctrl}
	mock.recorder = &MockPrinterIFMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrinterIF) EXPECT() *MockPrinterIFMockRecorder {
	return m.recorder
}

// Print mocks base method
func (m *MockPrinterIF) Print(issues []*issue.Issue, format string, quiet bool) {
	m.ctrl.Call(m, "Print", issues, format, quiet)
}

// Print indicates an expected call of Print
func (mr *MockPrinterIFMockRecorder) Print(issues, format, quiet interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockPrinterIF)(nil).Print), issues, format, quiet)
}
