package filestore

import "bitbucket.org/junglee_games/getsetgo/models"

type FileStore interface {
	Save(filesData *models.FileData) (string, error)

	// it will upload data to upload queue and ack will receive on ack chan
	SaveAsync(filesData *models.FileData)

	// it will close the uploading queue Note : don't call SaveAsync method after calling this method
	StopUploading()

	// it will wait until all the workers finish its uploading
	WaitForFinishingUpload()

	// download file
	DownloadFile(filename string) ([]byte, error)

	// it will return ack chan
	GetAckChan() chan *models.FileData

	//it will temporarily save file and returns signed url
	TemporarySave(filesData *models.FileData, expriryMinutes int) (string, error)

	// it will return signed url for already uploaded file
	GetSignedURL(filepath string, expriryMinutes int) (string, error)
}

type Config interface {
	GetS3BucketName() string
	GetMaxRetries() int
	GetUploaderCount() int
}
