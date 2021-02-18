package main

import (
	"log"

	"flag"

	"path/filepath"

	//	"github.com/davecgh/go-spew/spew"
	"github.com/schollz/jsonstore"

	"os"

	"github.com/mamemomonga/go-study-simple-toot/src/don"
)

func main() {
	// configファイル名を取得
	var configFilename string
	flag.StringVar(&configFilename, "c", "", "Config File")
	flag.Parse()

	// configファイルロード
	cfg, err := configLoad(configFilename)
	if err != nil {
		log.Fatal(err)
	}
	// spew.Dump(cfg)

	// configファイルと同じフォルダにservers.jsonを置く
	serversFilename := filepath.Dir(configFilename) + "/servers.json"

	// serversファイルのロードもしくは新規作成
	var servers *jsonstore.JSONStore
	if _, err := os.Stat(serversFilename); err == nil {
		servers, err = jsonstore.Open(serversFilename)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		servers = new(jsonstore.JSONStore)
	}

	// donインスタンス作成
	don := don.New(don.Config{
		ClientName: cfg.ClientName,
		UserLogin:  cfg.Mastodon,
		Store:      servers,
	})

	// アプリ登録
	if save, err := don.Register(); err == nil {
		if save {
			jsonstore.Save(servers, serversFilename)
		}
	} else {
		log.Fatal(err)
	}

	// トゥート
	don.Toot("テストだぬ〜ん")
}
