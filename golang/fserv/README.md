# Golang - FServ

Simple static server for uploading and serving files.

## Upload 

You can upload both `multipart/form-data` and `application/octet-stream`.

```bash

# Octet stream upload
curl -X POST localhost:9000 -H 'Content-Type: application/octet-stream' --data-binary '@example-file.txt'

# Multipart form data upload
curl -X POST localhost:9000 -H 'Content-Type: multipart/form-data' -F 'file=@example-file.txt'

```  