package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"net/http"
	"robbo-assets/app/modules"
	"time"
)

func NewServer(lifecycle fx.Lifecycle, handlers modules.HandlerModule) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				router := SetupGinRouter(handlers)
				server := &http.Server{
					Addr: viper.GetString("server.address"),
					Handler: cors.New(
						cors.Options{
							AllowedOrigins:   []string{"http://0.0.0.0:8601"},
							AllowCredentials: true,
							AllowedMethods: []string{
								http.MethodGet,
								http.MethodPost,
								http.MethodPut,
								http.MethodDelete,
								http.MethodOptions,
							},
							AllowedHeaders: []string{"*"},
						},
					).Handler(router),
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second,
					MaxHeaderBytes: 1 << 20,
				}

				go func() {
					if err = server.ListenAndServe(); err != nil {
						log.Fatalf("Failed to listen and serve")
					}
				}()
				return
			},
			OnStop: func(context.Context) error {
				return nil
			},
		})
}

func SetupGinRouter(handlers modules.HandlerModule) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	handlers.AssetsHandler.InitAssetsRoutes(router)
	return router
}
