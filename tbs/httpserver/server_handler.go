package httpserver

import (
	"http/gziphandler"
	"http/static/upload"
	"http/static/vstatic"
	"net/http"
	"template-builder/tbs/api"
	"template-builder/tbs/view"
	"regexp"
)

func webAssetGzipHandler(handler http.Handler) http.Handler {
	gzip := gziphandler.GzipHandler(handler)
	assetRegex, _ := regexp.Compile(".(js|css)$")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if assetRegex.MatchString(r.URL.Path) {
			gzip.ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func (phs *ProjectHttpServer) addStaticHandler(s *http.ServeMux) {
	p := phs.pc
	staticConfig := p.Station.Static

	// s.Handle("/", http.RedirectHandler("/app/", http.StatusFound))

	var app = vstatic.NewVersionStatic(staticConfig.AppFolder)

	s.Handle("/app/", http.StripPrefix("/app", webAssetGzipHandler(app)))
	var device = vstatic.NewVersionStatic(staticConfig.DeviceFolder)
	s.Handle("/device/", http.StripPrefix("/device", webAssetGzipHandler(device)))

	// storage
	// storageConfig := p.Station.Storage

	s.Handle("/static/", http.StripPrefix("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))

	// s.Handle("/upload/", upload.NewUploadProxy(p.Station.GetUploadLink(), storageConfig.Upload, "/static/upload"))

	var up = upload.NewUploadFileServer("static/upload", 40960000)
	s.Handle("/upload/", http.StripPrefix("/upload/", up))

	s.Handle("/", http.StripPrefix("/v", view.NewViewServer()))

}

func (phs *ProjectHttpServer) makeHandler() http.Handler {
	var server = http.NewServeMux()
	phs.addStaticHandler(server)
	// application specific
	apiServer := api.NewApiServer()
	server.Handle("/api/",
		gziphandler.GzipHandler(http.StripPrefix("/api", apiServer)),
	)
	go func() {
		phs.ready <- struct{}{}
	}()
	return server
}
