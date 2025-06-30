package es

import (
	"github.com/elastic/go-elasticsearch/v9"
	"sync"
)

var (
	esClient *elasticsearch.Client
	once     sync.Once
)

func GetESClient(url string) *elasticsearch.Client {
	once.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: []string{url},
		}
		client, err := elasticsearch.NewClient(cfg)
		if err != nil {
			panic(err)
		}
		esClient = client
	})
	return esClient
}
