package main

import "github.com/terc1997/go-cms/cmd"

func main() {
	server := cmd.NewConfig()

	server.Run("localhost:8080")

}
