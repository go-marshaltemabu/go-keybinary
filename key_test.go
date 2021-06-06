package naclkey_test

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"testing"

	naclkey "github.com/go-marshaltemabu/go-naclkey"
)

func TestKey32_Load(t *testing.T) {
	var rawKey [32]byte
	rand.Read(rawKey[:])
	k1 := naclkey.NewKey32(&rawKey)
	var k2 naclkey.Key32
	k2.Load(&rawKey)
	if *k1 != k2 {
		t.Errorf("unexpected key difference: %s vs. %s",
			k1.String(), k2.String())
	}
}

func TestKey32_Clear(t *testing.T) {
	var rawKey [32]byte
	rand.Read(rawKey[:])
	rawKey[0] = 'a'
	k1 := naclkey.NewKey32(&rawKey)
	var kEmpty naclkey.Key32
	if *k1 == kEmpty {
		t.Errorf("unexpected key equal: %s vs. %s",
			k1.String(), kEmpty.String())
	}
	k1.Clear()
	if *k1 != kEmpty {
		t.Errorf("unexpected key difference: %s vs. %s",
			k1.String(), kEmpty.String())
	}
}

func TestKey32_Binary(t *testing.T) {
	for attempt := 0; attempt < 20; attempt++ {
		var rawKey [32]byte
		rand.Read(rawKey[:])
		k1 := naclkey.NewKey32(&rawKey)
		buf, err := k1.MarshalBinary()
		if nil != err {
			t.Errorf("invoke MarshalBinary failed: %v", err)
			continue
		}
		if !bytes.Equal(rawKey[:], buf) {
			t.Errorf("binary encoded key different: %s vs. %s",
				base64.RawStdEncoding.EncodeToString(rawKey[:]),
				base64.RawStdEncoding.EncodeToString(buf))
		}
		k2 := naclkey.NewKey32(nil)
		if err = k2.UnmarshalBinary(buf); nil != err {
			t.Errorf("invoke UnmarshalBinary failed: %v", err)
			continue
		}
		if *k1 != *k2 {
			t.Errorf("binary decoded key different: %s vs. %s",
				k1.String(), k2.String())
		}
		t.Logf("attempt %02d: %s", attempt, k1.String())
	}
	var k0 *naclkey.Key32
	buf, err := k0.MarshalBinary()
	if nil != err {
		t.Errorf("invoke MarshalBinary failed: %v", err)
	}
	if len(buf) != 0 {
		t.Errorf("unexpect MarshalBinary result: %v", buf)
	}
	var rawKey [32]byte
	rand.Read(rawKey[:])
	k0 = naclkey.NewKey32(&rawKey)
	k0.UnmarshalBinary(nil)
	var kEmpty naclkey.Key32
	if kEmpty != *k0 {
		t.Errorf("empty binary key different: %s vs. %s",
			k0.String(), kEmpty.String())
	}
	var k13 *naclkey.Key32
	if err = k13.UnmarshalBinary([]byte{0, 1, 2}); nil == err {
		t.Error("expecting error for malformed binary pack")
	} else if _, ok := err.(*naclkey.ErrIncorrectDataSize); !ok {
		t.Errorf("unexpect error: %v", err)
	}
}

func TestKey32_Text(t *testing.T) {
	targets := []string{
		"o2kBLbktGE/DnRc0/1cWQolTu2hl/PkrDDoXyQKL6ZE",
		"7ilLOfMrfHgiumT4SrQ8oMbmuRwf076JkENBedOvRJE",
		"TrdknGyTR4AJedGDA1bypUw96rKktEddY6++j7Vph8c",
		"C/UFmHWSHmaKW98sf8SERZLSVyvNBmjS1sUvUFTi0IM",
	}
	for idx, data := range targets {
		var k naclkey.Key32
		if err := k.UnmarshalText([]byte(data)); nil != err {
			t.Errorf("invoke UnmarshalText failed (%d): %v", idx, err)
			continue
		}
		buf, err := k.MarshalText()
		if nil != err {
			t.Errorf("invoke MarshalText failed (%d): %v", idx, err)
			continue
		}
		if string(buf) != data {
			t.Errorf("text transcoded key different: %s vs. %s",
				string(buf), data)
		}
	}
	var k0 *naclkey.Key32
	buf, err := k0.MarshalText()
	if nil != err {
		t.Errorf("invoke MarshalText failed: %v", err)
	}
	if len(buf) != 0 {
		t.Errorf("unexpect MarshalText result: %v", buf)
	}
	var rawKey [32]byte
	rand.Read(rawKey[:])
	k0 = naclkey.NewKey32(&rawKey)
	k0.UnmarshalText(nil)
	var kEmpty naclkey.Key32
	if kEmpty != *k0 {
		t.Errorf("empty text key different: %s vs. %s",
			k0.String(), kEmpty.String())
	}
	var k13 *naclkey.Key32
	if err = k13.UnmarshalText([]byte{'a', 'b', 'c'}); nil == err {
		t.Error("expecting error for malformed text pack")
	} else if _, ok := err.(*naclkey.ErrIncorrectDataSize); !ok {
		t.Errorf("unexpect error: %v", err)
	}
}

func TestKey64_Load(t *testing.T) {
	var rawKey [64]byte
	rand.Read(rawKey[:])
	k1 := naclkey.NewKey64(&rawKey)
	var k2 naclkey.Key64
	k2.Load(&rawKey)
	if *k1 != k2 {
		t.Errorf("unexpected key difference: %s vs. %s",
			k1.String(), k2.String())
	}
}

func TestKey64_Clear(t *testing.T) {
	var rawKey [64]byte
	rand.Read(rawKey[:])
	rawKey[0] = 'a'
	k1 := naclkey.NewKey64(&rawKey)
	var kEmpty naclkey.Key64
	if *k1 == kEmpty {
		t.Errorf("unexpected key equal: %s vs. %s",
			k1.String(), kEmpty.String())
	}
	k1.Clear()
	if *k1 != kEmpty {
		t.Errorf("unexpected key difference: %s vs. %s",
			k1.String(), kEmpty.String())
	}
}

func TestKey64_Binary(t *testing.T) {
	for attempt := 0; attempt < 20; attempt++ {
		var rawKey [64]byte
		rand.Read(rawKey[:])
		k1 := naclkey.NewKey64(&rawKey)
		buf, err := k1.MarshalBinary()
		if nil != err {
			t.Errorf("invoke MarshalBinary failed: %v", err)
			continue
		}
		if !bytes.Equal(rawKey[:], buf) {
			t.Errorf("binary encoded key different: %s vs. %s",
				base64.RawStdEncoding.EncodeToString(rawKey[:]),
				base64.RawStdEncoding.EncodeToString(buf))
		}
		k2 := naclkey.NewKey64(nil)
		if err = k2.UnmarshalBinary(buf); nil != err {
			t.Errorf("invoke UnmarshalBinary failed: %v", err)
			continue
		}
		if *k1 != *k2 {
			t.Errorf("binary decoded key different: %s vs. %s",
				k1.String(), k2.String())
		}
		t.Logf("attempt %02d: %s", attempt, k1.String())
	}
	var k0 *naclkey.Key64
	buf, err := k0.MarshalBinary()
	if nil != err {
		t.Errorf("invoke MarshalBinary failed: %v", err)
	}
	if len(buf) != 0 {
		t.Errorf("unexpect MarshalBinary result: %v", buf)
	}
	var rawKey [64]byte
	rand.Read(rawKey[:])
	k0 = naclkey.NewKey64(&rawKey)
	k0.UnmarshalBinary(nil)
	var kEmpty naclkey.Key64
	if kEmpty != *k0 {
		t.Errorf("empty binary key different: %s vs. %s",
			k0.String(), kEmpty.String())
	}
	var k13 *naclkey.Key32
	if err = k13.UnmarshalBinary([]byte{0, 1, 2}); nil == err {
		t.Error("expecting error for malformed binary pack")
	} else if _, ok := err.(*naclkey.ErrIncorrectDataSize); !ok {
		t.Errorf("unexpect error: %v", err)
	}
}

func TestKey64_Text(t *testing.T) {
	targets := []string{
		"mDIutc9D1yvS5biH1GMPuNR0fq1uuCrNHFsHgUPuJqWGrSMTnVBBcjRwvySoZYN8kSNGHEH1/5mqmc4k6014hQ",
		"5QvhptwdV2joU3mI/dzlYum5SMkYu6PpM+XEAM3l5gxerW/Hrne6HSWbGIpLIchvvCPXKLRTR+raZQryTFbQgA",
		"nNzFlbzOPHvT2N+T+rfhJd3rr+ZaMb1dQeLSzpwrF4kvD+oZMaKQIgd3qTFD39y/poQG6HcHP/CINOGXpANKpA",
		"DvHDFAkPB8eab1ccJG8+msC3QT7xEL1YsAznO/9wb3/0tvRAkKMnEfMgjk5LictRZc5kACy9nCiHqhE98kaJKA",
	}
	for idx, data := range targets {
		var k naclkey.Key64
		if err := k.UnmarshalText([]byte(data)); nil != err {
			t.Errorf("invoke UnmarshalText failed (%d): %v", idx, err)
			continue
		}
		buf, err := k.MarshalText()
		if nil != err {
			t.Errorf("invoke MarshalText failed (%d): %v", idx, err)
			continue
		}
		if string(buf) != data {
			t.Errorf("text transcoded key different: %s vs. %s",
				string(buf), data)
		}
	}
	var k0 *naclkey.Key64
	buf, err := k0.MarshalText()
	if nil != err {
		t.Errorf("invoke MarshalText failed: %v", err)
	}
	if len(buf) != 0 {
		t.Errorf("unexpect MarshalText result: %v", buf)
	}
	var rawKey [64]byte
	rand.Read(rawKey[:])
	k0 = naclkey.NewKey64(&rawKey)
	k0.UnmarshalText(nil)
	var kEmpty naclkey.Key64
	if kEmpty != *k0 {
		t.Errorf("empty text key different: %s vs. %s",
			k0.String(), kEmpty.String())
	}
	var k13 *naclkey.Key64
	if err = k13.UnmarshalText([]byte{'a', 'b', 'c'}); nil == err {
		t.Error("expecting error for malformed text pack")
	} else if _, ok := err.(*naclkey.ErrIncorrectDataSize); !ok {
		t.Errorf("unexpect error: %v", err)
	}
}