// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gateway.proto

package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CMD_GW_ENUM int32

const (
	CMD_GW_INVALID      CMD_GW_ENUM = 0
	CMD_GW_VERIFY_TOKEN CMD_GW_ENUM = 1
)

var CMD_GW_ENUM_name = map[int32]string{
	0: "INVALID",
	1: "VERIFY_TOKEN",
}
var CMD_GW_ENUM_value = map[string]int32{
	"INVALID":      0,
	"VERIFY_TOKEN": 1,
}

func (x CMD_GW_ENUM) String() string {
	return proto.EnumName(CMD_GW_ENUM_name, int32(x))
}
func (CMD_GW_ENUM) EnumDescriptor() ([]byte, []int) { return fileDescriptorGateway, []int{0, 0} }

type CMD_GW struct {
}

func (m *CMD_GW) Reset()                    { *m = CMD_GW{} }
func (m *CMD_GW) String() string            { return proto.CompactTextString(m) }
func (*CMD_GW) ProtoMessage()               {}
func (*CMD_GW) Descriptor() ([]byte, []int) { return fileDescriptorGateway, []int{0} }

type MSG_GW_VERIFY_TOKEN struct {
	Id    *SERVER_ID `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Token string     `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (m *MSG_GW_VERIFY_TOKEN) Reset()                    { *m = MSG_GW_VERIFY_TOKEN{} }
func (m *MSG_GW_VERIFY_TOKEN) String() string            { return proto.CompactTextString(m) }
func (*MSG_GW_VERIFY_TOKEN) ProtoMessage()               {}
func (*MSG_GW_VERIFY_TOKEN) Descriptor() ([]byte, []int) { return fileDescriptorGateway, []int{1} }

func (m *MSG_GW_VERIFY_TOKEN) GetId() *SERVER_ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MSG_GW_VERIFY_TOKEN) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*CMD_GW)(nil), "protocol.CMD_GW")
	proto.RegisterType((*MSG_GW_VERIFY_TOKEN)(nil), "protocol.MSG_GW_VERIFY_TOKEN")
	proto.RegisterEnum("protocol.CMD_GW_ENUM", CMD_GW_ENUM_name, CMD_GW_ENUM_value)
}
func (m *CMD_GW) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CMD_GW) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MSG_GW_VERIFY_TOKEN) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MSG_GW_VERIFY_TOKEN) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGateway(dAtA, i, uint64(m.Id.Size()))
		n1, err := m.Id.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGateway(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func encodeFixed64Gateway(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Gateway(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGateway(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CMD_GW) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MSG_GW_VERIFY_TOKEN) Size() (n int) {
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovGateway(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovGateway(uint64(l))
	}
	return n
}

func sovGateway(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGateway(x uint64) (n int) {
	return sovGateway(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CMD_GW) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGateway
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
			return fmt.Errorf("proto: CMD_GW: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CMD_GW: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGateway
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
func (m *MSG_GW_VERIFY_TOKEN) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGateway
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
			return fmt.Errorf("proto: MSG_GW_VERIFY_TOKEN: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MSG_GW_VERIFY_TOKEN: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
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
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &SERVER_ID{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateway
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
				return ErrInvalidLengthGateway
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGateway
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
func skipGateway(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGateway
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
					return 0, ErrIntOverflowGateway
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
					return 0, ErrIntOverflowGateway
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
				return 0, ErrInvalidLengthGateway
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGateway
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
				next, err := skipGateway(dAtA[start:])
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
	ErrInvalidLengthGateway = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGateway   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("gateway.proto", fileDescriptorGateway) }

var fileDescriptorGateway = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39,
	0x52, 0x3c, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x10, 0x71, 0x25, 0x7d, 0x2e, 0x36, 0x67, 0x5f,
	0x97, 0x78, 0xf7, 0x70, 0x25, 0x55, 0x2e, 0x16, 0x57, 0xbf, 0x50, 0x5f, 0x21, 0x6e, 0x2e, 0x76,
	0x4f, 0xbf, 0x30, 0x47, 0x1f, 0x4f, 0x17, 0x01, 0x06, 0x21, 0x01, 0x2e, 0x9e, 0x30, 0xd7, 0x20,
	0x4f, 0xb7, 0xc8, 0xf8, 0x10, 0x7f, 0x6f, 0x57, 0x3f, 0x01, 0x46, 0xa5, 0x00, 0x2e, 0x61, 0xdf,
	0x60, 0xf7, 0x78, 0xf7, 0xf0, 0x78, 0x64, 0x09, 0x21, 0x65, 0x2e, 0x26, 0xcf, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x6e, 0x23, 0x61, 0x3d, 0x98, 0x65, 0x7a, 0xc1, 0xae, 0x41, 0x61, 0xae, 0x41,
	0xf1, 0x9e, 0x2e, 0x41, 0x4c, 0x9e, 0x29, 0x42, 0x22, 0x5c, 0xac, 0x21, 0xf9, 0xd9, 0xa9, 0x79,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x93, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0x58, 0xbf,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x00, 0xd3, 0x8e, 0xc4, 0x00, 0x00, 0x00,
}
