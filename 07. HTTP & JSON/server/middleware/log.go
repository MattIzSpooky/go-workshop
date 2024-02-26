package middleware

import (
	"fmt"
	"net/http"
	"workshop/http+json/server/log"
)

// A simple logging middleware.
func LogMiddleware(next http.Handler, log *log.ServerLogger) http.HandlerFunc {
	amountOfRequestThatCameIn := 0
	amountOfRequestThatWentOut := 0

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Before
		amountOfRequestThatCameIn++
		log.WriteInfo(fmt.Sprintf("Incoming requests: [%d]", amountOfRequestThatCameIn))

		next.ServeHTTP(w, r)

		// After
		amountOfRequestThatWentOut++
		log.WriteInfo(fmt.Sprintf("Successfully handled [%d] requests", amountOfRequestThatWentOut))
	})
}
