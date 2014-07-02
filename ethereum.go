
package main

import (
    "github.com/ethereum/eth-go"
    "github.com/ethereum/eth-go/ethutil"
    "github.com/ethereum/eth-go/ethpub"
    "github.com/ethereum/eth-go/ethlog"
    "github.com/ethereum/go-ethereum/utils"
    "log"
)

//Logging
var logger *ethlog.Logger = ethlog.NewLogger("C3D")


func EthConfig() {
    ethutil.ReadConfig(".test", "datadir", "", "c3d-go")
    utils.InitLogging("datadir", "", int(ethlog.DebugLevel), "")
}

func NewEthPEth() (*eth.Ethereum, *ethpub.PEthereum){
    // create a new ethereum node: init db, nat/upnp, ethereum struct, reactorEngine, txPool, blockChain, stateManager
    ethereum, err := eth.New(eth.CapDefault, false)
    if err != nil {
        log.Fatal("Could not start node: %s\n", err)
    }
    // initialize the public ethereum object. this is the interface QML gets, and it's mostly good enough for us to
    peth := ethpub.NewPEthereum(ethereum) 
    return ethereum, peth
}

func GetInfoHashStartTorrent(peth *ethpub.PEthereum, contract_addr string, storage_addr string){
    logger.Infoln("contract addr and storage", contract_addr, storage_addr)
    ret := peth.GetStorage(contract_addr, storage_addr) // returns a massive base-10 integer.
    infohash := BigNumStrToHex(ret)
    logger.Infoln("recovered infohash", infohash)
    StartTorrent(infohash)
}

func CurrentInfo(peth *ethpub.PEthereum){
    n_peers := peth.GetPeerCount()
    peers := peth.GetPeers()
    addr := peth.GetKey().Address
    state := peth.GetStateObject(addr)
    isMin := peth.GetIsMining()
    isLis := peth.GetIsListening()
    //coinbase := peth.GetCoinBase()
    txs := peth.GetTransactionsFor(addr, true)
    //tx_count := peth.GetTxCountAt(addr)

    logger.Infoln("Summary of Current Ethereum Node State")
    logger.Infoln("N Peers: \t", n_peers)
    for _, p := range(peers){
        logger.Infoln("\t\tPeer: ", p.Ip, ":", p.Port)
    }
    logger.Infoln("Top Address on KeyRing:")
    logger.Infoln("\tAddress:\t", addr)
    logger.Infoln("\tValue:\t", state.Value())
    logger.Infoln("\tNonce:\t", state.Nonce())
    logger.Infoln("Is Mining?\t", isMin)
    logger.Infoln("Is Listening?\t", isLis)
    //logger.Infoln("Coinbase: \t", coinbase)
    logger.Infoln("Txs for\t", txs)
}

