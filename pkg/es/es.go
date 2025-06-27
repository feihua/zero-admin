package es

import (
	"github.com/elastic/go-elasticsearch/v9"
	"sync"
)

var (
	esClient *elasticsearch.Client
	once     sync.Once
)

func GetESClient() *elasticsearch.Client {
	once.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: []string{"http://localhost:9200"},
		}
		client, err := elasticsearch.NewClient(cfg)
		if err != nil {
			panic(err)
		}
		esClient = client
	})
	return esClient
}
