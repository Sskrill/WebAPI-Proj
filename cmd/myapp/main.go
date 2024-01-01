package main

import (
	"time"

	"github.com/Sskrill/WebAPI-Proj/internal/cache"
	webAPIUsers "github.com/Sskrill/WebAPI-Proj/internal/pkg"
	_ "github.com/lib/pq"
)

func main() {
	cache := cache.NewCache(time.Second * 10)
	db := webAPIUsers.NewDB()

	handler := webAPIUsers.NewHandler(db, cache)

	webAPIUsers.NewRouting(handler)
}
