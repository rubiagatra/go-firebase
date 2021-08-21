package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func CloudStorage() *storage.BucketHandle {
	config := &firebase.Config{
		StorageBucket: "",
	}
	opt := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	// 'bucket' is an object defined in the cloud.google.com/go/storage package.
	// See https://godoc.org/cloud.google.com/go/storage#BucketHandle
	// for more details.
	// [END cloud_storage_golang]

	return bucket
}
