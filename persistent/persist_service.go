package persistent

import proto "github.com/golang/protobuf/proto"
import (
    "fmt"
    server "github.com/micro/go-micro/server"
    context "golang.org/x/net/context"
)

//Request type and staff
type  UserRequest struct {
    ID string `protobuf:"bytes,1,opt,name=id"`
    FirstName string `protobuf:"bytes,2,opt,name=firstname"`
    LastName string `protobuf:"bytes,3,opt,name=lastname"`
    EMail string `protobuf:"bytes,4,opt,name=email"`
}

func (m *UserRequest) ProtoMessage() { fmt.Println("proto ProtoMessage") }
func (m *UserRequest) Reset() { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }

//Response type and staff
type  UserResponse struct {
    ID string `protobuf:"bytes,1,opt,name=id"`
    FirstName string `protobuf:"bytes,2,opt,name=firstname"`
    LastName string `protobuf:"bytes,3,opt,name=lastname"`
    EMail string `protobuf:"bytes,4,opt,name=email"`
    StatusCode string `protobuf:"bytes,5,opt,name=statuscode"`
    Msg string `protobuf:"bytes,6,opt,name=msg"`
}

func (m *UserResponse) ProtoMessage() { fmt.Println("Proto message") }
func (m *UserResponse) Reset() { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }

func init() {
    proto.RegisterType((*UserRequest)(nil), "todoer.srv.persistent.UserRequest")
    proto.RegisterType((*UserResponse)(nil), "todoer.srv.persistent.UserResponse")
}

type PersistentServiceHandler interface {
    CreateUser(context.Context, *UserRequest, *UserResponse) error
    ReadUser  (context.Context, *UserRequest, *UserResponse) error
    UpdateUser(context.Context, *UserRequest, *UserResponse) error
    DeleteUser(context.Context, *UserRequest, *UserResponse) error
}

type PersistentServiceHandleInterface struct {
    PersistentServiceHandler
}

func RegisterPersistentServiceHandlers(s server.Server, hdlr PersistentServiceHandler) {
    s.Handle(s.NewHandler(&PersistentServiceHandleInterface{hdlr}))
}
