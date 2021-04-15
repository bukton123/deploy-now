package server

import (
	"context"
	"dpn/pkg/server/api"
	"golang.org/x/sync/errgroup"
)

type Serve interface {
	Start()
	Close(ctx context.Context)
}

type service interface {
	New() error
	Start() error
	Close(ctx context.Context) error
}

type server struct {
	serve []service
}

func New() Serve {
	srv := &server{}

	// registry
	srv.registry(api.New())

	return srv
}

func (s *server) Start() {
	srv, _ := errgroup.WithContext(context.Background())

	for _, fn := range s.serve {
		if err := fn.New(); err != nil {
			//logrus.Fatal(err)
		}

		srv.Go(fn.Start)
	}

	if err := srv.Wait(); err != nil {
		//logrus.Fatal(err)
	}
}

func (s *server) Close(ctx context.Context) {
	for _, fn := range s.serve {
		if fn == nil {
			continue
		}

		if err := fn.Close(ctx); err != nil {
			//logrus.Error(err)
		}
	}
}

func (s *server) registry(svc service) {
	s.serve = append(s.serve, svc)
}
