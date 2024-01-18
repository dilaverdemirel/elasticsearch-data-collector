package elasticsearch

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

var ES *elasticsearch.Client

func ConnectElasticsearch() {
	var cfg = elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9201",
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	}

	ES, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	info, err := ES.Info()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cluseter info : ", info)

}
