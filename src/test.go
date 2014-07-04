package main

import (
//	"flag"
	"fmt"
	"github.com/project-douglas/c3d-go/src/eth_json_rpc"
)


func main() {

//	fmt.Println(eth_json_rpc.Rpc_json("http://localhost:9090/", "procedures", map[string]string{}))
	rpc := eth_json_rpc.EthJsonRpcCpp{Addr:"http://localhost:9090/"}
	fmt.Println(rpc.GetKey())
	fmt.Println(rpc.GetKeys())
	fmt.Println(rpc.GetPeerCount())
	fmt.Println(rpc.GetPeers())
	fmt.Println(rpc.GetIsMining())
	fmt.Println(rpc.GetKeys())
	fmt.Println(rpc.GetCoinBase())

	//Note arbitrary and not guaranteed to be online.
	addr := "0x8d863d4cdd41c9ad1c6d01f961bd16590632cd68"
	fmt.Println(rpc.GetGasPrice(addr))
	fmt.Println("With args")
	fmt.Println(rpc.GetStorage(addr, "0x20"))
	fmt.Println(rpc.GetTxCountAt(addr))
	fmt.Println(rpc.GetIsContract(addr))
	fmt.Println(rpc.GetBalanceAt(addr))
}
