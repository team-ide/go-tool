package util

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestByte(t *testing.T) {
	var bs []byte
	b := intToBytes(1937006964)
	bs = append(bs, b...)
	b = intToBytes(1195725856)
	bs = append(bs, b...)
	b = intToBytes(369295618)
	bs = append(bs, b...)
	b = intToBytes(1937006964)
	bs = append(bs, b...)
	b = intToBytes(1195725856)
	bs = append(bs, b...)
	b = intToBytes(369295618)
	bs = append(bs, b...)
	fmt.Println("bs:", string(bs))

}
func intToBytes(n int) []byte {
	a := uint32(n)
	var bytes []byte = []byte{0, 0, 0, 0}
	//相当于把这个长度转成len
	binary.BigEndian.PutUint32(bytes[0:4], a)
	return bytes
}
func IntToBytes(a int) ([]byte, error) {
	if a > math.MaxInt32 {
		return nil, errors.New(fmt.Sprintf("a>math.MaxInt32, a is %d\n", a))
	}
	buf := make([]byte, 4)
	for i := 0; i < 4; i++ {
		var b uint8 = uint8(a & 0xff)
		buf[i] = b
		a = a >> 8
	}
	return buf, nil
}

func BytesToInt(buf []byte) (int, error) {
	if len(buf) != 4 {
		return -1, errors.New(fmt.Sprintf("BytesToInt len(buf) must be 4, but got %d\n", len(buf)))
	}
	result := 0
	for i := 0; i < 4; i++ {
		result += int(buf[i]) << (8 * i)
	}
	return result, nil
}
