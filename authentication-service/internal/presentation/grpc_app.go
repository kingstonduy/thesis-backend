package grpc_server

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authentication_proto "github.com/kingstonduy/authentication-service/proto"
)

type app struct {
	secretKey                                                     []byte
	authentication_proto.UnimplementedAuthenticationServiceServer // Embed this for compatibility
}

func NewApp() authentication_proto.AuthenticationServiceServer {
	// todo move it to local variable
	var secretKey = []byte("Duong Khanh Duy")

	return &app{
		secretKey: secretKey,
	}
}

// CreateToken implements authentication_proto.AuthenticationServiceServer.
func (a *app) CreateToken(ctx context.Context, req *authentication_proto.CreateTokenRequest) (res *authentication_proto.CreateTokenResponse, err error) {
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    req.UseID,                        // Subject (user identifier)
		"iss":    "authentication-service",         // Issuer
		"aud":    "user",                           // Audience (user role)
		"exp":    time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":    time.Now().Unix(),                // Issued at
		"userId": req.UseID,                        // custom field
	})

	tokenString, err := claims.SignedString(a.secretKey)
	if err != nil {
		return nil, err
	}

	return &authentication_proto.CreateTokenResponse{
		Token: tokenString,
	}, nil
}

// VerifyToken implements authentication_proto.AuthenticationServiceServer.
func (a *app) VerifyToken(ctx context.Context, req *authentication_proto.VerifyTokenRequest) (res *authentication_proto.VerifyTokenResponse, err error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(req.GetToken(), func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not parse claims")
	}

	userId := claims["userId"].(string)

	// Return the verified token
	return &authentication_proto.VerifyTokenResponse{
		Valid:  true,
		UserID: userId,
	}, nil
}
