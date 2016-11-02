// Automatically generated by MockGen. DO NOT EDIT!
// Source: handler/store/instancestore.go

package mocks

import (
	context "context"
	types0 "github.com/aws/amazon-ecs-event-stream-handler/handler/store/types"
	types "github.com/aws/amazon-ecs-event-stream-handler/handler/types"
	gomock "github.com/golang/mock/gomock"
)

// Mock of ContainerInstanceStore interface
type MockContainerInstanceStore struct {
	ctrl     *gomock.Controller
	recorder *_MockContainerInstanceStoreRecorder
}

// Recorder for MockContainerInstanceStore (not exported)
type _MockContainerInstanceStoreRecorder struct {
	mock *MockContainerInstanceStore
}

func NewMockContainerInstanceStore(ctrl *gomock.Controller) *MockContainerInstanceStore {
	mock := &MockContainerInstanceStore{ctrl: ctrl}
	mock.recorder = &_MockContainerInstanceStoreRecorder{mock}
	return mock
}

func (_m *MockContainerInstanceStore) EXPECT() *_MockContainerInstanceStoreRecorder {
	return _m.recorder
}

func (_m *MockContainerInstanceStore) AddContainerInstance(instance string) error {
	ret := _m.ctrl.Call(_m, "AddContainerInstance", instance)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockContainerInstanceStoreRecorder) AddContainerInstance(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddContainerInstance", arg0)
}

func (_m *MockContainerInstanceStore) GetContainerInstance(cluster string, instanceARN string) (*types.ContainerInstance, error) {
	ret := _m.ctrl.Call(_m, "GetContainerInstance", cluster, instanceARN)
	ret0, _ := ret[0].(*types.ContainerInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockContainerInstanceStoreRecorder) GetContainerInstance(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetContainerInstance", arg0, arg1)
}

func (_m *MockContainerInstanceStore) ListContainerInstances() ([]types.ContainerInstance, error) {
	ret := _m.ctrl.Call(_m, "ListContainerInstances")
	ret0, _ := ret[0].([]types.ContainerInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockContainerInstanceStoreRecorder) ListContainerInstances() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainerInstances")
}

func (_m *MockContainerInstanceStore) FilterContainerInstances(filterKey string, filterValue string) ([]types.ContainerInstance, error) {
	ret := _m.ctrl.Call(_m, "FilterContainerInstances", filterKey, filterValue)
	ret0, _ := ret[0].([]types.ContainerInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockContainerInstanceStoreRecorder) FilterContainerInstances(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilterContainerInstances", arg0, arg1)
}

func (_m *MockContainerInstanceStore) StreamContainerInstances(ctx context.Context) (chan types0.ContainerInstanceErrorWrapper, error) {
	ret := _m.ctrl.Call(_m, "StreamContainerInstances", ctx)
	ret0, _ := ret[0].(chan types0.ContainerInstanceErrorWrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockContainerInstanceStoreRecorder) StreamContainerInstances(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StreamContainerInstances", arg0)
}