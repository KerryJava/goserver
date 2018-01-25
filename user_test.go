package user_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	//"github.com/KerryJava/goserver/user"
	"net/http"
	//	"net/url"
	"testing"
)

func Benchmark_Request(b *testing.B) {

	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		requestControl()
	}
}

func Benchmark_LoginRequest(b *testing.B) {

	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		login()
	}
}

func Test_HttpConnect(t *testing.T) {

	response, buf, _ := requestControl()
	t.Logf("%s", buf)
	if response.StatusCode == http.StatusOK {
		t.Log("success")
	} else {
		t.Error("response no ok")
	}

}

func requestControl() (*http.Response, string, error) {

	reqParam := make(map[string]interface{})
	reqParam["id"] = "123456"
	reqParam["jsonrpc"] = "2.0"
	reqParam["method"] = "User.ContactList"
	reqParam["params"] = make(map[string]interface{})

	bytesData, err := json.Marshal(reqParam)
	if err != nil {
		fmt.Println(err.Error())
		return nil, "", err
	}

	reader := bytes.NewReader(bytesData)

	r, err := http.NewRequest("POST", "http://localhost:8085/control", reader)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDY3ODEzNTcsImlhdCI6MTUwNjc3Nzc1NywidXNlcmlkIjoxMTAwMDAwMDAwfQ.38i1lsbrrNNU309tpbfRKBPVU2Wj07Xj3pNiq7i9BWhmR3GhhBRgQToAhKXVae5NEP6MC4-rliHnzI4xaLHLkXGmFmB2PDEIlU4wP8_Js6RnBiDNr7lLpzyG_r4vwM9KCDVB9cIyMSaGGU0H2ACDKGhg8wpKqZNYF1yhxuaB1VwHOhTyw-PlfIrlHWagT-hEyRFfCHX6yi1F6cfHs9XQ1WxGmi37VDt1eYeorWx9BTj4T7gqT1XmzOASBk-c1rHC9ow7YMVhV4TcsXGWet7nqY-NZ8xpeFUYeNH2iW1i2SslBQigdl0bHKAMJy2Uyaq3hEUeJn2OlT7JhCqpndb4iQ"
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := http.DefaultClient.Do(r)

	buf := new(bytes.Buffer)
	io.Copy(buf, response.Body)
	response.Body.Close()
	return response, buf.String(), nil
}

func login() (*http.Response, string, error) {

	reqParam := make(map[string]interface{})
	reqParam["id"] = "123456"
	reqParam["jsonrpc"] = "2.0"
	reqParam["method"] = "Base.Login"
	param := make(map[string]interface{})

	param["phone"] = "13826595953"
	param["passwd"] = "123456"

	reqParam["params"] = param

	bytesData, err := json.Marshal(reqParam)
	if err != nil {
		fmt.Println(err.Error())
		return nil, "", err
	}

	reader := bytes.NewReader(bytesData)

	r, err := http.NewRequest("POST", "http://localhost:8085/control", reader)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDY3ODEzNTcsImlhdCI6MTUwNjc3Nzc1NywidXNlcmlkIjoxMTAwMDAwMDAwfQ.38i1lsbrrNNU309tpbfRKBPVU2Wj07Xj3pNiq7i9BWhmR3GhhBRgQToAhKXVae5NEP6MC4-rliHnzI4xaLHLkXGmFmB2PDEIlU4wP8_Js6RnBiDNr7lLpzyG_r4vwM9KCDVB9cIyMSaGGU0H2ACDKGhg8wpKqZNYF1yhxuaB1VwHOhTyw-PlfIrlHWagT-hEyRFfCHX6yi1F6cfHs9XQ1WxGmi37VDt1eYeorWx9BTj4T7gqT1XmzOASBk-c1rHC9ow7YMVhV4TcsXGWet7nqY-NZ8xpeFUYeNH2iW1i2SslBQigdl0bHKAMJy2Uyaq3hEUeJn2OlT7JhCqpndb4iQ"
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := http.DefaultClient.Do(r)

	buf := new(bytes.Buffer)
	io.Copy(buf, response.Body)
	response.Body.Close()
	return response, buf.String(), nil
}
