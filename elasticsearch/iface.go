package elasticsearch

import "github.com/olivere/elastic/v7"

type IService interface {
	Info() (res *elastic.NodesInfoResponse, err error)
	DeleteIndex(indexName string) (err error)
	CreateIndex(indexName string, bodyJSON map[string]interface{}) (err error)
	Indexes() (indexes []*IndexInfo, err error)
	GetMapping(indexName string) (res interface{}, err error)
	PutMapping(indexName string, bodyJSON map[string]interface{}) (err error)
	SetFieldType(indexName string, fieldName string, fieldType string) (err error)
	Search(indexName string, pageIndex int, pageSize int, whereList []*Where, orderList []*Order) (res *SearchResult, err error)
	Insert(indexName string, id string, doc interface{}) (res *InsertResponse, err error)
	InsertNotWait(indexName string, id string, doc interface{}) (res *InsertResponse, err error)
	BatchInsertNotWait(docs []*InsertDoc) (res *BulkResponse, err error)
	Update(indexName string, id string, doc interface{}) (res *UpdateResponse, err error)
	Delete(indexName string, id string) (res *DeleteResponse, err error)
	Reindex(sourceIndexName string, toIndexName string) (res *BulkIndexByScrollResponse, err error)
	IndexStat(indexName string) (res *IndicesStatsResponse, err error)
	Scroll(indexName string, scrollId string, pageSize int, whereList []*Where, orderList []*Order) (res *SearchResult, err error)
	IndexAlias(indexName string, aliasName string) (res *IndexAliasResponse, err error)
}
