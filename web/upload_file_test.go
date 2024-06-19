package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/upload.gohtml"))

	tmpl.ExecuteTemplate(w, "upload.gohtml", nil)
}

func UploadCore(w http.ResponseWriter, r *http.Request) {
	//? Under the hood
	// r.ParseMultipartForm(32 << 20)
	name := r.FormValue("name")
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileDestination, err := os.Create("./uploads/" + name + "-" + header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "File successfully uploaded")
}

//go:embed uploads/test.jpg
var img []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Adit")
	part, _ := writer.CreateFormFile("file", "./uploads/test.jpg")
	part.Write(img)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "/", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()
	UploadCore(recorder, request)

	res, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, "File successfully uploaded", string(res))
}
