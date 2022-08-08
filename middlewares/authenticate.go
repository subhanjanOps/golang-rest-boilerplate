package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang-rest/helpers"
	"golang-rest/pkg"
	statuscodes "golang-rest/statusCodes"
)

func (m Middleware) Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auths := r.Header.Values("authorization")
		if len(auths) <= 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(
				helpers.
					ParseResponse(
						statuscodes.
							ResponseBody{
							Status: *statuscodes.ErrorCodes["401"],
							Data: map[string]interface{}{
								"error": "'authorization' token is not provided. Please check.",
							},
						},
					),
			)
			return
		}
		accessToken := auths[0]
		manager := pkg.NewJWTManager(os.Getenv("SECRET_KEY"), 60*time.Minute)
		claims, err := manager.Verify(accessToken)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(
				helpers.
					ParseResponse(
						statuscodes.
							ResponseBody{
							Status: *statuscodes.ErrorCodes["401"],
							Data: map[string]interface{}{
								"error": fmt.Sprintf("%v. 'authorization' token is invalid. Please check.", err),
							},
						},
					),
			)
			log.Println(err)
			return
		}

		r.Header.Set("userId", fmt.Sprintf("%d", claims.ID))
		h.ServeHTTP(w, r)
	})
}
