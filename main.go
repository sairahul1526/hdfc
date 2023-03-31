package main

import (
	API "hdfc-backend/api"
	INIT "hdfc-backend/init"
)

func main() {
	// init
	INIT.Init()

	// start api server
	API.StartServer()
}
