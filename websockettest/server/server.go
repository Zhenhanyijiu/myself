package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"time"
)

var upgrade = websocket.Upgrader{}

type serverWeb struct {
	sa string
}

type InMsg struct {
	Name string
}

type OutMsg struct {
	ErrCode int
	ErrMsg  string
	Great   string
}

func NewServerWeb(sa string) *serverWeb {
	return &serverWeb{
		sa: sa,
	}
}
func (s *serverWeb) GetTest(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	//var upgrader = websocket.Upgrader{
	//	// 解决跨域问题
	//	CheckOrigin: func(r *http.Request) bool {
	//		return true
	//	},
	//}
	upgrade.CheckOrigin = func(r *http.Request) bool {
		if r.Host != "192.168.5.25" { //限制跨域
			return false
		}
		return true
	}
	c, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("ws upgrade error(%v)\n", err)
		return
	}
	defer c.Close()
	//defer c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	defer c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "nomal close"))
	fmt.Printf("FormatCloseMessage::[]byte(%x)\n", websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	/*
		in := InMsg{}
		if err := ctx.ShouldBindJSON(&in); err != nil {
			fmt.Printf("req is error(%v)\n", err)
			return
		}
	*/
	//read message
	ak := ctx.Query("appkey")
	fmt.Printf("====appkey(%v)\n", ak)

	ins := []InMsg{}
	for {
		start := time.Now()
		c.SetReadDeadline(start.Add(time.Second * 2))
		_, msg, err := c.ReadMessage()

		//最好从客户端发送个eof告知服务端结束
		if err != nil {
			fmt.Printf(">>>>>>>>%v\n", err)
			//out := OutMsg{ErrCode: 707, ErrMsg: err.Error()}
			//jsonB, _ := json.Marshal(&out)
			//if err := c.WriteMessage(websocket.TextMessage, jsonB); err != nil {
			//	fmt.Printf("WriteMessage:error(%v)\n", err)
			//}
			break
		}
		in := InMsg{}
		if err := json.Unmarshal(msg, &in); err != nil {
			out := OutMsg{ErrCode: 707, ErrMsg: err.Error()}
			jsonB, _ := json.Marshal(&out)
			if err := c.WriteMessage(websocket.TextMessage, jsonB); err != nil {
				fmt.Printf("WriteMessage:error(%v)\n", err)
			}
			return
		}
		ins = append(ins, in)
	}

	for _, in := range ins {
		names := strings.Split(in.Name, ",")
		for i, v := range names {
			if v != "" {
				out := OutMsg{
					ErrCode: 0,
					ErrMsg:  "ok",
					Great:   fmt.Sprintf("Greating hello %v (%v)", v, i+1),
				}
				jsonB, _ := json.Marshal(&out)
				if err := c.WriteMessage(websocket.TextMessage, jsonB); err != nil {
					fmt.Printf("WriteMessage:error(%v)\n", err)
					return
				}
			}
		}
	}

}

func (s *serverWeb) Router() {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	grp := engine.Group("/ws")
	grp.GET("/test", s.GetTest)
	engine.Run(s.sa)
}
