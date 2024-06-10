// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PackOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsPackOptions(buf []byte, offset flatbuffers.UOffsetT) *PackOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PackOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPackOptions(buf []byte, offset flatbuffers.UOffsetT) *PackOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PackOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PackOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PackOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PackOptions) ValuesCount() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PackOptions) MutateValuesCount(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *PackOptions) Axis() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PackOptions) MutateAxis(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func PackOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PackOptionsAddValuesCount(builder *flatbuffers.Builder, valuesCount int32) {
	builder.PrependInt32Slot(0, valuesCount, 0)
}
func PackOptionsAddAxis(builder *flatbuffers.Builder, axis int32) {
	builder.PrependInt32Slot(1, axis, 0)
}
func PackOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
