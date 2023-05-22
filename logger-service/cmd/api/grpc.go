package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

// this will be the server
type LogServer struct {
	logs.UnimplementedLogServiceServer
	//this is required for all the service which we write over grpc
	//to ensure backwards compatibility
	Models data.Models
}

// definition for function in remote grpc server
func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	//write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}
	//write this to mongo
	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err

	}
	//return response
	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}

// function for grpclistener
func (app *Config) gRpcListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Falied to listen to grpc: %w", err)
	}

	s := grpc.NewServer()
	//register the server
	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})
	log.Printf("grpc server started on the port %s", gRpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to listen for grpc: %w", err)
	}

}
