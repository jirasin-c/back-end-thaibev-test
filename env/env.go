package env

import "os"

var (
	PG_HOST     string = os.Getenv("PG_HOST")
	PG_NAME     string = os.Getenv("PG_NAME")
	PG_PORT     string = os.Getenv("PG_PORT")
	PG_USER     string = os.Getenv("PG_USER")
	PG_PASSWORD string = os.Getenv("PG_PASSWORD")
)
