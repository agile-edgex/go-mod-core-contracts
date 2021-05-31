// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"

	responses "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

// SubscriptionClient is an autogenerated mock type for the SubscriptionClient type
type SubscriptionClient struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, reqs
func (_m *SubscriptionClient) Add(ctx context.Context, reqs []requests.AddSubscriptionRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseWithIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.AddSubscriptionRequest) []common.BaseWithIdResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseWithIdResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.AddSubscriptionRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllSubscriptions provides a mock function with given fields: ctx, offset, limit
func (_m *SubscriptionClient) AllSubscriptions(ctx context.Context, offset int, limit int) (responses.MultiSubscriptionsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, offset, limit)

	var r0 responses.MultiSubscriptionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, int) responses.MultiSubscriptionsResponse); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiSubscriptionsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteSubscriptionByName provides a mock function with given fields: ctx, name
func (_m *SubscriptionClient) DeleteSubscriptionByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionByName provides a mock function with given fields: ctx, name
func (_m *SubscriptionClient) SubscriptionByName(ctx context.Context, name string) (responses.SubscriptionResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 responses.SubscriptionResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) responses.SubscriptionResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(responses.SubscriptionResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionsByCategory provides a mock function with given fields: ctx, category, offset, limit
func (_m *SubscriptionClient) SubscriptionsByCategory(ctx context.Context, category string, offset int, limit int) (responses.MultiSubscriptionsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, category, offset, limit)

	var r0 responses.MultiSubscriptionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiSubscriptionsResponse); ok {
		r0 = rf(ctx, category, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiSubscriptionsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, category, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionsByLabel provides a mock function with given fields: ctx, label, offset, limit
func (_m *SubscriptionClient) SubscriptionsByLabel(ctx context.Context, label string, offset int, limit int) (responses.MultiSubscriptionsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, label, offset, limit)

	var r0 responses.MultiSubscriptionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiSubscriptionsResponse); ok {
		r0 = rf(ctx, label, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiSubscriptionsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, label, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionsByReceiver provides a mock function with given fields: ctx, receiver, offset, limit
func (_m *SubscriptionClient) SubscriptionsByReceiver(ctx context.Context, receiver string, offset int, limit int) (responses.MultiSubscriptionsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, receiver, offset, limit)

	var r0 responses.MultiSubscriptionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiSubscriptionsResponse); ok {
		r0 = rf(ctx, receiver, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiSubscriptionsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, receiver, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, reqs
func (_m *SubscriptionClient) Update(ctx context.Context, reqs []requests.UpdateSubscriptionRequest) ([]common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.UpdateSubscriptionRequest) []common.BaseResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.UpdateSubscriptionRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}
