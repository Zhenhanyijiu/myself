package main

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main22() {
	strr := "  hello  "
	out := strings.TrimSpace(strr)
	fmt.Printf("===%v\n", out)
	ff := "abc好	花花花"
	r, n := utf8.DecodeLastRune([]byte(ff))
	bl := unicode.IsSpace(r)
	fmt.Printf("===%s,,%v,error(%v)\n", string(r), bl, n)
}
func mainaa() {
	engine := gin.New()
	engine.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPathsRegexs([]string{".*"})))
	engine.Use(gin.Logger(), gin.Recovery(), func(c *gin.Context) {
		fmt.Printf("1===enter use middle ware\n")
	}, func(c *gin.Context) {
		fmt.Printf("2===enter use middle ware\n")
	})
	subEngine := engine.Group("/test")
	subEngine.Use(func(c *gin.Context) {
		fmt.Printf("7===group /test before c.Next()\n")
		//c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
		//c.Abort()
		//return
		c.Next()
		fmt.Printf("7===group /test after c.Next()\n")
	})
	subEngine.GET("/*a", func(c *gin.Context) {
		fmt.Printf("1===router /a,,%v\n", c.Param("a"))
		c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
	}, func(c *gin.Context) {
		fmt.Printf("2===router /a\n")
	})
	fmt.Printf("===>>>>>>%v\n", subEngine.BasePath())
	fmt.Printf("start lesten ........\n")
	fmt.Println(engine.Run(":8777"))
}

type testrw struct {
	pw *io.PipeWriter
	pr *io.PipeReader
}

func main() {
	//"sh","-c","./httpFaceServer -db ${dburl} -db-name ${dbname} \
	//-redis ${redisaddr} -redis-auto ${redisauto} -redis-database ${redisdb} \
	//-ea ${eaddr} -ea-expire ${eaexpire} -sa ${saddr} -admin ${adim} \
	//-wa \"${waddr}\" -et ${exptime} -rp ${rep} -file-name ${filename} \
	//-url-image-size ${urlimagesize} -url-image-expire ${urlimageexpire} \
	//-logAgent \"${lagent}\" -firewall ${fireaddr} -level ${level}"
	fmt.Printf("====\n")
}
