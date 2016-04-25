package todoer

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
)

type MongoWrapper struct {
    s *mgo.Session
    c *mgo.Collection
}

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

func (w *MongoWrapper) Close() {
    w.s.Close()
}

func (w *MongoWrapper) Create(firstName, lastName, EMail string) error {
    result := User{}
    err := w.c.Find(bson.M{"firstname": firstName, "lastname": lastName, "email": EMail}).One(&result)
    if(err != nil) {
        err = w.c.Insert(&User{firstName, lastName, EMail})
    }
    return err
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

func (w *MongoWrapper) Remove(email string) error {
    colQuerier := bson.M{"email":email}
    err := w.c.Remove(colQuerier)
    return err
}
