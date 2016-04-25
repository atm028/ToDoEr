package todoer

import (
    "io/ioutil"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

type User struct {
    FirstName, LastName, EMail string
}

func persistGetHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("TMOROZOV: "+r.URL.Path)
    fmt.Println("TMOROZOV: "+r.Method)
    fmt.Println(r.URL.Query()["firstname"])
    w.WriteHeader(http.StatusOK)
}

func persistPutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("TMOROZOV: "+r.URL.Path)
    fmt.Println("TMOROZOV: "+r.Method)
    body, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(body))
    w.WriteHeader(http.StatusCreated)
}

func persistDeleteHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
}

func Handlers() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/persist/{user}", persistGetHandler).Methods("GET")
    r.HandleFunc("/persist/{user}", persistPutHandler).Methods("PUT")
    r.HandleFunc("/persist/{user}", persistDeleteHandler).Methods("DELETE")

    return r
}

func main() {
    router := Handlers()
    http.ListenAndServe(":8080", router)
}
