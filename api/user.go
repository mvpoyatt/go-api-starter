package api

import (
	"context"
	"errors"
	"net/mail"

	"github.com/bufbuild/connect-go"
	"github.com/mvpoyatt/go-api/database"
	userv1 "github.com/mvpoyatt/go-api/gen/proto/go/user/v1"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *UserServer) PutUser(
	ctx context.Context,
	req *connect.Request[userv1.PutUserRequest],
) (*connect.Response[userv1.PutUserResponse], error) {
	// Ensure email is correct format
	email := req.Msg.Email
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	// Hash user's password before saving
	password := req.Msg.Password
	if len(password) < 5 {
		err := errors.New("password must be at least 5 characters")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Create new user in database
	createResult := database.Db.Create(&database.User{
		Email:    email,
		Password: string(hash[:]),
	})
	if err := createResult.Error; err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	// Get user just created and return as protobuf user type
	var dbUser database.User
	searchResult := database.Db.Where("email = ?", email).First(&dbUser)
	if err := searchResult.Error; err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	// return jwt for future authorization by client
	jwt, err := NewToken(email)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&userv1.PutUserResponse{
		Token: jwt,
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (s *UserServer) LoginUser(
	ctx context.Context,
	req *connect.Request[userv1.LoginUserRequest],
) (*connect.Response[userv1.LoginUserResponse], error) {
	// Check that email is valid and exists in database
	email := req.Msg.Email
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	var dbUser database.User
	result := database.Db.Where("email = ?", email).First(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, result.Error)
	} else if err := result.Error; err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Check if given password matches hash
	password := req.Msg.Password
	dbPassword := dbUser.Password
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)); err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, err)
	} else {
		// return jwt for future authorization by client
		jwt, err := NewToken(email)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		res := connect.NewResponse(&userv1.LoginUserResponse{
			Token: jwt,
		})
		res.Header().Set("User-Version", "v1")
		return res, nil
	}
}

func (s *UserServer) GetUser(
	ctx context.Context,
	req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
	email := ctx.Value("email")
	// Search Database for email
	var dbUser database.User
	result := database.Db.Where("email = ?", email).First(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, result.Error)
	} else if err := result.Error; err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Return user info as protobuf object
	res := connect.NewResponse(&userv1.GetUserResponse{
		User: database.UserDatabaseToProtobuf(dbUser),
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (s *UserServer) DeleteUser(
	ctx context.Context,
	req *connect.Request[userv1.DeleteUserRequest],
) (*connect.Response[userv1.DeleteUserResponse], error) {
	email := ctx.Value("email")
	// Find user by email and delete by ID
	var dbUser database.User
	searchResult := database.Db.Where("email = ?", email).First(&dbUser)
	if errors.Is(searchResult.Error, gorm.ErrRecordNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, searchResult.Error)
	} else if err := searchResult.Error; err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	id := dbUser.ID
	deleteResult := database.Db.Delete(&database.User{}, id)
	if err := deleteResult.Error; err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Return deleted user
	res := connect.NewResponse(&userv1.DeleteUserResponse{
		User: database.UserDatabaseToProtobuf(dbUser),
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}
