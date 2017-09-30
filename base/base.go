package base

import (
	//	"crypto/md5"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/KerryJava/goserver/other"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
	//	"strconv"
)

//import "github.com/KerryJava/goserver/main/"
var (
	PrepareStmt *sql.Stmt
	db          *gorm.DB = OrmDB
)

type Base struct {
}

type LoginParam struct {
	Phone  int64  `json:"phone"`
	Passwd string `json:"Passwd"`

	Log LoginLog `json:log`
}

type CheckTokenReply struct {
	Status    int    `json:"status"`
	StatusMsg string `json:"statusmsg"`
	UserData  User   `json:userdata`
}

type User struct {
	ID     int64  `json:"userid"`
	Phone  int64  `json:"phone"`
	Name   string `json:"Name"`
	Passwd string `json:"-"`
	Token  string `json:token`
}

type LoginLog struct {
	Token       string    `json:"-"`
	Userid      int64     `json:"-"`
	Channel     int32     `json:channel`
	Device      string    `json:device`
	Screen      string    `json:screen`
	Create_Time time.Time `json:"-"`
}

func (LoginLog) TableName() string {
	return "login_log"
}

func (User) TableName() string {
	return "user"
}
func (user User) Print() {
	fmt.Println("%v\n", user)
}
func init() {
	fmt.Println("base init")
	flag.Parse()
	glog.Info("init prepare statement")
	PrepareStmt, _ = DB.Prepare("SELECT id, Name, phone FROM user WHERE Name=?")
	RawQuery(true)
}
func RawQuery(isPrepare bool) {
	// Prepare statement for select data
	var stmtOut *sql.Stmt
	if isPrepare {
		stmtOut = PrepareStmt
	} else {
		stmtOut, _ = DB.Prepare("SELECT id, Name, phone FROM user WHERE Name=?")
		defer stmtOut.Close()
	}
	//defer stmtOut.Close()
	var chkErr, err error
	/*
	   chkErr := checkErr(err)
	   if chkErr != nil {
	       glog.Info(chkErr)
	   }
	*/
	// Execute the query
	rows, err := stmtOut.Query("testuser")
	var id, Name, phone []byte
	for rows.Next() {
		// Scan the value to []byte
		err = rows.Scan(&id, &Name, &phone)
		chkErr = checkErr(err)
		if chkErr != nil {
			fmt.Println(chkErr)
		}
		// Use the string value
		//fmt.Println(string(id), string(Name), string(phone))
	}
}

func (h *Base) Login(r *http.Request, param *LoginParam, reply *CheckTokenReply) error {
	other.Desc()
	var user = User{}
	//	phone, err := strconv.ParseInt(param.Phone, 10, 64)
	phone := param.Phone
	passwd := param.Passwd

	if passwd == "" {
		return ErrParams
	}
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	OrmDB.Where(&User{Phone: phone, Passwd: param.Passwd}).First(&user)
	//isSucc := false
	msg := "fail"
	if user.Phone == param.Phone {
		//	isSucc = true
		msg = "success"
	} else {
		return ErrLogin
	}

	token, _ := LoginHandler(&user)
	user.Token = token

	fmt.Println(user.Phone)
	user.Print()

	logJson, _ := json.Marshal(param.Log)

	fmt.Println("%#v", logJson)
	loginlog := new(LoginLog)

	db.FirstOrCreate(&loginlog, LoginLog{Userid: user.ID})

	_ = json.Unmarshal(logJson, loginlog)

	//	loginlog.Userid = user.ID
	loginlog.Token = user.Token
	loginlog.Create_Time = time.Now()
	//db.Updates("screen")
	db.Where(&LoginLog{Userid: user.ID}).Delete(&LoginLog{})
	db.Save(&loginlog)

	reply.Status = 1
	reply.UserData = user
	reply.StatusMsg = msg
	return nil
}

func Gentoken() string {
	/*
		crutime := time.Now().Unix()
		fmt.Println("crutime-->", crutime)

		h := md5.New()
		fmt.Println("h-->", h)

		fmt.Println("strconv.FormatInt(crutime, 10)-->", strconv.FormatInt(crutime, 10))
		io.WriteString(h, strconv.FormatInt(crutime, 10))

		fmt.Println("h-->", h)

		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("token--->", token)

		fmt.Println(len("8e1a188743c6077110da3c9778183031"))
	*/
	return ""
}

func (h *Base) Contact(r *http.Request, param *LoginParam, reply *CheckTokenReply) error {
	other.Desc()
	var user = User{}
	//	phone, err := strconv.ParseInt(param.Phone, 10, 64)
	phone := param.Phone
	passwd := param.Passwd

	if passwd == "" {
		return ErrParams
	}
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	OrmDB.Where(&User{Phone: phone, Passwd: param.Passwd}).First(&user)
	//isSucc := false
	msg := "fail"
	if user.Phone == param.Phone {
		//	isSucc = true
		msg = "success"
	} else {
		return ErrLogin
	}
	fmt.Println(user.Phone)
	user.Print()
	reply.Status = 1
	reply.UserData = user
	reply.StatusMsg = msg
	return nil
}

func checkErr(err error) error {
	if err == nil || err == sql.ErrNoRows {
		return nil
	}
	return err
}
