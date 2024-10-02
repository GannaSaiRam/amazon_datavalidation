package main

import (
	"github.com/GannaSaiRam/data_validation/api"
)

func main() {
	server := api.StartServer(":8000")
	server.Run()
}
