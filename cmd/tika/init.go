package tika

import (
	"context"
	"github.com/google/go-tika/tika"
)

var server *tika.Server

func Init() error {
	var err error

	ocrLanguages = []string{"eng", "rus"}

	server, err = tika.NewServer("tika-server-1.24.1.jar", "")

	if err != nil {
		return err
	}

	err = server.Start(context.Background())

	return err
}

func getClient() *tika.Client {
	return tika.NewClient(nil, server.URL())
}

func Close() {
	server.Stop()
}
