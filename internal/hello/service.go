package hello

import "snaptrail/internal/structs"

type service interface {
	getHello() (hello structs.Hello, err error)
}

func newService() service {
	return svc{
		repo: newRepo(),
	}
}

type svc struct {
	repo repository
}

func (s svc) getHello() (hello structs.Hello, err error) {
	return s.repo.getHello()
}
