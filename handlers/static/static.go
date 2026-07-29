package static

import (
	"embed"
	"net/http"
	"strings"

	"codeberg.org/dankstuff/danktoolz/static/assets"
	"codeberg.org/dankstuff/danktoolz/static/meta"
	"github.com/tdewolff/minify/v2"
)

type staticHandler struct {
	minifyer *minify.M
	assetsFS embed.FS
	metaFS   embed.FS
}

func New(minifyer *minify.M) *staticHandler {
	return &staticHandler{
		minifyer: minifyer,
		assetsFS: assets.FS(),
		metaFS:   meta.FS(),
	}
}

func (s *staticHandler) HandleRobots(w http.ResponseWriter, r *http.Request) {
	robotsFile, _ := s.metaFS.ReadFile("robots.txt")
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write(robotsFile)
}

type sitemapValue struct {
	PublicId string
	LastMod  string
}

func (s *staticHandler) HandleSitemap(w http.ResponseWriter, r *http.Request) {
	sitemapFile, _ := s.metaFS.ReadFile("sitemap.xml")
	w.Header().Set("Content-Type", "text/plain")

	w.Header().Set("Cache-Control", "public, max-age=15")
	w.Header().Set("Content-Type", "application/xml")

	_, _ = w.Write(sitemapFile)
}

func (s *staticHandler) HandleFavicon(w http.ResponseWriter, r *http.Request) {
	faviconFile, _ := s.assetsFS.ReadFile("favicon.ico")
	w.Header().Set("Content-Type", "image/x-icon")
	_, _ = w.Write(faviconFile)
}

func (s *staticHandler) AssetsHandler() http.Handler {
	return http.StripPrefix("/static", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ".go") || strings.Contains(r.URL.Path, "admin") {
			http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusTemporaryRedirect)
			return
		}

		w.Header().Set("Cache-Control", "public, max-age=7200, stale-while-revalidate=5")

		s.minifyer.Middleware(http.FileServer(http.FS(s.assetsFS))).
			ServeHTTP(w, r)
	}))
}
