// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	command "github.com/determined-ai/determined/master/pkg/command"
	apiv1 "github.com/determined-ai/determined/proto/pkg/apiv1"

	jobv1 "github.com/determined-ai/determined/proto/pkg/jobv1"

	mock "github.com/stretchr/testify/mock"

	model "github.com/determined-ai/determined/master/pkg/model"

	sproto "github.com/determined-ai/determined/master/internal/sproto"
)

// ResourceManager is an autogenerated mock type for the ResourceManager type
type ResourceManager struct {
	mock.Mock
}

// Allocate provides a mock function with given fields: _a0
func (_m *ResourceManager) Allocate(_a0 sproto.AllocateRequest) (*sproto.ResourcesSubscription, error) {
	ret := _m.Called(_a0)

	var r0 *sproto.ResourcesSubscription
	var r1 error
	if rf, ok := ret.Get(0).(func(sproto.AllocateRequest) (*sproto.ResourcesSubscription, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(sproto.AllocateRequest) *sproto.ResourcesSubscription); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sproto.ResourcesSubscription)
		}
	}

	if rf, ok := ret.Get(1).(func(sproto.AllocateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteJob provides a mock function with given fields: _a0
func (_m *ResourceManager) DeleteJob(_a0 sproto.DeleteJob) (sproto.DeleteJobResponse, error) {
	ret := _m.Called(_a0)

	var r0 sproto.DeleteJobResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(sproto.DeleteJob) (sproto.DeleteJobResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(sproto.DeleteJob) sproto.DeleteJobResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(sproto.DeleteJobResponse)
	}

	if rf, ok := ret.Get(1).(func(sproto.DeleteJob) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableAgent provides a mock function with given fields: _a0
func (_m *ResourceManager) DisableAgent(_a0 *apiv1.DisableAgentRequest) (*apiv1.DisableAgentResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.DisableAgentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.DisableAgentRequest) (*apiv1.DisableAgentResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.DisableAgentRequest) *apiv1.DisableAgentResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.DisableAgentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.DisableAgentRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableSlot provides a mock function with given fields: _a0
func (_m *ResourceManager) DisableSlot(_a0 *apiv1.DisableSlotRequest) (*apiv1.DisableSlotResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.DisableSlotResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.DisableSlotRequest) (*apiv1.DisableSlotResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.DisableSlotRequest) *apiv1.DisableSlotResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.DisableSlotResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.DisableSlotRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableAgent provides a mock function with given fields: _a0
func (_m *ResourceManager) EnableAgent(_a0 *apiv1.EnableAgentRequest) (*apiv1.EnableAgentResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.EnableAgentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.EnableAgentRequest) (*apiv1.EnableAgentResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.EnableAgentRequest) *apiv1.EnableAgentResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.EnableAgentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.EnableAgentRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableSlot provides a mock function with given fields: _a0
func (_m *ResourceManager) EnableSlot(_a0 *apiv1.EnableSlotRequest) (*apiv1.EnableSlotResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.EnableSlotResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.EnableSlotRequest) (*apiv1.EnableSlotResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.EnableSlotRequest) *apiv1.EnableSlotResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.EnableSlotResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.EnableSlotRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExternalPreemptionPending provides a mock function with given fields: _a0
func (_m *ResourceManager) ExternalPreemptionPending(_a0 sproto.PendingPreemption) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(sproto.PendingPreemption) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAgent provides a mock function with given fields: _a0
func (_m *ResourceManager) GetAgent(_a0 *apiv1.GetAgentRequest) (*apiv1.GetAgentResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.GetAgentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.GetAgentRequest) (*apiv1.GetAgentResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.GetAgentRequest) *apiv1.GetAgentResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.GetAgentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.GetAgentRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgents provides a mock function with given fields: _a0
func (_m *ResourceManager) GetAgents(_a0 *apiv1.GetAgentsRequest) (*apiv1.GetAgentsResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.GetAgentsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.GetAgentsRequest) (*apiv1.GetAgentsResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.GetAgentsRequest) *apiv1.GetAgentsResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.GetAgentsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.GetAgentsRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllocationSummaries provides a mock function with given fields: _a0
func (_m *ResourceManager) GetAllocationSummaries(_a0 sproto.GetAllocationSummaries) (map[model.AllocationID]sproto.AllocationSummary, error) {
	ret := _m.Called(_a0)

	var r0 map[model.AllocationID]sproto.AllocationSummary
	var r1 error
	if rf, ok := ret.Get(0).(func(sproto.GetAllocationSummaries) (map[model.AllocationID]sproto.AllocationSummary, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(sproto.GetAllocationSummaries) map[model.AllocationID]sproto.AllocationSummary); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.AllocationID]sproto.AllocationSummary)
		}
	}

	if rf, ok := ret.Get(1).(func(sproto.GetAllocationSummaries) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllocationSummary provides a mock function with given fields: _a0
func (_m *ResourceManager) GetAllocationSummary(_a0 sproto.GetAllocationSummary) (*sproto.AllocationSummary, error) {
	ret := _m.Called(_a0)

	var r0 *sproto.AllocationSummary
	var r1 error
	if rf, ok := ret.Get(0).(func(sproto.GetAllocationSummary) (*sproto.AllocationSummary, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(sproto.GetAllocationSummary) *sproto.AllocationSummary); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sproto.AllocationSummary)
		}
	}

	if rf, ok := ret.Get(1).(func(sproto.GetAllocationSummary) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultAuxResourcePool provides a mock function with given fields: _a0
func (_m *ResourceManager) GetDefaultAuxResourcePool(_a0 sproto.GetDefaultAuxResourcePoolRequest) (sproto.GetDefaultAuxResourcePoolResponse, error) {
	ret := _m.Called(_a0)

	var r0 sproto.GetDefaultAuxResourcePoolResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(sproto.GetDefaultAuxResourcePoolRequest) (sproto.GetDefaultAuxResourcePoolResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(sproto.GetDefaultAuxResourcePoolRequest) sproto.GetDefaultAuxResourcePoolResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(sproto.GetDefaultAuxResourcePoolResponse)
	}

	if rf, ok := ret.Get(1).(func(sproto.GetDefaultAuxResourcePoolRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultComputeResourcePool provides a mock function with given fields: _a0
func (_m *ResourceManager) GetDefaultComputeResourcePool(_a0 sproto.GetDefaultComputeResourcePoolRequest) (sproto.GetDefaultComputeResourcePoolResponse, error) {
	ret := _m.Called(_a0)

	var r0 sproto.GetDefaultComputeResourcePoolResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(sproto.GetDefaultComputeResourcePoolRequest) (sproto.GetDefaultComputeResourcePoolResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(sproto.GetDefaultComputeResourcePoolRequest) sproto.GetDefaultComputeResourcePoolResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(sproto.GetDefaultComputeResourcePoolResponse)
	}

	if rf, ok := ret.Get(1).(func(sproto.GetDefaultComputeResourcePoolRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExternalJobs provides a mock function with given fields: _a0
func (_m *ResourceManager) GetExternalJobs(_a0 sproto.GetExternalJobs) ([]*jobv1.Job, error) {
	ret := _m.Called(_a0)

	var r0 []*jobv1.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(sproto.GetExternalJobs) ([]*jobv1.Job, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(sproto.GetExternalJobs) []*jobv1.Job); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*jobv1.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(sproto.GetExternalJobs) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobQ provides a mock function with given fields: _a0
func (_m *ResourceManager) GetJobQ(_a0 sproto.GetJobQ) (map[model.JobID]*sproto.RMJobInfo, error) {
	ret := _m.Called(_a0)

	var r0 map[model.JobID]*sproto.RMJobInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(sproto.GetJobQ) (map[model.JobID]*sproto.RMJobInfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(sproto.GetJobQ) map[model.JobID]*sproto.RMJobInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.JobID]*sproto.RMJobInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(sproto.GetJobQ) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobQueueStatsRequest provides a mock function with given fields: _a0
func (_m *ResourceManager) GetJobQueueStatsRequest(_a0 *apiv1.GetJobQueueStatsRequest) (*apiv1.GetJobQueueStatsResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.GetJobQueueStatsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.GetJobQueueStatsRequest) (*apiv1.GetJobQueueStatsResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.GetJobQueueStatsRequest) *apiv1.GetJobQueueStatsResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.GetJobQueueStatsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.GetJobQueueStatsRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcePools provides a mock function with given fields: _a0
func (_m *ResourceManager) GetResourcePools(_a0 *apiv1.GetResourcePoolsRequest) (*apiv1.GetResourcePoolsResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.GetResourcePoolsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.GetResourcePoolsRequest) (*apiv1.GetResourcePoolsResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.GetResourcePoolsRequest) *apiv1.GetResourcePoolsResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.GetResourcePoolsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.GetResourcePoolsRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSlot provides a mock function with given fields: _a0
func (_m *ResourceManager) GetSlot(_a0 *apiv1.GetSlotRequest) (*apiv1.GetSlotResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.GetSlotResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.GetSlotRequest) (*apiv1.GetSlotResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.GetSlotRequest) *apiv1.GetSlotResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.GetSlotResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.GetSlotRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSlots provides a mock function with given fields: _a0
func (_m *ResourceManager) GetSlots(_a0 *apiv1.GetSlotsRequest) (*apiv1.GetSlotsResponse, error) {
	ret := _m.Called(_a0)

	var r0 *apiv1.GetSlotsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*apiv1.GetSlotsRequest) (*apiv1.GetSlotsResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*apiv1.GetSlotsRequest) *apiv1.GetSlotsResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.GetSlotsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*apiv1.GetSlotsRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsReattachableOnlyAfterStarted provides a mock function with given fields:
func (_m *ResourceManager) IsReattachableOnlyAfterStarted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MoveJob provides a mock function with given fields: _a0
func (_m *ResourceManager) MoveJob(_a0 sproto.MoveJob) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(sproto.MoveJob) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NotifyContainerRunning provides a mock function with given fields: _a0
func (_m *ResourceManager) NotifyContainerRunning(_a0 sproto.NotifyContainerRunning) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(sproto.NotifyContainerRunning) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RecoverJobPosition provides a mock function with given fields: _a0
func (_m *ResourceManager) RecoverJobPosition(_a0 sproto.RecoverJobPosition) {
	_m.Called(_a0)
}

// Release provides a mock function with given fields: _a0
func (_m *ResourceManager) Release(_a0 sproto.ResourcesReleased) {
	_m.Called(_a0)
}

// ResolveResourcePool provides a mock function with given fields: name, workspace, slots
func (_m *ResourceManager) ResolveResourcePool(name string, workspace int, slots int) (string, error) {
	ret := _m.Called(name, workspace, slots)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) (string, error)); ok {
		return rf(name, workspace, slots)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) string); ok {
		r0 = rf(name, workspace, slots)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(name, workspace, slots)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAllocationName provides a mock function with given fields: _a0
func (_m *ResourceManager) SetAllocationName(_a0 sproto.SetAllocationName) {
	_m.Called(_a0)
}

// SetGroupMaxSlots provides a mock function with given fields: _a0
func (_m *ResourceManager) SetGroupMaxSlots(_a0 sproto.SetGroupMaxSlots) {
	_m.Called(_a0)
}

// SetGroupPriority provides a mock function with given fields: _a0
func (_m *ResourceManager) SetGroupPriority(_a0 sproto.SetGroupPriority) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(sproto.SetGroupPriority) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetGroupWeight provides a mock function with given fields: _a0
func (_m *ResourceManager) SetGroupWeight(_a0 sproto.SetGroupWeight) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(sproto.SetGroupWeight) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TaskContainerDefaults provides a mock function with given fields: resourcePoolName, fallbackConfig
func (_m *ResourceManager) TaskContainerDefaults(resourcePoolName string, fallbackConfig model.TaskContainerDefaultsConfig) (model.TaskContainerDefaultsConfig, error) {
	ret := _m.Called(resourcePoolName, fallbackConfig)

	var r0 model.TaskContainerDefaultsConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(string, model.TaskContainerDefaultsConfig) (model.TaskContainerDefaultsConfig, error)); ok {
		return rf(resourcePoolName, fallbackConfig)
	}
	if rf, ok := ret.Get(0).(func(string, model.TaskContainerDefaultsConfig) model.TaskContainerDefaultsConfig); ok {
		r0 = rf(resourcePoolName, fallbackConfig)
	} else {
		r0 = ret.Get(0).(model.TaskContainerDefaultsConfig)
	}

	if rf, ok := ret.Get(1).(func(string, model.TaskContainerDefaultsConfig) error); ok {
		r1 = rf(resourcePoolName, fallbackConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateCommandResources provides a mock function with given fields: _a0
func (_m *ResourceManager) ValidateCommandResources(_a0 sproto.ValidateCommandResourcesRequest) (sproto.ValidateCommandResourcesResponse, error) {
	ret := _m.Called(_a0)

	var r0 sproto.ValidateCommandResourcesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(sproto.ValidateCommandResourcesRequest) (sproto.ValidateCommandResourcesResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(sproto.ValidateCommandResourcesRequest) sproto.ValidateCommandResourcesResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(sproto.ValidateCommandResourcesResponse)
	}

	if rf, ok := ret.Get(1).(func(sproto.ValidateCommandResourcesRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateResourcePool provides a mock function with given fields: name
func (_m *ResourceManager) ValidateResourcePool(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateResourcePoolAvailability provides a mock function with given fields: name, slots
func (_m *ResourceManager) ValidateResourcePoolAvailability(name string, slots int) ([]command.LaunchWarning, error) {
	ret := _m.Called(name, slots)

	var r0 []command.LaunchWarning
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) ([]command.LaunchWarning, error)); ok {
		return rf(name, slots)
	}
	if rf, ok := ret.Get(0).(func(string, int) []command.LaunchWarning); ok {
		r0 = rf(name, slots)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]command.LaunchWarning)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(name, slots)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateResources provides a mock function with given fields: name, slots, _a2
func (_m *ResourceManager) ValidateResources(name string, slots int, _a2 bool) error {
	ret := _m.Called(name, slots, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, bool) error); ok {
		r0 = rf(name, slots, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewResourceManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewResourceManager creates a new instance of ResourceManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResourceManager(t mockConstructorTestingTNewResourceManager) *ResourceManager {
	mock := &ResourceManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
