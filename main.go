package main

import (
	_ "github.com/team-ide/go-dialect/dialect"
	_ "github.com/team-ide/go-dialect/worker"
	_ "github.com/team-ide/go-tool/db"
	_ "github.com/team-ide/go-tool/db/db_type_dm"
	_ "github.com/team-ide/go-tool/db/db_type_gbase"
	_ "github.com/team-ide/go-tool/db/db_type_kingbase"
	_ "github.com/team-ide/go-tool/db/db_type_mysql"
	_ "github.com/team-ide/go-tool/db/db_type_odbc"
	_ "github.com/team-ide/go-tool/db/db_type_opengauss"
	_ "github.com/team-ide/go-tool/db/db_type_oracle"
	_ "github.com/team-ide/go-tool/db/db_type_postgresql"
	_ "github.com/team-ide/go-tool/db/db_type_shentong"
	_ "github.com/team-ide/go-tool/db/db_type_sqlite"
	_ "github.com/team-ide/go-tool/elasticsearch"
	_ "github.com/team-ide/go-tool/javascript"
	_ "github.com/team-ide/go-tool/kafka"
	_ "github.com/team-ide/go-tool/redis"
	_ "github.com/team-ide/go-tool/util"
	_ "github.com/team-ide/go-tool/zookeeper"
)

func main() {

}
