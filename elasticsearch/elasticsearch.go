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
			ResponseHeaderTimeout: 180 * time.Second, //flush: net/http: timeout awaiting response headers hatası için arttırıldı
			DialContext:           (&net.Dialer{Timeout: 90 * time.Second}).DialContext,
		},
		RetryOnStatus: []int{502, 503, 504, 429},
		RetryBackoff:  func(i int) time.Duration { return time.Duration(i) * 2000 * time.Millisecond },
		MaxRetries:    5,
	}

	var es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println(err)
	}

	info, err := es.Info()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Cluseter info : ", info)

	ES = es
}
