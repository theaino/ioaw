package main

import (
	_ "ioaw/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")

	orm.RunSyncdb("default", false, true)

	beego.Run()
}

