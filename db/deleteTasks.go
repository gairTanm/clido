package db

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/boltdb/bolt"
)

func DeleteTasks(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := tx.Bucket(completedBucket)
		task := b.Get(itob(key))
		buffer := &bytes.Buffer{}
		t := &Task{}
		buffer.Write(task)
		d := gob.NewDecoder(buffer)
		err := d.Decode(t)
		if err != nil {
			return err
		}

		decTask := *t
		completedTask := Task{
			Key:   decTask.Key,
			Value: decTask.Value,
			Start: decTask.Start,
			Done: Completed{
				Complete: true,
				Date:     time.Now(),
			},
		}
		//fmt.Println(completedTask.Done.Date.YearDay())
		enBuffer := &bytes.Buffer{}
		e := gob.NewEncoder(enBuffer)
		err = e.Encode(completedTask)
		if err != nil {
			return err
		}
		err = c.Put(itob(key), enBuffer.Bytes())
		if err != nil {
			return err
		}
		return b.Delete(itob(key))
	})
}
