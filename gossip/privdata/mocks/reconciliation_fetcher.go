// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"

import privdata_common "github.com/hyperledger/udo/gossip/privdata/common"

// ReconciliationFetcher is an autogenerated mock type for the ReconciliationFetcher type
type ReconciliationFetcher struct {
	mock.Mock
}

// FetchReconciledItems provides a mock function with given fields: dig2collectionConfig
func (_m *ReconciliationFetcher) FetchReconciledItems(dig2collectionConfig privdata_common.Dig2CollectionConfig) (*privdata_common.FetchedPvtDataContainer, error) {
	ret := _m.Called(dig2collectionConfig)

	var r0 *privdata_common.FetchedPvtDataContainer
	if rf, ok := ret.Get(0).(func(privdata_common.Dig2CollectionConfig) *privdata_common.FetchedPvtDataContainer); ok {
		r0 = rf(dig2collectionConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*privdata_common.FetchedPvtDataContainer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(privdata_common.Dig2CollectionConfig) error); ok {
		r1 = rf(dig2collectionConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
