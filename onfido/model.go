package onfido

type Location struct {
	IPAddress          string `json:"ip_address"`
	CountryOfResidence string `json:"country_of_residence"`
}

type CreateApplicantRequest struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Location  Location `json:"location"`
}

type CreateApplicantResponse struct {
	ID    string      `json:"id"`
	Error interface{} `json:"error"`
}

type CreateCheckRequest struct {
	ApplicantId string   `json:"applicant_id"`
	ReportNames []string `json:"report_names"`
}

type CreateCheckResponse struct {
	ID        string      `json:"id"`
	ReportIds []string    `json:"report_ids"`
	Status    string      `json:"status"`
	Error     interface{} `json:"error"`
}

type UploadDocumentResponse struct {
	ID           string      `json:"id"`
	FileName     string      `json:"file_name"`
	FileType     string      `json:"file_type"`
	Type         string      `json:"type"`
	Size         string      `json:"size"`
	Side         string      `json:"side"`
	DownloadHref string      `json:"download_href"`
	Error        interface{} `json:"error"`
}

type ReportResponse struct {
	CheckId    string `json:"check_id"`
	CreatedAt  string `json:"created_at"`
	ReportId   string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		DOB             string `json:"date_of_birth"`
		DateOfExpiry    string `json:"date_of_expiry"`
		DocumentNumbers []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"document_numbers"`
		DocType        string `json:"document_type"`
		FirstName      string `json:"first_name"`
		LastName       string `json:"last_name"`
		Gender         string `json:"gender"`
		IssuingCountry string `json:"issuing_country"`
		Nationality    string `json:"nationality"`
	} `json:"properties"`
	Result string      `json:"result"`
	Error  interface{} `json:"error"`
}
