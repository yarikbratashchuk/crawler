package main

import (
	"io"

	"github.com/btcsuite/btclog"
	"github.com/yarikbratashchuk/crawler/crawler"
	"github.com/yarikbratashchuk/crawler/data"
	"github.com/yarikbratashchuk/crawler/linkstorage"
)

var log btclog.Logger

func setupLog(dest io.Writer, loglevel string) {
	logBackend := btclog.NewBackend(dest)
	lvl, _ := btclog.LevelFromString(loglevel)

	log = logBackend.Logger("MAIN")
	dataLog := logBackend.Logger("DATA")
	linkLog := logBackend.Logger("LNKS")
	crawlerLog := logBackend.Logger("CRWL")

	log.SetLevel(lvl)
	dataLog.SetLevel(lvl)
	linkLog.SetLevel(lvl)
	crawlerLog.SetLevel(lvl)

	data.UseLogger(dataLog)
	linkstorage.UseLogger(linkLog)
	crawler.UseLogger(crawlerLog)
}
