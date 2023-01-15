package models

type FileData struct {
	Name       string
	Data       string
	UserId     string
	UploadInfo struct {
		Path  string
		Error error
	}
}
