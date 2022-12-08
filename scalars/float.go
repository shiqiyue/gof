package scalars

import (
	"encoding/binary"
	"math"
)

func GetFloat64Ref(i float64) *float64 {
	return &i
}

func GetFloat64Value(i *float64) float64 {
	return *i
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

func Float64IfNilDefault(boolean *float64, def float64) float64 {
	if boolean == nil {
		return def
	}
	return *boolean
}

func Float32IfNilDefault(boolean *float32, def float32) float32 {
	if boolean == nil {
		return def
	}
	return *boolean
}
