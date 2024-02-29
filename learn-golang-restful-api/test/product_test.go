package test

import (
	"context"
	"encoding/json"
	"io"
	"learn-golang-restful-api/helper"
	"learn-golang-restful-api/model"
	"learn-golang-restful-api/repository"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductSuccess(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)
	productBody := strings.NewReader(`{
		"name": "Product-1",
		"description": "Description-1"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/products", productBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	helper.PanicError(err)

	assert.Equal(t, http.StatusCreated, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Product-1", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateProductFailed(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)
	productBody := strings.NewReader(`{
		"name": "",
		"description": ""
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/products", productBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	helper.PanicError(err)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestFindAllProductSuccess(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)
	ctx := context.Background()
	tx, err := db.Begin()
	helper.PanicError(err)
	productRepository := repository.NewProductRepository()
	product1 := productRepository.Save(ctx, tx, model.Product{Name: "Product-1", Description: "Description-1"})
	product2 := productRepository.Save(ctx, tx, model.Product{Name: "Product-2", Description: "Description-3"})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/products", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	helper.PanicError(err)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	products := responseBody["data"].([]interface{})

	productResponse1 := products[0].(map[string]interface{})
	productResponse2 := products[1].(map[string]interface{})

	assert.Equal(t, product1.Name, productResponse1["name"])
	assert.Equal(t, product2.Name, productResponse2["name"])
}

func TestFindAllProductEmpty(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/products", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	helper.PanicError(err)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}

func TestFindByIdProductSuccess(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)

	ctx := context.Background()
	tx, err := db.Begin()
	helper.PanicError(err)
	productRepository := repository.NewProductRepository()
	product := productRepository.Save(ctx, tx, model.Product{Name: "Product-1", Description: "Description-1"})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/products/"+strconv.Itoa(product.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	helper.PanicError(err)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, product.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestFindByIdProductNotFound(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/products/2", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	helper.PanicError(err)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
	assert.Equal(t, "product not found", responseBody["data"])
}

func TestUpdateProductSuccess(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)

	ctx := context.Background()
	tx, err := db.Begin()
	helper.PanicError(err)
	productRepository := repository.NewProductRepository()
	product := productRepository.Save(ctx, tx, model.Product{Name: "Product-1", Description: "Description-1"})
	tx.Commit()

	productBody := strings.NewReader(`{
		"name": "Product-2",
		"description": "Description-2"
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/products/"+strconv.Itoa(product.Id), productBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	helper.PanicError(err)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Product-2", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateProductNotFound(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)

	productBody := strings.NewReader(`{
		"name": "Product-1",
		"description": "Description-1"
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/products/20", productBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	helper.PanicError(err)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestUpdateProductBadRequest(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)

	ctx := context.Background()
	tx, err := db.Begin()
	helper.PanicError(err)
	productRepository := repository.NewProductRepository()
	product := productRepository.Save(ctx, tx, model.Product{Name: "Product-1", Description: "Description-1"})
	tx.Commit()

	productBody := strings.NewReader(`{
		"name": "",
		"description": ""
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/products/"+strconv.Itoa(product.Id), productBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	helper.PanicError(err)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestDeleteProductSuccess(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)

	ctx := context.Background()
	tx, err := db.Begin()
	helper.PanicError(err)
	productRepository := repository.NewProductRepository()
	product := productRepository.Save(ctx, tx, model.Product{Name: "Product-1", Description: "Description-1"})
	tx.Commit()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/products/"+strconv.Itoa(product.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	helper.PanicError(err)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteProductNotFound(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/products/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTHSECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	helper.PanicError(err)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestUnauthorized(t *testing.T) {
	db := SetupDbTest()
	TruncateDB(db)
	router := SetupRouter(db)

	productBody := strings.NewReader(`{
		"name": "Product-1",
		"description": "Description-1"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/products", productBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "AUTH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	helper.PanicError(err)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}
