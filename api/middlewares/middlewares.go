package middlewares

import(
	"errors"
	"net/http"

	"../fullstack/api/auth"
	"github.com/victorsteven/fullstack/api/responses"
)


func SetMiddlewareJSON(next, http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponsWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
