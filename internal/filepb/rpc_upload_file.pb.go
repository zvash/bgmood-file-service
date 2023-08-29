// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc_upload_file.proto

package filepb

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

type FileInfo_FileType int32

const (
	FileInfo_AVATAR_IMAGE     FileInfo_FileType = 0
	FileInfo_BACKGROUND_IMAGE FileInfo_FileType = 1
	FileInfo_ENCRYPTED_FILE   FileInfo_FileType = 2
)

// Enum value maps for FileInfo_FileType.
var (
	FileInfo_FileType_name = map[int32]string{
		0: "AVATAR_IMAGE",
		1: "BACKGROUND_IMAGE",
		2: "ENCRYPTED_FILE",
	}
	FileInfo_FileType_value = map[string]int32{
		"AVATAR_IMAGE":     0,
		"BACKGROUND_IMAGE": 1,
		"ENCRYPTED_FILE":   2,
	}
)

func (x FileInfo_FileType) Enum() *FileInfo_FileType {
	p := new(FileInfo_FileType)
	*p = x
	return p
}

func (x FileInfo_FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileInfo_FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_upload_file_proto_enumTypes[0].Descriptor()
}

func (FileInfo_FileType) Type() protoreflect.EnumType {
	return &file_rpc_upload_file_proto_enumTypes[0]
}

func (x FileInfo_FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileInfo_FileType.Descriptor instead.
func (FileInfo_FileType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_upload_file_proto_rawDescGZIP(), []int{0, 0}
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileType  FileInfo_FileType `protobuf:"varint,1,opt,name=file_type,json=fileType,proto3,enum=filepb.FileInfo_FileType" json:"file_type,omitempty"`
	Extension *string           `protobuf:"bytes,2,opt,name=extension,proto3,oneof" json:"extension,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_upload_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_upload_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_rpc_upload_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetFileType() FileInfo_FileType {
	if x != nil {
		return x.FileType
	}
	return FileInfo_AVATAR_IMAGE
}

func (x *FileInfo) GetExtension() string {
	if x != nil && x.Extension != nil {
		return *x.Extension
	}
	return ""
}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UploadFileRequest_Info
	//	*UploadFileRequest_ChunkData
	Data isUploadFileRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_upload_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_upload_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_rpc_upload_file_proto_rawDescGZIP(), []int{1}
}

func (m *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadFileRequest) GetInfo() *FileInfo {
	if x, ok := x.GetData().(*UploadFileRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadFileRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UploadFileRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_Info struct {
	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadFileRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadFileRequest_Info) isUploadFileRequest_Data() {}

func (*UploadFileRequest_ChunkData) isUploadFileRequest_Data() {}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_upload_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_upload_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_rpc_upload_file_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadFileResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UploadFileResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_rpc_upload_file_proto protoreflect.FileDescriptor

var file_rpc_upload_file_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x22,
	0xbb, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x22, 0x46, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x49, 0x4d,
	0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x41, 0x43, 0x4b, 0x47, 0x52, 0x4f,
	0x55, 0x4e, 0x44, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x45,
	0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x02, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x64, 0x0a,
	0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x7a, 0x76, 0x61, 0x73, 0x68, 0x2f, 0x62, 0x67, 0x6d, 0x6f, 0x6f, 0x64, 0x2d, 0x66,
	0x69, 0x6c, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_upload_file_proto_rawDescOnce sync.Once
	file_rpc_upload_file_proto_rawDescData = file_rpc_upload_file_proto_rawDesc
)

func file_rpc_upload_file_proto_rawDescGZIP() []byte {
	file_rpc_upload_file_proto_rawDescOnce.Do(func() {
		file_rpc_upload_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_upload_file_proto_rawDescData)
	})
	return file_rpc_upload_file_proto_rawDescData
}

var file_rpc_upload_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_upload_file_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_upload_file_proto_goTypes = []interface{}{
	(FileInfo_FileType)(0),     // 0: filepb.FileInfo.FileType
	(*FileInfo)(nil),           // 1: filepb.FileInfo
	(*UploadFileRequest)(nil),  // 2: filepb.UploadFileRequest
	(*UploadFileResponse)(nil), // 3: filepb.UploadFileResponse
}
var file_rpc_upload_file_proto_depIdxs = []int32{
	0, // 0: filepb.FileInfo.file_type:type_name -> filepb.FileInfo.FileType
	1, // 1: filepb.UploadFileRequest.info:type_name -> filepb.FileInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_upload_file_proto_init() }
func file_rpc_upload_file_proto_init() {
	if File_rpc_upload_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_upload_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_rpc_upload_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_rpc_upload_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
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
	file_rpc_upload_file_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_rpc_upload_file_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UploadFileRequest_Info)(nil),
		(*UploadFileRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_upload_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_upload_file_proto_goTypes,
		DependencyIndexes: file_rpc_upload_file_proto_depIdxs,
		EnumInfos:         file_rpc_upload_file_proto_enumTypes,
		MessageInfos:      file_rpc_upload_file_proto_msgTypes,
	}.Build()
	File_rpc_upload_file_proto = out.File
	file_rpc_upload_file_proto_rawDesc = nil
	file_rpc_upload_file_proto_goTypes = nil
	file_rpc_upload_file_proto_depIdxs = nil
}