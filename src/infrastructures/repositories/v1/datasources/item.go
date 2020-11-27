package datasources

import (
	"encoding/json"
	"errors"

	"tiket.vip/src/domains"
	elastics "tiket.vip/src/infrastructures/elastic"
)

const (
	itemsIndex = "items"
	docType    = "_doc"
)

type ItemRepos struct {
	Es elastics.ESClientInterface
}

func NewItemRepo(e elastics.ESClientInterface) domains.ItemRepo {
	return &ItemRepos{
		Es: e,
	}
}

func (i *ItemRepos) Search(query domains.EsQuery) ([]domains.Item, error) {
	result, err := i.Es.Search(itemsIndex, query.Build())
	if err != nil {
		return nil, errors.New("database error")
	}

	items := make([]domains.Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item domains.Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, errors.New("database error")
		}
		items[index] = item
	}
	if len(items) == 0 {
		return nil, errors.New("no item found")
	}

	return items, nil
}
