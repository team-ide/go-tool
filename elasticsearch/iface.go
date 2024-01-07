package elasticsearch

import "github.com/olivere/elastic/v7"

type IService interface {
	// Close  关闭 elasticsearch 客户端
	Close()
	// Info  获取 elasticsearch 信息
	Info() (res *elastic.NodesInfoResponse, err error)
	// DeleteIndex  删除 索引
	DeleteIndex(indexName string) (err error)
	// CreateIndex  创建 索引
	CreateIndex(indexName string, bodyJSON map[string]interface{}) (err error)
	// Indexes  查询 索引
	Indexes() (indexes []*IndexInfo, err error)
	// GetMapping  查询 索引 配置
	GetMapping(indexName string) (res interface{}, err error)
	// PutMapping  设置 索引 配置
	PutMapping(indexName string, bodyJSON map[string]interface{}) (err error)
	// SetFieldType  设置 索引 字段类型
	SetFieldType(indexName string, fieldName string, fieldType string) (err error)
	// Search  搜索
	Search(indexName string, pageIndex int, pageSize int, whereList []*Where, orderList []*Order) (res *SearchResult, err error)
	// Insert  插入数据 并且 等待刷新
	Insert(indexName string, id string, doc interface{}) (res *InsertResponse, err error)
	// InsertNotWait  插入数据 不 等待刷新
	InsertNotWait(indexName string, id string, doc interface{}) (res *InsertResponse, err error)
	// BatchInsertNotWait  批量插入数据 不 等待刷新
	BatchInsertNotWait(docs []*InsertDoc) (res *BulkResponse, err error)
	// Update  更新数据 并且 等待刷新
	Update(indexName string, id string, doc interface{}) (res *UpdateResponse, err error)
	// UpdateNotWait  更新数据 不 等待刷新
	UpdateNotWait(indexName string, id string, doc interface{}) (res *UpdateResponse, err error)
	// Delete  删除 并且 等待刷新
	Delete(indexName string, id string) (res *DeleteResponse, err error)
	// DeleteNotWait  删除 不 等待刷新
	DeleteNotWait(indexName string, id string) (res *DeleteResponse, err error)
	// Reindex 修改索引名称
	Reindex(sourceIndexName string, toIndexName string) (res *BulkIndexByScrollResponse, err error)
	// IndexStat 索引状态
	IndexStat(indexName string) (res *IndicesStatsResponse, err error)
	// Scroll 滚动查询
	Scroll(indexName string, scrollId string, pageSize int, whereList []*Where, orderList []*Order) (res *SearchResult, err error)
	// IndexAlias 索引别名
	IndexAlias(indexName string, aliasName string) (res *IndexAliasResponse, err error)
	PerformRequest(options PerformRequestOptions) (res *PerformResponse, err error)
	QuerySql(query string) (res *QuerySqlResult, err error)
}
