// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: pkg/driver/external/driver.proto

package external

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type InfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InfoJson      []byte                 `protobuf:"bytes,1,opt,name=info_json,json=infoJson,proto3" json:"info_json,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{0}
}

func (x *InfoResponse) GetInfoJson() []byte {
	if x != nil {
		return x.InfoJson
	}
	return nil
}

type StartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{1}
}

func (x *StartResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *StartResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SetConfigRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	InstanceConfigJson []byte                 `protobuf:"bytes,1,opt,name=instance_config_json,json=instanceConfigJson,proto3" json:"instance_config_json,omitempty"`
	SshLocalPort       int64                  `protobuf:"varint,3,opt,name=ssh_local_port,json=sshLocalPort,proto3" json:"ssh_local_port,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *SetConfigRequest) Reset() {
	*x = SetConfigRequest{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfigRequest) ProtoMessage() {}

func (x *SetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConfigRequest.ProtoReflect.Descriptor instead.
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{2}
}

func (x *SetConfigRequest) GetInstanceConfigJson() []byte {
	if x != nil {
		return x.InstanceConfigJson
	}
	return nil
}

func (x *SetConfigRequest) GetSshLocalPort() int64 {
	if x != nil {
		return x.SshLocalPort
	}
	return 0
}

type ChangeDisplayPasswordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Password      string                 `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangeDisplayPasswordRequest) Reset() {
	*x = ChangeDisplayPasswordRequest{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeDisplayPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeDisplayPasswordRequest) ProtoMessage() {}

func (x *ChangeDisplayPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeDisplayPasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangeDisplayPasswordRequest) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeDisplayPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetDisplayConnectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Connection    string                 `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDisplayConnectionResponse) Reset() {
	*x = GetDisplayConnectionResponse{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDisplayConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDisplayConnectionResponse) ProtoMessage() {}

func (x *GetDisplayConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDisplayConnectionResponse.ProtoReflect.Descriptor instead.
func (*GetDisplayConnectionResponse) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{4}
}

func (x *GetDisplayConnectionResponse) GetConnection() string {
	if x != nil {
		return x.Connection
	}
	return ""
}

type CreateSnapshotRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSnapshotRequest) Reset() {
	*x = CreateSnapshotRequest{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnapshotRequest) ProtoMessage() {}

func (x *CreateSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnapshotRequest.ProtoReflect.Descriptor instead.
func (*CreateSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSnapshotRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type ApplySnapshotRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApplySnapshotRequest) Reset() {
	*x = ApplySnapshotRequest{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplySnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplySnapshotRequest) ProtoMessage() {}

func (x *ApplySnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplySnapshotRequest.ProtoReflect.Descriptor instead.
func (*ApplySnapshotRequest) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{6}
}

func (x *ApplySnapshotRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type DeleteSnapshotRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSnapshotRequest) Reset() {
	*x = DeleteSnapshotRequest{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnapshotRequest) ProtoMessage() {}

func (x *DeleteSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnapshotRequest.ProtoReflect.Descriptor instead.
func (*DeleteSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSnapshotRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type ListSnapshotsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Snapshots     string                 `protobuf:"bytes,1,opt,name=snapshots,proto3" json:"snapshots,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSnapshotsResponse) Reset() {
	*x = ListSnapshotsResponse{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSnapshotsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnapshotsResponse) ProtoMessage() {}

func (x *ListSnapshotsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnapshotsResponse.ProtoReflect.Descriptor instead.
func (*ListSnapshotsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{8}
}

func (x *ListSnapshotsResponse) GetSnapshots() string {
	if x != nil {
		return x.Snapshots
	}
	return ""
}

type ForwardGuestAgentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShouldForward bool                   `protobuf:"varint,1,opt,name=should_forward,json=shouldForward,proto3" json:"should_forward,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ForwardGuestAgentResponse) Reset() {
	*x = ForwardGuestAgentResponse{}
	mi := &file_pkg_driver_external_driver_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForwardGuestAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardGuestAgentResponse) ProtoMessage() {}

func (x *ForwardGuestAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_external_driver_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardGuestAgentResponse.ProtoReflect.Descriptor instead.
func (*ForwardGuestAgentResponse) Descriptor() ([]byte, []int) {
	return file_pkg_driver_external_driver_proto_rawDescGZIP(), []int{9}
}

func (x *ForwardGuestAgentResponse) GetShouldForward() bool {
	if x != nil {
		return x.ShouldForward
	}
	return false
}

var File_pkg_driver_external_driver_proto protoreflect.FileDescriptor

const file_pkg_driver_external_driver_proto_rawDesc = "" +
	"\n" +
	" pkg/driver/external/driver.proto\x1a\x1bgoogle/protobuf/empty.proto\"+\n" +
	"\fInfoResponse\x12\x1b\n" +
	"\tinfo_json\x18\x01 \x01(\fR\binfoJson\"?\n" +
	"\rStartResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\"j\n" +
	"\x10SetConfigRequest\x120\n" +
	"\x14instance_config_json\x18\x01 \x01(\fR\x12instanceConfigJson\x12$\n" +
	"\x0essh_local_port\x18\x03 \x01(\x03R\fsshLocalPort\":\n" +
	"\x1cChangeDisplayPasswordRequest\x12\x1a\n" +
	"\bpassword\x18\x01 \x01(\tR\bpassword\">\n" +
	"\x1cGetDisplayConnectionResponse\x12\x1e\n" +
	"\n" +
	"connection\x18\x01 \x01(\tR\n" +
	"connection\")\n" +
	"\x15CreateSnapshotRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\"(\n" +
	"\x14ApplySnapshotRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\")\n" +
	"\x15DeleteSnapshotRequest\x12\x10\n" +
	"\x03tag\x18\x01 \x01(\tR\x03tag\"5\n" +
	"\x15ListSnapshotsResponse\x12\x1c\n" +
	"\tsnapshots\x18\x01 \x01(\tR\tsnapshots\"B\n" +
	"\x19ForwardGuestAgentResponse\x12%\n" +
	"\x0eshould_forward\x18\x01 \x01(\bR\rshouldForward2\xf8\b\n" +
	"\x06Driver\x12:\n" +
	"\bValidate\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x12<\n" +
	"\n" +
	"Initialize\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x12<\n" +
	"\n" +
	"CreateDisk\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x121\n" +
	"\x05Start\x12\x16.google.protobuf.Empty\x1a\x0e.StartResponse0\x01\x126\n" +
	"\x04Stop\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x128\n" +
	"\x06RunGUI\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x12N\n" +
	"\x15ChangeDisplayPassword\x12\x1d.ChangeDisplayPasswordRequest\x1a\x16.google.protobuf.Empty\x12M\n" +
	"\x14GetDisplayConnection\x12\x16.google.protobuf.Empty\x1a\x1d.GetDisplayConnectionResponse\x12@\n" +
	"\x0eCreateSnapshot\x12\x16.CreateSnapshotRequest\x1a\x16.google.protobuf.Empty\x12>\n" +
	"\rApplySnapshot\x12\x15.ApplySnapshotRequest\x1a\x16.google.protobuf.Empty\x12@\n" +
	"\x0eDeleteSnapshot\x12\x16.DeleteSnapshotRequest\x1a\x16.google.protobuf.Empty\x12?\n" +
	"\rListSnapshots\x12\x16.google.protobuf.Empty\x1a\x16.ListSnapshotsResponse\x12:\n" +
	"\bRegister\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x12<\n" +
	"\n" +
	"Unregister\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x12G\n" +
	"\x11ForwardGuestAgent\x12\x16.google.protobuf.Empty\x1a\x1a.ForwardGuestAgentResponse\x12@\n" +
	"\x0eGuestAgentConn\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x126\n" +
	"\tSetConfig\x12\x11.SetConfigRequest\x1a\x16.google.protobuf.Empty\x120\n" +
	"\aGetInfo\x12\x16.google.protobuf.Empty\x1a\r.InfoResponseB-Z+github.com/lima-vm/lima/pkg/driver/externalb\x06proto3"

var (
	file_pkg_driver_external_driver_proto_rawDescOnce sync.Once
	file_pkg_driver_external_driver_proto_rawDescData []byte
)

func file_pkg_driver_external_driver_proto_rawDescGZIP() []byte {
	file_pkg_driver_external_driver_proto_rawDescOnce.Do(func() {
		file_pkg_driver_external_driver_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_driver_external_driver_proto_rawDesc), len(file_pkg_driver_external_driver_proto_rawDesc)))
	})
	return file_pkg_driver_external_driver_proto_rawDescData
}

var file_pkg_driver_external_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_driver_external_driver_proto_goTypes = []any{
	(*InfoResponse)(nil),                 // 0: InfoResponse
	(*StartResponse)(nil),                // 1: StartResponse
	(*SetConfigRequest)(nil),             // 2: SetConfigRequest
	(*ChangeDisplayPasswordRequest)(nil), // 3: ChangeDisplayPasswordRequest
	(*GetDisplayConnectionResponse)(nil), // 4: GetDisplayConnectionResponse
	(*CreateSnapshotRequest)(nil),        // 5: CreateSnapshotRequest
	(*ApplySnapshotRequest)(nil),         // 6: ApplySnapshotRequest
	(*DeleteSnapshotRequest)(nil),        // 7: DeleteSnapshotRequest
	(*ListSnapshotsResponse)(nil),        // 8: ListSnapshotsResponse
	(*ForwardGuestAgentResponse)(nil),    // 9: ForwardGuestAgentResponse
	(*emptypb.Empty)(nil),                // 10: google.protobuf.Empty
}
var file_pkg_driver_external_driver_proto_depIdxs = []int32{
	10, // 0: Driver.Validate:input_type -> google.protobuf.Empty
	10, // 1: Driver.Initialize:input_type -> google.protobuf.Empty
	10, // 2: Driver.CreateDisk:input_type -> google.protobuf.Empty
	10, // 3: Driver.Start:input_type -> google.protobuf.Empty
	10, // 4: Driver.Stop:input_type -> google.protobuf.Empty
	10, // 5: Driver.RunGUI:input_type -> google.protobuf.Empty
	3,  // 6: Driver.ChangeDisplayPassword:input_type -> ChangeDisplayPasswordRequest
	10, // 7: Driver.GetDisplayConnection:input_type -> google.protobuf.Empty
	5,  // 8: Driver.CreateSnapshot:input_type -> CreateSnapshotRequest
	6,  // 9: Driver.ApplySnapshot:input_type -> ApplySnapshotRequest
	7,  // 10: Driver.DeleteSnapshot:input_type -> DeleteSnapshotRequest
	10, // 11: Driver.ListSnapshots:input_type -> google.protobuf.Empty
	10, // 12: Driver.Register:input_type -> google.protobuf.Empty
	10, // 13: Driver.Unregister:input_type -> google.protobuf.Empty
	10, // 14: Driver.ForwardGuestAgent:input_type -> google.protobuf.Empty
	10, // 15: Driver.GuestAgentConn:input_type -> google.protobuf.Empty
	2,  // 16: Driver.SetConfig:input_type -> SetConfigRequest
	10, // 17: Driver.GetInfo:input_type -> google.protobuf.Empty
	10, // 18: Driver.Validate:output_type -> google.protobuf.Empty
	10, // 19: Driver.Initialize:output_type -> google.protobuf.Empty
	10, // 20: Driver.CreateDisk:output_type -> google.protobuf.Empty
	1,  // 21: Driver.Start:output_type -> StartResponse
	10, // 22: Driver.Stop:output_type -> google.protobuf.Empty
	10, // 23: Driver.RunGUI:output_type -> google.protobuf.Empty
	10, // 24: Driver.ChangeDisplayPassword:output_type -> google.protobuf.Empty
	4,  // 25: Driver.GetDisplayConnection:output_type -> GetDisplayConnectionResponse
	10, // 26: Driver.CreateSnapshot:output_type -> google.protobuf.Empty
	10, // 27: Driver.ApplySnapshot:output_type -> google.protobuf.Empty
	10, // 28: Driver.DeleteSnapshot:output_type -> google.protobuf.Empty
	8,  // 29: Driver.ListSnapshots:output_type -> ListSnapshotsResponse
	10, // 30: Driver.Register:output_type -> google.protobuf.Empty
	10, // 31: Driver.Unregister:output_type -> google.protobuf.Empty
	9,  // 32: Driver.ForwardGuestAgent:output_type -> ForwardGuestAgentResponse
	10, // 33: Driver.GuestAgentConn:output_type -> google.protobuf.Empty
	10, // 34: Driver.SetConfig:output_type -> google.protobuf.Empty
	0,  // 35: Driver.GetInfo:output_type -> InfoResponse
	18, // [18:36] is the sub-list for method output_type
	0,  // [0:18] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_driver_external_driver_proto_init() }
func file_pkg_driver_external_driver_proto_init() {
	if File_pkg_driver_external_driver_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_driver_external_driver_proto_rawDesc), len(file_pkg_driver_external_driver_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_driver_external_driver_proto_goTypes,
		DependencyIndexes: file_pkg_driver_external_driver_proto_depIdxs,
		MessageInfos:      file_pkg_driver_external_driver_proto_msgTypes,
	}.Build()
	File_pkg_driver_external_driver_proto = out.File
	file_pkg_driver_external_driver_proto_goTypes = nil
	file_pkg_driver_external_driver_proto_depIdxs = nil
}
