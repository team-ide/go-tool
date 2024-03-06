package mongodb

import (
	"fmt"
	"github.com/team-ide/go-tool/util"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestMongodb(t *testing.T) {
	config := &Config{
		Address: "192.168.0.53:27017",
	}
	service, err := New(config)
	if err != nil {
		panic(err)
	}
	databases, totalSize, err := service.Databases()
	if err != nil {
		panic(err)
	}
	fmt.Println("totalSize:", totalSize)
	for _, database := range databases {
		fmt.Println("database:" + util.GetStringValue(database))

		var collections []*Collection
		collections, err = service.Collections(database.Name)
		if err != nil {
			panic(err)
		}
		for _, collection := range collections {
			fmt.Println("collection:" + util.GetStringValue(collection))

			var indexes []map[string]interface{}
			indexes, err = service.Indexes(database.Name, collection.Name)
			if err != nil {
				panic(err)
			}
			for _, one := range indexes {
				fmt.Println("index:" + util.GetStringValue(one))
			}
			var dataList []map[string]interface{}
			dataList, err = service.QueryMap(database.Name, collection.Name, bson.M{}, nil)
			if err != nil {
				panic(err)
			}
			for _, one := range dataList {
				fmt.Println("data:" + util.GetStringValue(one))
			}
		}
	}
	dataList, err := service.QueryMap("test_db", "test_tb", bson.M{}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("dataList:", dataList)
	for _, d := range dataList {
		fmt.Println("data:" + util.GetStringValue(d))
	}

}
