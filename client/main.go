package main

import (
	"context"
	"fmt"
	"log"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	userv1 "github.com/mvpoyatt/go-api/gen/proto/go/user/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to PetStoreService on %s: %w", connectTo, err)
	}
	defer conn.Close()
	log.Println("Connected to", connectTo)

	userService := userv1.NewUserServiceClient(conn)

	// if _, err := userService.PutUser(context.Background(), &userv1.PutUserRequest{
	// 	Email:    "testuser@fake.email",
	// 	Password: "hackme",
	// }); err != nil {
	// 	return fmt.Errorf("Failed to put new user: %w", err)
	// }

	var userFromDb *userv1.GetUserResponse
	userEmail := "testuser@fake.email"
	if dbResponse, err := userService.GetUser(context.Background(), &userv1.GetUserRequest{
		Email: userEmail,
	}); err != nil {
		log.Println(err)
		return fmt.Errorf("Failed to get user with email %s", userEmail)
	} else {
		userFromDb = dbResponse
	}

	if userFromDb != nil {
		log.Println("User retrieved by email: ", userFromDb)
	}

	return nil
}
