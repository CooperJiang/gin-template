package app

import (
	"context"
	stdErrors "errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"email-manage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

const shutdownTimeout = 10 * time.Second

// RunHTTPServer 启动HTTP服务并支持优雅停机
func RunHTTPServer(engine *gin.Engine, addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !stdErrors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signalCh)

	select {
	case err := <-errCh:
		return err
	case sig := <-signalCh:
		logger.Info("接收到退出信号: %s", sig.String())
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		_ = srv.Close()
		return err
	}

	logger.Info("HTTP服务已优雅关闭")
	return nil
}
