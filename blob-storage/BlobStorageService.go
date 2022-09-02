package blob_storage

import "go.uber.org/zap"

type BlobStorageService interface {
	PutWithCommand(request BlobStorageRequest) error
	Get(request BlobStorageRequest) error
}

type BlobStorageServiceImpl struct {
	logger *zap.SugaredLogger
}

func NewBlobStorageServiceImpl(logger *zap.SugaredLogger) *BlobStorageServiceImpl {
	impl := &BlobStorageServiceImpl{
		logger: logger,
	}
	return impl
}

func (impl *BlobStorageServiceImpl) PutWithCommand(request BlobStorageRequest) error {
	return nil
}

func (impl *BlobStorageServiceImpl) Get(request BlobStorageRequest) error {
	return nil
}
