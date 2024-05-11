package elasticsearch

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

var ES *elasticsearch.Client

func ConnectElasticsearch() {

	var address = os.Getenv("ES_DATA_COLLECTOR_ELASTICSEARH_ADDRESS")
	if address == "" {
		address = "http://localhost:9200"
	}

	var cfg = elasticsearch.Config{
		Addresses: []string{
			address,
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	}

	var es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	info, err := es.Info()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cluseter info : ", info)

	ES = es
}
