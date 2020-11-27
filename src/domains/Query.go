package domains

import "github.com/olivere/elastic"

type EsQuery struct {
	Equal []FieldValue `json:"equal"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	equalsQuery := make([]elastic.Query, 0)
	for _, eq := range q.Equal {
		equalsQuery = append(equalsQuery, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	query.Must(equalsQuery...)
	return query
}
