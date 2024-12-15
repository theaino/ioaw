package main

import (
	"io/fs"
	_ "ioaw/routers"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Masterminds/sprig/v3"
)

func loadLocales() {
	filepath.Walk("conf/locale", func(path string, info fs.FileInfo, err error) error {
		if err != nil { return err }

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".ini") {
			locale := strings.TrimSuffix(info.Name(), ".ini")
			err := i18n.SetMessage(locale, path)
			if err != nil {
				logs.Error("Failed to load locale", locale, err)
			}
		}

		return nil
	})
}

func setupDB() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		orm.RegisterDataBase("default", "sqlite3", "data.db")
	} else {
		orm.RegisterDataBase("default", "postgres", dbUrl)
	}

	forceSyncdb := os.Getenv("FORCE_DB") == "1"

	err := orm.RunSyncdb("default", forceSyncdb, true)
	if err != nil {
		logs.Error("Failed to syncdb", err)
	}
}

func templateFuncs() {
	web.AddFuncMap("i18n", i18n.Tr)

	for key, value := range sprig.FuncMap() {
		web.AddFuncMap(key, value)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port != "" {
		beego.BConfig.Listen.HTTPPort, _ = strconv.Atoi(port)
	}

	web.SetStaticPath("/dist", "dist")

	setupDB()
	loadLocales()

	templateFuncs()

	beego.Run()
}

