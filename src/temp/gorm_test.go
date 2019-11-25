package temp

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

const createTableTTs = `create table if not exists device_info_tts (` +
	`device_id varchar(100) primary key,` +
	`active int unsigned default 0,` +
	`cipher varchar(256) default "",` +
	`project_id varchar(256) default ""` +
	`up_date int unsigned default 0) engine=innodb default charset=utf8`

//
type InfoTest struct {
	DeviceID  string `gorm:"column:device_id;type:varchar(100);not null;primary key"`
	Active    int    `gorm:"column:active;type:int;unsigned;not null;default:0"`
	Cipher    string `gorm:"column:cipher;type:varchar(256);not null;default:''"`
	ProjectID string `gorm:"column:project_id;type:varchar(256);not null;default:''"`
	Update    int    `gorm:"column:up_date;type:int;unsigned;not null;default:0"`
}

func (i *InfoTest) TableName() string {
	return "device_info_test"
}

func testInsert(t *testing.T, dbG *gorm.DB) {
	//Insert: insert one piece of data
	info := InfoTest{
		DeviceID: "test1",
		Active:   1,
	}
	if err := dbG.Create(&info).Error; err != nil {
		t.Errorf("insert failed\n")
		return
	}
	t.Logf("insert ok\n")
}

func testQuery(t *testing.T, dbG *gorm.DB) {
	infoQ := InfoTest{}
	if err := dbG.Take(&infoQ).Error; err != nil {
		t.Errorf("query error(%v)\n", err)
	}
	t.Logf("Query:%v\n", infoQ)
}
func TestGormOpen(t *testing.T) {
	//open db
	dbG, err := gorm.Open("mysql", "root:root@tcp(192.168.5.30:3306)/device")
	if err != nil {
		t.Errorf("error(%v)\n", err)
		return
	}

	//db ping
	if err := dbG.DB().Ping(); err != nil {
		t.Errorf("error(%v)\n", err)
		return
	}
	t.Logf("ping ok...\n")
	defer dbG.Close()
	var info InfoTest

	//judge if the table exists or not
	hasTable := dbG.HasTable(&info)
	if hasTable {
		t.Logf("has table,ok\n")
	} else {
		t.Logf("no table\n")
		//create a new table in the db
		err := dbG.Set("gorm:table_options", "engine=innodb default "+
			"charset=utf8 comment='device_info_test example'").CreateTable(&info).Error
		//err := dbG.CreateTable(&info).Error
		if err != nil {
			t.Errorf("create table error(%v)\n", err)
			return
		}
		t.Logf("create table ok(%v)\n", info.TableName())
	}
	info = InfoTest{
		DeviceID: "test1",
		Active:   1,
	}

	//Insert: insert one piece of data
	//testInsert(t, dbG)

	//Query: query a piece of data
	testQuery(t, dbG)
	//Update: update a piece of data

	//Delete: delete a piece of data

}
