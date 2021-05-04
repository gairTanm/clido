package db

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
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
			if task.Done.Date.YearDay() == task.Start.YearDay() {
				completed = append(completed, *t)
			} else {
				err = c.Delete(k)
				if err != nil {
					return err
				}
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

func RemoveTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
