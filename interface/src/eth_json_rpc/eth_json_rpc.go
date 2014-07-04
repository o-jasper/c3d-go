package eth_json_rpc

import (
	"json_rpc"
	"github.com/ethereum/eth-go/ethpub"
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

func pkey_from_privkey(privkey string) ethpub.PKey {  //TODO... how to get the pubkey?
	return ethpub.PKey{PrivateKey:privkey, Address:eth.SecretToAddress(privkey)}
}

// Get the different things to interface.

func (eth EthJsonRpcCpp) GetKey() ethpub.PKey {
	return pkey_from_privkey(Rpc_("key")["result"].(string))
}
func (eth EthJsonRpcCpp) GetKeys() [] ethpub.PKey { //TODO reconstruct to string list?
	list := eth.Rpc_("keys")["result"].([]interface{})
	ret_list := [] ethpub.PKey{}
	for i := range list {
		ret_list = append(ret_list, pkey_from_privkey(list[i].(string)))
	}
	return ret_list
}

func (eth EthJsonRpcCpp) GetPeerCount() int {
	return int(eth.Rpc_("peerCount")["result"].(float64))
}
func (eth EthJsonRpcCpp) GetPeers() []ethpub.PPeer {
	return []ethpub.PPeer{ethpub.PPeer{Ip:"Peer getting not supported"}}  // TODO not supported.
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
func (eth EthJsonRpcCpp) GetStorage(hexAddress string, hexStorageAddress string) string {
	return eth.Rpc("storageAt", map[string]string{"a":hexAddress, "x":hexStorageAddress})["result"].(string)
}
func (eth EthJsonRpcCpp) GetTxCountAt(hexAddress string) string {
	return eth.Rpc("txCountAt", map[string]string{"a":hexAddress})["result"].(string)
}
func (eth EthJsonRpcCpp) GetIsContract(hexAddress string) bool {
	return eth.Rpc("isContractAt", map[string]string{"a":hexAddress})["result"].(bool)
}
func (eth EthJsonRpcCpp) SecretToAddress(sec string) string {
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


//{"method"=>"block", "params"=>{"a"=>""
//{"method"=>"check", "params"=>{"a"=>[]

//{"method"=>"keys", "params"=>nil, "returns"=>[]
//{"method"=>"lastBlock", "params"=>nil, "returns"=>{}
//{"method"=>"lll", "params"=>{"s"=>""
//"returns"=>""
//{"method"=>"procedures", "params"=>nil, "returns"=>[]
//{"method"=>"secretToAddress", "params"=>{"a"=>""
