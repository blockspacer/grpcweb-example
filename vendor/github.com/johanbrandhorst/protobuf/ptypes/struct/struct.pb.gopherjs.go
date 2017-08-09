// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: struct/struct.proto

/*
	Package structpb is a generated protocol buffer package.

	It is generated from these files:
		struct/struct.proto

	It has these top-level messages:
		Struct
		Value
		ListValue
*/
package structpb

import js "github.com/gopherjs/gopherjs/js"
import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion1

// `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//  The JSON representation for `NullValue` is JSON `null`.
type NullValue int

const (
	// Null value.
	NullValue_NULL_VALUE NullValue = 0
)

var NullValue_name = map[int]string{
	0: "NULL_VALUE",
}
var NullValue_value = map[string]int{
	"NULL_VALUE": 0,
}

func (x NullValue) String() string {
	return NullValue_name[int(x)]
}

// `Struct` represents a structured data value, consisting of fields
// which map to dynamically typed values. In some languages, `Struct`
// might be supported by a native representation. For example, in
// scripting languages like JS a struct is represented as an
// object. The details of that representation are described together
// with the proto support for the language.
//
// The JSON representation for `Struct` is JSON object.
type Struct struct {
	*js.Object
}

// GetFields gets the Fields of the Struct.
// Unordered map of dynamically typed values.
func (m *Struct) GetFields() (x map[string]*Value) {
	if m == nil {
		return x
	}
	x = map[string]*Value{}
	mapFunc := func(value *js.Object, key *js.Object) {
		x[key.String()] = &Value{Object: value}
	}
	m.Call("getFieldsMap").Call("forEach", mapFunc)
	return x
}

// SetFields sets the Fields of the Struct.
// Unordered map of dynamically typed values.
func (m *Struct) SetFields(v map[string]*Value) {
	m.Call("clearFieldsMap")
	mp := m.Call("getFieldsMap")
	for key, value := range v {
		mp.Call("set", key, value)
	}
}

// ClearFields clears the Fields of the Struct.
// Unordered map of dynamically typed values.
func (m *Struct) ClearFields() {
	m.Call("clearFieldsMap")
}

// New creates a new Struct.
// Unordered map of dynamically typed values.
func (m *Struct) New(fields map[string]*Value) *Struct {
	m = &Struct{
		Object: js.Global.Get("proto").Get("google").Get("protobuf").Get("Struct").New([]interface{}{
			js.Undefined,
		}),
	}

	mp := m.Call("getFieldsMap")
	for key, value := range fields {
		mp.Call("set", key, value)
	}

	return m
}

// Serialize marshals Struct to a slice of bytes.
func (m *Struct) Serialize() []byte {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a Struct from a slice of bytes.
func (m *Struct) Deserialize(rawBytes []byte) (*Struct, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("google").Get("protobuf").Get("Struct"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Struct{
		Object: obj,
	}, nil
}

// `Value` represents a dynamically typed value which can be either
// null, a number, a string, a boolean, a recursive struct value, or a
// list of values. A producer of value is expected to set one of that
// variants, absence of any variant indicates an error.
//
// The JSON representation for `Value` is JSON value.
type Value struct {
	*js.Object
}

// The kind of value.
//
// Types that are valid to be assigned to Kind:
//	*Value_NullValue
//	*Value_NumberValue
//	*Value_StringValue
//	*Value_BoolValue
//	*Value_StructValue
//	*Value_ListValue
type isValue_Kind interface {
	isValue_Kind()
}

type Value_NullValue struct {
	// Represents a null value.
	NullValue NullValue
}
type Value_NumberValue struct {
	// Represents a double value.
	NumberValue float64
}
type Value_StringValue struct {
	// Represents a string value.
	StringValue string
}
type Value_BoolValue struct {
	// Represents a boolean value.
	BoolValue bool
}
type Value_StructValue struct {
	// Represents a structured value.
	StructValue *Struct
}
type Value_ListValue struct {
	// Represents a repeated `Value`.
	ListValue *ListValue
}

func (*Value_NullValue) isValue_Kind()   {}
func (*Value_NumberValue) isValue_Kind() {}
func (*Value_StringValue) isValue_Kind() {}
func (*Value_BoolValue) isValue_Kind()   {}
func (*Value_StructValue) isValue_Kind() {}
func (*Value_ListValue) isValue_Kind()   {}

// GetKind gets the Kind of the Value.
func (m *Value) GetKind() (x isValue_Kind) {
	switch m.Call("getKindCase").Int() {
	case 1:
		x = &Value_NullValue{
			NullValue: m.GetNullValue(),
		}
	case 2:
		x = &Value_NumberValue{
			NumberValue: m.GetNumberValue(),
		}
	case 3:
		x = &Value_StringValue{
			StringValue: m.GetStringValue(),
		}
	case 4:
		x = &Value_BoolValue{
			BoolValue: m.GetBoolValue(),
		}
	case 5:
		x = &Value_StructValue{
			StructValue: m.GetStructValue(),
		}
	case 6:
		x = &Value_ListValue{
			ListValue: m.GetListValue(),
		}
	}

	return x
}

// SetKind sets the Kind of theValue.
// If the input is nil, SetKind does nothing.
func (m *Value) SetKind(kind isValue_Kind) {
	switch x := kind.(type) {
	case *Value_NullValue:
		m.SetNullValue(x.NullValue)
	case *Value_NumberValue:
		m.SetNumberValue(x.NumberValue)
	case *Value_StringValue:
		m.SetStringValue(x.StringValue)
	case *Value_BoolValue:
		m.SetBoolValue(x.BoolValue)
	case *Value_StructValue:
		m.SetStructValue(x.StructValue)
	case *Value_ListValue:
		m.SetListValue(x.ListValue)
	}
}

// GetNullValue gets the NullValue of the Value.
// Represents a null value.
func (m *Value) GetNullValue() (x NullValue) {
	if m == nil {
		return x
	}
	return NullValue(m.Call("getNullValue").Int())
}

// SetNullValue sets the NullValue of the Value.
// Represents a null value.
func (m *Value) SetNullValue(v NullValue) {
	m.Call("setNullValue", v)
}

// HasNullValue indicates whether the NullValue of the Value is set.
// Represents a null value.
func (m *Value) HasNullValue() bool {
	if m == nil {
		return false
	}
	return m.Call("hasNullValue").Bool()
}

// ClearNullValue clears the NullValue of the Value.
// Represents a null value.
func (m *Value) ClearNullValue() {
	m.Call("clearNullValue")
}

// GetNumberValue gets the NumberValue of the Value.
// Represents a double value.
func (m *Value) GetNumberValue() (x float64) {
	if m == nil {
		return x
	}
	return m.Call("getNumberValue").Float()
}

// SetNumberValue sets the NumberValue of the Value.
// Represents a double value.
func (m *Value) SetNumberValue(v float64) {
	m.Call("setNumberValue", v)
}

// HasNumberValue indicates whether the NumberValue of the Value is set.
// Represents a double value.
func (m *Value) HasNumberValue() bool {
	if m == nil {
		return false
	}
	return m.Call("hasNumberValue").Bool()
}

// ClearNumberValue clears the NumberValue of the Value.
// Represents a double value.
func (m *Value) ClearNumberValue() {
	m.Call("clearNumberValue")
}

// GetStringValue gets the StringValue of the Value.
// Represents a string value.
func (m *Value) GetStringValue() (x string) {
	if m == nil {
		return x
	}
	return m.Call("getStringValue").String()
}

// SetStringValue sets the StringValue of the Value.
// Represents a string value.
func (m *Value) SetStringValue(v string) {
	m.Call("setStringValue", v)
}

// HasStringValue indicates whether the StringValue of the Value is set.
// Represents a string value.
func (m *Value) HasStringValue() bool {
	if m == nil {
		return false
	}
	return m.Call("hasStringValue").Bool()
}

// ClearStringValue clears the StringValue of the Value.
// Represents a string value.
func (m *Value) ClearStringValue() {
	m.Call("clearStringValue")
}

// GetBoolValue gets the BoolValue of the Value.
// Represents a boolean value.
func (m *Value) GetBoolValue() (x bool) {
	if m == nil {
		return x
	}
	return m.Call("getBoolValue").Bool()
}

// SetBoolValue sets the BoolValue of the Value.
// Represents a boolean value.
func (m *Value) SetBoolValue(v bool) {
	m.Call("setBoolValue", v)
}

// HasBoolValue indicates whether the BoolValue of the Value is set.
// Represents a boolean value.
func (m *Value) HasBoolValue() bool {
	if m == nil {
		return false
	}
	return m.Call("hasBoolValue").Bool()
}

// ClearBoolValue clears the BoolValue of the Value.
// Represents a boolean value.
func (m *Value) ClearBoolValue() {
	m.Call("clearBoolValue")
}

// GetStructValue gets the StructValue of the Value.
// Represents a structured value.
func (m *Value) GetStructValue() (x *Struct) {
	if m == nil {
		return x
	}
	return &Struct{Object: m.Call("getStructValue")}
}

// SetStructValue sets the StructValue of the Value.
// Represents a structured value.
func (m *Value) SetStructValue(v *Struct) {
	if v != nil {
		m.Call("setStructValue", v)
	} else {
		m.ClearStructValue()
	}
}

// HasStructValue indicates whether the StructValue of the Value is set.
// Represents a structured value.
func (m *Value) HasStructValue() bool {
	if m == nil {
		return false
	}
	return m.Call("hasStructValue").Bool()
}

// ClearStructValue clears the StructValue of the Value.
// Represents a structured value.
func (m *Value) ClearStructValue() {
	m.Call("clearStructValue")
}

// GetListValue gets the ListValue of the Value.
// Represents a repeated `Value`.
func (m *Value) GetListValue() (x *ListValue) {
	if m == nil {
		return x
	}
	return &ListValue{Object: m.Call("getListValue")}
}

// SetListValue sets the ListValue of the Value.
// Represents a repeated `Value`.
func (m *Value) SetListValue(v *ListValue) {
	if v != nil {
		m.Call("setListValue", v)
	} else {
		m.ClearListValue()
	}
}

// HasListValue indicates whether the ListValue of the Value is set.
// Represents a repeated `Value`.
func (m *Value) HasListValue() bool {
	if m == nil {
		return false
	}
	return m.Call("hasListValue").Bool()
}

// ClearListValue clears the ListValue of the Value.
// Represents a repeated `Value`.
func (m *Value) ClearListValue() {
	m.Call("clearListValue")
}

// New creates a new Value.
// Represents a null value.
func (m *Value) New(kind isValue_Kind) *Value {
	m = &Value{
		Object: js.Global.Get("proto").Get("google").Get("protobuf").Get("Value").New([]interface{}{
			js.Undefined,
			js.Undefined,
			js.Undefined,
			js.Undefined,
			js.Undefined,
			js.Undefined,
		}),
	}

	m.SetKind(kind)

	return m
}

// Serialize marshals Value to a slice of bytes.
func (m *Value) Serialize() []byte {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a Value from a slice of bytes.
func (m *Value) Deserialize(rawBytes []byte) (*Value, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("google").Get("protobuf").Get("Value"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Value{
		Object: obj,
	}, nil
}

// `ListValue` is a wrapper around a repeated field of values.
//
// The JSON representation for `ListValue` is JSON array.
type ListValue struct {
	*js.Object
}

// GetValues gets the Values of the ListValue.
// Repeated field of dynamically typed values.
// Warning: mutating the returned slice will not be reflected in the message.
// Use the setter to make changes to the slice in the message.
func (m *ListValue) GetValues() (x []*Value) {
	if m == nil {
		return x
	}
	arrFunc := func(value *js.Object) {
		x = append(x, &Value{Object: value})
	}
	m.Call("getValuesList").Call("forEach", arrFunc)
	return x
}

// SetValues sets the Values of the ListValue.
// Repeated field of dynamically typed values.
func (m *ListValue) SetValues(v []*Value) {
	arr := js.Global.Get("Array").New(len(v))
	for i, value := range v {
		arr.SetIndex(i, value)
	}
	m.Call("setValuesList", arr)
}

// AddValues adds an entry to the Values slice of the ListValue
// at the specified index. If index is negative, inserts the element
// at the index counted from the end of the slice, with origin 1.
// Repeated field of dynamically typed values.
func (m *ListValue) AddValues(v *Value, index int) {
	m.Call("addValues", v, index)
}

// ClearValues clears the Values of the ListValue.
// Repeated field of dynamically typed values.
func (m *ListValue) ClearValues() {
	m.Call("clearValuesList")
}

// New creates a new ListValue.
// Repeated field of dynamically typed values.
func (m *ListValue) New(values []*Value) *ListValue {
	m = &ListValue{
		Object: js.Global.Get("proto").Get("google").Get("protobuf").Get("ListValue").New([]interface{}{
			js.Undefined,
		}),
	}

	arr := js.Global.Get("Array").New(len(values))
	for i, value := range values {
		arr.SetIndex(i, value)
	}
	m.Call("setValuesList", arr)

	return m
}

// Serialize marshals ListValue to a slice of bytes.
func (m *ListValue) Serialize() []byte {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a ListValue from a slice of bytes.
func (m *ListValue) Deserialize(rawBytes []byte) (*ListValue, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("google").Get("protobuf").Get("ListValue"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &ListValue{
		Object: obj,
	}, nil
}