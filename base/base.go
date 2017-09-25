package base

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/KerryJava/goserver/other"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"net/http"
	"strconv"
)

//import "github.com/KerryJava/goserver/main/"

var (
	PrepareStmt *sql.Stmt
)

type Base struct {
}

type LoginParam struct {
	Phone string `json:"phone"`
}

type CheckTokenReply struct {
	Status    int    `json:"status"`
	StatusMsg string `json:"statusmsg"`
	UserData  User   `json:userdata`
}

type User struct {
	ID    int64  `json:"userid"`
	Phone int64  `json:"phone"`
	Name  string `json:"Name"`
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

	phone, err := strconv.ParseInt(param.Phone, 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	OrmDB.Where(&User{Phone: phone}).First(&user)
	fmt.Println(user.Phone)
	user.Print()

	reply.Status = 1
	reply.UserData = user
	reply.StatusMsg = "success"

	return nil
}

func checkErr(err error) error {
	if err == nil || err == sql.ErrNoRows {
		return nil
	}

	return err
}
