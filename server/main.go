package main

import (
	"log"
	"net"

	"github.com/LordRadamanthys/grpc_chat_golang/pb/chat_server"
	"google.golang.org/grpc"
)

func main() {

	Port := ":8888"

	listen, err := net.Listen("tcp", Port)

	if err != nil {
		log.Fatalf("Could not listen port ", Port)
	}

	log.Println("Listening ", Port)

	//grpc instance
	grpcServer := grpc.NewServer()

	//register chatService
	cs := chat_server.ChatServer{}
	chat_server.RegisterServicesServer(grpcServer, &cs)

	//grpc listen and server
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed start server port: ", Port)
	}

}
