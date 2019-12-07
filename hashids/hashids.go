package hashids

import (
	"encoding/hex"
	"fmt"
)
import "github.com/speps/go-hashids"

type HashIdsHelper struct {
	HashID *hashids.HashID
}

func NewHashIdsHelper(salt string) (*HashIdsHelper, error) {
	hd := hashids.NewData()
	hd.MinLength = 30
	hd.Salt = salt
	hashId, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}
	return &HashIdsHelper{HashID: hashId}, err
}

func (h *HashIdsHelper) EncodeString(input string) (string, error) {
	return h.HashID.EncodeHex(hex.EncodeToString([]byte(input)))
}

func (h *HashIdsHelper) DecodeString(input string) (string, error) {
	d, err := h.HashID.DecodeHex(input)
	if err != nil {
		return "", err
	}
	b, err := hex.DecodeString(d)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func EncodeString(input, salt string) (string, error) {
	hd := hashids.NewData()
	hd.MinLength = 30
	hd.Salt = salt
	hd.Alphabet =  hashids.DefaultAlphabet
	hashId, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	fmt.Printf("raw : %s , hex.EncodeToString: %s\n", input, hex.EncodeToString([]byte(input)))
	return hashId.EncodeHex(hex.EncodeToString([]byte(input)))
}

func DecodeString(input, salt string) (string, error) {
	hd := hashids.NewData()
	hd.MinLength = 30
	hd.Salt = salt
	hashId, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	hexStr, err := hashId.DecodeHex(input)
	if err != nil {
		return "", err
	}
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return string(b), err
}
