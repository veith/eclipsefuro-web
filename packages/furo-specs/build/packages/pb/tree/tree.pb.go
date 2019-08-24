// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tree.proto

package tree

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

// Navigation tree type with recursive navigation nodes
type Tree struct {
	// description of the tree
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// String representation of the tree
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Id of the tree
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Rootnode of the tree
	Root                 *Navigationnode `protobuf:"bytes,10,opt,name=root" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Tree) Reset()         { *m = Tree{} }
func (m *Tree) String() string { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()    {}
func (*Tree) Descriptor() ([]byte, []int) {
	return fileDescriptor_tree_110ddb9b48c1436a, []int{0}
}
func (m *Tree) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tree.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Tree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tree.Merge(dst, src)
}
func (m *Tree) XXX_Size() int {
	return m.Size()
}
func (m *Tree) XXX_DiscardUnknown() {
	xxx_messageInfo_Tree.DiscardUnknown(m)
}

var xxx_messageInfo_Tree proto.InternalMessageInfo

func (m *Tree) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tree) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Tree) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tree) GetRoot() *Navigationnode {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*Tree)(nil), "tree.Tree")
}
func (m *Tree) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tree) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTree(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.DisplayName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTree(dAtA, i, uint64(len(m.DisplayName)))
		i += copy(dAtA[i:], m.DisplayName)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTree(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if m.Root != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintTree(dAtA, i, uint64(m.Root.Size()))
		n1, err := m.Root.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTree(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Tree) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTree(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovTree(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTree(uint64(l))
	}
	if m.Root != nil {
		l = m.Root.Size()
		n += 1 + l + sovTree(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTree(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTree(x uint64) (n int) {
	return sovTree(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tree) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTree
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
			return fmt.Errorf("proto: Tree: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tree: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
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
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Root == nil {
				m.Root = &Navigationnode{}
			}
			if err := m.Root.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTree
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
func skipTree(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTree
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
					return 0, ErrIntOverflowTree
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
					return 0, ErrIntOverflowTree
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
				return 0, ErrInvalidLengthTree
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTree
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
				next, err := skipTree(dAtA[start:])
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
	ErrInvalidLengthTree = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTree   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tree.proto", fileDescriptor_tree_110ddb9b48c1436a) }

var fileDescriptor_tree_110ddb9b48c1436a = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x29, 0x4a, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x44, 0xf2, 0x12, 0xcb, 0x32,
	0xd3, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xf2, 0xf2, 0x53, 0xa0, 0x72, 0x4a, 0xad, 0x8c, 0x5c, 0x2c,
	0x21, 0x45, 0xa9, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x4c, 0x99, 0x29, 0x42, 0x8a, 0x5c, 0x3c, 0x29, 0x99, 0xc5, 0x05, 0x39, 0x89, 0x95, 0xf1,
	0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x19, 0x6e, 0xa8, 0x98, 0x5f, 0x62, 0x6e, 0xaa, 0x90,
	0x02, 0x17, 0x77, 0x4a, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x01, 0xc8, 0x50, 0x09, 0x66, 0xa8, 0x0a,
	0x84, 0x90, 0x90, 0x06, 0x17, 0x4b, 0x51, 0x7e, 0x7e, 0x89, 0x04, 0x97, 0x02, 0xa3, 0x06, 0xb7,
	0x91, 0x88, 0x1e, 0xd8, 0x51, 0x7e, 0x28, 0xee, 0x08, 0x02, 0xab, 0x70, 0xe2, 0x39, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x93, 0xd8, 0xc0, 0x8e, 0x33, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x10, 0xaa, 0x4d, 0x9f, 0xc6, 0x00, 0x00, 0x00,
}
