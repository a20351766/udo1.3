// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import etcdraft "github.com/hyperledger/udo/orderer/consensus/etcdraft"
import mock "github.com/stretchr/testify/mock"

// ReceiverGetter is an autogenerated mock type for the ReceiverGetter type
type ReceiverGetter struct {
	mock.Mock
}

// ReceiverByChain provides a mock function with given fields: channelID
func (_m *ReceiverGetter) ReceiverByChain(channelID string) etcdraft.MessageReceiver {
	ret := _m.Called(channelID)

	var r0 etcdraft.MessageReceiver
	if rf, ok := ret.Get(0).(func(string) etcdraft.MessageReceiver); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(etcdraft.MessageReceiver)
		}
	}

	return r0
}
