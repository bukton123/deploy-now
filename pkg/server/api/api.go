package api

import (
	"context"
	"github.com/labstack/echo"
	"io/ioutil"
)

type service struct {
	serve *echo.Echo
}

func New() *service {
	return &service{}
}

func (s *service) New() error {
	//logrus.Infof("%s Dashboard enabled", logging.AppNameType)
	//logrus.Infof("%s Server started on %s", logging.ServeType, config.Agent.Dashboard.Listen)
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetOutput(ioutil.Discard)
	//e.Use(static.ServeRoot("/", &assetFileSystem.AssetFS{
	//	Asset:     browser.Asset,
	//	AssetDir:  browser.AssetDir,
	//	AssetInfo: browser.AssetInfo,
	//	Prefix:    "website/build",
	//	Fallback:  "index.html",
	//}))

	s.serve = e
	return nil
}

func (s *service) Start() error {
	if err := s.serve.Start(":1809"); err != nil {
		//return fmt.Errorf("%s Shutting down the server: %v", logging.ServeType, err)
	}

	return nil
}

func (s *service) Close(ctx context.Context) error {
	if s.serve == nil {
		return nil
	}

	if err := s.serve.Shutdown(ctx); err != nil {
		//return fmt.Errorf("%s Shutdown: %v", logging.ServeType, err)
	}

	return nil
}
