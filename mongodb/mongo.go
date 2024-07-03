package mongodb

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"os"
	"strings"
	"time"
)

type Service struct {
	*Config
	client   *mongo.Client
	isClosed bool
}

func (this_ *Service) init() (err error) {

	var minPoolSize = 10
	if this_.MinPoolSize > 0 {
		minPoolSize = this_.MinPoolSize
	}
	var maxPoolSize = 20
	if this_.MaxPoolSize >= minPoolSize {
		maxPoolSize = this_.MaxPoolSize
	}
	var connectTimeout = 10
	if this_.ConnectTimeout > 0 {
		connectTimeout = this_.ConnectTimeout
	}

	clientOptions := options.Client().SetHosts(this_.GetServers()).
		SetMinPoolSize(uint64(minPoolSize)).
		SetMaxPoolSize(uint64(maxPoolSize)).
		SetConnectTimeout(time.Second * time.Duration(connectTimeout))

	//设置用户名和密码
	username := this_.Username
	password := this_.Password

	if len(username) > 0 && len(password) > 0 {
		clientOptions.SetAuth(options.Credential{Username: username, Password: password})
	}
	if this_.CertPath != "" {
		TLSClientConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		certPool := x509.NewCertPool()
		var pemCerts []byte
		pemCerts, err = os.ReadFile(this_.CertPath)
		if err != nil {
			return
		}

		if !certPool.AppendCertsFromPEM(pemCerts) {
			err = errors.New("证书[" + this_.CertPath + "]解析失败")
			return
		}
		TLSClientConfig.RootCAs = certPool

		//TLSClientConfig.Certificates = []tls.Certificate{clicrt}
		clientOptions.TLSConfig = TLSClientConfig
	}

	this_.client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return
	}

	// Check the connection
	err = this_.client.Ping(context.TODO(), nil)
	if err != nil {
		return
	}
	this_.isClosed = false
	return
}

func (this_ *Service) GetServers() []string {
	var servers []string
	if this_.Address == "" {
		return servers
	}
	if strings.Contains(this_.Address, ",") {
		servers = strings.Split(this_.Address, ",")
	} else if strings.Contains(this_.Address, ";") {
		servers = strings.Split(this_.Address, ";")
	} else {
		servers = []string{this_.Address}
	}
	return servers
}

func (this_ *Service) Close() {
	if this_ == nil {
		return
	}
	if this_.isClosed {
		return
	}
	this_.isClosed = true
	if this_.client != nil {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = this_.client.Disconnect(ctx)
		this_.client = nil
	}
}

type Database struct {
	Name       string `json:"name"`
	SizeOnDisk int64  `json:"sizeOnDisk"`
	Empty      bool   `json:"empty"`
}

func (this_ *Service) Databases() (databases []*Database, totalSize int64, err error) {
	res, err := this_.client.ListDatabases(context.TODO(), bson.D{})
	if err != nil {
		return
	}
	for _, one := range res.Databases {
		d := &Database{}
		d.Name = one.Name
		d.SizeOnDisk = one.SizeOnDisk
		d.Empty = one.Empty
		databases = append(databases, d)
	}
	totalSize = res.TotalSize
	return
}

func (this_ *Service) DatabaseDelete(database string) (err error) {
	err = this_.client.Database(database).Drop(context.TODO())
	if err != nil {
		return
	}
	return
}

type Collection struct {
	Name    string                 `json:"name"`
	IdIndex map[string]interface{} `json:"idIndex"`
	Info    map[string]interface{} `json:"info"`
	Type    string                 `json:"type"`
	Options map[string]interface{} `json:"options"`
}

func (this_ *Service) Collections(database string) (collections []*Collection, err error) {
	rows, err := this_.client.Database(database).ListCollections(context.TODO(), bson.D{})
	if err != nil {
		return
	}
	defer func() { _ = rows.Close(context.Background()) }()

	ctx := context.TODO()
	for rows.Next(ctx) {
		one := &Collection{}
		err = rows.Decode(one)
		if err != nil {
			return
		}
		collections = append(collections, one)
	}
	return
}

func (this_ *Service) CollectionDelete(database string, collection string) (err error) {
	err = this_.client.Database(database).Collection(collection).Drop(context.TODO())
	if err != nil {
		return
	}
	return
}

func (this_ *Service) CollectionCreate(database string, collection string, opts ...*options.CreateCollectionOptions) (err error) {
	err = this_.client.Database(database).CreateCollection(context.TODO(), collection, opts...)
	if err != nil {
		return
	}
	return
}

func (this_ *Service) Indexes(database string, collection string) (indexes []map[string]interface{}, err error) {
	rows, err := this_.client.Database(database).Collection(collection).Indexes().List(context.TODO())
	if err != nil {
		return
	}
	defer func() { _ = rows.Close(context.Background()) }()

	ctx := context.TODO()
	for rows.Next(ctx) {
		one := map[string]interface{}{}
		err = rows.Decode(one)
		if err != nil {
			return
		}
		indexes = append(indexes, one)
	}
	return
}

func (this_ *Service) IndexCreate(database string, collection string, index mongo.IndexModel) (name string, err error) {
	name, err = this_.client.Database(database).Collection(collection).Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		return
	}
	return
}

func (this_ *Service) IndexesCreate(database string, collection string, indexes []mongo.IndexModel) (names []string, err error) {
	names, err = this_.client.Database(database).Collection(collection).Indexes().CreateMany(context.TODO(), indexes)
	if err != nil {
		return
	}
	return
}

func (this_ *Service) IndexDelete(database string, collection string, name string) (err error) {
	_, err = this_.client.Database(database).Collection(collection).Indexes().DropOne(context.TODO(), name)
	if err != nil {
		return
	}
	return
}

func (this_ *Service) IndexDeleteAll(database string, collection string) (err error) {
	_, err = this_.client.Database(database).Collection(collection).Indexes().DropAll(context.TODO())
	if err != nil {
		return
	}
	return
}

func (this_ *Service) Insert(database string, collection string, document interface{}, opts ...*options.InsertOneOptions) (insertedId interface{}, err error) {
	res, err := this_.client.Database(database).Collection(collection).InsertOne(context.TODO(), document, opts...)
	if err != nil {
		return
	}
	if res != nil {
		insertedId = res.InsertedID
	}
	return
}

func (this_ *Service) BatchInsert(database string, collection string, documents []interface{}, opts ...*options.InsertManyOptions) (insertedIds []interface{}, err error) {
	res, err := this_.client.Database(database).Collection(collection).InsertMany(context.TODO(), documents, opts...)
	if err != nil {
		return
	}
	if res != nil {
		insertedIds = res.InsertedIDs
	}
	return
}

type UpdateResult struct {
	MatchedCount  int64       `json:"matchedCount"`
	ModifiedCount int64       `json:"modifiedCount"`
	UpsertedCount int64       `json:"upsertedCount"`
	UpsertedID    interface{} `json:"upsertedID"`
}

func (this_ *Service) Update(database string, collection string, id interface{}, update interface{}, opts ...*options.UpdateOptions) (updateResult *UpdateResult, err error) {
	res, err := this_.client.Database(database).Collection(collection).UpdateByID(context.TODO(), id, update, opts...)
	if err != nil {
		return
	}
	updateResult = &UpdateResult{}
	if res != nil {
		updateResult.MatchedCount = res.MatchedCount
		updateResult.ModifiedCount = res.ModifiedCount
		updateResult.UpsertedCount = res.UpsertedCount
		updateResult.UpsertedID = res.UpsertedID
	}
	return
}

func (this_ *Service) UpdateOne(database string, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (updateResult *UpdateResult, err error) {
	res, err := this_.client.Database(database).Collection(collection).UpdateOne(context.TODO(), filter, update, opts...)
	if err != nil {
		return
	}
	updateResult = &UpdateResult{}
	if res != nil {
		updateResult.MatchedCount = res.MatchedCount
		updateResult.ModifiedCount = res.ModifiedCount
		updateResult.UpsertedCount = res.UpsertedCount
		updateResult.UpsertedID = res.UpsertedID
	}
	return
}

func (this_ *Service) BatchUpdate(database string, collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (updateResult *UpdateResult, err error) {
	res, err := this_.client.Database(database).Collection(collection).UpdateMany(context.TODO(), filter, update, opts...)
	if err != nil {
		return
	}
	updateResult = &UpdateResult{}
	if res != nil {
		updateResult.MatchedCount = res.MatchedCount
		updateResult.ModifiedCount = res.ModifiedCount
		updateResult.UpsertedCount = res.UpsertedCount
		updateResult.UpsertedID = res.UpsertedID
	}
	return
}

func (this_ *Service) QueryMap(database string, collection string, filter interface{}, opts *options.FindOptions) (list []map[string]interface{}, err error) {
	rows, err := this_.client.Database(database).Collection(collection).Find(context.TODO(), filter, opts)
	if err != nil {
		return
	}
	defer func() { _ = rows.Close(context.Background()) }()

	ctx := context.Background()
	for rows.Next(ctx) {
		one := map[string]interface{}{}
		err = rows.Decode(one)
		if err != nil {
			return
		}
		list = append(list, one)
	}
	return
}

type Page struct {
	PageSize   int64         `json:"pageSize"`
	PageNo     int64         `json:"pageNo"`
	TotalCount int64         `json:"totalCount"`
	TotalPage  int64         `json:"totalPage"`
	List       []interface{} `json:"list"`
}

func (this_ *Service) QueryMapPage(database string, collection string, filter interface{}, page *Page, opts *options.FindOptions) (list []map[string]interface{}, err error) {
	if opts == nil {
		opts = options.Find()
	}
	if opts.Limit == nil {
		opts.SetLimit(page.PageSize)
	}
	if opts.Skip == nil {
		opts.SetSkip(page.PageSize * (page.PageNo - 1))
	}
	rows, err := this_.client.Database(database).Collection(collection).Find(context.TODO(), filter, opts)
	if err != nil {
		return
	}
	defer func() { _ = rows.Close(context.Background()) }()

	ctx := context.Background()
	for rows.Next(ctx) {
		one := bson.M{}
		err = rows.Decode(one)
		if err != nil {
			return
		}
		list = append(list, one)
	}
	return
}

func (this_ *Service) QueryMapPageResult(database string, collection string, filter interface{}, page *Page, opts *options.FindOptions) (pageResult *Page, err error) {
	totalCount, err := this_.client.Database(database).Collection(collection).CountDocuments(context.TODO(), filter)
	if err != nil {
		return
	}
	pageResult = &Page{}
	pageResult.TotalCount = totalCount
	pageResult.TotalPage = int64(math.Ceil(float64(totalCount) / float64(page.PageSize))) // page 总数
	pageResult.PageSize = page.PageSize
	pageResult.PageNo = page.PageNo

	list, err := this_.QueryMapPage(database, collection, filter, page, opts)
	if err != nil {
		return
	}
	for _, one := range list {
		pageResult.List = append(pageResult.List, one)
	}
	return
}

func (this_ *Service) Count(database string, collection string, filter interface{}) (totalCount int64, err error) {
	totalCount, err = this_.client.Database(database).Collection(collection).CountDocuments(context.TODO(), filter)
	if err != nil {
		return
	}
	return
}

func (this_ *Service) DeleteOne(database string, collection string, filter interface{}) (deletedCount int64, err error) {
	res, err := this_.client.Database(database).Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return
	}
	if res != nil {
		deletedCount = res.DeletedCount
	}
	return
}

func (this_ *Service) DeleteMany(database string, collection string, filter interface{}) (deletedCount int64, err error) {
	res, err := this_.client.Database(database).Collection(collection).DeleteMany(context.TODO(), filter)
	if err != nil {
		return
	}
	if res != nil {
		deletedCount = res.DeletedCount
	}
	return
}
