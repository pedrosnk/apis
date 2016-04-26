package main

import (
	. "./db/mongo"
	. "./db/redis"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	json.NewEncoder(w).Encode(healthcheck)
}

func (s *Server) handleSchema(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		io.WriteString(w, "Saving item-schema")
		return
	}
	io.WriteString(w, "ok")
}

func (s *Server) handleCCCloud(w http.ResponseWriter, r *http.Request) {
	messageJson := map[string]string{"message": "foobar"}
	bodyJsonMapped, _ := json.Marshal(map[string]interface{}{
		"code":    CloudCode,
		"context": messageJson,
	})

	req, _ := http.NewRequest("POST", "http://localhost:5000/cc-run", bytes.NewBuffer(bodyJsonMapped))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	io.WriteString(w, string(body))
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

func (s *Server) handleCCRegistryRunner(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	io.WriteString(w, string(body))
}

func main() {
	s := &Server{
		mongodb: NewMongoDbConnection(),
		redisdb: NewRedisConnection(),
	}
	http.HandleFunc("/healthcheck", s.healthcheck)
	http.HandleFunc("/item-schemas", s.handleSchema)
	http.HandleFunc("/cmd-cc", s.handleCCCmdRunner)
	http.HandleFunc("/cloud-cc", s.handleCCCloud)
	http.HandleFunc("/registry-cc", s.handleCCRegistryRunner)
	fmt.Println("initilize App on port 8888")
	http.ListenAndServe(":8888", nil)
}
