// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/idirall22/crypto_app/account/service/model"
)

// IEventStore is an autogenerated mock type for the IEventStore type
type IEventStore struct {
	mock.Mock
}

// PublishEmail provides a mock function with given fields: ctx, queueName, emailChan
func (_m *IEventStore) PublishEmail(ctx context.Context, queueName string, emailChan <-chan model.EmailEvent) error {
	ret := _m.Called(ctx, queueName, emailChan)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, <-chan model.EmailEvent) error); ok {
		r0 = rf(ctx, queueName, emailChan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishNotification provides a mock function with given fields: ctx, queueName, emailChan
func (_m *IEventStore) PublishNotification(ctx context.Context, queueName string, emailChan <-chan model.NotificationEvent) error {
	ret := _m.Called(ctx, queueName, emailChan)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, <-chan model.NotificationEvent) error); ok {
		r0 = rf(ctx, queueName, emailChan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}