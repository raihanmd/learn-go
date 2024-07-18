package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/raihanmd/dependency_injection/app"
	"github.com/raihanmd/dependency_injection/controller"
	"github.com/raihanmd/dependency_injection/helper"
	"github.com/raihanmd/dependency_injection/middleware"
	"github.com/raihanmd/dependency_injection/model/entity"
	"github.com/raihanmd/dependency_injection/repository"
	"github.com/raihanmd/dependency_injection/service"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

var router http.Handler
var db *sql.DB

func setupTestDB() *sql.DB {
	db, err := sql.Open("sqlite", "database_test.db")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestMain(m *testing.M) {
	db = setupTestDB()
	db.ExecContext(context.Background(), "TRUNCATE category")
	router = setupRouter(db)

	m.Run()
}

func TestCreateCategory(t *testing.T) {
	t.Run("Should success", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": "Gadget"}`)
		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "SECRET")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 200.0, jsonResult["code"])
		assert.Equal(t, "OK", jsonResult["message"])
		assert.NotNil(t, jsonResult["data"].(map[string]any)["id"])
		assert.Equal(t, "Gadget", jsonResult["data"].(map[string]any)["name"])
	})

	t.Run("Should failed if invalid request", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": ""}`)
		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "SECRET")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400.0, jsonResult["code"])
		assert.Equal(t, "BAD REQUEST", jsonResult["message"])
		assert.NotNil(t, jsonResult["data"])
	})

	t.Run("Should failed if invalid API KEY", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": "Gadget"}`)
		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "WRONG")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 401, response.StatusCode)
		assert.Equal(t, 401.0, jsonResult["code"])
		assert.Equal(t, "UNAUTHORIZED", jsonResult["message"])
		assert.Nil(t, jsonResult["data"])
	})
}

func TestUpdateCategory(t *testing.T) {
	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entity.Category{Name: "Gadget"})

	tx.Commit()

	t.Run("Should success", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": "New Gadget"}`)
		request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "SECRET")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 200.0, jsonResult["code"])
		assert.Equal(t, "OK", jsonResult["message"])
		assert.Equal(t, float64(category.Id), jsonResult["data"].(map[string]any)["id"])
		assert.Equal(t, "New Gadget", jsonResult["data"].(map[string]any)["name"])
	})

	t.Run("Should failed if invalid request", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": ""}`)
		request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "SECRET")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400.0, jsonResult["code"])
		assert.Equal(t, "BAD REQUEST", jsonResult["message"])
		assert.NotNil(t, jsonResult["data"])
	})

	t.Run("Should failed if invalid API KEY", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": "Gadget"}`)
		request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories", requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "WRONG")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 401, response.StatusCode)
		assert.Equal(t, 401.0, jsonResult["code"])
		assert.Equal(t, "UNAUTHORIZED", jsonResult["message"])
		assert.Nil(t, jsonResult["data"])
	})
}

func TestGetByIdCategory(t *testing.T) {
	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entity.Category{Name: "Gadget"})

	tx.Commit()

	t.Run("Should success", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
		request.Header.Add("X-API-Key", "SECRET")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 200.0, jsonResult["code"])
		assert.Equal(t, "OK", jsonResult["message"])
		assert.Equal(t, float64(category.Id), jsonResult["data"].(map[string]any)["id"])
		assert.Equal(t, category.Name, jsonResult["data"].(map[string]any)["name"])
	})

	t.Run("Should not found", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
		request.Header.Add("X-API-Key", "SECRET")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 404, response.StatusCode)
		assert.Equal(t, 404.0, jsonResult["code"])
		assert.Equal(t, "NOT FOUND", jsonResult["message"])
	})

	t.Run("Should failed if invalid API KEY", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/404", nil)
		request.Header.Add("X-API-Key", "WRONG")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 401, response.StatusCode)
		assert.Equal(t, 401.0, jsonResult["code"])
		assert.Equal(t, "UNAUTHORIZED", jsonResult["message"])
		assert.Nil(t, jsonResult["data"])
	})
}

func TestDeleteCategory(t *testing.T) {
	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entity.Category{Name: "Gadget"})

	tx.Commit()

	t.Run("Should success", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
		request.Header.Add("X-API-Key", "SECRET")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 200.0, jsonResult["code"])
		assert.Equal(t, "OK", jsonResult["message"])
	})

	t.Run("Should failed if invalid request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
		request.Header.Add("X-API-Key", "SECRET")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 404, response.StatusCode)
		assert.Equal(t, 404.0, jsonResult["code"])
		assert.Equal(t, "NOT FOUND", jsonResult["message"])
	})

	t.Run("Should failed if invalid API KEY", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories"+strconv.Itoa(category.Id), nil)
		request.Header.Add("X-API-Key", "WRONG")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 401, response.StatusCode)
		assert.Equal(t, 401.0, jsonResult["code"])
		assert.Equal(t, "UNAUTHORIZED", jsonResult["message"])
		assert.Nil(t, jsonResult["data"])
	})
}

func TestListCategory(t *testing.T) {
	db.ExecContext(context.Background(), "DELETE FROM category")

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entity.Category{Name: "Gadget"})
	category2 := categoryRepository.Create(context.Background(), tx, entity.Category{Name: "Drink"})

	tx.Commit()

	t.Run("Should success", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
		request.Header.Add("X-API-Key", "SECRET")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		fmt.Println(jsonResult)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 200.0, jsonResult["code"])
		assert.Equal(t, "OK", jsonResult["message"])
		assert.Equal(t, category.Name, jsonResult["data"].([]any)[0].(map[string]any)["name"])
		assert.Equal(t, category2.Name, jsonResult["data"].([]any)[1].(map[string]any)["name"])
	})

	t.Run("Should failed if invalid API KEY", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
		request.Header.Add("X-API-Key", "WRONG")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 401, response.StatusCode)
		assert.Equal(t, 401.0, jsonResult["code"])
		assert.Equal(t, "UNAUTHORIZED", jsonResult["message"])
		assert.Nil(t, jsonResult["data"])
	})
}
