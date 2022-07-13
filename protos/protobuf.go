package protos

/**
 * @title protobuf
 * @author xiongshao
 * @date 2022-06-22 15:04:04
 */

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
type ClassName int32

const (
	ClassName_class1 ClassName = 0
	ClassName_class2 ClassName = 1
	ClassName_class3 ClassName = 2
)

var ClassName_name = map[int32]string{
	0: "class1",
	1: "class2",
	2: "class3",
}
var ClassName_value = map[string]int32{
	"class1": 0,
	"class2": 1,
	"class3": 2,
}

func (x ClassName) String() string {
	return proto.EnumName(ClassName_name, int32(x))
}
func (ClassName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

type Student struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32     `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Address              string    `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Cn                   ClassName `protobuf:"varint,4,opt,name=cn,proto3,enum=main.ClassName" json:"cn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}
func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Student) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Student) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Student) GetCn() ClassName {
	if m != nil {
		return m.Cn
	}
	return ClassName_class1
}

type Students struct {
	Person               []*Student `protobuf:"bytes,1,rep,name=person,proto3" json:"person,omitempty"`
	School               string     `protobuf:"bytes,2,opt,name=school,proto3" json:"school,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Students) Reset()         { *m = Students{} }
func (m *Students) String() string { return proto.CompactTextString(m) }
func (*Students) ProtoMessage()    {}
func (*Students) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *Students) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Students.Unmarshal(m, b)
}
func (m *Students) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Students.Marshal(b, m, deterministic)
}
func (m *Students) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Students.Merge(m, src)
}
func (m *Students) XXX_Size() int {
	return xxx_messageInfo_Students.Size(m)
}
func (m *Students) XXX_DiscardUnknown() {
	xxx_messageInfo_Students.DiscardUnknown(m)
}

var xxx_messageInfo_Students proto.InternalMessageInfo

func (m *Students) GetPerson() []*Student {
	if m != nil {
		return m.Person
	}
	return nil
}

func (m *Students) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func init() {
	proto.RegisterEnum("main.ClassName", ClassName_name, ClassName_value)
	proto.RegisterType((*Student)(nil), "main.Student")
	proto.RegisterType((*Students)(nil), "main.Students")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xcd, 0x4e, 0xc5, 0x20,
	0x10, 0x85, 0x85, 0x56, 0xae, 0x1d, 0xa3, 0x92, 0x59, 0x18, 0x76, 0x92, 0x9b, 0x98, 0x10, 0x17,
	0x35, 0xf6, 0x3e, 0x82, 0x2b, 0x37, 0x2e, 0xf0, 0x09, 0x90, 0x12, 0x7f, 0xd2, 0x42, 0xd3, 0xc1,
	0xf7, 0x37, 0x45, 0xec, 0xee, 0xfb, 0x72, 0xe0, 0xcc, 0x0c, 0x40, 0x0e, 0x94, 0xfb, 0x65, 0x4d,
	0x39, 0x61, 0x3b, 0xbb, 0xaf, 0x78, 0xfc, 0x86, 0xc3, 0x5b, 0xfe, 0x19, 0x43, 0xcc, 0x88, 0xd0,
	0x46, 0x37, 0x07, 0xc5, 0x34, 0x33, 0x9d, 0x2d, 0x8c, 0x12, 0x1a, 0xf7, 0x11, 0x14, 0xd7, 0xcc,
	0x9c, 0xdb, 0x0d, 0x51, 0xc1, 0xc1, 0x8d, 0xe3, 0x1a, 0x88, 0x54, 0x53, 0x1e, 0xfe, 0x2b, 0xde,
	0x01, 0xf7, 0x51, 0xb5, 0x9a, 0x99, 0xeb, 0xe1, 0xa6, 0xdf, 0xda, 0xfb, 0xe7, 0xc9, 0x11, 0xbd,
	0xba, 0x39, 0x58, 0xee, 0xe3, 0xf1, 0x05, 0x2e, 0xea, 0x2c, 0xc2, 0x7b, 0x10, 0x4b, 0x58, 0x29,
	0x45, 0xc5, 0x74, 0x63, 0x2e, 0x87, 0xab, 0xbf, 0x0f, 0x35, 0xb7, 0x35, 0xc4, 0x5b, 0x10, 0xe4,
	0x3f, 0x53, 0x9a, 0xca, 0x0a, 0x9d, 0xad, 0xf6, 0xf0, 0x08, 0xdd, 0xde, 0x8d, 0x00, 0xc2, 0x6f,
	0xf2, 0x24, 0xcf, 0x76, 0x1e, 0x24, 0xdb, 0xf9, 0x24, 0xf9, 0xbb, 0x28, 0x47, 0x9f, 0x7e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x67, 0x97, 0x87, 0x75, 0x02, 0x01, 0x00, 0x00,
}

// protobuf 测试
func protobufTest() {
	s1 := &Student{} // 第一个测试
	s1.Name = "jz01"
	s1.Age = 23
	s1.Address = "cq"
	s1.Cn = ClassName_class2 //枚举类型赋值
	ss := &Students{}
	ss.Person = append(ss.Person, s1) //将第一个学生信息添加到Students对应的切片中
	s2 := &Student{}                  //第二个学生信息
	s2.Name = "jz02"
	s2.Age = 25
	s2.Address = "cd"
	s2.Cn = ClassName_class3
	ss.Person = append(ss.Person, s2) //将第二个学生信息添加到Students对应的切片中
	ss.School = "cqu"
	fmt.Println("Students信息为：", ss)
	// Marshal takes a protocol buffer message
	// and encodes it into the wire format, returning the data.
	buffer, _ := proto.Marshal(ss)
	fmt.Println("序列化之后的信息为：", buffer)
	// Use UnmarshalMerge to preserve and append to existing data.
	data := &Students{}
	proto.Unmarshal(buffer, data)
	fmt.Println("反序列化之后的信息为：", data)
}
