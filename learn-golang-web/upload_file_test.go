package learn_golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func ViewForm(writer http.ResponseWriter, request *http.Request) {

	templateGlobal.ExecuteTemplate(writer, "upload-file.html", nil)

}

func UploadFile(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(100 << 20) // 100MB
	file, fileHeader, err := request.FormFile("file")

	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)

	if err != nil {
		panic(err)
	}

	_, err2 := io.Copy(fileDestination, file)

	if err2 != nil {
		panic(err2)
	}

	name := request.PostFormValue("name")

	templateGlobal.ExecuteTemplate(writer, "upload-success.html", map[string]any{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

func TestUploadFileServer(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", ViewForm)
	mux.HandleFunc("/upload", UploadFile)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources/image.jpg
var imageFile []byte

func TestUploadFile(t *testing.T) {

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	err := writer.WriteField("name", "Hutama Trirahmanto")
	if err != nil {
		panic(err)
	}
	w, err2 := writer.CreateFormFile("file", "IMAGE.jpg")
	if err2 != nil {
		panic(err2)
	}

	w.Write(imageFile)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	UploadFile(recorder, request)

	b, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
