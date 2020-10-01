/*
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddInsight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetInsightById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	registerHandler := new(handlers.RegisterHandler)
	body, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Error(err)
		json.NewEncoder(w).Encode(models.ApiResponse{Code: 400, Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var model = models.RegisterRequestModel{}
	parseError := json.Unmarshal(body, interface{}(&model))
	if parseError!= nil {
		log.Info(parseError)
		json.NewEncoder(w).Encode(models.ApiResponse{Code: 400, Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := registerHandler.RegisterUserCodeServer(model)
	if err != nil {
		log.Error(err)
		json.NewEncoder(w).Encode(models.ApiResponse{Code: 400, Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(models.ApiResponse{Code: 200, Message: response})
	w.WriteHeader(http.StatusOK)
}

