// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/recommendation.proto

package containers_ai_alameda_v1alpha1_datahub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// Represents a resource configuration recommendation made by the AI Engine
//
// It includes recommended limits and requests for the initial stage (a container which is just started) and after the initial strage
//
type ContainerRecommendation struct {
	Name                          string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitRecommendations          []*MetricData `protobuf:"bytes,2,rep,name=limit_recommendations,json=limitRecommendations,proto3" json:"limit_recommendations,omitempty"`
	RequestRecommendations        []*MetricData `protobuf:"bytes,3,rep,name=request_recommendations,json=requestRecommendations,proto3" json:"request_recommendations,omitempty"`
	InitialLimitRecommendations   []*MetricData `protobuf:"bytes,4,rep,name=initial_limit_recommendations,json=initialLimitRecommendations,proto3" json:"initial_limit_recommendations,omitempty"`
	InitialRequestRecommendations []*MetricData `protobuf:"bytes,5,rep,name=initial_request_recommendations,json=initialRequestRecommendations,proto3" json:"initial_request_recommendations,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}      `json:"-"`
	XXX_unrecognized              []byte        `json:"-"`
	XXX_sizecache                 int32         `json:"-"`
}

func (m *ContainerRecommendation) Reset()         { *m = ContainerRecommendation{} }
func (m *ContainerRecommendation) String() string { return proto.CompactTextString(m) }
func (*ContainerRecommendation) ProtoMessage()    {}
func (*ContainerRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_626eeb1447246c93, []int{0}
}
func (m *ContainerRecommendation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRecommendation.Unmarshal(m, b)
}
func (m *ContainerRecommendation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRecommendation.Marshal(b, m, deterministic)
}
func (dst *ContainerRecommendation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRecommendation.Merge(dst, src)
}
func (m *ContainerRecommendation) XXX_Size() int {
	return xxx_messageInfo_ContainerRecommendation.Size(m)
}
func (m *ContainerRecommendation) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRecommendation.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRecommendation proto.InternalMessageInfo

func (m *ContainerRecommendation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerRecommendation) GetLimitRecommendations() []*MetricData {
	if m != nil {
		return m.LimitRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetRequestRecommendations() []*MetricData {
	if m != nil {
		return m.RequestRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetInitialLimitRecommendations() []*MetricData {
	if m != nil {
		return m.InitialLimitRecommendations
	}
	return nil
}

func (m *ContainerRecommendation) GetInitialRequestRecommendations() []*MetricData {
	if m != nil {
		return m.InitialRequestRecommendations
	}
	return nil
}

// *
// Represents a recommended pod-to-node assignment (i.e. pod placement) made by the AI Engine
//
type AssignPodPolicy struct {
	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Types that are valid to be assigned to Policy:
	//	*AssignPodPolicy_NodePriority
	//	*AssignPodPolicy_NodeSelector
	//	*AssignPodPolicy_NodeName
	Policy               isAssignPodPolicy_Policy `protobuf_oneof:"policy"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AssignPodPolicy) Reset()         { *m = AssignPodPolicy{} }
func (m *AssignPodPolicy) String() string { return proto.CompactTextString(m) }
func (*AssignPodPolicy) ProtoMessage()    {}
func (*AssignPodPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_626eeb1447246c93, []int{1}
}
func (m *AssignPodPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignPodPolicy.Unmarshal(m, b)
}
func (m *AssignPodPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignPodPolicy.Marshal(b, m, deterministic)
}
func (dst *AssignPodPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignPodPolicy.Merge(dst, src)
}
func (m *AssignPodPolicy) XXX_Size() int {
	return xxx_messageInfo_AssignPodPolicy.Size(m)
}
func (m *AssignPodPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignPodPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_AssignPodPolicy proto.InternalMessageInfo

func (m *AssignPodPolicy) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type isAssignPodPolicy_Policy interface {
	isAssignPodPolicy_Policy()
}

type AssignPodPolicy_NodePriority struct {
	NodePriority *NodePriority `protobuf:"bytes,2,opt,name=node_priority,json=nodePriority,proto3,oneof"`
}

type AssignPodPolicy_NodeSelector struct {
	NodeSelector *Selector `protobuf:"bytes,3,opt,name=node_selector,json=nodeSelector,proto3,oneof"`
}

type AssignPodPolicy_NodeName struct {
	NodeName string `protobuf:"bytes,4,opt,name=node_name,json=nodeName,proto3,oneof"`
}

func (*AssignPodPolicy_NodePriority) isAssignPodPolicy_Policy() {}

func (*AssignPodPolicy_NodeSelector) isAssignPodPolicy_Policy() {}

func (*AssignPodPolicy_NodeName) isAssignPodPolicy_Policy() {}

func (m *AssignPodPolicy) GetPolicy() isAssignPodPolicy_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *AssignPodPolicy) GetNodePriority() *NodePriority {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodePriority); ok {
		return x.NodePriority
	}
	return nil
}

func (m *AssignPodPolicy) GetNodeSelector() *Selector {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodeSelector); ok {
		return x.NodeSelector
	}
	return nil
}

func (m *AssignPodPolicy) GetNodeName() string {
	if x, ok := m.GetPolicy().(*AssignPodPolicy_NodeName); ok {
		return x.NodeName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AssignPodPolicy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AssignPodPolicy_OneofMarshaler, _AssignPodPolicy_OneofUnmarshaler, _AssignPodPolicy_OneofSizer, []interface{}{
		(*AssignPodPolicy_NodePriority)(nil),
		(*AssignPodPolicy_NodeSelector)(nil),
		(*AssignPodPolicy_NodeName)(nil),
	}
}

func _AssignPodPolicy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AssignPodPolicy)
	// policy
	switch x := m.Policy.(type) {
	case *AssignPodPolicy_NodePriority:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodePriority); err != nil {
			return err
		}
	case *AssignPodPolicy_NodeSelector:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodeSelector); err != nil {
			return err
		}
	case *AssignPodPolicy_NodeName:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NodeName)
	case nil:
	default:
		return fmt.Errorf("AssignPodPolicy.Policy has unexpected type %T", x)
	}
	return nil
}

func _AssignPodPolicy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AssignPodPolicy)
	switch tag {
	case 2: // policy.node_priority
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NodePriority)
		err := b.DecodeMessage(msg)
		m.Policy = &AssignPodPolicy_NodePriority{msg}
		return true, err
	case 3: // policy.node_selector
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Selector)
		err := b.DecodeMessage(msg)
		m.Policy = &AssignPodPolicy_NodeSelector{msg}
		return true, err
	case 4: // policy.node_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Policy = &AssignPodPolicy_NodeName{x}
		return true, err
	default:
		return false, nil
	}
}

func _AssignPodPolicy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AssignPodPolicy)
	// policy
	switch x := m.Policy.(type) {
	case *AssignPodPolicy_NodePriority:
		s := proto.Size(x.NodePriority)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AssignPodPolicy_NodeSelector:
		s := proto.Size(x.NodeSelector)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AssignPodPolicy_NodeName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.NodeName)))
		n += len(x.NodeName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// *
// Represents a set of container resource configuration recommenation of a pod made by the AI Engine
//
type PodRecommendation struct {
	NamespacedName           *NamespacedName            `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ApplyRecommendationNow   bool                       `protobuf:"varint,2,opt,name=apply_recommendation_now,json=applyRecommendationNow,proto3" json:"apply_recommendation_now,omitempty"`
	AssignPodPolicy          *AssignPodPolicy           `protobuf:"bytes,3,opt,name=assign_pod_policy,json=assignPodPolicy,proto3" json:"assign_pod_policy,omitempty"`
	ContainerRecommendations []*ContainerRecommendation `protobuf:"bytes,4,rep,name=container_recommendations,json=containerRecommendations,proto3" json:"container_recommendations,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *PodRecommendation) Reset()         { *m = PodRecommendation{} }
func (m *PodRecommendation) String() string { return proto.CompactTextString(m) }
func (*PodRecommendation) ProtoMessage()    {}
func (*PodRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_recommendation_626eeb1447246c93, []int{2}
}
func (m *PodRecommendation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodRecommendation.Unmarshal(m, b)
}
func (m *PodRecommendation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodRecommendation.Marshal(b, m, deterministic)
}
func (dst *PodRecommendation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodRecommendation.Merge(dst, src)
}
func (m *PodRecommendation) XXX_Size() int {
	return xxx_messageInfo_PodRecommendation.Size(m)
}
func (m *PodRecommendation) XXX_DiscardUnknown() {
	xxx_messageInfo_PodRecommendation.DiscardUnknown(m)
}

var xxx_messageInfo_PodRecommendation proto.InternalMessageInfo

func (m *PodRecommendation) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *PodRecommendation) GetApplyRecommendationNow() bool {
	if m != nil {
		return m.ApplyRecommendationNow
	}
	return false
}

func (m *PodRecommendation) GetAssignPodPolicy() *AssignPodPolicy {
	if m != nil {
		return m.AssignPodPolicy
	}
	return nil
}

func (m *PodRecommendation) GetContainerRecommendations() []*ContainerRecommendation {
	if m != nil {
		return m.ContainerRecommendations
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerRecommendation)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation")
	proto.RegisterType((*AssignPodPolicy)(nil), "containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy")
	proto.RegisterType((*PodRecommendation)(nil), "containers_ai.alameda.v1alpha1.datahub.PodRecommendation")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/recommendation.proto", fileDescriptor_recommendation_626eeb1447246c93)
}

var fileDescriptor_recommendation_626eeb1447246c93 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x14, 0x6c, 0x93, 0x50, 0xa5, 0x0e, 0x10, 0xd5, 0x82, 0x76, 0x09, 0xaa, 0x5a, 0xe5, 0x80, 0x8a,
	0x90, 0x1c, 0x12, 0x10, 0x70, 0x43, 0x7c, 0x1c, 0x7a, 0x80, 0x28, 0x32, 0x48, 0x1c, 0x38, 0x58,
	0x2f, 0xbb, 0x26, 0xb5, 0x58, 0x7f, 0xb0, 0x76, 0x5a, 0x05, 0xf1, 0x27, 0xf8, 0x07, 0xfc, 0x4b,
	0xae, 0x68, 0x1d, 0x6f, 0xc4, 0x2e, 0x49, 0x59, 0xe5, 0x66, 0xaf, 0x3d, 0x33, 0x6f, 0xde, 0x8e,
	0x1f, 0x1a, 0x42, 0x0a, 0x92, 0x27, 0xc0, 0xc0, 0x88, 0xc1, 0xe5, 0x10, 0x52, 0x73, 0x01, 0xc3,
	0x41, 0x02, 0x0e, 0x2e, 0xe6, 0xd3, 0x41, 0xc6, 0x63, 0x2d, 0x25, 0x57, 0x09, 0x38, 0xa1, 0x15,
	0x31, 0x99, 0x76, 0x1a, 0x3f, 0x88, 0xb5, 0x72, 0x20, 0x14, 0xcf, 0x2c, 0x03, 0x41, 0x02, 0x01,
	0x29, 0xc0, 0x24, 0x80, 0x7b, 0x27, 0x33, 0xad, 0x67, 0x29, 0x1f, 0x78, 0xd4, 0x74, 0xfe, 0x65,
	0xe0, 0x84, 0xe4, 0xd6, 0x81, 0x34, 0x4b, 0xa2, 0xde, 0xa3, 0x6b, 0xb5, 0x25, 0x77, 0x90, 0xaf,
	0xc3, 0xe5, 0xeb, 0x0b, 0x35, 0x3a, 0x61, 0x60, 0xad, 0x98, 0x29, 0xc9, 0x95, 0x0b, 0x90, 0x87,
	0xff, 0xe3, 0xcf, 0x44, 0xbc, 0xbc, 0xda, 0xff, 0xdd, 0x44, 0x47, 0x6f, 0x0a, 0x5b, 0xb4, 0xe4,
	0x1a, 0x63, 0xd4, 0x52, 0x20, 0x79, 0xb4, 0x7b, 0xba, 0x7b, 0xb6, 0x4f, 0xfd, 0x1a, 0xcf, 0xd0,
	0xdd, 0x54, 0x48, 0xe1, 0x58, 0xb9, 0x43, 0x36, 0x6a, 0x9c, 0x36, 0xcf, 0x3a, 0xa3, 0x11, 0xa9,
	0xd7, 0x23, 0xf2, 0xde, 0x17, 0xf1, 0x16, 0x1c, 0xd0, 0x3b, 0x9e, 0xb0, 0xac, 0x6d, 0xf1, 0x57,
	0x74, 0x94, 0xf1, 0x6f, 0x73, 0x6e, 0xff, 0x95, 0x6a, 0x6e, 0x2d, 0x75, 0x18, 0x28, 0xab, 0x62,
	0x97, 0xe8, 0x58, 0x28, 0xe1, 0x04, 0xa4, 0x6c, 0xbd, 0xbb, 0xd6, 0xd6, 0x92, 0xf7, 0x03, 0xf1,
	0xbb, 0x75, 0x26, 0xbf, 0xa3, 0x93, 0x42, 0x77, 0x93, 0xd9, 0x1b, 0x5b, 0x2b, 0x17, 0x96, 0xe8,
	0x5a, 0xcf, 0xfd, 0x5f, 0x0d, 0xd4, 0x7d, 0xe5, 0x93, 0x33, 0xd1, 0xc9, 0x44, 0xa7, 0x22, 0x5e,
	0x60, 0x82, 0x5a, 0x79, 0x56, 0xfd, 0x1f, 0xef, 0x8c, 0x7a, 0x64, 0x19, 0x64, 0x52, 0x04, 0x99,
	0x7c, 0x2c, 0x82, 0x4c, 0xfd, 0x3d, 0xfc, 0x19, 0xdd, 0x52, 0x3a, 0xe1, 0xcc, 0x64, 0x42, 0x67,
	0xc2, 0x2d, 0xa2, 0x86, 0x07, 0x3e, 0xad, 0x5b, 0xed, 0x58, 0x27, 0x7c, 0x12, 0xb0, 0xe7, 0x3b,
	0xf4, 0xa6, 0xfa, 0x6b, 0x8f, 0x3f, 0x05, 0x72, 0xcb, 0x53, 0x1e, 0x3b, 0x9d, 0x45, 0x4d, 0x4f,
	0xfe, 0xb8, 0x2e, 0xf9, 0x87, 0x80, 0x2b, 0x88, 0x8b, 0x3d, 0x3e, 0x46, 0xfb, 0x9e, 0xd8, 0x87,
	0xbb, 0x95, 0x87, 0xfb, 0x7c, 0x87, 0xb6, 0xf3, 0x4f, 0x63, 0x90, 0xfc, 0x75, 0x1b, 0xed, 0x19,
	0xdf, 0x8e, 0xfe, 0xcf, 0x26, 0x3a, 0x98, 0xe8, 0xa4, 0xf2, 0x2c, 0x18, 0xea, 0xe6, 0x48, 0x6b,
	0x20, 0xe6, 0x09, 0x5b, 0xbd, 0x90, 0xce, 0xe8, 0x59, 0x6d, 0xdb, 0x2b, 0x78, 0xbe, 0xa2, 0xb7,
	0x55, 0x69, 0x8f, 0x5f, 0xa0, 0x08, 0x8c, 0x49, 0x17, 0x95, 0x2c, 0x30, 0xa5, 0xaf, 0x7c, 0x83,
	0xdb, 0xf4, 0xd0, 0x9f, 0x97, 0xeb, 0x1a, 0xeb, 0x2b, 0x1c, 0xa3, 0x83, 0xe5, 0x30, 0x60, 0xf9,
	0x5c, 0x58, 0xba, 0x08, 0x6d, 0x7b, 0x5e, 0xb7, 0xb8, 0x4a, 0x26, 0x68, 0x17, 0x2a, 0x21, 0xf9,
	0x81, 0xee, 0xad, 0xa8, 0x36, 0x3c, 0x94, 0x97, 0x75, 0xc5, 0x36, 0x8c, 0x1e, 0x1a, 0xc5, 0xeb,
	0x0f, 0xec, 0x74, 0xcf, 0x87, 0xf1, 0xc9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x3b, 0x2c,
	0xd9, 0xc0, 0x05, 0x00, 0x00,
}
