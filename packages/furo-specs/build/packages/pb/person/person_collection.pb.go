// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: person_collection.proto

package person

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import furo "furo"

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

// PersonCollection with repeated PersonEntity
type PersonCollection struct {
	// Contains a person.PersonEntity repeated
	Entities []*PersonEntity `protobuf:"bytes,4,rep,name=entities" json:"entities,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,3,rep,name=links" json:"links,omitempty"`
	// Meta for the response
	Meta                 *furo.Meta `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PersonCollection) Reset()         { *m = PersonCollection{} }
func (m *PersonCollection) String() string { return proto.CompactTextString(m) }
func (*PersonCollection) ProtoMessage()    {}
func (*PersonCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_person_collection_d97243f2d842b7f8, []int{0}
}
func (m *PersonCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PersonCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PersonCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PersonCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonCollection.Merge(dst, src)
}
func (m *PersonCollection) XXX_Size() int {
	return m.Size()
}
func (m *PersonCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonCollection.DiscardUnknown(m)
}

var xxx_messageInfo_PersonCollection proto.InternalMessageInfo

func (m *PersonCollection) GetEntities() []*PersonEntity {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *PersonCollection) GetLinks() []*furo.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PersonCollection) GetMeta() *furo.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*PersonCollection)(nil), "person.PersonCollection")
}
func (m *PersonCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PersonCollection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Meta != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPersonCollection(dAtA, i, uint64(m.Meta.Size()))
		n1, err := m.Meta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Links) > 0 {
		for _, msg := range m.Links {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintPersonCollection(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Entities) > 0 {
		for _, msg := range m.Entities {
			dAtA[i] = 0x22
			i++
			i = encodeVarintPersonCollection(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPersonCollection(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PersonCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovPersonCollection(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovPersonCollection(uint64(l))
		}
	}
	if len(m.Entities) > 0 {
		for _, e := range m.Entities {
			l = e.Size()
			n += 1 + l + sovPersonCollection(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPersonCollection(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPersonCollection(x uint64) (n int) {
	return sovPersonCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PersonCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPersonCollection
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
			return fmt.Errorf("proto: PersonCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PersonCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonCollection
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
				return ErrInvalidLengthPersonCollection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &furo.Meta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonCollection
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
				return ErrInvalidLengthPersonCollection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &furo.Link{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPersonCollection
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
				return ErrInvalidLengthPersonCollection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entities = append(m.Entities, &PersonEntity{})
			if err := m.Entities[len(m.Entities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPersonCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPersonCollection
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
func skipPersonCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPersonCollection
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
					return 0, ErrIntOverflowPersonCollection
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
					return 0, ErrIntOverflowPersonCollection
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
				return 0, ErrInvalidLengthPersonCollection
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPersonCollection
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
				next, err := skipPersonCollection(dAtA[start:])
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
	ErrInvalidLengthPersonCollection = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPersonCollection   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("person_collection.proto", fileDescriptor_person_collection_d97243f2d842b7f8)
}

var fileDescriptor_person_collection_d97243f2d842b7f8 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0x8b, 0x4f, 0xce, 0xcf, 0xc9, 0x49, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x48, 0x48, 0x09, 0x43, 0x15, 0xa4, 0xe6, 0x95, 0x64, 0x96,
	0x54, 0x42, 0x24, 0xa5, 0xf8, 0xd3, 0x4a, 0x8b, 0xf2, 0xf5, 0x73, 0x53, 0x4b, 0x12, 0x51, 0x04,
	0x72, 0x32, 0xf3, 0xb2, 0x21, 0x02, 0x4a, 0x6d, 0x8c, 0x5c, 0x02, 0x01, 0x60, 0x9d, 0xce, 0x70,
	0x93, 0x85, 0xe4, 0xb8, 0x58, 0x40, 0x7a, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xb8, 0xf4,
	0x40, 0x9a, 0xf4, 0x7c, 0x53, 0x4b, 0x12, 0x83, 0xc0, 0xe2, 0x42, 0x0a, 0x5c, 0xac, 0x20, 0x23,
	0x8a, 0x25, 0x98, 0x15, 0x98, 0x11, 0x0a, 0x7c, 0x32, 0xf3, 0xb2, 0x83, 0x20, 0x12, 0x42, 0x06,
	0x5c, 0x1c, 0x60, 0x87, 0x64, 0xa6, 0x16, 0x4b, 0xb0, 0x80, 0x15, 0x89, 0xe8, 0x41, 0x1c, 0xa8,
	0x07, 0xb1, 0xcd, 0x15, 0xec, 0xcc, 0x20, 0xb8, 0x2a, 0x27, 0x9e, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0x89, 0x0d, 0xec, 0x3a, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x46, 0x5d, 0xfa, 0xbe, 0xf7, 0x00, 0x00, 0x00,
}
