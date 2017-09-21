package base

import "fmt"
import "net/http"
import "goserver/main"

type Base struct {
}

type CheckTokenParam struct {
	Token string `json:"token"`
}

type CheckTokenReply struct {
	Status    int    `json:"status"`
	StatusMsg string `json:"statusmsg"`
}

func (h *Base) Login(r *http.Request, param *CheckTokenParam, reply *CheckTokenReply) error {
	reply.Status = 1
	reply.StatusMsg = "pass"
	return nil
}

func init() {
	fmt.Println("base init")
}
