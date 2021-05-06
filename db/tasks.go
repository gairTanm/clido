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
	Key      int
	Value    string
	Priority float64
	Start    time.Time
	Done     Completed
}

type ByPriority []Task

func (t ByPriority) Len() int { return len(t) }

func (t ByPriority) Less(i, j int) bool { return t[i].Priority > t[j].Priority }

func (t ByPriority) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

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
