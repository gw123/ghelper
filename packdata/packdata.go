package packdata

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"net"
)

const HeaderLen = 4
const HeaderDataLen = 2
const HeaderDataTypeLen = 2

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

func (msg *Msg) GetMsgSign(signer Signer) ([]byte, error) {
	return signer.Sign(msg.Body)
}

func (msg *Msg) GetBody() []byte {
	return msg.Body
}

type DataPackV1 struct {
	signer Signer
}

func NewDataPackV1(signer Signer) *DataPackV1 {
	return &DataPackV1{signer: signer}
}

func (dataPack *DataPackV1) GetHeadLen() uint32 {
	return 8
}

func (dataPack DataPackV1) Pack(msg *Msg) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	var msgType, datalen uint32
	msgType = msg.MsgType
	datalen = msg.Length + SignLength

	if err := binary.Write(dataBuff, binary.BigEndian, datalen); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.BigEndian, msgType); err != nil {
		return nil, err
	}

	sign, err := msg.GetMsgSign(dataPack.signer)
	if err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.BigEndian, sign[0:SignLength]); err != nil {
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

	trueSign, err := msg.GetMsgSign(dataPack.signer)
	if err != nil {
		return msg, err
	}

	if len(trueSign) < SignLength {
		return msg, errors.New("长度校验失败")
	}

	if !bytes.Equal(sign, trueSign[0:SignLength]) {
		return msg, errors.New("签名错误...")
	}

	return msg, nil
}

func (dataPack DataPackV1) UnPackFromConn(conn net.Conn) (*Msg, error) {
	var msgType, datalen uint32

	if err := binary.Read(conn, binary.BigEndian, &datalen); err != nil {
		return nil, err
	}

	if datalen < SignLength {
		return nil, errors.New("报文数据异常")
	}
	if err := binary.Read(conn, binary.BigEndian, &msgType); err != nil {
		return nil, err
	}

	sign := make([]byte, SignLength)
	if err := binary.Read(conn, binary.BigEndian, sign); err != nil {
		return nil, err
	}

	msg := &Msg{}
	msg.Body = make([]byte, datalen-SignLength)
	if err := binary.Read(conn, binary.BigEndian, msg.Body); err != nil {
		return nil, err
	}

	trueSign, err := msg.GetMsgSign(dataPack.signer)
	if err != nil {
		return msg, err
	}

	if len(trueSign) < SignLength {
		return msg, errors.New("长度校验失败")
	}

	if !bytes.Equal(sign, trueSign[0:SignLength]) {
		return msg, errors.New("签名错误...")
	}

	return msg, nil
}
