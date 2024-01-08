package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
)

func (this_ *Executor) dbToDb() (err error) {

	return
}

func (this_ *Executor) dbToSql() (err error) {
	util.Logger.Info("db to sql start")
	err = this_.forEachOwnersTables(true, func(owner *DbOwner, table *DbTable, from *DataSourceDb) (err error) {

		return
	})
	util.Logger.Info("db to sql end")
	return
}

func (this_ *Executor) dbToTxt() (err error) {

	util.Logger.Info("db to txt start")
	err = this_.forEachOwnersTables(true, func(owner *DbOwner, table *DbTable, from *DataSourceDb) (err error) {
		to := NewDataSourceTxt()
		to.ColumnList = from.ColumnList
		to.FilePath = this_.Dir + owner.TargetName + "." + table.TargetName + ".txt"
		err = DateMove(this_.Progress, from, to)
		return
	})
	util.Logger.Info("db to txt end")

	return
}

func (this_ *Executor) dbToExcel() (err error) {

	return
}

func (this_ *Executor) forEachOwnersTables(isSource bool, on func(owner *DbOwner, table *DbTable, datasource *DataSourceDb) (err error)) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			util.Logger.Error("for each owners to do panic error", zap.Error(err))
		}
	}()
	util.Logger.Info("for each owners to do", zap.Any("isSource", isSource), zap.Any("allOwner", this_.AllOwner))
	owners := this_.Owners

	if this_.AllOwner && isSource {
		var list []*dialect.OwnerModel
		var service db.IService
		service, err = this_.newDbService(*this_.Source.DbConfig, "", "", "")
		if err != nil {
			return
		}
		defer func() {
			if service != nil {
				service.Close()
			}
		}()
		list, err = worker.OwnersSelect(service.GetDb(), service.GetDialect(), this_.GetDialectParam())
		if err != nil {
			return
		}

		service.Close()
		service = nil

		for _, one := range list {
			var find *DbOwner

			for _, o := range owners {
				if strings.ToLower(o.SourceName) == strings.ToLower(one.OwnerName) {
					find = o
					break
				}
			}
			if find == nil {
				owners = append(owners, &DbOwner{
					SourceName: one.OwnerName,
					AllTable:   true,
				})
			} else {
				find.SourceName = one.OwnerName
			}

		}
	}

	if isSource {
		var newOwners []*DbOwner
		for _, o := range owners {
			if util.StringIndexOf(this_.SkipOwnerNames, o.SourceName) >= 0 {
				continue
			}
			newOwners = append(newOwners, o)
		}
		owners = newOwners
	}

	util.Logger.Info("for each owners to do", zap.Any("owners", len(owners)))
	if len(owners) == 0 {
		return
	}

	this_.OwnerTotal += int64(len(owners))

	for _, owner := range owners {
		if owner.SourceName == "" {
			e := errors.New("库名未配置")
			this_.OwnerCount.AddError(1, e)
			if !this_.ErrorContinue {
				err = e
				return
			}
		}
		if owner.TargetName == "" {
			owner.TargetName = owner.SourceName
		}

		e := this_.forEachOwnerTables(isSource, owner, on)
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

func (this_ *Executor) forEachOwnerTables(isSource bool, owner *DbOwner, on func(owner *DbOwner, table *DbTable, datasource *DataSourceDb) (err error)) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			util.Logger.Error("for each owner tables to do panic error", zap.Error(err))
		}
	}()
	util.Logger.Info("for each owner tables to do", zap.Any("isSource", isSource), zap.Any("owner", owner.SourceName))

	if isSource {
		owner.sourceService, err = this_.newDbService(*this_.Source.DbConfig, owner.Username, owner.Username, owner.SourceName)
	} else {
		owner.targetService, err = this_.newDbService(*this_.Target.DbConfig, owner.Username, owner.Username, owner.TargetName)
	}
	if err != nil {
		return
	}
	defer func() {
		if isSource {
			if owner.sourceService != nil {
				owner.sourceService.Close()
				owner.sourceService = nil
			}
		} else {
			if owner.targetService != nil {
				owner.targetService.Close()
				owner.targetService = nil
			}
		}
	}()

	tables := owner.Tables

	if owner.AllTable && isSource {
		var list []*dialect.TableModel
		list, err = worker.TablesSelect(owner.sourceService.GetDb(), owner.sourceService.GetDialect(), this_.GetDialectParam(), owner.SourceName)
		if err != nil {
			return
		}

		for _, one := range list {
			var find *DbTable

			for _, o := range tables {
				if strings.ToLower(o.SourceName) == strings.ToLower(one.TableName) {
					find = o
					break
				}
			}
			if find == nil {
				tables = append(tables, &DbTable{
					SourceName: one.TableName,
					AllColumn:  true,
				})
			} else {
				find.SourceName = one.TableName
			}
		}
	}

	if isSource {
		var newList []*DbTable
		for _, o := range tables {
			if util.StringIndexOf(owner.SkipTableNames, o.SourceName) >= 0 {
				continue
			}
			newList = append(newList, o)
		}
		tables = newList
	}

	util.Logger.Info("for each owner tables to do", zap.Any("tables", len(tables)))
	if len(tables) == 0 {
		return
	}

	this_.TableTotal += int64(len(tables))

	for _, table := range tables {

		if table.SourceName == "" {
			e := errors.New("表名未配置")
			this_.TableCount.AddError(1, e)
			if !this_.ErrorContinue {
				err = e
				return
			}
		}
		if table.TargetName == "" {
			table.TargetName = table.SourceName
		}

		e := this_.doOwnerTable(isSource, owner, table, on)
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

func (this_ *Executor) doOwnerTable(isSource bool, owner *DbOwner, table *DbTable, on func(owner *DbOwner, table *DbTable, datasource *DataSourceDb) (err error)) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			util.Logger.Error("owner table to do panic error", zap.Error(err))
		}
	}()
	util.Logger.Info("owner table to do", zap.Any("isSource", isSource), zap.Any("ownerName", owner.SourceName), zap.Any("tableName", table.SourceName))

	columns := table.Columns

	var detail *dialect.TableModel
	detail, _ = worker.TableDetail(owner.sourceService.GetDb(), owner.sourceService.GetDialect(), this_.GetDialectParam(), owner.SourceName, table.SourceName, true)

	if detail != nil {

		for _, one := range detail.ColumnList {
			var find *DbColumn

			for _, o := range columns {
				if strings.ToLower(o.SourceName) == strings.ToLower(one.TableName) {
					find = o
					break
				}
			}
			if find == nil {
				if table.AllColumn && isSource {
					columns = append(columns, &DbColumn{
						SourceName: one.OwnerName,
						Column: &Column{
							ColumnModel: one,
						},
					})
				}
			} else {
				find.SourceName = one.ColumnName
				if find.Column == nil {
					find.Column = &Column{}
				}
				find.Column.ColumnModel = one
			}
		}
	}

	if isSource {
		util.Logger.Info("owner table to do", zap.Any("columns", len(columns)))

		var newList []*DbColumn
		for _, o := range columns {
			if util.StringIndexOf(table.SkipColumnNames, o.SourceName) >= 0 {
				continue
			}
			newList = append(newList, o)
		}
		columns = newList
	}

	if len(columns) == 0 {
		return
	}

	if isSource {
		datasource := NewDataSourceDb()
		datasource.ParamModel = this_.GetDialectParam()
		datasource.OwnerName = owner.SourceName
		datasource.TableName = table.SourceName
		datasource.Service = owner.sourceService

		for _, c := range columns {
			datasource.ColumnList = append(datasource.ColumnList, c.Column)
		}

		err = on(owner, table, datasource)
		if err != nil {
			return
		}
	}

	return
}

func (this_ *Executor) newDb() (err error) {

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
