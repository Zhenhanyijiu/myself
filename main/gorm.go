package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Data1 struct {
	ID        int64  `gorm:"column:id;not null;primary_key;AUTO_INCREMENT;comment:'auto incrementt'"`
	Text1     string `gorm:"column:text_1;type:varchar(128);not null;default:'';unique_index:unq_t;comment:'text 1'"`
	IntValue1 int32  `gorm:"column:int_value_1;type:int unsigned;not null;default:0;unique_index:unq_t;comment:'int value 1'"`
	Context1  string `gorm:"column:context_1;type:varchar(1024);not null;default:'';comment:'context 1'"`
	TimeStamp string `gorm:"column:time_stamp;type:datetime;not null;default:current_timestamp on update current_timestamp;comment:'更新时间'"`
}

func (d Data1) TableName() string {
	return "table_data_1"
}

type Data2 struct {
	ID        int64  `gorm:"column:id;not null;primary_key;AUTO_INCREMENT;comment:'auto incrementt'"`
	Text2     string `gorm:"column:text_2;type:varchar(128);not null;default:'';unique_index:unq_t;comment:'text 2'"`
	IntValue2 int32  `gorm:"column:int_value_2;type:int unsigned;not null;default:0;unique_index:unq_t;comment:'int value 2'"`
	Context2  string `gorm:"column:context_2;type:varchar(1024);not null;default:'';comment:'context 2'"`
	TimeStamp string `gorm:"column:time_stamp;type:datetime;not null;default:current_timestamp on update current_timestamp;comment:'更新时间'"`
}

func (d Data2) TableName() string {
	return "table_data_2"
}

type TestGorm struct {
	mysqlCli *gorm.DB
}

func NewTestGorm() (*TestGorm, error) {
	mysqlCli, err := gorm.Open("mysql", "root:123456@tcp(192.168.5.25:3306)/test_gorm?charset=utf8&interpolateParams=true")
	if err != nil {
		return nil, err
	}
	if err := mysqlCli.DB().Ping(); err != nil {
		return nil, err
	}
	mysqlCli.DB().SetMaxIdleConns(100)
	mysqlCli.DB().SetMaxOpenConns(100)
	mysqlCli.DB().SetConnMaxLifetime(time.Second * 100)
	mysqlCli = mysqlCli.Debug()
	return &TestGorm{mysqlCli: mysqlCli}, nil

}
func (t *TestGorm) Close() {
	t.mysqlCli.Close()
}
func (t *TestGorm) CreateTable() error {
	hasTable := t.mysqlCli.HasTable(&Data1{})
	d1 := Data1{}
	d2 := Data2{}
	fmt.Printf("%s,hasTable:%v\n", d1.TableName(), hasTable)
	if !hasTable {
		err := t.mysqlCli.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 COMMENT='test_1 table'").CreateTable(&Data1{}).Error
		if err != nil {
			return err
		}
	}
	hasTable = t.mysqlCli.HasTable(&Data2{})
	fmt.Printf("%s,hasTable:%v\n", d2.TableName(), hasTable)
	if !hasTable {
		err := t.mysqlCli.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 COMMENT='test_2 table'").CreateTable(&Data2{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func (t *TestGorm) Insert(v interface{}) error {
	return t.mysqlCli.Create(v).Error
}
func (t *TestGorm) Add(text string) error {
	if text == "" {
		return fmt.Errorf("text==null")
	}
	res := t.mysqlCli.Model(&Data1{}).Where(&Data1{Text1: text, IntValue1: 1}).Update(&Data1{Context1: "golang ...."})
	if res.Error != nil {
		return res.Error
	}
	fmt.Printf("res.RowsAffected:%v\n", res.RowsAffected)
	if res.RowsAffected > 0 {

	} else {
		t.mysqlCli.Create(&Data1{Text1: text, IntValue1: 1, Context1: "ccc"})
	}
	return nil

}
func main() {
	test, err := NewTestGorm()
	if err != nil {
		fmt.Printf("===open mysql error(%v)\n", err)
		return
	}
	defer test.Close()

	if err := test.CreateTable(); err != nil {
		fmt.Printf("===open mysql error(%v)\n", err)
		return
	}

	/*
		err = test.Insert(&Data1{Text1: "aa", IntValue1: 1, Context1: "hello"})
		if err != nil {
			fmt.Printf("insert 1 error(%v)\n", err)
		}
		err = test.Insert(&Data2{Text2: "aa", IntValue2: 1, Context2: "hello"})
		if err != nil {
			fmt.Printf("insert 2 error(%v)\n", err)
		}
	*/
	test.Add("aaq")
}
