package data

import (
	"context"
	"iosync/ent"

	_ "github.com/lib/pq"
)

func Init(c context.Context) (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=iosync password=password sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = client.Schema.Create(c)
	return client, err
}
