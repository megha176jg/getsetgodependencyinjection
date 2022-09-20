package onfido

type Onfido interface {
	CreateApplicant(firstName, lastName string, location bool) (*CreateApplicantResponse, error)
	UploadDocument(applicantId, fileType, filePath, side string) (*UploadDocumentResponse, error)
	CreateCheck(applicantId string, reportNames []string) (*CreateCheckResponse, error)
	RetriveReport(reportId string) (*ReportResponse, error)
	DownloadDocument(documentId string, destPath string) (err error)
}
