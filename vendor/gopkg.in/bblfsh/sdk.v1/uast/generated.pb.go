// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gopkg.in/bblfsh/sdk.v1/uast/generated.proto

/*
	Package uast is a generated protocol buffer package.

	It is generated from these files:
		gopkg.in/bblfsh/sdk.v1/uast/generated.proto

	It has these top-level messages:
		Node
		Position
*/
package uast

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Role is the main UAST annotation. It indicates that a node in an AST can
// be interpreted as acting with certain language-independent role.
//
// go:generate stringer -type=Role
var Role_name = map[int32]string{
	0:   "INVALID",
	1:   "IDENTIFIER",
	2:   "QUALIFIED",
	3:   "OPERATOR",
	4:   "BINARY",
	5:   "UNARY",
	6:   "LEFT",
	7:   "RIGHT",
	8:   "INFIX",
	9:   "POSTFIX",
	10:  "BITWISE",
	11:  "BOOLEAN",
	12:  "UNSIGNED",
	13:  "LEFT_SHIFT",
	14:  "RIGHT_SHIFT",
	15:  "OR",
	16:  "XOR",
	17:  "AND",
	18:  "EXPRESSION",
	19:  "STATEMENT",
	20:  "EQUAL",
	21:  "NOT",
	22:  "LESS_THAN",
	23:  "LESS_THAN_OR_EQUAL",
	24:  "GREATER_THAN",
	25:  "GREATER_THAN_OR_EQUAL",
	26:  "IDENTICAL",
	27:  "CONTAINS",
	28:  "INCREMENT",
	29:  "DECREMENT",
	30:  "NEGATIVE",
	31:  "POSITIVE",
	32:  "DEREFERENCE",
	33:  "TAKE_ADDRESS",
	34:  "FILE",
	35:  "ADD",
	36:  "SUBSTRACT",
	37:  "MULTIPLY",
	38:  "DIVIDE",
	39:  "MODULO",
	40:  "PACKAGE",
	41:  "DECLARATION",
	42:  "IMPORT",
	43:  "PATHNAME",
	44:  "ALIAS",
	45:  "FUNCTION",
	46:  "BODY",
	47:  "NAME",
	48:  "RECEIVER",
	49:  "ARGUMENT",
	50:  "VALUE",
	51:  "ARGS_LIST",
	52:  "BASE",
	53:  "IMPLEMENTS",
	54:  "INSTANCE",
	55:  "SUBTYPE",
	56:  "SUBPACKAGE",
	57:  "MODULE",
	58:  "FRIEND",
	59:  "WORLD",
	60:  "IF",
	61:  "CONDITION",
	62:  "THEN",
	63:  "ELSE",
	64:  "SWITCH",
	65:  "CASE",
	66:  "DEFAULT",
	67:  "FOR",
	68:  "INITIALIZATION",
	69:  "UPDATE",
	70:  "ITERATOR",
	71:  "WHILE",
	72:  "DO_WHILE",
	73:  "BREAK",
	74:  "CONTINUE",
	75:  "GOTO",
	76:  "BLOCK",
	77:  "SCOPE",
	78:  "RETURN",
	79:  "TRY",
	80:  "CATCH",
	81:  "FINALLY",
	82:  "THROW",
	83:  "ASSERT",
	84:  "CALL",
	85:  "CALLEE",
	86:  "POSITIONAL",
	87:  "NOOP",
	88:  "LITERAL",
	89:  "BYTE",
	90:  "BYTE_STRING",
	91:  "CHARACTER",
	92:  "LIST",
	93:  "MAP",
	94:  "NULL",
	95:  "NUMBER",
	96:  "REGEXP",
	97:  "SET",
	98:  "STRING",
	99:  "TUPLE",
	100: "TYPE",
	101: "ENTRY",
	102: "KEY",
	103: "PRIMITIVE",
	104: "ASSIGNMENT",
	105: "THIS",
	106: "COMMENT",
	107: "DOCUMENTATION",
	108: "WHITESPACE",
	109: "INCOMPLETE",
	110: "UNANNOTATED",
	111: "VISIBILITY",
	112: "ANNOTATION",
	113: "ANONYMOUS",
	114: "ENUMERATION",
	115: "ARITHMETIC",
	116: "RELATIONAL",
	117: "VARIABLE",
}
var Role_value = map[string]int32{
	"INVALID":               0,
	"IDENTIFIER":            1,
	"QUALIFIED":             2,
	"OPERATOR":              3,
	"BINARY":                4,
	"UNARY":                 5,
	"LEFT":                  6,
	"RIGHT":                 7,
	"INFIX":                 8,
	"POSTFIX":               9,
	"BITWISE":               10,
	"BOOLEAN":               11,
	"UNSIGNED":              12,
	"LEFT_SHIFT":            13,
	"RIGHT_SHIFT":           14,
	"OR":                    15,
	"XOR":                   16,
	"AND":                   17,
	"EXPRESSION":            18,
	"STATEMENT":             19,
	"EQUAL":                 20,
	"NOT":                   21,
	"LESS_THAN":             22,
	"LESS_THAN_OR_EQUAL":    23,
	"GREATER_THAN":          24,
	"GREATER_THAN_OR_EQUAL": 25,
	"IDENTICAL":             26,
	"CONTAINS":              27,
	"INCREMENT":             28,
	"DECREMENT":             29,
	"NEGATIVE":              30,
	"POSITIVE":              31,
	"DEREFERENCE":           32,
	"TAKE_ADDRESS":          33,
	"FILE":                  34,
	"ADD":                   35,
	"SUBSTRACT":             36,
	"MULTIPLY":              37,
	"DIVIDE":                38,
	"MODULO":                39,
	"PACKAGE":               40,
	"DECLARATION":           41,
	"IMPORT":                42,
	"PATHNAME":              43,
	"ALIAS":                 44,
	"FUNCTION":              45,
	"BODY":                  46,
	"NAME":                  47,
	"RECEIVER":              48,
	"ARGUMENT":              49,
	"VALUE":                 50,
	"ARGS_LIST":             51,
	"BASE":                  52,
	"IMPLEMENTS":            53,
	"INSTANCE":              54,
	"SUBTYPE":               55,
	"SUBPACKAGE":            56,
	"MODULE":                57,
	"FRIEND":                58,
	"WORLD":                 59,
	"IF":                    60,
	"CONDITION":             61,
	"THEN":                  62,
	"ELSE":                  63,
	"SWITCH":                64,
	"CASE":                  65,
	"DEFAULT":               66,
	"FOR":                   67,
	"INITIALIZATION":        68,
	"UPDATE":                69,
	"ITERATOR":              70,
	"WHILE":                 71,
	"DO_WHILE":              72,
	"BREAK":                 73,
	"CONTINUE":              74,
	"GOTO":                  75,
	"BLOCK":                 76,
	"SCOPE":                 77,
	"RETURN":                78,
	"TRY":                   79,
	"CATCH":                 80,
	"FINALLY":               81,
	"THROW":                 82,
	"ASSERT":                83,
	"CALL":                  84,
	"CALLEE":                85,
	"POSITIONAL":            86,
	"NOOP":                  87,
	"LITERAL":               88,
	"BYTE":                  89,
	"BYTE_STRING":           90,
	"CHARACTER":             91,
	"LIST":                  92,
	"MAP":                   93,
	"NULL":                  94,
	"NUMBER":                95,
	"REGEXP":                96,
	"SET":                   97,
	"STRING":                98,
	"TUPLE":                 99,
	"TYPE":                  100,
	"ENTRY":                 101,
	"KEY":                   102,
	"PRIMITIVE":             103,
	"ASSIGNMENT":            104,
	"THIS":                  105,
	"COMMENT":               106,
	"DOCUMENTATION":         107,
	"WHITESPACE":            108,
	"INCOMPLETE":            109,
	"UNANNOTATED":           110,
	"VISIBILITY":            111,
	"ANNOTATION":            112,
	"ANONYMOUS":             113,
	"ENUMERATION":           114,
	"ARITHMETIC":            115,
	"RELATIONAL":            116,
	"VARIABLE":              117,
}

func (Role) EnumDescriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *Node) Reset()                    { *m = Node{} }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func init() {
	proto.RegisterType((*Node)(nil), "gopkg.in.bblfsh.sdk.v1.uast.Node")
	proto.RegisterType((*Position)(nil), "gopkg.in.bblfsh.sdk.v1.uast.Position")
	proto.RegisterEnum("gopkg.in.bblfsh.sdk.v1.uast.Role", Role_name, Role_value)
}
func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.InternalType) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.InternalType)))
		i += copy(dAtA[i:], m.InternalType)
	}
	if len(m.Properties) > 0 {
		for k, _ := range m.Properties {
			dAtA[i] = 0x12
			i++
			v := m.Properties[k]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Children) > 0 {
		for _, msg := range m.Children {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if m.StartPosition != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.StartPosition.ProtoSize()))
		n1, err := m.StartPosition.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.EndPosition != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.EndPosition.ProtoSize()))
		n2, err := m.EndPosition.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Roles) > 0 {
		dAtA4 := make([]byte, len(m.Roles)*10)
		var j3 int
		for _, num := range m.Roles {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x3a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	return i, nil
}

func (m *Position) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Position) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Offset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Offset))
	}
	if m.Line != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Line))
	}
	if m.Col != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Col))
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Node) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.InternalType)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Properties) > 0 {
		for k, v := range m.Properties {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if len(m.Children) > 0 {
		for _, e := range m.Children {
			l = e.ProtoSize()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.StartPosition != nil {
		l = m.StartPosition.ProtoSize()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.EndPosition != nil {
		l = m.EndPosition.ProtoSize()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Roles) > 0 {
		l = 0
		for _, e := range m.Roles {
			l += sovGenerated(uint64(e))
		}
		n += 1 + sovGenerated(uint64(l)) + l
	}
	return n
}

func (m *Position) ProtoSize() (n int) {
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + sovGenerated(uint64(m.Offset))
	}
	if m.Line != 0 {
		n += 1 + sovGenerated(uint64(m.Line))
	}
	if m.Col != 0 {
		n += 1 + sovGenerated(uint64(m.Col))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InternalType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Properties == nil {
				m.Properties = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
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
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthGenerated
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Properties[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Properties[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, &Node{})
			if err := m.Children[len(m.Children)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartPosition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartPosition == nil {
				m.StartPosition = &Position{}
			}
			if err := m.StartPosition.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndPosition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndPosition == nil {
				m.EndPosition = &Position{}
			}
			if err := m.EndPosition.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType == 0 {
				var v Role
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (Role(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Roles = append(m.Roles, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenerated
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v Role
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (Role(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Roles = append(m.Roles, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *Position) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Position: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Position: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			m.Line = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Line |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Col", wireType)
			}
			m.Col = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Col |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("gopkg.in/bblfsh/sdk.v1/uast/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 2196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xd9, 0x92, 0xe3, 0xb6,
	0xd5, 0xee, 0x7d, 0x41, 0x2f, 0x03, 0xd3, 0xcb, 0x2f, 0xcb, 0x63, 0x8d, 0xbc, 0xfe, 0x13, 0x3b,
	0xee, 0xce, 0xd8, 0x49, 0xec, 0x4c, 0xe2, 0x24, 0x10, 0x09, 0x49, 0x48, 0x53, 0xa4, 0x06, 0x04,
	0x7b, 0x71, 0x96, 0x0e, 0x25, 0x41, 0x12, 0xd2, 0x14, 0x29, 0x93, 0x54, 0xdb, 0x9d, 0x27, 0x98,
	0xd2, 0x13, 0xb8, 0x52, 0xa5, 0xaa, 0x49, 0xc5, 0x17, 0x79, 0x8c, 0x5c, 0xe6, 0x32, 0x8f, 0x90,
	0x1a, 0xbf, 0x40, 0xf2, 0x06, 0x29, 0xe0, 0x50, 0x3d, 0x53, 0xb9, 0xb0, 0x2b, 0x77, 0xe2, 0x39,
	0x07, 0x1f, 0xce, 0xf6, 0x1d, 0x1c, 0xa1, 0xf7, 0x47, 0xe9, 0xf4, 0x6a, 0x74, 0xa4, 0x92, 0xe3,
	0x5e, 0x2f, 0x1e, 0xe6, 0xe3, 0xe3, 0x7c, 0x70, 0x75, 0x74, 0xfd, 0xe0, 0x78, 0x16, 0xe5, 0xc5,
	0xf1, 0x48, 0x26, 0x32, 0x8b, 0x0a, 0x39, 0x38, 0x9a, 0x66, 0x69, 0x91, 0x5a, 0xaf, 0x2d, 0x8d,
	0x8f, 0xc0, 0xf8, 0x08, 0x8c, 0x8f, 0xb4, 0x71, 0xf5, 0x83, 0x91, 0x2a, 0xc6, 0xb3, 0xde, 0x51,
	0x3f, 0x9d, 0x1c, 0x8f, 0xd2, 0x51, 0x7a, 0x6c, 0xce, 0xf4, 0x66, 0x43, 0xf3, 0x65, 0x3e, 0xcc,
	0x2f, 0xc0, 0x7a, 0xf3, 0xdf, 0xeb, 0x68, 0xc3, 0x4b, 0x07, 0xd2, 0x7a, 0x0b, 0x1d, 0xa8, 0xa4,
	0x90, 0x59, 0x12, 0xc5, 0x97, 0xc5, 0xcd, 0x54, 0x56, 0x56, 0xeb, 0xab, 0xf7, 0x77, 0xf9, 0xfe,
	0x52, 0x28, 0x6e, 0xa6, 0xd2, 0x7a, 0x84, 0xd0, 0x34, 0x4b, 0xa7, 0x32, 0x2b, 0x94, 0xcc, 0x2b,
	0x6b, 0xf5, 0xf5, 0xfb, 0x7b, 0x1f, 0x3e, 0x38, 0xfa, 0x16, 0x77, 0x8e, 0x34, 0xf6, 0x51, 0xf7,
	0xf6, 0x0c, 0x4d, 0x8a, 0xec, 0x86, 0x3f, 0x07, 0x62, 0x7d, 0x8a, 0x76, 0xfa, 0x63, 0x15, 0x0f,
	0x32, 0x99, 0x54, 0xd6, 0x0d, 0xe0, 0x1b, 0xdf, 0x09, 0xc8, 0x6f, 0x8f, 0x58, 0x2f, 0xa1, 0xcd,
	0x22, 0xbd, 0x92, 0x49, 0x65, 0xc3, 0xb8, 0x0b, 0x1f, 0x96, 0x8b, 0x0e, 0xf3, 0x22, 0xca, 0x8a,
	0xcb, 0x69, 0x9a, 0xab, 0x42, 0xa5, 0x49, 0x65, 0xb3, 0xbe, 0x7a, 0x7f, 0xef, 0xc3, 0x77, 0xbe,
	0x15, 0xba, 0x5b, 0x1a, 0xf3, 0x03, 0x73, 0x78, 0xf9, 0x69, 0xb5, 0xd1, 0xbe, 0x4c, 0x06, 0xcf,
	0xb0, 0xb6, 0xfe, 0x17, 0xac, 0x3d, 0x99, 0x0c, 0x6e, 0x91, 0x3e, 0x46, 0x9b, 0x59, 0x1a, 0xcb,
	0xbc, 0xb2, 0x5d, 0x5f, 0xbf, 0x7f, 0xf8, 0x1d, 0x91, 0xf2, 0x34, 0x96, 0x1c, 0xec, 0xab, 0x9f,
	0xa2, 0x3b, 0xff, 0x95, 0x44, 0x0b, 0xa3, 0xf5, 0x2b, 0x79, 0x53, 0x96, 0x49, 0xff, 0xd4, 0xb9,
	0xb8, 0x8e, 0xe2, 0x99, 0xac, 0xac, 0x41, 0x2e, 0xcc, 0xc7, 0xc3, 0xb5, 0x4f, 0x56, 0x1f, 0xee,
	0x3f, 0x7e, 0x72, 0x6f, 0xe5, 0xab, 0x27, 0xf7, 0x56, 0xfe, 0xf5, 0xe7, 0x7b, 0x2b, 0x6f, 0x72,
	0xb4, 0x73, 0xeb, 0xd1, 0x2b, 0x68, 0x2b, 0x1d, 0x0e, 0x73, 0x59, 0x18, 0xa0, 0x03, 0x5e, 0x7e,
	0x59, 0x16, 0xda, 0x88, 0x55, 0x02, 0x50, 0x07, 0xdc, 0xfc, 0xd6, 0x37, 0xf6, 0xd3, 0xb8, 0xb2,
	0x6e, 0x44, 0xfa, 0xe7, 0xc3, 0x9d, 0xc7, 0x25, 0xe6, 0x7b, 0x7f, 0xaa, 0xa0, 0x0d, 0xed, 0xb0,
	0x55, 0x41, 0xdb, 0xcc, 0x3b, 0x25, 0x2e, 0x73, 0xf0, 0x4a, 0x75, 0x6f, 0xbe, 0xa8, 0x6f, 0xb3,
	0xe4, 0x3a, 0x8a, 0xd5, 0xc0, 0xaa, 0x21, 0xc4, 0x1c, 0xea, 0x09, 0xd6, 0x64, 0x94, 0xe3, 0xd5,
	0xea, 0xe1, 0x7c, 0x51, 0x47, 0x6c, 0x20, 0x93, 0x42, 0x0d, 0x95, 0xcc, 0xac, 0xbb, 0x68, 0xf7,
	0x51, 0x48, 0x5c, 0xad, 0x76, 0xf0, 0x5a, 0xf5, 0x60, 0xbe, 0xa8, 0xef, 0x3e, 0x9a, 0x45, 0xb1,
	0xd6, 0x0e, 0xac, 0x2a, 0xda, 0xf1, 0xbb, 0x94, 0x13, 0xe1, 0x73, 0xbc, 0x5e, 0xdd, 0x9f, 0x2f,
	0xea, 0x3b, 0xfe, 0x54, 0xd3, 0x22, 0xcd, 0x74, 0x10, 0x0d, 0xe6, 0x11, 0x7e, 0x81, 0x37, 0xaa,
	0x68, 0xbe, 0xa8, 0x6f, 0x35, 0x54, 0x12, 0x65, 0x26, 0x21, 0xa1, 0x11, 0x6f, 0x56, 0x77, 0xe7,
	0x8b, 0xfa, 0x66, 0x68, 0xa4, 0x16, 0xda, 0x70, 0x69, 0x53, 0xe0, 0xad, 0xea, 0xce, 0x7c, 0x51,
	0xdf, 0x70, 0xe5, 0xb0, 0xd0, 0x96, 0x9c, 0xb5, 0xda, 0x02, 0x6f, 0x83, 0x25, 0x57, 0xa3, 0xb1,
	0x91, 0x32, 0xaf, 0xc9, 0xce, 0xf1, 0x0e, 0x48, 0x59, 0x32, 0x54, 0x5f, 0xea, 0x08, 0xbb, 0x7e,
	0x20, 0xb4, 0x7c, 0x17, 0x22, 0xec, 0xa6, 0x79, 0x51, 0x6a, 0x1a, 0x4c, 0x9c, 0xb1, 0x80, 0x62,
	0x04, 0x9a, 0x86, 0x2a, 0xbe, 0x50, 0xb9, 0xc9, 0x4a, 0xc3, 0xf7, 0x5d, 0x4a, 0x3c, 0xbc, 0x57,
	0x6a, 0xd2, 0x34, 0x96, 0x51, 0xa2, 0xe3, 0x0a, 0xbd, 0x80, 0xb5, 0x3c, 0xea, 0xe0, 0x7d, 0x88,
	0x2b, 0x4c, 0x72, 0x35, 0x4a, 0xe4, 0xc0, 0x7a, 0x1d, 0x21, 0xed, 0xe9, 0x65, 0xd0, 0x66, 0x4d,
	0x81, 0x0f, 0x20, 0x25, 0xda, 0xdf, 0x60, 0xac, 0x86, 0x85, 0x75, 0x0f, 0xed, 0x19, 0xa7, 0x4b,
	0xfd, 0x21, 0x64, 0xd4, 0xb8, 0x0e, 0x06, 0x87, 0x68, 0xcd, 0xe7, 0xf8, 0x4e, 0x75, 0x6b, 0xbe,
	0xa8, 0xaf, 0xf9, 0x99, 0x2e, 0xe0, 0xb9, 0xcf, 0x31, 0xae, 0x6e, 0xcf, 0x17, 0xf5, 0xf5, 0xf3,
	0xd4, 0x48, 0x88, 0xe7, 0xe0, 0x17, 0x40, 0x42, 0x12, 0x53, 0x25, 0x7a, 0xde, 0xe5, 0x34, 0x08,
	0x98, 0xef, 0x61, 0x0b, 0x30, 0xe9, 0x97, 0xd3, 0x4c, 0xe6, 0xb9, 0x6e, 0x98, 0xbb, 0x68, 0x37,
	0x10, 0x44, 0xd0, 0x0e, 0xf5, 0x04, 0x7e, 0x11, 0x5c, 0x0a, 0x8a, 0xa8, 0x90, 0x13, 0x99, 0x98,
	0x8c, 0x51, 0x5d, 0x44, 0xfc, 0x12, 0x64, 0x8c, 0x7e, 0x3e, 0x8b, 0x62, 0x7d, 0x8b, 0xe7, 0x0b,
	0xfc, 0x32, 0xdc, 0xe2, 0xa5, 0x85, 0xf5, 0x1a, 0xda, 0x75, 0x69, 0x10, 0x5c, 0x8a, 0x36, 0xf1,
	0xf0, 0x2b, 0x10, 0xb6, 0x2b, 0xf3, 0x5c, 0x8c, 0xa3, 0xc4, 0x7a, 0x1f, 0x59, 0xb7, 0xca, 0x4b,
	0x9f, 0x5f, 0x02, 0xe2, 0xff, 0x55, 0x5f, 0x9c, 0x2f, 0xea, 0x77, 0x96, 0x56, 0x7e, 0x06, 0xd8,
	0x6f, 0xa0, 0xfd, 0x16, 0xa7, 0x44, 0x50, 0x0e, 0x60, 0x95, 0xea, 0x9d, 0xf9, 0xa2, 0xbe, 0xd7,
	0xca, 0x64, 0x54, 0xc8, 0xcc, 0xe0, 0x3d, 0x40, 0x2f, 0x3f, 0x6f, 0xf2, 0x0c, 0xf2, 0xd5, 0xea,
	0x2b, 0xf3, 0x45, 0xdd, 0x7a, 0xce, 0x76, 0x89, 0x7a, 0x17, 0xed, 0x42, 0xaf, 0xda, 0xc4, 0xc5,
	0x55, 0x88, 0x12, 0x5a, 0xb5, 0x1f, 0xc5, 0xba, 0x66, 0xb6, 0xef, 0x09, 0xc2, 0xbc, 0x00, 0xbf,
	0x06, 0xce, 0xdb, 0x69, 0x52, 0x44, 0x2a, 0xc9, 0xcd, 0x49, 0xcf, 0xe6, 0x90, 0x9f, 0xbb, 0xe5,
	0xc9, 0xa4, 0x9f, 0x41, 0x7e, 0xee, 0xa2, 0x5d, 0x87, 0x2e, 0xb5, 0xaf, 0x83, 0xd6, 0x91, 0x4b,
	0x6d, 0x15, 0xed, 0x78, 0xb4, 0x45, 0x04, 0x3b, 0xa5, 0xb8, 0x06, 0xb8, 0x9e, 0x1c, 0x45, 0x85,
	0xba, 0x96, 0x5a, 0xd7, 0xf5, 0x03, 0x66, 0x74, 0xf7, 0x40, 0x07, 0x24, 0xbe, 0x96, 0x56, 0x1d,
	0xed, 0x39, 0x94, 0xd3, 0x26, 0xe5, 0xd4, 0xb3, 0x29, 0xae, 0x43, 0x0a, 0x1c, 0x99, 0xc9, 0xa1,
	0xcc, 0x64, 0xd2, 0x97, 0x3a, 0x4b, 0x82, 0x9c, 0xd0, 0x4b, 0xe2, 0x38, 0xba, 0xb4, 0xf8, 0x0d,
	0x30, 0x11, 0xd1, 0x95, 0x24, 0x83, 0x81, 0xae, 0xad, 0xa6, 0x45, 0x93, 0xb9, 0x14, 0xbf, 0x09,
	0xb4, 0x68, 0xaa, 0xd8, 0x30, 0x9e, 0x38, 0x0e, 0x7e, 0xab, 0x6c, 0x8f, 0xc1, 0xc0, 0x94, 0x3f,
	0x6c, 0x04, 0x82, 0x13, 0x5b, 0xe0, 0xb7, 0xcb, 0xf2, 0xcf, 0x7a, 0x79, 0x91, 0x45, 0x7d, 0x13,
	0x40, 0x27, 0x74, 0x05, 0xeb, 0xba, 0x17, 0xf8, 0x1d, 0x70, 0xb2, 0x33, 0x8b, 0x0b, 0x35, 0x8d,
	0x6f, 0x34, 0x49, 0x1d, 0x76, 0xca, 0x1c, 0x8a, 0xdf, 0x05, 0x92, 0x3a, 0xea, 0x5a, 0x0d, 0xa4,
	0x96, 0x77, 0x7c, 0x27, 0x74, 0x7d, 0xfc, 0xff, 0x20, 0xef, 0xa4, 0x83, 0x59, 0x9c, 0x1a, 0x9a,
	0x11, 0xfb, 0x84, 0xb4, 0x28, 0xbe, 0x5f, 0xd2, 0x2c, 0xea, 0x5f, 0x45, 0xa3, 0x32, 0x5c, 0xdb,
	0x25, 0x9c, 0x08, 0xdd, 0xa3, 0xdf, 0x5b, 0x86, 0xdb, 0x8f, 0xa3, 0x2c, 0x5a, 0x4e, 0x35, 0xd6,
	0xe9, 0xfa, 0x5c, 0xe0, 0xf7, 0x00, 0x93, 0x4d, 0xa6, 0x69, 0x66, 0xfc, 0xeb, 0x12, 0xd1, 0xf6,
	0x48, 0x87, 0xe2, 0xf7, 0xcb, 0x24, 0x46, 0xc5, 0x38, 0x89, 0x26, 0x52, 0xb7, 0x2e, 0x71, 0x19,
	0x09, 0xf0, 0xf7, 0xa1, 0x75, 0x49, 0xac, 0xa2, 0x5c, 0x9f, 0x68, 0x86, 0x9e, 0x6d, 0x2e, 0xfa,
	0x00, 0x4e, 0x34, 0x67, 0x49, 0xdf, 0xdc, 0x62, 0xa1, 0x8d, 0x86, 0xef, 0x5c, 0xe0, 0x23, 0xc8,
	0x58, 0x23, 0x1d, 0x98, 0xe1, 0x62, 0xd0, 0x8f, 0x41, 0xe6, 0x69, 0xe4, 0x2a, 0xda, 0xe1, 0xd4,
	0xa6, 0xec, 0x94, 0x72, 0xfc, 0x03, 0xc0, 0xe0, 0xb2, 0x2f, 0xd5, 0xb5, 0xcc, 0xb4, 0x8e, 0xf0,
	0x56, 0x68, 0xfa, 0xe1, 0x01, 0xe8, 0x48, 0x36, 0x9a, 0x2d, 0xc9, 0x74, 0x4a, 0xdc, 0x90, 0xe2,
	0x0f, 0xc1, 0xa3, 0x53, 0x3d, 0xcf, 0x35, 0x75, 0x08, 0x6f, 0x05, 0x97, 0x2e, 0x0b, 0x04, 0xfe,
	0xe8, 0xf6, 0x48, 0xee, 0xaa, 0xdc, 0x8c, 0xed, 0x06, 0x09, 0x28, 0xfe, 0x61, 0xe9, 0x52, 0x94,
	0x4b, 0x33, 0x77, 0x3b, 0x5d, 0xd7, 0xf4, 0x5c, 0x80, 0x7f, 0x54, 0xce, 0xdd, 0xc9, 0x34, 0x36,
	0x4d, 0x67, 0x42, 0x64, 0x5e, 0x20, 0x88, 0x6e, 0x9d, 0x1f, 0x03, 0x1e, 0x4b, 0xf2, 0x22, 0xd2,
	0x7d, 0x53, 0x41, 0xdb, 0x41, 0xd8, 0x10, 0x17, 0x5d, 0x8a, 0x3f, 0x86, 0x22, 0x04, 0xb3, 0x9e,
	0x5e, 0x0f, 0x34, 0x6a, 0x10, 0x36, 0x96, 0x15, 0xfa, 0x04, 0x50, 0x83, 0x59, 0x6f, 0x5a, 0x16,
	0x69, 0x59, 0x56, 0x8a, 0x7f, 0xf2, 0x5c, 0x59, 0x8d, 0xbc, 0xc9, 0x19, 0xf5, 0x1c, 0xfc, 0x10,
	0xe4, 0xcd, 0x4c, 0xc9, 0x64, 0xa0, 0x83, 0x3d, 0xf3, 0xb9, 0xeb, 0xe0, 0x9f, 0x42, 0xb0, 0x67,
	0x69, 0x16, 0x0f, 0xf4, 0x04, 0x63, 0x4d, 0xfc, 0x33, 0x98, 0x60, 0x6c, 0xa8, 0xdb, 0xcf, 0xf6,
	0x3d, 0x87, 0x99, 0x7a, 0x7c, 0x0a, 0xed, 0x67, 0xa7, 0xc9, 0x40, 0x2d, 0x0b, 0x22, 0xda, 0xd4,
	0xc3, 0x3f, 0x87, 0xe8, 0xc5, 0x58, 0x1a, 0x19, 0x75, 0x03, 0x8a, 0x7f, 0x01, 0x32, 0x1a, 0xe7,
	0xc6, 0x87, 0xe0, 0x8c, 0x09, 0xbb, 0x8d, 0x7f, 0x09, 0x3e, 0x04, 0x5f, 0xa8, 0xa2, 0x3f, 0xd6,
	0xb6, 0xb6, 0xce, 0x1e, 0x01, 0x5b, 0x3b, 0x82, 0xc9, 0xed, 0xd0, 0x26, 0x09, 0x5d, 0x81, 0x1b,
	0x90, 0x01, 0x47, 0x0e, 0xa3, 0x59, 0x5c, 0x68, 0x72, 0x34, 0x7d, 0x8e, 0x6d, 0x20, 0x47, 0x33,
	0xcd, 0xac, 0x77, 0xd1, 0x21, 0xf3, 0x98, 0x60, 0xc4, 0x65, 0x9f, 0x41, 0x6f, 0x3a, 0x55, 0x6b,
	0xbe, 0xa8, 0x1f, 0xb2, 0x44, 0x15, 0x2a, 0x8a, 0xd5, 0x1f, 0x6f, 0xdb, 0x33, 0xec, 0x3a, 0x44,
	0x50, 0x4c, 0xe1, 0xfe, 0x70, 0x3a, 0x88, 0x0a, 0xd3, 0x28, 0x4c, 0x94, 0x6f, 0x5c, 0xb3, 0xac,
	0x44, 0x51, 0xbe, 0x71, 0x3a, 0x3f, 0x6d, 0xcd, 0xcf, 0x56, 0x99, 0x9f, 0xb1, 0x26, 0xe8, 0xab,
	0x68, 0xc7, 0xf1, 0x2f, 0x41, 0xd1, 0x2e, 0xdd, 0x4b, 0x41, 0xf5, 0x12, 0xda, 0x6c, 0x70, 0x4a,
	0x4e, 0x30, 0x83, 0x03, 0x8d, 0x4c, 0x46, 0x57, 0xcb, 0xd1, 0xc5, 0xbc, 0x90, 0xe2, 0x5f, 0x3d,
	0x1b, 0x5d, 0x2a, 0x99, 0x49, 0x1d, 0x7e, 0xcb, 0x17, 0x3e, 0x3e, 0x81, 0xf0, 0x5b, 0x7a, 0xd7,
	0xd4, 0x28, 0xae, 0x6f, 0x9f, 0x60, 0xb7, 0x44, 0x89, 0xd3, 0xfe, 0x95, 0x96, 0x06, 0xb6, 0xdf,
	0xa5, 0xb8, 0x03, 0xd2, 0xa0, 0x9f, 0x4e, 0x4d, 0x5a, 0x39, 0x15, 0x21, 0xf7, 0xb0, 0x07, 0x61,
	0x71, 0x59, 0xcc, 0xb2, 0x44, 0x27, 0x4a, 0xf0, 0x0b, 0xec, 0x43, 0xa2, 0x04, 0x3c, 0xcc, 0x36,
	0xd1, 0xf9, 0xef, 0xc2, 0x79, 0x3b, 0xd2, 0xe9, 0xaf, 0xa0, 0xed, 0x26, 0xf3, 0x88, 0xeb, 0x5e,
	0xe0, 0x47, 0x10, 0x4b, 0x53, 0x25, 0x51, 0x1c, 0x1b, 0x7b, 0xd1, 0xe6, 0xfe, 0x19, 0xe6, 0x60,
	0x2f, 0xc6, 0x59, 0xfa, 0x85, 0xbe, 0x8f, 0x04, 0x01, 0xe5, 0x02, 0x07, 0x70, 0x1f, 0xc9, 0x73,
	0x99, 0x15, 0x50, 0x46, 0xd7, 0xc5, 0x62, 0x59, 0xc6, 0x38, 0xd6, 0xb6, 0x5a, 0x46, 0x29, 0x0e,
	0xc1, 0x56, 0x4b, 0xa5, 0x69, 0x63, 0x18, 0xab, 0xbe, 0x47, 0x5c, 0x7c, 0x0a, 0x6d, 0xbc, 0xdc,
	0x8e, 0xa2, 0xd8, 0xf0, 0xd9, 0xf7, 0xbb, 0xf8, 0xac, 0xe4, 0x73, 0x9a, 0x4e, 0xb5, 0x9f, 0xae,
	0xa9, 0x93, 0x8b, 0xcf, 0xc1, 0x4f, 0x57, 0xe9, 0x32, 0x19, 0xeb, 0xc6, 0x85, 0xa0, 0xf8, 0xa2,
	0xa4, 0xdf, 0x4d, 0x21, 0xf5, 0x2b, 0xad, 0x65, 0x97, 0x81, 0xe0, 0xcc, 0x6b, 0xe1, 0xcf, 0xe0,
	0x0a, 0xad, 0x0a, 0x8a, 0x4c, 0x25, 0x23, 0xd3, 0xd3, 0x6d, 0xa2, 0x07, 0x2a, 0xe5, 0xf8, 0xd7,
	0x65, 0x4f, 0x8f, 0x23, 0x3d, 0x50, 0x65, 0x66, 0xb6, 0x15, 0xcd, 0xf4, 0xdf, 0x94, 0xdb, 0x8a,
	0x66, 0x39, 0x46, 0xeb, 0x1d, 0xd2, 0xc5, 0xbf, 0x85, 0x84, 0x76, 0xa2, 0xa9, 0x71, 0x33, 0x74,
	0x5d, 0xfc, 0xbb, 0xd2, 0xcd, 0x19, 0x84, 0xec, 0x85, 0x9d, 0x06, 0xe5, 0xf8, 0x12, 0x42, 0xf6,
	0x66, 0x93, 0x9e, 0xcc, 0xa0, 0x4c, 0x2d, 0x7a, 0xde, 0xc5, 0xbf, 0x5f, 0x96, 0x69, 0x24, 0xbf,
	0x9c, 0x6a, 0xd4, 0x80, 0x0a, 0x1c, 0x01, 0x6a, 0x20, 0x0b, 0xc3, 0x13, 0xf0, 0xba, 0x57, 0xf2,
	0x04, 0x3c, 0xd6, 0xe5, 0x08, 0xbb, 0x2e, 0xc5, 0xfd, 0xb2, 0x1c, 0xb3, 0x69, 0x6c, 0xda, 0xc7,
	0x0c, 0x8a, 0x41, 0xc9, 0x3e, 0x3d, 0x25, 0xf4, 0x3e, 0xe0, 0xe9, 0xe2, 0xcb, 0x72, 0x1f, 0x58,
	0xae, 0xae, 0x27, 0xf4, 0x02, 0x0f, 0xe1, 0xa6, 0x13, 0x79, 0xa3, 0x73, 0xd0, 0xe5, 0xac, 0x03,
	0xcf, 0xdb, 0x08, 0x72, 0xd0, 0xcd, 0xd4, 0x04, 0xde, 0xb7, 0x1a, 0x42, 0x24, 0xd0, 0x3b, 0x92,
	0x19, 0x93, 0x63, 0xc8, 0x20, 0xc9, 0xf5, 0x96, 0x64, 0x06, 0xa5, 0xe1, 0x3d, 0x0b, 0xb0, 0x5a,
	0xf2, 0x5e, 0xe5, 0xba, 0x48, 0xb6, 0xdf, 0x31, 0x07, 0xfe, 0x00, 0x45, 0xb2, 0xd3, 0x89, 0xb1,
	0x7e, 0x1b, 0x1d, 0x38, 0xbe, 0x6d, 0x46, 0x2e, 0x90, 0xf4, 0xaa, 0xfa, 0xc2, 0x7c, 0x51, 0x3f,
	0x70, 0xd2, 0xbe, 0x99, 0xbb, 0xc0, 0xd1, 0x1a, 0x42, 0x67, 0x6d, 0x26, 0x68, 0xd0, 0x25, 0x36,
	0xc5, 0x31, 0xdc, 0x79, 0x36, 0x56, 0x85, 0xcc, 0xa7, 0x51, 0x1f, 0xa6, 0xaa, 0x67, 0xfb, 0x7a,
	0xb0, 0x0a, 0x8a, 0x27, 0xe5, 0x54, 0x4d, 0xfa, 0xa9, 0x1e, 0xac, 0x85, 0x79, 0xa4, 0x42, 0x8f,
	0x78, 0x9e, 0xaf, 0x97, 0x25, 0x07, 0x27, 0xf0, 0x48, 0x85, 0x49, 0x94, 0x24, 0xa9, 0x5e, 0x97,
	0xcc, 0xa6, 0x75, 0xca, 0x02, 0xd6, 0x60, 0x2e, 0x13, 0x17, 0x38, 0x05, 0x84, 0x53, 0x95, 0xab,
	0x9e, 0x8a, 0x55, 0x71, 0x63, 0xa2, 0x86, 0xf3, 0xda, 0xc9, 0x69, 0x19, 0x35, 0x1c, 0x2f, 0x37,
	0x31, 0xe2, 0xf9, 0xde, 0x45, 0xc7, 0x0f, 0x03, 0xfc, 0x39, 0xe4, 0x8c, 0x24, 0x69, 0x72, 0x33,
	0x49, 0x67, 0xb9, 0xbe, 0x9f, 0x7a, 0x61, 0x87, 0x96, 0x8f, 0x64, 0x06, 0xf7, 0xd3, 0x64, 0x36,
	0x91, 0xd9, 0x6d, 0x84, 0x84, 0x33, 0xd1, 0xee, 0x50, 0xc1, 0x6c, 0x9c, 0x97, 0xf8, 0x99, 0x2a,
	0xc6, 0x13, 0x59, 0xa8, 0xbe, 0xd6, 0x73, 0xea, 0x92, 0x92, 0x1a, 0x45, 0xb9, 0x5d, 0xca, 0x38,
	0x2a, 0xa9, 0x51, 0x45, 0x3b, 0xa7, 0x84, 0x33, 0xd2, 0x70, 0x29, 0x9e, 0xc1, 0x28, 0x39, 0x8d,
	0x32, 0x15, 0xf5, 0x62, 0x59, 0xdd, 0x7f, 0xfc, 0x97, 0xda, 0xca, 0x5f, 0xbf, 0xae, 0xad, 0xfc,
	0xed, 0xeb, 0xda, 0x4a, 0xa3, 0xfa, 0xf7, 0xa7, 0xb5, 0xd5, 0x7f, 0x3c, 0xad, 0xad, 0xfe, 0xf3,
	0x69, 0x6d, 0xe5, 0xab, 0x6f, 0x6a, 0x2b, 0x4f, 0xbe, 0xa9, 0xad, 0x7e, 0xb6, 0xa1, 0xff, 0xe5,
	0xf4, 0xb6, 0xcc, 0xff, 0xd0, 0x8f, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x47, 0x84, 0x0e, 0xb0,
	0x02, 0x0f, 0x00, 0x00,
}
