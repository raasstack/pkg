// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/messaging.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message is the container for all message types
type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Type  string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//
	//	*Message_Command
	//	*Message_Heartbeat
	//	*Message_ClientInfo
	//	*Message_CommandResult
	Payload       isMessage_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_proto_messaging_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaging_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_messaging_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetPayload() isMessage_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Message) GetCommand() *Command {
	if x != nil {
		if x, ok := x.Payload.(*Message_Command); ok {
			return x.Command
		}
	}
	return nil
}

func (x *Message) GetHeartbeat() *Heartbeat {
	if x != nil {
		if x, ok := x.Payload.(*Message_Heartbeat); ok {
			return x.Heartbeat
		}
	}
	return nil
}

func (x *Message) GetClientInfo() *ClientInfo {
	if x != nil {
		if x, ok := x.Payload.(*Message_ClientInfo); ok {
			return x.ClientInfo
		}
	}
	return nil
}

func (x *Message) GetCommandResult() *CommandResult {
	if x != nil {
		if x, ok := x.Payload.(*Message_CommandResult); ok {
			return x.CommandResult
		}
	}
	return nil
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Command struct {
	Command *Command `protobuf:"bytes,2,opt,name=command,proto3,oneof"`
}

type Message_Heartbeat struct {
	Heartbeat *Heartbeat `protobuf:"bytes,3,opt,name=heartbeat,proto3,oneof"`
}

type Message_ClientInfo struct {
	ClientInfo *ClientInfo `protobuf:"bytes,4,opt,name=client_info,json=clientInfo,proto3,oneof"`
}

type Message_CommandResult struct {
	CommandResult *CommandResult `protobuf:"bytes,5,opt,name=command_result,json=commandResult,proto3,oneof"`
}

func (*Message_Command) isMessage_Payload() {}

func (*Message_Heartbeat) isMessage_Payload() {}

func (*Message_ClientInfo) isMessage_Payload() {}

func (*Message_CommandResult) isMessage_Payload() {}

// Command represents a command to be executed by a client
type Command struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Command       string                 `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Args          []string               `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	RequirePty    bool                   `protobuf:"varint,4,opt,name=require_pty,json=requirePty,proto3" json:"require_pty,omitempty"`
	SenderId      string                 `protobuf:"bytes,5,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Timestamp     int64                  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Command) Reset() {
	*x = Command{}
	mi := &file_proto_messaging_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaging_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_proto_messaging_proto_rawDescGZIP(), []int{1}
}

func (x *Command) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Command) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Command) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Command) GetRequirePty() bool {
	if x != nil {
		return x.RequirePty
	}
	return false
}

func (x *Command) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Command) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// CommandResult represents the result of a command execution
type CommandResult struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommandId       string                 `protobuf:"bytes,2,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	ClientId        string                 `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Success         bool                   `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
	Output          string                 `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	Error           string                 `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	ExecutionTimeMs int64                  `protobuf:"varint,7,opt,name=execution_time_ms,json=executionTimeMs,proto3" json:"execution_time_ms,omitempty"`
	Timestamp       int64                  `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CommandResult) Reset() {
	*x = CommandResult{}
	mi := &file_proto_messaging_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResult) ProtoMessage() {}

func (x *CommandResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaging_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResult.ProtoReflect.Descriptor instead.
func (*CommandResult) Descriptor() ([]byte, []int) {
	return file_proto_messaging_proto_rawDescGZIP(), []int{2}
}

func (x *CommandResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommandResult) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *CommandResult) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CommandResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CommandResult) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *CommandResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CommandResult) GetExecutionTimeMs() int64 {
	if x != nil {
		return x.ExecutionTimeMs
	}
	return 0
}

func (x *CommandResult) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// ClientInfo represents information about a client
type ClientInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hostname      string                 `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Os            string                 `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
	OsVersion     string                 `protobuf:"bytes,4,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	Architecture  string                 `protobuf:"bytes,5,opt,name=architecture,proto3" json:"architecture,omitempty"`
	IpAddress     string                 `protobuf:"bytes,6,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	ConnectedAt   int64                  `protobuf:"varint,7,opt,name=connected_at,json=connectedAt,proto3" json:"connected_at,omitempty"`
	LastSeenAt    int64                  `protobuf:"varint,8,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
	Capabilities  []string               `protobuf:"bytes,9,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	Metrics       *SystemMetrics         `protobuf:"bytes,10,opt,name=metrics,proto3" json:"metrics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	mi := &file_proto_messaging_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaging_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_proto_messaging_proto_rawDescGZIP(), []int{3}
}

func (x *ClientInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ClientInfo) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *ClientInfo) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *ClientInfo) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *ClientInfo) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *ClientInfo) GetConnectedAt() int64 {
	if x != nil {
		return x.ConnectedAt
	}
	return 0
}

func (x *ClientInfo) GetLastSeenAt() int64 {
	if x != nil {
		return x.LastSeenAt
	}
	return 0
}

func (x *ClientInfo) GetCapabilities() []string {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *ClientInfo) GetMetrics() *SystemMetrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

// Heartbeat represents a client heartbeat
type Heartbeat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Timestamp     int64                  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Metrics       *SystemMetrics         `protobuf:"bytes,3,opt,name=metrics,proto3" json:"metrics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	mi := &file_proto_messaging_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaging_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_proto_messaging_proto_rawDescGZIP(), []int{4}
}

func (x *Heartbeat) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Heartbeat) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Heartbeat) GetMetrics() *SystemMetrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

// SystemMetrics represents system metrics
type SystemMetrics struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	CpuUsagePercent    float64                `protobuf:"fixed64,1,opt,name=cpu_usage_percent,json=cpuUsagePercent,proto3" json:"cpu_usage_percent,omitempty"`
	MemoryUsagePercent float64                `protobuf:"fixed64,2,opt,name=memory_usage_percent,json=memoryUsagePercent,proto3" json:"memory_usage_percent,omitempty"`
	TotalMemoryBytes   int64                  `protobuf:"varint,3,opt,name=total_memory_bytes,json=totalMemoryBytes,proto3" json:"total_memory_bytes,omitempty"`
	UsedMemoryBytes    int64                  `protobuf:"varint,4,opt,name=used_memory_bytes,json=usedMemoryBytes,proto3" json:"used_memory_bytes,omitempty"`
	FreeMemoryBytes    int64                  `protobuf:"varint,5,opt,name=free_memory_bytes,json=freeMemoryBytes,proto3" json:"free_memory_bytes,omitempty"`
	LoadAvg_1M         float64                `protobuf:"fixed64,6,opt,name=load_avg_1m,json=loadAvg1m,proto3" json:"load_avg_1m,omitempty"`
	LoadAvg_5M         float64                `protobuf:"fixed64,7,opt,name=load_avg_5m,json=loadAvg5m,proto3" json:"load_avg_5m,omitempty"`
	LoadAvg_15M        float64                `protobuf:"fixed64,8,opt,name=load_avg_15m,json=loadAvg15m,proto3" json:"load_avg_15m,omitempty"`
	UptimeSeconds      int64                  `protobuf:"varint,9,opt,name=uptime_seconds,json=uptimeSeconds,proto3" json:"uptime_seconds,omitempty"`
	Disk               *DiskMetrics           `protobuf:"bytes,10,opt,name=disk,proto3" json:"disk,omitempty"`
	Network            *NetworkMetrics        `protobuf:"bytes,11,opt,name=network,proto3" json:"network,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *SystemMetrics) Reset() {
	*x = SystemMetrics{}
	mi := &file_proto_messaging_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SystemMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemMetrics) ProtoMessage() {}

func (x *SystemMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaging_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemMetrics.ProtoReflect.Descriptor instead.
func (*SystemMetrics) Descriptor() ([]byte, []int) {
	return file_proto_messaging_proto_rawDescGZIP(), []int{5}
}

func (x *SystemMetrics) GetCpuUsagePercent() float64 {
	if x != nil {
		return x.CpuUsagePercent
	}
	return 0
}

func (x *SystemMetrics) GetMemoryUsagePercent() float64 {
	if x != nil {
		return x.MemoryUsagePercent
	}
	return 0
}

func (x *SystemMetrics) GetTotalMemoryBytes() int64 {
	if x != nil {
		return x.TotalMemoryBytes
	}
	return 0
}

func (x *SystemMetrics) GetUsedMemoryBytes() int64 {
	if x != nil {
		return x.UsedMemoryBytes
	}
	return 0
}

func (x *SystemMetrics) GetFreeMemoryBytes() int64 {
	if x != nil {
		return x.FreeMemoryBytes
	}
	return 0
}

func (x *SystemMetrics) GetLoadAvg_1M() float64 {
	if x != nil {
		return x.LoadAvg_1M
	}
	return 0
}

func (x *SystemMetrics) GetLoadAvg_5M() float64 {
	if x != nil {
		return x.LoadAvg_5M
	}
	return 0
}

func (x *SystemMetrics) GetLoadAvg_15M() float64 {
	if x != nil {
		return x.LoadAvg_15M
	}
	return 0
}

func (x *SystemMetrics) GetUptimeSeconds() int64 {
	if x != nil {
		return x.UptimeSeconds
	}
	return 0
}

func (x *SystemMetrics) GetDisk() *DiskMetrics {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *SystemMetrics) GetNetwork() *NetworkMetrics {
	if x != nil {
		return x.Network
	}
	return nil
}

// DiskMetrics represents disk usage metrics
type DiskMetrics struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	TotalDiskBytes   int64                  `protobuf:"varint,1,opt,name=total_disk_bytes,json=totalDiskBytes,proto3" json:"total_disk_bytes,omitempty"`
	UsedDiskBytes    int64                  `protobuf:"varint,2,opt,name=used_disk_bytes,json=usedDiskBytes,proto3" json:"used_disk_bytes,omitempty"`
	FreeDiskBytes    int64                  `protobuf:"varint,3,opt,name=free_disk_bytes,json=freeDiskBytes,proto3" json:"free_disk_bytes,omitempty"`
	DiskUsagePercent float64                `protobuf:"fixed64,4,opt,name=disk_usage_percent,json=diskUsagePercent,proto3" json:"disk_usage_percent,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *DiskMetrics) Reset() {
	*x = DiskMetrics{}
	mi := &file_proto_messaging_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiskMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskMetrics) ProtoMessage() {}

func (x *DiskMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaging_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskMetrics.ProtoReflect.Descriptor instead.
func (*DiskMetrics) Descriptor() ([]byte, []int) {
	return file_proto_messaging_proto_rawDescGZIP(), []int{6}
}

func (x *DiskMetrics) GetTotalDiskBytes() int64 {
	if x != nil {
		return x.TotalDiskBytes
	}
	return 0
}

func (x *DiskMetrics) GetUsedDiskBytes() int64 {
	if x != nil {
		return x.UsedDiskBytes
	}
	return 0
}

func (x *DiskMetrics) GetFreeDiskBytes() int64 {
	if x != nil {
		return x.FreeDiskBytes
	}
	return 0
}

func (x *DiskMetrics) GetDiskUsagePercent() float64 {
	if x != nil {
		return x.DiskUsagePercent
	}
	return 0
}

// NetworkMetrics represents network usage metrics
type NetworkMetrics struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	BytesSent       int64                  `protobuf:"varint,1,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	BytesReceived   int64                  `protobuf:"varint,2,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	PacketsSent     int64                  `protobuf:"varint,3,opt,name=packets_sent,json=packetsSent,proto3" json:"packets_sent,omitempty"`
	PacketsReceived int64                  `protobuf:"varint,4,opt,name=packets_received,json=packetsReceived,proto3" json:"packets_received,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *NetworkMetrics) Reset() {
	*x = NetworkMetrics{}
	mi := &file_proto_messaging_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkMetrics) ProtoMessage() {}

func (x *NetworkMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messaging_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkMetrics.ProtoReflect.Descriptor instead.
func (*NetworkMetrics) Descriptor() ([]byte, []int) {
	return file_proto_messaging_proto_rawDescGZIP(), []int{7}
}

func (x *NetworkMetrics) GetBytesSent() int64 {
	if x != nil {
		return x.BytesSent
	}
	return 0
}

func (x *NetworkMetrics) GetBytesReceived() int64 {
	if x != nil {
		return x.BytesReceived
	}
	return 0
}

func (x *NetworkMetrics) GetPacketsSent() int64 {
	if x != nil {
		return x.PacketsSent
	}
	return 0
}

func (x *NetworkMetrics) GetPacketsReceived() int64 {
	if x != nil {
		return x.PacketsReceived
	}
	return 0
}

var File_proto_messaging_proto protoreflect.FileDescriptor

const file_proto_messaging_proto_rawDesc = "" +
	"\n" +
	"\x15proto/messaging.proto\x12\braastack\"\x87\x02\n" +
	"\aMessage\x12\x12\n" +
	"\x04type\x18\x01 \x01(\tR\x04type\x12-\n" +
	"\acommand\x18\x02 \x01(\v2\x11.raastack.CommandH\x00R\acommand\x123\n" +
	"\theartbeat\x18\x03 \x01(\v2\x13.raastack.HeartbeatH\x00R\theartbeat\x127\n" +
	"\vclient_info\x18\x04 \x01(\v2\x14.raastack.ClientInfoH\x00R\n" +
	"clientInfo\x12@\n" +
	"\x0ecommand_result\x18\x05 \x01(\v2\x17.raastack.CommandResultH\x00R\rcommandResultB\t\n" +
	"\apayload\"\xa3\x01\n" +
	"\aCommand\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\acommand\x18\x02 \x01(\tR\acommand\x12\x12\n" +
	"\x04args\x18\x03 \x03(\tR\x04args\x12\x1f\n" +
	"\vrequire_pty\x18\x04 \x01(\bR\n" +
	"requirePty\x12\x1b\n" +
	"\tsender_id\x18\x05 \x01(\tR\bsenderId\x12\x1c\n" +
	"\ttimestamp\x18\x06 \x01(\x03R\ttimestamp\"\xed\x01\n" +
	"\rCommandResult\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"command_id\x18\x02 \x01(\tR\tcommandId\x12\x1b\n" +
	"\tclient_id\x18\x03 \x01(\tR\bclientId\x12\x18\n" +
	"\asuccess\x18\x04 \x01(\bR\asuccess\x12\x16\n" +
	"\x06output\x18\x05 \x01(\tR\x06output\x12\x14\n" +
	"\x05error\x18\x06 \x01(\tR\x05error\x12*\n" +
	"\x11execution_time_ms\x18\a \x01(\x03R\x0fexecutionTimeMs\x12\x1c\n" +
	"\ttimestamp\x18\b \x01(\x03R\ttimestamp\"\xc6\x02\n" +
	"\n" +
	"ClientInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\bhostname\x18\x02 \x01(\tR\bhostname\x12\x0e\n" +
	"\x02os\x18\x03 \x01(\tR\x02os\x12\x1d\n" +
	"\n" +
	"os_version\x18\x04 \x01(\tR\tosVersion\x12\"\n" +
	"\farchitecture\x18\x05 \x01(\tR\farchitecture\x12\x1d\n" +
	"\n" +
	"ip_address\x18\x06 \x01(\tR\tipAddress\x12!\n" +
	"\fconnected_at\x18\a \x01(\x03R\vconnectedAt\x12 \n" +
	"\flast_seen_at\x18\b \x01(\x03R\n" +
	"lastSeenAt\x12\"\n" +
	"\fcapabilities\x18\t \x03(\tR\fcapabilities\x121\n" +
	"\ametrics\x18\n" +
	" \x01(\v2\x17.raastack.SystemMetricsR\ametrics\"y\n" +
	"\tHeartbeat\x12\x1b\n" +
	"\tclient_id\x18\x01 \x01(\tR\bclientId\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\x03R\ttimestamp\x121\n" +
	"\ametrics\x18\x03 \x01(\v2\x17.raastack.SystemMetricsR\ametrics\"\xdb\x03\n" +
	"\rSystemMetrics\x12*\n" +
	"\x11cpu_usage_percent\x18\x01 \x01(\x01R\x0fcpuUsagePercent\x120\n" +
	"\x14memory_usage_percent\x18\x02 \x01(\x01R\x12memoryUsagePercent\x12,\n" +
	"\x12total_memory_bytes\x18\x03 \x01(\x03R\x10totalMemoryBytes\x12*\n" +
	"\x11used_memory_bytes\x18\x04 \x01(\x03R\x0fusedMemoryBytes\x12*\n" +
	"\x11free_memory_bytes\x18\x05 \x01(\x03R\x0ffreeMemoryBytes\x12\x1e\n" +
	"\vload_avg_1m\x18\x06 \x01(\x01R\tloadAvg1m\x12\x1e\n" +
	"\vload_avg_5m\x18\a \x01(\x01R\tloadAvg5m\x12 \n" +
	"\fload_avg_15m\x18\b \x01(\x01R\n" +
	"loadAvg15m\x12%\n" +
	"\x0euptime_seconds\x18\t \x01(\x03R\ruptimeSeconds\x12)\n" +
	"\x04disk\x18\n" +
	" \x01(\v2\x15.raastack.DiskMetricsR\x04disk\x122\n" +
	"\anetwork\x18\v \x01(\v2\x18.raastack.NetworkMetricsR\anetwork\"\xb5\x01\n" +
	"\vDiskMetrics\x12(\n" +
	"\x10total_disk_bytes\x18\x01 \x01(\x03R\x0etotalDiskBytes\x12&\n" +
	"\x0fused_disk_bytes\x18\x02 \x01(\x03R\rusedDiskBytes\x12&\n" +
	"\x0ffree_disk_bytes\x18\x03 \x01(\x03R\rfreeDiskBytes\x12,\n" +
	"\x12disk_usage_percent\x18\x04 \x01(\x01R\x10diskUsagePercent\"\xa4\x01\n" +
	"\x0eNetworkMetrics\x12\x1d\n" +
	"\n" +
	"bytes_sent\x18\x01 \x01(\x03R\tbytesSent\x12%\n" +
	"\x0ebytes_received\x18\x02 \x01(\x03R\rbytesReceived\x12!\n" +
	"\fpackets_sent\x18\x03 \x01(\x03R\vpacketsSent\x12)\n" +
	"\x10packets_received\x18\x04 \x01(\x03R\x0fpacketsReceivedB Z\x1egithub.com/raasstack/pkg/protob\x06proto3"

var (
	file_proto_messaging_proto_rawDescOnce sync.Once
	file_proto_messaging_proto_rawDescData []byte
)

func file_proto_messaging_proto_rawDescGZIP() []byte {
	file_proto_messaging_proto_rawDescOnce.Do(func() {
		file_proto_messaging_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_messaging_proto_rawDesc), len(file_proto_messaging_proto_rawDesc)))
	})
	return file_proto_messaging_proto_rawDescData
}

var file_proto_messaging_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_messaging_proto_goTypes = []any{
	(*Message)(nil),        // 0: raastack.Message
	(*Command)(nil),        // 1: raastack.Command
	(*CommandResult)(nil),  // 2: raastack.CommandResult
	(*ClientInfo)(nil),     // 3: raastack.ClientInfo
	(*Heartbeat)(nil),      // 4: raastack.Heartbeat
	(*SystemMetrics)(nil),  // 5: raastack.SystemMetrics
	(*DiskMetrics)(nil),    // 6: raastack.DiskMetrics
	(*NetworkMetrics)(nil), // 7: raastack.NetworkMetrics
}
var file_proto_messaging_proto_depIdxs = []int32{
	1, // 0: raastack.Message.command:type_name -> raastack.Command
	4, // 1: raastack.Message.heartbeat:type_name -> raastack.Heartbeat
	3, // 2: raastack.Message.client_info:type_name -> raastack.ClientInfo
	2, // 3: raastack.Message.command_result:type_name -> raastack.CommandResult
	5, // 4: raastack.ClientInfo.metrics:type_name -> raastack.SystemMetrics
	5, // 5: raastack.Heartbeat.metrics:type_name -> raastack.SystemMetrics
	6, // 6: raastack.SystemMetrics.disk:type_name -> raastack.DiskMetrics
	7, // 7: raastack.SystemMetrics.network:type_name -> raastack.NetworkMetrics
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_messaging_proto_init() }
func file_proto_messaging_proto_init() {
	if File_proto_messaging_proto != nil {
		return
	}
	file_proto_messaging_proto_msgTypes[0].OneofWrappers = []any{
		(*Message_Command)(nil),
		(*Message_Heartbeat)(nil),
		(*Message_ClientInfo)(nil),
		(*Message_CommandResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_messaging_proto_rawDesc), len(file_proto_messaging_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_messaging_proto_goTypes,
		DependencyIndexes: file_proto_messaging_proto_depIdxs,
		MessageInfos:      file_proto_messaging_proto_msgTypes,
	}.Build()
	File_proto_messaging_proto = out.File
	file_proto_messaging_proto_goTypes = nil
	file_proto_messaging_proto_depIdxs = nil
}
