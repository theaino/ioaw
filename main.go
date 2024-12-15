package main

import (
	_ "ioaw/routers"
	"os"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
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

	web.SetStaticPath("/static", "dist")

	beego.Run()
}

