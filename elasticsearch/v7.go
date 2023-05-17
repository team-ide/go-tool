package elasticsearch

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"sync"
)

// V7Service 注册处理器在线信息等
type V7Service struct {
	*Config
	client     *elastic.Client
	clientLock sync.Mutex
}

func (this_ *V7Service) init() error {
	var err error
	return err
}

func (this_ *V7Service) GetClient() (client *elastic.Client, err error) {
	if this_ == nil {
		err = errors.New("elasticsearch service is null")
		return
	}
	this_.clientLock.Lock()
	defer this_.clientLock.Unlock()
	if this_.client != nil && this_.client.IsRunning() {
		client = this_.client
		return
	}
	util.Logger.Info("es client is null or not running,now to create client", zap.Any("url", this_.Url))
	var urls []string
	if strings.Contains(this_.Url, ",") {
		urls = strings.Split(this_.Url, ",")
	} else if strings.Contains(this_.Url, ";") {
		urls = strings.Split(this_.Url, ";")
	} else {
		urls = []string{this_.Url}
	}
	var isHttps bool
	for _, one := range urls {
		if strings.HasPrefix(one, "https") {
			isHttps = true
		}
	}

	var options []elastic.ClientOptionFunc

	options = append(options, elastic.SetURL(urls...))
	options = append(options, elastic.SetSniff(false))
	if isHttps {
		httpClient := &http.Client{}
		TLSClientConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		if this_.CertPath != "" {
			certPool := x509.NewCertPool()
			var pemCerts []byte
			pemCerts, err = ioutil.ReadFile(this_.CertPath)
			if err != nil {
				return
			}

			if !certPool.AppendCertsFromPEM(pemCerts) {
				err = errors.New("证书[" + this_.CertPath + "]解析失败")
				return
			}
			TLSClientConfig.RootCAs = certPool

			//TLSClientConfig.Certificates = []tls.Certificate{clicrt}

		}
		httpClient.Transport = &http.Transport{
			TLSClientConfig: TLSClientConfig,
		}
		options = append(options, elastic.SetHttpClient(httpClient))
	}
	if this_.Username != "" {
		options = append(options, elastic.SetBasicAuth(this_.Username, this_.Password))
	}
	client, err = elastic.NewClient(options...)
	if err != nil {
		if client != nil {
			client.Stop()
		}
		return
	}
	this_.client = client
	return
}

func (this_ *V7Service) Close() {
	if this_ != nil && this_.client != nil {
		this_.client.Stop()
		this_.client = nil
	}
}

func (this_ *V7Service) Info() (res *elastic.NodesInfoResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()
	res, err = client.NodesInfo().Do(context.Background())
	if err != nil {
		return
	}

	return
}

func (this_ *V7Service) DeleteIndex(indexName string) (err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()
	_, err = client.DeleteIndex(indexName).Do(context.Background())
	if err != nil {
		return
	}
	return
}

func (this_ *V7Service) CreateIndex(indexName string, bodyJSON map[string]interface{}) (err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()
	_, err = client.CreateIndex(indexName).BodyJson(bodyJSON).Do(context.Background())
	if err != nil {
		return
	}
	return
}

type IndexInfo struct {
	IndexName string `json:"indexName"`
}

func (this_ *V7Service) Indexes() (indexes []*IndexInfo, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()
	indexNames, err := client.IndexNames()
	if err != nil {
		return
	}

	sort.Slice(indexNames, func(i, j int) bool {
		return strings.ToLower(indexNames[i]) < strings.ToLower(indexNames[j]) //升序  即前面的值比后面的小 忽略大小写排序
	})
	for _, indexName := range indexNames {
		info := &IndexInfo{
			IndexName: indexName,
		}
		indexes = append(indexes, info)
	}
	return
}

func (this_ *V7Service) GetMapping(indexName string) (res interface{}, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()
	mappingMap, err := client.GetMapping().Index(indexName).Do(context.Background())
	if err != nil {
		return
	}
	for key, value := range mappingMap {
		if key == indexName {
			res = value
		}
	}
	return
}

func (this_ *V7Service) PutMapping(indexName string, bodyJSON map[string]interface{}) (err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()
	_, err = client.PutMapping().Index(indexName).BodyJson(bodyJSON).Do(context.Background())
	if err != nil {
		return
	}
	return
}

func (this_ *V7Service) SetFieldType(indexName string, fieldName string, fieldType string) (err error) {
	bodyJSON := map[string]interface{}{}
	bodyJSON["properties"] = map[string]interface{}{
		fieldName: map[string]interface{}{
			"type": fieldType,
		},
	}
	err = this_.PutMapping(indexName, bodyJSON)
	if err != nil {
		return
	}
	return
}

type SearchResult struct {
	TotalHits *elastic.TotalHits `json:"total,omitempty"`     // total number of hits found
	MaxScore  *float64           `json:"max_score,omitempty"` // maximum score of all hits
	Hits      []*HitData         `json:"hits,omitempty"`      // the actual hits returned
}

type HitData struct {
	Index   string `json:"_index,omitempty"`   // index name
	Type    string `json:"_type,omitempty"`    // type meta field
	Id      string `json:"_id,omitempty"`      // external or internal
	Uid     string `json:"_uid,omitempty"`     // uid meta field (see MapperService.java for all meta fields)
	Version *int64 `json:"_version,omitempty"` // version number, when Version is set to true in SearchService
	Source  string `json:"_source,omitempty"`  // stored document source
}

type Where struct {
	Name                    string `json:"name"`
	Value                   string `json:"value"`
	Before                  string `json:"before"`
	After                   string `json:"after"`
	CustomSql               string `json:"customSql"`
	SqlConditionalOperation string `json:"sqlConditionalOperation"`
	AndOr                   string `json:"andOr"`
}

type Order struct {
	Name    string `json:"name"`
	AscDesc string `json:"ascDesc"`
}

func (this_ *V7Service) Search(indexName string, pageIndex int, pageSize int, whereList []*Where, orderList []*Order) (res *SearchResult, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()

	doer := client.Search(indexName)
	var query = elastic.NewBoolQuery()
	for _, where := range whereList {
		var q elastic.Query
		var isNot = false
		switch where.SqlConditionalOperation {
		case "like":
			q = elastic.NewWildcardQuery(where.Name, "*"+where.Value+"*")
		case "not like":
			q = elastic.NewWildcardQuery(where.Name, "*"+where.Value+"*")
			isNot = true
		case "like start":
			q = elastic.NewWildcardQuery(where.Name, where.Value+"*")
		case "not like start":
			q = elastic.NewWildcardQuery(where.Name, where.Value+"*")
			isNot = true
		case "like end":
			q = elastic.NewWildcardQuery(where.Name, "*"+where.Value)
		case "not like end":
			q = elastic.NewWildcardQuery(where.Name, "*"+where.Value)
			isNot = true
		case "between":
			q = elastic.NewRangeQuery(where.Name).Gte(where.Before).Lte(where.After)
		case "not between":
			q = elastic.NewRangeQuery(where.Name).Gte(where.Before).Lte(where.After)
			isNot = true
		case "in":
			q = elastic.NewTermsQuery(where.Name, strings.Split(where.Value, ","))
		case "not in":
			q = elastic.NewTermsQuery(where.Name, strings.Split(where.Value, ","))
			isNot = true
		default:
			q = elastic.NewTermQuery(where.Name, where.Value)
		}
		var addQ elastic.Query
		if strings.Contains(where.Name, ".") {
			addQ = elastic.NewNestedQuery(where.Name[0:strings.LastIndex(where.Name, ".")], q)
		} else {
			addQ = q
		}
		if isNot {
			query.MustNot(addQ)
		} else {
			query.Must(addQ)
		}
	}

	doer.Query(query)

	for _, one := range orderList {
		switch one.AscDesc {
		case "ASC":
			doer.Sort(one.Name, true)
			break
		default:
			doer.Sort(one.Name, false)
			break
		}
	}
	//ss, _ := query.Source()
	//util.Logger.Info("es search", zap.Any("query", ss))

	doer.TrackTotalHits(true)

	searchResult, err := doer.Size(pageSize).From((pageIndex - 1) * pageSize).Do(context.Background())
	if err != nil {
		return
	}
	res = &SearchResult{}
	if searchResult.Hits != nil {
		res.TotalHits = searchResult.Hits.TotalHits
		res.MaxScore = searchResult.Hits.MaxScore
		for _, one := range searchResult.Hits.Hits {
			data := &HitData{
				Id:      one.Id,
				Type:    one.Type,
				Index:   one.Index,
				Uid:     one.Uid,
				Version: one.Version,
			}
			if one.Source != nil {
				bs, _ := json.Marshal(one.Source)
				if bs != nil {
					data.Source = string(bs)
				}

			}
			res.Hits = append(res.Hits, data)
		}
	}

	return
}

type InsertResponse struct {
	*elastic.IndexResponse
}

func (this_ *V7Service) Insert(indexName string, id string, doc interface{}) (res *InsertResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()
	doer := client.Index()
	indexResponse, err := doer.Index(indexName).Id(id).BodyJson(doc).Refresh("wait_for").Do(context.Background())
	if err != nil {
		return
	}
	res = &InsertResponse{
		IndexResponse: indexResponse,
	}
	return
}

func (this_ *V7Service) InsertNotWait(indexName string, id string, doc interface{}) (res *InsertResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()
	doer := client.Index()
	indexResponse, err := doer.Index(indexName).Id(id).BodyJson(doc).Do(context.Background())
	if err != nil {
		return
	}
	res = &InsertResponse{
		IndexResponse: indexResponse,
	}
	return
}

type InsertDoc struct {
	IndexName string
	Id        string
	Doc       interface{}
}

type BulkResponse struct {
	*elastic.BulkResponse
}

func (this_ *V7Service) BatchInsertNotWait(docs []*InsertDoc) (res *BulkResponse, err error) {
	if len(docs) == 0 {
		return
	}
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()
	bulk := client.Bulk()
	for _, doc := range docs {
		bulk.Add(elastic.NewBulkIndexRequest().Index(doc.IndexName).Id(doc.Id).Doc(doc.Doc))
	}

	bulkResponse, err := bulk.Do(context.Background())
	if err != nil {
		return
	}
	res = &BulkResponse{
		BulkResponse: bulkResponse,
	}
	return
}

type UpdateResponse struct {
	*elastic.UpdateResponse
}

func (this_ *V7Service) Update(indexName string, id string, doc interface{}) (res *UpdateResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()

	doer := client.Update()
	updateResponse, err := doer.Index(indexName).Id(id).Doc(doc).Refresh("wait_for").Do(context.Background())
	if err != nil {
		return
	}
	res = &UpdateResponse{
		UpdateResponse: updateResponse,
	}

	return
}

func (this_ *V7Service) UpdateNotWait(indexName string, id string, doc interface{}) (res *UpdateResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()

	doer := client.Update()
	updateResponse, err := doer.Index(indexName).Id(id).Doc(doc).Do(context.Background())
	if err != nil {
		return
	}
	res = &UpdateResponse{
		UpdateResponse: updateResponse,
	}

	return
}

type DeleteResponse struct {
	*elastic.DeleteResponse
}

func (this_ *V7Service) Delete(indexName string, id string) (res *DeleteResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()

	doer := client.Delete()
	deleteResponse, err := doer.Index(indexName).Id(id).Refresh("wait_for").Do(context.Background())
	if err != nil {
		return
	}
	res = &DeleteResponse{
		DeleteResponse: deleteResponse,
	}

	return
}

func (this_ *V7Service) DeleteNotWait(indexName string, id string) (res *DeleteResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()

	doer := client.Delete()
	deleteResponse, err := doer.Index(indexName).Id(id).Do(context.Background())
	if err != nil {
		return
	}
	res = &DeleteResponse{
		DeleteResponse: deleteResponse,
	}

	return
}

type BulkIndexByScrollResponse struct {
	*elastic.BulkIndexByScrollResponse
}

func (this_ *V7Service) Reindex(sourceIndexName string, toIndexName string) (res *BulkIndexByScrollResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()

	doer := client.Reindex()
	bulkIndexByScrollResponse, err := doer.Source(elastic.NewReindexSource().Index(sourceIndexName)).DestinationIndex(toIndexName).Refresh("true").Do(context.Background())
	if err != nil {
		return
	}
	res = &BulkIndexByScrollResponse{
		BulkIndexByScrollResponse: bulkIndexByScrollResponse,
	}

	return
}

type IndicesStatsResponse struct {
	*elastic.IndicesStatsResponse
}

func (this_ *V7Service) IndexStat(indexName string) (res *IndicesStatsResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()

	doer := client.IndexStats()
	response, err := doer.Index(indexName).Do(context.Background())
	if err != nil {
		return
	}
	res = &IndicesStatsResponse{
		IndicesStatsResponse: response,
	}

	return
}

func (this_ *V7Service) Scroll(indexName string, scrollId string, pageSize int, whereList []*Where, orderList []*Order) (res *SearchResult, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()

	doer := client.Scroll(indexName)
	query := elastic.NewBoolQuery()
	searchResult, err := doer.Query(query).Size(pageSize).ScrollId(scrollId).Do(context.Background())
	if err != nil {
		return
	}
	if searchResult.Hits != nil {
		res.TotalHits = searchResult.Hits.TotalHits
		res.MaxScore = searchResult.Hits.MaxScore
		for _, one := range searchResult.Hits.Hits {
			data := &HitData{
				Id:      one.Id,
				Type:    one.Type,
				Index:   one.Index,
				Uid:     one.Uid,
				Version: one.Version,
			}
			if one.Source != nil {
				bs, _ := json.Marshal(one.Source)
				if bs != nil {
					data.Source = string(bs)
				}

			}
			res.Hits = append(res.Hits, data)
		}
	}

	return
}

type IndexAliasResponse struct {
	*elastic.AliasResult
}

func (this_ *V7Service) IndexAlias(indexName string, aliasName string) (res *IndexAliasResponse, err error) {
	client, err := this_.GetClient()
	if err != nil {
		return
	}
	//defer client.Stop()

	doer := client.Alias()
	aliasResult, err := doer.Add(indexName, aliasName).Do(context.Background())
	if err != nil {
		return
	}
	res = &IndexAliasResponse{
		AliasResult: aliasResult,
	}

	return
}
