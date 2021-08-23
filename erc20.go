package ethrpc

import (
	"strings"

	"github.com/ethereum/go-ethereum/core/asm"
	"github.com/ethereum/go-ethereum/core/vm"
)

// 0x95d89b41 symbol()
// 0x06fdde03 name()
// 0x313ce567 decimals()

// 0x18160ddd totalSupply()
// 0x70a08231 balanceOf(address)
// 0xdd62ed3e allowance(address,address)
// 0xa9059cbb transfer(address,uint256)
// 0x095ea7b3 approve(address,uint256)
// 0x23b872dd transferFrom(address,address,uint256)

// 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925 Approval(address,address,uint256)
// 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef Transfer(address,address,uint256)

type ERC20Token struct {
	Contract Address

	Symbol   *string
	Name     *string
	Decimals *uint8

	Methods struct {
		HasTotalSupply  bool
		HasBalanceOf    bool
		HasAllowance    bool
		HasTransfer     bool
		HasApprove      bool
		HasTransferFrom bool
	}

	Events struct {
		HasTransfer bool
		HasAppoval  bool
	}
}

//
// Functions are identified by their apearance in function selector table.
// Here is an example from BAT token for balanceOf(address) entry:
//
//		DUP1
//		PUSH4 0x70a08231
//		EQ
//		PUSH2 0x0394
//		JUMPI
//
// Events are identified by a call to appropriate log function.
// Here is an example from BAT token for Transfer event:
//
//		PUSH32 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
//		SWAP3
//		DUP2
//		SWAP1
//		SUB
//		SWAP1
//		SWAP2
//		ADD
//		SWAP1
//		LOG3
//
// For brevity only PUSH4 and PUSH32 are being collected, which provides pretty
// error prone results.
//
type info struct {
	Func  [][]byte
	Event [][]byte
}

func contractInfo(code []byte) (*info, error) {
	var out info

	iter := asm.NewInstructionIterator(code)
	for iter.Next() {
		switch {
		case iter.Op() == vm.PUSH4:
			out.Func = append(out.Func, iter.Arg())
		case iter.Op() == vm.PUSH32:
			out.Event = append(out.Event, iter.Arg())
		}
	}
	if err := iter.Error(); err != nil {
		if strings.Contains(err.Error(), "incomplete push") {
			// NOTE: this error happens for some of malformed
			// contracts from time to time, so simply ignore
			// it
		} else {
			return nil, err
		}
	}
	return &out, nil
}
