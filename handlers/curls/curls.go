package curls

import (
	"encoding/json"
	"net/http"
	"strings"

	"codeberg.org/dankstuff/danktoolz/actions"
)

type curlPagesHandler struct {
	usecases *actions.Actions
}

func New(usecases *actions.Actions) *curlPagesHandler {
	return &curlPagesHandler{
		usecases: usecases,
	}
}

func (p *curlPagesHandler) HandleMyIP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".go") || strings.Contains(r.URL.Path, "admin") {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusTemporaryRedirect)
		return
	}

	ipInfo := actions.ExtractClientMetadata(r)
	w.Header().Set("Content-Type", "application/json")
	body, _ := json.MarshalIndent(ipInfo, "", "\t")
	body = append(body, '\n')
	_, _ = w.Write(body)
}
