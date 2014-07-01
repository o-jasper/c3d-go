package main

import (
    "github.com/ethereum/eth-go/ethutil"
    "io/ioutil"
    "fmt"
    "strings"
    "log"
)



func newKeyPair(){
    keyPair, err := ethutil.GenerateNewKeyPair()
    if err != nil{
        log.Println("Tragedy! Could not generating keypair! Ba-bye!")
        log.Fatal(err)
    }
    fmt.Println(keyPair)
    ethutil.GetKeyRing().Add(keyPair)
//    keyRing.NewKeyPair(keyPair.PrivateKey)
}

// private keys in plain-text hex format one per line
func loadKeys(filename string){
    keyData, err := ioutil.ReadFile(filename)
    if err != nil{
        log.Println("Could not find keys file. Creating new keypair...")        
        newKeyPair()
    } else { 
        keys := strings.Split(string(keyData), "\n")
        fmt.Println(keys)
        for _, k := range keys{
            if len(k) == 64{
                keyPair, err := ethutil.NewKeyPairFromSec(ethutil.FromHex(k))
                if err == nil{
                    ethutil.GetKeyRing().Add(keyPair)
                }
            }
        }
    }
    if ethutil.GetKeyRing().Len() == 0{
        newKeyPair()
    }
    fmt.Println(ethutil.GetKeyRing().Len())
}

