// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: user/v1/user.proto

package userv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/mvpoyatt/go-api/gen/proto/go/user/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "user.v1.UserService"
)

// UserServiceClient is a client for the user.v1.UserService service.
type UserServiceClient interface {
	PutUser(context.Context, *connect_go.Request[v1.PutUserRequest]) (*connect_go.Response[v1.PutUserResponse], error)
	LoginUser(context.Context, *connect_go.Request[v1.LoginUserRequest]) (*connect_go.Response[v1.LoginUserResponse], error)
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error)
	DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error)
}

// NewUserServiceClient constructs a client for the user.v1.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		putUser: connect_go.NewClient[v1.PutUserRequest, v1.PutUserResponse](
			httpClient,
			baseURL+"/user.v1.UserService/PutUser",
			opts...,
		),
		loginUser: connect_go.NewClient[v1.LoginUserRequest, v1.LoginUserResponse](
			httpClient,
			baseURL+"/user.v1.UserService/LoginUser",
			opts...,
		),
		getUser: connect_go.NewClient[v1.GetUserRequest, v1.GetUserResponse](
			httpClient,
			baseURL+"/user.v1.UserService/GetUser",
			opts...,
		),
		deleteUser: connect_go.NewClient[v1.DeleteUserRequest, v1.DeleteUserResponse](
			httpClient,
			baseURL+"/user.v1.UserService/DeleteUser",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	putUser    *connect_go.Client[v1.PutUserRequest, v1.PutUserResponse]
	loginUser  *connect_go.Client[v1.LoginUserRequest, v1.LoginUserResponse]
	getUser    *connect_go.Client[v1.GetUserRequest, v1.GetUserResponse]
	deleteUser *connect_go.Client[v1.DeleteUserRequest, v1.DeleteUserResponse]
}

// PutUser calls user.v1.UserService.PutUser.
func (c *userServiceClient) PutUser(ctx context.Context, req *connect_go.Request[v1.PutUserRequest]) (*connect_go.Response[v1.PutUserResponse], error) {
	return c.putUser.CallUnary(ctx, req)
}

// LoginUser calls user.v1.UserService.LoginUser.
func (c *userServiceClient) LoginUser(ctx context.Context, req *connect_go.Request[v1.LoginUserRequest]) (*connect_go.Response[v1.LoginUserResponse], error) {
	return c.loginUser.CallUnary(ctx, req)
}

// GetUser calls user.v1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// DeleteUser calls user.v1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the user.v1.UserService service.
type UserServiceHandler interface {
	PutUser(context.Context, *connect_go.Request[v1.PutUserRequest]) (*connect_go.Response[v1.PutUserResponse], error)
	LoginUser(context.Context, *connect_go.Request[v1.LoginUserRequest]) (*connect_go.Response[v1.LoginUserResponse], error)
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error)
	DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/user.v1.UserService/PutUser", connect_go.NewUnaryHandler(
		"/user.v1.UserService/PutUser",
		svc.PutUser,
		opts...,
	))
	mux.Handle("/user.v1.UserService/LoginUser", connect_go.NewUnaryHandler(
		"/user.v1.UserService/LoginUser",
		svc.LoginUser,
		opts...,
	))
	mux.Handle("/user.v1.UserService/GetUser", connect_go.NewUnaryHandler(
		"/user.v1.UserService/GetUser",
		svc.GetUser,
		opts...,
	))
	mux.Handle("/user.v1.UserService/DeleteUser", connect_go.NewUnaryHandler(
		"/user.v1.UserService/DeleteUser",
		svc.DeleteUser,
		opts...,
	))
	return "/user.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) PutUser(context.Context, *connect_go.Request[v1.PutUserRequest]) (*connect_go.Response[v1.PutUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.PutUser is not implemented"))
}

func (UnimplementedUserServiceHandler) LoginUser(context.Context, *connect_go.Request[v1.LoginUserRequest]) (*connect_go.Response[v1.LoginUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.LoginUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.GetUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.DeleteUser is not implemented"))
}