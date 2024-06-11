// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite_metadata

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NormalizationOptionsT struct {
	Mean []float32
	Std  []float32
}

func (t *NormalizationOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	meanOffset := flatbuffers.UOffsetT(0)
	if t.Mean != nil {
		meanLength := len(t.Mean)
		NormalizationOptionsStartMeanVector(builder, meanLength)
		for j := meanLength - 1; j >= 0; j-- {
			builder.PrependFloat32(t.Mean[j])
		}
		meanOffset = builder.EndVector(meanLength)
	}
	stdOffset := flatbuffers.UOffsetT(0)
	if t.Std != nil {
		stdLength := len(t.Std)
		NormalizationOptionsStartStdVector(builder, stdLength)
		for j := stdLength - 1; j >= 0; j-- {
			builder.PrependFloat32(t.Std[j])
		}
		stdOffset = builder.EndVector(stdLength)
	}
	NormalizationOptionsStart(builder)
	NormalizationOptionsAddMean(builder, meanOffset)
	NormalizationOptionsAddStd(builder, stdOffset)
	return NormalizationOptionsEnd(builder)
}

func (rcv *NormalizationOptions) UnPackTo(t *NormalizationOptionsT) {
	meanLength := rcv.MeanLength()
	t.Mean = make([]float32, meanLength)
	for j := 0; j < meanLength; j++ {
		t.Mean[j] = rcv.Mean(j)
	}
	stdLength := rcv.StdLength()
	t.Std = make([]float32, stdLength)
	for j := 0; j < stdLength; j++ {
		t.Std[j] = rcv.Std(j)
	}
}

func (rcv *NormalizationOptions) UnPack() *NormalizationOptionsT {
	if rcv == nil {
		return nil
	}
	t := &NormalizationOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type NormalizationOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsNormalizationOptions(buf []byte, offset flatbuffers.UOffsetT) *NormalizationOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NormalizationOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNormalizationOptions(buf []byte, offset flatbuffers.UOffsetT) *NormalizationOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NormalizationOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NormalizationOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NormalizationOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NormalizationOptions) Mean(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *NormalizationOptions) MeanLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *NormalizationOptions) MutateMean(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *NormalizationOptions) Std(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *NormalizationOptions) StdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *NormalizationOptions) MutateStd(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func NormalizationOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func NormalizationOptionsAddMean(builder *flatbuffers.Builder, mean flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(mean), 0)
}
func NormalizationOptionsStartMeanVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func NormalizationOptionsAddStd(builder *flatbuffers.Builder, std flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(std), 0)
}
func NormalizationOptionsStartStdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func NormalizationOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}