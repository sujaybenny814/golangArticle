package validation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	model "github.com/micro1/models"
)

var validate *validator.Validate

func AddArticle(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		var article model.Article
		json.Unmarshal(body, &article)

		err := validate.Struct(&article)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Failed to create an article")
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
