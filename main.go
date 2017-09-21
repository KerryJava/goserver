package main

import "fmt"

//import "log"
//import "github.com/golang/glog"
import "github.com/gorilla/rpc/v2"
import "github.com/gorilla/rpc/v2/json2"

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
)

func main() {
	fmt.Println("hello world")
	//glog.info("hello golang")

	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCustomCodec(&rpc.CompressionSelector{}), "application/json")

}

func init() {
	fmt.Printf("init ..........")
	fmt.Printf("%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION)

}
