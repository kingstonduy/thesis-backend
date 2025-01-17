package isomessage

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/logger"
	"github.com/moov-io/iso8583"
)

type IsoMessage struct {
	hmac      string
	isoMsgLib iso8583.Message
}

func NewIsoMessage(hmac string) *IsoMessage {
	return &IsoMessage{
		hmac:      hmac,
		isoMsgLib: *iso8583.NewMessage(Spec),
	}
}

func (i *IsoMessage) Verify() error {
	_, err := i.isoMsgLib.Pack()
	if err != nil {
		return err
	}
	return nil
}

// this function already add the size to the head of the string
func (i *IsoMessage) ToString() string {
	dataByte, err := i.isoMsgLib.Pack()
	if err != nil {
		logger.Error(context.Background(), err.Error())
	}

	res := padLeft(fmt.Sprintf("%d", len(string(dataByte))), 4, '0') + string(dataByte)
	return res
}

func (i *IsoMessage) SetField(idx int, val string) error {
	return i.isoMsgLib.Field(idx, val)
}

func (i *IsoMessage) GetMTI() string {
	return i.GetField(0)
}

func (i *IsoMessage) SetMTI(s string) error {
	return i.SetField(0, s)
}

func (i *IsoMessage) GetField(idx int) string {
	var s string

	s, err := i.isoMsgLib.GetField(idx).String()
	if err != nil {
		logger.Error(context.Background(), err)
	}

	return s
}

func (i *IsoMessage) GenHMAC() {
	i.SetField(128, GenHMAC(&i.isoMsgLib, i.hmac))
}

func (i *IsoMessage) GetNapasID() string {
	var res string

	var fields []int = []int{2, 7, 11, 32, 37, 41}

	for _, f := range fields {
		res += i.GetField(f)
	}

	return res
}

func ToIsoMessage(s string, hmac string) (*IsoMessage, error) {
	s = s[4:]
	msg := iso8583.NewMessage(Spec)
	err := msg.Unpack([]byte(s))
	if err != nil {
		return nil, err
	}

	isoMsg := NewIsoMessage(hmac)
	err = isoMsg.isoMsgLib.Unpack([]byte(s))
	if err != nil {
		return nil, err
	}

	return isoMsg, nil
}
