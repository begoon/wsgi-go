package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"cloud.google.com/go/storage"
)

func StorageSave(bucket string, name string, data []byte) {
	c := context.Background()
	w := Must(storage.NewClient(c)).Bucket(bucket).Object(name).NewWriter(c)
	defer Success(w.Close())
	w.Write(data)
}

func TestStorageWrite(t *testing.T) {
	bucket := "chiro-alexander-testing"

	StorageSave(bucket, "just a json", []byte("just a text"))
	StorageSave(bucket, "just a json", Must(os.ReadFile("nel.jpeg")))

	// json
	payload := struct {
		Name  string
		Age   int
		Stuff []string
	}{"alice", 30, []string{"one", "two", "three"}}

	StorageSave(bucket, "just a json", Must(json.Marshal(&payload)))
}
