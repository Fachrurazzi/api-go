package test

import (
	postgres "api-go/pkg/db"
	server "api-go/pkg/http"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic.gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err.Error())
	}
	gin.SetMode(gin.ReleaseMode)
}

var (
	id string
	token string
	idQuote int
)

func setupTest(method, path, keyHeader, valueHeader string, body io.Reader) *httptest.ResponseRecorder {
	db, err := postgres.Connect("test"); if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	r := server.SetupServer(db, "test")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(keyHeader, valueHeader)
	r.ServeHTTP(w, req)
	return w
}

func TestGetQuotesRandom(t *testing.T) {
	w := setupTest("GET", "/v1/quotes", "", "", nil)
	w.Result()

	var response map[string]string
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	id = response["id"]
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegisterUserSuccess(t *testing.T) {
	body := gin.H{
		"name": "Fachrurazzi",
		"email": "testing@gmail.com",
		"password": "12345678",
	}

	s, _ := json.Marshal(body)
	b := bytes.NewBuffer(s)

	w := setupTest("POST", "/v1/user/create/", "", "", b)
	w.Result()
	assert.Equal(t, 200, w.Code)
}

func TestRegisterUserAlreadyExist(t *testing.T) {
	body := gin.H{
		"name": "Fachrurazzi",
		"email": "testing@gmail.com",
		"password": "12345678",
	}

	s, _ := json.Marshal(body)
	b := bytes.NewBuffer(s)

	w := setupTest("POST", "/v1/user/create/", "", "", b)
	w.Result()
	assert.Equal(t, 200, w.Code)
}

func TestRegisterUserInvalidEmail(t *testing.T) {
	body := gin.H{
		"name": "Fachrurazzi",
		"email": "testingemail.com",
		"password": "12345678",
	}

	s, _ := json.Marshal(body)
	b := bytes.NewBuffer(s)

	w := setupTest("POST", "/v1/user/create/", "", "", b)
	w.Result()
	assert.Equal(t, 400, w.Code)
}

func TestRegisterUserInvalidLengthPassword(t *testing.T) {
	body := gin.H{
		"name": "Fachrurazzi",
		"email": "test@gmail.com",
		"password": "123",
	}

	s, _ := json.Marshal(body)
	b := bytes.NewBuffer(s)

	w := setupTest("POST", "/v1/user/create", "", "", b)
	w.Result()
	assert.Equal(t, 400, w.Code)
}

func TestRegisterUserLogin(t *testing.T) {
	body := gin.H{
		"email": "testing@gmail.com",
		"password": "12345678",
	}

	s, _ := json.Marshal(body)
	b := bytes.NewBuffer(s)

	w := setupTest("POST", "/v1/user/login", "", "", b)
	var response map[string]string
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	token = response["token"]
	w.Result()
	assert.Equal(t, 200, w.Code)
}

func TestRegisterUserNotExist(t *testing.T) {
	body := gin.H{
		"email": "notExist@gmail.com",
		"password": "12345678",
	}

	s, _ := json.Marshal(body)
	b := bytes.NewBuffer(s)

	w := setupTest("POST", "/v1/user/login", "", "", b)
	w.Result()
	assert.Equal(t, 400, w.Code)
}

func TestRegisterUserBadCredentials(t *testing.T) {
	body := gin.H{
		"email": "testing@gmail.com",
		"password": "12345678900",
	}

	s, _ = json.Marshal(body)
	b := bytes.NewBuffer(s)

	w := setupTest("POST", "/v1/user/login", "", "", b)
	w.Result()
	assert.Equal(t, 404, w.Code)
}

func TestLovesQuote(t *testing.T) {
	w := setupTest("POST", fmt.Sprintf("/v1/favoritequotes/%s", id), "token", token, nil)
	w.Result()
	assert.Equal(t, 200, w.Code)
}

func TestGetUserQoutes(t *testing.T) {
	w := setupTest("GET", "/v1/user/quotes", "token", token, nil)
	w.Result()

	var response []map[string]int
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	idQuote = response[0]["id"]
	assert.Equal(t, 200, w.Code)
}

func TestGetAllQuotesAllUsers(t *testing.T) {
	w := setupTest("GET", "/v1/userquotes", "", "", nil)
	w.Result()
	assert.Equal(t, 200, w.Code)
}

func TestDeleteQuoteUser(t *testing.T) {
	w := setupTest("DELETE", fmt.Sprintf("/v1/deletequote/%s", strconv.Itoa(idQuote)), "token", token, nil)
	w.Result()
	assert.Equal(t, 200, w.Code)
}