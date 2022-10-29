package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticDocs struct {
	Name      string
	Id        int32
	IsStudent bool
}

func main() {
	ctx := context.Background()

	client, err := elasticsearch.NewDefaultClient()

	checkError(err)

	res, err := client.Info()

	// Deserialize the response into a map.
	if err != nil {
		log.Fatalln("client.Info() ERROR:", err)
	} else {
		// log.Println("client response:", res)
	}

	doc1 := ElasticDocs{}
	doc1.Name = "Andrey Shkunov"
	doc1.Id = 123456
	doc1.IsStudent = true

	docStr1 := jsonStruct(doc1)

	req := esapi.IndexRequest{
		Index:      "person",
		DocumentID: strconv.Itoa(1),
		Body:       strings.NewReader(docStr1),
		Refresh:    "true",
	}

	res, err = req.Do(ctx, client)
	checkError(err)
	defer res.Body.Close()

	var resMap map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resMap)
	checkError(err)

	res, err = client.Search(client.Search.WithIndex("person"))
	checkError(err)
	defer res.Body.Close()

	fmt.Println("\n\nGet\n\n", res.String())

	_, err = client.Update("person", "1", strings.NewReader(jsonStruct(ElasticDocs{
		IsStudent: false,
		Id:        123456,
		Name:      "Andrey Shkunov",
	})), client.Update.WithRefresh("true"))
	checkError(err)
	defer res.Body.Close()

	res, err = client.Search(client.Search.WithIndex("person"))
	checkError(err)
	defer res.Body.Close()

	fmt.Println("\n\nUpdate\n\n", res.String())

	_, err = client.Delete("person", "1", client.Delete.WithRefresh("true"))
	checkError(err)
	defer res.Body.Close()

	res, err = client.Search(client.Search.WithIndex("person"))
	checkError(err)
	defer res.Body.Close()

	fmt.Println("\n\nDelete\n\n", res.String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func jsonStruct(doc ElasticDocs) string {

	// Create struct instance of the Elasticsearch fields struct object
	docStruct := &ElasticDocs{
		Name:      doc.Name,
		Id:        doc.Id,
		IsStudent: doc.IsStudent,
	}

	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(docStruct)
	checkError(err)
	return string(b)
}
