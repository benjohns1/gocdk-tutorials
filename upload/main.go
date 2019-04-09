package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: upload BUCKET-URL FILE")
	}
	bucketURL := os.Args[1]
	file := os.Args[2]

	ctx := context.Background()
	// Open connection to bucket
	b, err := blob.OpenBucket(ctx, bucketURL)
	if err != nil {
		log.Fatalf("Failed to setup bucket: %v", err)
	}
	defer b.Close()

	// Prepare file for upload
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Write to bucket
	w, err := b.NewWriter(ctx, file, nil)
	if err != nil {
		log.Fatalf("Failed to obtain writer: %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to bucket: %v", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("Failed to close: %s", err)
	}
}
