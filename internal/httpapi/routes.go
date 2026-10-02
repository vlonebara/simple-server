package httpapi

import "net/http"

func (a *App) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", a.health)
	mux.HandleFunc("GET /api/v1/hello", a.hello)
	return a.recoverer(a.requestLogger(mux))
}
