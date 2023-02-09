// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/networking/v1alpha1 (interfaces: NetworkingV1alpha1Interface)

// Package v1alpha1 is a generated GoMock package.
package v1alpha1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
	rest "k8s.io/client-go/rest"
)

// MockNetworkingV1alpha1Interface is a mock of NetworkingV1alpha1Interface interface.
type MockNetworkingV1alpha1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkingV1alpha1InterfaceMockRecorder
}

// MockNetworkingV1alpha1InterfaceMockRecorder is the mock recorder for MockNetworkingV1alpha1Interface.
type MockNetworkingV1alpha1InterfaceMockRecorder struct {
	mock *MockNetworkingV1alpha1Interface
}

// NewMockNetworkingV1alpha1Interface creates a new mock instance.
func NewMockNetworkingV1alpha1Interface(ctrl *gomock.Controller) *MockNetworkingV1alpha1Interface {
	mock := &MockNetworkingV1alpha1Interface{ctrl: ctrl}
	mock.recorder = &MockNetworkingV1alpha1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkingV1alpha1Interface) EXPECT() *MockNetworkingV1alpha1InterfaceMockRecorder {
	return m.recorder
}

// ClusterCIDRs mocks base method.
func (m *MockNetworkingV1alpha1Interface) ClusterCIDRs() v1alpha1.ClusterCIDRInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCIDRs")
	ret0, _ := ret[0].(v1alpha1.ClusterCIDRInterface)
	return ret0
}

// ClusterCIDRs indicates an expected call of ClusterCIDRs.
func (mr *MockNetworkingV1alpha1InterfaceMockRecorder) ClusterCIDRs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCIDRs", reflect.TypeOf((*MockNetworkingV1alpha1Interface)(nil).ClusterCIDRs))
}

// RESTClient mocks base method.
func (m *MockNetworkingV1alpha1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockNetworkingV1alpha1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockNetworkingV1alpha1Interface)(nil).RESTClient))
}