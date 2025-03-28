// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	metrics "github.com/jaegertracing/jaeger/internal/metrics"
	dependencystore "github.com/jaegertracing/jaeger/internal/storage/v1/api/dependencystore"

	mock "github.com/stretchr/testify/mock"

	spanstore "github.com/jaegertracing/jaeger/internal/storage/v1/api/spanstore"

	zap "go.uber.org/zap"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// CreateDependencyReader provides a mock function with no fields
func (_m *Factory) CreateDependencyReader() (dependencystore.Reader, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateDependencyReader")
	}

	var r0 dependencystore.Reader
	var r1 error
	if rf, ok := ret.Get(0).(func() (dependencystore.Reader, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() dependencystore.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dependencystore.Reader)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSpanReader provides a mock function with no fields
func (_m *Factory) CreateSpanReader() (spanstore.Reader, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateSpanReader")
	}

	var r0 spanstore.Reader
	var r1 error
	if rf, ok := ret.Get(0).(func() (spanstore.Reader, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() spanstore.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spanstore.Reader)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSpanWriter provides a mock function with no fields
func (_m *Factory) CreateSpanWriter() (spanstore.Writer, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateSpanWriter")
	}

	var r0 spanstore.Writer
	var r1 error
	if rf, ok := ret.Get(0).(func() (spanstore.Writer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() spanstore.Writer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spanstore.Writer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: metricsFactory, logger
func (_m *Factory) Initialize(metricsFactory metrics.Factory, logger *zap.Logger) error {
	ret := _m.Called(metricsFactory, logger)

	if len(ret) == 0 {
		panic("no return value specified for Initialize")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(metrics.Factory, *zap.Logger) error); ok {
		r0 = rf(metricsFactory, logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFactory creates a new instance of Factory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *Factory {
	mock := &Factory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
