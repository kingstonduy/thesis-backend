package authen

import (
	"context"
	"errors"
	"fmt"
	"log"

	authentication_proto "github.com/kingstonduy/user-service/internal/pkg/authen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ErrAuthenticationFailed = errors.New("the token is invalid")
)

const (
	AUTHENTICATION_SERVICE_PORT = "7006"
)

func VerifyToken(port string, token string) (useID string, err error) {
	conn, err := grpc.NewClient(fmt.Sprintf(":%s", AUTHENTICATION_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	c := authentication_proto.NewAuthenticationServiceClient(conn)
	verifyRequest := authentication_proto.VerifyTokenRequest{
		Token: token,
	}

	verifyResponse, err := c.VerifyToken(context.Background(), &verifyRequest)
	if err != nil {
		return "", err
	}

	if !verifyResponse.GetValid() {
		return "", ErrAuthenticationFailed
	}

	return verifyResponse.GetUserID(), nil
}
