package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "net/http/pprof"
	"time"
)

func print() {
	time.Sleep(8 * time.Second)
	fmt.Println("niho")
}
func gort() string {
	go print()
	return "hudi"
}

type dbServer struct {
	isOpen bool
	name   *string
	db     *sql.DB
	dbUrl  *string
	//*logrus.Log
	sAddr *string
}

func check() gin.HandlerFunc {
	return func(c *gin.Context) {
		return
	}
}

const (
	createDb       = `create database if not exists %s character set UTF8;`
	createTbMale   = `create table if not exists %s (name varchar(20) not null, age varchar(10) not null)ENGINE=InnoDB DEFAULT CHARSET=utf8;`                                                                                                         //person_male
	createTbFemale = `create table if not exists %s (num_id int unsigned not null auto_increment primary key,number varchar(25) not null, name varchar(20) not null, age varchar(10) not null,unique key(number))ENGINE=InnoDB DEFAULT CHARSET=utf8;` //person_female

	//createTbFemale = "create table if not exists `%s` (number varchar(25) not null, name varchar(20) not null, age varchar(10) not null,unique key(nember))ENGINE=InnoDB DEGAULT CHARSET=utf8;" //person_female

	createDataBase = `create database if not exists %s character set UTF8;`
	createTable    = "create table if not exists `appkey_v2`( " +
		"`appkey` varchar(65) not null, " +
		"`active` varchar(10) not null, " +
		"`action` varchar(10) not null, " +
		"`company` varchar(65) not null, " +
		"`start_time` DateTime , " +
		"`end_time` DateTime, " +
		"`limit_num` int(11) default -1," +
		"`json` json," +
		" unique key `appkey` (`appkey`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8;"
	insertTable   = `insert into %s.appkey_v2(%s) values(%s)`
	deleteTable   = `delete from %s.appkey_v2 where appkey=?`
	queryTable    = `select * from %s.appkey_v2 where appkey=?`
	queryAllTable = `select * from %s.appkey_v2 limit ?, ?`
	updateTable   = `update %s.appkey_v2 set %s where appkey=?`

	createSecketTable = "create table if not exists `appkey_secret_v2`( " +
		"`appkey_secret` varchar(115) not null, " +
		"`secret` varchar(50) not null, " +
		"`appkey` varchar(65) not null, " +
		"`status` varchar(10) not null, " +
		" unique key `appkey_secret` (`appkey_secret`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	insertSecretTable = `insert into %s.appkey_secret_v2(%s) values(%s)`
	querySecretTable  = `select * from %s.appkey_secret_v2 where appkey_secret=?`
	deleteSecretTable = `delete from %s.appkey_secret_v2 where appkey=?`
	//truncate table,, `delete from %s.appkey_secret_v2`
	deleteAllSecretTable = `truncate %s.appkey_secret_v2`
	updateSecretTable    = `update %s.appkey_secret_v2 set %s where appkey_secret=?`
	selectSecretTable    = `select * from %s.appkey_secret_v2 where appkey=?`
)

//curl 192.168.6.95:7777/db/open
func (ds *dbServer) ConnDb(c *gin.Context) {
	val := c.Query("key")
	if val != "123456" {
		c.JSON(200, gin.H{"errcode": -1, "errmsg": "key wrong"})
		return
	}
	if !ds.isOpen {
		db, err := sql.Open("mysql", *ds.dbUrl)
		if err != nil {
			c.JSON(200, gin.H{"errcode": -1, "errmsg": err.Error()})
			return
		}
		//querysql := fmt.Sprintf(createDb, ds.name)
		db.SetMaxIdleConns(100)
		db.SetMaxOpenConns(100)
		ds.db = db

		ds.isOpen = true
		c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok", "isopen": "success"})
		return
	}
	c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok", "isopen": "true"})

}

//curl 192.168.6.95:7777/db/close
func (ds *dbServer) CloseDb(c *gin.Context) {
	//if ds.isOpen == true {
	//	ds.db.Close()
	//	ds.isOpen = false
	//	c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
	//	return
	//}
	val := c.Query("key")
	if val != "123456" {
		c.JSON(200, gin.H{"errcode": -1, "errmsg": "key wrong"})
		return
	}
	if ds.db == nil {
		c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok,db not open"})
		return
	}
	err := ds.db.Ping()
	if err == nil {
		ds.db.Close()
		ds.isOpen = false
		c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok,ping,close"})
		return
	}
	//has close
	ds.isOpen = false
	c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok,close"})
}

func (ds *dbServer) CreateDb(c *gin.Context) {

}

//curl -X GET  "192.168.6.95:7777/db/create/table?tableName=student"
func (ds *dbServer) CreateTb(c *gin.Context) {
	tn := c.Query("tableName")
	fmt.Println("###tn=", tn)
	if tn == "" {
		c.JSON(200, gin.H{"errcode": -1, "errmsg": "tableName is null"})
		return
	}
	sqlquery := fmt.Sprintf(createTbFemale, tn)
	fmt.Println("###sqlquery=", sqlquery)
	if _, err := ds.db.Exec(sqlquery); err != nil {
		c.JSON(200, gin.H{"errcode": -1, "errmsg": fmt.Sprintf("%s table create failed,%s", tn, err.Error())})
		return
	}
	c.JSON(200, gin.H{"errcode": 0, "errmsg": "ok"})
}
func (ds *dbServer) Insert(c *gin.Context) {

}
func (ds *dbServer) Select(c *gin.Context) {

}

func (ds *dbServer) Update(c *gin.Context) {

}

func (ds *dbServer) Delete(c *gin.Context) {

}
func (ds *dbServer) route() {
	fmt.Println("start......")
	g := gin.Default()
	//g:=gin.New()
	g.Use(check())
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	gp := g.Group("/db")
	gp.GET("/open", ds.ConnDb)
	gp.GET("/close", ds.CloseDb)
	gp.GET("/create", ds.CreateDb)
	gp.GET("/create/table", ds.CreateTb)
	gp.GET("/insert", ds.Insert)
	gp.GET("/select", ds.Select)
	gp.GET("/update", ds.Update)
	gp.DELETE("/delete", ds.Delete)
	g.Run(*ds.sAddr)
}
func main() {
	dbs := dbServer{}
	dbs.name = flag.String("dbname", "person", "(opt)db name")
	dbs.dbUrl = flag.String("dburl", "", "(must)mysql connect url")
	dbs.sAddr = flag.String("sa", "", "(must)db lesten ip:port")
	flag.Parse()
	fmt.Printf("url=%v,sa=%v\n", *dbs.dbUrl, *dbs.sAddr)
	if *dbs.dbUrl == "" || *dbs.sAddr == "" {
		flag.Usage()
		return
	}
	dbs.isOpen = false
	db, err := sql.Open("mysql", *dbs.dbUrl)
	if err != nil {
		fmt.Println("db open failed")
		return
	}
	defer db.Close()
	if db.Ping() != nil {
		fmt.Println("db is not pong")
		return
	}

	sqlQuery := fmt.Sprintf(createDb, *dbs.name)
	_, err = db.Exec(sqlQuery)
	if err != nil {
		fmt.Println("create db failed")
		return
	}
	if _, err := db.Exec("use " + *dbs.name); err != nil {
		fmt.Println("use failed")
		return
	}
	if db.Ping() != nil {
		fmt.Println("db is not pong")
		return
	}
	fmt.Printf("use database %v ok\n", *dbs.name)
	dbs.db = db
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(time.Second * 100)

	dbs.route()
}
