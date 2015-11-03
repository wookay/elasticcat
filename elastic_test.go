package elasticcat

import (
	"encoding/json"
	"gopkg.in/olivere/elastic.v2"
	"log"
	"testing"
)

func TestElasticV2(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.PerformRequest("GET", "/", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res == nil {
		t.Fatal("expected response to be != nil")
	}

	ret := new(elastic.PingResult)
	if err := json.Unmarshal(res.Body, ret); err != nil {
		t.Fatalf("expected no error on decode; got: %v", err)
	}
	log.Println(ret)
	if ret.Status != 200 {
		t.Errorf("expected HTTP status 200; got: %d", ret.Status)
	}
}
