package main

import (
	_ "ioaw/routers"
	"os"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port != "" {
		beego.BConfig.Listen.HTTPPort, _ = strconv.Atoi(port)
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		orm.RegisterDataBase("default", "sqlite3", "data.db")
	} else {
		orm.RegisterDataBase("default", "postgres", dbUrl)
	}

	orm.RunSyncdb("default", false, true)

	beego.Run()
}

