// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/quintilesims/layer0/common/aws/autoscaling (interfaces: Provider)

package mock_autoscaling

import (
	gomock "github.com/golang/mock/gomock"
	autoscaling "github.com/quintilesims/layer0/common/aws/autoscaling"
)

// Mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *_MockProviderRecorder
}

// Recorder for MockProvider (not exported)
type _MockProviderRecorder struct {
	mock *MockProvider
}

func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &_MockProviderRecorder{mock}
	return mock
}

func (_m *MockProvider) EXPECT() *_MockProviderRecorder {
	return _m.recorder
}

func (_m *MockProvider) AttachLoadBalancer(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "AttachLoadBalancer", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) AttachLoadBalancer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AttachLoadBalancer", arg0, arg1)
}

func (_m *MockProvider) CreateAutoScalingGroup(_param0 string, _param1 string, _param2 string, _param3 int, _param4 int) error {
	ret := _m.ctrl.Call(_m, "CreateAutoScalingGroup", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) CreateAutoScalingGroup(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAutoScalingGroup", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockProvider) CreateLaunchConfiguration(_param0 *string, _param1 *string, _param2 *string, _param3 *string, _param4 *string, _param5 *string, _param6 []*string, _param7 map[string]int) error {
	ret := _m.ctrl.Call(_m, "CreateLaunchConfiguration", _param0, _param1, _param2, _param3, _param4, _param5, _param6, _param7)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) CreateLaunchConfiguration(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateLaunchConfiguration", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (_m *MockProvider) DeleteAutoScalingGroup(_param0 *string) error {
	ret := _m.ctrl.Call(_m, "DeleteAutoScalingGroup", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) DeleteAutoScalingGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAutoScalingGroup", arg0)
}

func (_m *MockProvider) DeleteLaunchConfiguration(_param0 *string) error {
	ret := _m.ctrl.Call(_m, "DeleteLaunchConfiguration", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) DeleteLaunchConfiguration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteLaunchConfiguration", arg0)
}

func (_m *MockProvider) DescribeAutoScalingGroup(_param0 string) (*autoscaling.Group, error) {
	ret := _m.ctrl.Call(_m, "DescribeAutoScalingGroup", _param0)
	ret0, _ := ret[0].(*autoscaling.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeAutoScalingGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeAutoScalingGroup", arg0)
}

func (_m *MockProvider) DescribeAutoScalingGroups(_param0 []*string) ([]*autoscaling.Group, error) {
	ret := _m.ctrl.Call(_m, "DescribeAutoScalingGroups", _param0)
	ret0, _ := ret[0].([]*autoscaling.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeAutoScalingGroups(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeAutoScalingGroups", arg0)
}

func (_m *MockProvider) DescribeLaunchConfiguration(_param0 string) (*autoscaling.LaunchConfiguration, error) {
	ret := _m.ctrl.Call(_m, "DescribeLaunchConfiguration", _param0)
	ret0, _ := ret[0].(*autoscaling.LaunchConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeLaunchConfiguration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeLaunchConfiguration", arg0)
}

func (_m *MockProvider) DescribeLaunchConfigurations(_param0 []*string) ([]*autoscaling.LaunchConfiguration, error) {
	ret := _m.ctrl.Call(_m, "DescribeLaunchConfigurations", _param0)
	ret0, _ := ret[0].([]*autoscaling.LaunchConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) DescribeLaunchConfigurations(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeLaunchConfigurations", arg0)
}

func (_m *MockProvider) SetDesiredCapacity(_param0 string, _param1 int) error {
	ret := _m.ctrl.Call(_m, "SetDesiredCapacity", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) SetDesiredCapacity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDesiredCapacity", arg0, arg1)
}

func (_m *MockProvider) TerminateInstanceInAutoScalingGroup(_param0 string, _param1 bool) (*autoscaling.Activity, error) {
	ret := _m.ctrl.Call(_m, "TerminateInstanceInAutoScalingGroup", _param0, _param1)
	ret0, _ := ret[0].(*autoscaling.Activity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) TerminateInstanceInAutoScalingGroup(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TerminateInstanceInAutoScalingGroup", arg0, arg1)
}

func (_m *MockProvider) UpdateAutoScalingGroupMaxSize(_param0 string, _param1 int) error {
	ret := _m.ctrl.Call(_m, "UpdateAutoScalingGroupMaxSize", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) UpdateAutoScalingGroupMaxSize(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateAutoScalingGroupMaxSize", arg0, arg1)
}

func (_m *MockProvider) UpdateAutoScalingGroupMinSize(_param0 string, _param1 int) error {
	ret := _m.ctrl.Call(_m, "UpdateAutoScalingGroupMinSize", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProviderRecorder) UpdateAutoScalingGroupMinSize(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateAutoScalingGroupMinSize", arg0, arg1)
}
