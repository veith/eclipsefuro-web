// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: project.proto

package project

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import person "../person"
import _type "google/type"

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

// Project description
type Project struct {
	// Project cost limit
	CostLimit *_type.Money `protobuf:"bytes,7,opt,name=cost_limit,json=costLimit" json:"cost_limit,omitempty"`
	// Short project description
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Localized String representation of a project
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Prospective end date of the project
	End *_type.Date `protobuf:"bytes,4,opt,name=end" json:"end,omitempty"`
	// Identity of a project
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// List of project members
	Members []*person.Person `protobuf:"bytes,6,rep,name=members" json:"members,omitempty"`
	// Start date of the project
	Start                *_type.Date `protobuf:"bytes,3,opt,name=start" json:"start,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_4810e3e837155f51, []int{0}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Project.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(dst, src)
}
func (m *Project) XXX_Size() int {
	return m.Size()
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetCostLimit() *_type.Money {
	if m != nil {
		return m.CostLimit
	}
	return nil
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Project) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Project) GetEnd() *_type.Date {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetMembers() []*person.Person {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Project) GetStart() *_type.Date {
	if m != nil {
		return m.Start
	}
	return nil
}

func init() {
	proto.RegisterType((*Project)(nil), "project.Project")
}
func (m *Project) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Project) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProject(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.DisplayName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProject(dAtA, i, uint64(len(m.DisplayName)))
		i += copy(dAtA[i:], m.DisplayName)
	}
	if m.Start != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProject(dAtA, i, uint64(m.Start.Size()))
		n1, err := m.Start.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.End != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProject(dAtA, i, uint64(m.End.Size()))
		n2, err := m.End.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProject(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Members) > 0 {
		for _, msg := range m.Members {
			dAtA[i] = 0x32
			i++
			i = encodeVarintProject(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.CostLimit != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintProject(dAtA, i, uint64(m.CostLimit.Size()))
		n3, err := m.CostLimit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintProject(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Project) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if m.Start != nil {
		l = m.Start.Size()
		n += 1 + l + sovProject(uint64(l))
	}
	if m.End != nil {
		l = m.End.Size()
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	if m.CostLimit != nil {
		l = m.CostLimit.Size()
		n += 1 + l + sovProject(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProject(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProject(x uint64) (n int) {
	return sovProject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Project) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
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
			return fmt.Errorf("proto: Project: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Project: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Start == nil {
				m.Start = &_type.Date{}
			}
			if err := m.Start.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.End == nil {
				m.End = &_type.Date{}
			}
			if err := m.End.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &person.Person{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CostLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CostLimit == nil {
				m.CostLimit = &_type.Money{}
			}
			if err := m.CostLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProject
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
func skipProject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProject
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
					return 0, ErrIntOverflowProject
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
					return 0, ErrIntOverflowProject
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
				return 0, ErrInvalidLengthProject
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProject
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
				next, err := skipProject(dAtA[start:])
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
	ErrInvalidLengthProject = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProject   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("project.proto", fileDescriptor_project_4810e3e837155f51) }

var fileDescriptor_project_4810e3e837155f51 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x49, 0xc7, 0x99, 0x32, 0xe9, 0x38, 0x60, 0x04, 0x0d, 0xb3, 0x28, 0x55, 0x17, 0x76,
	0xd5, 0xe2, 0x78, 0x03, 0x71, 0xa9, 0x32, 0xf4, 0x02, 0x43, 0xa6, 0x79, 0x0c, 0x91, 0xa6, 0x09,
	0xc9, 0xdb, 0xf4, 0x08, 0xde, 0xcc, 0xa5, 0x47, 0x90, 0x9e, 0x44, 0xda, 0x54, 0xd4, 0x85, 0xab,
	0x47, 0xfe, 0xef, 0x27, 0x1f, 0xef, 0xd1, 0x53, 0xeb, 0xcc, 0x2b, 0xd4, 0x58, 0x58, 0x67, 0xd0,
	0xb0, 0x78, 0x7a, 0x6e, 0x2e, 0x8f, 0xc6, 0x1c, 0x1b, 0x28, 0xb1, 0xb3, 0x50, 0x6a, 0xd3, 0x42,
	0x17, 0x1a, 0x9b, 0x8b, 0xdf, 0x40, 0x0a, 0x84, 0x29, 0x3f, 0xb7, 0xe0, 0xbc, 0x69, 0xcb, 0x30,
	0x42, 0x78, 0xfd, 0x16, 0xd1, 0x78, 0x17, 0x7e, 0x64, 0x6b, 0x1a, 0x29, 0xc9, 0x49, 0x46, 0xf2,
	0x65, 0x15, 0x29, 0xc9, 0xae, 0xe8, 0x4a, 0x2a, 0x6f, 0x1b, 0xd1, 0xed, 0x5b, 0xa1, 0x81, 0x47,
	0x23, 0x49, 0xa6, 0xec, 0x45, 0x68, 0x60, 0xb7, 0x74, 0xee, 0x51, 0x38, 0xe4, 0xb3, 0x8c, 0xe4,
	0xc9, 0xf6, 0xac, 0x08, 0xee, 0x62, 0x70, 0x17, 0x8f, 0x02, 0xa1, 0x0a, 0x9c, 0xdd, 0xd0, 0x19,
	0xb4, 0x92, 0x9f, 0xfc, 0x57, 0x1b, 0x28, 0xcb, 0x68, 0x22, 0xc1, 0xd7, 0x4e, 0x59, 0x54, 0xa6,
	0xe5, 0xf3, 0xc9, 0xf7, 0x13, 0xb1, 0x9c, 0xc6, 0x1a, 0xf4, 0x01, 0x9c, 0xe7, 0x8b, 0x6c, 0x96,
	0x27, 0xdb, 0x75, 0x31, 0xad, 0xb3, 0x1b, 0x47, 0xf5, 0x8d, 0xd9, 0x1d, 0xa5, 0xb5, 0xf1, 0xb8,
	0x6f, 0x94, 0x56, 0xc8, 0xe3, 0xd1, 0xcb, 0xfe, 0x78, 0x9f, 0x87, 0x9b, 0x55, 0xcb, 0xa1, 0xf5,
	0x34, 0x94, 0x1e, 0x56, 0xef, 0x7d, 0x4a, 0x3e, 0xfa, 0x94, 0x7c, 0xf6, 0x29, 0x39, 0x2c, 0xc6,
	0x03, 0xdd, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x23, 0x42, 0xda, 0x80, 0x01, 0x00, 0x00,
}
