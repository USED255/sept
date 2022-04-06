package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	ts := GetUnixMillisTimestamp()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/ping", Ping)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	j := loadJSON(w.Body.String())
	assert.LessOrEqual(t, ts, int64(j["timestamp"].(float64)))
	assert.Equal(t, http.StatusOK, int(j["status"].(float64)))
	assert.Equal(t, "pong", j["message"])
}

func TestPing2(t *testing.T) {
	ts := GetUnixMillisTimestamp()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/ping", Ping)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", strings.NewReader(""))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	j := loadJSON(w.Body.String())
	assert.LessOrEqual(t, ts, int64(j["timestamp"].(float64)))
	assert.Equal(t, http.StatusOK, int(j["status"].(float64)))
	assert.Equal(t, "pong", j["message"])
}

func TestPingSay(t *testing.T) {
	s := randString(5)
	ts := GetUnixMillisTimestamp()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/ping", Ping)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", strings.NewReader(dumpJSON(gin.H{"say": s})))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	j := loadJSON(w.Body.String())
	assert.LessOrEqual(t, ts, int64(j["timestamp"].(float64)))
	assert.Equal(t, http.StatusOK, int(j["status"].(float64)))
	assert.Equal(t, "pong", j["message"])
	assert.Equal(t, s, j["say"])
}

func TestPingJsonError(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/ping", Ping)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", strings.NewReader("a"))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	j := loadJSON(w.Body.String())
	assert.Equal(t, http.StatusBadRequest, int(j["status"].(float64)))
	assert.Equal(t, "Invalid JSON", j["message"])
}

func loadJSON(s string) gin.H {
	var g gin.H
	json.Unmarshal([]byte(s), &g)
	return g
}

func dumpJSON(g gin.H) string {
	b, _ := json.Marshal(g)
	return string(b)
}

func randString(l int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, l)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
