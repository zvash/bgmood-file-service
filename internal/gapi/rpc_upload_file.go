package gapi

import (
	"bytes"
	"context"
	"fmt"
	"github.com/zvash/bgmood-file-service/internal/filepb"
	"github.com/zvash/bgmood-file-service/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
)

func (server *Server) UploadFile(stream filepb.File_UploadFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive file info"))
	}
	fileName := util.RandomFileName(req.GetInfo().GetExtension())
	filePath := ""
	ft := req.GetInfo().GetFileType()
	switch ft {
	case filepb.FileInfo_AVATAR_IMAGE:
		filePath = fmt.Sprintf("files/avatars/%s", fileName)
	case filepb.FileInfo_ENCRYPTED_FILE:
		filePath = fmt.Sprintf("files/encrypted/%s", fileName)
	case filepb.FileInfo_BACKGROUND_IMAGE:
		filePath = fmt.Sprintf("files/backgrounds/%s", fileName)
	}
	log.Printf("receive an upload file request with type %s", req.GetInfo().GetFileType())

	fileData := bytes.Buffer{}
	fileSize := 0

	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetChunkData()
		size := len(chunk)
		log.Printf("received a chunk with size: %d", size)
		fileSize += size
		log.Printf("total received size: %d", size)

		_, err = fileData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return status.Errorf(codes.Unknown, "couldn't create the file: %s err: %v", filePath, err)
	}
	_, err = fileData.WriteTo(file)
	if err != nil {
		return status.Errorf(codes.Unknown, "couldn't write data to file: %s err: %v", filePath, err)
	}
	resp := &filepb.UploadFileResponse{
		Name: fileName,
		Path: filePath,
		Url:  "",
	}
	err = stream.SendAndClose(resp)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
