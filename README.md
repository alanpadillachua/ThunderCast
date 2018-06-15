# GoUDPCast

GoUDPCast is a golang wrapper for UDPCast.

## Usage 
### Current usage 
`curl -F "file=@[filename.type]" localhost:8080/GoUPDCast/v1/upload/`
### Ideal Usage 
`curl -F "file=@test" -F "host=[IP]" -F "user=[username]" -F "pwd=[pwd]" localhost:8080/GoUPDCast/v1/upload/`