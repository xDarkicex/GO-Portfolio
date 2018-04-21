package gzip

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// GzipResponseWriter is a Struct for manipulating io writer
type GzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (res GzipResponseWriter) Write(b []byte) (int, error) {
	if "" == res.Header().Get("Content-Type") {
		// If no content type, apply sniffing algorithm to un-gzipped body.
		res.Header().Add("Content-Type", http.DetectContentType(b))
	}
	return res.Writer.Write(b)
}

// Middleware force - bool, whether or not to force Gzip regardless of the sent headers.
func Middleware(fn httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, pm httprouter.Params) {
		res.Header().Add("Server", "Golang")
		if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			fn(res, req, pm)
			return
		}
		res.Header().Set("Vary", "Accept-Encoding")
		res.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(res)
		defer gz.Close()
		gzr := GzipResponseWriter{Writer: gz, ResponseWriter: res}
		fn(gzr, req, pm)
	}
}
