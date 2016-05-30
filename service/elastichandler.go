package elasticservice

import (
    "fmt"
    //"encoding/json"
)

type ElasticHandler struct {
    h *ESService
}

func NewElasticHandler() (*ElasticHandler, error) {
    w := new(ElasticHandler)
    var err error
    w.h, err = NewESService("localhost", "9200")
    if(err != nil) {
        return nil, err
    }
    return w, nil
}

func (w *ElasticHandler) Create(index, tp, query string) (string, int) {
    _, err := w.h.Create(index, tp, query)
    if(err != nil) {
        return `{"reason": "Problem with document creation:", "index": "`+index+
        `", "type":"`+tp+`", "value": "`+query+`", "error": "`+err.Error()+`"}`, 500
    }
    return "", 201
}

func (w *ElasticHandler) Search(index, key, value string) (string, int) {
    fmt.Println("ElasticHandler:Search")
    rsp, err := w.h.Search(index, key, value)
    if(err != nil) {
        fmt.Println("ElasticHandler:Search Error")
        return `{"reason": "Problem with document quering", "index": "`+index+
        `", "key": "`+key+`", "value":"`+value+`", "error": "`+err.Error()+`"}`, 500
    }
    return rsp, 200
}

