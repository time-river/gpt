package hello

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"net/http"
	"scancenter/models"
	"scancenter/models/scan"
	"scancenter/repository/db"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	var req scan.GetSdkappidReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("json decode err")
		return
	}

	dbCtlClient :=db.NewDbClient()
	infos,err:=dbCtlClient.GetSdkAppidByAppid(req)
	if err != nil {
		logger.Error("db get data err")
		return
	}
	response := models.Response{
		Code : "2000",
		Message: "success",
		Result: infos,
	}

	json.NewEncoder(w).Encode(response)
}
