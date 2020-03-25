package main

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
	"log"
)

func main() {
	bits
	opt := nutsdb.DefaultOptions
	opt.Dir = "./tmp/nutsdb" //这边数据库会自动创建这个目录文件
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < 10; i++ {
		err := db.Update(func(tx *nutsdb.Tx) error {
			bucket := "test"
			key := []byte("myList")
			val := []byte(fmt.Sprintf("myList: %d\n", i))
			return tx.RPush(bucket, key, val)
		})
		if err != nil {
			log.Println(err)
		}
	}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			bucket := "test"
			key := []byte("myList")
			if items, err := tx.LRange(bucket, key, 0, -1); err != nil {
				return err
			} else {
				//fmt.Println(items)
				for _, item := range items {
					fmt.Print(string(item))
				}
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
}
