package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-rest/app"
	"go-rest/controller"
	"go-rest/entity/domain"
	"go-rest/helper"
	"go-rest/middleware"
	"go-rest/repository"
	"go-rest/service"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/gorest_test")
	helper.CheckError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	hewanRepository := repository.NewHewanRepository()
	hewanService := service.NewHewanService(hewanRepository, db, validate)
	hewanController := controller.NewHewanController(hewanService)

	router := app.NewRouter(hewanController)

	return middleware.NewAuthMiddleware(router)
}

func truncateHewan(db *sql.DB) {
	db.Exec("TRUNCATE hewan")
}

func TestCreateHewanSuccess(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Biawak"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:6969/api/hewan", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", os.Getenv("APIKey"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 200, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])
	assert.Equal(t, "Biawak", resBody["data"].(map[string]interface{})["name"])
}

func TestCreateHewanFailed(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:6969/api/hewan", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", os.Getenv("APIKey"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 400, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 400, int(resBody["code"].(float64)))
	assert.Equal(t, "Bad Request", resBody["status"])
}

func TestUpdateHewanSuccess(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)

	tx, _ := db.Begin()
	hewanRepository := repository.NewHewanRepository()
	hewan := hewanRepository.Save(context.Background(), tx, domain.Hewan{
		Name: "Biawak",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Anjing"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:6969/api/hewan/"+strconv.Itoa(hewan.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", os.Getenv("APIKey"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 200, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])
	assert.Equal(t, hewan.Id, int(resBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Anjing", resBody["data"].(map[string]interface{})["name"])
}

func TestUpdateHewanFailed(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)

	tx, _ := db.Begin()
	hewanRepository := repository.NewHewanRepository()
	hewan := hewanRepository.Save(context.Background(), tx, domain.Hewan{
		Name: "Biawak",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:6969/api/hewan/"+strconv.Itoa(hewan.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", os.Getenv("APIKey"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 400, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 400, int(resBody["code"].(float64)))
	assert.Equal(t, "Bad Request", resBody["status"])
}

func TestGetHewanByIdSuccess(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)

	tx, _ := db.Begin()
	hewanRepository := repository.NewHewanRepository()
	hewan := hewanRepository.Save(context.Background(), tx, domain.Hewan{
		Name: "Biawak",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:6969/api/hewan/"+strconv.Itoa(hewan.Id), nil)
	request.Header.Add("X-API-KEY", os.Getenv("APIKey"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 200, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])
	assert.Equal(t, hewan.Id, int(resBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, hewan.Name, resBody["data"].(map[string]interface{})["name"])
}

func TestGetHewanByIdFailed(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:6969/api/hewan/10", nil)
	request.Header.Add("X-API-KEY", os.Getenv("APIKey"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 404, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 404, int(resBody["code"].(float64)))
	assert.Equal(t, "Not Found", resBody["status"])
}

func TestDeleteHewanSuccess(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)

	tx, _ := db.Begin()
	hewanRepository := repository.NewHewanRepository()
	hewan := hewanRepository.Save(context.Background(), tx, domain.Hewan{
		Name: "Biawak",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:6969/api/hewan/"+strconv.Itoa(hewan.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", os.Getenv("APIKey"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 200, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])
}

func TestDeleteHewanFailed(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:6969/api/hewan/10", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", os.Getenv("APIKey"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 404, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 404, int(resBody["code"].(float64)))
	assert.Equal(t, "Not Found", resBody["status"])
}

func TestGetAllHewanSuccess(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)

	tx, _ := db.Begin()
	hewanRepository := repository.NewHewanRepository()
	hewan1 := hewanRepository.Save(context.Background(), tx, domain.Hewan{
		Name: "Biawak",
	})
	hewan2 := hewanRepository.Save(context.Background(), tx, domain.Hewan{
		Name: "Anjing",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:6969/api/hewan", nil)
	request.Header.Add("X-API-KEY", os.Getenv("APIKey"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 200, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])

	semua_hewan := resBody["data"].([]interface{})
	hewan1Response := semua_hewan[0].(map[string]interface{})
	hewan2Response := semua_hewan[1].(map[string]interface{})

	assert.Equal(t, hewan1.Id, int(hewan1Response["id"].(float64)))
	assert.Equal(t, hewan1.Name, hewan1Response["name"])

	assert.Equal(t, hewan2.Id, int(hewan2Response["id"].(float64)))
	assert.Equal(t, hewan2.Name, hewan2Response["name"])
}

func TestUnauthrized(t *testing.T) {
	db := setupTestDB()
	truncateHewan(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:6969/api/hewan", nil)
	request.Header.Add("X-API-KEY", "wrong api key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	assert.Equal(t, 401, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)
	fmt.Println(resBody)

	assert.Equal(t, 401, int(resBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", resBody["status"])
}
