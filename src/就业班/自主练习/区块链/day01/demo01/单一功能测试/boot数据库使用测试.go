package 单一功能测试
//
//import (
//	"github.com/bolt-master"
//	"log"
//	"fmt"
//)
//
//func main() {
//	db, err := bolt.Open("test.db", 0600, nil)
//	if err != nil {
//		log.Panic("dbOpen", err)
//	}
//	db.Update(func(tx *bolt.Tx) error {
//		bucket := tx.Bucket([]byte("dbtest"))
//		if bucket == nil {
//			bucket, err = tx.CreateBucket([]byte("dbtest"))
//		}
//		err := bucket.Put([]byte("publicKey"), []byte("dfsdfgsdfgsdfgsdfgsd"))
//		return err
//	})
//	//读数据
//	db.View(func(tx *bolt.Tx) error {
//		bucket := tx.Bucket([]byte("dbtest"))
//		if bucket == nil {
//			fmt.Println("数据库不存在")
//			return err
//		}
//		res := bucket.Get([]byte("publicKey"))
//		fmt.Println("读数据", res)
//		return err
//	})
//	defer db.Close()
//}
