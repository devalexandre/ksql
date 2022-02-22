// Code generated by MockGen. DO NOT EDIT.
// Source: contracts.go

// Package exampleservice is a generated GoMock package.
package exampleservice

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ksql "github.com/vingarcia/ksql"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockProvider) Delete(ctx context.Context, table ksql.Table, idOrRecord interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, table, idOrRecord)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProviderMockRecorder) Delete(ctx, table, idOrRecord interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProvider)(nil).Delete), ctx, table, idOrRecord)
}

// Exec mocks base method.
func (m *MockProvider) Exec(ctx context.Context, query string, params ...interface{}) (ksql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(ksql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockProviderMockRecorder) Exec(ctx, query interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockProvider)(nil).Exec), varargs...)
}

// Insert mocks base method.
func (m *MockProvider) Insert(ctx context.Context, table ksql.Table, record interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, table, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockProviderMockRecorder) Insert(ctx, table, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockProvider)(nil).Insert), ctx, table, record)
}

// Patch mocks base method.
func (m *MockProvider) Patch(ctx context.Context, table ksql.Table, record interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", ctx, table, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockProviderMockRecorder) Patch(ctx, table, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockProvider)(nil).Patch), ctx, table, record)
}

// Query mocks base method.
func (m *MockProvider) Query(ctx context.Context, records interface{}, query string, params ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, records, query}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockProviderMockRecorder) Query(ctx, records, query interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, records, query}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockProvider)(nil).Query), varargs...)
}

// QueryChunks mocks base method.
func (m *MockProvider) QueryChunks(ctx context.Context, parser ksql.ChunkParser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryChunks", ctx, parser)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryChunks indicates an expected call of QueryChunks.
func (mr *MockProviderMockRecorder) QueryChunks(ctx, parser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryChunks", reflect.TypeOf((*MockProvider)(nil).QueryChunks), ctx, parser)
}

// QueryOne mocks base method.
func (m *MockProvider) QueryOne(ctx context.Context, record interface{}, query string, params ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, record, query}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryOne", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryOne indicates an expected call of QueryOne.
func (mr *MockProviderMockRecorder) QueryOne(ctx, record, query interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, record, query}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOne", reflect.TypeOf((*MockProvider)(nil).QueryOne), varargs...)
}

// Transaction mocks base method.
func (m *MockProvider) Transaction(ctx context.Context, fn func(ksql.Provider) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockProviderMockRecorder) Transaction(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockProvider)(nil).Transaction), ctx, fn)
}

// Update mocks base method.
func (m *MockProvider) Update(ctx context.Context, table ksql.Table, record interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, table, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProviderMockRecorder) Update(ctx, table, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProvider)(nil).Update), ctx, table, record)
}
