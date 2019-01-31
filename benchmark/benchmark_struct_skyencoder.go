// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package benchmark

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher/encoder"
)

// EncodeSizeBenchmarkStruct computes the size of an encoded object of type BenchmarkStruct
func EncodeSizeBenchmarkStruct(obj *BenchmarkStruct) int {
	i0 := 0

	// obj.Int64
	i0 += 8

	// obj.String
	i0 += 4 + len(obj.String)

	// obj.StringSlice
	i0 += 4
	for _, x := range obj.StringSlice {
		i1 := 0

		// x
		i1 += 4 + len(x)

		i0 += i1
	}

	// obj.StaticStructArray
	{
		i1 := 0

		// x.A
		i1++

		// x.B
		i1 += 8

		i0 += 3 * i1
	}

	// obj.DynamicStructSlice
	i0 += 4
	for _, x := range obj.DynamicStructSlice {
		i1 := 0

		// x.C
		i1 += 4 + len(x.C)

		i0 += i1
	}

	// obj.ByteArray
	i0 += 3

	// obj.ByteSlice
	i0 += 4 + len(obj.ByteSlice)

	return i0
}

// EncodeBenchmarkStruct encodes an object of type BenchmarkStruct to the buffer in encoder.Encoder
func EncodeBenchmarkStruct(e *encoder.Encoder, obj *BenchmarkStruct) error {

	// obj.Int64
	e.Int64(obj.Int64)

	// obj.String length check
	if len(obj.String) > math.MaxUint32 {
		return errors.New("obj.String length exceeds math.MaxUint32")
	}

	// obj.String
	e.ByteSlice([]byte(obj.String))

	// obj.StringSlice length check
	if len(obj.StringSlice) > math.MaxUint32 {
		return errors.New("obj.StringSlice length exceeds math.MaxUint32")
	}

	// obj.StringSlice length
	e.Uint32(uint32(len(obj.StringSlice)))

	// obj.StringSlice
	for _, x := range obj.StringSlice {

		// x length check
		if len(x) > math.MaxUint32 {
			return errors.New("x length exceeds math.MaxUint32")
		}

		// x
		e.ByteSlice([]byte(x))

	}

	// obj.StaticStructArray
	for _, x := range obj.StaticStructArray {

		// x.A
		e.Uint8(x.A)

		// x.B
		e.Uint64(x.B)

	}

	// obj.DynamicStructSlice length check
	if len(obj.DynamicStructSlice) > math.MaxUint32 {
		return errors.New("obj.DynamicStructSlice length exceeds math.MaxUint32")
	}

	// obj.DynamicStructSlice length
	e.Uint32(uint32(len(obj.DynamicStructSlice)))

	// obj.DynamicStructSlice
	for _, x := range obj.DynamicStructSlice {

		// x.C length check
		if len(x.C) > math.MaxUint32 {
			return errors.New("x.C length exceeds math.MaxUint32")
		}

		// x.C
		e.ByteSlice([]byte(x.C))

	}

	// obj.ByteArray
	e.CopyBytes(obj.ByteArray[:])

	// obj.ByteSlice length check
	if len(obj.ByteSlice) > math.MaxUint32 {
		return errors.New("obj.ByteSlice length exceeds math.MaxUint32")
	}

	// obj.ByteSlice length
	e.Uint32(uint32(len(obj.ByteSlice)))

	// obj.ByteSlice copy
	e.CopyBytes(obj.ByteSlice)

	return nil
}

// DecodeBenchmarkStruct decodes an object of type BenchmarkStruct from the buffer in encoder.Decoder
func DecodeBenchmarkStruct(d *encoder.Decoder, obj *BenchmarkStruct) error {
	{
		// obj.Int64
		i, err := d.Int64()
		if err != nil {
			return err
		}
		obj.Int64 = i
	}

	{
		// obj.String

		if len(d.Buffer) < 4 {
			return encoder.ErrBufferUnderflow
		}

		ul, err := d.Uint32()
		if err != nil {
			return err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return encoder.ErrBufferUnderflow
		}

		obj.String = string(d.Buffer[:length])
		d.Buffer = d.Buffer[length:]
	}

	{
		// obj.StringSlice

		if len(d.Buffer) < 4 {
			return encoder.ErrBufferUnderflow
		}

		ul, err := d.Uint32()
		if err != nil {
			return err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return encoder.ErrBufferUnderflow
		}

		obj.StringSlice = make([]string, length)

		for z1 := range obj.StringSlice {
			{
				// obj.StringSlice[z1]

				if len(d.Buffer) < 4 {
					return encoder.ErrBufferUnderflow
				}

				ul, err := d.Uint32()
				if err != nil {
					return err
				}

				length := int(ul)
				if length < 0 || length > len(d.Buffer) {
					return encoder.ErrBufferUnderflow
				}

				obj.StringSlice[z1] = string(d.Buffer[:length])
				d.Buffer = d.Buffer[length:]
			}
		}

	}

	{
		// obj.StaticStructArray
		for z1 := range obj.StaticStructArray {
			{
				// obj.StaticStructArray[z1].A
				i, err := d.Uint8()
				if err != nil {
					return err
				}
				obj.StaticStructArray[z1].A = i
			}

			{
				// obj.StaticStructArray[z1].B
				i, err := d.Uint64()
				if err != nil {
					return err
				}
				obj.StaticStructArray[z1].B = i
			}

		}
	}

	{
		// obj.DynamicStructSlice

		if len(d.Buffer) < 4 {
			return encoder.ErrBufferUnderflow
		}

		ul, err := d.Uint32()
		if err != nil {
			return err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return encoder.ErrBufferUnderflow
		}

		obj.DynamicStructSlice = make([]DynamicStruct, length)

		for z1 := range obj.DynamicStructSlice {
			{
				// obj.DynamicStructSlice[z1].C

				if len(d.Buffer) < 4 {
					return encoder.ErrBufferUnderflow
				}

				ul, err := d.Uint32()
				if err != nil {
					return err
				}

				length := int(ul)
				if length < 0 || length > len(d.Buffer) {
					return encoder.ErrBufferUnderflow
				}

				obj.DynamicStructSlice[z1].C = string(d.Buffer[:length])
				d.Buffer = d.Buffer[length:]
			}
		}

	}

	{
		// obj.ByteArray
		if len(d.Buffer) < len(obj.ByteArray) {
			return encoder.ErrBufferUnderflow
		}
		copy(obj.ByteArray[:], d.Buffer[:len(obj.ByteArray)])
		d.Buffer = d.Buffer[len(obj.ByteArray):]
	}

	{
		// obj.ByteSlice

		if len(d.Buffer) < 4 {
			return encoder.ErrBufferUnderflow
		}

		ul, err := d.Uint32()
		if err != nil {
			return err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return encoder.ErrBufferUnderflow
		}

		obj.ByteSlice = make([]byte, length)
		copy(obj.ByteSlice[:], d.Buffer[:length])
		d.Buffer = d.Buffer[length:]
	}

	if len(d.Buffer) != 0 {
		return encoder.ErrRemainingBytes
	}

	return nil
}
