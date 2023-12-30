package main

import (
	webAPIUsers "github.com/Sskrill/WebAPI-Proj/internal/pkg"
	_ "github.com/lib/pq"
)

func main() {

	db := webAPIUsers.NewDB()

	handler := webAPIUsers.NewHandler(db)

	webAPIUsers.NewRouting(handler)
}
