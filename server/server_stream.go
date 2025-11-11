package main

import (
	"log"
	"time"

	pb "github.com/iRittikSharma/go-grpc-demo/proto"
)

func (s *helloServer) SayHelloServerSideStreaming(req *pb.NameList , stream pb.GreetService_SayHelloServerSideStreamingServer ) error{

	log.Printf("got the req with the names %v" , req.Names)

	for _,name := range req.Names{
		res := &pb.HelloResponse{
			Message : "Hello" + name,
		}

		if err :=  stream.Send(res); err != nil{
			return  err
		}
		time.Sleep(2* time.Second)
	}
	return nil
}
