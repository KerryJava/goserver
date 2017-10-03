package user

import (
	"fmt"
	"github.com/KerryJava/goserver/base"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
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

type ListParam struct {
	List []Contact `json:"list"`
}

type Contact struct {
	gorm.Model // add soft delete feature

	Userid int64  `json:"-"`
	Phone  int64  `json:"phone"`
	Name   string `json:"name"`
}

func (Contact) TableName() string {
	return "contact"
}

func (h *User) ContactList(r *http.Request, param *ReqParam, reply *Reply) error {
	msg := "msg"

	userid, err := base.DecodeSign(r)
	if err != nil {
		return err
	}
	if userid == 0 {
		return base.ErrParams
	} else {
		glog.V(10).Info("userid ", userid)
	}

	reply.Contacts = make([]Contact, 0)
	//userid := int64(1100000000)
	//userid := int64(0000001100000000)
	base.OrmDB.Where(&Contact{Userid: userid}).Find(&reply.Contacts)
	reply.Status = 1
	reply.StatusMsg = msg

	return nil
}

func (h *User) UpdateContactList(r *http.Request, param *ListParam, reply *Reply) error {
	msg := "msg"

	userid, err := base.DecodeSign(r)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if userid == 0 {
		return base.ErrParams
	} else {
		glog.V(10).Info("userid ", userid)
	}

	glog.V(10).Infoln("contatlist update")
	reply.Contacts = make([]Contact, 0)

	//userid := int64(1100000000)
	//userid := int64(0000001100000000)
	//userid := int64(0000001100000000)
	base.OrmDB.Where(&Contact{Userid: userid}).Delete(&Contact{})

	for _, contact := range param.List {
		contact.Userid = userid
		base.OrmDB.Create(&contact)
	}

	reply.Status = 1
	reply.StatusMsg = msg
	reply.Contacts = param.List

	return nil
}
