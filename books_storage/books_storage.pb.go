// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/books_storage.proto

package books_storage

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

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorName string `protobuf:"bytes,1,opt,name=AuthorName,proto3" json:"AuthorName,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_proto_books_storage_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookName string `protobuf:"bytes,1,opt,name=BookName,proto3" json:"BookName,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_proto_books_storage_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

type Authors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorNames []string `protobuf:"bytes,1,rep,name=AuthorNames,proto3" json:"AuthorNames,omitempty"`
}

func (x *Authors) Reset() {
	*x = Authors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authors) ProtoMessage() {}

func (x *Authors) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authors.ProtoReflect.Descriptor instead.
func (*Authors) Descriptor() ([]byte, []int) {
	return file_proto_books_storage_proto_rawDescGZIP(), []int{2}
}

func (x *Authors) GetAuthorNames() []string {
	if x != nil {
		return x.AuthorNames
	}
	return nil
}

type Books struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookNames []string `protobuf:"bytes,1,rep,name=BookNames,proto3" json:"BookNames,omitempty"`
}

func (x *Books) Reset() {
	*x = Books{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Books) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Books) ProtoMessage() {}

func (x *Books) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Books.ProtoReflect.Descriptor instead.
func (*Books) Descriptor() ([]byte, []int) {
	return file_proto_books_storage_proto_rawDescGZIP(), []int{3}
}

func (x *Books) GetBookNames() []string {
	if x != nil {
		return x.BookNames
	}
	return nil
}

var File_proto_books_storage_proto protoreflect.FileDescriptor

var file_proto_books_storage_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x06, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x0a, 0x08,
	0x42, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x42, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x05, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x94, 0x01, 0x0a,
	0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_books_storage_proto_rawDescOnce sync.Once
	file_proto_books_storage_proto_rawDescData = file_proto_books_storage_proto_rawDesc
)

func file_proto_books_storage_proto_rawDescGZIP() []byte {
	file_proto_books_storage_proto_rawDescOnce.Do(func() {
		file_proto_books_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_books_storage_proto_rawDescData)
	})
	return file_proto_books_storage_proto_rawDescData
}

var file_proto_books_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_books_storage_proto_goTypes = []interface{}{
	(*Author)(nil),  // 0: books_storage.Author
	(*Book)(nil),    // 1: books_storage.Book
	(*Authors)(nil), // 2: books_storage.Authors
	(*Books)(nil),   // 3: books_storage.Books
}
var file_proto_books_storage_proto_depIdxs = []int32{
	0, // 0: books_storage.BooksStorage.GetBooksByAuthor:input_type -> books_storage.Author
	1, // 1: books_storage.BooksStorage.GetAuthorsByBook:input_type -> books_storage.Book
	3, // 2: books_storage.BooksStorage.GetBooksByAuthor:output_type -> books_storage.Books
	2, // 3: books_storage.BooksStorage.GetAuthorsByBook:output_type -> books_storage.Authors
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_books_storage_proto_init() }
func file_proto_books_storage_proto_init() {
	if File_proto_books_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_books_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_proto_books_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_proto_books_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authors); i {
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
		file_proto_books_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Books); i {
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
			RawDescriptor: file_proto_books_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_books_storage_proto_goTypes,
		DependencyIndexes: file_proto_books_storage_proto_depIdxs,
		MessageInfos:      file_proto_books_storage_proto_msgTypes,
	}.Build()
	File_proto_books_storage_proto = out.File
	file_proto_books_storage_proto_rawDesc = nil
	file_proto_books_storage_proto_goTypes = nil
	file_proto_books_storage_proto_depIdxs = nil
}
