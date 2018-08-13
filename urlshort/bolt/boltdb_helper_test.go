package bolt

import (
	"testing"
	"github.com/boltdb/bolt"
	"time"
)


func TestWriteToDb(t *testing.T){
	paths := []string {
		"/one",
		"/two",
		"/three",
		"/four",
		"/five",
		"/six",
		"/seven",
	}
	urls := []string {
		"https://godoc.org/github.com/gophercises/urlshort",
		"https://godoc.org/gopkg.in/yaml.v2",
		"https://www.google.com.ua",
		"https://gophercises.com/exercises",
		"https://github.com/gophercises/urlshort",
		"https://github.com/gophercises/urlshort/tree/solution",
		"https://github.com/boltdb/bolt",
	}
	if len(paths) != len(urls){
		t.Fatal("URLs and Paths amounts should be equal")
	}
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{Timeout:1 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func (tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("URLs_and_Paths"))
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < len(paths); i++{
			retreivedUrl := bucket.Get([]byte(paths[i]))
			if retreivedUrl == nil {
				bucket.Put([]byte(paths[i]), []byte(urls[i]))
			}
		}
		return nil
	}); if err != nil {
		t.Fatal(err)
	}
}
