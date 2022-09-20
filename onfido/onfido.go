package onfido

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"bitbucket.org/junglee_games/getsetgo/httpclient"
	"bitbucket.org/junglee_games/getsetgo/newrelic"

	"github.com/pkg/errors"
)

type OnfidoSDK struct {
	Endpoint   string
	AuthToken  string
	nr         newrelic.Agent
	httpClient httpclient.HTTPClient
}

func New(endpoint, authToken string, client httpclient.HTTPClient, nr newrelic.Agent) *OnfidoSDK {
	return &OnfidoSDK{Endpoint: endpoint, AuthToken: authToken, httpClient: client, nr: nr}
}

var (
	local_ip_address = ""
)

func getLocalIP() string {
	if local_ip_address != "" {
		return local_ip_address
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				local_ip_address = ipnet.IP.String()
				return local_ip_address
			}
		}
	}
	return ""
}
func (osdk *OnfidoSDK) addHeaders(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", osdk.AuthToken))
	req.Header.Add("Content-Type", "application/json")

}
func (osdk *OnfidoSDK) CreateApplicant(firstName, lastName string, location bool) (*CreateApplicantResponse, error) {
	tr := osdk.nr.StartTransaction(ONFIDO_CREATE_APPLICANT_CALL)
	defer tr.End()
	url := osdk.Endpoint + "/v3.4/applicants"
	method := "POST"
	bytes, _ := json.Marshal(CreateApplicantRequest{FirstName: firstName, LastName: lastName, Location: Location{CountryOfResidence: "IND", IPAddress: getLocalIP()}})
	payload := strings.NewReader(string(bytes))

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, errors.Wrap(ErrRequestCreation, err.Error())
	}
	osdk.addHeaders(req)
	res, err := osdk.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingOnfido, err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(ErrReadingResponseBody, err.Error())
	}
	var result CreateApplicantResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshalingResponse, err.Error())
	}
	return &result, nil
}

func (osdk *OnfidoSDK) UploadDocument(applicantId, fileType, filePath, side string) (*UploadDocumentResponse, error) {
	tr := osdk.nr.StartTransaction(ONFIDO_UPLOAD_DOCUMENT_CALL)
	defer tr.End()

	url := osdk.Endpoint + "/v3.4/documents"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("type", fileType)
	file, errFile2 := os.Open(filePath)
	if errFile2 != nil {
		return nil, errors.Wrap(ErrReadingFile, errFile2.Error())
	}
	defer file.Close()
	part2, errFile2 := writer.CreateFormFile("file", filepath.Base(filePath))
	if errFile2 != nil {
		return nil, errors.Wrap(ErrReadingFile, errFile2.Error())
	}
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		return nil, errors.Wrap(ErrReadingFile, errFile2.Error())
	}
	_ = writer.WriteField("applicant_id", applicantId)
	err := writer.Close()
	if err != nil {

		return nil, errors.Wrap(ErrWritingFormField, err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, errors.Wrap(ErrRequestCreation, err.Error())
	}
	osdk.addHeaders(req)
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingOnfido, err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(ErrReadingResponseBody, err.Error())
	}
	var result UploadDocumentResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshalingResponse, err.Error())
	}
	return &result, nil
}

func (osdk *OnfidoSDK) CreateCheck(applicantId string, reportNames []string) (*CreateCheckResponse, error) {
	tr := osdk.nr.StartTransaction(ONFIDO_CREATE_CHECK_CALL)
	defer tr.End()

	url := osdk.Endpoint + "/v3.4/checks"
	method := "POST"
	bytes, _ := json.Marshal(&CreateCheckRequest{ApplicantId: applicantId, ReportNames: reportNames})
	payload := strings.NewReader(string(bytes))

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, errors.Wrap(ErrRequestCreation, err.Error())
	}
	osdk.addHeaders(req)
	res, err := osdk.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingOnfido, err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(ErrReadingResponseBody, err.Error())
	}
	var result CreateCheckResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshalingResponse, err.Error())
	}
	return &result, nil
}

func (osdk *OnfidoSDK) RetriveReport(reportId string) (*ReportResponse, error) {
	tr := osdk.nr.StartTransaction(ONFIDO_RETRIVE_REPORT_CALL)
	defer tr.End()

	url := fmt.Sprintf("%s/v3.4/reports/%s", osdk.Endpoint, reportId)
	method := "GET"

	payload := strings.NewReader("")

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, errors.Wrap(ErrRequestCreation, err.Error())
	}

	osdk.addHeaders(req)

	res, err := osdk.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrCallingOnfido, err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(ErrReadingResponseBody, err.Error())
	}
	var result ReportResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Wrap(ErrUnmarshalingResponse, err.Error())
	}
	return &result, nil
}

func (osdk *OnfidoSDK) DownloadDocument(documentId string, destPath string) (err error) {
	tr := osdk.nr.StartTransaction(ONFIDO_DOWNLOAD_DOCUMENT_CALL)
	defer tr.End()

	url := fmt.Sprintf("%s/v3.4/documents/%s/download", osdk.Endpoint, documentId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return errors.Wrap(ErrRequestCreation, err.Error())
	}
	osdk.addHeaders(req)
	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(ErrCallingOnfido, err.Error())
	}
	defer res.Body.Close()

	err = ensureBaseDir(destPath)
	if err != nil {
		return errors.Wrap(ErrFileOrDirectoryNotFound, err.Error())
	}
	file, err := os.Create(destPath)
	if err != nil {
		return errors.Wrapf(ErrFileOrDirectoryNotFound, err.Error())
	}
	defer file.Close()
	// Writer the body to file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return errors.Wrap(ErrReadingResponseBody, err.Error())
	}
	return nil
}

func ensureBaseDir(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}
