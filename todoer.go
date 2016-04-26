package todoer

import (
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "strings"
)

type User struct {
    ID, FirstName, LastName, EMail string
}

func userGetHandler(w http.ResponseWriter, r *http.Request) {
    persistHandler, err := NewPersistHandler()
    if(err != nil) {
        panic(err)
    }

    req := `{`
    for key := range r.URL.Query() {
        req = req + `"`+key+`":"`+r.URL.Query().Get(key)+`",`
    }
    req = strings.TrimRight(req, ",")
    req = req + `}`
    res, code := persistHandler.persistHandlerRead(req)
    w.WriteHeader(code)
    w.Write([]byte(res))
}

func userPutHandler(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)

    persistHandler, err := NewPersistHandler()
    if(err != nil) {
        panic(err)
    }
    res, code := persistHandler.persistHandlerCreate(string(body))
    w.WriteHeader(code)
    w.Write([]byte(res))
}

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
    persistHandler, err := NewPersistHandler()
    if(err != nil) {
        panic(err)
    }
    res, code := persistHandler.persistHandlerDelete(`{"id": "`+r.URL.Query().Get("id")+`"}`)
    w.WriteHeader(code)
    w.Write([]byte(res))
}

func userUpdateHandler(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)
    persistHandler, err := NewPersistHandler()
    if(err != nil) {
        panic(err)
    }
    res, code := persistHandler.persistHandlerUpdate(`{"id": "`+r.URL.Query().Get("id")+`"}`, string(body))
    w.WriteHeader(code)
    w.Write([]byte(res))
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
}

func Handlers() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/persist/users", userGetHandler).Methods("GET")
    r.HandleFunc("/persist/users", userPutHandler).Methods("PUT")
    r.HandleFunc("/persist/users", userUpdateHandler).Methods("POST")
    r.HandleFunc("/persist/users", userDeleteHandler).Methods("DELETE")

    return r
}

func main() {
    router := Handlers()
    http.ListenAndServe(":8080", router)
}
