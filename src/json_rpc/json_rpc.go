package json_rpc

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"encoding/json"
)

//Note: encoding/json library was not clear enough to me when i made it.
// but it seems good anyway, so TODO: replace this.
func Mk_json_params(params map[string]string) string {

	first, param_str := true, ""
	for k,v := range params {
		if first {
			first = false
			param_str = fmt.Sprintf("%s\"%s\":\"%s\"", param_str, k, v)
		} else {
			param_str = fmt.Sprintf("%s, \"%s\":\"%s\"",param_str, k, v)
		}
	}
	return param_str
}

func Mk_json(method string, params map[string]string) string {
	return fmt.Sprintf("{\"method\":\"%s\", \"params\":{%s}, \"jsonrpc\":\"2.0\", \"id\":\"eth_go_interface_rpc\"}", method, Mk_json_params(params))
}

const wait_time = 4000000  //four milliseconds.
const interval_time = 1000000 //1 millisecond.

func Rpc_json(addr string, method string, params map[string]string) map[string] interface{} {
	reader := strings.NewReader(Mk_json(method, params))
	start := time.Now().UnixNano()
	for time.Now().UnixNano() < start + wait_time {
		resp, err := http.Post(addr, "application/json", reader)
		if err == nil {
			defer resp.Body.Close()

			into := map[string] interface{}{}
			decoder := json.NewDecoder(resp.Body)
			
			err = decoder.Decode(&into)
//			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Some err.. ", into, " err: ", err)
			}
			return into
		} else {
			//fmt.Println("Some err.. ", addr, " err: ", err)
		}
		time.Sleep(interval_time)
	}
	return map[string] interface{}{}
}

func Rpc_json_(addr string, method string) map[string] interface{} {
	return Rpc_json(addr, method, map[string]string{})
}
