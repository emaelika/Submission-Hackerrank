package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
)

func main() {
	fmt.Println("begin")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	store, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	a, err := store.Bucket("aaa")
	if err != nil {
		log.Fatalln(err)
	}
	a.ACL()
}
