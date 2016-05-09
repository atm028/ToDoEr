package elasticservice

import proto "github.com/golang/protobuf/proto"
import (
    "fmt"
    server "github.com/micro/go-micro/server"
    context "golang.org/x/net/context"
)

type Request struct {
    Index string `protobuf:"bytes,1,opt,name=index"`
    Type string `protobuf:"bytes,2,opt,name=type"`
    ID string `protobuf:"bytes,3,opt,name=id"`
    Key string `protobuf:"bytes,4,opt,name=key"`
    Value string `protobuf:"bytes,5,opt,name=value"`
}
func (m *Request) ProtoMessage() {}
func (m *Request) Reset() { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }

type Response struct {
    Msg string `protobuf:"bytes,1,opt,name=msg"`
    StatusCode string `protobuf:"bytes,2,opt,name=statuscode"`
}
func (m *Response) ProtoMessage() {}
func (m *Response) Reset() { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }

func init() {
    proto.RegisterType((*Request)(nil), "todoer.srv.elastic.Request")
    proto.RegisterType((*Response)(nil), "todoer.srv.elastic.Response")
}

type ElasticServiceHandler interface {
    ElasticServiceHandler
}

type ElasticServiceHandleInterface struct {
    ElasticServiceHandler
}

func RegisterElasticServiceHandlers(s server.Server, hdlr ElasticServiceHandler) {
    s.Handle(s.NewHandler(&ElasticServiceHandleInterface{hdlr}))
}
