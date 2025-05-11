package thumbnail

import (
	"snaptrail/internal/structs"
)

type service interface {
	getThumbnailById(sessionId uint64, thumbnailId uint64, published bool) (thumbnail structs.Thumbnail, err error)
}

func newService() service {
	return svc{
		repo: newRepo(),
	}
}

type svc struct {
	repo repository
}

func (s svc) getThumbnailById(sessionId uint64, thumbnailId uint64, published bool) (thumbnail structs.Thumbnail, err error) {
	return s.repo.getThumbnailById(sessionId, thumbnailId, published)
}
