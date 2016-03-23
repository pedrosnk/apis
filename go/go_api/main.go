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

func main() {
	s := &Server{
		mongodb: NewMongoDbConnection(),
		redisdb: NewRedisConnection(),
	}
	http.HandleFunc("/healtchcheck", s.healthcheck)
	fmt.Println("initilize App on port 8888")
	http.ListenAndServe(":8888", nil)
}
