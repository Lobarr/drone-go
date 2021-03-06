package core

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/gosuri/uiprogress"
	"github.com/kpango/glg"
	"golang.org/x/sync/errgroup"
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
	workerGroup := errgroup.Group{}

	uiprogress.Start()

	for _, filePath := range filePaths {
		workerGroup.Go(func() error {
			return sendFile(filePath, fragmentSize, client)
		})
	}

	if err = workerGroup.Wait(); err != nil {
		return glg.Get().Errorf(droneServerLogTemplate, err.Error())
	}
	return nil
}

func sendFile(fp string, fragmentSize int, client DroneClient) error {
	fileName := filepath.Base(fp)
	fileSize := getFileSize(fp)
	fileContentBuffer := make([]byte, fragmentSize)
	makeErr := func(err error) error {
		return fmt.Errorf("[%s] %s", fp, err.Error())
	}

	stream, err := client.ReceiveFile(context.Background())
	if err != nil {
		return err
	}

	file, err := os.Open(fp)
	if err != nil {
		return makeErr(err)
	}
	defer file.Close()

	var (
		offset         int64
		fragmentID     int
		totalFragments = int32(math.Ceil(float64(fileSize) / float64(fragmentSize)))
		transactionID  = uuid.New().String()
		bar            = uiprogress.AddBar(int(totalFragments)).AppendCompleted().PrependElapsed()
	)

	for offset < fileSize {
		err = getFileFragmentByID(file, fragmentID, fileContentBuffer)
		if err != nil {
			return makeErr(err)
		}

		err = stream.Send(&FileFragment{
			FileName:        fileName,
			FragmentID:      int32(fragmentID),
			FragmentContent: sanitizeBytes(fileContentBuffer),
			TotalFragments:  totalFragments,
			TransactionID:   transactionID,
		})
		if err != nil {
			return makeErr(err)
		}

		offset += int64(fragmentSize)
		fragmentID++
		fileContentBuffer = make([]byte, fragmentSize)
		bar.Incr()
	}

	reply, err := stream.CloseAndRecv()
	if err != nil || reply.GetStatusCode() != 200 {
		return makeErr(err)
	}

	return nil
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
