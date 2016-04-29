package persistent

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
    "math/rand"
)

type MongoWrapper struct {
    s *mgo.Session
    c *mgo.Collection
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewMongoWrapper(addr, dbname, colname string) (*MongoWrapper, error) {
    w := new(MongoWrapper)
    var err error
    w.s, err = mgo.Dial(addr)
    if(err != nil) {
        return nil, err
    }
    w.s.SetMode(mgo.Monotonic, true)
    w.c = w.s.DB(dbname).C(colname)
    return w, nil
}

func createID() string {
    b := make([]rune, 12)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func (w *MongoWrapper) Close() {
    w.s.Close()
}

func (w *MongoWrapper) Create(firstName, lastName, EMail string) (string, error) {
    result := User{}
    id := createID()
    err := w.c.Find(bson.M{"firstname": firstName, "lastname": lastName, "email": EMail}).One(&result)
    if(err != nil) {
        err = w.c.Insert(&User{id, firstName, lastName, EMail})
        return id, err
    }
    id = result.ID
    return id, err
}

func (w *MongoWrapper) Find(email string) (string, error) {
    var result = User{}
    err := w.c.Find(bson.M{"email": email}).One(&result)
    if(err != nil) {
        return "", err
    }
    b, err := json.Marshal(result)
    if(err != nil) {
       return "", err
    }
    return string(b), nil
}

func (w *MongoWrapper) Remove(id string) error {
    colQuerier := bson.M{"id":id}
    err := w.c.Remove(colQuerier)
    return err
}
