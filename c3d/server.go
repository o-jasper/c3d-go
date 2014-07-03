package c3d

import (
    "github.com/ethereum/eth-go/ethpub"
    "github.com/ethereum/eth-go/ethutil"
    "net/http"
    "html/template"
    "strings"
    "log"
    "fmt"
)

type account struct{
    Addr string
    Priv []byte
    Value string
    Nonce int   
    Storage map[string]string // maps hex addrs to hex values
    Code []byte
}

type torrent struct{
    InfoHash string
    Done bool
    Contract string
}

type Session struct{
    //Accounters map[string]int //map from addresses to account numbers
   // Accounts map[int]account //map from account number to account
    Accounts []account
    AccountMap map[string]int //map from addr to account number
    Contracts []account
    Torrents []torrent
    peth *ethpub.PEthereum
}

type Config struct{
    EthPort string
    EthDataDir string
    EthLogFile string
    EthConfigFile string
    EthKeyFile string
}


var templates = template.Must(template.ParseFiles("views/index.html", "views/config.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}){
    //we already parsed the html templates
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func updateSession(s *Session){
    for i:=0; i<len(s.Accounts);i++{
        addr := (*s).Accounts[i].Addr
        state := (*s).peth.GetStateObject(addr)
        val := state.Value()
        nonce := state.Nonce()
        (*s).Accounts[i].Value = val
        (*s).Accounts[i].Nonce = nonce
    }
}

func loadConfig(peth *ethpub.PEthereum) *Config{
    conf := &Config{}
    conf.EthPort = *EthPort
    conf.EthDataDir = *EthDataDir
    conf.EthLogFile = *EthLogFile
    conf.EthConfigFile = *EthConfigFile
    conf.EthKeyFile = *KeyFile

    return conf
}

func loadSession(peth *ethpub.PEthereum) *Session {
     keyRing := ethutil.GetKeyRing()
     session := &Session{}
     session.peth = peth
     (*session).AccountMap = make(map[string]int)
     for i:=0;i<keyRing.Len();i++{
        key := keyRing.Get(i)
        addr := ethutil.Hex(key.Address())
        priv := key.PrivateKey
        state := peth.GetStateObject(addr)
        val := state.Value()
        nonce := state.Nonce()
        ac := account{Addr: addr, Value:val, Nonce:nonce, Priv:priv}
        (*session).Accounts = append((*session).Accounts, ac)
        (*session).AccountMap[addr] = i
     }
    return session
}

func (s *Session) handleTransact(w http.ResponseWriter, r *http.Request){
        to := r.FormValue("recipient")
        val := r.FormValue("amount")
        g := r.FormValue("gas")
        gp := r.FormValue("gasprice")
        from := r.FormValue("from_addr")
        acc_num := (*s).AccountMap[from]
        priv := (*s).Accounts[acc_num].Priv
        log.Println(to, val, g, gp, from, acc_num)
        p, err := (*s).peth.Transact(ethutil.Hex(priv), to, val, g, gp, "")
        if err != nil{
            log.Println(err)
        }
        log.Println(p)
        renderTemplate(w, "index", s)
}

func updateConfig(c *Config){
    //TODO
}

func (c *Config) handleConfig(w http.ResponseWriter, r *http.Request){
        updateConfig(c)
        renderTemplate(w, "config", c)
}

func (s *Session) handleIndex(w http.ResponseWriter, r *http.Request){
        if r.FormValue("reset_config") == "1"{
            // reset everything with new config :)
        }
        updateSession(s)
        renderTemplate(w, "index", s)
}

func (s *Session) serveFile(w http.ResponseWriter, r *http.Request){
    if !strings.Contains(r.URL.Path, "."){
        s.handleIndex(w, r)
    }else{
        path := fmt.Sprintf("../", r.URL.Path[1:])
        http.ServeFile(w, r, path)
    }
}

func StartServer(peth *ethpub.PEthereum){
    sesh := loadSession(peth)
    conf := loadConfig(peth)
    http.HandleFunc("/assets/", sesh.serveFile)
    http.HandleFunc("/", sesh.handleIndex)
    http.HandleFunc("/transact", sesh.handleTransact)
    http.HandleFunc("/config", conf.handleConfig)
    http.ListenAndServe(":9099", nil)
}
