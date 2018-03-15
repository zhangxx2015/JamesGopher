package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func proc(w http.ResponseWriter, req *http.Request) {
	//args := make(map[string]interface{})

	bdBytes, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	body := string(bdBytes)
	fmt.Println("body:", body)
	//args["body"] = body

	//	queryForm, err := url.ParseQuery(req.URL.RawQuery)
	//	if err == nil {
	//		for k, v := range queryForm {
	//			fmt.Printf("Query:k=%v, v=%v\n", k, v)
	//		}
	//	}

	req.ParseForm()
	//for k, v := range req.Form {
	//	fmt.Printf("Form:k=%v, v=%v\n", k, v)
	//	args[k] = v
	//}
	//bs, _ := json.Marshal(args)
	//_json := string(bs)
	//fmt.Println("bs:", _json)
	bs, _ := json.Marshal(req.Form)
	_json := string(bs)

	path := req.URL.EscapedPath()
	//fmt.Println(strings.HasPrefix("my string", "prefix"))  // false
	//fmt.Println(strings.HasPrefix("my string", "my"))      // true
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	page := "404.php"
	if len(path) == 0 {
		path = "index.php"
	}
	if PathExists(path) {
		page = path
	}
	fmt.Println("path:", path)

	st := time.Now()

	cmd := exec.Command("php", page, "form="+_json+"&body="+body)
	f, _ := cmd.Output()

	fmt.Printf("elapse:%f seconds\r\n", time.Now().Sub(st).Seconds())
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("X-Powered-By", "Fools' web server")
	w.Write(f)
}
func main() {
	scheme := "http://"
	host := "127.0.0.1"
	port := "80"
	route := "/info.php?id=abc&age=35"
	url := scheme + host + ":" + port + route
	go func() {
		time.Sleep(time.Microsecond * 300)
		//resp, _ := http.Get(url)
		//resp, _ := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("nick=fools"))

		json := `{"hobby":"Fishing", "gender": true}`
		req := bytes.NewBuffer([]byte(json))
		resp, _ := http.Post(url, "application/json;charset=utf-8", req)

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Printf("# request [%s]:%s\r\n", url, string(body))
	}()
	fmt.Println("start service is succeed")
	http.HandleFunc("/", proc)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
