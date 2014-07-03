package main

import (
    "github.com/ethereum/eth-go"
    "github.com/ethereum/eth-go/ethutil"
    "github.com/ethereum/eth-go/ethpub"
    "github.com/ethereum/go-ethereum/utils"
    "os"
    "log"
    "time"
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
        logger.Infoln("hex addr ", hexAddr)
        GetInfoHashStartTorrent(peth, hexAddr, "0")
    }
}

/*
    Demonstration of simplest functionality.  
    Start everything up, stick an infohash in the blockchain, retreive it, plop into Transmission, download files over BT

*/
func main() {
    // parse flags.
    Init()

    // check if transmission is running. if not, start 'er up
    CheckStartTransmission()

    // basic ethereum config.  let's put this in a big file
    EthConfig()

    ethereum, peth := NewEthPEth()
    ethereum.Port = "10101"
    ethereum.MaxPeers = 10

    //start the node
    ethereum.Start(false)

    // deal with keys :) the two genesis block keys are in keys.txt.  loadKeys will get them both for you.
    // if there are more keys, having 0 balance, funds will be transfered to them
    loadKeys(*keyFile)

    // start mining
    utils.StartMining(ethereum)

    // checks if any addrs have 0 balance, tops them up
    checkZeroBalance(peth)

   
    keyRing := ethutil.GetKeyRing()
    priv := ethutil.Hex(keyRing.Get(0).PrivateKey)
    //addrHex := ethutil.Hex(keyRing.Get(0).Address())

    //time.Sleep(time.Second*10)    

    //store an infohash at storage[0]
    infohash := "0x1183596810fbca83fce8e12d98234aaaf38eb7cd"
    p, err := peth.Create(priv, "271", "2000", "1000000", "this.store[0] = " + infohash)
    if err != nil{
        log.Fatal(err)
    }
    logger.Infoln("created contract with address ", p.Address, " to store the infohash ", infohash)
    time.Sleep(time.Second)
    CurrentInfo(peth)

    /* The storage is not available until we've mined. We'll ultimately need access to the txPool
        for now, we use a callback that triggers when our contracts state changes
    */
    go callback(peth, p.Address, ethereum)
    ethereum.WaitForShutdown()

    os.Exit(0)

}
