package main

import jsonpd "github.com/golang/protobuf/jsonpb"
import (
    "html/template"
    "fmt"
//    "io"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/metadata"
    persistent "github.com/atm028/GoToDoEr/persistent"
    "net/http"
    "golang.org/x/net/context"
)

type Page struct {
    Title string
    Body string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    req := client.NewRequest("todoer.srv.persistent", "PersistentServiceHandleInterface.ReadUser", &persistent.UserRequest{EMail:"john.doe@mail.com"})
    ctx := metadata.NewContext(context.Background(), map[string]string{
        "X-User-Id":"john",
    })
    rsp := &persistent.UserResponse{}
    if err := client.Call(ctx, req, rsp); err != nil {
        fmt.Println(err)
        return
    }
    mlr := jsonpd.Marshaler{}
    jsRsp, _ := mlr.MarshalToString(rsp)

    t, err := template.ParseFiles("static/index.html")
    if err != nil {
        panic(err)
    }
    p := &Page{}
    p.Title = "Title"
    p.Body = jsRsp

    t.Execute(w, p)
//    io.WriteString(w, jsRsp)
//    w.WriteHeader(200)
}

func main() {
    router := persistent.Handlers()
    router.HandleFunc("/", mainHandler)

    http.ListenAndServe(":8080", router)
}
