package ethrpc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type BlockNumber int64

const (
	EarliestBlock BlockNumber = 0
	PendingBlock              = -1
	LatestBlock               = -2
)

func (n BlockNumber) String() string {
	switch n {
	case EarliestBlock:
		return "earliest"
	case PendingBlock:
		return "pending"
	case LatestBlock:
		return "latest"
	default:
		return fmt.Sprintf("%#x", int64(n))
	}
}

func (n BlockNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *BlockNumber) UnmarshalJSON(b []byte) error {
	var u Number
	err := json.Unmarshal(b, &u)
	if err != nil {
		return err
	}
	*n = BlockNumber(int64(u))
	return nil
}

type Address []byte

func (a Address) String() string {
	return "0x" + hex.EncodeToString([]byte(a))
}

func (a Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a *Address) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*a, err = AddressFromString(s)
	return err
}

func AddressFromString(s string) (Address, error) {
	s = strings.TrimPrefix(s, "0x")
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	if len(b) != 20 {
		return nil, errors.New("address must be 20 bytes long")
	}
	return Address(b), nil
}

type Hash []byte

func (h Hash) String() string {
	return "0x" + hex.EncodeToString([]byte(h))
}

func (h Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

func (h *Hash) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*h, err = HashFromString(s)
	return err
}

func HashFromString(s string) (Hash, error) {
	s = strings.TrimPrefix(s, "0x")
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	if len(b) != 32 {
		return nil, errors.New("hash must be 32 bytes long")
	}
	return Hash(b), nil
}

func MustHashFromString(s string) Hash {
	hash, err := HashFromString(s)
	if err != nil {
		panic(err)
	}
	return hash
}

type LogTopic []byte

func (t LogTopic) String() string {
	return "0x" + hex.EncodeToString([]byte(t))
}

func (t LogTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *LogTopic) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t, err = LogTopicFromString(s)
	return err
}

func LogTopicFromString(s string) (LogTopic, error) {
	s = strings.TrimPrefix(s, "0x")
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	if len(b) != 32 {
		return nil, errors.New("log topic must be 32 bytes long")
	}
	return LogTopic(b), nil
}

func MustLogTopicFromString(s string) LogTopic {
	topic, err := LogTopicFromString(s)
	if err != nil {
		panic(err)
	}
	return topic
}

type Number uint64

func (n Number) String() string {
	if n == 0 {
		return "0x"
	}
	return "0x" + strconv.FormatUint(uint64(n), 16)
}

func (n Number) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *Number) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*n, err = NumberFromString(s)
	return err
}

func NumberFromString(s string) (Number, error) {
	s = strings.TrimPrefix(s, "0x")
	if s == "" {
		return Number(0), nil
	}
	u, err := strconv.ParseUint(s, 16, 64)
	return Number(u), err
}

func MustNumberFromString(s string) Number {
	n, err := NumberFromString(s)
	if err != nil {
		panic(err)
	}
	return n
}

type Data []byte

func (d Data) String() string {
	if d == nil {
		return "0x"
	}
	return "0x" + hex.EncodeToString([]byte(d))
}

func (d Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Data) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*d, err = DataFromString(s)
	return err
}

func DataFromString(s string) ([]byte, error) {
	s = strings.TrimPrefix(s, "0x")
	if s == "" {
		return nil, nil
	}
	return hex.DecodeString(s)
}

func MustDataFromString(s string) Data {
	d, err := DataFromString(s)
	if err != nil {
		panic(err)
	}
	return d
}
