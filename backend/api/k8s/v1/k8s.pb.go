// Code generated by protoc-gen-go. DO NOT EDIT.
// source: k8s/v1/k8s.proto

package k8sv1

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#containerstate-v1-core
type Container_State int32

const (
	Container_UNSPECIFIED Container_State = 0
	Container_UNKNOWN     Container_State = 1
	Container_TERMINATED  Container_State = 2
	Container_RUNNING     Container_State = 3
	Container_WAITING     Container_State = 4
)

var Container_State_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "TERMINATED",
	3: "RUNNING",
	4: "WAITING",
}

var Container_State_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"TERMINATED":  2,
	"RUNNING":     3,
	"WAITING":     4,
}

func (x Container_State) String() string {
	return proto.EnumName(Container_State_name, int32(x))
}

func (Container_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{2, 0}
}

// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/
type Pod_State int32

const (
	Pod_UNSPECIFIED Pod_State = 0
	Pod_UNKNOWN     Pod_State = 1
	Pod_PENDING     Pod_State = 2
	Pod_RUNNING     Pod_State = 3
	Pod_SUCCEEDED   Pod_State = 4
	Pod_FAILED      Pod_State = 5
)

var Pod_State_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "RUNNING",
	4: "SUCCEEDED",
	5: "FAILED",
}

var Pod_State_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PENDING":     2,
	"RUNNING":     3,
	"SUCCEEDED":   4,
	"FAILED":      5,
}

func (x Pod_State) String() string {
	return proto.EnumName(Pod_State_name, int32(x))
}

func (Pod_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{3, 0}
}

type DescribePodRequest struct {
	Clientset            string            `protobuf:"bytes,1,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Cluster              string            `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace            string            `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DescribePodRequest) Reset()         { *m = DescribePodRequest{} }
func (m *DescribePodRequest) String() string { return proto.CompactTextString(m) }
func (*DescribePodRequest) ProtoMessage()    {}
func (*DescribePodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{0}
}

func (m *DescribePodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribePodRequest.Unmarshal(m, b)
}
func (m *DescribePodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribePodRequest.Marshal(b, m, deterministic)
}
func (m *DescribePodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribePodRequest.Merge(m, src)
}
func (m *DescribePodRequest) XXX_Size() int {
	return xxx_messageInfo_DescribePodRequest.Size(m)
}
func (m *DescribePodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribePodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribePodRequest proto.InternalMessageInfo

func (m *DescribePodRequest) GetClientset() string {
	if m != nil {
		return m.Clientset
	}
	return ""
}

func (m *DescribePodRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *DescribePodRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DescribePodRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DescribePodRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type DescribePodResponse struct {
	Pod                  *Pod     `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribePodResponse) Reset()         { *m = DescribePodResponse{} }
func (m *DescribePodResponse) String() string { return proto.CompactTextString(m) }
func (*DescribePodResponse) ProtoMessage()    {}
func (*DescribePodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{1}
}

func (m *DescribePodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribePodResponse.Unmarshal(m, b)
}
func (m *DescribePodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribePodResponse.Marshal(b, m, deterministic)
}
func (m *DescribePodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribePodResponse.Merge(m, src)
}
func (m *DescribePodResponse) XXX_Size() int {
	return xxx_messageInfo_DescribePodResponse.Size(m)
}
func (m *DescribePodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribePodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribePodResponse proto.InternalMessageInfo

func (m *DescribePodResponse) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

// TODO(maybe): Identify with resource annotations.
type Container struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image                string          `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	State                Container_State `protobuf:"varint,3,opt,name=state,proto3,enum=clutch.k8s.v1.Container_State" json:"state,omitempty"`
	Ready                bool            `protobuf:"varint,4,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{2}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetState() Container_State {
	if m != nil {
		return m.State
	}
	return Container_UNSPECIFIED
}

func (m *Container) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

type Pod struct {
	Cluster              string               `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace            string               `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Containers           []*Container         `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	NodeIp               string               `protobuf:"bytes,5,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
	PodIp                string               `protobuf:"bytes,6,opt,name=pod_ip,json=podIp,proto3" json:"pod_ip,omitempty"`
	State                Pod_State            `protobuf:"varint,7,opt,name=state,proto3,enum=clutch.k8s.v1.Pod_State" json:"state,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Labels               map[string]string    `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations          map[string]string    `protobuf:"bytes,10,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{3}
}

func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (m *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(m, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Pod) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Pod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pod) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Pod) GetNodeIp() string {
	if m != nil {
		return m.NodeIp
	}
	return ""
}

func (m *Pod) GetPodIp() string {
	if m != nil {
		return m.PodIp
	}
	return ""
}

func (m *Pod) GetState() Pod_State {
	if m != nil {
		return m.State
	}
	return Pod_UNSPECIFIED
}

func (m *Pod) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Pod) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Pod) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type ListPodsOptions struct {
	Labels               map[string]string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListPodsOptions) Reset()         { *m = ListPodsOptions{} }
func (m *ListPodsOptions) String() string { return proto.CompactTextString(m) }
func (*ListPodsOptions) ProtoMessage()    {}
func (*ListPodsOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{4}
}

func (m *ListPodsOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPodsOptions.Unmarshal(m, b)
}
func (m *ListPodsOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPodsOptions.Marshal(b, m, deterministic)
}
func (m *ListPodsOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPodsOptions.Merge(m, src)
}
func (m *ListPodsOptions) XXX_Size() int {
	return xxx_messageInfo_ListPodsOptions.Size(m)
}
func (m *ListPodsOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPodsOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ListPodsOptions proto.InternalMessageInfo

func (m *ListPodsOptions) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type DeletePodRequest struct {
	Clientset            string   `protobuf:"bytes,1,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Cluster              string   `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace            string   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePodRequest) Reset()         { *m = DeletePodRequest{} }
func (m *DeletePodRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePodRequest) ProtoMessage()    {}
func (*DeletePodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{5}
}

func (m *DeletePodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePodRequest.Unmarshal(m, b)
}
func (m *DeletePodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePodRequest.Marshal(b, m, deterministic)
}
func (m *DeletePodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePodRequest.Merge(m, src)
}
func (m *DeletePodRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePodRequest.Size(m)
}
func (m *DeletePodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePodRequest proto.InternalMessageInfo

func (m *DeletePodRequest) GetClientset() string {
	if m != nil {
		return m.Clientset
	}
	return ""
}

func (m *DeletePodRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *DeletePodRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeletePodRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeletePodResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePodResponse) Reset()         { *m = DeletePodResponse{} }
func (m *DeletePodResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePodResponse) ProtoMessage()    {}
func (*DeletePodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{6}
}

func (m *DeletePodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePodResponse.Unmarshal(m, b)
}
func (m *DeletePodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePodResponse.Marshal(b, m, deterministic)
}
func (m *DeletePodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePodResponse.Merge(m, src)
}
func (m *DeletePodResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePodResponse.Size(m)
}
func (m *DeletePodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePodResponse proto.InternalMessageInfo

type HPA struct {
	Cluster              string            `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace            string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Sizing               *HPA_Sizing       `protobuf:"bytes,4,opt,name=sizing,proto3" json:"sizing,omitempty"`
	Labels               map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations          map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HPA) Reset()         { *m = HPA{} }
func (m *HPA) String() string { return proto.CompactTextString(m) }
func (*HPA) ProtoMessage()    {}
func (*HPA) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{7}
}

func (m *HPA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HPA.Unmarshal(m, b)
}
func (m *HPA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HPA.Marshal(b, m, deterministic)
}
func (m *HPA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HPA.Merge(m, src)
}
func (m *HPA) XXX_Size() int {
	return xxx_messageInfo_HPA.Size(m)
}
func (m *HPA) XXX_DiscardUnknown() {
	xxx_messageInfo_HPA.DiscardUnknown(m)
}

var xxx_messageInfo_HPA proto.InternalMessageInfo

func (m *HPA) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *HPA) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *HPA) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HPA) GetSizing() *HPA_Sizing {
	if m != nil {
		return m.Sizing
	}
	return nil
}

func (m *HPA) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *HPA) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type HPA_Sizing struct {
	MinReplicas          uint32   `protobuf:"varint,1,opt,name=min_replicas,json=minReplicas,proto3" json:"min_replicas,omitempty"`
	MaxReplicas          uint32   `protobuf:"varint,2,opt,name=max_replicas,json=maxReplicas,proto3" json:"max_replicas,omitempty"`
	CurrentReplicas      uint32   `protobuf:"varint,3,opt,name=current_replicas,json=currentReplicas,proto3" json:"current_replicas,omitempty"`
	DesiredReplicas      uint32   `protobuf:"varint,4,opt,name=desired_replicas,json=desiredReplicas,proto3" json:"desired_replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HPA_Sizing) Reset()         { *m = HPA_Sizing{} }
func (m *HPA_Sizing) String() string { return proto.CompactTextString(m) }
func (*HPA_Sizing) ProtoMessage()    {}
func (*HPA_Sizing) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{7, 0}
}

func (m *HPA_Sizing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HPA_Sizing.Unmarshal(m, b)
}
func (m *HPA_Sizing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HPA_Sizing.Marshal(b, m, deterministic)
}
func (m *HPA_Sizing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HPA_Sizing.Merge(m, src)
}
func (m *HPA_Sizing) XXX_Size() int {
	return xxx_messageInfo_HPA_Sizing.Size(m)
}
func (m *HPA_Sizing) XXX_DiscardUnknown() {
	xxx_messageInfo_HPA_Sizing.DiscardUnknown(m)
}

var xxx_messageInfo_HPA_Sizing proto.InternalMessageInfo

func (m *HPA_Sizing) GetMinReplicas() uint32 {
	if m != nil {
		return m.MinReplicas
	}
	return 0
}

func (m *HPA_Sizing) GetMaxReplicas() uint32 {
	if m != nil {
		return m.MaxReplicas
	}
	return 0
}

func (m *HPA_Sizing) GetCurrentReplicas() uint32 {
	if m != nil {
		return m.CurrentReplicas
	}
	return 0
}

func (m *HPA_Sizing) GetDesiredReplicas() uint32 {
	if m != nil {
		return m.DesiredReplicas
	}
	return 0
}

type ResizeHPARequest struct {
	Clientset            string                   `protobuf:"bytes,1,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Cluster              string                   `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace            string                   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Sizing               *ResizeHPARequest_Sizing `protobuf:"bytes,5,opt,name=sizing,proto3" json:"sizing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ResizeHPARequest) Reset()         { *m = ResizeHPARequest{} }
func (m *ResizeHPARequest) String() string { return proto.CompactTextString(m) }
func (*ResizeHPARequest) ProtoMessage()    {}
func (*ResizeHPARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{8}
}

func (m *ResizeHPARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResizeHPARequest.Unmarshal(m, b)
}
func (m *ResizeHPARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResizeHPARequest.Marshal(b, m, deterministic)
}
func (m *ResizeHPARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResizeHPARequest.Merge(m, src)
}
func (m *ResizeHPARequest) XXX_Size() int {
	return xxx_messageInfo_ResizeHPARequest.Size(m)
}
func (m *ResizeHPARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResizeHPARequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResizeHPARequest proto.InternalMessageInfo

func (m *ResizeHPARequest) GetClientset() string {
	if m != nil {
		return m.Clientset
	}
	return ""
}

func (m *ResizeHPARequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *ResizeHPARequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ResizeHPARequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResizeHPARequest) GetSizing() *ResizeHPARequest_Sizing {
	if m != nil {
		return m.Sizing
	}
	return nil
}

type ResizeHPARequest_Sizing struct {
	Min                  uint32   `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max                  uint32   `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResizeHPARequest_Sizing) Reset()         { *m = ResizeHPARequest_Sizing{} }
func (m *ResizeHPARequest_Sizing) String() string { return proto.CompactTextString(m) }
func (*ResizeHPARequest_Sizing) ProtoMessage()    {}
func (*ResizeHPARequest_Sizing) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{8, 0}
}

func (m *ResizeHPARequest_Sizing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResizeHPARequest_Sizing.Unmarshal(m, b)
}
func (m *ResizeHPARequest_Sizing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResizeHPARequest_Sizing.Marshal(b, m, deterministic)
}
func (m *ResizeHPARequest_Sizing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResizeHPARequest_Sizing.Merge(m, src)
}
func (m *ResizeHPARequest_Sizing) XXX_Size() int {
	return xxx_messageInfo_ResizeHPARequest_Sizing.Size(m)
}
func (m *ResizeHPARequest_Sizing) XXX_DiscardUnknown() {
	xxx_messageInfo_ResizeHPARequest_Sizing.DiscardUnknown(m)
}

var xxx_messageInfo_ResizeHPARequest_Sizing proto.InternalMessageInfo

func (m *ResizeHPARequest_Sizing) GetMin() uint32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *ResizeHPARequest_Sizing) GetMax() uint32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type ResizeHPAResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResizeHPAResponse) Reset()         { *m = ResizeHPAResponse{} }
func (m *ResizeHPAResponse) String() string { return proto.CompactTextString(m) }
func (*ResizeHPAResponse) ProtoMessage()    {}
func (*ResizeHPAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_479456f95e4a3015, []int{9}
}

func (m *ResizeHPAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResizeHPAResponse.Unmarshal(m, b)
}
func (m *ResizeHPAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResizeHPAResponse.Marshal(b, m, deterministic)
}
func (m *ResizeHPAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResizeHPAResponse.Merge(m, src)
}
func (m *ResizeHPAResponse) XXX_Size() int {
	return xxx_messageInfo_ResizeHPAResponse.Size(m)
}
func (m *ResizeHPAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResizeHPAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResizeHPAResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("clutch.k8s.v1.Container_State", Container_State_name, Container_State_value)
	proto.RegisterEnum("clutch.k8s.v1.Pod_State", Pod_State_name, Pod_State_value)
	proto.RegisterType((*DescribePodRequest)(nil), "clutch.k8s.v1.DescribePodRequest")
	proto.RegisterMapType((map[string]string)(nil), "clutch.k8s.v1.DescribePodRequest.LabelsEntry")
	proto.RegisterType((*DescribePodResponse)(nil), "clutch.k8s.v1.DescribePodResponse")
	proto.RegisterType((*Container)(nil), "clutch.k8s.v1.Container")
	proto.RegisterType((*Pod)(nil), "clutch.k8s.v1.Pod")
	proto.RegisterMapType((map[string]string)(nil), "clutch.k8s.v1.Pod.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "clutch.k8s.v1.Pod.LabelsEntry")
	proto.RegisterType((*ListPodsOptions)(nil), "clutch.k8s.v1.ListPodsOptions")
	proto.RegisterMapType((map[string]string)(nil), "clutch.k8s.v1.ListPodsOptions.LabelsEntry")
	proto.RegisterType((*DeletePodRequest)(nil), "clutch.k8s.v1.DeletePodRequest")
	proto.RegisterType((*DeletePodResponse)(nil), "clutch.k8s.v1.DeletePodResponse")
	proto.RegisterType((*HPA)(nil), "clutch.k8s.v1.HPA")
	proto.RegisterMapType((map[string]string)(nil), "clutch.k8s.v1.HPA.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "clutch.k8s.v1.HPA.LabelsEntry")
	proto.RegisterType((*HPA_Sizing)(nil), "clutch.k8s.v1.HPA.Sizing")
	proto.RegisterType((*ResizeHPARequest)(nil), "clutch.k8s.v1.ResizeHPARequest")
	proto.RegisterType((*ResizeHPARequest_Sizing)(nil), "clutch.k8s.v1.ResizeHPARequest.Sizing")
	proto.RegisterType((*ResizeHPAResponse)(nil), "clutch.k8s.v1.ResizeHPAResponse")
}

func init() {
	proto.RegisterFile("k8s/v1/k8s.proto", fileDescriptor_479456f95e4a3015)
}

var fileDescriptor_479456f95e4a3015 = []byte{
	// 1080 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xc1, 0x6f, 0xe3, 0xc4,
	0x17, 0xfe, 0xd9, 0x8e, 0x9d, 0xe6, 0x65, 0x77, 0xeb, 0x4e, 0xf7, 0xc7, 0x9a, 0x10, 0xed, 0xb6,
	0x66, 0x41, 0x4b, 0x05, 0x8e, 0xd2, 0x45, 0x90, 0xed, 0x01, 0x29, 0x69, 0x5c, 0x1a, 0x6d, 0x49,
	0x23, 0xb7, 0xd5, 0x4a, 0x7b, 0xa9, 0xdc, 0x78, 0x28, 0x56, 0x13, 0xdb, 0x78, 0x9c, 0xa8, 0xed,
	0x6a, 0x2f, 0x1c, 0x39, 0x21, 0x71, 0xe3, 0xcc, 0x69, 0x8f, 0x48, 0xfc, 0x21, 0xc0, 0x89, 0x0b,
	0x52, 0x25, 0x6e, 0xf0, 0x17, 0xf4, 0x84, 0x66, 0x3c, 0x76, 0x1c, 0xa7, 0x2d, 0x4b, 0x17, 0x89,
	0xbd, 0xcd, 0x9b, 0xf9, 0xe6, 0xbd, 0x79, 0xdf, 0xfb, 0xde, 0xd3, 0x80, 0x7a, 0xd4, 0x20, 0xb5,
	0x71, 0xbd, 0x76, 0xd4, 0x20, 0x46, 0x10, 0xfa, 0x91, 0x8f, 0x6e, 0xf6, 0x07, 0xa3, 0xa8, 0xff,
	0x85, 0x41, 0x77, 0xc6, 0xf5, 0x4a, 0xf5, 0xd0, 0xf7, 0x0f, 0x07, 0xb8, 0x66, 0x07, 0x6e, 0xcd,
	0xf6, 0x3c, 0x3f, 0xb2, 0x23, 0xd7, 0xf7, 0x38, 0xb8, 0x72, 0x8f, 0x9f, 0x32, 0xeb, 0x60, 0xf4,
	0x79, 0x2d, 0x72, 0x87, 0x98, 0x44, 0xf6, 0x30, 0xe0, 0x80, 0x3b, 0x63, 0x7b, 0xe0, 0x3a, 0x76,
	0x84, 0x6b, 0xc9, 0x82, 0x1f, 0x68, 0xd4, 0xe1, 0xb8, 0x3e, 0xeb, 0x53, 0xff, 0x43, 0x04, 0xd4,
	0xc6, 0xa4, 0x1f, 0xba, 0x07, 0xb8, 0xe7, 0x3b, 0x16, 0xfe, 0x72, 0x84, 0x49, 0x84, 0xde, 0x81,
	0x52, 0x7f, 0xe0, 0x62, 0x2f, 0x22, 0x38, 0xd2, 0x84, 0x25, 0xe1, 0x41, 0xa9, 0x55, 0x3c, 0x6f,
	0x15, 0x42, 0x71, 0x49, 0xb0, 0x26, 0x27, 0x68, 0x19, 0x8a, 0xfd, 0xc1, 0x88, 0x44, 0x38, 0xd4,
	0xc4, 0x69, 0x50, 0xb2, 0x4f, 0x3d, 0x79, 0xf6, 0x10, 0x93, 0xc0, 0xee, 0x63, 0x4d, 0xca, 0x79,
	0x4a, 0x4f, 0xd0, 0x5b, 0x50, 0xa0, 0x86, 0x56, 0x98, 0x46, 0xb0, 0x4d, 0xb4, 0x03, 0xca, 0xc0,
	0x3e, 0xc0, 0x03, 0xa2, 0xc9, 0x4b, 0xd2, 0x83, 0xf2, 0xea, 0x07, 0xc6, 0x14, 0x6d, 0xc6, 0x6c,
	0x02, 0xc6, 0x16, 0xc3, 0x9b, 0x5e, 0x14, 0x9e, 0xb4, 0x6e, 0x9c, 0xb7, 0x4a, 0xdf, 0x09, 0x8a,
	0x1e, 0xbb, 0xe4, 0xae, 0x2a, 0x8f, 0xa0, 0x9c, 0x01, 0x21, 0x15, 0xa4, 0x23, 0x7c, 0x12, 0xe7,
	0x6a, 0xd1, 0x25, 0xba, 0x0d, 0xf2, 0xd8, 0x1e, 0x8c, 0x70, 0x9c, 0x9a, 0x15, 0x1b, 0x6b, 0x62,
	0x43, 0x58, 0xfb, 0xf8, 0x87, 0xb3, 0xea, 0x43, 0xa8, 0xc3, 0xc2, 0xf4, 0x33, 0x7a, 0xbe, 0x83,
	0xaa, 0xcf, 0x78, 0xe2, 0xcf, 0x6b, 0xcf, 0xd2, 0xec, 0xf8, 0xfa, 0xb9, 0xbe, 0x01, 0x8b, 0x53,
	0x6f, 0x25, 0x81, 0xef, 0x11, 0x8c, 0xee, 0x83, 0x14, 0xf8, 0x0e, 0x8b, 0x5d, 0x5e, 0x45, 0xc6,
	0x8c, 0x57, 0x8b, 0x1e, 0xaf, 0x95, 0x5e, 0x9c, 0x55, 0x65, 0x86, 0xd4, 0x7f, 0x16, 0xa0, 0xb4,
	0xee, 0x7b, 0x91, 0xed, 0x7a, 0x38, 0x44, 0x88, 0x73, 0x17, 0xbf, 0x3d, 0xa6, 0xec, 0x36, 0xc8,
	0xee, 0xd0, 0x3e, 0x4c, 0x1f, 0xcf, 0x0c, 0xf4, 0x21, 0xc8, 0x24, 0xb2, 0xa3, 0xb8, 0x10, 0xb7,
	0x56, 0xef, 0xe6, 0x42, 0xa5, 0x2e, 0x8d, 0x1d, 0x8a, 0xb2, 0x62, 0x30, 0xf5, 0x15, 0x62, 0xdb,
	0x39, 0x61, 0xc5, 0x99, 0xb3, 0x62, 0x43, 0xdf, 0x06, 0x99, 0xa1, 0xd0, 0x3c, 0x94, 0xf7, 0xba,
	0x3b, 0x3d, 0x73, 0xbd, 0xb3, 0xd1, 0x31, 0xdb, 0xea, 0xff, 0x50, 0x19, 0x8a, 0x7b, 0xdd, 0xc7,
	0xdd, 0xed, 0x27, 0x5d, 0x55, 0x40, 0xb7, 0x00, 0x76, 0x4d, 0xeb, 0xb3, 0x4e, 0xb7, 0xb9, 0x6b,
	0xb6, 0x55, 0x91, 0x1e, 0x5a, 0x7b, 0xdd, 0x6e, 0xa7, 0xfb, 0xa9, 0x2a, 0x51, 0xe3, 0x49, 0xb3,
	0xb3, 0x4b, 0x8d, 0x82, 0xfe, 0xa3, 0x0c, 0x12, 0xa5, 0x50, 0x9b, 0x88, 0x2a, 0xce, 0x28, 0xd5,
	0x52, 0x35, 0xab, 0xa5, 0x38, 0xb1, 0x8c, 0x84, 0x12, 0x1a, 0xa4, 0x0c, 0x0d, 0x0d, 0x80, 0x7e,
	0x92, 0x14, 0xd1, 0x0a, 0x4c, 0x3d, 0xda, 0x65, 0x59, 0x5b, 0x19, 0x2c, 0xba, 0x03, 0x45, 0xcf,
	0x77, 0xf0, 0xbe, 0x1b, 0x68, 0x32, 0x73, 0xa8, 0x50, 0xb3, 0x13, 0xa0, 0xff, 0x83, 0x12, 0xf8,
	0x0e, 0xdd, 0x57, 0x62, 0x6a, 0x03, 0xdf, 0xe9, 0x04, 0xc8, 0x48, 0xa8, 0x2d, 0x32, 0x6a, 0xb5,
	0xd9, 0x2a, 0x4e, 0x93, 0xfa, 0x08, 0x80, 0x44, 0x76, 0x18, 0xed, 0xd3, 0x26, 0xd6, 0xe6, 0x58,
	0xe9, 0x2b, 0x46, 0xdc, 0xe1, 0x46, 0xd2, 0xe1, 0xc6, 0x6e, 0xd2, 0xe1, 0x56, 0x89, 0xa1, 0xa9,
	0x8d, 0x3e, 0x4a, 0xdb, 0xa1, 0xc4, 0x12, 0xba, 0x7b, 0x41, 0xac, 0x8c, 0xb4, 0x13, 0xc5, 0x23,
	0x13, 0xca, 0x99, 0x01, 0xa0, 0x01, 0xbb, 0xfc, 0xf6, 0x05, 0x97, 0x9b, 0x13, 0x54, 0xec, 0x21,
	0x7b, 0xef, 0x15, 0x1a, 0xa7, 0xf2, 0x09, 0xa8, 0x79, 0xdf, 0xff, 0xe4, 0xbe, 0xfe, 0xf4, 0xe5,
	0x34, 0x57, 0x86, 0x62, 0xcf, 0xec, 0xb6, 0xa9, 0xac, 0x72, 0x82, 0xbb, 0x09, 0xa5, 0x9d, 0xbd,
	0xf5, 0x75, 0xd3, 0x6c, 0x9b, 0x6d, 0xb5, 0x80, 0x00, 0x94, 0x8d, 0x66, 0x67, 0xcb, 0x6c, 0xab,
	0xf2, 0xf5, 0x9b, 0xfa, 0x1b, 0x01, 0xe6, 0xb7, 0x5c, 0x12, 0xf5, 0x7c, 0x87, 0x6c, 0x07, 0x2c,
	0x33, 0xd4, 0xca, 0x95, 0x68, 0x25, 0xc7, 0x72, 0x0e, 0x7f, 0x51, 0xb9, 0x5e, 0x81, 0x67, 0xfd,
	0x57, 0x01, 0xd4, 0x36, 0x1e, 0xe0, 0xe8, 0xb5, 0x9d, 0xe9, 0xd7, 0xa7, 0x7b, 0x11, 0x16, 0x32,
	0xa9, 0xc5, 0x13, 0x54, 0xff, 0xad, 0x00, 0xd2, 0x66, 0xaf, 0xf9, 0xaf, 0xce, 0x8e, 0x3a, 0x28,
	0xc4, 0x3d, 0x75, 0xbd, 0x43, 0x96, 0x40, 0x79, 0xf5, 0xcd, 0x5c, 0x0d, 0x37, 0x7b, 0x4d, 0x63,
	0x87, 0x01, 0x2c, 0x0e, 0xcc, 0x74, 0xa6, 0x7c, 0x61, 0x67, 0xd2, 0x2b, 0x2f, 0xd1, 0x99, 0xca,
	0x85, 0x9d, 0x49, 0x2f, 0x5f, 0xdd, 0x99, 0xdf, 0x0b, 0xa0, 0xc4, 0x2f, 0x42, 0xcb, 0x70, 0x63,
	0xe8, 0x7a, 0xfb, 0x21, 0x0e, 0x06, 0x6e, 0xdf, 0x26, 0x8c, 0x8d, 0x9b, 0x56, 0x79, 0xe8, 0x7a,
	0x16, 0xdf, 0x62, 0x10, 0xfb, 0x78, 0x02, 0x11, 0x39, 0xc4, 0x3e, 0x4e, 0x21, 0xef, 0x81, 0xda,
	0x1f, 0x85, 0x21, 0xf6, 0xa2, 0x09, 0x4c, 0x62, 0xb0, 0x79, 0xbe, 0x9f, 0x85, 0x3a, 0x98, 0xb8,
	0x21, 0x76, 0x26, 0xd0, 0x42, 0x0c, 0xe5, 0xfb, 0x09, 0xf4, 0x3f, 0x1c, 0x20, 0x97, 0xa9, 0x8e,
	0x4a, 0xe7, 0x6a, 0xd5, 0xfd, 0x24, 0x82, 0x6a, 0x61, 0xe2, 0x9e, 0xe2, 0xcd, 0x5e, 0xf3, 0xf5,
	0xfc, 0x25, 0x6d, 0xa6, 0x7a, 0x95, 0x99, 0x5e, 0xdf, 0xcd, 0xe9, 0x27, 0xff, 0x7c, 0x2e, 0xde,
	0xd6, 0xdc, 0x79, 0x4b, 0xfe, 0x5a, 0x10, 0x55, 0x21, 0x91, 0x71, 0xe5, 0xfd, 0x54, 0x46, 0x2a,
	0x48, 0x43, 0xd7, 0xe3, 0xea, 0xa1, 0x4b, 0xb6, 0x63, 0x1f, 0x73, 0xb1, 0xd0, 0xe5, 0xf5, 0x39,
	0x5d, 0x84, 0x85, 0xcc, 0x9b, 0xe2, 0x4e, 0x5e, 0xfd, 0x53, 0x04, 0xe5, 0x71, 0x83, 0x34, 0x7b,
	0x1d, 0x74, 0x0a, 0xe5, 0xcc, 0x6f, 0x09, 0x2d, 0xff, 0xed, 0xaf, 0xaf, 0xa2, 0x5f, 0x05, 0xe1,
	0xa3, 0xe2, 0xfe, 0x8b, 0xb3, 0xaa, 0x38, 0x27, 0x7e, 0xf5, 0xcb, 0xef, 0xdf, 0x8a, 0x9a, 0xbe,
	0xc8, 0xff, 0xe3, 0x35, 0x67, 0x02, 0x5d, 0x13, 0x56, 0x50, 0x08, 0xa5, 0x74, 0xca, 0xa0, 0x7b,
	0x33, 0x6e, 0xa7, 0x47, 0x6b, 0x65, 0xe9, 0x72, 0x00, 0x8f, 0xaa, 0xb3, 0xa8, 0x05, 0x16, 0xf5,
	0x0d, 0x7d, 0x61, 0x12, 0x95, 0x03, 0x79, 0xcc, 0x94, 0x8f, 0x99, 0x98, 0xf9, 0xea, 0xcd, 0xc4,
	0x9c, 0xa1, 0x92, 0xc7, 0x94, 0xf2, 0x31, 0xc3, 0x04, 0xb8, 0x26, 0xac, 0xb4, 0x8a, 0x4f, 0xe5,
	0xa3, 0x06, 0x19, 0xd7, 0x0f, 0x14, 0xf6, 0xe7, 0x78, 0xf8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd8, 0xf0, 0x80, 0x8d, 0xa4, 0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// K8SAPIClient is the client API for K8SAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type K8SAPIClient interface {
	DescribePod(ctx context.Context, in *DescribePodRequest, opts ...grpc.CallOption) (*DescribePodResponse, error)
	DeletePod(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error)
	ResizeHPA(ctx context.Context, in *ResizeHPARequest, opts ...grpc.CallOption) (*ResizeHPAResponse, error)
}

type k8SAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewK8SAPIClient(cc grpc.ClientConnInterface) K8SAPIClient {
	return &k8SAPIClient{cc}
}

func (c *k8SAPIClient) DescribePod(ctx context.Context, in *DescribePodRequest, opts ...grpc.CallOption) (*DescribePodResponse, error) {
	out := new(DescribePodResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DescribePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeletePod(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error) {
	out := new(DeletePodResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeletePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) ResizeHPA(ctx context.Context, in *ResizeHPARequest, opts ...grpc.CallOption) (*ResizeHPAResponse, error) {
	out := new(ResizeHPAResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/ResizeHPA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K8SAPIServer is the server API for K8SAPI service.
type K8SAPIServer interface {
	DescribePod(context.Context, *DescribePodRequest) (*DescribePodResponse, error)
	DeletePod(context.Context, *DeletePodRequest) (*DeletePodResponse, error)
	ResizeHPA(context.Context, *ResizeHPARequest) (*ResizeHPAResponse, error)
}

// UnimplementedK8SAPIServer can be embedded to have forward compatible implementations.
type UnimplementedK8SAPIServer struct {
}

func (*UnimplementedK8SAPIServer) DescribePod(ctx context.Context, req *DescribePodRequest) (*DescribePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePod not implemented")
}
func (*UnimplementedK8SAPIServer) DeletePod(ctx context.Context, req *DeletePodRequest) (*DeletePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePod not implemented")
}
func (*UnimplementedK8SAPIServer) ResizeHPA(ctx context.Context, req *ResizeHPARequest) (*ResizeHPAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeHPA not implemented")
}

func RegisterK8SAPIServer(s *grpc.Server, srv K8SAPIServer) {
	s.RegisterService(&_K8SAPI_serviceDesc, srv)
}

func _K8SAPI_DescribePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DescribePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DescribePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DescribePod(ctx, req.(*DescribePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeletePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeletePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeletePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeletePod(ctx, req.(*DeletePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_ResizeHPA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeHPARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).ResizeHPA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/ResizeHPA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).ResizeHPA(ctx, req.(*ResizeHPARequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _K8SAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.k8s.v1.K8sAPI",
	HandlerType: (*K8SAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePod",
			Handler:    _K8SAPI_DescribePod_Handler,
		},
		{
			MethodName: "DeletePod",
			Handler:    _K8SAPI_DeletePod_Handler,
		},
		{
			MethodName: "ResizeHPA",
			Handler:    _K8SAPI_ResizeHPA_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "k8s/v1/k8s.proto",
}
