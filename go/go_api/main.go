package main

import (
	. "./db/mongo"
	. "./db/redis"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Server struct {
	mongodb *MongoDB
	redisdb *RedisDB
}

type healthcheckRes struct {
	mongodb string `json:"mongodb"`
	redisdb string `json:"redisdb"`
}

func (s *Server) healthcheck(w http.ResponseWriter, r *http.Request) {
	healthcheck := &healthcheckRes{
		mongodb: s.mongodb.Ping(),
		redisdb: s.redisdb.Ping(),
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

func main() {
	s := &Server{
		mongodb: NewMongoDbConnection(),
		redisdb: NewRedisConnection(),
	}
	http.HandleFunc("/healthcheck", s.healthcheck)
	http.HandleFunc("/item-schemas", s.handleSchema)
	fmt.Println("initilize App on port 8888")
	http.ListenAndServe(":8888", nil)
}
