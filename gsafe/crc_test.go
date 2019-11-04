package gsafe

import (
	"bytes"
	"testing"
)

type TestItem struct {
	Val           []byte
	CRC_8_Result  []byte
	CRC_32_Result []byte
}

func TestCrc8(t *testing.T) {
	testDatas := []TestItem{
		{
			Val:           []byte("hello"),
			CRC_8_Result:  []byte{0x92},
			CRC_32_Result: []byte{0x36, 0x10, 0xA6, 0x86},
		},
		{
			Val:           []byte("hello world"),
			CRC_8_Result:  []byte{0xa8},
			CRC_32_Result: []byte{0x0D, 0x4A, 0x11, 0x85},
		},
		{
			Val:           []byte("1234567890qwertyuiop"),
			CRC_8_Result:  []byte{0x72},
			CRC_32_Result: []byte{0xF6, 0xB7, 0x9B, 0x3A},
		},
		{
			Val:           []byte{0x10, 0x20, 0x30, 0x40, 0x50, 0x60},
			CRC_8_Result:  []byte{0xfe},
			CRC_32_Result: []byte{0x04, 0x6D, 0xC5, 0x51},
		},
	}

	for _, item := range testDatas {
		res := Crc8(item.Val)
		if res != item.CRC_8_Result[0] {
			t.Fail()
		}

		res2 := Crc32Byte(item.Val)
		if bytes.Compare(res2, item.CRC_32_Result) != 0 {
			t.Fail()
		}

	}
}
