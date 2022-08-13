package main

import (
	"calvarado2004/microservices-go/log-service/data"
	"calvarado2004/microservices-go/log-service/logs"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer // required for service on grpc
	Models                             data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {

	input := req.GetLogEntry()

	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{
			Result: "logging failed with grpc",
		}
		return res, err
	}

	// return response
	res := &logs.LogResponse{Result: "success logging with gRPC"}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen with gRPC: %v", err)
	}

	s := grpc.NewServer()
	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve with gRPC: %v", err)
	}

	log.Printf("gRPC server listening on port %s", gRpcPort)

}
