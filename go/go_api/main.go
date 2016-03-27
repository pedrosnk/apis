package main

import (
	. "./db/mongo"
	. "./db/redis"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

var CloudCode string = `function() {
		var leftpad = (str, len) => {
			str = String(str)
			var i = -1
			len = len - str.length
			while (++i < len) {
				str = '#' + str
			}
			return str
		}
		return leftpad(message, 10)
}()`

type Server struct {
	mongodb *MongoDB
	redisdb *RedisDB
}

type healthcheckRes struct {
	Mongodb string `json:"mongodb"`
	Redisdb string `json:"redisdb"`
}

func (s *Server) healthcheck(w http.ResponseWriter, r *http.Request) {
	pingRedis := make(chan string, 1)
	pingMongo := make(chan string, 1)
	go func() {
		pingRedis <- s.redisdb.Ping()
	}()
	go func() {
		pingMongo <- s.mongodb.Ping()
	}()

	healthcheck := &healthcheckRes{
		Mongodb: <-pingMongo,
		Redisdb: <-pingRedis,
	}
	hcString, _ := json.Marshal(healthcheck)
	io.WriteString(w, string(hcString))
}

func (s *Server) handleSchema(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		io.WriteString(w, "Saving item-schema")
		return
	}
	io.WriteString(w, "ok")
}

func (s *Server) handleCCCmdRunner(w http.ResponseWriter, r *http.Request) {
	messageJson := map[string]string{"message": "foobar"}
	codeContext, _ := json.Marshal(messageJson)
	cmd := exec.Command("node",
		"../cc_runner/runner.js",
		CloudCode,
		string(codeContext),
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("error running cloud code")
	}
	io.WriteString(w, out.String())
}

func main() {
	s := &Server{
		mongodb: NewMongoDbConnection(),
		redisdb: NewRedisConnection(),
	}
	http.HandleFunc("/healthcheck", s.healthcheck)
	http.HandleFunc("/item-schemas", s.handleSchema)
	http.HandleFunc("/run-cc", s.handleCCCmdRunner)
	fmt.Println("initilize App on port 8888")
	http.ListenAndServe(":8888", nil)
}
