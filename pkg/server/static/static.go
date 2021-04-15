package static

import (
	"net/http"
	"strings"

	assetFileSystem "github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
)

type (
	FileSystem interface {
		http.FileSystem
		Exist(prefix, pathname string) bool
	}

	fileSystem struct {
		fs http.FileSystem
	}
)

func (fs *fileSystem) Open(name string) (http.File, error) {
	return fs.fs.Open(name)
}

func (fs *fileSystem) Exist(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(fs *assetFileSystem.AssetFS) *fileSystem {
	return &fileSystem{fs}
}

func ServeRoot(urlPrefix string, fs *assetFileSystem.AssetFS) echo.MiddlewareFunc {
	return Serve(urlPrefix, BinaryFileSystem(fs))
}

// Serve Static returns a middleware handler that serves static files in the given directory.
func Serve(urlPrefix string, fs FileSystem) echo.MiddlewareFunc {
	fileServer := http.FileServer(fs)
	if urlPrefix != "" {
		fileServer = http.StripPrefix(urlPrefix, fileServer)
	}
	return func(before echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := before(c)
			if err != nil {
				if c, ok := err.(*echo.HTTPError); !ok || c.Code != http.StatusNotFound {
					return err
				}
			}

			w, r := c.Response(), c.Request()
			if fs.Exist(urlPrefix, r.URL.Path) {
				fileServer.ServeHTTP(w, r)
				return nil
			}
			return err
		}
	}
}
