package main

import (
	"time"

	"github.com/terc1997/go-cms/cmd"
	"github.com/terc1997/go-cms/internal/cache"
)

func main() {
	app := cmd.NewConfig()
	var initialData []string
	cache.NewCache(initialData, 10*time.Second)

	app.Run("localhost:8080")

}
