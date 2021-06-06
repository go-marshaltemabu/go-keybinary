package naclkey

import (
	"encoding/base64"
)

const (
	expectEncodedKey32 = 43 // base64.RawStdEncoding.EncodedLen(32)
	expectEncodedKey64 = 86 // base64.RawStdEncoding.EncodedLen(64)
)

// Key32 contain a 32 bytes key.
type Key32 struct {
	key [32]byte
}

// NewKey32 create new instance of Key32 with given key.
// If key is nil the resulted instance will have key fill with empty (0/zero) value.
func NewKey32(key *[32]byte) (k *Key32) {
	if key == nil {
		k = &Key32{}
	} else {
		k = &Key32{
			key: *key,
		}
	}
	return
}

// Load copy given key into instance.
func (k *Key32) Load(key *[32]byte) {
	k.key = *key
}

// Clear empty key content.
func (k *Key32) Clear() {
	copy(k.key[:], make([]byte, 32))
}

// MarshalBinary implement encoding.BinaryMarshaler interface.
func (k *Key32) MarshalBinary() (data []byte, err error) {
	if k == nil {
		return
	}
	data = make([]byte, 32)
	copy(data, k.key[:])
	return
}

// UnmarshalBinary implement encoding.BinaryUnmarshaler interface.
func (k *Key32) UnmarshalBinary(data []byte) (err error) {
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
	copy(k.key[:], data)
	return
}

// MarshalText implement encoding.TextMarshaler interface.
func (k *Key32) MarshalText() (text []byte, err error) {
	if k == nil {
		return
	}
	text = make([]byte, expectEncodedKey32)
	base64.RawStdEncoding.Encode(text, k.key[:])
	return
}

// UnmarshalText implement encoding.TextUnmarshaler interface.
func (k *Key32) UnmarshalText(text []byte) (err error) {
	if l := len(text); l == 0 {
		k.Clear()
		return
	} else if l != expectEncodedKey32 {
		err = &ErrIncorrectDataSize{
			ExpectSize:   expectEncodedKey32,
			ReceivedSize: l,
		}
		return
	}
	_, err = base64.RawStdEncoding.Decode(k.key[:], text)
	return
}

// String convert k into string.
// Resulted string will be base64.RawStdEncoding encoded or empty string if k is nil.
func (k *Key32) String() string {
	buf, _ := k.MarshalText()
	return string(buf)
}

// Key64 contain a 64 bytes key.
type Key64 struct {
	key [64]byte
}

// NewKey64 create new instance of Key32 with given key.
// If key is nil the resulted instance will have key fill with empty (0/zero) value.
func NewKey64(key *[64]byte) (k *Key64) {
	if key == nil {
		k = &Key64{}
	} else {
		k = &Key64{
			key: *key,
		}
	}
	return
}

// Load copy given key into instance.
func (k *Key64) Load(key *[64]byte) {
	k.key = *key
}

// Clear empty key content.
func (k *Key64) Clear() {
	copy(k.key[:], make([]byte, 64))
}

// MarshalBinary implement encoding.BinaryMarshaler interface.
func (k *Key64) MarshalBinary() (data []byte, err error) {
	if k == nil {
		return
	}
	data = make([]byte, 64)
	copy(data, k.key[:])
	return
}

// UnmarshalBinary implement encoding.BinaryUnmarshaler interface.
func (k *Key64) UnmarshalBinary(data []byte) (err error) {
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
	copy(k.key[:], data)
	return
}

// MarshalText implement encoding.TextMarshaler interface.
func (k *Key64) MarshalText() (text []byte, err error) {
	if k == nil {
		return
	}
	text = make([]byte, expectEncodedKey64)
	base64.RawStdEncoding.Encode(text, k.key[:])
	return
}

// UnmarshalText implement encoding.TextUnmarshaler interface.
func (k *Key64) UnmarshalText(text []byte) (err error) {
	if l := len(text); l == 0 {
		k.Clear()
		return
	} else if l != expectEncodedKey64 {
		err = &ErrIncorrectDataSize{
			ExpectSize:   expectEncodedKey64,
			ReceivedSize: l,
		}
		return
	}
	_, err = base64.RawStdEncoding.Decode(k.key[:], text)
	return
}

// String convert k into string.
// Resulted string will be base64.RawStdEncoding encoded or empty string if k is nil.
func (k *Key64) String() string {
	buf, _ := k.MarshalText()
	return string(buf)
}
