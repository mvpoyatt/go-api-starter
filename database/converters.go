package database

import (
	"strconv"

	pbUser "github.com/mvpoyatt/go-api/gen/proto/go/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserDatabaseToProtobuf(user User) *pbUser.User {
	return &pbUser.User{
		Id:        strconv.Itoa(int(user.ID)),
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
