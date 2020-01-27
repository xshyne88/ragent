package main

import (
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/hashicorp/yamux"
	pb "github.com/xshyne88/ragent/proto"

	"google.golang.org/grpc"
)

type server struct{}

func (s server) SendCommands(srv pb.CommandIssuer_SendCommandsServer) error {
	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()

		if err == io.EOF {
			log.Println("io.EOF Reached...Exiting now")
			return nil
		}

		if err != nil {
			log.Printf("agent received an error from server: %v", err)
			continue
		}

		log.Printf("Got Request: {\n ReqId: %s, \n Req.CommandType: %s \n}", req.Id, req.CommandType)

		// spin up go routine to handle work, block until completed
		// then send response

		resp := pb.CommandResponse{Id: req.Id, Acknowledged: true}

		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
	}
}

func main() {
	log.Println("Agent started.")
	log.Println("---------------")

	log.Println("Dialing Server Home.....")
	conn, err := net.DialTimeout("tcp", ":3003", time.Second*5)
	if err != nil {
		log.Fatalf("error dialing", err)
	}

	serverConnection, err := yamux.Server(conn, yamux.DefaultConfig())

	grpcServer := grpc.NewServer()
	pb.RegisterCommandIssuerServer(grpcServer, server{})
	log.Println("---------------")
	log.Println("Connected To Server Home, Setting up Service...")

	if err := grpcServer.Serve(serverConnection); err != nil {
		log.Fatalf("failed to serve grpc server: ", err)
	}
}

func checkError(err error, resp string) {
	if err != nil {
		log.Fatal(resp, err)
	}
}

func getRandomCommand() string {
	num := rand.Intn(3)

	switch num {
	case 1:
		return "Restart"
	case 2:
		return "Shut Down"
	default:
		return "Start Up"
	}
}
