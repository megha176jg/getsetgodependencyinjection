package filestore

import (
	"log"

	"bitbucket.org/junglee_games/getsetgo/clients/aws"
)

type Factory struct {
	config Config
	logger *log.Logger
}

func NewFileStoreFactory(config Config, logger *log.Logger) *Factory {
	return &Factory{config: config, logger: logger}
}

func (fsf *Factory) GetFileStore(name string) (FileStore, error) {
	switch name {
	case AMAZON_S3:
		return aws.NewS3Store(fsf.config.GetS3BucketName(), fsf.config.GetUploaderCount(), fsf.config.GetMaxRetries(), fsf.logger)
	}
	return nil, ErrFileStoreNotFound
}
