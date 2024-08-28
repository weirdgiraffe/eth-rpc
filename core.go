package ethrpc

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Number uint64

func (n Number) String() string {
	if n == 0 {
		return "0x"
	}
	return "0x" + strconv.FormatUint(uint64(n), 16)
}

func (n Number) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *Number) UnmarshalText(b []byte) (err error) {
	s := string(b)
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

type BlockNumber int64

// BlockNumberFromInt64 returns a BlockNumber from an int64. It panics if the
// block number is negative.
// Use [PendingBlock], [EarliestBlock] or [LatestBlock] for special block numbers.
func BlockNumberFromInt64(i int64) BlockNumber {
	if i < 0 {
		panic("block number must be non-negative")
	}
	return BlockNumber(i)
}

func (n BlockNumber) Int64() int64 {
	return int64(n)
}

const (
	BadBlock      BlockNumber = 0
	EarliestBlock BlockNumber = -3
	PendingBlock  BlockNumber = -2
	LatestBlock   BlockNumber = -1
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

func (n BlockNumber) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *BlockNumber) UnmarshalText(b []byte) error {
	s := string(b)
	switch s {
	case "earliest":
		*n = EarliestBlock
		return nil
	case "pending":
		*n = PendingBlock
		return nil
	case "latest":
		*n = LatestBlock
		return nil
	default:
		var u Number
		err := u.UnmarshalText(b)
		if err != nil {
			return err
		}
		*n = BlockNumber(int64(u))
		return nil
	}
}

type Address []byte

func (a Address) String() string {
	return "0x" + hex.EncodeToString([]byte(a))
}

func (a Address) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

func (a *Address) UnmarshalText(b []byte) (err error) {
	*a, err = AddressFromString(string(b))
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

func MustAddressFromString(s string) Address {
	a, err := AddressFromString(s)
	if err != nil {
		panic(err)
	}
	return a
}

type Hash []byte

func (h Hash) String() string {
	return "0x" + hex.EncodeToString([]byte(h))
}

func (h Hash) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

func (h *Hash) UnmarshalText(b []byte) (err error) {
	*h, err = HashFromString(string(b))
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

type Data []byte

func (d Data) String() string {
	if d == nil {
		return "0x"
	}
	return "0x" + hex.EncodeToString([]byte(d))
}

func (d Data) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Data) UnmarshalText(b []byte) (err error) {
	*d, err = DataFromString(string(b))
	return err
}

func DataFromString(s string) ([]byte, error) {
	s = strings.TrimPrefix(s, "0x")
	if s == "" {
		return nil, nil
	}
	if len(s)%2 != 0 {
		s = "0" + s
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
