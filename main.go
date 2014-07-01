package main

import (
    "github.com/ethereum/eth-go"
    "github.com/ethereum/eth-go/ethutil"
    "github.com/ethereum/eth-go/ethpub"
    "github.com/ethereum/eth-go/ethlog"
    "github.com/ethereum/go-ethereum/utils"
    "fmt"
    "os"
    "log"
    "time"
    "flag"
)

// Flags
var (
    killTransmission = flag.Bool("killTransmission", false, "kill transmission and die")
    downloadTorrent = flag.String("downloadTorrent", "", "download torrent from infohash and die")
    isDone = flag.String("isDone", "", "check if torrent is done")
    lookupDownloadTorrent = flag.String("lookupDownloadTorrent", "", "lookup this contract address for an infohash, using storageAt flag for sotrage address")
    storageAt = flag.String("storageAt", "", "storage address in contract")
)


func GetInfoHashStartTorrent(peth *ethpub.PEthereum, contract_addr string, storage_addr string){
    ret := peth.GetStorage(contract_addr, storage_addr) // returns a massive base-10 integer.
    infohash := BigNumStrToHex(ret)
    log.Println("recovered infohash", infohash)
    StartTorrent(infohash)
}

// monitor for state change at addr using reactor. get info hash from contract, load into transmission
func callback(peth *ethpub.PEthereum, addr string, ethereum *eth.Ethereum){
    addr = string(ethutil.FromHex(addr))
    ch := make(chan ethutil.React, 1)
    reactor := ethereum.Reactor()
    reactor.Subscribe("object:"+addr, ch) // when the state at addr changes, this channel will receive
    for {
        _ = <- ch
        hexAddr := ethutil.Hex([]byte(addr))
        fmt.Println("hex addr", hexAddr)
        GetInfoHashStartTorrent(peth, addr, "0")
    }
}


/*
    Demonstration of simplest functionality.  
    Start everything up, stick an infohash in the blockchain, retreive it, plop into Transmission, download files over BT

*/

func EthConfig(){
    ethutil.ReadConfig(".test", "datadir", "", "c3d-go")
    utils.InitLogging("datadir", "", int(ethlog.DebugLevel), "")
}

func NewEthPEth() (*eth.Ethereum, *ethpub.PEthereum){
    // create a new ethereum node: init db, nat/upnp, ethereum struct, reactorEngine, txPool, blockChain, stateManager
    ethereum, err := eth.New(eth.CapDefault, false)
    if err != nil {
        panic(fmt.Sprintf("Could not start node: %s\n", err))
    }

    // initialize the public ethereum object. this is the interface QML gets, and it's mostly good enough for us to
    peth := ethpub.NewPEthereum(ethereum) 

    return ethereum, peth

}

func main() {
    flag.Parse()
    if *killTransmission{
        KillPidByName("transmission")
        os.Exit(0)
    }
    if *downloadTorrent != ""{
        DownloadTorrent(*downloadTorrent)
        os.Exit(0)
    }
    if *isDone != ""{
        done := IsTorrentDone(*isDone)
        log.Println("\tIs done:", done)
        os.Exit(0)
    }
    if *lookupDownloadTorrent != ""{
       if *storageAt == ""{
            *storageAt = "0"
       }
       EthConfig()
       _ , peth := NewEthPEth()
       GetInfoHashStartTorrent(peth, *lookupDownloadTorrent, *storageAt)
       os.Exit(0)
    }


    // check if transmission is running. if not, start 'er up
    CheckStartTransmission()

    // basic ethereum config.  let's put this in a big file
    EthConfig()

    ethereum, peth := NewEthPEth()

    // deal with keys :) the two genesis block keys are in keys.txt.  loadKeys will get the both for you.
    loadKeys("keys.txt")
   
    // start the node, and start mining 
    ethereum.Start(false)
    utils.StartMining(ethereum)
   
    ethereum.Port = "10101"
    ethereum.MaxPeers = 10


    keyRing := ethutil.GetKeyRing()
    priv := ethutil.Hex(keyRing.Get(0).PrivateKey)
    //addrHex := ethutil.Hex(keyRing.Get(0).Address())

    time.Sleep(time.Second)    

    //store an infohash at storage[0]
    infohash := "0x61f6beb929ffc6ccffca4e2250bb8f5edb727dd2"
    p, err := peth.Create(priv, "271", "10000", "20000000000000", "this.store[0] = " + infohash)
    if err != nil{
        log.Fatal(err)
    }
    log.Println("created contract with address", p.Address, "to store the infohash", infohash)

    /* The storage is not available until we've mined. We'll ultimately need access to the txPool
        for now, we use a callback that triggers when our contracts state changes
    */
    go callback(peth, p.Address, ethereum)
    ethereum.WaitForShutdown()

    os.Exit(0)

}
