package main

import (
	"log"
	"net/http"
	"regexp"

	"codeberg.org/dankstuff/danktoolz/actions"
	"codeberg.org/dankstuff/danktoolz/config"
	"codeberg.org/dankstuff/danktoolz/handlers/curls"
	"codeberg.org/dankstuff/danktoolz/handlers/middlewares/contenttype"
	"codeberg.org/dankstuff/danktoolz/handlers/middlewares/ismobile"
	"codeberg.org/dankstuff/danktoolz/handlers/middlewares/pathname"
	"codeberg.org/dankstuff/danktoolz/handlers/pages"
	"codeberg.org/dankstuff/danktoolz/handlers/static"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	mjson "github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

var (
	minifyer *minify.M
	usecases *actions.Actions
)

func init() {
	minifyer = minify.New()
	minifyer.AddFunc("text/css", css.Minify)
	minifyer.AddFunc("text/html", html.Minify)
	minifyer.AddFunc("image/svg+xml", svg.Minify)
	minifyer.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	minifyer.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	minifyer.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)

	usecases = actions.New()
}

func main() {
	///
	/// PAGES
	///

	pages := pages.New(usecases)

	pagesHandler := http.NewServeMux()
	pagesHandler.HandleFunc("/", pages.HandleIndex)
	pagesHandler.HandleFunc("/my-ip", pages.HandleMyIP)
	pagesHandler.HandleFunc("/passwords", pages.HandlePasswords)
	pagesHandler.HandleFunc("/paste", pages.HandlePastebin)
	pagesHandler.HandleFunc("/urls", pages.HandleUrlShortner)
	pagesHandler.HandleFunc("/privacy", pages.HandlePrivacy)

	///
	/// CURL "PAGES"
	///

	curlPages := curls.New(usecases)

	curlsHandler := http.NewServeMux()
	curlsHandler.HandleFunc("/", curlPages.HandleMyIP)

	///
	/// WEB APIS
	///

	/*
		webApis := webapis.New(usecases)

		webApisHandler := http.NewServeMux()
		webApisHandler.HandleFunc("GET /lyrics", webApis.HandleGetSongLyrics)
		webApisHandler.HandleFunc("POST /lyrics", webApis.HandleSubmitLyrics)
		webApisHandler.HandleFunc("POST /auth", webApis.HandleAuthSubmitLyrics)
		webApisHandler.HandleFunc("GET /auth/confirm", webApis.HandleConfirmAuthSubmitLyrics)
	*/

	///
	/// APPLICATION HANDLER
	///

	applicationHandler := http.NewServeMux()

	applicationHandler.Handle("/", minifyer.Middleware(pathname.Handler(ismobile.Handler(contenttype.Html(pagesHandler)))))
	applicationHandler.Handle("/curl", curlsHandler)
	// applicationHandler.Handle("/api/", contenttype.Json(http.StripPrefix("/api", v1ApiHandler)))
	// applicationHandler.Handle("/api/web/", ismobile.Handler(http.StripPrefix("/api/web", webApisHandler)))

	///
	/// STATICS HANDLER
	///

	statics := static.New(minifyer)
	applicationHandler.HandleFunc("/robots.txt", statics.HandleRobots)
	applicationHandler.HandleFunc("/sitemap.xml", statics.HandleSitemap)
	applicationHandler.HandleFunc("/favicon.ico", statics.HandleFavicon)
	applicationHandler.Handle("/static/", statics.AssetsHandler())

	log.Printf("Starting web server at port %s", config.Env().Port)
	log.Fatalln(http.ListenAndServe(":"+config.Env().Port, applicationHandler))
}
