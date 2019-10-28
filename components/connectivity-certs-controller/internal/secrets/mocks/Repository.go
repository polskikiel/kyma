// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Get provides a mock function with given fields: name
func (_m *Repository) Get(name types.NamespacedName) (map[string][]byte, error) {
	ret := _m.Called(name)

	var r0 map[string][]byte
	if rf, ok := ret.Get(0).(func(types.NamespacedName) map[string][]byte); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.NamespacedName) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertWithMerge provides a mock function with given fields: name, data
func (_m *Repository) UpsertWithMerge(name types.NamespacedName, data map[string][]byte) error {
	ret := _m.Called(name, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.NamespacedName, map[string][]byte) error); ok {
		r0 = rf(name, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertWithReplace provides a mock function with given fields: name, data
func (_m *Repository) UpsertWithReplace(name types.NamespacedName, data map[string][]byte) error {
	ret := _m.Called(name, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.NamespacedName, map[string][]byte) error); ok {
		r0 = rf(name, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
