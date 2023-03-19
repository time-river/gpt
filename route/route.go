package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"scancenter/api/hello"
	"scancenter/api/scan"
)

const (
	Hello  = "/hello"         //测试
	GetTopCustomerVuls = "/getTopCustomerVuls" //获取前10的用户漏洞数量
)

var Method = []string{http.MethodGet, http.MethodPost}

func SetRoute(r *mux.Router){
	r.HandleFunc(Hello,hello.Hello).Methods(Method...)
	r.HandleFunc(GetTopCustomerVuls,scan.GetTopCustomerVuls).Methods(Method...)

}
