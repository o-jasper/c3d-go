package main

import (
    "os"
    "flag"
)

// Flags
var (
    kill = flag.String("kill", "", "kill a process and die")
    downloadTorrent = flag.String("downloadTorrent", "", "download torrent from infohash and die")
    isDone = flag.String("isDone", "", "check if torrent is done")
    lookupDownloadTorrent = flag.String("lookupDownloadTorrent", "", "lookup this contract address for an infohash, using storageAt flag for sotrage address")
    storageAt = flag.String("storageAt", "", "storage address in contract")
)


func Init(){
    flag.Parse()
    if *kill != ""{
        KillPidByName(*kill)
        os.Exit(0)
    }
    if *downloadTorrent != ""{
        DownloadTorrent(*downloadTorrent)
        os.Exit(0)
    }
    if *isDone != ""{
        done := IsTorrentDone(*isDone)
        logger.Infoln("\tIs done:", done)
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

}
