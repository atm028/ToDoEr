package main

import jsonpd "github.com/golang/protobuf/jsonpb"
import (
    persistent "github.com/atm028/GoToDoEr/persistent"
    "strconv"
    "github.com/micro/go-micro"
    "encoding/json"
    "time"
    "log"
    "golang.org/x/net/context"
)

type PersistentServiceHandleInterface struct {}

func(h *PersistentServiceHandleInterface) CreateUser(ctx context.Context, req *persistent.UserRequest, rsp *persistent.UserResponse) error {
    pl, err := persistent.NewPersistHandler()
    if(err != nil) {
        log.Fatal(err)
    }
    mlr := jsonpd.Marshaler{}
    jsReq, _ := mlr.MarshalToString(req)
    res, code := pl.PersistHandlerCreate(string(jsReq))
    rsp.StatusCode = strconv.Itoa(code)
    rsp.Msg = string(res)
    return nil
}

func(h *PersistentServiceHandleInterface) ReadUser(ctx context.Context, req *persistent.UserRequest, rsp *persistent.UserResponse) error {
    pl, err := persistent.NewPersistHandler()
    if(err != nil) {
        log.Fatal(err)
    }
    mlr := jsonpd.Marshaler{}
    jsReq, _ := mlr.MarshalToString(req)
    res, code := pl.PersistHandlerRead(string(jsReq))
    var dat persistent.User
    err = json.Unmarshal([]byte(res), &dat)

    rsp.StatusCode = strconv.Itoa(code)
    if(code == 200) {
        rsp.ID = dat.ID
        rsp.FirstName = dat.FirstName
        rsp.LastName = dat.LastName
        rsp.EMail = dat.EMail
    }
    if(code != 200) {
        rsp.Msg = res
    }
    return nil
}

func(h *PersistentServiceHandleInterface) UpdateUser(ctx context.Context, req *persistent.UserRequest, rsp *persistent.UserResponse) error {
    pl, err := persistent.NewPersistHandler()
    if(err != nil) {
        log.Fatal(err)
    }
    mlr := jsonpd.Marshaler{}
    jsReq, _ := mlr.MarshalToString(req)
    res, code := pl.PersistHandlerRead(string(jsReq))
    var dat persistent.User
    err = json.Unmarshal([]byte(res), &dat)

    res, code = pl.PersistHandlerUpdate(`{"id": "`+dat.ID+`"}`, string(jsReq))
    rsp.StatusCode = strconv.Itoa(code)
    rsp.Msg = string(res)
    return nil
}

func(h *PersistentServiceHandleInterface) DeleteUser(ctx context.Context, req *persistent.UserRequest, rsp *persistent.UserResponse) error {
    pl, err := persistent.NewPersistHandler()
    if(err != nil) {
        log.Fatal(err)
    }
    mlr := jsonpd.Marshaler{}
    jsReq, _ := mlr.MarshalToString(req)
    res, code := pl.PersistHandlerDelete(string(jsReq))
    rsp.StatusCode = strconv.Itoa(code)
    rsp.Msg = string(res)
    return nil
}

func main() {
    pService := micro.NewService(
        micro.Name("todoer.srv.persistent"),
        micro.RegisterTTL(time.Second*30),
        micro.RegisterInterval(time.Second*10),
    )
    pService.Init()
    persistent.RegisterPersistentServiceHandlers(pService.Server(), new(PersistentServiceHandleInterface))
    if err := pService.Run(); err != nil {
        log.Fatal(err)
    }
}
