package main

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
	"log"
)

func main() {
	//bits
	opt := nutsdb.DefaultOptions
	opt.Dir = "./tmp/nutsdb" //这边数据库会自动创建这个目录文件
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//for i := 0; i < 5; i++ {
	//	err := db.Update(func(tx *nutsdb.Tx) error {
	//		bucket := "test"
	//		key := []byte("myList")
	//		val := []byte(fmt.Sprintf("myList: %d\n", i))
	//		return tx.RPush(bucket, key, val)
	//	})
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}

	//if err := db.View(
	//	func(tx *nutsdb.Tx) error {
	//		bucket := "test"
	//		key := []byte("myList")
	//		if items, err := tx.LRange(bucket, key, 0, -1); err != nil {
	//			fmt.Println("===========")
	//			return err
	//		} else {
	//			//fmt.Println(items)
	//			for _, item := range items {
	//				//fmt.Println("###")
	//				fmt.Print(string(item))
	//			}
	//		}
	//		return nil
	//	}); err != nil {
	//	log.Fatal(err)
	//}

	key := []byte("key001")
	val := []byte("val001\n")

	bucket001 := "bucket001"
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(bucket001, key, val, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			//bucket := bucket001
			//key := []byte("myList")
			if e, err := tx.Get(bucket001, key); err != nil {
				return err
			} else {
				fmt.Println(string(e.Value)) // "val1-modify"
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}

	////
	//bucket002 := "bucket002"
	//if err := db.Update(
	//	func(tx *nutsdb.Tx) error {
	//		if err := tx.Put(bucket002, key, val, 0); err != nil {
	//			return err
	//		}
	//		return nil
	//	}); err != nil {
	//	log.Fatal(err)
	//}
	err := db.Update(func(tx *nutsdb.Tx) error {

	})
}

type name struct {
	AgeNum string `json:"age_num"`
}
