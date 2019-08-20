package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"github.com/fatedier/frp/utils/log"
	"github.com/pkg/errors"
)

const SignLength = 20

const SingMethod = "hmac_sha1"

var signKey = "lhpostoken"

type Msg struct {
	Length  uint32
	MsgType uint32
	Body    []byte
}

func NewMsg(msgType uint32, body []byte) *Msg {
	return &Msg{
		Length:  uint32(len(body)),
		MsgType: msgType,
		Body:    body,
	}
}

func (msg *Msg) GetMsgSign() []byte {
	mac := hmac.New(sha1.New, []byte(signKey))
	mac.Write(msg.Body)
	return mac.Sum(nil)
}

func (msg *Msg) CheckSign(sign []byte) bool {
	mac := hmac.New(sha1.New, []byte(signKey))
	mac.Write(msg.Body)
	MsgSign := mac.Sum(nil)
	return bytes.Equal(MsgSign, sign)
}

func (msg *Msg) GetBody() []byte {
	return msg.Body
}

type DataPackV1 struct {
}

func NewDataPackV1() *DataPackV1 {
	return &DataPackV1{}
}

func (dataPack *DataPackV1) GetHeadLen() uint32 {
	return 8
}

func (dataPack DataPackV1) Pack(msg *Msg) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	var msgType, datalen uint32
	msgType = 1
	datalen = msg.Length + SignLength

	if err := binary.Write(dataBuff, binary.BigEndian, datalen); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.BigEndian, msgType); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetMsgSign()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetBody()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (dataPack DataPackV1) UnPack(data []byte) (*Msg, error) {
	dataBuf := bytes.NewBuffer(data)
	var msgType, datalen uint32

	if err := binary.Read(dataBuf, binary.BigEndian, &datalen); err != nil {
		return nil, err
	}

	if datalen < SignLength {
		return nil, errors.New("报文数据异常")
	}
	if err := binary.Read(dataBuf, binary.BigEndian, &msgType); err != nil {
		return nil, err
	}

	sign := make([]byte, SignLength)
	if err := binary.Read(dataBuf, binary.BigEndian, sign); err != nil {

		return nil, err
	}

	msg := &Msg{}
	msg.Body = make([]byte, datalen-SignLength)
	if err := binary.Read(dataBuf, binary.BigEndian, msg.Body); err != nil {
		return nil, err
	}

	if !msg.CheckSign(sign) {
		return msg, errors.New("数据签名出错")
	}

	return msg, nil
}

func main() {
	msg := NewMsg(1, []byte("123456"))
	dataPack := NewDataPackV1()
	pdata, err := dataPack.Pack(msg)
	if err != nil {
		log.Debug("pack %s", err.Error())
	}
	log.Debug("pack Data: %x", pdata)
	unPackData, err := dataPack.UnPack(pdata)
	if err != nil {
		log.Debug("unpack %s", err.Error())
		return
	}
	log.Debug("upPackData %s", unPackData.Body)

}
