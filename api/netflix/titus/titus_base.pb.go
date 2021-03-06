// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netflix/titus/titus_base.proto

/*
Package titus is a generated protocol buffer package.

It is generated from these files:
	netflix/titus/titus_base.proto
	netflix/titus/titus_agent_api.proto
	netflix/titus/agent.proto

It has these top-level messages:
	CallMetadata
	Page
	Pagination
	RetryPolicy
	ResourceDimension
	InstanceLifecycleStatus
	InstanceOverrideStatus
	HealthStatus
	InstanceGroupLifecycleStatus
	AgentInstance
	AutoScaleRule
	AgentInstanceGroup
	AgentChangeEvent
	AgentQuery
	Id
	AgentInstanceGroups
	AgentInstances
	TierUpdate
	AutoScalingRuleUpdate
	InstanceGroupLifecycleStateUpdate
	InstanceGroupAttributesUpdate
	InstanceOverrideStateUpdate
	TaskStatusData
	StateUpdate
	StateUpdates
	TaskInfo
	ContainerInfo
*/
package titus

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// / Titus scheduler tier (see Titus scheduler documentation for more information).
type Tier int32

const (
	// / Tier for running latency insensitive batch workloads.
	Tier_Flex Tier = 0
	// / Tier for running latency sensitive workloads.
	Tier_Critical Tier = 1
)

var Tier_name = map[int32]string{
	0: "Flex",
	1: "Critical",
}
var Tier_value = map[string]int32{
	"Flex":     0,
	"Critical": 1,
}

func (x Tier) String() string {
	return proto.EnumName(Tier_name, int32(x))
}
func (Tier) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// / Disk mount permission mask
type MountPerm int32

const (
	// / Read only
	MountPerm_RO MountPerm = 0
	// Write only
	MountPerm_WO MountPerm = 1
	// Read/write
	MountPerm_RW MountPerm = 2
)

var MountPerm_name = map[int32]string{
	0: "RO",
	1: "WO",
	2: "RW",
}
var MountPerm_value = map[string]int32{
	"RO": 0,
	"WO": 1,
	"RW": 2,
}

func (x MountPerm) String() string {
	return proto.EnumName(MountPerm_name, int32(x))
}
func (MountPerm) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// / Call metadata (caller's identity, call path).
type CallMetadata struct {
	// / (Required) The original caller identifier. Depending on the user identity (a user or an application), the format
	// of the id may be different. For example a user's id could be encoded as 'user:<email-address>', and
	// application id as 'app:<aws-account-id>:<app-name>:<app-stack>'.
	CallerId string `protobuf:"bytes,1,opt,name=callerId" json:"callerId,omitempty"`
	// / (Optional) The reason why a call was made.
	CallReason string `protobuf:"bytes,2,opt,name=callReason" json:"callReason,omitempty"`
	// / (Optional) The list of intermediaries via which a call was relayed.
	CallPath []string `protobuf:"bytes,3,rep,name=callPath" json:"callPath,omitempty"`
	// / (Optional) If set to true, a diagnostic information is provided if a request fails.
	Debug bool `protobuf:"varint,4,opt,name=debug" json:"debug,omitempty"`
}

func (m *CallMetadata) Reset()                    { *m = CallMetadata{} }
func (m *CallMetadata) String() string            { return proto.CompactTextString(m) }
func (*CallMetadata) ProtoMessage()               {}
func (*CallMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CallMetadata) GetCallerId() string {
	if m != nil {
		return m.CallerId
	}
	return ""
}

func (m *CallMetadata) GetCallReason() string {
	if m != nil {
		return m.CallReason
	}
	return ""
}

func (m *CallMetadata) GetCallPath() []string {
	if m != nil {
		return m.CallPath
	}
	return nil
}

func (m *CallMetadata) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

// / An entity representing single page of a collection. Either pageNumber or cursor must be set on requests.
type Page struct {
	// / (Optional) Requested page number, starting from 0 (defaults to 0 if not specified).
	PageNumber int32 `protobuf:"varint,1,opt,name=pageNumber" json:"pageNumber,omitempty"`
	// / (Required) Requested page size (if not specified, default size is operation specific).
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize" json:"pageSize,omitempty"`
	// / (Optional) The position in the collection from which the next page should be returned. If the cursor value is set, it is
	// always used, irrespective of the page number set.
	Cursor string `protobuf:"bytes,3,opt,name=cursor" json:"cursor,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Page) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *Page) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Page) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

// / An entity representing pagination information returned to a client iterating over its elements.
// It includes current page that the client requested, and the total collection size.
// As not always pageCount * pageSize == itemCount, the item count is included as well.
type Pagination struct {
	// / (Required) Requested page details.
	CurrentPage *Page `protobuf:"bytes,1,opt,name=currentPage" json:"currentPage,omitempty"`
	// / Are there any more items to return.
	HasMore bool `protobuf:"varint,2,opt,name=hasMore" json:"hasMore,omitempty"`
	// / Total number of pages.
	TotalPages int32 `protobuf:"varint,3,opt,name=totalPages" json:"totalPages,omitempty"`
	// / Total number of items.
	TotalItems int32 `protobuf:"varint,4,opt,name=totalItems" json:"totalItems,omitempty"`
	// / The last retrieved item's position in the collection. The cursor value can be sent on a subsequent request to
	// get the next page of items. Using cursors, instead of page numbers, will guarantee that all items are
	// retrieved with a potential of items being duplicated.
	Cursor string `protobuf:"bytes,5,opt,name=cursor" json:"cursor,omitempty"`
	// / Position of the cursor relative to totalItems. It can be used to determine what pageNumber would overlap with a
	// cursor, or to provide an idea of progress when walking all pages. Valid values are [0, totalItems-1].
	CursorPosition int32 `protobuf:"varint,6,opt,name=cursorPosition" json:"cursorPosition,omitempty"`
}

func (m *Pagination) Reset()                    { *m = Pagination{} }
func (m *Pagination) String() string            { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()               {}
func (*Pagination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Pagination) GetCurrentPage() *Page {
	if m != nil {
		return m.CurrentPage
	}
	return nil
}

func (m *Pagination) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *Pagination) GetTotalPages() int32 {
	if m != nil {
		return m.TotalPages
	}
	return 0
}

func (m *Pagination) GetTotalItems() int32 {
	if m != nil {
		return m.TotalItems
	}
	return 0
}

func (m *Pagination) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *Pagination) GetCursorPosition() int32 {
	if m != nil {
		return m.CursorPosition
	}
	return 0
}

// / Retry polices.
type RetryPolicy struct {
	// / (Required) Retry policy.
	//
	// Types that are valid to be assigned to Policy:
	//	*RetryPolicy_Immediate_
	//	*RetryPolicy_Delayed_
	//	*RetryPolicy_ExponentialBackOff_
	Policy isRetryPolicy_Policy `protobuf_oneof:"Policy"`
}

func (m *RetryPolicy) Reset()                    { *m = RetryPolicy{} }
func (m *RetryPolicy) String() string            { return proto.CompactTextString(m) }
func (*RetryPolicy) ProtoMessage()               {}
func (*RetryPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isRetryPolicy_Policy interface {
	isRetryPolicy_Policy()
}

type RetryPolicy_Immediate_ struct {
	Immediate *RetryPolicy_Immediate `protobuf:"bytes,1,opt,name=immediate,oneof"`
}
type RetryPolicy_Delayed_ struct {
	Delayed *RetryPolicy_Delayed `protobuf:"bytes,2,opt,name=delayed,oneof"`
}
type RetryPolicy_ExponentialBackOff_ struct {
	ExponentialBackOff *RetryPolicy_ExponentialBackOff `protobuf:"bytes,3,opt,name=exponentialBackOff,oneof"`
}

func (*RetryPolicy_Immediate_) isRetryPolicy_Policy()          {}
func (*RetryPolicy_Delayed_) isRetryPolicy_Policy()            {}
func (*RetryPolicy_ExponentialBackOff_) isRetryPolicy_Policy() {}

func (m *RetryPolicy) GetPolicy() isRetryPolicy_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *RetryPolicy) GetImmediate() *RetryPolicy_Immediate {
	if x, ok := m.GetPolicy().(*RetryPolicy_Immediate_); ok {
		return x.Immediate
	}
	return nil
}

func (m *RetryPolicy) GetDelayed() *RetryPolicy_Delayed {
	if x, ok := m.GetPolicy().(*RetryPolicy_Delayed_); ok {
		return x.Delayed
	}
	return nil
}

func (m *RetryPolicy) GetExponentialBackOff() *RetryPolicy_ExponentialBackOff {
	if x, ok := m.GetPolicy().(*RetryPolicy_ExponentialBackOff_); ok {
		return x.ExponentialBackOff
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RetryPolicy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RetryPolicy_OneofMarshaler, _RetryPolicy_OneofUnmarshaler, _RetryPolicy_OneofSizer, []interface{}{
		(*RetryPolicy_Immediate_)(nil),
		(*RetryPolicy_Delayed_)(nil),
		(*RetryPolicy_ExponentialBackOff_)(nil),
	}
}

func _RetryPolicy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RetryPolicy)
	// Policy
	switch x := m.Policy.(type) {
	case *RetryPolicy_Immediate_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Immediate); err != nil {
			return err
		}
	case *RetryPolicy_Delayed_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Delayed); err != nil {
			return err
		}
	case *RetryPolicy_ExponentialBackOff_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExponentialBackOff); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RetryPolicy.Policy has unexpected type %T", x)
	}
	return nil
}

func _RetryPolicy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RetryPolicy)
	switch tag {
	case 1: // Policy.immediate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RetryPolicy_Immediate)
		err := b.DecodeMessage(msg)
		m.Policy = &RetryPolicy_Immediate_{msg}
		return true, err
	case 2: // Policy.delayed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RetryPolicy_Delayed)
		err := b.DecodeMessage(msg)
		m.Policy = &RetryPolicy_Delayed_{msg}
		return true, err
	case 3: // Policy.exponentialBackOff
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RetryPolicy_ExponentialBackOff)
		err := b.DecodeMessage(msg)
		m.Policy = &RetryPolicy_ExponentialBackOff_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RetryPolicy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RetryPolicy)
	// Policy
	switch x := m.Policy.(type) {
	case *RetryPolicy_Immediate_:
		s := proto.Size(x.Immediate)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RetryPolicy_Delayed_:
		s := proto.Size(x.Delayed)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RetryPolicy_ExponentialBackOff_:
		s := proto.Size(x.ExponentialBackOff)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// / Re-run immediately.
type RetryPolicy_Immediate struct {
	// / Maximum number of retries.
	Retries uint32 `protobuf:"varint,1,opt,name=retries" json:"retries,omitempty"`
}

func (m *RetryPolicy_Immediate) Reset()                    { *m = RetryPolicy_Immediate{} }
func (m *RetryPolicy_Immediate) String() string            { return proto.CompactTextString(m) }
func (*RetryPolicy_Immediate) ProtoMessage()               {}
func (*RetryPolicy_Immediate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *RetryPolicy_Immediate) GetRetries() uint32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

// / Re-run after a given delay.
type RetryPolicy_Delayed struct {
	// / Initial delay in milliseconds.
	InitialDelayMs uint64 `protobuf:"varint,1,opt,name=initialDelayMs" json:"initialDelayMs,omitempty"`
	// / Subsequenet delays in milliseconds.
	DelayMs uint64 `protobuf:"varint,2,opt,name=delayMs" json:"delayMs,omitempty"`
	// / Maximum number of retries.
	Retries uint32 `protobuf:"varint,3,opt,name=retries" json:"retries,omitempty"`
}

func (m *RetryPolicy_Delayed) Reset()                    { *m = RetryPolicy_Delayed{} }
func (m *RetryPolicy_Delayed) String() string            { return proto.CompactTextString(m) }
func (*RetryPolicy_Delayed) ProtoMessage()               {}
func (*RetryPolicy_Delayed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 1} }

func (m *RetryPolicy_Delayed) GetInitialDelayMs() uint64 {
	if m != nil {
		return m.InitialDelayMs
	}
	return 0
}

func (m *RetryPolicy_Delayed) GetDelayMs() uint64 {
	if m != nil {
		return m.DelayMs
	}
	return 0
}

func (m *RetryPolicy_Delayed) GetRetries() uint32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

// / Exponential back-off retry policy.
type RetryPolicy_ExponentialBackOff struct {
	// / Initial delay in milliseconds.
	InitialDelayMs uint64 `protobuf:"varint,1,opt,name=initialDelayMs" json:"initialDelayMs,omitempty"`
	// / Upper bound on delay interval.
	MaxDelayIntervalMs uint64 `protobuf:"varint,2,opt,name=maxDelayIntervalMs" json:"maxDelayIntervalMs,omitempty"`
	// / Maximum number of retries.
	Retries uint32 `protobuf:"varint,3,opt,name=retries" json:"retries,omitempty"`
}

func (m *RetryPolicy_ExponentialBackOff) Reset()         { *m = RetryPolicy_ExponentialBackOff{} }
func (m *RetryPolicy_ExponentialBackOff) String() string { return proto.CompactTextString(m) }
func (*RetryPolicy_ExponentialBackOff) ProtoMessage()    {}
func (*RetryPolicy_ExponentialBackOff) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 2}
}

func (m *RetryPolicy_ExponentialBackOff) GetInitialDelayMs() uint64 {
	if m != nil {
		return m.InitialDelayMs
	}
	return 0
}

func (m *RetryPolicy_ExponentialBackOff) GetMaxDelayIntervalMs() uint64 {
	if m != nil {
		return m.MaxDelayIntervalMs
	}
	return 0
}

func (m *RetryPolicy_ExponentialBackOff) GetRetries() uint32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

// / Instance resource dimensions
type ResourceDimension struct {
	Cpu         uint32 `protobuf:"varint,1,opt,name=cpu" json:"cpu,omitempty"`
	Gpu         uint32 `protobuf:"varint,2,opt,name=gpu" json:"gpu,omitempty"`
	MemoryMB    uint32 `protobuf:"varint,3,opt,name=memoryMB" json:"memoryMB,omitempty"`
	DiskMB      uint32 `protobuf:"varint,4,opt,name=diskMB" json:"diskMB,omitempty"`
	NetworkMbps uint32 `protobuf:"varint,5,opt,name=networkMbps" json:"networkMbps,omitempty"`
}

func (m *ResourceDimension) Reset()                    { *m = ResourceDimension{} }
func (m *ResourceDimension) String() string            { return proto.CompactTextString(m) }
func (*ResourceDimension) ProtoMessage()               {}
func (*ResourceDimension) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResourceDimension) GetCpu() uint32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *ResourceDimension) GetGpu() uint32 {
	if m != nil {
		return m.Gpu
	}
	return 0
}

func (m *ResourceDimension) GetMemoryMB() uint32 {
	if m != nil {
		return m.MemoryMB
	}
	return 0
}

func (m *ResourceDimension) GetDiskMB() uint32 {
	if m != nil {
		return m.DiskMB
	}
	return 0
}

func (m *ResourceDimension) GetNetworkMbps() uint32 {
	if m != nil {
		return m.NetworkMbps
	}
	return 0
}

func init() {
	proto.RegisterType((*CallMetadata)(nil), "com.netflix.titus.CallMetadata")
	proto.RegisterType((*Page)(nil), "com.netflix.titus.Page")
	proto.RegisterType((*Pagination)(nil), "com.netflix.titus.Pagination")
	proto.RegisterType((*RetryPolicy)(nil), "com.netflix.titus.RetryPolicy")
	proto.RegisterType((*RetryPolicy_Immediate)(nil), "com.netflix.titus.RetryPolicy.Immediate")
	proto.RegisterType((*RetryPolicy_Delayed)(nil), "com.netflix.titus.RetryPolicy.Delayed")
	proto.RegisterType((*RetryPolicy_ExponentialBackOff)(nil), "com.netflix.titus.RetryPolicy.ExponentialBackOff")
	proto.RegisterType((*ResourceDimension)(nil), "com.netflix.titus.ResourceDimension")
	proto.RegisterEnum("com.netflix.titus.Tier", Tier_name, Tier_value)
	proto.RegisterEnum("com.netflix.titus.MountPerm", MountPerm_name, MountPerm_value)
}

func init() { proto.RegisterFile("netflix/titus/titus_base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6f, 0x1a, 0x3b,
	0x10, 0x66, 0xf9, 0x15, 0x18, 0xc2, 0x13, 0xb1, 0x9e, 0xde, 0xa3, 0x1c, 0x52, 0x44, 0xd5, 0x28,
	0xca, 0x81, 0xa8, 0xa9, 0x7a, 0xe8, 0x75, 0x93, 0x56, 0x70, 0xa0, 0x41, 0x6e, 0xa4, 0x48, 0xb9,
	0x54, 0x66, 0x77, 0xd8, 0x58, 0xd9, 0xb5, 0x91, 0xed, 0x6d, 0xa1, 0xea, 0xb9, 0xb7, 0xfe, 0x11,
	0xfd, 0xd3, 0xfa, 0x9f, 0x54, 0xf6, 0xee, 0xd2, 0x6d, 0x48, 0x95, 0x5e, 0xc0, 0xdf, 0x37, 0xdf,
	0xcc, 0x37, 0x33, 0xd8, 0xc0, 0xa1, 0x40, 0xb3, 0x8c, 0xf9, 0xfa, 0xd4, 0x70, 0x93, 0xea, 0xec,
	0xf3, 0xc3, 0x82, 0x69, 0x1c, 0xaf, 0x94, 0x34, 0x92, 0x1c, 0x04, 0x32, 0x19, 0xe7, 0x9a, 0xb1,
	0x8b, 0x0e, 0x9e, 0x44, 0x52, 0x46, 0x31, 0x9e, 0x3a, 0xc1, 0x22, 0x5d, 0x9e, 0x32, 0xb1, 0xc9,
	0xd4, 0xa3, 0x2f, 0xb0, 0x7f, 0xce, 0xe2, 0x78, 0x86, 0x86, 0x85, 0xcc, 0x30, 0x32, 0x80, 0x56,
	0xc0, 0xe2, 0x18, 0xd5, 0x34, 0xec, 0x7b, 0x43, 0xef, 0xb8, 0x4d, 0xb7, 0x98, 0x1c, 0x02, 0xd8,
	0x33, 0x45, 0xa6, 0xa5, 0xe8, 0x57, 0x5d, 0xb4, 0xc4, 0x14, 0xb9, 0x73, 0x66, 0x6e, 0xfb, 0xb5,
	0x61, 0xad, 0xc8, 0xb5, 0x98, 0xfc, 0x0b, 0x8d, 0x10, 0x17, 0x69, 0xd4, 0xaf, 0x0f, 0xbd, 0xe3,
	0x16, 0xcd, 0xc0, 0xe8, 0x06, 0xea, 0x73, 0x16, 0xa1, 0xad, 0xbc, 0x62, 0x11, 0xbe, 0x4b, 0x93,
	0x05, 0x2a, 0xe7, 0xdb, 0xa0, 0x25, 0xc6, 0x56, 0xb6, 0xe8, 0x3d, 0xff, 0x8c, 0xce, 0xb7, 0x41,
	0xb7, 0x98, 0xfc, 0x07, 0xcd, 0x20, 0x55, 0x5a, 0xaa, 0x7e, 0xcd, 0x75, 0x94, 0xa3, 0xd1, 0x0f,
	0x0f, 0x60, 0xce, 0x22, 0x2e, 0x98, 0xe1, 0x52, 0x90, 0xd7, 0xd0, 0x09, 0x52, 0xa5, 0x50, 0x18,
	0xeb, 0xe8, 0x3c, 0x3a, 0x67, 0xff, 0x8f, 0x77, 0x96, 0x35, 0xb6, 0x61, 0x5a, 0xd6, 0x92, 0x3e,
	0xec, 0xdd, 0x32, 0x3d, 0x93, 0x2a, 0x33, 0x6f, 0xd1, 0x02, 0xda, 0xbe, 0x8d, 0x34, 0x2c, 0xb6,
	0x32, 0xed, 0xfc, 0x1b, 0xb4, 0xc4, 0x6c, 0xe3, 0x53, 0x83, 0x89, 0x76, 0xa3, 0x17, 0x71, 0xc7,
	0x94, 0x7a, 0x6f, 0x94, 0x7b, 0x27, 0x47, 0xf0, 0x4f, 0x76, 0x9a, 0x4b, 0xcd, 0x6d, 0xfb, 0xfd,
	0xa6, 0xcb, 0xbd, 0xc7, 0x8e, 0xbe, 0xd7, 0xa1, 0x43, 0xd1, 0xa8, 0xcd, 0x5c, 0xc6, 0x3c, 0xd8,
	0x90, 0x09, 0xb4, 0x79, 0x92, 0x60, 0xc8, 0x99, 0x29, 0x46, 0x3c, 0x7e, 0x60, 0xc4, 0x52, 0xca,
	0x78, 0x5a, 0xe8, 0x27, 0x15, 0xfa, 0x2b, 0x99, 0xf8, 0xb0, 0x17, 0x62, 0xcc, 0x36, 0x18, 0xba,
	0x99, 0x3b, 0x67, 0x47, 0x8f, 0xd4, 0xb9, 0xc8, 0xd4, 0x93, 0x0a, 0x2d, 0x12, 0x49, 0x00, 0x04,
	0xd7, 0x2b, 0x29, 0x50, 0x18, 0xce, 0x62, 0x9f, 0x05, 0x77, 0x97, 0xcb, 0xa5, 0xdb, 0x52, 0xe7,
	0xec, 0xc5, 0x23, 0xe5, 0xde, 0xec, 0x24, 0x4e, 0x2a, 0xf4, 0x81, 0x72, 0x83, 0xe7, 0xd0, 0xde,
	0x8e, 0x60, 0x7f, 0x29, 0x85, 0x46, 0x71, 0xd4, 0x6e, 0xfa, 0x2e, 0x2d, 0xe0, 0x00, 0x61, 0x2f,
	0xef, 0xd0, 0x2e, 0x97, 0x0b, 0x6e, 0x6b, 0x38, 0x66, 0x96, 0x69, 0xeb, 0xf4, 0x1e, 0x6b, 0x8b,
	0x85, 0xb9, 0xa0, 0xea, 0x04, 0x05, 0x2c, 0xdb, 0xd4, 0x7e, 0xb7, 0xf9, 0xea, 0x01, 0xd9, 0x6d,
	0xfd, 0xaf, 0x2d, 0xc7, 0x40, 0x12, 0xb6, 0x76, 0x68, 0x2a, 0x0c, 0xaa, 0x8f, 0x2c, 0xde, 0xba,
	0x3f, 0x10, 0xf9, 0x73, 0x23, 0x7e, 0x0b, 0x9a, 0xd9, 0x26, 0x47, 0xdf, 0x3c, 0x38, 0xa0, 0xa8,
	0x65, 0xaa, 0x02, 0xbc, 0xe0, 0x09, 0x0a, 0x6d, 0x9f, 0x43, 0x0f, 0x6a, 0xc1, 0x2a, 0xcd, 0xb7,
	0x64, 0x8f, 0x96, 0x89, 0x56, 0xa9, 0x33, 0xeb, 0x52, 0x7b, 0xb4, 0xaf, 0x2e, 0xc1, 0x44, 0xaa,
	0xcd, 0xcc, 0xcf, 0xcb, 0x6f, 0xb1, 0xbd, 0xb9, 0x21, 0xd7, 0x77, 0x33, 0xdf, 0xdd, 0xea, 0x2e,
	0xcd, 0x11, 0x19, 0x42, 0x47, 0xa0, 0xf9, 0x24, 0xd5, 0xdd, 0x6c, 0xb1, 0xd2, 0xee, 0x5a, 0x77,
	0x69, 0x99, 0x3a, 0x39, 0x84, 0xfa, 0x15, 0x47, 0x45, 0x5a, 0x50, 0x7f, 0x1b, 0xe3, 0xba, 0x57,
	0x21, 0xfb, 0xd0, 0x3a, 0x57, 0xdc, 0xf0, 0x80, 0xc5, 0x3d, 0xef, 0xe4, 0x19, 0xb4, 0x67, 0x32,
	0x15, 0x66, 0x8e, 0x2a, 0x21, 0x4d, 0xa8, 0xd2, 0xcb, 0x5e, 0xc5, 0x7e, 0x5f, 0x5f, 0xf6, 0x3c,
	0x87, 0xaf, 0x7b, 0x55, 0xff, 0x15, 0x3c, 0xdd, 0xbd, 0x3f, 0x91, 0x5a, 0x05, 0xd9, 0xdf, 0x5a,
	0x84, 0xc2, 0x6f, 0x5f, 0x59, 0xd2, 0x67, 0x1a, 0xe7, 0xde, 0x4d, 0xc3, 0x29, 0x16, 0x4d, 0x17,
	0x7d, 0xf9, 0x33, 0x00, 0x00, 0xff, 0xff, 0x58, 0x94, 0x86, 0x39, 0x44, 0x05, 0x00, 0x00,
}
