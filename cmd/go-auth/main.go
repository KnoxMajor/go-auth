package main

import (
	"github.com/knoxmajor/go-auth/api"
	"github.com/knoxmajor/go-auth/config"
	"github.com/knoxmajor/go-auth/internal/db"
)

func main() {
	config.Database = db.Connect()
	api.StartServer()
}
