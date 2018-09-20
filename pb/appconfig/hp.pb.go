// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hp.proto

package appconfig

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

type SingInput struct {
	Query         string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int64  `protobuf:"zigzag64,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage int64  `protobuf:"zigzag64,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
}

func (m *SingInput) Reset()         { *m = SingInput{} }
func (m *SingInput) String() string { return proto.CompactTextString(m) }
func (*SingInput) ProtoMessage()    {}
func (*SingInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_hp_339a8fa1582be41f, []int{0}
}
func (m *SingInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SingInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SingInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SingInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingInput.Merge(dst, src)
}
func (m *SingInput) XXX_Size() int {
	return m.Size()
}
func (m *SingInput) XXX_DiscardUnknown() {
	xxx_messageInfo_SingInput.DiscardUnknown(m)
}

var xxx_messageInfo_SingInput proto.InternalMessageInfo

func (m *SingInput) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SingInput) GetPageNumber() int64 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SingInput) GetResultPerPage() int64 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

type SingOutput struct {
	Url      string        `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title    string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets []string      `protobuf:"bytes,3,rep,name=snippets" json:"snippets,omitempty"`
	Results  []*SingResult `protobuf:"bytes,4,rep,name=results" json:"results,omitempty"`
}

func (m *SingOutput) Reset()         { *m = SingOutput{} }
func (m *SingOutput) String() string { return proto.CompactTextString(m) }
func (*SingOutput) ProtoMessage()    {}
func (*SingOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_hp_339a8fa1582be41f, []int{1}
}
func (m *SingOutput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SingOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SingOutput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SingOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingOutput.Merge(dst, src)
}
func (m *SingOutput) XXX_Size() int {
	return m.Size()
}
func (m *SingOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_SingOutput.DiscardUnknown(m)
}

var xxx_messageInfo_SingOutput proto.InternalMessageInfo

func (m *SingOutput) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SingOutput) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SingOutput) GetSnippets() []string {
	if m != nil {
		return m.Snippets
	}
	return nil
}

func (m *SingOutput) GetResults() []*SingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type SingResult struct {
	Url      string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets []string `protobuf:"bytes,3,rep,name=snippets" json:"snippets,omitempty"`
}

func (m *SingResult) Reset()         { *m = SingResult{} }
func (m *SingResult) String() string { return proto.CompactTextString(m) }
func (*SingResult) ProtoMessage()    {}
func (*SingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_hp_339a8fa1582be41f, []int{2}
}
func (m *SingResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SingResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingResult.Merge(dst, src)
}
func (m *SingResult) XXX_Size() int {
	return m.Size()
}
func (m *SingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SingResult.DiscardUnknown(m)
}

var xxx_messageInfo_SingResult proto.InternalMessageInfo

func (m *SingResult) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SingResult) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SingResult) GetSnippets() []string {
	if m != nil {
		return m.Snippets
	}
	return nil
}

func init() {
	proto.RegisterType((*SingInput)(nil), "hprose.SingInput")
	proto.RegisterType((*SingOutput)(nil), "hprose.SingOutput")
	proto.RegisterType((*SingResult)(nil), "hprose.SingResult")
}
func (m *SingInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SingInput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Query) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHp(dAtA, i, uint64(len(m.Query)))
		i += copy(dAtA[i:], m.Query)
	}
	if m.PageNumber != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintHp(dAtA, i, uint64((uint64(m.PageNumber)<<1)^uint64((m.PageNumber>>63))))
	}
	if m.ResultPerPage != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHp(dAtA, i, uint64((uint64(m.ResultPerPage)<<1)^uint64((m.ResultPerPage>>63))))
	}
	return i, nil
}

func (m *SingOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SingOutput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHp(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHp(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Snippets) > 0 {
		for _, s := range m.Snippets {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			dAtA[i] = 0x22
			i++
			i = encodeVarintHp(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SingResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SingResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHp(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHp(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Snippets) > 0 {
		for _, s := range m.Snippets {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintHp(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SingInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovHp(uint64(l))
	}
	if m.PageNumber != 0 {
		n += 1 + sozHp(uint64(m.PageNumber))
	}
	if m.ResultPerPage != 0 {
		n += 1 + sozHp(uint64(m.ResultPerPage))
	}
	return n
}

func (m *SingOutput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovHp(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovHp(uint64(l))
	}
	if len(m.Snippets) > 0 {
		for _, s := range m.Snippets {
			l = len(s)
			n += 1 + l + sovHp(uint64(l))
		}
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovHp(uint64(l))
		}
	}
	return n
}

func (m *SingResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovHp(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovHp(uint64(l))
	}
	if len(m.Snippets) > 0 {
		for _, s := range m.Snippets {
			l = len(s)
			n += 1 + l + sovHp(uint64(l))
		}
	}
	return n
}

func sovHp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHp(x uint64) (n int) {
	return sovHp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SingInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHp
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
			return fmt.Errorf("proto: SingInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SingInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
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
				return ErrInvalidLengthHp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageNumber", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.PageNumber = int64(v)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultPerPage", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.ResultPerPage = int64(v)
		default:
			iNdEx = preIndex
			skippy, err := skipHp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHp
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
func (m *SingOutput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHp
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
			return fmt.Errorf("proto: SingOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SingOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
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
				return ErrInvalidLengthHp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
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
				return ErrInvalidLengthHp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snippets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
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
				return ErrInvalidLengthHp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Snippets = append(m.Snippets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
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
				return ErrInvalidLengthHp
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &SingResult{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHp
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
func (m *SingResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHp
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
			return fmt.Errorf("proto: SingResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SingResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
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
				return ErrInvalidLengthHp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
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
				return ErrInvalidLengthHp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snippets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHp
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
				return ErrInvalidLengthHp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Snippets = append(m.Snippets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHp
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
func skipHp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHp
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
					return 0, ErrIntOverflowHp
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
					return 0, ErrIntOverflowHp
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
				return 0, ErrInvalidLengthHp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHp
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
				next, err := skipHp(dAtA[start:])
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
	ErrInvalidLengthHp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHp   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("hp.proto", fileDescriptor_hp_339a8fa1582be41f) }

var fileDescriptor_hp_339a8fa1582be41f = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd0, 0x3d, 0x4e, 0xc4, 0x30,
	0x10, 0x05, 0xe0, 0x18, 0xc3, 0xb2, 0x99, 0x15, 0x02, 0x59, 0x14, 0xd6, 0x16, 0x26, 0x4a, 0x81,
	0x52, 0xa0, 0x14, 0x70, 0x03, 0x2a, 0x68, 0x20, 0x32, 0x07, 0x88, 0x76, 0xa5, 0x51, 0x36, 0x28,
	0x24, 0xc6, 0x3f, 0x05, 0x0d, 0x67, 0xe0, 0x58, 0x94, 0x29, 0x29, 0x51, 0x72, 0x11, 0x64, 0x1b,
	0x38, 0xc1, 0x76, 0x7e, 0x6f, 0x46, 0xfa, 0xe4, 0x81, 0xe5, 0x4e, 0x95, 0x4a, 0x0f, 0x76, 0x60,
	0x8b, 0x9d, 0xd2, 0x83, 0xc1, 0xfc, 0x19, 0xd2, 0xa7, 0xb6, 0x6f, 0xee, 0x7b, 0xe5, 0x2c, 0x3b,
	0x87, 0xa3, 0x57, 0x87, 0xfa, 0x8d, 0x93, 0x8c, 0x14, 0xa9, 0x8c, 0x81, 0x5d, 0xc0, 0x4a, 0x6d,
	0x1a, 0xac, 0x7b, 0xf7, 0xb2, 0x45, 0xcd, 0x0f, 0x32, 0x52, 0x30, 0x09, 0xbe, 0x7a, 0x08, 0x0d,
	0xbb, 0x84, 0x53, 0x8d, 0xc6, 0x75, 0xb6, 0x56, 0xa8, 0x6b, 0x3f, 0xe0, 0x34, 0x2c, 0x9d, 0xc4,
	0xba, 0x42, 0x5d, 0x6d, 0x1a, 0xcc, 0xdf, 0x01, 0xbc, 0xf5, 0xe8, 0xac, 0xc7, 0xce, 0x80, 0x3a,
	0xdd, 0xfd, 0x52, 0xfe, 0xe9, 0x79, 0xdb, 0xda, 0x0e, 0x03, 0x91, 0xca, 0x18, 0xd8, 0x1a, 0x96,
	0xa6, 0x6f, 0x95, 0x42, 0x6b, 0x38, 0xcd, 0x68, 0x91, 0xca, 0xff, 0xcc, 0xae, 0xe0, 0x38, 0x12,
	0x86, 0x1f, 0x66, 0xb4, 0x58, 0x5d, 0xb3, 0x32, 0xfe, 0xab, 0xf4, 0x90, 0x0c, 0x23, 0xf9, 0xb7,
	0x92, 0x57, 0xd1, 0x8f, 0xf5, 0x3e, 0xfc, 0xdb, 0xf5, 0xe7, 0x24, 0xc8, 0x38, 0x09, 0xf2, 0x3d,
	0x09, 0xf2, 0x31, 0x8b, 0x64, 0x9c, 0x45, 0xf2, 0x35, 0x8b, 0xe4, 0x8e, 0x6c, 0x17, 0xe1, 0xd0,
	0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x0c, 0x35, 0x73, 0x74, 0x01, 0x00, 0x00,
}
