package config

import "fmt"
import "github.com/KerryJava/configor"

type ContentType struct {
	APPName      string `default:"app name"`
	ListenAddr   string `default:"0.0.0.0:80"`
	DSN          string `default:"app:appapp@tcp(0.0.0.0:3306)/golden"`
	MaxIdleConns int    `default:100`
	MaxOpenConns int    `default:100`
}

var Content ContentType

func init() {
	configor.Load(&Content, "config.yml")
	fmt.Printf("config: %#v", Content)
	fmt.Println("")
	fmt.Println("config init")
}
