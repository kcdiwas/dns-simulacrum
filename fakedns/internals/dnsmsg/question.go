package dnsmsg

import (
	"encoding/binary"
	"errors"
	"strings"
)

type Question struct {
	Name  string
	Type  uint16
	Class uint16
}

func EncodeQuestion(q Question) []byte {
	qByte := encodeName(q.Name)
	typeByte := make([]byte, 4)
	binary.BigEndian.PutUint16(typeByte[0:], q.Type)
	binary.BigEndian.PutUint16(typeByte[2:], q.Class)
	qByte = append(qByte, typeByte...)

	return qByte
}

func DecodeQuestion(dc []byte, offset int) (Question, int, error) {
	name, newOffset, err := decodeName(dc, offset)
	if err != nil {
		return Question{}, 0, err
	}

	if newOffset+4 > len(dc) {
		return Question{}, 0, errors.New("Offset is greater than data")
	}

	t := binary.BigEndian.Uint16(dc[newOffset:])
	class := binary.BigEndian.Uint16(dc[newOffset+2:])

	question := Question{name, t, class}

	return question, newOffset + 4, nil
}

func encodeName(name string) []byte {
	result := make([]byte, 0)

	namearr := strings.Split(name, ".")

	for _, val := range namearr {
		result = append(result, byte(len(val)))
		result = append(result, []byte(val)...)
	}
	result = append(result, 0x00)

	return result
}

func decodeName(eName []byte, offset int) (string, int, error) {
	cursor := offset
	labels := []string{}

	if cursor >= len(eName) {
		return "", 0, errors.New("Out of bound  Cursor")
	}
	length := int(eName[cursor])

	for length != 0 {
		if cursor+length+1 > len(eName) {
			return "", 0, errors.New("Out of bound length byte")
		}
		labels = append(labels, string(eName[cursor+1:cursor+length+1]))

		cursor = cursor + length + 1
		if cursor >= len(eName) {
			return "", 0, errors.New("Out of bound Cursor")
		}
		length = int(eName[cursor])

	}
	cursor = cursor + 1

	return strings.Join(labels, "."), cursor, nil

}
