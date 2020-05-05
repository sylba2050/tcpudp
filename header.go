package pubsub

// MessageType is MessageType in header
type MessageType uint16

const HeaderSize = 8

const (
	Connect                         MessageType = 0x0001
	ConnectAckSuccess               MessageType = 0x0002
	ConnectAckFailure               MessageType = 0x0003
	ReconnectRequest                MessageType = 0x0004
	Disconnect                      MessageType = 0x0005
	CreateNewTopicRequest           MessageType = 0x0010
	Subscribe                       MessageType = 0x0020
	UnSubscribe                     MessageType = 0x0021
	SubscribeAckSuccess             MessageType = 0x0022
	SubscribePermisionError         MessageType = 0x0023
	Publish                         MessageType = 0x0040
	PublishAckSuccess               MessageType = 0x0041
	PublishPermissionError          MessageType = 0x0042
	Ping                            MessageType = 0x0060
	Pong                            MessageType = 0x0061
	AddSubscriber                   MessageType = 0x0100
	AddSubscriberPermissionError    MessageType = 0x0101
	RemoveSubscriber                MessageType = 0x0102
	RemoveSubscriberPermissionError MessageType = 0x0103
	GetSubscribersRequest           MessageType = 0x1000
	GetSubscribersPermissionError   MessageType = 0x1001
	SubscribersList                 MessageType = 0x1002
	CloseRequest                    MessageType = 0x2000
)

type Header interface {
	SetType(MessageType) error
	GetType() (MessageType, error)
	SetLength(uint16) error
	GetLength() (uint16, error)
	SetSenderTimestamp(uint32) error
	GetSenderTimestamp() (uint32, error)
	SetReceiverTimestamp(uint32) error
	GetReceiverTimestamp() (uint32, error)
	ToBytes() ([]byte, error)
}

type H struct {
	typ               MessageType
	length            uint16
	senderTimestamp   uint32
	receiverTimestamp uint32
}

func (h *H) SetType(m MessageType) error {
	return nil
}

func (h H) GetType() (MessageType, error) {
	return 0x0000, nil
}

func (h *H) SetLength(length uint16) error {
	return nil
}

func (h H) GetLength() (uint16, error) {
	return 0, nil
}

func (h *H) SetSenderTimestamp(timestamp uint32) error {
	return nil
}

func (h H) GetSenderTimestamp() (uint32, error) {
	return 0, nil
}

func (h *H) SetReceiverTimestamp(timestamp uint32) error {
	return nil
}

func (h H) GetReceiverTimestamp() (uint32, error) {
	return 0, nil
}

func (h *H) ToBytes() ([]byte, error) {
	return nil, nil
}