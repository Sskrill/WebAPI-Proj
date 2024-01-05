package main

import (
	"time"

	webAPIUsers "github.com/Sskrill/WebAPI-Proj/internal/pkg"
	"github.com/Sskrill/WebAPI-Proj/pkg/cache"
	_ "github.com/lib/pq"
)

func main() {

	cache := cache.NewCache(time.Second * 10)
	db := webAPIUsers.NewDB()

	handler := webAPIUsers.NewHandler(db, cache)

	webAPIUsers.NewRouting(handler)
}
