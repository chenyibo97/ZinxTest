package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"studygo2/zinxtest/utils"
	"studygo2/zinxtest/ziface"
)

type DataPack struct {
}

//datalen|dataid|data
func (d *DataPack) GetHeadLen() uint32 {
	return 8
}
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
	//先写datalen，再写msgid，在写data
}

func (d *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	reader := bytes.NewReader(data)
	//先读4位datalen 和data id
	msg := &Message{}
	err := binary.Read(reader, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}
	if msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New(fmt.Sprintf("data out of max,data:%d", msg.DataLen))
	}
	return msg, nil

}

func NewDataPack() *DataPack {
	return &DataPack{}
}
