package eth_json_rpc

import (
	"json_rpc"
)

type EthJsonRpcCpp struct {
	Addr string
}

func (eth EthJsonRpcCpp) Rpc(method string, params map[string]string) map[string] interface{} {	return json_rpc.Rpc_json(eth.Addr, method, params)
}
func (eth EthJsonRpcCpp) Rpc_(method string) map[string] interface{} {
	return json_rpc.Rpc_json(eth.Addr, method, map[string]string{})
}

//Code repeat ahead!

// Get the different things to interface.
func (eth EthJsonRpcCpp) GetKey() string {
	return eth.Rpc_("key")["result"].(string)
}
func (eth EthJsonRpcCpp) GetPeerCount() int64 {
	return int64(eth.Rpc_("peerCount")["result"].(float64))
}
func (eth EthJsonRpcCpp) GetIsMining() bool {
	return eth.Rpc_("isMining")["result"].(bool)
}
func (eth EthJsonRpcCpp) GetIsListening() bool {
	return eth.Rpc_("isListening")["result"].(bool)
}
func (eth EthJsonRpcCpp) GetCoinBase() string {
	return eth.Rpc_("coinbase")["result"].(string)
}
func (eth EthJsonRpcCpp) GetStorageAt(hexAddress string, hexStorageAddress string) string {
	return eth.Rpc("storageAt", map[string]string{"a":hexAddress, "x":hexStorageAddress})["result"].(string)
}
func (eth EthJsonRpcCpp) GetTxCountAt(hexAddress string) string {
	return eth.Rpc("txCountAt", map[string]string{"a":hexAddress})["result"].(string)
}
func (eth EthJsonRpcCpp) GetIsContractAt(hexAddress string) bool {
	return eth.Rpc("isContractAt", map[string]string{"a":hexAddress})["result"].(bool)
}
func (eth EthJsonRpcCpp) SecretToAddresss(sec string) string {
	return eth.Rpc("secretToAddress", map[string]string{"a":sec})["result"].(string)
}

func (eth EthJsonRpcCpp) Transact(seckey string, to string, value string, gas string, gasPrice string, data string) map[string] interface{} {
	return eth.Rpc("transact", map[string]string{"sec":seckey, "aDest":to, "xValue":value, "xGas":gas, "xGasPrice":gasPrice, "bData":data})
}

func (eth EthJsonRpcCpp) Create(secKey string, value string, gas string, gasPrice string, init string, hexCode string) map[string] interface{} {
	return eth.Rpc("create", map[string]string{"sec":secKey, "xEndowment":value, "xGas":gas, "xGasPrice":gas, "bCode":hexCode})
}

func (eth EthJsonRpcCpp) GetBalanceAt(addr string) string {
	return eth.Rpc("balanceAt", map[string]string{"a":addr})["result"].(string)
}

func (eth EthJsonRpcCpp) GetGasPrice(addr string) interface{} {
	return eth.Rpc_("gasPrice")["result"]
}
func (eth EthJsonRpcCpp) GetKeys() [] interface{} { //TODO reconstruct to string list?
	return eth.Rpc_("keys")["result"].([]interface{})
}


//{"method"=>"block", "params"=>{"a"=>""
//{"method"=>"check", "params"=>{"a"=>[]

//{"method"=>"keys", "params"=>nil, "returns"=>[]
//{"method"=>"lastBlock", "params"=>nil, "returns"=>{}
//{"method"=>"lll", "params"=>{"s"=>""
//"returns"=>""
//{"method"=>"procedures", "params"=>nil, "returns"=>[]
//{"method"=>"secretToAddress", "params"=>{"a"=>""
