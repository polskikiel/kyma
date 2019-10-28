// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	mock "github.com/stretchr/testify/mock"
	types "k8s.io/apimachinery/pkg/types"
)

// SecretRepository is an autogenerated mock type for the SecretRepository type
type SecretRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: name
func (_m *SecretRepository) Get(name types.NamespacedName) (map[string][]byte, error) {
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

// Upsert provides a mock function with given fields: name, data
func (_m *SecretRepository) Upsert(name types.NamespacedName, data map[string][]byte) error {
	ret := _m.Called(name, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.NamespacedName, map[string][]byte) error); ok {
		r0 = rf(name, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValuesProvided provides a mock function with given fields: secretName, keys
func (_m *SecretRepository) ValuesProvided(secretName types.NamespacedName, keys []string) (bool, error) {
	ret := _m.Called(secretName, keys)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.NamespacedName, []string) bool); ok {
		r0 = rf(secretName, keys)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.NamespacedName, []string) error); ok {
		r1 = rf(secretName, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
