package core

import (
	"bytes"
	"context"
	"errors"
	"math"
	"os"
	"path/filepath"
	"sync"

	"github.com/gosuri/uiprogress"
	"github.com/kpango/glg"
	"google.golang.org/grpc"
)

var droneClientLogTemplate = "[DroneClient] %s"

//SendFiles sends requested files to recipient
func SendFiles(filePaths []string, recipient string, fragmentSize int) error {
	if !checkFiles(filePaths) {
		return errors.New("Input files must exist")
	}

	conn, err := grpc.Dial(recipient, grpc.WithInsecure())
	if err != nil {
		return err
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
		glg.Errorf(droneServerLogTemplate, err.Error())
	}

	file, err := os.Open(fp)
	if err != nil {
		glg.Errorf(droneServerLogTemplate, err.Error())
	}

	var offset int64
	var fragmentID int
	var totalFragments = int32(math.Ceil(float64(fileSize) / float64(fragmentSize)))
	var bar = uiprogress.AddBar(int(totalFragments)).AppendCompleted().PrependElapsed()

	for offset < fileSize {
		err = getFileFragmentByID(file, fragmentID, fileContentBuffer)
		if err != nil {
			glg.Errorf(droneServerLogTemplate, err.Error())
		}

		err = stream.Send(&FileFragment{
			FileName:        fileName,
			FragmentId:      int32(fragmentID),
			FragmentContent: sanitizeBytes(fileContentBuffer),
			TotalFragments:  totalFragments,
		})
		if err != nil {
			glg.Errorf(droneServerLogTemplate, err.Error())
		}

		offset += int64(fragmentSize)
		fragmentID++
		fileContentBuffer = make([]byte, fragmentSize)
		bar.Incr()
	}

	reply, err := stream.CloseAndRecv()
	if err != nil || reply.GetStatusCode() != 200 {
		glg.Errorf(droneServerLogTemplate, err.Error())
	}

	wg.Done()
}

func getFileFragmentByID(file *os.File, fragmentID int, fileContent []byte) error {
	fragmentSize := len(fileContent)
	offset := int64(fragmentID * fragmentSize)

	if _, err := file.Seek(offset, 0); err != nil {
		return err
	}

	if _, err := file.Read(fileContent); err != nil {
		return err
	}

	return nil
}

func checkFiles(files []string) bool {
	for _, file := range files {
		if !fileExists(file) || !isFile(file) {
			return false
		}
	}
	return true
}

func sanitizeBytes(_bytes []byte) []byte {
	return bytes.TrimRightFunc(_bytes, func(r rune) bool {
		return r == rune(0)
	})
}
