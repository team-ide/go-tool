package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/elasticsearch"
	"github.com/team-ide/go-tool/kafka"
	"github.com/team-ide/go-tool/redis"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
)

func (this_ *Executor) dbToRedis() (err error) {

	util.Logger.Info("db to redis start")
	err = this_.forEachOwnersTables(func(owner *DbOwner, table *DbTable, from *DataSourceDb) (err error) {
		to := NewDataSourceRedis()
		to.ColumnList = from.ColumnList
		to.Service, err = redis.New(this_.To.RedisConfig)
		if err != nil {
			util.Logger.Error("redis client new error", zap.Error(err))
			return
		}
		err = DateMove(this_.Progress, from, to)
		return
	})
	util.Logger.Info("db to redis end")

	return
}

func (this_ *Executor) dbToKafka() (err error) {

	util.Logger.Info("db to kafka start")
	err = this_.forEachOwnersTables(func(owner *DbOwner, table *DbTable, from *DataSourceDb) (err error) {
		to := NewDataSourceKafka()
		to.ColumnList = from.ColumnList
		to.TopicName = table.To.TableName
		to.TopicGroupName = table.TopicGroupName
		to.TopicKey = table.TopicKey
		to.TopicValue = table.TopicValue
		to.TopicValueByData = this_.To.TopicValueByData
		to.FillColumn = this_.To.FillColumn
		to.Service, err = kafka.New(this_.To.KafkaConfig)
		if err != nil {
			util.Logger.Error("kafka client new error", zap.Error(err))
			return
		}
		err = DateMove(this_.Progress, from, to)
		return
	})
	util.Logger.Info("db to kafka end")

	return
}

func (this_ *Executor) dbToEs() (err error) {

	util.Logger.Info("db to es start")
	err = this_.forEachOwnersTables(func(owner *DbOwner, table *DbTable, from *DataSourceDb) (err error) {
		to := NewDataSourceEs()
		to.ColumnList = from.ColumnList
		to.IndexName = table.To.TableName
		to.IndexIdName = table.IndexIdName
		to.IndexIdScript = table.IndexIdScript
		to.FillColumn = this_.To.FillColumn
		to.Service, err = elasticsearch.New(this_.To.EsConfig)
		if err != nil {
			util.Logger.Error("elasticsearch client new error", zap.Error(err))
			return
		}
		err = DateMove(this_.Progress, from, to)
		return
	})
	util.Logger.Info("db to es end")

	return
}
func (this_ *Executor) dbToDb() (err error) {

	util.Logger.Info("db to db start")
	err = this_.forEachOwnersTables(func(owner *DbOwner, table *DbTable, from *DataSourceDb) (err error) {
		to := NewDataSourceDb()
		to.ColumnList = from.ColumnList
		to.ParamModel = this_.To.GetDialectParam()
		to.OwnerName = owner.To.OwnerName
		to.TableName = table.To.TableName
		to.Service = owner.toService
		to.FillColumn = this_.To.FillColumn
		err = DateMove(this_.Progress, from, to)
		return
	})
	util.Logger.Info("db to db end")

	return
}

func (this_ *Executor) dbToSql() (err error) {
	util.Logger.Info("db to sql start")
	err = this_.forEachOwnersTables(func(owner *DbOwner, table *DbTable, from *DataSourceDb) (err error) {
		to := NewDataSourceSql()
		to.ParamModel = this_.To.GetDialectParam()
		to.ColumnList = from.ColumnList
		to.OwnerName = owner.To.OwnerName
		to.TableName = table.To.TableName
		to.DialectType = this_.To.DialectType
		to.FillColumn = this_.To.FillColumn
		if this_.From.BySql {
			to.FilePath = this_.getFilePath("", this_.To.GetFileName(), "sql")
		} else {
			switch this_.To.SqlFileMergeType {
			case "", "owner":
				to.FilePath = this_.getFilePath("", owner.To.OwnerName, "sql")
				break
			case "one":
				to.FilePath = this_.getFilePath("", this_.To.GetFileName(), "sql")
				break
			case "table":
				owner.appended = false
				to.FilePath = this_.getFilePath(owner.To.OwnerName, table.To.TableName, "sql")
				break
			default:
				err = errors.New("不支持的SQL文件合并类型")
				return
			}
		}

		var sqlList []string
		// 需要建库
		if this_.From.ShouldOwner && !owner.appended {
			owner.appended = true
			ss, e := to.GetDialect().OwnerCreateSql(to.ParamModel, &dialect.OwnerModel{
				OwnerName: owner.To.OwnerName,
			})
			if e != nil {
				util.Logger.Error("建库语句生成失败", zap.Error(e))
			}
			sqlList = append(sqlList, ss...)
		}
		// 需要建表
		if this_.From.ShouldTable && !table.appended {
			table.appended = true
			ss, e := to.GetDialect().TableCreateSql(to.ParamModel, owner.To.OwnerName, table.GetToDialectTable())
			if e != nil {
				util.Logger.Error("建表语句生成失败", zap.Error(e))
			}
			sqlList = append(sqlList, ss...)
		}
		if len(sqlList) > 0 {
			this_.ReadCount.AddSuccess(int64(len(sqlList)))
			err = to.Write(this_.Progress, &Data{
				DataType: DataTypeSql,
				Total:    int64(len(sqlList)),
				SqlList:  sqlList,
			})
			if err != nil {
				return
			}
		}
		err = DateMove(this_.Progress, from, to)
		return
	})
	util.Logger.Info("db to sql end")
	return
}

func (this_ *Executor) dbToTxt() (err error) {

	util.Logger.Info("db to txt start")
	err = this_.forEachOwnersTables(func(owner *DbOwner, table *DbTable, from *DataSourceDb) (err error) {
		to := NewDataSourceTxt()
		to.ColumnList = from.ColumnList
		to.ColSeparator = this_.To.ColSeparator
		to.ReplaceCol = this_.To.ReplaceCol
		to.ReplaceLine = this_.To.ReplaceLine
		to.ShouldTrimSpace = this_.To.ShouldTrimSpace
		to.FillColumn = this_.To.FillColumn
		if this_.From.BySql {
			to.FilePath = this_.getFilePath("", this_.To.GetFileName(), this_.To.GetTxtFileType())
		} else {
			switch this_.To.FileNameSplice {
			case "", "/":
				to.FilePath = this_.getFilePath(owner.To.OwnerName, table.To.TableName, this_.To.GetTxtFileType())
				break
			default:
				to.FilePath = this_.getFilePath("", owner.To.OwnerName+this_.To.FileNameSplice+table.To.TableName, this_.To.GetTxtFileType())
				break
			}
		}
		err = DateMove(this_.Progress, from, to)
		return
	})
	util.Logger.Info("db to txt end")

	return
}

func (this_ *Executor) dbToExcel() (err error) {

	util.Logger.Info("db to excel start")
	err = this_.forEachOwnersTables(func(owner *DbOwner, table *DbTable, from *DataSourceDb) (err error) {
		to := NewDataSourceExcel()
		to.ColumnList = from.ColumnList
		to.ShouldTrimSpace = this_.To.ShouldTrimSpace
		to.FillColumn = this_.To.FillColumn
		if this_.From.BySql {
			to.FilePath = this_.getFilePath("", this_.To.GetFileName(), "xlsx")
		} else {
			switch this_.To.FileNameSplice {
			case "", "/":
				to.FilePath = this_.getFilePath(owner.To.OwnerName, table.To.TableName, "xlsx")
				break
			default:
				to.FilePath = this_.getFilePath("", owner.To.OwnerName+this_.To.FileNameSplice+table.To.TableName, "xlsx")
				break
			}
		}
		err = DateMove(this_.Progress, from, to)
		return
	})
	util.Logger.Info("db to excel end")

	return
}

func (this_ *Executor) forEachOwnersTables(on func(owner *DbOwner, table *DbTable, datasource *DataSourceDb) (err error)) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			util.Logger.Error("for each owners to do panic error", zap.Error(err))
		}
	}()
	if this_.ShouldStop() {
		return
	}

	this_.From.dbService, err = this_.newDbService(*this_.From.DbConfig, this_.From.Username, this_.From.Password, this_.From.OwnerName)
	if err != nil {
		return
	}
	defer func() {
		if this_.From.dbService != nil {
			_ = this_.From.dbService.GetDb().Close()
		}
	}()
	if this_.To.IsDb() {
		this_.To.dbService, err = this_.newDbService(*this_.To.DbConfig, this_.To.Username, this_.To.Password, this_.To.OwnerName)
		if err != nil {
			return
		}
		defer func() {
			if this_.To.dbService != nil {
				_ = this_.To.dbService.GetDb().Close()
			}
		}()
	}

	if this_.From.BySql {

		from := NewDataSourceDb()
		from.ShouldSelectPage = this_.From.ShouldSelectPage
		if len(this_.From.ColumnList) == 0 && len(this_.To.ColumnList) == 0 {

			str := strings.TrimSpace(this_.From.SelectSql)
			if this_.From.ShouldSelectPage {
				str = this_.From.GetDialect().PackPageSql(str, 1, 1)
			}

			rows, e := this_.From.dbService.GetDb().Query(str)
			if e == nil {
				defer func() { _ = rows.Close() }()
				columnTypes, _ := rows.ColumnTypes()
				if columnTypes != nil {
					for _, columnType := range columnTypes {
						column := &Column{
							ColumnModel: &dialect.ColumnModel{},
						}
						column.ColumnName = columnType.Name()
						if precision, scale, ok := columnType.DecimalSize(); ok {
							column.ColumnPrecision = int(precision)
							column.ColumnScale = int(scale)
						}
						column.ColumnDataType = columnType.DatabaseTypeName()
						if length, ok := columnType.Length(); ok {
							column.ColumnLength = int(length)
						}
						if nullable, ok := columnType.Nullable(); ok {
							column.ColumnNotNull = !nullable
						}

						this_.From.ColumnList = append(this_.From.ColumnList, column)
						this_.To.ColumnList = append(this_.To.ColumnList, column)
					}
				}
			}

		}

		owner := &DbOwner{
			From: &dialect.OwnerModel{
				OwnerName: this_.From.OwnerName,
			},
			To: &dialect.OwnerModel{
				OwnerName: this_.To.OwnerName,
			},
		}
		table := &DbTable{
			From: &dialect.TableModel{
				TableName: this_.From.TableName,
			},
			To: &dialect.TableModel{
				TableName: this_.To.TableName,
			},
			IndexIdName:   this_.To.IndexIdName,
			IndexIdScript: this_.To.IndexIdScript,
		}
		from.ParamModel = this_.From.GetDialectParam()
		from.OwnerName = owner.From.OwnerName
		from.TableName = table.From.TableName
		from.ColumnList = this_.From.ColumnList
		from.Service = this_.From.dbService
		from.SelectSql = this_.From.SelectSql
		from.FillColumn = this_.From.FillColumn
		err = on(owner, table, from)

		return
	}

	util.Logger.Info("for each owners to do", zap.Any("allOwner", this_.From.AllOwner))
	owners := this_.From.Owners

	if this_.From.AllOwner {
		var list []*dialect.OwnerModel
		list, err = worker.OwnersSelect(this_.From.dbService.GetDb(), this_.From.GetDialect(), this_.From.GetDialectParam())
		if err != nil {
			return
		}

		for _, one := range list {
			var find *DbOwner

			for _, o := range owners {
				if o.From == nil {
					continue
				}
				if strings.ToLower(o.From.OwnerName) == strings.ToLower(one.OwnerName) {
					find = o
					break
				}
			}
			if find == nil {
				owners = append(owners, &DbOwner{
					From:     one,
					AllTable: true,
				})
			} else {
				find.From.OwnerName = one.OwnerName
			}

		}
	}

	var newOwners []*DbOwner
	for _, o := range owners {
		if o.From == nil || o.From.OwnerName == "" {
			continue
		}
		if util.StringIndexOf(this_.From.SkipOwnerNames, o.From.OwnerName) >= 0 {
			continue
		}
		if o.To == nil {
			o.To = &dialect.OwnerModel{}
		}
		if o.To.OwnerName == "" {
			o.To.OwnerName = o.From.OwnerName
		}
		newOwners = append(newOwners, o)
	}
	owners = newOwners

	util.Logger.Info("for each owners to do", zap.Any("owners", len(owners)))
	if len(owners) == 0 {
		return
	}

	this_.OwnerTotal += int64(len(owners))

	for _, owner := range owners {
		if this_.ShouldStop() {
			return
		}
		e := this_.forEachOwnerTables(owner, on)
		if e != nil {
			this_.OwnerCount.AddError(1, e)
			if !this_.ErrorContinue {
				err = e
				return
			}
		} else {
			this_.OwnerCount.AddSuccess(1)
		}
	}
	return
}

func (this_ *Executor) forEachOwnerTables(owner *DbOwner, on func(owner *DbOwner, table *DbTable, datasource *DataSourceDb) (err error)) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			util.Logger.Error("for each owner tables to do panic error", zap.Error(err))
		}
	}()
	if this_.ShouldStop() {
		return
	}
	util.Logger.Info("for each owner tables to do", zap.Any("owner", owner.From.OwnerName))

	owner.fromService, err = this_.newDbService(*this_.From.DbConfig, owner.From.OwnerUsername, owner.From.OwnerPassword, owner.From.OwnerName)
	if err != nil {
		return
	}
	if this_.To.IsDb() {

		// 需要建库
		if this_.From.ShouldOwner && !owner.appended {
			owner.appended = true
			// 同步结构
			var toOwnerOne *dialect.OwnerModel
			toOwnerOne, err = worker.OwnerSelect(this_.To.dbService.GetDb(), this_.To.GetDialect(), this_.To.GetDialectParam(), owner.To.OwnerName)
			if err != nil {
				return
			}
			if toOwnerOne == nil {
				_, err = worker.OwnerCreate(this_.To.dbService.GetDb(), this_.To.GetDialect(), this_.To.GetDialectParam(), &dialect.OwnerModel{
					OwnerName:             owner.To.OwnerName,
					OwnerPassword:         owner.To.OwnerPassword,
					OwnerCharacterSetName: "utf8mb4",
				})
				if err != nil {
					return
				}
			}
		}

		owner.toService, err = this_.newDbService(*this_.To.DbConfig, owner.To.OwnerUsername, owner.To.OwnerPassword, owner.To.OwnerName)
		if err != nil {
			return
		}
	}
	defer func() {

		if owner.fromService != nil {
			_ = owner.fromService.GetDb().Close()
			owner.fromService = nil
		}

		if owner.toService != nil {
			_ = owner.toService.GetDb().Close()
			owner.toService = nil
		}

	}()

	tables := owner.Tables

	if owner.AllTable {
		var list []*dialect.TableModel
		list, err = worker.TablesSelect(owner.fromService.GetDb(), owner.fromService.GetDialect(), this_.From.GetDialectParam(), owner.From.OwnerName)
		if err != nil {
			return
		}

		for _, one := range list {
			var find *DbTable

			for _, o := range tables {
				if o.From == nil || o.From.TableName == "" {
					continue
				}
				if strings.ToLower(o.From.TableName) == strings.ToLower(one.TableName) {
					find = o
					break
				}
			}
			if find == nil {
				if owner.AllTable {
					tables = append(tables, &DbTable{
						From:      one,
						AllColumn: true,
					})
				}
			} else {
				find.From.TableName = one.TableName
			}
		}
	}

	var newList []*DbTable
	for _, o := range tables {
		if o.From == nil || o.From.TableName == "" {
			continue
		}
		if util.StringIndexOf(owner.SkipTableNames, o.From.TableName) >= 0 {
			continue
		}
		if o.To == nil {
			o.To = &dialect.TableModel{}
		}
		if o.To.TableName == "" {
			o.To.TableName = o.From.TableName
		}
		newList = append(newList, o)
	}
	tables = newList

	util.Logger.Info("for each owner tables to do", zap.Any("tables", len(tables)))
	if len(tables) == 0 {
		return
	}

	this_.TableTotal += int64(len(tables))

	for _, table := range tables {
		if this_.ShouldStop() {
			return
		}

		e := this_.doOwnerTable(owner, table, on)
		if e != nil {
			this_.TableCount.AddError(1, e)
			if !this_.ErrorContinue {
				err = e
				return
			}
		} else {
			this_.TableCount.AddSuccess(1)
		}
	}

	return
}

func (this_ *Executor) doOwnerTable(owner *DbOwner, table *DbTable, on func(owner *DbOwner, table *DbTable, datasource *DataSourceDb) (err error)) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			util.Logger.Error("owner table to do panic error", zap.Error(err))
		}
	}()
	if this_.ShouldStop() {
		return
	}
	util.Logger.Info("owner table to do", zap.Any("ownerName", owner.From.OwnerName), zap.Any("tableName", table.From.TableName))

	var detail *dialect.TableModel
	detail, _ = worker.TableDetail(owner.fromService.GetDb(), owner.fromService.GetDialect(), this_.From.GetDialectParam(), owner.From.OwnerName, table.From.TableName, true)

	if detail != nil {
		table.From = detail

		for _, one := range detail.ColumnList {
			var find *DbColumn

			for _, o := range table.Columns {
				if o.From == nil || o.From.ColumnName == "" {
					continue
				}
				if strings.ToLower(o.From.ColumnName) == strings.ToLower(one.ColumnName) {
					find = o
					break
				}
			}
			if find == nil {
				if table.AllColumn {
					table.Columns = append(table.Columns, &DbColumn{
						From: one,
					})
				}
			} else {
				find.From = one
			}
		}
	}

	util.Logger.Info("owner table to do", zap.Any("columns", len(table.Columns)))

	var newList []*DbColumn
	for _, o := range table.Columns {
		if o.From == nil || o.From.ColumnName == "" {
			continue
		}
		if util.StringIndexOf(table.SkipColumnNames, o.From.ColumnName) >= 0 {
			continue
		}
		if o.To == nil {
			o.To = o.From
		}
		newList = append(newList, o)
	}
	table.Columns = newList

	if len(table.Columns) == 0 {
		return
	}

	datasource := NewDataSourceDb()
	datasource.ParamModel = this_.From.GetDialectParam()
	datasource.OwnerName = owner.From.OwnerName
	datasource.TableName = table.From.TableName
	datasource.Service = owner.fromService
	datasource.FillColumn = this_.From.FillColumn

	for _, c := range table.Columns {
		datasource.ColumnList = append(datasource.ColumnList, &Column{
			ColumnModel: c.From,
			Value:       c.Value,
		})
	}

	if this_.To.IsDb() {
		// 需要建表
		if this_.From.ShouldTable && !table.appended {
			table.appended = true
			// 同步结构
			var oldTableDetail *dialect.TableModel
			oldTableDetail, err = worker.TableDetail(
				this_.To.dbService.GetDb(),
				this_.To.GetDialect(),
				this_.To.GetDialectParam(),
				owner.To.OwnerName,
				table.To.TableName,
				false,
			)
			if err != nil {
				util.Logger.Error("to owner table select detail error", zap.Error(err))
				return
			}
			if oldTableDetail == nil {
				oldTableDetail = &dialect.TableModel{}
			}
			if len(oldTableDetail.ColumnList) == 0 {
				err = worker.TableCreate(
					this_.To.dbService.GetDb(),
					this_.To.GetDialect(),
					this_.To.GetDialectParam(),
					owner.To.OwnerName,
					table.GetToDialectTable(),
				)
				if err != nil {
					util.Logger.Error("to owner table create table error", zap.Error(err))
					return
				}
			} else {
				err = worker.TableUpdate(
					this_.To.dbService.GetDb(),
					this_.To.GetDialect(),
					oldTableDetail,
					this_.To.GetDialect(),
					table.GetToDialectTable(),
				)
				if err != nil {
					util.Logger.Error("to owner table update table error", zap.Error(err))
					return
				}
			}
		}
	}

	err = on(owner, table, datasource)
	if err != nil {
		return
	}

	return
}

func (this_ *Executor) newDbService(config db.Config, username string, password string, ownerName string) (service db.IService, err error) {
	config.MaxIdleConn = 3
	config.MaxOpenConn = 3
	if username != "" {
		config.Username = username
	}
	if password != "" {
		config.Password = password
	}
	databaseType := db.GetDatabaseType(config.Type)
	switch databaseType.GetDialect().DialectType() {
	case dialect.TypeMysql:
		config.Database = ownerName
		break
	case dialect.TypeGBase:

		if ownerName != "" {
			var keyL = len("db=")
			index := strings.Index(strings.ToLower(config.OdbcDsn), "db=")

			if index < 0 {
				index = strings.Index(strings.ToLower(config.OdbcDsn), "database=")
				keyL = len("database=")
			}
			if index >= 0 {
				beforeStr := config.OdbcDsn[0 : index+keyL]
				afterStr := ""
				str := config.OdbcDsn[index+keyL:]
				index = strings.Index(str, ";")
				if index >= 0 {
					afterStr = str[index:]
				}
				config.OdbcDsn = beforeStr + ownerName + afterStr
			}
		}

		break
	default:
		config.Schema = ownerName
		break
	}

	service, err = db.New(&config)
	if err != nil {
		util.Logger.Error("db service new error", zap.Any("host", config.Host), zap.Any("username", config.Username), zap.Any("ownerName", ownerName), zap.Error(err))
		return
	}
	return
}
