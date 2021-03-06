package main

//import "strings"

//import "log"
import (
	"flag"
	"fmt"
	"github.com/KerryJava/goserver/base"
	"github.com/KerryJava/goserver/config"
	"github.com/KerryJava/goserver/other"
	"github.com/KerryJava/goserver/user"
	"github.com/codegangsta/negroni"
	"github.com/golang/glog"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	//	"github.com/gorilla/rpc"
	//	"github.com/gorilla/rpc/json"
	"net/http"
)

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

	fmt.Println("welcome to goserver")
	glog.Info("hello golang")

	s := rpc.NewServer()
	//s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json2.NewCustomCodec(&rpc.CompressionSelector{}), "application/json")
	s.RegisterService(new(base.Base), "")
	s.RegisterService(new(other.Other), "")
	//s.RegisterInterceptFunc(SpecificMiddleware1)

	control := rpc.NewServer()
	//s.RegisterCodec(json.NewCodec(), "application/json")
	control.RegisterCodec(json2.NewCustomCodec(&rpc.CompressionSelector{}), "application/json")
	control.RegisterService(new(user.User), "")

	var common *negroni.Negroni = negroni.New()
	common.Use(negroni.HandlerFunc(base.SpecificMiddlewareSign))
	common.UseHandler(control)

	http.Handle("/", s)
	http.Handle("/control", common)
	http.HandleFunc("/hello/", sayhelloName)

	listenAddr := config.Content.ListenAddr
	e := http.ListenAndServe(listenAddr, nil)

	if e != nil {
		fmt.Println(e)
	}
}

func init() {
	fmt.Println("init main ..........")
	fmt.Printf("%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION)

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数, 默认是不会解析的
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	/*
		fmt.Println(r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
	*/
	fmt.Fprintf(w, "Hello ") //输出到客户端的信息
}

func SpecificMiddleware1(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	fmt.Printf("middleware %#v", r)
	fmt.Printf("next is %#v", next)
	next(rw, r)
	// do some stuff after

}
