// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extensions/pkg/gardener/terraformer (interfaces: Interface,Initializer,Factory)

// Package terraformer is a generated GoMock package.
package terraformer

import (
	terraformer "github.com/gardener/gardener-extensions/pkg/gardener/terraformer"
	terraformer0 "github.com/gardener/gardener/pkg/operation/terraformer"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method
func (m *MockInterface) Apply() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply")
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockInterfaceMockRecorder) Apply() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockInterface)(nil).Apply))
}

// ConfigExists mocks base method
func (m *MockInterface) ConfigExists() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigExists")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigExists indicates an expected call of ConfigExists
func (mr *MockInterfaceMockRecorder) ConfigExists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigExists", reflect.TypeOf((*MockInterface)(nil).ConfigExists))
}

// Destroy mocks base method
func (m *MockInterface) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockInterfaceMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockInterface)(nil).Destroy))
}

// GetStateOutputVariables mocks base method
func (m *MockInterface) GetStateOutputVariables(arg0 ...string) (map[string]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStateOutputVariables", varargs...)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateOutputVariables indicates an expected call of GetStateOutputVariables
func (mr *MockInterfaceMockRecorder) GetStateOutputVariables(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateOutputVariables", reflect.TypeOf((*MockInterface)(nil).GetStateOutputVariables), arg0...)
}

// InitializeWith mocks base method
func (m *MockInterface) InitializeWith(arg0 terraformer.Initializer) terraformer.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeWith", arg0)
	ret0, _ := ret[0].(terraformer.Interface)
	return ret0
}

// InitializeWith indicates an expected call of InitializeWith
func (mr *MockInterfaceMockRecorder) InitializeWith(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeWith", reflect.TypeOf((*MockInterface)(nil).InitializeWith), arg0)
}

// SetVariablesEnvironment mocks base method
func (m *MockInterface) SetVariablesEnvironment(arg0 map[string]string) terraformer.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVariablesEnvironment", arg0)
	ret0, _ := ret[0].(terraformer.Interface)
	return ret0
}

// SetVariablesEnvironment indicates an expected call of SetVariablesEnvironment
func (mr *MockInterfaceMockRecorder) SetVariablesEnvironment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVariablesEnvironment", reflect.TypeOf((*MockInterface)(nil).SetVariablesEnvironment), arg0)
}

// MockInitializer is a mock of Initializer interface
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockInitializerMockRecorder
}

// MockInitializerMockRecorder is the mock recorder for MockInitializer
type MockInitializerMockRecorder struct {
	mock *MockInitializer
}

// NewMockInitializer creates a new mock instance
func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &MockInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return m.recorder
}

// Initialize mocks base method
func (m *MockInitializer) Initialize(arg0 *terraformer0.InitializerConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockInitializerMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockInitializer)(nil).Initialize), arg0)
}

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// DefaultInitializer mocks base method
func (m *MockFactory) DefaultInitializer(arg0 client.Client, arg1, arg2 string, arg3 []byte) terraformer.Initializer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultInitializer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(terraformer.Initializer)
	return ret0
}

// DefaultInitializer indicates an expected call of DefaultInitializer
func (mr *MockFactoryMockRecorder) DefaultInitializer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultInitializer", reflect.TypeOf((*MockFactory)(nil).DefaultInitializer), arg0, arg1, arg2, arg3)
}

// New mocks base method
func (m *MockFactory) New(arg0 logrus.FieldLogger, arg1 client.Client, arg2 v1.CoreV1Interface, arg3, arg4, arg5, arg6 string) terraformer.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(terraformer.Interface)
	return ret0
}

// New indicates an expected call of New
func (mr *MockFactoryMockRecorder) New(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockFactory)(nil).New), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// NewForConfig mocks base method
func (m *MockFactory) NewForConfig(arg0 logrus.FieldLogger, arg1 *rest.Config, arg2, arg3, arg4, arg5 string) (terraformer.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewForConfig", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(terraformer.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewForConfig indicates an expected call of NewForConfig
func (mr *MockFactoryMockRecorder) NewForConfig(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewForConfig", reflect.TypeOf((*MockFactory)(nil).NewForConfig), arg0, arg1, arg2, arg3, arg4, arg5)
}
