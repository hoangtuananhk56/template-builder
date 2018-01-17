
package main

import (
	// other packages
	// "github.com/golang/glog"
	// "os"
	// "os/signal"

	// 1. Config
	// _ "template-builder/tbs/config"

	// "template-builder/tbs/httpserver"
	// "template-builder/tbs/x/runtime"

	"context"
	"template-builder/tbs/config"
	"template-builder/tbs/init2"
)

// func main() {
// 	stop := make(chan os.Signal)
// 	signal.Notify(stop, os.Interrupt)
// 	runtime.MaxProc()
// 	httpserver.Run(stop)
// 	glog.Flush()
// }

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	init2.Start(ctx, config.ReadConfig())
	init2.Wait()
}
