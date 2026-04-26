package main

import "github.com/go-chi/chi/v5"

func mount() {
	r := chi.NewRouter()
}

type application struct {
	config config
	//logger
	//db driver

}

//run
//mount

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
