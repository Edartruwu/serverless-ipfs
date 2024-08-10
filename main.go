package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ipfs/go-ipfs-api"
	"os"
)

type Request struct {
	ID       string `json:"id"`       // Unique ID provided by the caller
	FileName string `json:"fileName"` // Name of the file being uploaded
	FileData []byte `json:"fileData"` // File content in bytes
}

type Response struct {
	URI      string `json:"uri"`      // IPFS URI of the uploaded file
	FileName string `json:"filename"` // Name of the file
	ID       string `json:"id"`       // The ID that was provided in the request
}

// uploadToIPFS function to upload the file to IPFS and return the URI
func uploadToIPFS(fileData []byte) (string, error) {
	// Retrieve the IPFS node URL from the environment variable
	ipfsNodeURL := os.Getenv("IPFS_NODE_URL")
	if ipfsNodeURL == "" {
		//throw an error if env var is missing
		return "", fmt.Errorf("missing IPFS_NODE_URL environment variable")
	}

	sh := shell.NewShell(ipfsNodeURL)
	reader := bytes.NewReader(fileData)

	// Upload the file to our IPFS node
	cid, err := sh.Add(reader)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to IPFS: %v", err)
	}

	// Pin the CID to ensure it is not garbage collected lol
	err = sh.Pin(cid)
	if err != nil {
		return "", fmt.Errorf("failed to pin file on IPFS: %v", err)
	}

	// Return the IPFS URI
	ipfsURI := fmt.Sprintf("ipfs://%s", cid)
	return ipfsURI, nil
}

func handler(ctx context.Context, req Request) (Response, error) {
	ipfsURI, err := uploadToIPFS(req.FileData)
	if err != nil {
		return Response{}, fmt.Errorf("error handling request: %v", err)
	}

	resp := Response{
		URI:      ipfsURI,
		FileName: req.FileName,
		ID:       req.ID,
	}

	return resp, nil
}

func main() {
	lambda.Start(handler)
}
