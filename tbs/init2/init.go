package init2

import (
	"context"
	"db/mgo"
	"template-builder/tbs/config"
	"template-builder/tbs/httpserver"
	"util/runtime"

	"github.com/golang/glog"
)

func initialize(ctx context.Context) {
	mgo.Start(ctx)
}

func Start(ctx context.Context, p *config.ProjectConfig) {
	runtime.MaxProc()
	server = httpserver.NewProjectHttpServer(p)
	initialize(ctx)
}

func Wait() {
	defer beforeExit()
	server.Wait()
}

func beforeExit() {
	runtime.Recover()
	glog.Flush()
}
