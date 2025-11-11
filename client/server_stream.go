package main

import (
	"context"
	"io"
	"log"

	pb "github.com/iRittikSharma/go-grpc-demo/proto"
)


func callSayHelloServerStream(client pb.GreetServiceClient , names *pb.NameList){
	log.Printf("streaming has started")

	stream, err := client.SayHelloServerSideStreaming(context.Background() , names)

	if err != nil {
		log.Fatalf("could not send names %v" , err)
	}

	for{
		message, err := stream.Recv()

		if err  == io.EOF{
			break
		}

		if err != nil{
			log.Fatalf("error while streaming %v" , err)
		}

		log.Println(message)
	}
	log.Printf("streaming end")
}