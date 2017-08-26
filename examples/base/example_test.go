package examples

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func compareMsg(t *testing.T, m1, m2 proto.Message) {
	if !reflect.DeepEqual(m1, m2) {
		t.Fatalf("not equal\nmsg: %#v\nexp:%#v", m1, m2)
	}
}

func TestSimple(t *testing.T) {
	msg := &SimpleMessage{
		StringValue: "xyz",
		BoolValue:   true,
	}
	msg.Anonymize()
	compareMsg(t, msg, msg)
}

func TestNumberMessage(t *testing.T) {
	msg := &NumberMessage{
		FloatValue:    0.5,
		DoubleValue:   2.2,
		Int32Value:    -3,
		Int64Value:    -4,
		Uint32Value:   5,
		Uint64Value:   6,
		Sint32Value:   -7,
		Sint64Value:   -8,
		Fixed32Value:  9,
		Fixed64Value:  10,
		Sfixed32Value: -11,
		Sfixed64Value: -12,
	}
	msg.Anonymize()
	compareMsg(t, msg, msg)
}

func TestRepeatedNumberMessage(t *testing.T) {
	msg := &RepeatedNumberMessage{
		FloatValues:    []float32{0.5, 1},
		DoubleValues:   []float64{2.2, 1},
		Int32Values:    []int32{-3, 3},
		Int64Values:    []int64{-4, 4},
		Uint32Values:   []uint32{5, 55},
		Uint64Values:   []uint64{6, 66},
		Sint32Values:   []int32{-7, 7},
		Sint64Values:   []int64{-8, 8},
		Fixed32Values:  []uint32{9, 99},
		Fixed64Values:  []uint64{10, 100},
		Sfixed32Values: []int32{-11, 11},
		Sfixed64Values: []int64{-12, 12},
	}
	msg.Anonymize()
	compareMsg(t, msg, msg)
}

func TestNestedMessage(t *testing.T) {
	msg := &NestedMessage{
		NestedValue: &NestedMessage_Nested{
			Int32Value:  100,
			StringValue: "xxx",
		},
		RepeatedNestedValues: []*NestedMessage_Nested{
			{
				Int32Value:  200,
				StringValue: "yyy",
			},
			{
				Int32Value:  300,
				StringValue: "zzz",
			},
		},
	}
	msg.Anonymize()
	compareMsg(t, msg, msg)
}

func TestEnumMessage(t *testing.T) {
	msg := &EnumMessage{
		NumericEnumValue: NumericEnum_ONE,
		RepeatedNumericEnumValues: []NumericEnum{
			NumericEnum_ONE,
			NumericEnum_TWO,
		},
		AliasedEnumValue: AliasedEnum_RUNNING,
		NestedEnumValue:  EnumMessage_PENDING,
		RepeatedNestedEnumValues: []EnumMessage_Nested{
			EnumMessage_PENDING,
			EnumMessage_COMPLETED,
		},
	}
	msg.Anonymize()
	compareMsg(t, msg, msg)
}

func TestOneofMessage(t *testing.T) {
	msg := &OneofMessage{
		OneofValue: &OneofMessage_Int32Value{Int32Value: 1000},
		RepeatedOneofValues: []*Oneof{
			{
				OneofValue: &Oneof_Int32Value{Int32Value: 1000},
			},
			{
				OneofValue: &Oneof_StringValue{StringValue: "xyz"},
			},
		},
	}
	msg.Anonymize()
	compareMsg(t, msg, msg)
}

func TestMapMessage(t *testing.T) {
	msg := &MapMessage{
		MappedValue: map[int32]string{
			1: "foo",
			2: "bar",
		},
		MappedEnumValue: map[string]NumericEnum{
			"one": NumericEnum_ONE,
			"two": NumericEnum_TWO,
		},
		MappedNestedValue: map[string]*NestedMessage{
			"foo": &NestedMessage{
				NestedValue: &NestedMessage_Nested{
					Int32Value:  100,
					StringValue: "xxx",
				},
				RepeatedNestedValues: []*NestedMessage_Nested{
					{
						Int32Value:  200,
						StringValue: "yyy",
					},
					{
						Int32Value:  300,
						StringValue: "zzz",
					},
				},
			},
		},
	}
	msg.Anonymize()
	compareMsg(t, msg, msg)
}

func TestWellKnownTypeMessage(t *testing.T) {
	d := ptypes.DurationProto(10 * time.Minute)
	ts, _ := ptypes.TimestampProto(time.Unix(1502533013, 125892275))
	msg := &WellKnownTypeMessage{
		Duration:  d,
		Timestamp: ts,
	}
	msg.Anonymize()
	compareMsg(t, msg, msg)
}
