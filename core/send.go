package core

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sync"

	"github.com/gosuri/uiprogress"
	"github.com/kpango/glg"
	"google.golang.org/grpc"
)

//SendFiles sends requested files to recipient
func SendFiles(filePaths []string, recipient string, fragmentSize int) error {
	if !checkFiles(filePaths) {
		return errors.New("[DroneClient] Input files must exist")
	}

	conn, err := grpc.Dial(recipient, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("[DroneClient] %s", err.Error())
	}
	defer conn.Close()

	client := NewDroneClient(conn)
	wg := sync.WaitGroup{}

	uiprogress.Start()
	for _, filePath := range filePaths {
		wg.Add(1)
		go sendFile(filePath, fragmentSize, client, &wg)
	}

	wg.Wait()
	return nil
}

func sendFile(fp string, fragmentSize int, client DroneClient, wg *sync.WaitGroup) {
	fileName := filepath.Base(fp)
	fileSize := getFileSize(fp)
	fileContentBuffer := make([]byte, fragmentSize)

	stream, err := client.ReceiveFile(context.Background())
	if err != nil {
		glg.Errorf("[DroneClient] %s", err.Error())
	}

	file, err := os.Open(fp)
	if err != nil {
		glg.Errorf("[DroneClient] %s", err.Error())
	}

	var offset int64 = 0
	var fragmentID int = 0
	var totalFragments int32 = int32(math.Ceil(float64(fileSize) / float64(fragmentSize)))
	var bar = uiprogress.AddBar(int(totalFragments)).AppendCompleted().PrependElapsed()

	for offset < fileSize {
		_, err := file.Seek(offset, 0)
		if err != nil {
			glg.Errorf("[DroneClient] %s", err.Error())
		}

		_, err = file.Read(fileContentBuffer)
		if err != nil {
			glg.Errorf("[DroneClient] %s", err.Error())
		}

		err = stream.Send(&FileFragment{
			FileName:        fileName,
			FragmentId:      int32(fragmentID),
			FragmentContent: sanitizeBytes(fileContentBuffer),
			TotalFragments:  totalFragments,
		})
		if err != nil {
			glg.Errorf("[DroneClient] %s", err.Error())
		}

		offset += int64(fragmentSize)
		fragmentID++
		fileContentBuffer = make([]byte, fragmentSize)
		bar.Incr()
	}

	reply, err := stream.CloseAndRecv()
	if err != nil || reply.GetStatusCode() != 200 {
		glg.Errorf("[DroneClient] %s", err.Error())
	}

	wg.Done()
}

func checkFiles(files []string) bool {
	for _, file := range files {
		if !fileExists(file) || !isFile(file) {
			return false
		}
	}
	return true
}

func sanitizeBytes(bytes []byte) (filtered []byte) {
	for _, b := range bytes {
		if b != 0 {
			filtered = append(filtered, b)
		}
	}
	return
}
