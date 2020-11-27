package elastics

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic"
)

// var (
// 	Client ESClientInterface = &ESClient{}
// )

type ESClientInterface interface {
	SetClient(*elastic.Client) ESClientInterface
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type ESClient struct {
	Client *elastic.Client
}

func Init() ESClientInterface {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		// elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetRetrier(NewCustomRetrier()),
		// elastic.SetGzip(true),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		// elastic.SetHeaders(http.Header{
		// 	"X-Caller-Id": []string{"..."},
		// }),
	)
	if err != nil {
		panic(err)
	}

	var es ESClient
	es.SetClient(client)

	return &es
}

func (c *ESClient) SetClient(client *elastic.Client) ESClientInterface {
	c.Client = client
	return c
}

func (c *ESClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.Client.Index().Index(index).Type(docType).BodyJson(doc).Do(ctx)
	if err != nil {
		fmt.Errorf("error when creating item with index %s. %s", index, err.Error())
		return nil, err
	}
	return result, nil

}

func (c *ESClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.Client.Get().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {
		fmt.Errorf("error when trying to get id %s. %s", id, err.Error())
		return nil, err
	}
	return result, nil
}

func (c *ESClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	fmt.Println("ES -----------")

	ctx := context.Background()

	result, err := c.Client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	// result, err := c.Client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		fmt.Errorf("error when searching in document with query %s. %s", query, err.Error())
		return nil, err
	}

	return result, nil
}
