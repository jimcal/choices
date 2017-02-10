// Code generated by protoc-gen-gogo.
// source: elwin.proto
// DO NOT EDIT!

/*
	Package elwin is a generated protocol buffer package.

	It is generated from these files:
		elwin.proto

	It has these top-level messages:
		Identifier
		Experiments
		Experiment
		Param
*/
package elwin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Identifier struct {
	TeamID string `protobuf:"bytes,1,opt,name=teamID,proto3" json:"teamID,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (m *Identifier) Reset()                    { *m = Identifier{} }
func (*Identifier) ProtoMessage()               {}
func (*Identifier) Descriptor() ([]byte, []int) { return fileDescriptorElwin, []int{0} }

func (m *Identifier) GetTeamID() string {
	if m != nil {
		return m.TeamID
	}
	return ""
}

func (m *Identifier) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type Experiments struct {
	Experiments map[string]*Experiment `protobuf:"bytes,1,rep,name=experiments" json:"experiments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Experiments) Reset()                    { *m = Experiments{} }
func (*Experiments) ProtoMessage()               {}
func (*Experiments) Descriptor() ([]byte, []int) { return fileDescriptorElwin, []int{1} }

func (m *Experiments) GetExperiments() map[string]*Experiment {
	if m != nil {
		return m.Experiments
	}
	return nil
}

type Experiment struct {
	Namespace string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Params    []*Param `protobuf:"bytes,2,rep,name=params" json:"params,omitempty"`
}

func (m *Experiment) Reset()                    { *m = Experiment{} }
func (*Experiment) ProtoMessage()               {}
func (*Experiment) Descriptor() ([]byte, []int) { return fileDescriptorElwin, []int{2} }

func (m *Experiment) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Experiment) GetParams() []*Param {
	if m != nil {
		return m.Params
	}
	return nil
}

type Param struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Param) Reset()                    { *m = Param{} }
func (*Param) ProtoMessage()               {}
func (*Param) Descriptor() ([]byte, []int) { return fileDescriptorElwin, []int{3} }

func (m *Param) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Param) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Identifier)(nil), "elwin.Identifier")
	proto.RegisterType((*Experiments)(nil), "elwin.Experiments")
	proto.RegisterType((*Experiment)(nil), "elwin.Experiment")
	proto.RegisterType((*Param)(nil), "elwin.Param")
}
func (this *Identifier) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Identifier)
	if !ok {
		that2, ok := that.(Identifier)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.TeamID != that1.TeamID {
		return false
	}
	if this.UserID != that1.UserID {
		return false
	}
	return true
}
func (this *Experiments) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Experiments)
	if !ok {
		that2, ok := that.(Experiments)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Experiments) != len(that1.Experiments) {
		return false
	}
	for i := range this.Experiments {
		if !this.Experiments[i].Equal(that1.Experiments[i]) {
			return false
		}
	}
	return true
}
func (this *Experiment) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Experiment)
	if !ok {
		that2, ok := that.(Experiment)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if len(this.Params) != len(that1.Params) {
		return false
	}
	for i := range this.Params {
		if !this.Params[i].Equal(that1.Params[i]) {
			return false
		}
	}
	return true
}
func (this *Param) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Param)
	if !ok {
		that2, ok := that.(Param)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *Identifier) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&elwin.Identifier{")
	s = append(s, "TeamID: "+fmt.Sprintf("%#v", this.TeamID)+",\n")
	s = append(s, "UserID: "+fmt.Sprintf("%#v", this.UserID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Experiments) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&elwin.Experiments{")
	keysForExperiments := make([]string, 0, len(this.Experiments))
	for k, _ := range this.Experiments {
		keysForExperiments = append(keysForExperiments, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForExperiments)
	mapStringForExperiments := "map[string]*Experiment{"
	for _, k := range keysForExperiments {
		mapStringForExperiments += fmt.Sprintf("%#v: %#v,", k, this.Experiments[k])
	}
	mapStringForExperiments += "}"
	if this.Experiments != nil {
		s = append(s, "Experiments: "+mapStringForExperiments+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Experiment) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&elwin.Experiment{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	if this.Params != nil {
		s = append(s, "Params: "+fmt.Sprintf("%#v", this.Params)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Param) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&elwin.Param{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringElwin(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringElwin(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Elwin service

type ElwinClient interface {
	GetNamespaces(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Experiments, error)
}

type elwinClient struct {
	cc *grpc.ClientConn
}

func NewElwinClient(cc *grpc.ClientConn) ElwinClient {
	return &elwinClient{cc}
}

func (c *elwinClient) GetNamespaces(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Experiments, error) {
	out := new(Experiments)
	err := grpc.Invoke(ctx, "/elwin.Elwin/GetNamespaces", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Elwin service

type ElwinServer interface {
	GetNamespaces(context.Context, *Identifier) (*Experiments, error)
}

func RegisterElwinServer(s *grpc.Server, srv ElwinServer) {
	s.RegisterService(&_Elwin_serviceDesc, srv)
}

func _Elwin_GetNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElwinServer).GetNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elwin.Elwin/GetNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElwinServer).GetNamespaces(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _Elwin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elwin.Elwin",
	HandlerType: (*ElwinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNamespaces",
			Handler:    _Elwin_GetNamespaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elwin.proto",
}

func (m *Identifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Identifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TeamID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintElwin(dAtA, i, uint64(len(m.TeamID)))
		i += copy(dAtA[i:], m.TeamID)
	}
	if len(m.UserID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintElwin(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	return i, nil
}

func (m *Experiments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Experiments) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Experiments) > 0 {
		for k, _ := range m.Experiments {
			dAtA[i] = 0xa
			i++
			v := m.Experiments[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovElwin(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovElwin(uint64(len(k))) + msgSize
			i = encodeVarintElwin(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintElwin(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintElwin(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	return i, nil
}

func (m *Experiment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Experiment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintElwin(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Params) > 0 {
		for _, msg := range m.Params {
			dAtA[i] = 0x12
			i++
			i = encodeVarintElwin(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Param) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Param) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintElwin(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintElwin(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func encodeFixed64Elwin(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Elwin(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintElwin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Identifier) Size() (n int) {
	var l int
	_ = l
	l = len(m.TeamID)
	if l > 0 {
		n += 1 + l + sovElwin(uint64(l))
	}
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovElwin(uint64(l))
	}
	return n
}

func (m *Experiments) Size() (n int) {
	var l int
	_ = l
	if len(m.Experiments) > 0 {
		for k, v := range m.Experiments {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovElwin(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovElwin(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovElwin(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Experiment) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovElwin(uint64(l))
	}
	if len(m.Params) > 0 {
		for _, e := range m.Params {
			l = e.Size()
			n += 1 + l + sovElwin(uint64(l))
		}
	}
	return n
}

func (m *Param) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovElwin(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovElwin(uint64(l))
	}
	return n
}

func sovElwin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozElwin(x uint64) (n int) {
	return sovElwin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Identifier) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Identifier{`,
		`TeamID:` + fmt.Sprintf("%v", this.TeamID) + `,`,
		`UserID:` + fmt.Sprintf("%v", this.UserID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Experiments) String() string {
	if this == nil {
		return "nil"
	}
	keysForExperiments := make([]string, 0, len(this.Experiments))
	for k, _ := range this.Experiments {
		keysForExperiments = append(keysForExperiments, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForExperiments)
	mapStringForExperiments := "map[string]*Experiment{"
	for _, k := range keysForExperiments {
		mapStringForExperiments += fmt.Sprintf("%v: %v,", k, this.Experiments[k])
	}
	mapStringForExperiments += "}"
	s := strings.Join([]string{`&Experiments{`,
		`Experiments:` + mapStringForExperiments + `,`,
		`}`,
	}, "")
	return s
}
func (this *Experiment) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Experiment{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Params:` + strings.Replace(fmt.Sprintf("%v", this.Params), "Param", "Param", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Param) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Param{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringElwin(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Identifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowElwin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Identifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Identifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthElwin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TeamID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthElwin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipElwin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthElwin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Experiments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowElwin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Experiments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Experiments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Experiments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthElwin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthElwin
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Experiments == nil {
				m.Experiments = make(map[string]*Experiment)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowElwin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowElwin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthElwin
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthElwin
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := &Experiment{}
				if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.Experiments[mapkey] = mapvalue
			} else {
				var mapvalue *Experiment
				m.Experiments[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipElwin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthElwin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Experiment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowElwin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Experiment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Experiment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthElwin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthElwin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Params = append(m.Params, &Param{})
			if err := m.Params[len(m.Params)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipElwin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthElwin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Param) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowElwin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Param: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Param: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthElwin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthElwin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipElwin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthElwin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipElwin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowElwin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowElwin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthElwin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowElwin
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipElwin(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthElwin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowElwin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("elwin.proto", fileDescriptorElwin) }

var fileDescriptorElwin = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x4a, 0xe3, 0x50,
	0x14, 0xc6, 0x73, 0xd3, 0x49, 0xa1, 0x27, 0x33, 0xd0, 0xb9, 0x0c, 0x25, 0x74, 0xca, 0xa5, 0x93,
	0x19, 0x98, 0x32, 0xcc, 0xb4, 0xb4, 0xb3, 0x91, 0xe2, 0x4a, 0x1a, 0xa4, 0x1b, 0xa9, 0xd9, 0xb9,
	0xbc, 0xd5, 0x63, 0x09, 0x36, 0x37, 0x21, 0xb9, 0xad, 0x76, 0x27, 0x3e, 0x81, 0xe0, 0x43, 0xe8,
	0xa3, 0xb8, 0x2c, 0xb8, 0x71, 0x69, 0xa3, 0x0b, 0x97, 0x7d, 0x04, 0xc9, 0x9f, 0x9a, 0x60, 0x77,
	0xe7, 0xfb, 0xce, 0xb9, 0xdf, 0xfd, 0x25, 0xe7, 0x82, 0x8e, 0xd3, 0x73, 0x47, 0xb4, 0xfd, 0xc0,
	0x93, 0x1e, 0xd5, 0x12, 0x51, 0x6f, 0x4c, 0x3c, 0x6f, 0x32, 0xc5, 0x0e, 0xf7, 0x9d, 0x0e, 0x17,
	0xc2, 0x93, 0x5c, 0x3a, 0x9e, 0x08, 0xd3, 0x21, 0x73, 0x17, 0x60, 0x78, 0x82, 0x42, 0x3a, 0xa7,
	0x0e, 0x06, 0xb4, 0x06, 0x65, 0x89, 0xdc, 0x1d, 0x0e, 0x0c, 0xd2, 0x24, 0xad, 0x8a, 0x9d, 0xa9,
	0xd8, 0x9f, 0x85, 0x18, 0x0c, 0x07, 0x86, 0x9a, 0xfa, 0xa9, 0x32, 0x6f, 0x09, 0xe8, 0xd6, 0x85,
	0x8f, 0x81, 0xe3, 0xa2, 0x90, 0x21, 0xb5, 0x40, 0xc7, 0x5c, 0x1a, 0xa4, 0x59, 0x6a, 0xe9, 0xbd,
	0x9f, 0xed, 0x94, 0xaa, 0x30, 0x58, 0xac, 0x2d, 0x21, 0x83, 0x85, 0x5d, 0x3c, 0x57, 0x3f, 0x84,
	0xea, 0xc7, 0x01, 0x5a, 0x85, 0xd2, 0x19, 0x2e, 0x32, 0xae, 0xb8, 0xa4, 0xbf, 0x41, 0x9b, 0xf3,
	0xe9, 0x0c, 0x13, 0x26, 0xbd, 0xf7, 0x75, 0xeb, 0x1a, 0x3b, 0xed, 0xf7, 0xd5, 0x1d, 0x62, 0x8e,
	0x00, 0xf2, 0x06, 0x6d, 0x40, 0x45, 0x70, 0x17, 0x43, 0x9f, 0x1f, 0x63, 0x16, 0x99, 0x1b, 0xf4,
	0x17, 0x94, 0x7d, 0x1e, 0x70, 0x37, 0x34, 0xd4, 0xe4, 0x03, 0x3e, 0x67, 0xc9, 0xa3, 0xd8, 0xb4,
	0xb3, 0x9e, 0xd9, 0x05, 0x2d, 0x31, 0x28, 0x85, 0x4f, 0xf1, 0xd9, 0x2c, 0x27, 0xa9, 0xe9, 0xb7,
	0x22, 0x5b, 0x25, 0x03, 0xe9, 0x8d, 0x41, 0xb3, 0xe2, 0x24, 0x7a, 0x04, 0x5f, 0xf6, 0x51, 0x1e,
	0x6c, 0x6e, 0x0c, 0xe9, 0x06, 0x3e, 0xdf, 0x45, 0x9d, 0x6e, 0xff, 0x36, 0xf3, 0xc7, 0xd5, 0xc3,
	0xcb, 0x8d, 0xfa, 0xdd, 0xac, 0x25, 0xdb, 0x9c, 0x77, 0x3b, 0x13, 0x94, 0xff, 0xde, 0xc1, 0xc3,
	0x3e, 0xf9, 0xb3, 0xf7, 0x77, 0xb9, 0x62, 0xca, 0xe3, 0x8a, 0x29, 0xeb, 0x15, 0x23, 0x97, 0x11,
	0x23, 0x77, 0x11, 0x23, 0xf7, 0x11, 0x23, 0xcb, 0x88, 0x91, 0xa7, 0x88, 0x91, 0xd7, 0x88, 0x29,
	0xeb, 0x88, 0x91, 0xeb, 0x67, 0xa6, 0x8c, 0xcb, 0xc9, 0x2b, 0xf8, 0xff, 0x16, 0x00, 0x00, 0xff,
	0xff, 0x44, 0x88, 0xf0, 0xdf, 0x39, 0x02, 0x00, 0x00,
}
