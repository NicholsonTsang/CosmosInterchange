// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibcdex/buy_order_book.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BuyOrderBook struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index       string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	AmountDenom string `protobuf:"bytes,3,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	PriceDenom  string `protobuf:"bytes,4,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
}

func (m *BuyOrderBook) Reset()         { *m = BuyOrderBook{} }
func (m *BuyOrderBook) String() string { return proto.CompactTextString(m) }
func (*BuyOrderBook) ProtoMessage()    {}
func (*BuyOrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_82585ed974178ec3, []int{0}
}
func (m *BuyOrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BuyOrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BuyOrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BuyOrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyOrderBook.Merge(m, src)
}
func (m *BuyOrderBook) XXX_Size() int {
	return m.Size()
}
func (m *BuyOrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyOrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_BuyOrderBook proto.InternalMessageInfo

func (m *BuyOrderBook) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *BuyOrderBook) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *BuyOrderBook) GetAmountDenom() string {
	if m != nil {
		return m.AmountDenom
	}
	return ""
}

func (m *BuyOrderBook) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*BuyOrderBook)(nil), "username.interchange.ibcdex.BuyOrderBook")
}

func init() { proto.RegisterFile("ibcdex/buy_order_book.proto", fileDescriptor_82585ed974178ec3) }

var fileDescriptor_82585ed974178ec3 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0xc6, 0xad, 0xfe, 0xa5, 0x6a, 0x27, 0x91, 0xc1, 0x34, 0x20, 0x42, 0xa7, 0x4e, 0xd2, 0xd0,
	0x37, 0x30, 0x1d, 0xba, 0x15, 0x3a, 0x76, 0x09, 0x96, 0x7c, 0x38, 0x22, 0x58, 0x67, 0x2e, 0x12,
	0xd8, 0x5b, 0x1f, 0xa1, 0x8f, 0xd5, 0x31, 0x63, 0xc7, 0x62, 0xbf, 0x48, 0x89, 0x1c, 0x83, 0xb7,
	0xfb, 0xbe, 0xfb, 0xbe, 0x83, 0xdf, 0xf1, 0xb5, 0x33, 0xb6, 0x82, 0x4e, 0x9b, 0xd8, 0x6f, 0x91,
	0x2a, 0xa0, 0xad, 0x41, 0xdc, 0xab, 0x96, 0x30, 0xa0, 0x58, 0xc7, 0x03, 0x90, 0x2f, 0x1b, 0x50,
	0xce, 0x07, 0x20, 0xbb, 0x2b, 0x7d, 0x0d, 0x6a, 0x6a, 0x3c, 0xae, 0x6a, 0xac, 0x31, 0xe5, 0xf4,
	0x69, 0x9a, 0x2a, 0x4f, 0x5f, 0x8c, 0x3f, 0x14, 0xb1, 0x7f, 0x3f, 0x9d, 0x2a, 0x10, 0xf7, 0x22,
	0xe7, 0xb7, 0x96, 0xa0, 0x0c, 0x48, 0x39, 0xdb, 0xb0, 0xe7, 0xbb, 0x8f, 0x59, 0x8a, 0x15, 0xbf,
	0x76, 0xbe, 0x82, 0x2e, 0xbf, 0x48, 0xfe, 0x24, 0xc4, 0x86, 0xdf, 0x97, 0x0d, 0x46, 0x1f, 0x5e,
	0xc1, 0x63, 0x93, 0x5f, 0xa6, 0xdd, 0xd2, 0x12, 0x92, 0xf3, 0x96, 0x9c, 0x85, 0x29, 0x70, 0x95,
	0x02, 0x0b, 0xa7, 0x78, 0xfb, 0x19, 0x24, 0x3b, 0x0e, 0x92, 0xfd, 0x0d, 0x92, 0x7d, 0x8f, 0x32,
	0x3b, 0x8e, 0x32, 0xfb, 0x1d, 0x65, 0xf6, 0xa9, 0x6a, 0x17, 0x76, 0xd1, 0x28, 0x8b, 0x8d, 0x9e,
	0xd1, 0xf4, 0x02, 0x4d, 0x77, 0xfa, 0xfc, 0x8e, 0xd0, 0xb7, 0x70, 0x30, 0x37, 0x89, 0xe9, 0xe5,
	0x3f, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x9d, 0x52, 0xe5, 0x25, 0x01, 0x00, 0x00,
}

func (m *BuyOrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BuyOrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BuyOrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AmountDenom) > 0 {
		i -= len(m.AmountDenom)
		copy(dAtA[i:], m.AmountDenom)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.AmountDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBuyOrderBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovBuyOrderBook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BuyOrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	l = len(m.AmountDenom)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	return n
}

func sovBuyOrderBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBuyOrderBook(x uint64) (n int) {
	return sovBuyOrderBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BuyOrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuyOrderBook
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
			return fmt.Errorf("proto: BuyOrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BuyOrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuyOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBuyOrderBook
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
func skipBuyOrderBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBuyOrderBook
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
					return 0, ErrIntOverflowBuyOrderBook
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
					return 0, ErrIntOverflowBuyOrderBook
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
				return 0, ErrInvalidLengthBuyOrderBook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBuyOrderBook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBuyOrderBook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBuyOrderBook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBuyOrderBook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBuyOrderBook = fmt.Errorf("proto: unexpected end of group")
)