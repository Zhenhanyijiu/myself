package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

const (
	dbUrl       = "root:root@tcp(192.168.5.30:3306)/device_activate?charset=utf8&parseTime=True&loc=Local"
	dbName      = "device_activate"
	tableName   = "device_info"
	createData  = `create database if not exists %s character set UTF8;`
	createTable = `create table if not exists %s (
device_id varchar(100) default "",
active int unsigned default 0,
engine_type varchar(30) default "",
raw_access varchar(2048) default "",
time_start  varchar(30) default "",
time_end  varchar(30) default "",
time_stamp varchar(30) default "",
max_count int unsigned default 0,
primary key (device_id,engine_type)
) engine=innodb default charset=utf8`
	insertSql = `insert into %s.%s(%s) values(%s)`
)

//user:pwd@tcp(192.168.5.30:3306)
type DBParam struct {
	dbName      string
	tableName   string
	createData  string
	createTable string
	db          *sql.DB
}

func newDBParam() *DBParam {
	return &DBParam{
		dbName:      dbName,
		tableName:   tableName,
		createData:  fmt.Sprintf(createData, dbName),
		createTable: fmt.Sprintf(createTable, tableName),
	}
}

type Data struct {
	DeviceID   string `gorm:"column:device_id;type:varchar(100);default:'';primary_key;comment:'DeviceID'"`
	Active     int    `gorm:"column:active;type:int unsigned;default:0;"`
	EngineType string `gorm:"column:engine_type;type:varchar(30);default:'';primary_key;comment:'EngineType'"`
	RawAccess  string `gorm:"column:raw_access;type:varchar(2048);default:'';comment:'RawAccess'"`
	TimeStart  string `gorm:"column:time_start;type:varchar(30);default:'';comment:'TimeStart'"`
	TimeEnd    string `gorm:"column:time_end;type:varchar(30);default:'';comment:'TimeEnd';"`
	TimeStamp  string `gorm:"column:time_stamp;type:varchar(30);default:'';comment:'TimeStamp'"`
	MaxCount   int    `gorm:"column:max_count;type:int unsigned;default:0;comment:'MaxCount'"`
	//Test       int    `gorm:"column:test;type:int unsigned;default:0;comment:'Test'"`
}

func createTableFunc(db *gorm.DB, models interface{}) error {
	return db.CreateTable(models).Error
}
func insertFunc(db *gorm.DB, value interface{}) error {
	return db.Create(value).Error
}
func insertTest(dba *gorm.DB) error {
	value1 := Data{
		DeviceID:   "111",
		EngineType: "e",
		TimeStamp:  time.Now().Format("2006-01-02 15:04:05"),
	}
	value2 := Data{
		DeviceID:   "111",
		EngineType: "f",
		TimeStamp:  time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := insertFunc(dba, &value1); err != nil {
		fmt.Printf("insertFunc error(%v)\n", err)
		return err
	}
	if err := insertFunc(dba, &value2); err != nil {
		fmt.Printf("insertFunc error(%v)\n", err)
		return err
	}
	return nil
}
func update1(dba *gorm.DB) error {
	var d Data
	if err := dba.Where("device_id=? and engine_type=?", "116", "e").First(&d).Error; err != nil {
		fmt.Printf("upate error(%v)\n", err)
		return err
	}
	d.MaxCount = 73
	if err := dba.Table(tableName).Save(&d).Error; err != nil {
		fmt.Printf("upate save error(%v)\n", err)
		return err
	}
	fmt.Printf("data:%v\n", d)
	return nil
}
func main() {
	//fmt.Printf("gorm\n")
	dba, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
		return
	}
	defer dba.Close()
	if err := dba.DB().Ping(); err != nil {
		fmt.Printf("error(%v)\n", err)
		return
	}

	/*
		//create table
		if err := createTableFunc(dba, &Data{}); err != nil {
			fmt.Printf("createTableFunc error(%v)\n", err)
			return
		}
		fmt.Printf("create table ok...\n")
	*/
	dba = dba.Table(tableName)
	//insert data
	//if err := insertTest(dba); err != nil {
	//	return
	//}
	fmt.Printf("insert data ok...\n")

	//delete data operation
	if err := dba.Where("device_id=? and engine_type=?", "116", "ee").Delete(&Data{}).Error; err != nil {
		fmt.Printf("delete error(%v)\n", err)
		return
	}
	fmt.Printf("delete data ok...\n")
	//update data
	if err := update1(dba); err != nil {
		return
	}
	////v := Data{
	////	DeviceID:   "115",
	////	EngineType: "f",
	////	TimeStamp:  time.Now().Format("2006-01-02 15:04:05"),
	////}
	////dba.Set()
	//fmt.Printf("--------\n")
	////create table
	//dba.Table(tableName).CreateTable(&Data{})
	//fmt.Printf("--------\n")
	////rw, err := dba.Table(tableName).Rows()
	////if err != nil {
	////	fmt.Printf("error(%v)\n", err)
	////	return
	////}
	//d := Data{}
	////err = rw.Scan(&d)
	////if err != nil {
	////	fmt.Printf("error(%v)\n", err)
	////	return
	////}
	////fmt.Printf("%v\n", d)
	//dba.Table(tableName).Where("device_id=?", "111").First(&d)
	//fmt.Printf("%v\n", d)
	////if rw.Next() {
	////	err = rw.Scan(&d)
	////	if err != nil {
	////		fmt.Printf("error(%v)\n", err)
	////		return
	////	}
	////	fmt.Printf("%v\n", d)
	////}
	//fmt.Printf("--------\n")
}
func main11() {
	dba, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Printf("open:%v\n", err)
		return
	}
	if err := dba.Ping(); err != nil {
		fmt.Printf("ping:%v\n", err)
		return
	}
	defer dba.Close()
	sqlString := fmt.Sprintf(createData, dbName)
	_, err = dba.Exec(sqlString)
	if err != nil {
		fmt.Printf("create data:%v\n", err)
		return
	}
	_, err = dba.Exec("use " + dbName)
	if err != nil {
		fmt.Printf("create data:%v\n", err)
		return
	}
	sqlString = fmt.Sprintf(createTable, tableName)
	_, err = dba.Exec(sqlString)
	if err != nil {
		fmt.Printf("create table:%v\n", err)
		return
	}
	sqlString = fmt.Sprintf(insertSql, dbName, tableName, "device_id,engine_type,time_stamp", "?,?,?")
	insertStmt, err := dba.Prepare(sqlString)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
		return
	}
	defer insertStmt.Close()
	//lo, err := time.LoadLocation("")
	//now := time.Now()
	nowStr := ""
	//strconv.Itoa(int(now))
	//sqlString = fmt.Sprintf("select from_unixtime(%d);", now)
	//dba.Exec(sqlString, &nowStr)
	//fmt.Printf("now:%v, lo:%v,\n", now.String(), now.Unix())
	//
	sqlString = fmt.Sprintf("select time_stamp from %s.%s where device_id=? and engine_type=?", dbName, tableName)
	row := dba.QueryRow(sqlString, "113", "e")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	err = row.Scan(&nowStr)
	fmt.Printf("row:%v, %v,error(%v)\n", row, nowStr, err)
	tn, _ := time.ParseInLocation("2006-01-02 15:04:05", nowStr, time.Local)
	fmt.Printf("time:%v, %v\n", tn.Unix(), tn.String())
	//tm, _ := strconv.Atoi(nowStr)
	//zeng := time.Unix(int64(tm), 0)
	//fmt.Printf("Zeng:%v\n", zeng.Format("2006-01-02 15:04:05"))
	//insertStmt.Exec("113", "e", now.Format("2006-01-02 15:04:05"))
}
func (d *DBParam) Create() {
	dba, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Printf("open:%v\n", err)
		return
	}
	if err := dba.Ping(); err != nil {
		fmt.Printf("ping:%v\n", err)
		return
	}
	defer dba.Close()
	//sqlString := fmt.Sprintf(createData, dbName)
	_, err = dba.Exec(d.createData)
	if err != nil {
		fmt.Printf("create data:%v\n", err)
		return
	}
	_, err = dba.Exec("use " + d.dbName)
	if err != nil {
		fmt.Printf("create data:%v\n", err)
		return
	}
	//sqlString = fmt.Sprintf(createTable, tableName)
	_, err = dba.Exec(d.createTable)
	if err != nil {
		fmt.Printf("create table:%v\n", err)
		return
	}
}
