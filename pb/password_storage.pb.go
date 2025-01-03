// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: pb/password_storage.proto

package passwordstoragepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A single password entry in the password storage
type PasswordEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Group       string `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PasswordEntry) Reset() {
	*x = PasswordEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_password_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordEntry) ProtoMessage() {}

func (x *PasswordEntry) ProtoReflect() protoreflect.Message {
	mi := &file_pb_password_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordEntry.ProtoReflect.Descriptor instead.
func (*PasswordEntry) Descriptor() ([]byte, []int) {
	return file_pb_password_storage_proto_rawDescGZIP(), []int{0}
}

func (x *PasswordEntry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PasswordEntry) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PasswordEntry) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PasswordEntry) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *PasswordEntry) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// The data structure used to store the password entries
type PasswordEntryList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PasswordEntries []*PasswordEntry `protobuf:"bytes,1,rep,name=password_entries,json=passwordEntries,proto3" json:"password_entries,omitempty"`
	// Always include id to fix decryption problem when message is empty
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PasswordEntryList) Reset() {
	*x = PasswordEntryList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_password_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordEntryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordEntryList) ProtoMessage() {}

func (x *PasswordEntryList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_password_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordEntryList.ProtoReflect.Descriptor instead.
func (*PasswordEntryList) Descriptor() ([]byte, []int) {
	return file_pb_password_storage_proto_rawDescGZIP(), []int{1}
}

func (x *PasswordEntryList) GetPasswordEntries() []*PasswordEntry {
	if x != nil {
		return x.PasswordEntries
	}
	return nil
}

func (x *PasswordEntryList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pb_password_storage_proto protoreflect.FileDescriptor

var file_pb_password_storage_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x22, 0x8f,
	0x01, 0x0a, 0x0d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x70, 0x0a, 0x11, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_password_storage_proto_rawDescOnce sync.Once
	file_pb_password_storage_proto_rawDescData = file_pb_password_storage_proto_rawDesc
)

func file_pb_password_storage_proto_rawDescGZIP() []byte {
	file_pb_password_storage_proto_rawDescOnce.Do(func() {
		file_pb_password_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_password_storage_proto_rawDescData)
	})
	return file_pb_password_storage_proto_rawDescData
}

var file_pb_password_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_password_storage_proto_goTypes = []interface{}{
	(*PasswordEntry)(nil),     // 0: passwordstoragepb.PasswordEntry
	(*PasswordEntryList)(nil), // 1: passwordstoragepb.PasswordEntryList
}
var file_pb_password_storage_proto_depIdxs = []int32{
	0, // 0: passwordstoragepb.PasswordEntryList.password_entries:type_name -> passwordstoragepb.PasswordEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_password_storage_proto_init() }
func file_pb_password_storage_proto_init() {
	if File_pb_password_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_password_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_password_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordEntryList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_password_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_password_storage_proto_goTypes,
		DependencyIndexes: file_pb_password_storage_proto_depIdxs,
		MessageInfos:      file_pb_password_storage_proto_msgTypes,
	}.Build()
	File_pb_password_storage_proto = out.File
	file_pb_password_storage_proto_rawDesc = nil
	file_pb_password_storage_proto_goTypes = nil
	file_pb_password_storage_proto_depIdxs = nil
}
