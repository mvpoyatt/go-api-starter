package api

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/mvpoyatt/go-api/gen/proto/go/user/v1/userv1connect"
	"github.com/mvpoyatt/go-api/utils/logger"
)

type ServerConfig struct {
	HostName       string `mapstructure:"hostname"`
	Port           string `mapstructure:"port"`
	PasswordLength int    `mapstructure:"password_length"`
}

var PasswordLength int

type UserServer struct{}

func StartServer(configs ServerConfig) {
	PasswordLength = configs.PasswordLength
	mux := http.NewServeMux()
	user := &UserServer{}
	interceptors := connect.WithInterceptors(AuthInterceptor())
	path, handler := userv1connect.NewUserServiceHandler(user, interceptors)
	compress1KB := connect.WithCompressMinBytes(1024)

	mux.Handle(path, handler)
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(userv1connect.UserServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(userv1connect.UserServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(userv1connect.UserServiceName),
		compress1KB,
	))

	listenOn := configs.HostName + ":" + configs.Port
	srv := &http.Server{
		Addr: listenOn,
		Handler: h2c.NewHandler(
			newCORS().Handler(mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	stopSignals := make(chan os.Signal, 1)
	signal.Notify(stopSignals, os.Interrupt, syscall.SIGTERM)

	logger.Log.Info("Server setup succeeded")
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Panicf("HTTP listen and serve failed: %v", err)
		}
	}()

	<-stopSignals
	logger.Log.Info("Stopping server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Panicf("Failed to properly shutdown server: %v", err)
	}
}

func newCORS() *cors.Cors {
	// To let web developers play with the demo service from browsers, we need a
	// very permissive CORS setup.
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
			"Refresh-Token",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}
