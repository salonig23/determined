// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	bun "github.com/uptrace/bun"

	mock "github.com/stretchr/testify/mock"

	model "github.com/determined-ai/determined/master/pkg/model"

	projectv1 "github.com/determined-ai/determined/proto/pkg/projectv1"

	rbacv1 "github.com/determined-ai/determined/proto/pkg/rbacv1"
)

// ExperimentAuthZ is an autogenerated mock type for the ExperimentAuthZ type
type ExperimentAuthZ struct {
	mock.Mock
}

// CanCreateExperiment provides a mock function with given fields: ctx, curUser, proj, e
func (_m *ExperimentAuthZ) CanCreateExperiment(ctx context.Context, curUser model.User, proj *projectv1.Project, e *model.Experiment) error {
	ret := _m.Called(ctx, curUser, proj, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *projectv1.Project, *model.Experiment) error); ok {
		r0 = rf(ctx, curUser, proj, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanDeleteExperiment provides a mock function with given fields: ctx, curUser, e
func (_m *ExperimentAuthZ) CanDeleteExperiment(ctx context.Context, curUser model.User, e *model.Experiment) error {
	ret := _m.Called(ctx, curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment) error); ok {
		r0 = rf(ctx, curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanEditExperiment provides a mock function with given fields: ctx, curUser, e
func (_m *ExperimentAuthZ) CanEditExperiment(ctx context.Context, curUser model.User, e *model.Experiment) error {
	ret := _m.Called(ctx, curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment) error); ok {
		r0 = rf(ctx, curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanEditExperimentsMetadata provides a mock function with given fields: ctx, curUser, e
func (_m *ExperimentAuthZ) CanEditExperimentsMetadata(ctx context.Context, curUser model.User, e *model.Experiment) error {
	ret := _m.Called(ctx, curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment) error); ok {
		r0 = rf(ctx, curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanForkFromExperiment provides a mock function with given fields: ctx, curUser, e
func (_m *ExperimentAuthZ) CanForkFromExperiment(ctx context.Context, curUser model.User, e *model.Experiment) error {
	ret := _m.Called(ctx, curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment) error); ok {
		r0 = rf(ctx, curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanGetExperiment provides a mock function with given fields: ctx, curUser, e
func (_m *ExperimentAuthZ) CanGetExperiment(ctx context.Context, curUser model.User, e *model.Experiment) error {
	ret := _m.Called(ctx, curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment) error); ok {
		r0 = rf(ctx, curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanGetExperimentArtifacts provides a mock function with given fields: ctx, curUser, e
func (_m *ExperimentAuthZ) CanGetExperimentArtifacts(ctx context.Context, curUser model.User, e *model.Experiment) error {
	ret := _m.Called(ctx, curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment) error); ok {
		r0 = rf(ctx, curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanPreviewHPSearch provides a mock function with given fields: ctx, curUser
func (_m *ExperimentAuthZ) CanPreviewHPSearch(ctx context.Context, curUser model.User) error {
	ret := _m.Called(ctx, curUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(ctx, curUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanRunCustomSearch provides a mock function with given fields: ctx, curUser, e
func (_m *ExperimentAuthZ) CanRunCustomSearch(ctx context.Context, curUser model.User, e *model.Experiment) error {
	ret := _m.Called(ctx, curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment) error); ok {
		r0 = rf(ctx, curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetExperimentsCheckpointGCPolicy provides a mock function with given fields: ctx, curUser, e
func (_m *ExperimentAuthZ) CanSetExperimentsCheckpointGCPolicy(ctx context.Context, curUser model.User, e *model.Experiment) error {
	ret := _m.Called(ctx, curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment) error); ok {
		r0 = rf(ctx, curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetExperimentsMaxSlots provides a mock function with given fields: ctx, curUser, e, slots
func (_m *ExperimentAuthZ) CanSetExperimentsMaxSlots(ctx context.Context, curUser model.User, e *model.Experiment, slots int) error {
	ret := _m.Called(ctx, curUser, e, slots)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment, int) error); ok {
		r0 = rf(ctx, curUser, e, slots)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetExperimentsPriority provides a mock function with given fields: ctx, curUser, e, priority
func (_m *ExperimentAuthZ) CanSetExperimentsPriority(ctx context.Context, curUser model.User, e *model.Experiment, priority int) error {
	ret := _m.Called(ctx, curUser, e, priority)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment, int) error); ok {
		r0 = rf(ctx, curUser, e, priority)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetExperimentsWeight provides a mock function with given fields: ctx, curUser, e, weight
func (_m *ExperimentAuthZ) CanSetExperimentsWeight(ctx context.Context, curUser model.User, e *model.Experiment, weight float64) error {
	ret := _m.Called(ctx, curUser, e, weight)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *model.Experiment, float64) error); ok {
		r0 = rf(ctx, curUser, e, weight)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterExperimentLabelsQuery provides a mock function with given fields: ctx, curUser, proj, query
func (_m *ExperimentAuthZ) FilterExperimentLabelsQuery(ctx context.Context, curUser model.User, proj *projectv1.Project, query *bun.SelectQuery) (*bun.SelectQuery, error) {
	ret := _m.Called(ctx, curUser, proj, query)

	var r0 *bun.SelectQuery
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *projectv1.Project, *bun.SelectQuery) (*bun.SelectQuery, error)); ok {
		return rf(ctx, curUser, proj, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *projectv1.Project, *bun.SelectQuery) *bun.SelectQuery); ok {
		r0 = rf(ctx, curUser, proj, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bun.SelectQuery)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User, *projectv1.Project, *bun.SelectQuery) error); ok {
		r1 = rf(ctx, curUser, proj, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterExperimentsQuery provides a mock function with given fields: ctx, curUser, proj, query, permissions
func (_m *ExperimentAuthZ) FilterExperimentsQuery(ctx context.Context, curUser model.User, proj *projectv1.Project, query *bun.SelectQuery, permissions []rbacv1.PermissionType) (*bun.SelectQuery, error) {
	ret := _m.Called(ctx, curUser, proj, query, permissions)

	var r0 *bun.SelectQuery
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *projectv1.Project, *bun.SelectQuery, []rbacv1.PermissionType) (*bun.SelectQuery, error)); ok {
		return rf(ctx, curUser, proj, query, permissions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *projectv1.Project, *bun.SelectQuery, []rbacv1.PermissionType) *bun.SelectQuery); ok {
		r0 = rf(ctx, curUser, proj, query, permissions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bun.SelectQuery)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User, *projectv1.Project, *bun.SelectQuery, []rbacv1.PermissionType) error); ok {
		r1 = rf(ctx, curUser, proj, query, permissions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExperimentAuthZ interface {
	mock.TestingT
	Cleanup(func())
}

// NewExperimentAuthZ creates a new instance of ExperimentAuthZ. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExperimentAuthZ(t mockConstructorTestingTNewExperimentAuthZ) *ExperimentAuthZ {
	mock := &ExperimentAuthZ{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
