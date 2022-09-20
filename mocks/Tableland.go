// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	tableland "github.com/textileio/go-tableland/internal/tableland"

	tables "github.com/textileio/go-tableland/pkg/tables"
)

// Tableland is an autogenerated mock type for the Tableland type
type Tableland struct {
	mock.Mock
}

type Tableland_Expecter struct {
	mock *mock.Mock
}

func (_m *Tableland) EXPECT() *Tableland_Expecter {
	return &Tableland_Expecter{mock: &_m.Mock}
}

// GetReceipt provides a mock function with given fields: ctx, chainID, txnHash
func (_m *Tableland) GetReceipt(ctx context.Context, chainID tableland.ChainID, txnHash string) (bool, *tableland.TxnReceipt, error) {
	ret := _m.Called(ctx, chainID, txnHash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, tableland.ChainID, string) bool); ok {
		r0 = rf(ctx, chainID, txnHash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *tableland.TxnReceipt
	if rf, ok := ret.Get(1).(func(context.Context, tableland.ChainID, string) *tableland.TxnReceipt); ok {
		r1 = rf(ctx, chainID, txnHash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*tableland.TxnReceipt)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, tableland.ChainID, string) error); ok {
		r2 = rf(ctx, chainID, txnHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Tableland_GetReceipt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReceipt'
type Tableland_GetReceipt_Call struct {
	*mock.Call
}

// GetReceipt is a helper method to define mock.On call
//   - ctx context.Context
//   - chainID tableland.ChainID
//   - txnHash string
func (_e *Tableland_Expecter) GetReceipt(ctx interface{}, chainID interface{}, txnHash interface{}) *Tableland_GetReceipt_Call {
	return &Tableland_GetReceipt_Call{Call: _e.mock.On("GetReceipt", ctx, chainID, txnHash)}
}

func (_c *Tableland_GetReceipt_Call) Run(run func(ctx context.Context, chainID tableland.ChainID, txnHash string)) *Tableland_GetReceipt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(tableland.ChainID), args[2].(string))
	})
	return _c
}

func (_c *Tableland_GetReceipt_Call) Return(_a0 bool, _a1 *tableland.TxnReceipt, _a2 error) *Tableland_GetReceipt_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

// RelayWriteQuery provides a mock function with given fields: ctx, chainID, caller, stmt
func (_m *Tableland) RelayWriteQuery(ctx context.Context, chainID tableland.ChainID, caller common.Address, stmt string) (tables.Transaction, error) {
	ret := _m.Called(ctx, chainID, caller, stmt)

	var r0 tables.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, tableland.ChainID, common.Address, string) tables.Transaction); ok {
		r0 = rf(ctx, chainID, caller, stmt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tables.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, tableland.ChainID, common.Address, string) error); ok {
		r1 = rf(ctx, chainID, caller, stmt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tableland_RelayWriteQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RelayWriteQuery'
type Tableland_RelayWriteQuery_Call struct {
	*mock.Call
}

// RelayWriteQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - chainID tableland.ChainID
//   - caller common.Address
//   - stmt string
func (_e *Tableland_Expecter) RelayWriteQuery(ctx interface{}, chainID interface{}, caller interface{}, stmt interface{}) *Tableland_RelayWriteQuery_Call {
	return &Tableland_RelayWriteQuery_Call{Call: _e.mock.On("RelayWriteQuery", ctx, chainID, caller, stmt)}
}

func (_c *Tableland_RelayWriteQuery_Call) Run(run func(ctx context.Context, chainID tableland.ChainID, caller common.Address, stmt string)) *Tableland_RelayWriteQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(tableland.ChainID), args[2].(common.Address), args[3].(string))
	})
	return _c
}

func (_c *Tableland_RelayWriteQuery_Call) Return(_a0 tables.Transaction, _a1 error) *Tableland_RelayWriteQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// RunReadQuery provides a mock function with given fields: ctx, stmt
func (_m *Tableland) RunReadQuery(ctx context.Context, stmt string) (*tableland.TableData, error) {
	ret := _m.Called(ctx, stmt)

	var r0 *tableland.TableData
	if rf, ok := ret.Get(0).(func(context.Context, string) *tableland.TableData); ok {
		r0 = rf(ctx, stmt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tableland.TableData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, stmt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tableland_RunReadQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunReadQuery'
type Tableland_RunReadQuery_Call struct {
	*mock.Call
}

// RunReadQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - stmt string
func (_e *Tableland_Expecter) RunReadQuery(ctx interface{}, stmt interface{}) *Tableland_RunReadQuery_Call {
	return &Tableland_RunReadQuery_Call{Call: _e.mock.On("RunReadQuery", ctx, stmt)}
}

func (_c *Tableland_RunReadQuery_Call) Run(run func(ctx context.Context, stmt string)) *Tableland_RunReadQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Tableland_RunReadQuery_Call) Return(_a0 *tableland.TableData, _a1 error) *Tableland_RunReadQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SetController provides a mock function with given fields: ctx, chainID, caller, controller, tableID
func (_m *Tableland) SetController(ctx context.Context, chainID tableland.ChainID, caller common.Address, controller common.Address, tableID tables.TableID) (tables.Transaction, error) {
	ret := _m.Called(ctx, chainID, caller, controller, tableID)

	var r0 tables.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, tableland.ChainID, common.Address, common.Address, tables.TableID) tables.Transaction); ok {
		r0 = rf(ctx, chainID, caller, controller, tableID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tables.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, tableland.ChainID, common.Address, common.Address, tables.TableID) error); ok {
		r1 = rf(ctx, chainID, caller, controller, tableID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tableland_SetController_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetController'
type Tableland_SetController_Call struct {
	*mock.Call
}

// SetController is a helper method to define mock.On call
//   - ctx context.Context
//   - chainID tableland.ChainID
//   - caller common.Address
//   - controller common.Address
//   - tableID tables.TableID
func (_e *Tableland_Expecter) SetController(ctx interface{}, chainID interface{}, caller interface{}, controller interface{}, tableID interface{}) *Tableland_SetController_Call {
	return &Tableland_SetController_Call{Call: _e.mock.On("SetController", ctx, chainID, caller, controller, tableID)}
}

func (_c *Tableland_SetController_Call) Run(run func(ctx context.Context, chainID tableland.ChainID, caller common.Address, controller common.Address, tableID tables.TableID)) *Tableland_SetController_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(tableland.ChainID), args[2].(common.Address), args[3].(common.Address), args[4].(tables.TableID))
	})
	return _c
}

func (_c *Tableland_SetController_Call) Return(_a0 tables.Transaction, _a1 error) *Tableland_SetController_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ValidateCreateTable provides a mock function with given fields: ctx, chainID, stmt
func (_m *Tableland) ValidateCreateTable(ctx context.Context, chainID tableland.ChainID, stmt string) (string, error) {
	ret := _m.Called(ctx, chainID, stmt)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, tableland.ChainID, string) string); ok {
		r0 = rf(ctx, chainID, stmt)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, tableland.ChainID, string) error); ok {
		r1 = rf(ctx, chainID, stmt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tableland_ValidateCreateTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateCreateTable'
type Tableland_ValidateCreateTable_Call struct {
	*mock.Call
}

// ValidateCreateTable is a helper method to define mock.On call
//   - ctx context.Context
//   - chainID tableland.ChainID
//   - stmt string
func (_e *Tableland_Expecter) ValidateCreateTable(ctx interface{}, chainID interface{}, stmt interface{}) *Tableland_ValidateCreateTable_Call {
	return &Tableland_ValidateCreateTable_Call{Call: _e.mock.On("ValidateCreateTable", ctx, chainID, stmt)}
}

func (_c *Tableland_ValidateCreateTable_Call) Run(run func(ctx context.Context, chainID tableland.ChainID, stmt string)) *Tableland_ValidateCreateTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(tableland.ChainID), args[2].(string))
	})
	return _c
}

func (_c *Tableland_ValidateCreateTable_Call) Return(_a0 string, _a1 error) *Tableland_ValidateCreateTable_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ValidateWriteQuery provides a mock function with given fields: ctx, chainID, stmt
func (_m *Tableland) ValidateWriteQuery(ctx context.Context, chainID tableland.ChainID, stmt string) (tables.TableID, error) {
	ret := _m.Called(ctx, chainID, stmt)

	var r0 tables.TableID
	if rf, ok := ret.Get(0).(func(context.Context, tableland.ChainID, string) tables.TableID); ok {
		r0 = rf(ctx, chainID, stmt)
	} else {
		r0 = ret.Get(0).(tables.TableID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, tableland.ChainID, string) error); ok {
		r1 = rf(ctx, chainID, stmt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tableland_ValidateWriteQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateWriteQuery'
type Tableland_ValidateWriteQuery_Call struct {
	*mock.Call
}

// ValidateWriteQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - chainID tableland.ChainID
//   - stmt string
func (_e *Tableland_Expecter) ValidateWriteQuery(ctx interface{}, chainID interface{}, stmt interface{}) *Tableland_ValidateWriteQuery_Call {
	return &Tableland_ValidateWriteQuery_Call{Call: _e.mock.On("ValidateWriteQuery", ctx, chainID, stmt)}
}

func (_c *Tableland_ValidateWriteQuery_Call) Run(run func(ctx context.Context, chainID tableland.ChainID, stmt string)) *Tableland_ValidateWriteQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(tableland.ChainID), args[2].(string))
	})
	return _c
}

func (_c *Tableland_ValidateWriteQuery_Call) Return(_a0 tables.TableID, _a1 error) *Tableland_ValidateWriteQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewTableland interface {
	mock.TestingT
	Cleanup(func())
}

// NewTableland creates a new instance of Tableland. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTableland(t mockConstructorTestingTNewTableland) *Tableland {
	mock := &Tableland{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
