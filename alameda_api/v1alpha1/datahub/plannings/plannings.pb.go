// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/plannings/plannings.proto

package plannings

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Planning struct {
	LimitPlannings          []*common.MetricData `protobuf:"bytes,1,rep,name=limit_plannings,json=limitPlannings,proto3" json:"limit_plannings,omitempty"`
	RequestPlannings        []*common.MetricData `protobuf:"bytes,2,rep,name=request_plannings,json=requestPlannings,proto3" json:"request_plannings,omitempty"`
	InitialLimitPlannings   []*common.MetricData `protobuf:"bytes,3,rep,name=initial_limit_plannings,json=initialLimitPlannings,proto3" json:"initial_limit_plannings,omitempty"`
	InitialRequestPlannings []*common.MetricData `protobuf:"bytes,4,rep,name=initial_request_plannings,json=initialRequestPlannings,proto3" json:"initial_request_plannings,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *Planning) Reset()         { *m = Planning{} }
func (m *Planning) String() string { return proto.CompactTextString(m) }
func (*Planning) ProtoMessage()    {}
func (*Planning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{0}
}

func (m *Planning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Planning.Unmarshal(m, b)
}
func (m *Planning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Planning.Marshal(b, m, deterministic)
}
func (m *Planning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Planning.Merge(m, src)
}
func (m *Planning) XXX_Size() int {
	return xxx_messageInfo_Planning.Size(m)
}
func (m *Planning) XXX_DiscardUnknown() {
	xxx_messageInfo_Planning.DiscardUnknown(m)
}

var xxx_messageInfo_Planning proto.InternalMessageInfo

func (m *Planning) GetLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.LimitPlannings
	}
	return nil
}

func (m *Planning) GetRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.RequestPlannings
	}
	return nil
}

func (m *Planning) GetInitialLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialLimitPlannings
	}
	return nil
}

func (m *Planning) GetInitialRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialRequestPlannings
	}
	return nil
}

type ContainerPlanning struct {
	Name                    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitPlannings          []*common.MetricData `protobuf:"bytes,2,rep,name=limit_plannings,json=limitPlannings,proto3" json:"limit_plannings,omitempty"`
	RequestPlannings        []*common.MetricData `protobuf:"bytes,3,rep,name=request_plannings,json=requestPlannings,proto3" json:"request_plannings,omitempty"`
	InitialLimitPlannings   []*common.MetricData `protobuf:"bytes,4,rep,name=initial_limit_plannings,json=initialLimitPlannings,proto3" json:"initial_limit_plannings,omitempty"`
	InitialRequestPlannings []*common.MetricData `protobuf:"bytes,5,rep,name=initial_request_plannings,json=initialRequestPlannings,proto3" json:"initial_request_plannings,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *ContainerPlanning) Reset()         { *m = ContainerPlanning{} }
func (m *ContainerPlanning) String() string { return proto.CompactTextString(m) }
func (*ContainerPlanning) ProtoMessage()    {}
func (*ContainerPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{1}
}

func (m *ContainerPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerPlanning.Unmarshal(m, b)
}
func (m *ContainerPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerPlanning.Marshal(b, m, deterministic)
}
func (m *ContainerPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerPlanning.Merge(m, src)
}
func (m *ContainerPlanning) XXX_Size() int {
	return xxx_messageInfo_ContainerPlanning.Size(m)
}
func (m *ContainerPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerPlanning proto.InternalMessageInfo

func (m *ContainerPlanning) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerPlanning) GetLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.LimitPlannings
	}
	return nil
}

func (m *ContainerPlanning) GetRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.RequestPlannings
	}
	return nil
}

func (m *ContainerPlanning) GetInitialLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialLimitPlannings
	}
	return nil
}

func (m *ContainerPlanning) GetInitialRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialRequestPlannings
	}
	return nil
}

type PodPlanning struct {
	ObjectMeta           *resources.ObjectMeta      `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType               `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                     `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64                    `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                       `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp       `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp       `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	AssignPodPolicy      *resources.AssignPodPolicy `protobuf:"bytes,8,opt,name=assign_pod_policy,json=assignPodPolicy,proto3" json:"assign_pod_policy,omitempty"`
	TopController        *resources.Controller      `protobuf:"bytes,9,opt,name=top_controller,json=topController,proto3" json:"top_controller,omitempty"`
	ContainerPlannings   []*ContainerPlanning       `protobuf:"bytes,10,rep,name=container_plannings,json=containerPlannings,proto3" json:"container_plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PodPlanning) Reset()         { *m = PodPlanning{} }
func (m *PodPlanning) String() string { return proto.CompactTextString(m) }
func (*PodPlanning) ProtoMessage()    {}
func (*PodPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{2}
}

func (m *PodPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodPlanning.Unmarshal(m, b)
}
func (m *PodPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodPlanning.Marshal(b, m, deterministic)
}
func (m *PodPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodPlanning.Merge(m, src)
}
func (m *PodPlanning) XXX_Size() int {
	return xxx_messageInfo_PodPlanning.Size(m)
}
func (m *PodPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_PodPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_PodPlanning proto.InternalMessageInfo

func (m *PodPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *PodPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *PodPlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *PodPlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *PodPlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *PodPlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *PodPlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *PodPlanning) GetAssignPodPolicy() *resources.AssignPodPolicy {
	if m != nil {
		return m.AssignPodPolicy
	}
	return nil
}

func (m *PodPlanning) GetTopController() *resources.Controller {
	if m != nil {
		return m.TopController
	}
	return nil
}

func (m *PodPlanning) GetContainerPlannings() []*ContainerPlanning {
	if m != nil {
		return m.ContainerPlannings
	}
	return nil
}

type ControllerPlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Kind                 resources.Kind        `protobuf:"varint,2,opt,name=kind,proto3,enum=containersai.alameda.v1alpha1.datahub.resources.Kind" json:"kind,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,3,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,4,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,5,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,6,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,9,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ControllerPlanning) Reset()         { *m = ControllerPlanning{} }
func (m *ControllerPlanning) String() string { return proto.CompactTextString(m) }
func (*ControllerPlanning) ProtoMessage()    {}
func (*ControllerPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{3}
}

func (m *ControllerPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerPlanning.Unmarshal(m, b)
}
func (m *ControllerPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerPlanning.Marshal(b, m, deterministic)
}
func (m *ControllerPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerPlanning.Merge(m, src)
}
func (m *ControllerPlanning) XXX_Size() int {
	return xxx_messageInfo_ControllerPlanning.Size(m)
}
func (m *ControllerPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerPlanning proto.InternalMessageInfo

func (m *ControllerPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ControllerPlanning) GetKind() resources.Kind {
	if m != nil {
		return m.Kind
	}
	return resources.Kind_KIND_UNDEFINED
}

func (m *ControllerPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *ControllerPlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *ControllerPlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *ControllerPlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *ControllerPlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ControllerPlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ControllerPlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type ApplicationPlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,8,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ApplicationPlanning) Reset()         { *m = ApplicationPlanning{} }
func (m *ApplicationPlanning) String() string { return proto.CompactTextString(m) }
func (*ApplicationPlanning) ProtoMessage()    {}
func (*ApplicationPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{4}
}

func (m *ApplicationPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationPlanning.Unmarshal(m, b)
}
func (m *ApplicationPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationPlanning.Marshal(b, m, deterministic)
}
func (m *ApplicationPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationPlanning.Merge(m, src)
}
func (m *ApplicationPlanning) XXX_Size() int {
	return xxx_messageInfo_ApplicationPlanning.Size(m)
}
func (m *ApplicationPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationPlanning proto.InternalMessageInfo

func (m *ApplicationPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ApplicationPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *ApplicationPlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *ApplicationPlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *ApplicationPlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *ApplicationPlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ApplicationPlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ApplicationPlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type NamespacePlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,8,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NamespacePlanning) Reset()         { *m = NamespacePlanning{} }
func (m *NamespacePlanning) String() string { return proto.CompactTextString(m) }
func (*NamespacePlanning) ProtoMessage()    {}
func (*NamespacePlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{5}
}

func (m *NamespacePlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespacePlanning.Unmarshal(m, b)
}
func (m *NamespacePlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespacePlanning.Marshal(b, m, deterministic)
}
func (m *NamespacePlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespacePlanning.Merge(m, src)
}
func (m *NamespacePlanning) XXX_Size() int {
	return xxx_messageInfo_NamespacePlanning.Size(m)
}
func (m *NamespacePlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespacePlanning.DiscardUnknown(m)
}

var xxx_messageInfo_NamespacePlanning proto.InternalMessageInfo

func (m *NamespacePlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *NamespacePlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *NamespacePlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *NamespacePlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *NamespacePlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *NamespacePlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *NamespacePlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *NamespacePlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type NodePlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,8,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NodePlanning) Reset()         { *m = NodePlanning{} }
func (m *NodePlanning) String() string { return proto.CompactTextString(m) }
func (*NodePlanning) ProtoMessage()    {}
func (*NodePlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{6}
}

func (m *NodePlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePlanning.Unmarshal(m, b)
}
func (m *NodePlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePlanning.Marshal(b, m, deterministic)
}
func (m *NodePlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePlanning.Merge(m, src)
}
func (m *NodePlanning) XXX_Size() int {
	return xxx_messageInfo_NodePlanning.Size(m)
}
func (m *NodePlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePlanning.DiscardUnknown(m)
}

var xxx_messageInfo_NodePlanning proto.InternalMessageInfo

func (m *NodePlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *NodePlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *NodePlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *NodePlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *NodePlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *NodePlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *NodePlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *NodePlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type ClusterPlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType          `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	PlanningId           string                `protobuf:"bytes,3,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64               `protobuf:"fixed64,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	ApplyPlanningNow     bool                  `protobuf:"varint,5,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Plannings            []*Planning           `protobuf:"bytes,8,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClusterPlanning) Reset()         { *m = ClusterPlanning{} }
func (m *ClusterPlanning) String() string { return proto.CompactTextString(m) }
func (*ClusterPlanning) ProtoMessage()    {}
func (*ClusterPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{7}
}

func (m *ClusterPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterPlanning.Unmarshal(m, b)
}
func (m *ClusterPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterPlanning.Marshal(b, m, deterministic)
}
func (m *ClusterPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterPlanning.Merge(m, src)
}
func (m *ClusterPlanning) XXX_Size() int {
	return xxx_messageInfo_ClusterPlanning.Size(m)
}
func (m *ClusterPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterPlanning proto.InternalMessageInfo

func (m *ClusterPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ClusterPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *ClusterPlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *ClusterPlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *ClusterPlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *ClusterPlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ClusterPlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ClusterPlanning) GetPlannings() []*Planning {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type ReadPlanning struct {
	SchemaMeta           *schemas.SchemaMeta      `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	Plannings            []*common.ReadMetricData `protobuf:"bytes,2,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ReadPlanning) Reset()         { *m = ReadPlanning{} }
func (m *ReadPlanning) String() string { return proto.CompactTextString(m) }
func (*ReadPlanning) ProtoMessage()    {}
func (*ReadPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{8}
}

func (m *ReadPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPlanning.Unmarshal(m, b)
}
func (m *ReadPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPlanning.Marshal(b, m, deterministic)
}
func (m *ReadPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPlanning.Merge(m, src)
}
func (m *ReadPlanning) XXX_Size() int {
	return xxx_messageInfo_ReadPlanning.Size(m)
}
func (m *ReadPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPlanning proto.InternalMessageInfo

func (m *ReadPlanning) GetSchemaMeta() *schemas.SchemaMeta {
	if m != nil {
		return m.SchemaMeta
	}
	return nil
}

func (m *ReadPlanning) GetPlannings() []*common.ReadMetricData {
	if m != nil {
		return m.Plannings
	}
	return nil
}

type WritePlanning struct {
	SchemaMeta           *schemas.SchemaMeta       `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	Plannings            []*common.WriteMetricData `protobuf:"bytes,2,rep,name=plannings,proto3" json:"plannings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *WritePlanning) Reset()         { *m = WritePlanning{} }
func (m *WritePlanning) String() string { return proto.CompactTextString(m) }
func (*WritePlanning) ProtoMessage()    {}
func (*WritePlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{9}
}

func (m *WritePlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WritePlanning.Unmarshal(m, b)
}
func (m *WritePlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WritePlanning.Marshal(b, m, deterministic)
}
func (m *WritePlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WritePlanning.Merge(m, src)
}
func (m *WritePlanning) XXX_Size() int {
	return xxx_messageInfo_WritePlanning.Size(m)
}
func (m *WritePlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_WritePlanning.DiscardUnknown(m)
}

var xxx_messageInfo_WritePlanning proto.InternalMessageInfo

func (m *WritePlanning) GetSchemaMeta() *schemas.SchemaMeta {
	if m != nil {
		return m.SchemaMeta
	}
	return nil
}

func (m *WritePlanning) GetPlannings() []*common.WriteMetricData {
	if m != nil {
		return m.Plannings
	}
	return nil
}

func init() {
	proto.RegisterType((*Planning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.Planning")
	proto.RegisterType((*ContainerPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ContainerPlanning")
	proto.RegisterType((*PodPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.PodPlanning")
	proto.RegisterType((*ControllerPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanning")
	proto.RegisterType((*ApplicationPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ApplicationPlanning")
	proto.RegisterType((*NamespacePlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.NamespacePlanning")
	proto.RegisterType((*NodePlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.NodePlanning")
	proto.RegisterType((*ClusterPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ClusterPlanning")
	proto.RegisterType((*ReadPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ReadPlanning")
	proto.RegisterType((*WritePlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.WritePlanning")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/plannings/plannings.proto", fileDescriptor_36c6a4155f49dbb5)
}

var fileDescriptor_36c6a4155f49dbb5 = []byte{
	// 844 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0x1b, 0xa7, 0x4d, 0x4e, 0xfa, 0x97, 0xa9, 0x10, 0xa6, 0x12, 0x6a, 0x94, 0xab, 0x5c,
	0x80, 0x4d, 0x83, 0x5a, 0x51, 0x41, 0x25, 0xda, 0xc0, 0x45, 0x05, 0x2d, 0x95, 0xa9, 0x54, 0xa9,
	0x20, 0x59, 0x13, 0x7b, 0x48, 0x07, 0x6c, 0xcf, 0xe0, 0x99, 0x50, 0xe5, 0x11, 0x90, 0x10, 0xe2,
	0x81, 0xb8, 0x46, 0xbc, 0xc2, 0xee, 0x83, 0xec, 0xf5, 0xca, 0x13, 0xff, 0xa5, 0xa9, 0x5a, 0xa7,
	0xdd, 0x66, 0x77, 0xa5, 0x5c, 0x75, 0x66, 0xea, 0xf3, 0x7d, 0xf3, 0xcd, 0xf9, 0xce, 0x51, 0x66,
	0x60, 0x1f, 0xfb, 0x38, 0x20, 0x1e, 0x76, 0x30, 0xa7, 0xd6, 0x1f, 0xbb, 0xd8, 0xe7, 0xd7, 0x78,
	0xd7, 0xf2, 0xb0, 0xc4, 0xd7, 0xc3, 0xbe, 0xc5, 0x7d, 0x1c, 0x86, 0x34, 0x1c, 0x88, 0x7c, 0x64,
	0xf2, 0x88, 0x49, 0x86, 0x2c, 0x97, 0x85, 0x12, 0xd3, 0x90, 0x44, 0x02, 0x53, 0x33, 0x01, 0x31,
	0x53, 0x00, 0x33, 0x01, 0x30, 0xb3, 0xb0, 0xed, 0xdd, 0x7b, 0x89, 0x5c, 0x16, 0x04, 0x2c, 0xb4,
	0x02, 0x22, 0x23, 0xea, 0x26, 0x1c, 0xdb, 0xdd, 0x92, 0x7b, 0x93, 0x23, 0x4e, 0xd2, 0x98, 0xbd,
	0x7b, 0x63, 0x22, 0x22, 0xd8, 0x30, 0x72, 0x89, 0x88, 0x99, 0x70, 0xbc, 0x3a, 0x63, 0x18, 0x67,
	0x3e, 0x75, 0x69, 0xc6, 0xb6, 0x5f, 0x32, 0x2c, 0x1b, 0x25, 0x71, 0x9f, 0xdd, 0x1b, 0x27, 0xdc,
	0x6b, 0x12, 0xe0, 0x49, 0x5d, 0x3b, 0x03, 0xc6, 0x06, 0x3e, 0xb1, 0xd4, 0xac, 0x3f, 0xfc, 0xc5,
	0x92, 0x34, 0x20, 0x42, 0xe2, 0x80, 0x8f, 0x3f, 0x68, 0xff, 0x5b, 0x81, 0xda, 0x79, 0x72, 0x24,
	0x08, 0xc3, 0x86, 0x4f, 0x03, 0x2a, 0x9d, 0xec, 0x90, 0x0c, 0xad, 0x55, 0xe9, 0x34, 0xba, 0x5f,
	0x98, 0xe5, 0xf2, 0x36, 0xce, 0x87, 0x79, 0xaa, 0xf2, 0xf1, 0x0d, 0x96, 0xd8, 0x5e, 0x57, 0x80,
	0x29, 0x83, 0x40, 0x04, 0x9a, 0x11, 0xf9, 0x7d, 0x48, 0x44, 0x91, 0x64, 0xe9, 0x89, 0x24, 0x9b,
	0x09, 0x64, 0x4e, 0xc3, 0xe1, 0x43, 0x1a, 0x52, 0x49, 0xb1, 0xef, 0xdc, 0x56, 0x54, 0x79, 0x22,
	0xd9, 0x07, 0x09, 0xf0, 0xf7, 0x93, 0xc2, 0x24, 0x7c, 0x94, 0x32, 0x4e, 0x0b, 0xd4, 0x9f, 0xc8,
	0x99, 0x8a, 0xb1, 0x6f, 0xe9, 0x6c, 0xbf, 0xa8, 0x40, 0xb3, 0x97, 0x82, 0x66, 0x79, 0x44, 0xa0,
	0x87, 0x38, 0x20, 0x86, 0xd6, 0xd2, 0x3a, 0x75, 0x5b, 0x8d, 0xef, 0xca, 0xed, 0xd2, 0x3c, 0x72,
	0x5b, 0x99, 0x67, 0x6e, 0xf5, 0xb7, 0x90, 0xdb, 0xea, 0x73, 0xe5, 0xf6, 0x55, 0x15, 0x1a, 0xe7,
	0xcc, 0xcb, 0xb2, 0xfa, 0x33, 0x34, 0x58, 0xff, 0x57, 0xe2, 0x4a, 0x27, 0xee, 0x42, 0x2a, 0xb9,
	0x8d, 0xee, 0x97, 0x25, 0x79, 0xf3, 0x56, 0xf2, 0x83, 0xc2, 0x38, 0x25, 0x12, 0xdb, 0xc0, 0xb2,
	0x31, 0xea, 0xc3, 0x5a, 0xaa, 0xc9, 0x89, 0x3b, 0x88, 0xb1, 0xd4, 0xd2, 0x3a, 0xeb, 0xdd, 0x43,
	0x73, 0xc6, 0x8e, 0x6d, 0xa6, 0xfb, 0xbd, 0x18, 0x71, 0x62, 0xaf, 0xf2, 0xc2, 0x0c, 0xed, 0x40,
	0x23, 0xe3, 0xa0, 0x9e, 0x51, 0x51, 0xf6, 0x84, 0x74, 0xe9, 0xc4, 0x43, 0x1f, 0x03, 0x48, 0x26,
	0xb1, 0xef, 0xb8, 0x4c, 0x48, 0x43, 0x6f, 0x69, 0x1d, 0xcd, 0xae, 0xab, 0x95, 0x1e, 0x13, 0x12,
	0x7d, 0x02, 0x08, 0x73, 0xee, 0x8f, 0xb2, 0xd3, 0x77, 0x42, 0x76, 0x63, 0x54, 0x5b, 0x5a, 0xa7,
	0x66, 0x6f, 0xaa, 0xff, 0xa4, 0xe4, 0x67, 0xec, 0x06, 0x1d, 0x00, 0x08, 0x89, 0x23, 0xe9, 0xc4,
	0x3d, 0xcf, 0x58, 0x56, 0xc7, 0xb5, 0x6d, 0x8e, 0x1b, 0xa2, 0x99, 0x36, 0x44, 0xf3, 0x22, 0x6d,
	0x88, 0x76, 0x5d, 0x7d, 0x1d, 0xcf, 0xd1, 0x1e, 0xd4, 0x48, 0xe8, 0x8d, 0x03, 0x57, 0x1e, 0x0c,
	0x5c, 0x21, 0xa1, 0xa7, 0xc2, 0x7c, 0x68, 0x62, 0x21, 0xe8, 0x20, 0x74, 0x38, 0xf3, 0x1c, 0xd5,
	0xf4, 0x47, 0x46, 0x4d, 0xc5, 0x7f, 0x3d, 0x73, 0x9e, 0x8e, 0x14, 0x52, 0x6c, 0x00, 0x85, 0x63,
	0x6f, 0xe0, 0xc9, 0x05, 0xd4, 0x87, 0x75, 0xc9, 0xb8, 0x13, 0xe3, 0x46, 0xcc, 0xf7, 0x49, 0x64,
	0xd4, 0x1f, 0x69, 0x89, 0x5e, 0x06, 0x61, 0xaf, 0x49, 0xc6, 0xf3, 0x29, 0x12, 0xb0, 0x95, 0x81,
	0x15, 0x3c, 0x0f, 0xca, 0xf3, 0xc7, 0x33, 0x7b, 0x63, 0xaa, 0x55, 0xd9, 0xc8, 0xbd, 0xbd, 0x24,
	0xda, 0x2f, 0x75, 0x40, 0xf9, 0x1e, 0xe6, 0xe4, 0xff, 0x13, 0xd0, 0x7f, 0xa3, 0xa1, 0x97, 0xd8,
	0x7e, 0x6f, 0x66, 0xd8, 0xef, 0x68, 0xe8, 0xd9, 0x0a, 0x62, 0xba, 0x94, 0x2a, 0xcf, 0x5e, 0x4a,
	0xfa, 0x03, 0xa5, 0x54, 0x2d, 0x57, 0x4a, 0xcb, 0xa5, 0x4a, 0x69, 0xe5, 0xb1, 0xa5, 0x54, 0x2b,
	0x5f, 0x4a, 0x97, 0x50, 0xcf, 0xed, 0x56, 0x57, 0x76, 0x3b, 0x78, 0xf4, 0xf9, 0xd9, 0x39, 0x56,
	0xfb, 0x1f, 0x1d, 0xb6, 0x8e, 0x38, 0xf7, 0xa9, 0x8b, 0x25, 0x65, 0xe1, 0xa2, 0xbb, 0xbe, 0x37,
	0xdd, 0x75, 0xc2, 0x12, 0xb5, 0x37, 0x68, 0x89, 0xbf, 0x75, 0x68, 0x9e, 0xe1, 0x80, 0x08, 0x8e,
	0x5d, 0xb2, 0x30, 0xc4, 0xc2, 0x10, 0x7f, 0xea, 0xb0, 0x7a, 0xc6, 0xbc, 0x85, 0x17, 0x16, 0x5e,
	0x68, 0xff, 0xa5, 0xc3, 0x46, 0xcf, 0x1f, 0x0a, 0x39, 0xb7, 0x5f, 0x22, 0x0b, 0x3b, 0xbc, 0xc3,
	0x76, 0xf8, 0x4f, 0x83, 0x55, 0x9b, 0xe0, 0xfc, 0x56, 0x76, 0x05, 0x8d, 0xf1, 0xc3, 0x4b, 0xd1,
	0x0b, 0x65, 0xb9, 0x92, 0x27, 0x1b, 0xf3, 0x47, 0xf5, 0x77, 0xec, 0x04, 0x91, 0x8d, 0xd1, 0x55,
	0x51, 0xc5, 0xf8, 0xb6, 0xfe, 0xd5, 0x6c, 0xf7, 0xcc, 0x78, 0xab, 0x85, 0xbb, 0x66, 0x41, 0xc8,
	0xff, 0x1a, 0xac, 0x5d, 0x46, 0x54, 0x92, 0xb9, 0x28, 0xf9, 0x69, 0x5a, 0xc9, 0xe1, 0x6c, 0x4a,
	0xd4, 0x5e, 0xef, 0x94, 0x72, 0xfc, 0xed, 0x55, 0x6f, 0x40, 0x65, 0xf2, 0x6d, 0xe1, 0x85, 0xf1,
	0x53, 0x4c, 0x2d, 0xcc, 0xa9, 0x55, 0xee, 0x3d, 0xb0, 0xbf, 0xac, 0x0c, 0xf5, 0xf9, 0xeb, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3e, 0x65, 0xb5, 0x17, 0xdc, 0x14, 0x00, 0x00,
}
