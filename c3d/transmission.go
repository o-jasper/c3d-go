package c3d

import (
    "fmt"
    "log"
    "time"
    "os/exec"
    "bytes"
    "strings"
    "net/http"
    "errors"
)


func StartTransmission(){
    cmd := exec.Command("transmission-daemon")
    err := cmd.Run()
    if err != nil{
        logger.Infoln("Couldn't start transmission...")
        log.Fatal(err)
    }
    logger.Infoln("Successfully started Transmission.  Watch it at http://localhost:9091")
}


func StartTorrentCmd(infohash string){
    logger.Infoln("Starting torrent with infohash", infohash)
    cmd := exec.Command("transmission-remote", "--add", "magnet:?xt=urn:btih:"+infohash, "--dht")
    err := cmd.Run()
    if err != nil {
        logger.Infoln("Error! Couldn't start torrent", infohash)
    } else {
        logger.Infoln("torrent download successfully started. Monitor at http://localhost:9091")
    }
}

func StartTorrent(infohash string){
    url := "http://localhost:9091/transmission/rpc/"
    link := fmt.Sprintf("magnet:?xt=urn:btih:%s",infohash)
    json := fmt.Sprintf(`{"arguments":{"filename":"%s"}, "method": "torrent-add"}`, link)
    logger.Infoln(json)
    header := make(map[string]string)
    err := http_post(url, header, json)
    if err != nil{
        logger.Infoln("Torrent start unsuccessful: ", err)
    } else {
        logger.Infoln("Successfully started torrent ", infohash, ". Monitor its progress at http://localhost:9091")
    }
}


func http_post(url string, header map[string]string, body string) error {
    b := strings.NewReader(body)
    client := &http.Client{}

    req, err := http.NewRequest("POST", url, b)
    if err != nil{
        logger.Infoln(err)
    }
    for k, v := range header{
        req.Header.Add(k, v)
    }
    resp, err := client.Do(req)
    if err != nil{
        logger.Infoln(err)
    }
    if strings.Contains(resp.Status, "409"){
        header["X-Transmission-Session-Id"] = resp.Header["X-Transmission-Session-Id"][0]
        http_post(url, header, body)
    } else if !strings.Contains(resp.Status, "200"){
        logger.Infoln("Could not connect!")
        logger.Infoln(resp)
        return errors.New(resp.Status)
    } else {
        logger.Infoln("Connection successful ", resp)
        return nil
    }
    return nil
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

func DownloadTorrent(infohash string){
    CheckStartTransmission() 
    StartTorrent(infohash)
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
