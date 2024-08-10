# Serverless IPFS Upload

This project provides an AWS Lambda function written in Go to upload files to IPFS. The Lambda function accepts a file and metadata, uploads the file to an IPFS node, pins it to ensure persistence, and returns a structured response with the IPFS URI, file name, and a provided ID.

## Prerequisites

- An IPFS node or a public IPFS service endpoint

## Invoke

- to invoke this lambda function you should send this

```json
{
  "id": "12345",
  "fileName": "example.jpg",
  "fileData": "SGVsbG8gd29ybGQ=" // Base64 encoded file content
}
```

have fun :3
