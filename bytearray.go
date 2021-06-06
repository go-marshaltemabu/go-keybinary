package keybinary

import (
	"encoding/base64"
)

const (
	expectEncodedByteArray32 = 43 // base64.RawStdEncoding.EncodedLen(32)
	expectEncodedByteArray64 = 86 // base64.RawStdEncoding.EncodedLen(64)
)

var emptyByteArray32 [32]byte
var emptyByteArray64 [64]byte

// ByteArray32 contain a 32 bytes array.
type ByteArray32 struct {
	a [32]byte
}

// NewByteArray32 create new instance of ByteArray32 with given array reference.
// If arrayRef is nil the resulted instance will have array fill with empty (0/zero) value.
func NewByteArray32(arrayRef *[32]byte) (k *ByteArray32) {
	if arrayRef == nil {
		k = &ByteArray32{}
	} else {
		k = &ByteArray32{
			a: *arrayRef,
		}
	}
	return
}

// Ref return pointer references array.
func (k *ByteArray32) Ref() (ref *[32]byte) {
	ref = &k.a
	return
}

// Load copy given array into instance.
func (k *ByteArray32) Load(arrayRef *[32]byte) {
	k.a = *arrayRef
}

// Clear empty key content.
func (k *ByteArray32) Clear() {
	copy(k.a[:], make([]byte, 32))
}

// MarshalBinary implement encoding.BinaryMarshaler interface.
func (k *ByteArray32) MarshalBinary() (data []byte, err error) {
	if k == nil {
		return
	}
	data = make([]byte, 32)
	copy(data, k.a[:])
	return
}

// UnmarshalBinary implement encoding.BinaryUnmarshaler interface.
func (k *ByteArray32) UnmarshalBinary(data []byte) (err error) {
	if l := len(data); l == 0 {
		k.Clear()
		return
	} else if l != 32 {
		err = &ErrIncorrectDataSize{
			ExpectSize:   32,
			ReceivedSize: len(data),
		}
		return
	}
	copy(k.a[:], data)
	return
}

// MarshalText implement encoding.TextMarshaler interface.
func (k *ByteArray32) MarshalText() (text []byte, err error) {
	if (k == nil) || (k.a == emptyByteArray32) {
		return
	}
	text = make([]byte, expectEncodedByteArray32)
	base64.RawStdEncoding.Encode(text, k.a[:])
	return
}

// UnmarshalText implement encoding.TextUnmarshaler interface.
func (k *ByteArray32) UnmarshalText(text []byte) (err error) {
	if l := len(text); l == 0 {
		k.Clear()
		return
	} else if l != expectEncodedByteArray32 {
		err = &ErrIncorrectDataSize{
			ExpectSize:   expectEncodedByteArray32,
			ReceivedSize: l,
		}
		return
	}
	_, err = base64.RawStdEncoding.Decode(k.a[:], text)
	return
}

// String convert k into string.
// Resulted string will be base64.RawStdEncoding encoded or empty string if k is nil.
func (k *ByteArray32) String() string {
	buf, _ := k.MarshalText()
	return string(buf)
}

// ByteArray64 contain a 64 bytes key.
type ByteArray64 struct {
	a [64]byte
}

// NewByteArray64 create new instance of ByteArray64 with given key.
// If arrayRef is nil the resulted instance will have array fill with empty (0/zero) value.
func NewByteArray64(arrayRef *[64]byte) (k *ByteArray64) {
	if arrayRef == nil {
		k = &ByteArray64{}
	} else {
		k = &ByteArray64{
			a: *arrayRef,
		}
	}
	return
}

// Ref return pointer references array.
func (k *ByteArray64) Ref() (ref *[64]byte) {
	ref = &k.a
	return
}

// Load copy given key into instance.
func (k *ByteArray64) Load(arrayRef *[64]byte) {
	k.a = *arrayRef
}

// Clear empty key content.
func (k *ByteArray64) Clear() {
	copy(k.a[:], make([]byte, 64))
}

// MarshalBinary implement encoding.BinaryMarshaler interface.
func (k *ByteArray64) MarshalBinary() (data []byte, err error) {
	if k == nil {
		return
	}
	data = make([]byte, 64)
	copy(data, k.a[:])
	return
}

// UnmarshalBinary implement encoding.BinaryUnmarshaler interface.
func (k *ByteArray64) UnmarshalBinary(data []byte) (err error) {
	if l := len(data); l == 0 {
		k.Clear()
		return
	} else if l != 64 {
		err = &ErrIncorrectDataSize{
			ExpectSize:   64,
			ReceivedSize: len(data),
		}
		return
	}
	copy(k.a[:], data)
	return
}

// MarshalText implement encoding.TextMarshaler interface.
func (k *ByteArray64) MarshalText() (text []byte, err error) {
	if (k == nil) || (k.a == emptyByteArray64) {
		return
	}
	text = make([]byte, expectEncodedByteArray64)
	base64.RawStdEncoding.Encode(text, k.a[:])
	return
}

// UnmarshalText implement encoding.TextUnmarshaler interface.
func (k *ByteArray64) UnmarshalText(text []byte) (err error) {
	if l := len(text); l == 0 {
		k.Clear()
		return
	} else if l != expectEncodedByteArray64 {
		err = &ErrIncorrectDataSize{
			ExpectSize:   expectEncodedByteArray64,
			ReceivedSize: l,
		}
		return
	}
	_, err = base64.RawStdEncoding.Decode(k.a[:], text)
	return
}

// String convert k into string.
// Resulted string will be base64.RawStdEncoding encoded or empty string if k is nil.
func (k *ByteArray64) String() string {
	buf, _ := k.MarshalText()
	return string(buf)
}
