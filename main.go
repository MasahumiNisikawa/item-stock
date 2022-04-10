package main

import (
	"item-stock/db"
	"item-stock/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
