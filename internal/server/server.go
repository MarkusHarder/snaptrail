package server

import (
	"fmt"
	"net/http"
	"snaptrail/internal/config"
	"snaptrail/internal/hello"
	"snaptrail/internal/middleware"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	apiPrefix = "/api/v1"
	helloPath = "/hello"
	adminPath = "/admin"
)

type Server struct {
	uiDir  string
	router *gin.Engine
}

func New(uiDir string) *Server {
	if !config.Get().Dev {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	return &Server{
		uiDir:  uiDir,
		router: r,
	}
}

func (s *Server) setupRoutes() {
	apiRouter := s.router.Group(apiPrefix)
	if !config.Get().Dev {
		apiRouter.Use(cors.New(cors.Config{
			AllowOriginFunc: func(origin string) bool {
				//TODO: get suffix from domain
				return strings.HasSuffix(origin, ".snaptrail.markusharder.com")
			},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}
	helloHandler := hello.New()
	apiRouter.GET(helloPath, helloHandler.Hello)
	s.router.GET("/ui/*filepath", middleware.StaticUi("/ui", s.uiDir))
}

func (s *Server) Start() {
	go func() {
		s.setupRoutes()
		err := s.router.Run(fmt.Sprintf(":%d", config.Get().Port))
		if err != nil {
			log.Fatal().Stack().Err(err).Msg("server exited with error")
		}
		routes := s.router.Routes()
		log.Info().Msg("registered routes")
		for _, r := range routes {
			log.Info().Msgf("%s %s", r.Method, r.Path)
		}

		log.Info().Msg("server stopped")
	}()

	httpMux := http.NewServeMux()
	httpMux.HandleFunc(adminPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	server := &http.Server{
		Addr:    config.Get().AdminPort,
		Handler: httpMux,
	}

	log.Info().Msgf("Liveness probe server running on %s", config.Get().AdminPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("Liveness server failed: %v", err)
	}
}
