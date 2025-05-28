package server

import (
	"fmt"
	"net/http"
	adminsession "snaptrail/internal/admin-session"
	"snaptrail/internal/appuser"
	"snaptrail/internal/config"
	"snaptrail/internal/middleware"
	"snaptrail/internal/session"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	apiPrefix     = "/api/v1"
	sessionPath   = "/sessions"
	adminPath     = "/admin"
	loginPath     = "/login"
	thumbnailPath = "/thumbnails"
	userPath      = "/users"
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
				return strings.HasSuffix(origin, config.Get().DomainSuffix)
			},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	sessionHandler := session.New()
	apiRouter.GET(sessionPath, sessionHandler.Session)

	userHandler := appuser.New()
	apiRouter.POST(loginPath, userHandler.Login)

	protectedRouter := s.router.Group(apiPrefix + adminPath)
	protectedRouter.Use(middleware.JwtAuthMiddleware())

	adminSessionHandler := adminsession.New()
	protectedRouter.GET(sessionPath, adminSessionHandler.Session)
	protectedRouter.POST(sessionPath, adminSessionHandler.CreateOrUpdateSession)
	protectedRouter.POST(sessionPath+"/:id", adminSessionHandler.CreateOrUpdateSession)
	protectedRouter.DELETE(sessionPath+"/:id", adminSessionHandler.DeleteSession)

	protectedRouter.POST(userPath, userHandler.PasswordChange)

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
		_, _ = fmt.Fprintln(w, "ok")
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
