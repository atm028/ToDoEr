package elasticservice

import (
    elastic "gopkg.in/olivere/elastic.v3"
    "math/rand"
)

type ESService struct {
    c *elastic.Client
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewESService(addr, port string) (*ESService, error) {
    w := New(ESService)
    var err error
    w.c, err := elastic.NewClient()
    if err != nil {
        nil, err
    }
    return w, nil
}

func (w *ESService) Create(index, tp, msg string) error {
    _. err := w.c.Index().Index(index).Type(tp).Id(letters[rand.Intn(len(letters))].BodyJson(msg).Do()
    return err
}

func (w *ESService) Search(index, key, val string) (string, error) {
    termQuery := elastc.NewTermQuery(key, val)
    res, err := w.c.Search().Index(index).Query(termQuery).Do()
    return res, err
}
