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
