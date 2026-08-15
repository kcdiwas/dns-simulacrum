package dnsmsg

import (
	"bytes"
	"testing"
)

func TestEncodeName(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []byte
	}{
		{"simple", "www.example.test", []byte{3, 'w', 'w', 'w', 7, 'e', 'x', 'a', 'm', 'p', 'l', 'e', 4, 't', 'e', 's', 't', 0x00}},
		{"single label", "test", []byte{4, 't', 'e', 's', 't', 0x00}},
	}

	for _, c := range cases {
		got := encodeName(c.in)
		if !bytes.Equal(got, c.want) {
			t.Errorf("%s: got % X, want % X", c.name, got, c.want)
		}
	}
}

func TestDecodeName(t *testing.T) {
	cases := []struct {
		name       string
		in         []byte
		wantName   string
		wantOffset int
	}{
		{
			"it should return correct name and offset",
			[]byte{3, 'w', 'w', 'w', 7, 'e', 'x', 'a', 'm', 'p', 'l', 'e', 4, 't', 'e', 's', 't', 0},
			"www.example.test",
			18,
		},
		{
			"it should return error if bad offset or out of bound length is provided",
			[]byte{5, 't', 'e', 's', 't'},
			"",
			0,
		},
	}

	for _, c := range cases {
		got_name, got_offset, _ := decodeName(c.in, 0)
		if got_name != c.wantName {
			t.Errorf("%s : got Name % X, want Name % X", c.name, got_name, c.wantName)
		}
		if got_offset != c.wantOffset {
			t.Errorf("%s : got Offset % X, want Offset % X", c.name, got_offset, c.wantOffset)
		}
	}
}
