# ThunderCast

ThunderCast is a golang wrapper for UDPCast.
![ThunderCast logo ](./gosender/Public/images/Logo.png)
![ThunderCast Gopher](./gosender/Public/images/gopher.png)

## Usage

### Current usage

`curl -F "file=@[filename.type]" localhost:8080/GoUPDCast/v1/upload/`

### Ideal Usage

`curl -F "file=@test" -F "host=[IP]" -F "user=[username]" -F "pwd=[pwd]" localhost:8080/GoUPDCast/v1/upload/`