package main

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
