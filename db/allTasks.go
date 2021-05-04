package db

import (
	"bytes"
	"encoding/gob"

	"github.com/boltdb/bolt"
)

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			buffer := &bytes.Buffer{}
			buffer.Write(v)
			dec := gob.NewDecoder(buffer)

			t := &Task{}
			err := dec.Decode(t)
			if err != nil {
				return err
			}
			tasks = append(tasks, *t)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
