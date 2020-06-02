package storage

import (
	"checkaem_server/cmd/tika"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
	"mime/multipart"
	"path"
	"strings"
	"sync"
	"time"
)

func Download(uuid string) (string, error) {
	filepath := path.Join(uuid, "file")

	return downloadInner(filepath)
}

func Upload(file multipart.File, uuid string) <-chan error {

	start := time.Now()

	errChan := make(chan error, 1)

	uploader := newUploader()

	filepath := path.Join(uuid, "file")

	mime, err := tika.GetType(file)

	if err != nil {
		errChan <- err
		return errChan
	}

	text, err := tika.ParseToStr(file)

	if err != nil {
		errChan <- err
		return errChan
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go uploadInner(uploader, filepath, file, mime, errChan, wg)

	textfile := strings.NewReader(text)

	textpath := path.Join(uuid, "text")
	wg.Add(1)
	go uploadInner(uploader, textpath, textfile, "text/plain", errChan, wg)

	go func() {
		wg.Wait()
		close(errChan)
		log.Println(time.Since(start))
	}()

	return errChan
}

func downloadInner(filepath string) (string, error) {
	svc := s3.New(s3Session)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath),
	})

	url, err := req.Presign(15 * time.Minute)

	if err != nil {
		return "", err
	}

	return url, err
}

func uploadInner(u *s3manager.Uploader, filepath string, file io.Reader, mime string, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := u.Upload(&s3manager.UploadInput{
		ContentType: aws.String(mime),
		Bucket:      aws.String(bucketName),
		Key:         aws.String(filepath),
		Body:        file,
	})

	if err != nil {
		errChan <- err
		return
	}

}
