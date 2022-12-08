package shutdowns

import (
	"context"
	"github.com/shiqiyue/gof/loggers"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ShutdownAble interface {
	Shutdown(ctx context.Context) error
}

func GracefulShutdown(s ShutdownAble) {
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGTERM)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			loggers.Error(context.Background(), "Server Shutdown", zap.Error(err))
			os.Exit(1)
		}
	}()
}
