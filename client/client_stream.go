package main

import (
	"context"
	"log"
	"time"

	pb "github.com/iRittikSharma/go-grpc-demo/proto"
)


func callSayHelloClientStream(client pb.GreetServiceClient , names *pb.NameList){
	log.Printf("streaming has started from client")

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil{
		log.Fatalf("could not send names through streaming : %v")
	}

	for _,name := range names.Names {

		req := &pb.HelloRequest{
			Name : name,
		}

		if err := stream.Send(req) ; err != nil {
			log.Fatalf("Error while sending %v" , err)
		}

		log.Printf("send the request with name : %s" , name)

		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	log.Printf("Client streaming finished")

	if err != nil {
		log.Fatalf("Error while recieving %v" , err)
	}

	log.Printf("%v",  res.Messages)

}