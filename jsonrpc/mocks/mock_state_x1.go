// Code generated by mockery v2.39.0. DO NOT EDIT.

package mocks

import (
	context "context"

	pgx "github.com/jackc/pgx/v4"
)

// GetFinalizedL2BlockNumber provides a mock function with given fields: ctx, l1FinalizedBlockNumber, dbTx
func (_m *StateMock) GetFinalizedL2BlockNumber(ctx context.Context, l1FinalizedBlockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, l1FinalizedBlockNumber, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, l1FinalizedBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, l1FinalizedBlockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, l1FinalizedBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}


// GetSafeL2BlockNumber provides a mock function with given fields: ctx, l1SafeBlockNumber, dbTx
func (_m *StateMock) GetSafeL2BlockNumber(ctx context.Context, l1SafeBlockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, l1SafeBlockNumber, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, l1SafeBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, l1SafeBlockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, l1SafeBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}