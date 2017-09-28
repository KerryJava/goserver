package user

import (
	"github.com/KerryJava/goserver/base"
	"net/http"
)

type User struct {
}

type ReqParam struct {
	phone int64 `json:"phone"`
}

type Reply struct {
	Status    int       `json:"status"`
	StatusMsg string    `json:"statusmsg"`
	Contacts  []Contact `json:"contacts"`
}

type Contact struct {
	Userid int64  `json:"-"`
	Phone  int64  `json:"phone"`
	Name   string `json:"Name"`
}

func (Contact) TableName() string {
	return "contact"
}
func (h *User) ContactList(r *http.Request, param *ReqParam, reply *Reply) error {
	msg := "msg"

	reply.Contacts = make([]Contact, 0)
	userid := int64(1100000000)
	//userid := int64(0000001100000000)
	base.OrmDB.Where(&Contact{Userid: userid}).Find(&reply.Contacts)
	reply.Status = 1
	reply.StatusMsg = msg

	return nil
}
