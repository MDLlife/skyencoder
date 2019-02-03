// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package tests

import (
	"bytes"
	"fmt"
	mathrand "math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/cipher/encoder/encodertest"
)

func newEmptyMaxLenStringStruct1ForEncodeTest() *MaxLenStringStruct1 {
	var obj MaxLenStringStruct1
	return &obj
}

func newRandomMaxLenStringStruct1ForEncodeTest(t *testing.T, rand *mathrand.Rand) *MaxLenStringStruct1 {
	var obj MaxLenStringStruct1
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen: 4,
		MinRandLen: 1,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func newRandomZeroLenMaxLenStringStruct1ForEncodeTest(t *testing.T, rand *mathrand.Rand) *MaxLenStringStruct1 {
	var obj MaxLenStringStruct1
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen:    0,
		MinRandLen:    0,
		EmptySliceNil: false,
		EmptyMapNil:   false,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func newRandomZeroLenNilMaxLenStringStruct1ForEncodeTest(t *testing.T, rand *mathrand.Rand) *MaxLenStringStruct1 {
	var obj MaxLenStringStruct1
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen:    0,
		MinRandLen:    0,
		EmptySliceNil: true,
		EmptyMapNil:   true,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func testSkyencoderMaxLenStringStruct1(t *testing.T, obj *MaxLenStringStruct1) {
	// EncodeSize

	n1, err := encoder.Size(obj)
	if err != nil {
		t.Fatalf("encoder.Size failed: %v", err)
	}

	n2 := EncodeSizeMaxLenStringStruct1(obj)

	if n1 != n2 {
		t.Fatalf("encoder.Size() != EncodeSizeMaxLenStringStruct1() (%d != %d)", n1, n2)
	}

	// Encode

	data1 := encoder.Serialize(obj)

	data2 := make([]byte, n2)
	err = EncodeMaxLenStringStruct1(data2, obj)
	if err != nil {
		t.Fatalf("EncodeMaxLenStringStruct1 failed: %v", err)
	}

	if len(data1) != len(data2) {
		t.Fatalf("len(encoder.Serialize()) != len(EncodeMaxLenStringStruct1()) (%d != %d)", len(data1), len(data2))
	}

	if !bytes.Equal(data1, data2) {
		t.Fatal("encoder.Serialize() != EncodeMaxLenStringStruct1()")
	}

	// Decode

	var obj2 MaxLenStringStruct1
	err = encoder.DeserializeRaw(data1, &obj2)
	if err != nil {
		t.Fatalf("encoder.DeserializeRaw failed: %v", err)
	}

	if !cmp.Equal(*obj, obj2, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw result wrong")
	}

	var obj3 MaxLenStringStruct1
	err = DecodeMaxLenStringStruct1(data2, &obj3)
	if err != nil {
		t.Fatalf("DecodeMaxLenStringStruct1 failed: %v", err)
	}

	if !cmp.Equal(obj2, obj3, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != DecodeMaxLenStringStruct1()")
	}
}

func TestSkyencoderMaxLenStringStruct1(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))

	type testCase struct {
		name string
		obj  *MaxLenStringStruct1
	}

	cases := []testCase{
		{
			name: "empty object",
			obj:  newEmptyMaxLenStringStruct1ForEncodeTest(),
		},
	}

	nRandom := 10

	for i := 0; i < nRandom; i++ {
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d", i),
			obj:  newRandomMaxLenStringStruct1ForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents", i),
			obj:  newRandomZeroLenMaxLenStringStruct1ForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents set to nil", i),
			obj:  newRandomZeroLenNilMaxLenStringStruct1ForEncodeTest(t, rand),
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testSkyencoderMaxLenStringStruct1(t, tc.obj)
		})
	}
}

func decodeMaxLenStringStruct1ExpectError(t *testing.T, buf []byte, expectedErr error) {
	var obj MaxLenStringStruct1
	err := DecodeMaxLenStringStruct1(buf, &obj)

	if err == nil {
		t.Fatal("DecodeMaxLenStringStruct1: expected error, got nil")
	}

	if err != expectedErr {
		t.Fatalf("DecodeMaxLenStringStruct1: expected error %q, got %q", expectedErr, err)
	}
}

func testSkyencoderMaxLenStringStruct1DecodeErrors(t *testing.T, k int, obj *MaxLenStringStruct1) {
	isEncodableField := func(f reflect.StructField) bool {
		// Skip unexported fields
		if f.PkgPath != "" {
			return false
		}

		// Skip fields disabled with and enc:"- struct tag
		tag := f.Tag.Get("enc")
		return !strings.HasPrefix(tag, "-,") && tag != "-"
	}

	numEncodableFields := func(obj interface{}) int {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()

			n := 0
			for i := 0; i < v.NumField(); i++ {
				f := t.Field(i)
				if !isEncodableField(f) {
					continue
				}
				n++
			}
			return n
		default:
			return 0
		}
	}

	hasOmitEmptyField := func(obj interface{}) bool {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()
			n := v.NumField()
			f := t.Field(n - 1)
			tag := f.Tag.Get("enc")
			return isEncodableField(f) && strings.Contains(tag, ",omitempty")
		default:
			return false
		}
	}

	// returns the number of bytes encoded by an omitempty field on a given object
	omitEmptyLen := func(obj interface{}) int {
		if !hasOmitEmptyField(obj) {
			return 0
		}

		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			n := v.NumField()
			f := v.Field(n - 1)
			if f.Len() == 0 {
				return 0
			}
			return 4 + f.Len()

		default:
			return 0
		}
	}

	n := EncodeSizeMaxLenStringStruct1(obj)
	buf := make([]byte, n)
	err := EncodeMaxLenStringStruct1(buf, obj)
	if err != nil {
		t.Fatalf("EncodeMaxLenStringStruct1 failed: %v", err)
	}

	// A nil buffer cannot decode, unless the object is a struct with a single omitempty field
	if hasOmitEmptyField(obj) && numEncodableFields(obj) > 1 {
		t.Run(fmt.Sprintf("%d buffer underflow nil", k), func(t *testing.T) {
			decodeMaxLenStringStruct1ExpectError(t, nil, encoder.ErrBufferUnderflow)
		})
	}

	// Test all possible truncations of the encoded byte array, but skip
	// a truncation that would be valid where omitempty is removed
	skipN := n - omitEmptyLen(obj)
	for i := 0; i < n; i++ {
		if i == skipN {
			continue
		}
		t.Run(fmt.Sprintf("%d buffer underflow bytes=%d", k, i), func(t *testing.T) {
			decodeMaxLenStringStruct1ExpectError(t, buf[:i], encoder.ErrBufferUnderflow)
		})
	}

	// Append 5 bytes for omit empty with a 0 length prefix, to cause an ErrRemainingBytes.
	// If only 1 byte is appended, the decoder will try to read the 4-byte length prefix,
	// and return an ErrBufferUnderflow instead
	if hasOmitEmptyField(obj) {
		buf = append(buf, []byte{0, 0, 0, 0, 0}...)
	} else {
		buf = append(buf, 0)
	}

	// Buffer too long
	buf = append(buf, 0)
	t.Run(fmt.Sprintf("%d remaining bytes", k), func(t *testing.T) {
		decodeMaxLenStringStruct1ExpectError(t, buf[:], encoder.ErrRemainingBytes)
	})
}

func TestSkyencoderMaxLenStringStruct1DecodeErrors(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))
	n := 10

	for i := 0; i < n; i++ {
		emptyObj := newEmptyMaxLenStringStruct1ForEncodeTest()
		fullObj := newRandomMaxLenStringStruct1ForEncodeTest(t, rand)
		testSkyencoderMaxLenStringStruct1DecodeErrors(t, i, emptyObj)
		testSkyencoderMaxLenStringStruct1DecodeErrors(t, i, fullObj)
	}
}
