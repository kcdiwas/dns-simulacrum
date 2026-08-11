package dnsmsg

import (
	"encoding/binary"
	"errors"
)

type Header struct {
	ID                                 uint16
	Flags                              uint16
	QDCount, ANCount, NSCount, ARCount uint16
}

func (h Header) QR() bool {
	return h.Flags&0x8000 != 0
}

func (h *Header) SetQR(v bool) {
	if v {
		h.Flags = h.Flags | 0x8000
	} else {
		h.Flags = h.Flags &^ 0x8000
	}
}

func (h Header) Opcode() uint16 {
	return (h.Flags & 0x7800) >> 0x000B
}

func (h *Header) SetOpcode(v uint16) {
	h.Flags = h.Flags & 0x87FF
	h.Flags = h.Flags | (v << 11)
}

func (h Header) AA() bool {
	return h.Flags&0x0400 != 0
}

func (h *Header) SetAA(v bool) {
	if v {
		h.Flags = h.Flags | 0x0400
	} else {
		h.Flags = h.Flags &^ 0x0400
	}
}

func (h Header) TC() bool {
	return h.Flags&0x0200 != 0
}

func (h *Header) SetTC(v bool) {
	if v {
		h.Flags = h.Flags | 0x0200
	} else {
		h.Flags = h.Flags &^ 0x0200
	}
}

func (h Header) RD() bool {
	return h.Flags&0x0100 != 0
}

func (h *Header) SetRD(v bool) {
	if v {
		h.Flags = h.Flags | 0x0100
	} else {
		h.Flags = h.Flags &^ 0x0100
	}
}

func (h Header) RA() bool {
	return h.Flags&0x0080 != 0
}

func (h *Header) SetRA(v bool) {
	if v {
		h.Flags = h.Flags | 0x0080
	} else {
		h.Flags = h.Flags &^ 0x0080
	}
}

func (h Header) Z() uint16 {
	return (h.Flags & 0x0070) >> 4
}

func (h *Header) SetZ(v uint16) {
	h.Flags = h.Flags & 0xFF8F
	h.Flags = h.Flags | (v << 4)
}

func (h Header) RCode() uint16 {
	return (h.Flags & 0x000F)
}

func (h *Header) SetRCode(v uint16) {
	h.Flags = h.Flags & 0xFFF0
	h.Flags = h.Flags | v
}

// Encode Header

func EncodeHeader(h Header) []byte {
	buf := make([]byte, 12)
	binary.BigEndian.PutUint16(buf[0:], h.ID)
	binary.BigEndian.PutUint16(buf[2:], h.Flags)
	binary.BigEndian.PutUint16(buf[4:], h.QDCount)
	binary.BigEndian.PutUint16(buf[6:], h.ANCount)
	binary.BigEndian.PutUint16(buf[8:], h.NSCount)
	binary.BigEndian.PutUint16(buf[10:], h.ARCount)

	return buf
}

func DecodeHeader(data []byte) (Header, error) {
	if len(data) < 12 {
		return Header{}, errors.New("Header too short")
	}

	header := Header{}
	header.ID = binary.BigEndian.Uint16(data[0:])
	header.Flags = binary.BigEndian.Uint16(data[2:])
	header.QDCount = binary.BigEndian.Uint16(data[4:])
	header.ANCount = binary.BigEndian.Uint16(data[6:])
	header.NSCount = binary.BigEndian.Uint16(data[8:])
	header.ARCount = binary.BigEndian.Uint16(data[10:])

	return header, nil

}
