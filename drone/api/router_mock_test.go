// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package main

import (
	"context"
	"sync"
)

// Ensure, that StorerMock does implement Storer.
// If this is not the case, regenerate this file with moq.
var _ Storer = &StorerMock{}

// StorerMock is a mock implementation of Storer.
//
//	func TestSomethingThatUsesStorer(t *testing.T) {
//
//		// make and configure a mocked Storer
//		mockedStorer := &StorerMock{
//			createAuthorizationFunc: func(contextMoqParam context.Context, authorization *Authorization) error {
//				panic("mock out the createAuthorization method")
//			},
//			createRequestFunc: func(contextMoqParam context.Context, request *Request) error {
//				panic("mock out the createRequest method")
//			},
//			findAuthorizationFunc: func(contextMoqParam context.Context, s string) (*Authorization, error) {
//				panic("mock out the findAuthorization method")
//			},
//			findRequestFunc: func(contextMoqParam context.Context, s string) (*Request, error) {
//				panic("mock out the findRequest method")
//			},
//		}
//
//		// use mockedStorer in code that requires Storer
//		// and then make assertions.
//
//	}
type StorerMock struct {
	// createAuthorizationFunc mocks the createAuthorization method.
	createAuthorizationFunc func(contextMoqParam context.Context, authorization *Authorization) error

	// createRequestFunc mocks the createRequest method.
	createRequestFunc func(contextMoqParam context.Context, request *Request) error

	// findAuthorizationFunc mocks the findAuthorization method.
	findAuthorizationFunc func(contextMoqParam context.Context, s string) (*Authorization, error)

	// findRequestFunc mocks the findRequest method.
	findRequestFunc func(contextMoqParam context.Context, s string) (*Request, error)

	// calls tracks calls to the methods.
	calls struct {
		// createAuthorization holds details about calls to the createAuthorization method.
		createAuthorization []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Authorization is the authorization argument value.
			Authorization *Authorization
		}
		// createRequest holds details about calls to the createRequest method.
		createRequest []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Request is the request argument value.
			Request *Request
		}
		// findAuthorization holds details about calls to the findAuthorization method.
		findAuthorization []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
		}
		// findRequest holds details about calls to the findRequest method.
		findRequest []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
		}
	}
	lockcreateAuthorization sync.RWMutex
	lockcreateRequest       sync.RWMutex
	lockfindAuthorization   sync.RWMutex
	lockfindRequest         sync.RWMutex
}

// createAuthorization calls createAuthorizationFunc.
func (mock *StorerMock) createAuthorization(contextMoqParam context.Context, authorization *Authorization) error {
	if mock.createAuthorizationFunc == nil {
		panic("StorerMock.createAuthorizationFunc: method is nil but Storer.createAuthorization was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Authorization   *Authorization
	}{
		ContextMoqParam: contextMoqParam,
		Authorization:   authorization,
	}
	mock.lockcreateAuthorization.Lock()
	mock.calls.createAuthorization = append(mock.calls.createAuthorization, callInfo)
	mock.lockcreateAuthorization.Unlock()
	return mock.createAuthorizationFunc(contextMoqParam, authorization)
}

// createAuthorizationCalls gets all the calls that were made to createAuthorization.
// Check the length with:
//
//	len(mockedStorer.createAuthorizationCalls())
func (mock *StorerMock) createAuthorizationCalls() []struct {
	ContextMoqParam context.Context
	Authorization   *Authorization
} {
	var calls []struct {
		ContextMoqParam context.Context
		Authorization   *Authorization
	}
	mock.lockcreateAuthorization.RLock()
	calls = mock.calls.createAuthorization
	mock.lockcreateAuthorization.RUnlock()
	return calls
}

// createRequest calls createRequestFunc.
func (mock *StorerMock) createRequest(contextMoqParam context.Context, request *Request) error {
	if mock.createRequestFunc == nil {
		panic("StorerMock.createRequestFunc: method is nil but Storer.createRequest was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Request         *Request
	}{
		ContextMoqParam: contextMoqParam,
		Request:         request,
	}
	mock.lockcreateRequest.Lock()
	mock.calls.createRequest = append(mock.calls.createRequest, callInfo)
	mock.lockcreateRequest.Unlock()
	return mock.createRequestFunc(contextMoqParam, request)
}

// createRequestCalls gets all the calls that were made to createRequest.
// Check the length with:
//
//	len(mockedStorer.createRequestCalls())
func (mock *StorerMock) createRequestCalls() []struct {
	ContextMoqParam context.Context
	Request         *Request
} {
	var calls []struct {
		ContextMoqParam context.Context
		Request         *Request
	}
	mock.lockcreateRequest.RLock()
	calls = mock.calls.createRequest
	mock.lockcreateRequest.RUnlock()
	return calls
}

// findAuthorization calls findAuthorizationFunc.
func (mock *StorerMock) findAuthorization(contextMoqParam context.Context, s string) (*Authorization, error) {
	if mock.findAuthorizationFunc == nil {
		panic("StorerMock.findAuthorizationFunc: method is nil but Storer.findAuthorization was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
	}
	mock.lockfindAuthorization.Lock()
	mock.calls.findAuthorization = append(mock.calls.findAuthorization, callInfo)
	mock.lockfindAuthorization.Unlock()
	return mock.findAuthorizationFunc(contextMoqParam, s)
}

// findAuthorizationCalls gets all the calls that were made to findAuthorization.
// Check the length with:
//
//	len(mockedStorer.findAuthorizationCalls())
func (mock *StorerMock) findAuthorizationCalls() []struct {
	ContextMoqParam context.Context
	S               string
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
	}
	mock.lockfindAuthorization.RLock()
	calls = mock.calls.findAuthorization
	mock.lockfindAuthorization.RUnlock()
	return calls
}

// findRequest calls findRequestFunc.
func (mock *StorerMock) findRequest(contextMoqParam context.Context, s string) (*Request, error) {
	if mock.findRequestFunc == nil {
		panic("StorerMock.findRequestFunc: method is nil but Storer.findRequest was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
	}
	mock.lockfindRequest.Lock()
	mock.calls.findRequest = append(mock.calls.findRequest, callInfo)
	mock.lockfindRequest.Unlock()
	return mock.findRequestFunc(contextMoqParam, s)
}

// findRequestCalls gets all the calls that were made to findRequest.
// Check the length with:
//
//	len(mockedStorer.findRequestCalls())
func (mock *StorerMock) findRequestCalls() []struct {
	ContextMoqParam context.Context
	S               string
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
	}
	mock.lockfindRequest.RLock()
	calls = mock.calls.findRequest
	mock.lockfindRequest.RUnlock()
	return calls
}
