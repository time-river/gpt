package scan

import (
	"encoding/json"
	"net/http"
	"scancenter/models"
	_"go.mongodb.org/mongo-driver/bson"
	_"go.mongodb.org/mongo-driver/mongo"
	_"go.mongodb.org/mongo-driver/mongo/options"

)

func GetTopCustomerVuls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	response := models.Response{
		Code : "2000",
		Message: "success",
		Result: "test",
	}
	json.NewEncoder(w).Encode(response)
}

