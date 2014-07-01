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
)


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
        ret := peth.GetStorage(hexAddr, "0") // returns a massive base-10 integer.
        infohash := BigNumStrToHex(ret)
        log.Println("recovered infohash", infohash)
        StartTorrent(infohash)
    }
}


/*
    Demonstration of simplest functionality.  
    Start everything up, stick an infohash in the blockchain, retreive it, plop into Transmission, download files over BT

*/

func main() {
    flag.Parse()
    if *killTransmission{
        KillPidByName("transmission")
        os.Exit(0)
    }

    // check if transmission is running. if not, start 'er up
    CheckStartTransmission()

    // basic ethereum config.  let's put this in a big file
    ethutil.ReadConfig(".test", "datadir", "", "c3d-go")
    utils.InitLogging("datadir", "", int(ethlog.DebugLevel), "")

    // create a new ethereum node: init db, nat/upnp, ethereum struct, reactorEngine, txPool, blockChain, stateManager
    ethereum, err := eth.New(eth.CapDefault, false)
    if err != nil {
        panic(fmt.Sprintf("Could not start node: %s\n", err))
    }

    // deal with keys :)
    // the two genesis block keys are in keys.txt.  loadKeys will get the both for you.
    // Let's print their balances
    loadKeys("keys.txt")

    // initialize the public ethereum object. this is the interface QML gets, and it's mostly good enough for us to
    peth := ethpub.NewPEthereum(ethereum) 

    
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
