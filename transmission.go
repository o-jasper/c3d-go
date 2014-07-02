package main

import (
    "log"
    "time"
    "os/exec"
    "bytes"
    "strings"
)

func KillPidByName(name string){
    /* should be cross platform
       `ps aux | grep name | awk '{print $2}' | xargs kill -9`
    */
    c1 := exec.Command("ps", "aux")
    c2 := exec.Command("grep", name)
    c3 := exec.Command("awk", "{print $2}")
    c4 := exec.Command("xargs", "kill", "-9")
    c2.Stdin, _ = c1.StdoutPipe()
    c3.Stdin, _ = c2.StdoutPipe()
    c4.Stdin, _ = c3.StdoutPipe()
    
    c1.Start()
    c2.Start()
    c3.Start()
    c4.Start()
    c1.Wait()
    c2.Wait()
    c3.Wait()
    c4.Wait()
}

func DownloadTorrent(infohash string){
    CheckStartTransmission() 
    StartTorrent(infohash)
}

func StartTransmission(){
    cmd := exec.Command("transmission-daemon")
    err := cmd.Run()
    if err != nil{
        logger.Infoln("Couldn't start transmission...")
        log.Fatal(err)
    }
    logger.Infoln("Successfully started Transmission.  Watch it at http://localhost:9091")
}

func StartTorrent(infohash string){
    logger.Infoln("Starting torrent with infohash", infohash)
    cmd := exec.Command("transmission-remote", "--add", "magnet:?xt=urn:btih:"+infohash, "--dht")
    err := cmd.Run()
    if err != nil {
        logger.Infoln("Error! Couldn't start torrent", infohash)
    } else {
        logger.Infoln("torrent download successfully started. Monitor at http://localhost:9091")
    }
}

func GetTorrentInfo(infohash string) []string{
    cmd := exec.Command("transmission-remote", "--torrent", infohash, "--info")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        logger.Infoln("Couldn't get info for", infohash)
        logger.Infoln(err)
    }
    outstr := strings.Split(out.String(), "\n")
    return outstr
}

func IsTorrentDone(infohash string) bool {
    outstr := GetTorrentInfo(infohash)
    donestr := ""
    for _, o := range(outstr){
        if strings.Contains(o, "Done"){
            donestr = o
            break
        }
    }
    logger.Infoln(donestr)
    if strings.Contains(donestr, "100"){
        return true
    }
    return false
}

func IsTransmissionRunning() bool{
    cmd := exec.Command("pgrep", "-l", "transmission")
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Run()
    return len(out.String()) > 0
}

func CheckStartTransmission(){
    if !IsTransmissionRunning(){
        StartTransmission()
    }
    time.Sleep(time.Second)
}

/*
func main() {
    KillPidByName("transmission")
    time.Sleep(time.Second)
    if !IsTransmissionRunning(){
        StartTransmission()
    }
    time.Sleep(time.Second)
    StartTorrent("61f6beb929ffc6ccffca4e2250bb8f5edb727dd2")
    
}*/
