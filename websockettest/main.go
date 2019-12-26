package main

import (
	"fmt"
	"myself/websockettest/server"
)

func main() {
	sa := "192.168.6.95:7700"
	ws := server.NewServerWeb(sa)
	fmt.Printf("start ws(%v) ...\n", sa)
	ws.Router()
}
