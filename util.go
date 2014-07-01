package main

import (
    "github.com/ethereum/eth-go/ethutil"
)


func BigNumStrToHex(s string) string{
    bignum := ethutil.Big(s)
    bignum_bytes := ethutil.BigToBytes(bignum, 16)
    return ethutil.Hex(bignum_bytes)
}
