package controllers

import (
	"BTS/helpers"
	"BTS/models"
	"BTS/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateChecklist(e echo.Context) error {
	log.Println("Starting process Create Checklist")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	/* Decode JWT */
	token := e.Request().Header.Get("Authorization")
	if token == "" || len(token) <= 7 {
		log.Println("error token is empty")
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Token is not valid",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	/* Binding Request Payload to Struct */
	var dataReq map[string]interface{}
	var req structs.Checklist
	if err := e.Bind(&req); err != nil {
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	var UserId int64
	if data, isSuccess := helpers.DecodeTokenJwt(token[7:]); !isSuccess {
		log.Println("error decode token:")
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Token is not valid",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	} else {
		UserId = int64(data["UserId"].(float64))
	}
	req.UserId = UserId

	/* convert Struct to Map */
	jsonreq, _ := json.Marshal(req)
	Jsonerr := json.Unmarshal([]byte(jsonreq), &dataReq)
	if Jsonerr != nil {
		fmt.Println("error:", Jsonerr)
	}

	/* Proses Insert Checklist */
	rsp, tx := models.CreateChecklist(dataReq)
	if tx != nil {
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest

		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Create Checklist Successfully"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}

func ListChecklist(e echo.Context) error {
	log.Println("Starting process Get List Checklist")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	/* Decode JWT */
	token := e.Request().Header.Get("Authorization")
	if token == "" || len(token) <= 7 {
		log.Println("error token is empty")
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Token is not valid",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	/* Binding Request Payload to Struct */
	var req structs.Filter
	if err := e.Bind(&req); err != nil {
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	/* Proses List Checklist */
	rsp, tx := models.ListChecklist(req)
	if tx != nil {
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest

		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Get List Checklist Successfully"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}

func DetailChecklist(e echo.Context) error {
	log.Println("Starting process Get Detail Checklist")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	/* Get ChecklistId */
	ChecklistsId, _ := strconv.Atoi(e.Param("ChecklistId"))

	/* Proses Update Checklist */
	rsp, tx := models.DetailChecklist(int64(ChecklistsId))
	if tx != nil {
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest

		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Get Detail Checklist Successfully"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}

func UpdateChecklist(e echo.Context) error {
	log.Println("Starting process Update Checklist")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	/* Decode JWT */
	token := e.Request().Header.Get("Authorization")
	if token == "" || len(token) <= 7 {
		log.Println("error token is empty")
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Token is not valid",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	/* Binding Request Payload to Struct */
	var dataReq map[string]interface{}
	var req structs.Checklist
	if err := e.Bind(&req); err != nil {
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	/* Get ChecklistId */
	ChecklistId, _ := strconv.Atoi(e.Param("ChecklistId"))
	req.ChecklistId = int64(ChecklistId)

	/* convert Struct to Map */
	jsonreq, _ := json.Marshal(req)
	Jsonerr := json.Unmarshal([]byte(jsonreq), &dataReq)
	if Jsonerr != nil {
		fmt.Println("error:", Jsonerr)
	}

	log.Println(dataReq)

	/* Proses Update Checklist */
	rsp, tx := models.UpdateChecklist(dataReq)
	if tx != nil {
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest

		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Update Checklist Successfully"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}
