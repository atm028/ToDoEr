package elasticservice

import (
    "encoding/json"
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

func (w *ElasticHandler) ElasticHandlerCreate(index, tp, query string) (string, error) {
    rsp, err := w.h.Create(index, tp, query)
    if(err != nill) {
        return `{"reason": "Problem with document creation:", "index": "`+index+
        `", "type":"`+tp+`", "value": "`+query+`", "error": "`+err.Erro()+`"}`, 500
    }
    return rsp, 201
}

func (w *ElasticHandler) ElasticHandlerSearch(index, key, value string) (string, error) {
    rsp, err := w.h.Search(index, key, value)
    if(err != nil) {
        return `{"reason": "Problem with document quering", "index": "`+index+
        `", "key": "`+key+`", "value":"`+value+`", "error": "`+err.Error()+`"}`, 500
    }
    return rsp, 200
}

