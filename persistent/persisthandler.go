package persistent

import (
    "encoding/json"
)

type PersistHandler struct {
    h *MongoWrapper
}

func NewPersistHandler() (*PersistHandler, error) {
    //TODO:read DB access from config and/or from consul
    w := new(PersistHandler)
    var err error
    w.h, err = NewMongoWrapper("127.0.0.1", "test", "users")
    if(err != nil) {
        return nil, err
    }
    return w, nil
}

func (w *PersistHandler)PersistHandlerCreate(query string) (string ,int) {
    var dat User
    err := json.Unmarshal([]byte(query), &dat)
    if(err != nil) {
        return `{"reason": "Problem with query encoding", "error": "`+err.Error()+`"}`, 500
    }

    rsp, err := w.h.Create(dat.FirstName, dat.LastName, dat.EMail)
    if(err != nil) {
        return `{"reason": "Problem with creation of user with EMail:"`+dat.EMail+
            `"error": "`+err.Error()+`"}`, 500
    }
    return rsp, 201
}

func (w *PersistHandler)PersistHandlerRead(query string) (string, int) {
    var dat User
    err := json.Unmarshal([]byte(query), &dat)
    if(err != nil) {
        return `{"reason": "Problem with handling of user with EMail:"`+dat.EMail+
            `"error": "`+err.Error()+`"}`, 500
    }

    res, err := w.h.Find(dat.EMail)
    if(err != nil) {
        return `{"reason": "Problem with reading of user with EMail:"`+dat.EMail+
            `"error": "`+err.Error()+`"}`, 500
    }
    return res, 200
}

func (w *PersistHandler)PersistHandlerUpdate(query, value string) (string, int) {
    return "", 200
}

func (w *PersistHandler)PersistHandlerDelete(query string) (string, int) {
    var dat User
    err := json.Unmarshal([]byte(query), &dat)
    if(err != nil) {
        return `{"reason": "Problem with query encoding", "error": "`+err.Error()+`"}`, 500
    }
    err = w.h.Remove(dat.ID)
    if(err != nil) {
        return `{"reason": "Problem with handling of user with EMail:"`+dat.EMail+
            `"error": "`+err.Error()+`"}`, 501
    }
    return "", 200
}
