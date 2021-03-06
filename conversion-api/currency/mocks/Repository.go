// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "conversion-api/models"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetCurrency provides a mock function with given fields: ctx, _a1
func (_m *Repository) GetCurrency(ctx context.Context, _a1 string) (models.Currency, error) {
	ret := _m.Called(ctx, _a1)

	var r0 models.Currency
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Currency); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(models.Currency)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertCurrency provides a mock function with given fields: ctx, _a1
func (_m *Repository) InsertCurrency(ctx context.Context, _a1 models.Currency) (models.Currency, error) {
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

// UpdateBallast provides a mock function with given fields: ctx, _a1
func (_m *Repository) UpdateBallast(ctx context.Context, _a1 models.Currency) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Currency) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
