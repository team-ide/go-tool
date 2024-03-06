package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IService interface {
	Close()
	Databases() (databases []*Database, totalSize int64, err error)
	Collections(database string) (collections []*Collection, err error)
	CollectionCreate(database string, collection string) (err error)
	CollectionDelete(database string, collection string) (err error)
	Indexes(database string, collection string) (indexes []map[string]interface{}, err error)
	IndexCreate(database string, collection string, index mongo.IndexModel) (name string, err error)
	IndexesCreate(database string, collection string, indexes []mongo.IndexModel) (names []string, err error)
	IndexDelete(database string, collection string, name string) (err error)
	IndexDeleteAll(database string, collection string) (err error)
	Insert(database string, collection string, document interface{}) (insertedId interface{}, err error)
	BatchInsert(database string, collection string, documents []interface{}) (insertedIds []interface{}, err error)
	Update(database string, collection string, id interface{}, update interface{}) (updateResult *UpdateResult, err error)
	UpdateOne(database string, collection string, filter interface{}, update interface{}) (updateResult *UpdateResult, err error)
	BatchUpdate(database string, collection string, filter interface{}, update interface{}) (updateResult *UpdateResult, err error)
	Count(database string, collection string, filter interface{}) (totalCount int64, err error)
	QueryMap(database string, collection string, filter interface{}, opts *options.FindOptions) (list []map[string]interface{}, err error)
	QueryMapPage(database string, collection string, filter interface{}, page *Page, opts *options.FindOptions) (list []map[string]interface{}, err error)
	QueryMapPageResult(database string, collection string, filter interface{}, page *Page, opts *options.FindOptions) (pageResult *Page, err error)
}
