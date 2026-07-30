package pages

import (
	"log"
	"net/http"
	"strings"

	"codeberg.org/dankstuff/danktoolz/actions"
	"codeberg.org/dankstuff/danktoolz/views/layouts"
	"codeberg.org/dankstuff/danktoolz/views/pages"
	"codeberg.org/dankstuff/danktoolz/views/types"
	_ "github.com/a-h/templ"
)

type pagesHandler struct {
	usecases *actions.Actions
}

func New(usecases *actions.Actions) *pagesHandler {
	return &pagesHandler{
		usecases: usecases,
	}
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

func (p *pagesHandler) HandleMyIP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".go") || strings.Contains(r.URL.Path, "admin") {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusTemporaryRedirect)
		return
	}

	ipInfo := actions.ExtractClientMetadata(r)

	w.Header().Set("Content-Type", "text/html")
	err := layouts.Default(types.PageProps{
		PageId: types.MyIpPage,
		Title:  "My IP",
	}, pages.MyIP(ipInfo)).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (p *pagesHandler) HandlePasswords(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".go") || strings.Contains(r.URL.Path, "admin") {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err := layouts.Default(types.PageProps{
		PageId: types.PasswordsPage,
		Title:  "My IP",
	}, pages.Passwords()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (p *pagesHandler) HandleClip(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".go") || strings.Contains(r.URL.Path, "admin") {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err := layouts.Default(types.PageProps{
		PageId: types.ClipPage,
		Title:  "Clip",
	}, pages.Clip()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (p *pagesHandler) HandleUrlShortner(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".go") || strings.Contains(r.URL.Path, "admin") {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err := layouts.Default(types.PageProps{
		PageId: types.UrlsPage,
		Title:  "URL Sohrtner",
	}, pages.UrlShortner()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (p *pagesHandler) HandlePrivacy(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".go") || strings.Contains(r.URL.Path, "admin") {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err := layouts.Default(types.PageProps{
		PageId: types.PrivacyPage,
		Title:  "Privacy",
	}, pages.Privacy()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
