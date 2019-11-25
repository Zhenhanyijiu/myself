package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
	"websockettest/server"
)

var wsurl = "ws://192.168.6.95:7700/ws/test?appkey=zpbfrnoef722gn6q435beiebqm2cgxyhcofwi3il"

func main() {
	dial := websocket.DefaultDialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c, _, err := dial.DialContext(ctx, wsurl, nil)
	if err != nil {
		fmt.Printf("DialContext error(%v)\n", err)
		return
	}
	defer c.Close()
	//rsp.Body
	//流式发送
	for i := 0; i < 2; i++ { //
		in := server.InMsg{
			Name: "hudi1,hudi2,hudi3",
		}
		data, _ := json.Marshal(&in)
		c.WriteMessage(websocket.TextMessage, data)
	}
	c.WriteMessage(websocket.CloseNormalClosure, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	for {
		start := time.Now()
		c.SetReadDeadline(start.Add(time.Second * 2))
		msgType, out, err := c.ReadMessage()
		if err != nil {
			fmt.Printf("%v(%v)\n", err, msgType)
			return
		}
		fmt.Printf("TestMsg(%v),msgType(%v)\n", websocket.TextMessage, msgType)
		if msgType == websocket.CloseMessage {
			break
		}
		fmt.Printf("%v", string(out))
	}
}
