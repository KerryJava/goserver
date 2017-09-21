package main

import "fmt"
import "flag"

import "base"
import "net/http"

//import "log"
import "github.com/golang/glog"
import "github.com/gorilla/rpc/v2"
import "github.com/gorilla/rpc/v2/json2"

type Main struct {
}

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
)

func main() {

	flag.Parse()

	defer glog.Flush()

	fmt.Println("hello world")
	glog.Info("hello golang")

	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCustomCodec(&rpc.CompressionSelector{}), "application/json")
	s.RegisterService(new(base.Base), "")
	http.Handle("/", s)

	listenAddr := "0.0.0.0:8082"
	e := http.ListenAndServe(listenAddr, nil)

	if e != nil {
		fmt.Println(e)
	}
}

func init() {
	fmt.Printf("init ..........")
	fmt.Printf("%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION)

}
