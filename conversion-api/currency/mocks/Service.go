// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "conversion-api/models"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateCurrency provides a mock function with given fields: ctx, _a1
func (_m *Service) CreateCurrency(ctx context.Context, _a1 models.Currency) (models.Currency, error) {
	ret := _m.Called(ctx, _a1)

	var r0 models.Currency
	if rf, ok := ret.Get(0).(func(context.Context, models.Currency) models.Currency); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(models.Currency)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Currency) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCurrency provides a mock function with given fields: ctx, _a1
func (_m *Service) ExchangeCurrency(ctx context.Context, _a1 models.Currency) (models.CurrencyExchange, error) {
	ret := _m.Called(ctx, _a1)

	var r0 models.CurrencyExchange
	if rf, ok := ret.Get(0).(func(context.Context, models.Currency) models.CurrencyExchange); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(models.CurrencyExchange)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Currency) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCurrency provides a mock function with given fields: ctx, _a1
func (_m *Service) UpdateCurrency(ctx context.Context, _a1 models.Currency) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Currency) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
