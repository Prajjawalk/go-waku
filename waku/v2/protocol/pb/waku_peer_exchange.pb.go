// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: waku_peer_exchange.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type PeerInfo struct {
	ENR                  []byte   `protobuf:"bytes,1,opt,name=ENR,proto3" json:"ENR,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerInfo) Reset()         { *m = PeerInfo{} }
func (m *PeerInfo) String() string { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()    {}
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce50192ba54b780f, []int{0}
}
func (m *PeerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerInfo.Merge(m, src)
}
func (m *PeerInfo) XXX_Size() int {
	return m.Size()
}
func (m *PeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerInfo proto.InternalMessageInfo

func (m *PeerInfo) GetENR() []byte {
	if m != nil {
		return m.ENR
	}
	return nil
}

type PeerExchangeQuery struct {
	NumPeers             uint64   `protobuf:"varint,1,opt,name=numPeers,proto3" json:"numPeers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerExchangeQuery) Reset()         { *m = PeerExchangeQuery{} }
func (m *PeerExchangeQuery) String() string { return proto.CompactTextString(m) }
func (*PeerExchangeQuery) ProtoMessage()    {}
func (*PeerExchangeQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce50192ba54b780f, []int{1}
}
func (m *PeerExchangeQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeerExchangeQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeerExchangeQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeerExchangeQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerExchangeQuery.Merge(m, src)
}
func (m *PeerExchangeQuery) XXX_Size() int {
	return m.Size()
}
func (m *PeerExchangeQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerExchangeQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PeerExchangeQuery proto.InternalMessageInfo

func (m *PeerExchangeQuery) GetNumPeers() uint64 {
	if m != nil {
		return m.NumPeers
	}
	return 0
}

type PeerExchangeResponse struct {
	PeerInfos            []*PeerInfo `protobuf:"bytes,1,rep,name=peerInfos,proto3" json:"peerInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PeerExchangeResponse) Reset()         { *m = PeerExchangeResponse{} }
func (m *PeerExchangeResponse) String() string { return proto.CompactTextString(m) }
func (*PeerExchangeResponse) ProtoMessage()    {}
func (*PeerExchangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce50192ba54b780f, []int{2}
}
func (m *PeerExchangeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeerExchangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeerExchangeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeerExchangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerExchangeResponse.Merge(m, src)
}
func (m *PeerExchangeResponse) XXX_Size() int {
	return m.Size()
}
func (m *PeerExchangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerExchangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeerExchangeResponse proto.InternalMessageInfo

func (m *PeerExchangeResponse) GetPeerInfos() []*PeerInfo {
	if m != nil {
		return m.PeerInfos
	}
	return nil
}

type PeerExchangeRPC struct {
	Query                *PeerExchangeQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Response             *PeerExchangeResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PeerExchangeRPC) Reset()         { *m = PeerExchangeRPC{} }
func (m *PeerExchangeRPC) String() string { return proto.CompactTextString(m) }
func (*PeerExchangeRPC) ProtoMessage()    {}
func (*PeerExchangeRPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce50192ba54b780f, []int{3}
}
func (m *PeerExchangeRPC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeerExchangeRPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeerExchangeRPC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeerExchangeRPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerExchangeRPC.Merge(m, src)
}
func (m *PeerExchangeRPC) XXX_Size() int {
	return m.Size()
}
func (m *PeerExchangeRPC) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerExchangeRPC.DiscardUnknown(m)
}

var xxx_messageInfo_PeerExchangeRPC proto.InternalMessageInfo

func (m *PeerExchangeRPC) GetQuery() *PeerExchangeQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *PeerExchangeRPC) GetResponse() *PeerExchangeResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerInfo)(nil), "PeerInfo")
	proto.RegisterType((*PeerExchangeQuery)(nil), "PeerExchangeQuery")
	proto.RegisterType((*PeerExchangeResponse)(nil), "PeerExchangeResponse")
	proto.RegisterType((*PeerExchangeRPC)(nil), "PeerExchangeRPC")
}

func init() { proto.RegisterFile("waku_peer_exchange.proto", fileDescriptor_ce50192ba54b780f) }

var fileDescriptor_ce50192ba54b780f = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x4f, 0xcc, 0x2e,
	0x8d, 0x2f, 0x48, 0x4d, 0x2d, 0x8a, 0x4f, 0xad, 0x48, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe1, 0xe2, 0x08, 0x48, 0x4d, 0x2d, 0xf2, 0xcc, 0x4b, 0xcb,
	0x17, 0x12, 0xe0, 0x62, 0x76, 0xf5, 0x0b, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0x31,
	0x95, 0xf4, 0xb9, 0x04, 0x41, 0xb2, 0xae, 0x50, 0x3d, 0x81, 0xa5, 0xa9, 0x45, 0x95, 0x42, 0x52,
	0x5c, 0x1c, 0x79, 0xa5, 0xb9, 0x20, 0xf1, 0x62, 0xb0, 0x5a, 0x96, 0x20, 0x38, 0x5f, 0xc9, 0x9e,
	0x4b, 0x04, 0x59, 0x43, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x3a, 0x17, 0x67,
	0x01, 0xd4, 0x1a, 0x90, 0x26, 0x66, 0x0d, 0x6e, 0x23, 0x4e, 0x3d, 0x98, 0xc5, 0x41, 0x08, 0x39,
	0xa5, 0x3c, 0x2e, 0x7e, 0x14, 0x03, 0x02, 0x9c, 0x85, 0x34, 0xb8, 0x58, 0x0b, 0x41, 0x16, 0x83,
	0x2d, 0xe3, 0x36, 0x12, 0xd2, 0xc3, 0x70, 0x52, 0x10, 0x44, 0x81, 0x90, 0x21, 0x17, 0x47, 0x11,
	0xd4, 0x46, 0x09, 0x26, 0xb0, 0x62, 0x51, 0x3d, 0x6c, 0xce, 0x09, 0x82, 0x2b, 0x73, 0x12, 0x38,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63,
	0x48, 0x62, 0x03, 0x07, 0x8c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x99, 0xd5, 0xbb, 0xb6, 0x34,
	0x01, 0x00, 0x00,
}

func (m *PeerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ENR) > 0 {
		i -= len(m.ENR)
		copy(dAtA[i:], m.ENR)
		i = encodeVarintWakuPeerExchange(dAtA, i, uint64(len(m.ENR)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PeerExchangeQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeerExchangeQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeerExchangeQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.NumPeers != 0 {
		i = encodeVarintWakuPeerExchange(dAtA, i, uint64(m.NumPeers))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PeerExchangeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeerExchangeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeerExchangeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PeerInfos) > 0 {
		for iNdEx := len(m.PeerInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeerInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWakuPeerExchange(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PeerExchangeRPC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeerExchangeRPC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeerExchangeRPC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Response != nil {
		{
			size, err := m.Response.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWakuPeerExchange(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Query != nil {
		{
			size, err := m.Query.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWakuPeerExchange(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWakuPeerExchange(dAtA []byte, offset int, v uint64) int {
	offset -= sovWakuPeerExchange(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PeerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ENR)
	if l > 0 {
		n += 1 + l + sovWakuPeerExchange(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PeerExchangeQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumPeers != 0 {
		n += 1 + sovWakuPeerExchange(uint64(m.NumPeers))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PeerExchangeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PeerInfos) > 0 {
		for _, e := range m.PeerInfos {
			l = e.Size()
			n += 1 + l + sovWakuPeerExchange(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PeerExchangeRPC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovWakuPeerExchange(uint64(l))
	}
	if m.Response != nil {
		l = m.Response.Size()
		n += 1 + l + sovWakuPeerExchange(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWakuPeerExchange(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWakuPeerExchange(x uint64) (n int) {
	return sovWakuPeerExchange(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PeerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWakuPeerExchange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ENR", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuPeerExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ENR = append(m.ENR[:0], dAtA[iNdEx:postIndex]...)
			if m.ENR == nil {
				m.ENR = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWakuPeerExchange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PeerExchangeQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWakuPeerExchange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeerExchangeQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeerExchangeQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumPeers", wireType)
			}
			m.NumPeers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuPeerExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumPeers |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWakuPeerExchange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PeerExchangeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWakuPeerExchange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeerExchangeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeerExchangeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuPeerExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerInfos = append(m.PeerInfos, &PeerInfo{})
			if err := m.PeerInfos[len(m.PeerInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWakuPeerExchange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PeerExchangeRPC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWakuPeerExchange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeerExchangeRPC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeerExchangeRPC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuPeerExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Query == nil {
				m.Query = &PeerExchangeQuery{}
			}
			if err := m.Query.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWakuPeerExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &PeerExchangeResponse{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWakuPeerExchange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWakuPeerExchange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWakuPeerExchange(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWakuPeerExchange
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
					return 0, ErrIntOverflowWakuPeerExchange
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWakuPeerExchange
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
			if length < 0 {
				return 0, ErrInvalidLengthWakuPeerExchange
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWakuPeerExchange
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWakuPeerExchange
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWakuPeerExchange        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWakuPeerExchange          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWakuPeerExchange = fmt.Errorf("proto: unexpected end of group")
)
