package db

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/boltdb/bolt"
)

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)

		key := itob(int(id64))
		t := Task{
			Key:   id,
			Value: task,
			Start: time.Now(),
			Done: Completed{
				Complete: false,
				Date:     time.Now(),
			},
		}
		buffer := &bytes.Buffer{}
		err := gob.NewEncoder(buffer).Encode(t)
		if err != nil {
			return err
		}
		return b.Put(key, buffer.Bytes())
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}
