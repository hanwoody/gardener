// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/extensions/pkg/terraformer (interfaces: Terraformer,Initializer,Factory)

// Package terraformer is a generated GoMock package.
package terraformer

import (
	context "context"
	reflect "reflect"
	time "time"

	terraformer "github.com/gardener/gardener/extensions/pkg/terraformer"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	rest "k8s.io/client-go/rest"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockTerraformer is a mock of Terraformer interface.
type MockTerraformer struct {
	ctrl     *gomock.Controller
	recorder *MockTerraformerMockRecorder
}

// MockTerraformerMockRecorder is the mock recorder for MockTerraformer.
type MockTerraformerMockRecorder struct {
	mock *MockTerraformer
}

// NewMockTerraformer creates a new mock instance.
func NewMockTerraformer(ctrl *gomock.Controller) *MockTerraformer {
	mock := &MockTerraformer{ctrl: ctrl}
	mock.recorder = &MockTerraformerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTerraformer) EXPECT() *MockTerraformerMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockTerraformer) Apply() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply")
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockTerraformerMockRecorder) Apply() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockTerraformer)(nil).Apply))
}

// CleanupConfiguration mocks base method.
func (m *MockTerraformer) CleanupConfiguration(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupConfiguration", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupConfiguration indicates an expected call of CleanupConfiguration.
func (mr *MockTerraformerMockRecorder) CleanupConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupConfiguration", reflect.TypeOf((*MockTerraformer)(nil).CleanupConfiguration), arg0)
}

// ConfigExists mocks base method.
func (m *MockTerraformer) ConfigExists() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigExists")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigExists indicates an expected call of ConfigExists.
func (mr *MockTerraformerMockRecorder) ConfigExists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigExists", reflect.TypeOf((*MockTerraformer)(nil).ConfigExists))
}

// Destroy mocks base method.
func (m *MockTerraformer) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockTerraformerMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockTerraformer)(nil).Destroy))
}

// EnsureCleanedUp mocks base method.
func (m *MockTerraformer) EnsureCleanedUp(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureCleanedUp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureCleanedUp indicates an expected call of EnsureCleanedUp.
func (mr *MockTerraformerMockRecorder) EnsureCleanedUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureCleanedUp", reflect.TypeOf((*MockTerraformer)(nil).EnsureCleanedUp), arg0)
}

// GetRawState mocks base method.
func (m *MockTerraformer) GetRawState(arg0 context.Context) (*terraformer.RawState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawState", arg0)
	ret0, _ := ret[0].(*terraformer.RawState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawState indicates an expected call of GetRawState.
func (mr *MockTerraformerMockRecorder) GetRawState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawState", reflect.TypeOf((*MockTerraformer)(nil).GetRawState), arg0)
}

// GetState mocks base method.
func (m *MockTerraformer) GetState() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState.
func (mr *MockTerraformerMockRecorder) GetState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockTerraformer)(nil).GetState))
}

// GetStateOutputVariables mocks base method.
func (m *MockTerraformer) GetStateOutputVariables(arg0 ...string) (map[string]string, error) {
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

// GetStateOutputVariables indicates an expected call of GetStateOutputVariables.
func (mr *MockTerraformerMockRecorder) GetStateOutputVariables(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateOutputVariables", reflect.TypeOf((*MockTerraformer)(nil).GetStateOutputVariables), arg0...)
}

// InitializeWith mocks base method.
func (m *MockTerraformer) InitializeWith(arg0 terraformer.Initializer) terraformer.Terraformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeWith", arg0)
	ret0, _ := ret[0].(terraformer.Terraformer)
	return ret0
}

// InitializeWith indicates an expected call of InitializeWith.
func (mr *MockTerraformerMockRecorder) InitializeWith(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeWith", reflect.TypeOf((*MockTerraformer)(nil).InitializeWith), arg0)
}

// IsStateEmpty mocks base method.
func (m *MockTerraformer) IsStateEmpty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStateEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStateEmpty indicates an expected call of IsStateEmpty.
func (mr *MockTerraformerMockRecorder) IsStateEmpty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStateEmpty", reflect.TypeOf((*MockTerraformer)(nil).IsStateEmpty))
}

// NumberOfResources mocks base method.
func (m *MockTerraformer) NumberOfResources(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumberOfResources", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NumberOfResources indicates an expected call of NumberOfResources.
func (mr *MockTerraformerMockRecorder) NumberOfResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumberOfResources", reflect.TypeOf((*MockTerraformer)(nil).NumberOfResources), arg0)
}

// SetDeadlineCleaning mocks base method.
func (m *MockTerraformer) SetDeadlineCleaning(arg0 time.Duration) terraformer.Terraformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadlineCleaning", arg0)
	ret0, _ := ret[0].(terraformer.Terraformer)
	return ret0
}

// SetDeadlineCleaning indicates an expected call of SetDeadlineCleaning.
func (mr *MockTerraformerMockRecorder) SetDeadlineCleaning(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadlineCleaning", reflect.TypeOf((*MockTerraformer)(nil).SetDeadlineCleaning), arg0)
}

// SetDeadlinePod mocks base method.
func (m *MockTerraformer) SetDeadlinePod(arg0 time.Duration) terraformer.Terraformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadlinePod", arg0)
	ret0, _ := ret[0].(terraformer.Terraformer)
	return ret0
}

// SetDeadlinePod indicates an expected call of SetDeadlinePod.
func (mr *MockTerraformerMockRecorder) SetDeadlinePod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadlinePod", reflect.TypeOf((*MockTerraformer)(nil).SetDeadlinePod), arg0)
}

// SetLogLevel mocks base method.
func (m *MockTerraformer) SetLogLevel(arg0 string) terraformer.Terraformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLogLevel", arg0)
	ret0, _ := ret[0].(terraformer.Terraformer)
	return ret0
}

// SetLogLevel indicates an expected call of SetLogLevel.
func (mr *MockTerraformerMockRecorder) SetLogLevel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MockTerraformer)(nil).SetLogLevel), arg0)
}

// SetTerminationGracePeriodSeconds mocks base method.
func (m *MockTerraformer) SetTerminationGracePeriodSeconds(arg0 int64) terraformer.Terraformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTerminationGracePeriodSeconds", arg0)
	ret0, _ := ret[0].(terraformer.Terraformer)
	return ret0
}

// SetTerminationGracePeriodSeconds indicates an expected call of SetTerminationGracePeriodSeconds.
func (mr *MockTerraformerMockRecorder) SetTerminationGracePeriodSeconds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTerminationGracePeriodSeconds", reflect.TypeOf((*MockTerraformer)(nil).SetTerminationGracePeriodSeconds), arg0)
}

// SetVariablesEnvironment mocks base method.
func (m *MockTerraformer) SetVariablesEnvironment(arg0 map[string]string) terraformer.Terraformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVariablesEnvironment", arg0)
	ret0, _ := ret[0].(terraformer.Terraformer)
	return ret0
}

// SetVariablesEnvironment indicates an expected call of SetVariablesEnvironment.
func (mr *MockTerraformerMockRecorder) SetVariablesEnvironment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVariablesEnvironment", reflect.TypeOf((*MockTerraformer)(nil).SetVariablesEnvironment), arg0)
}

// UseV2 mocks base method.
func (m *MockTerraformer) UseV2(arg0 bool) terraformer.Terraformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseV2", arg0)
	ret0, _ := ret[0].(terraformer.Terraformer)
	return ret0
}

// UseV2 indicates an expected call of UseV2.
func (mr *MockTerraformerMockRecorder) UseV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseV2", reflect.TypeOf((*MockTerraformer)(nil).UseV2), arg0)
}

// WaitForCleanEnvironment mocks base method.
func (m *MockTerraformer) WaitForCleanEnvironment(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForCleanEnvironment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForCleanEnvironment indicates an expected call of WaitForCleanEnvironment.
func (mr *MockTerraformerMockRecorder) WaitForCleanEnvironment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForCleanEnvironment", reflect.TypeOf((*MockTerraformer)(nil).WaitForCleanEnvironment), arg0)
}

// MockInitializer is a mock of Initializer interface.
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockInitializerMockRecorder
}

// MockInitializerMockRecorder is the mock recorder for MockInitializer.
type MockInitializerMockRecorder struct {
	mock *MockInitializer
}

// NewMockInitializer creates a new mock instance.
func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &MockInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return m.recorder
}

// Initialize mocks base method.
func (m *MockInitializer) Initialize(arg0 *terraformer.InitializerConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockInitializerMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockInitializer)(nil).Initialize), arg0)
}

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// DefaultInitializer mocks base method.
func (m *MockFactory) DefaultInitializer(arg0 client.Client, arg1, arg2 string, arg3 []byte, arg4 terraformer.StateConfigMapInitializer) terraformer.Initializer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultInitializer", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(terraformer.Initializer)
	return ret0
}

// DefaultInitializer indicates an expected call of DefaultInitializer.
func (mr *MockFactoryMockRecorder) DefaultInitializer(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultInitializer", reflect.TypeOf((*MockFactory)(nil).DefaultInitializer), arg0, arg1, arg2, arg3, arg4)
}

// New mocks base method.
func (m *MockFactory) New(arg0 logrus.FieldLogger, arg1 client.Client, arg2 v1.CoreV1Interface, arg3, arg4, arg5, arg6 string) terraformer.Terraformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(terraformer.Terraformer)
	return ret0
}

// New indicates an expected call of New.
func (mr *MockFactoryMockRecorder) New(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockFactory)(nil).New), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// NewForConfig mocks base method.
func (m *MockFactory) NewForConfig(arg0 logrus.FieldLogger, arg1 *rest.Config, arg2, arg3, arg4, arg5 string) (terraformer.Terraformer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewForConfig", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(terraformer.Terraformer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewForConfig indicates an expected call of NewForConfig.
func (mr *MockFactoryMockRecorder) NewForConfig(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewForConfig", reflect.TypeOf((*MockFactory)(nil).NewForConfig), arg0, arg1, arg2, arg3, arg4, arg5)
}
