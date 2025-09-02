package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kkl-v2/bootstrap"
	"kkl-v2/core/logs"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type UploadReponse struct {
	Data   UploadReponseData `json:"data"`
	Status bool              `json:"status"`
}
type UploadReponseData struct {
	Bucket   string `json:"bucket"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

type DeleteFileRequest struct {
	Bucket   string `json:"bucket"`
	FileName string `json:"file_name"`
}
type DeleteFileReponse struct {
	Data   string `json:"data"`
	Status bool   `json:"status"`
}

func SumPathName(fileName string) string {
	if fileName == "" {
		return ""
	}

	return bootstrap.GlobalEnv.Files.PathIp + fileName
}

func UploadFile(fileName string, file *multipart.FileHeader) (*UploadReponseData, error) {
	bucket := bootstrap.GlobalEnv.Files.Bucket
	key := bootstrap.GlobalEnv.Files.Key
	client := &http.Client{}

	// Open the file
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Create a new multipart writer
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	err = writer.WriteField("BUCKET", bucket)
	if err != nil {
		return nil, err
	}

	// Change the filename of the first file
	file.Filename = fmt.Sprintf("%v.png", fileName)
	// fmt.Println(file.Filename)
	// Create a new form file field
	part, err := writer.CreateFormFile("FILES", filepath.Base(file.Filename))
	if err != nil {
		return nil, err
	}

	// Copy the file content to the form file field
	_, err = io.Copy(part, f)
	if err != nil {
		return nil, err
	}

	// Close the multipart writer
	writer.Close()
	// Create a new POST request to the filestore.com/upload URL
	url := fmt.Sprintf("http://%v:%v/api/v1/file/upload", bootstrap.GlobalEnv.Files.Host, bootstrap.GlobalEnv.Files.Port)
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return nil, err
	}

	// Set the Content-Type header to multipart/form-data
	req.Header.Set("x-api-key", key)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle the response if needed
	// For example, you can check the response status code
	if resp.StatusCode != http.StatusOK {
		// Handle non-OK response
		return nil, fmt.Errorf("upload fail")
	}
	readAll, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	responseBody := UploadReponse{}
	err = json.Unmarshal(readAll, &responseBody)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if !responseBody.Status {
		logs.Error("FAIL TO UPLOAD FILE")
	}

	// responseBody.Data.FilePath = "http://10.150.1.88:3047" + responseBody.Data.FilePath
	return &responseBody.Data, nil
}

func DeleteFile(fileName string) (*DeleteFileReponse, error) {
	bucket := bootstrap.GlobalEnv.Files.Bucket
	key := bootstrap.GlobalEnv.Files.Key
	client := &http.Client{}

	body := DeleteFileRequest{
		Bucket:   bucket,
		FileName: fileName,
	}

	marshal, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	// Create a new DELETE request to the filestore.com/delete URL
	url := fmt.Sprintf("http://%v:%v/api/v1/file/delete", bootstrap.GlobalEnv.Files.Host, bootstrap.GlobalEnv.Files.Port)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}

	// Add the Content-Type header to multipart/form-data
	req.Header.Set("x-api-key", key)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle the response if needed
	// For example, you can check the response status code
	if resp.StatusCode != http.StatusOK {
		// Handle non-OK response
		return nil, fmt.Errorf("delete fail")
	}
	readAll, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	responseBody := DeleteFileReponse{}
	err = json.Unmarshal(readAll, &responseBody)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if !responseBody.Status {
		logs.Error("FAIL TO DELETE FILE")
	}

	return &responseBody, nil
}
