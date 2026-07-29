package pages

import (
	"log"
	"net/http"
	"strings"

	"codeberg.org/dankstuff/danktoolz/views/layouts"
	"codeberg.org/dankstuff/danktoolz/views/pages"
	"codeberg.org/dankstuff/danktoolz/views/types"
	_ "github.com/a-h/templ"
)

type pagesHandler struct {
}

func New() *pagesHandler {
	return &pagesHandler{}
}

func (p *pagesHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".go") || strings.Contains(r.URL.Path, "admin") {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err := layouts.Default(types.PageProps{
		PageId: types.IndexPage,
		Title:  "Index",
	}, pages.Index()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

/*
func (p *pages) HandleAbout(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".go") || strings.Contains(r.URL.Path, "admin") {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err := layouts.Default(types.PageProps{
		PageId: types.AboutPage,
		Title:  "About",
	}, partials.About()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
*/
