package base

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//	"strings"
	"time"

	"crypto/rsa"
	"github.com/KerryJava/goserver/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/golang/glog"
	"strconv"
	//	"runtime"
)

const (
	// For simplicity these files are in the same folder as the app binary.
	// You shouldn't do this in production.
	privKeyPath = "app.rsa"
	pubKeyPath  = "app.rsa.pub"
)

var (
	verifyKey      *rsa.PublicKey
	signKey        *rsa.PrivateKey
	defaultKeyFunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return verifyKey, nil }
	KeyFunc        jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return verifyKey, nil }
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)

	}
}

func init() {
	initKeys()
}

func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)

}

var token *jwt.Token = jwt.New(jwt.SigningMethodRS256)
var claims jwt.MapClaims = make(jwt.MapClaims)

func LoginHandler(user *User) (string, error) {

	//fmt.Println("loginHandle")
	//token := jwt.New(jwt.SigningMethodRS256)
	//claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userid"] = user.ID
	token.Claims = claims
	tokenString, err := token.SignedString(signKey)
	//	fmt.Println(tokenString)
	if err != nil {
		fatal(err)
		return "", err
	}
	//	fmt.Println(tokenString)

	//user.Token = tokenString
	return tokenString, nil
}

func DecodeSign(r *http.Request) (int64, error) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, nil)

	if err != nil && err.Error() != "no Keyfunc was provided." {
		return 0, err
	}
	var mapClaims jwt.MapClaims = token.Claims.(jwt.MapClaims)
	//var userid int64 = mapClaims["userid"].(int64)
	glog.V(10).Infof("decode %#v", mapClaims)
	glog.V(10).Infof("decode end %#v", mapClaims["userid"])
	valStr := strconv.FormatFloat(mapClaims["userid"].(float64), 'E', -1, 64)

	var userid int64
	var val float64
	_, err = fmt.Sscanf(valStr, "%e", &val)

	if err != nil {
		return 0, err
	}

	userid = int64(val)
	//userid := int64(0000001)
	return userid, err
}

func SpecificMiddlewareSign(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	glog.V(10).Infof("sign middleware %#v\n", r)
	//fmt.Printf("sign next is %#v\n", next)
	//	name := runtime.FuncForPC(reflect.ValueOf(next).Pointer()).Name()
	//	fmt.Println(name)

	if config.Content.AuthEnable == 0 {
		next(w, r)
		return
	}

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, defaultKeyFunc)
	glog.V(10).Infof("%#v", token)

	if err == nil {
		if token.Valid {
			next(w, r)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")

		}

	} else {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")

	}
	// do some stuff after
}
