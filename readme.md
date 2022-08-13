# Microservices in Go

This is the repository for microservices course in Go

- Built in Go version 1.18.1
- Uses the [chi browser](github.com/go-chi/chi/v5)


![Reservation](./images/microservices.png)


# MongoDB connection

MongoDB connection string

```
mongodb://mongo:mongo@localhost:27017/logs?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false
```

# Kubernetes objects

It requires Portworx

# Install gRPC on Go

Change your .bashrc file
```
export GOPATH="/opt/go"
PATH="$GOPATH/bin:/opt/protoc-21.5-linux-x86_64/bin:$PATH"
export PATH

source .bashrc

```

Install gRPC with the following commands:
```
 go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Download and install protoc 
```
wget https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-x86_64.zip
unzip https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-x86_64.zip

```



Compile the proto file
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto
```
