package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"sales-go/helpers/bcrypt"
	"sales-go/helpers/gin-rest"
	logger "sales-go/helpers/logging"
)

func CORSMiddleware() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-location-lat, x-location-long, x-unique-id")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
        if r.Method == "OPTIONS" {
            w.Write([]byte("allowed"))
            return
        }
	})
}

func HeaderVerificationMiddleware(ctx *gin.Context) {
	hashKeyStr := ctx.GetHeader("key")

	err := bcrypt.ComparePassword(hashKeyStr, "phincon")
	if err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, fmt.Errorf("keyEnc is empty or not authorized"))
		ctx.Abort()
		return
	}
}

func LoggingMiddleware(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-location-lat, x-location-long, x-unique-id")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
        if r.Method == "OPTIONS" {
            w.Write([]byte("allowed"))
            return
        }

		logger.Infof(fmt.Sprintf("Started %s localhost:5000%s", r.Method, r.URL.Path), r)

		mux.ServeHTTP(w, r)

		// handle panic error middleware
		defer func() {
			fmt.Println("MIDDLEWARE PASS 1")
			err := recover()
			if err != nil {
				fmt.Println("MIDDLEWARE PASS 2")
				w.WriteHeader(http.StatusInternalServerError)
				logger.Errorf(err.(error), r)
			} else {
				logger.Infof(fmt.Sprintf("Completed %s localhost:5000%s in %v", r.Method, r.URL.Path, time.Since(time.Now())), r)
			}
		}()
	})
}

func Use(middleware ...http.Handler) []http.Handler {
	var handlers []http.Handler
	handlers = append(handlers, middleware...)
	return handlers
}
