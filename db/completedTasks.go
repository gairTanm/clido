package db

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/boltdb/bolt"
)

func CompletedTasks() ([]Task, error) {
	var completed []Task
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(completedBucket)
		cur := c.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			buffer := &bytes.Buffer{}
			buffer.Write(v)
			d := gob.NewDecoder(buffer)
			t := &Task{}
			err := d.Decode(t)
			if err != nil {
				return err
			}
			task := *t
			//fmt.Println(time.Now().YearDay(), task.Done.Date.YearDay(), task.Start.YearDay())
			if task.Done.Date.YearDay() == time.Now().YearDay() {
				completed = append(completed, *t)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	for i, j := 0, len(completed)-1; i < j; i, j = i+1, j-1 {
		completed[i], completed[j] = completed[j], completed[i]
	}
	return completed, nil
}
