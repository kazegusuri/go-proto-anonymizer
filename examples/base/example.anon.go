// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/base/example.proto

/*
Package examples is a generated protocol buffer package.

It is generated from these files:
	examples/base/example.proto

It has these top-level messages:
	SimpleMessage
	NumberMessage
	RepeatedNumberMessage
	NestedMessage
	EnumMessage
	Oneof
	OneofMessage
	MapMessage
	WellKnownTypeMessage
*/
package examples

import github_com_kazegusuri_go_proto_anonymizer "github.com/kazegusuri/go-proto-anonymizer"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/kazegusuri/go-proto-anonymizer"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *SimpleMessage) Anonymize() {
	if m == nil {
		return
	}
	_ = m.StringValue
	_ = m.BoolValue
}

func (m *NumberMessage) Anonymize() {
	if m == nil {
		return
	}
	_ = m.FloatValue
	_ = m.DoubleValue
	_ = m.Int32Value
	_ = m.Int64Value
	_ = m.Uint32Value
	_ = m.Uint64Value
	_ = m.Sint32Value
	_ = m.Sint64Value
	_ = m.Fixed32Value
	_ = m.Fixed64Value
	_ = m.Sfixed32Value
	_ = m.Sfixed64Value
}

func (m *RepeatedNumberMessage) Anonymize() {
	if m == nil {
		return
	}
	for i := range m.FloatValues {
		_ = m.FloatValues[i]
	}
	for i := range m.DoubleValues {
		_ = m.DoubleValues[i]
	}
	for i := range m.Int32Values {
		_ = m.Int32Values[i]
	}
	for i := range m.Int64Values {
		_ = m.Int64Values[i]
	}
	for i := range m.Uint32Values {
		_ = m.Uint32Values[i]
	}
	for i := range m.Uint64Values {
		_ = m.Uint64Values[i]
	}
	for i := range m.Sint32Values {
		_ = m.Sint32Values[i]
	}
	for i := range m.Sint64Values {
		_ = m.Sint64Values[i]
	}
	for i := range m.Fixed32Values {
		_ = m.Fixed32Values[i]
	}
	for i := range m.Fixed64Values {
		_ = m.Fixed64Values[i]
	}
	for i := range m.Sfixed32Values {
		_ = m.Sfixed32Values[i]
	}
	for i := range m.Sfixed64Values {
		_ = m.Sfixed64Values[i]
	}
}

func (m *NestedMessage) Anonymize() {
	if m == nil {
		return
	}
	github_com_kazegusuri_go_proto_anonymizer.Anonymize(m.NestedValue)
	for i := range m.RepeatedNestedValues {
		github_com_kazegusuri_go_proto_anonymizer.Anonymize(m.RepeatedNestedValues[i])
	}
}

func (m *NestedMessage_Nested) Anonymize() {
	if m == nil {
		return
	}
	_ = m.Int32Value
	_ = m.StringValue
}

func (m *EnumMessage) Anonymize() {
	if m == nil {
		return
	}
	_ = m.NumericEnumValue
	for i := range m.RepeatedNumericEnumValues {
		_ = m.RepeatedNumericEnumValues[i]
	}
	_ = m.AliasedEnumValue
	_ = m.NestedEnumValue
	for i := range m.RepeatedNestedEnumValues {
		_ = m.RepeatedNestedEnumValues[i]
	}
}

func (m *Oneof) Anonymize() {
	if m == nil {
		return
	}
	if ov, ok := m.GetOneofValue().(*Oneof_Int32Value); ok {
		_ = ov.Int32Value
	}
	if ov, ok := m.GetOneofValue().(*Oneof_StringValue); ok {
		_ = ov.StringValue
	}
}

func (m *OneofMessage) Anonymize() {
	if m == nil {
		return
	}
	if ov, ok := m.GetOneofValue().(*OneofMessage_Int32Value); ok {
		_ = ov.Int32Value
	}
	if ov, ok := m.GetOneofValue().(*OneofMessage_StringValue); ok {
		_ = ov.StringValue
	}
	for i := range m.RepeatedOneofValues {
		github_com_kazegusuri_go_proto_anonymizer.Anonymize(m.RepeatedOneofValues[i])
	}
}

func (m *MapMessage) Anonymize() {
	if m == nil {
		return
	}
	for mk := range m.MappedValue {
		_ = m.MappedValue[mk]
	}
	for mk := range m.MappedEnumValue {
		_ = m.MappedEnumValue[mk]
	}
	for mk := range m.MappedNestedValue {
		github_com_kazegusuri_go_proto_anonymizer.Anonymize(m.MappedNestedValue[mk])
	}
}

func (m *WellKnownTypeMessage) Anonymize() {
	if m == nil {
		return
	}
}
