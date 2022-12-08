package https

import (
	"github.com/shiqiyue/gof/shutdowns"
	"net/http"
)

func GracefulShutdown(s *http.Server) {
	shutdowns.GracefulShutdown(s)
}
