package other

import "fmt"
import "net/http"

//import "github.com/KerryJava/goserver/main/"

type Other struct {
}

var (
	VERSION    string
	GO_VERSION string
	BUILD_TIME string
	COMMIT     string
)

type CheckTokenParam struct {
	Token string `json:"token"`
}

type CheckTokenReply struct {
	Status    int    `json:"status"`
	StatusMsg string `json:"statusmsg"`
}

func Desc() {
	fmt.Println("Other desc")
}

func (h *Other) Version(r *http.Request, param *CheckTokenParam, reply *CheckTokenReply) error {
	s := fmt.Sprintf("%s %s %s %s", VERSION, GO_VERSION, BUILD_TIME, COMMIT)
	reply.Status = 1
	reply.StatusMsg = s
	return nil
}

func init() {
	fmt.Printf("other %s\n%s\n%s\n%s\n", VERSION, GO_VERSION, BUILD_TIME, COMMIT)
	fmt.Println("other init")
}
