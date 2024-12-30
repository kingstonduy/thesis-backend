package main

import (
	"context"
	"fmt"
	"log"

	authentication_proto "github.com/kingstonduy/authentication-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":7006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	c := authentication_proto.NewAuthenticationServiceClient(conn)

	request := authentication_proto.CreateTokenRequest{
		UseID: "duydk3",
	}
	response, err := c.CreateToken(context.Background(), &request)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Token created: %s\n", response.Token)

	verifyRequest := authentication_proto.VerifyTokenRequest{
		Token: response.Token,
	}
	verifyResponse, err := c.VerifyToken(context.Background(), &verifyRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Token is valid: %t, userID= %s\n", verifyResponse.Valid, verifyResponse.UserID)
}
