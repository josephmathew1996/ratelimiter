package echo

import (
	"ratelimiter/internal/config"
	"ratelimiter/internal/logger"
	"ratelimiter/internal/middlewares"
	"ratelimiter/internal/ratelimiter"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type EchoServer struct {
	Config      config.Config
	Logger      *zap.SugaredLogger
	RateLimiter ratelimiter.Limiter
	Echo        *echo.Echo
}

func NewEchoServer() (*EchoServer, error) {
	// Initialize configuration
	cfg := config.InitializeConfig()

	// Initialize logger
	log, err := logger.InitializeLogger(cfg.App.LogLevel)
	if err != nil {
		return nil, err
	}

	// Rate limiter for per client
	rateLimiter := ratelimiter.NewFixedWindowClientsRateLimiter(
		cfg.RateLimiter.RequestsPerTimeFrame,
		cfg.RateLimiter.TimeFrame,
	)

	// Rate limiter for whole server
	// rateLimiter := ratelimiter.NewFixedWindowRateLimiter(
	// 	cfg.RateLimiter.RequestsPerTimeFrame,
	// 	cfg.RateLimiter.TimeFrame,
	// )

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.RateLimiterMiddleware(rateLimiter, cfg.RateLimiter.Enabled, log))
	// e.Use(middlewares.SampleEchoMiddleware)
	// e.Use(middlewares.SampleEchoMiddlewareWithArgs("arg1"))

	return &EchoServer{
		Config:      cfg,
		Logger:      log,
		RateLimiter: rateLimiter,
		Echo:        e,
	}, nil
}

func (s *EchoServer) Start() {
	port := s.Config.App.Port
	portStr := strconv.Itoa(port)
	s.Logger.Info("Starting server on port " + portStr)
	if err := s.Echo.Start(":" + portStr); err != nil {
		s.Logger.Fatal("Failed to start echo server", zap.Error(err))
	}
}
