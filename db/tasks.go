package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var completedBucket = []byte("completed")
var db *bolt.DB

type Completed struct {
	Complete bool
	Date     time.Time
}

type Task struct {
	Key   int
	Value string
	Start time.Time
	Done  Completed
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		_, err = tx.CreateBucketIfNotExists(completedBucket)
		return err
	})
}
