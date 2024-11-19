package routes

import (
	"cart-service/config"
	"cart-service/util/middleware"
	"context"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	cart "cart-service/handlers/cart"

	"github.com/spf13/viper"
)

type Routes struct {
	Router *http.ServeMux
	Cart   *cart.Handler
}

func URLRewriter(baseURLPath string, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, baseURLPath)

		next.ServeHTTP(w, r)
	}
}

func (r *Routes) setupRouter() {
	r.Router = http.NewServeMux()
	r.cartRoutes()
}

func (r *Routes) cartRoutes() {
	r.Router.HandleFunc("DELETE /cart/{user_id}/{product_id}", middleware.ApplyMiddleware(r.Cart.DeleteCart, middleware.EnabledCors, middleware.LoggerMiddleware()))
}

func (r *Routes) Run(port string, wg *sync.WaitGroup) {
	defer wg.Done()

	r.setupRouter()

	log.Printf("[Running-Success] clients on localhost on port :%s", port)
	srv := &http.Server{
		Handler:      r.Router,
		Addr:         "localhost:" + port,
		WriteTimeout: config.WriteTimeout() * time.Second,
		ReadTimeout:  config.ReadTimeout() * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panicf("Failed to start HTTP server: %v", err)
	}
}

func (r *Routes) Shutdown(ctx context.Context) error {
	srv := &http.Server{
		Handler: r.Router,
		Addr:    "localhost:" + viper.GetString("HTTP_PORT"),
	}

	return srv.Shutdown(ctx)
}

func (r *Routes) ShutdownHTTP() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
		return
	}

	log.Println("HTTP server stopped")
}
