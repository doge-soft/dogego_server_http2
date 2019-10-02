package dogego_server_http2

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

// HTTP/2.0 服务器
func HTTP2ServerProtocol(router *gin.Engine) error {
	log.Printf("HTTP/2.0 Server started on %s", os.Getenv("ADDR_HTTP2"))
	err := http.ListenAndServeTLS(
		os.Getenv("ADDR_HTTP2"),
		os.Getenv("HTTP2_TLS_CERT"),
		os.Getenv("HTTP2_TLS_KEY"), router)

	if err != nil {
		return err
	}

	return nil
}
