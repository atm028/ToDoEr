package elasticservice

import (
    elastic "gopkg.in/olivere/elastic.v3"
    "math/rand"
    "fmt"
)

type ESService struct {
    c *elastic.Client
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewESService(addr, port string) (*ESService, error) {
    w := new(ESService)
    var err error
    w.c, err = elastic.NewClient()
    if err != nil {
        return nil, err
    }
    return w, nil
}

func (w *ESService) Create(index, tp, msg string) (string, error) {
    _id := string(letters[(rand.Intn(len(letters)))])
    _, err := w.c.Index().Index(index).Type(tp).Id(_id).BodyJson(msg).Do()
    return "", err
}

func (w *ESService) Search(index, key, val string) (string, error) {
    fmt.Println("ESService:Search: key = "+key+" val="+val+" in index "+index)
    termQuery := elastic.NewMatchQuery(key, val)
    res, err := w.c.Search().Index().Index(index).Query(termQuery).Pretty(true).Do()
    fmt.Println(res.Hits.TotalHits)
    var ret string
    var cnt int64
    cnt = 1
    ret += "["
    if res.Hits.TotalHits > 0 {
        for _, hit := range res.Hits.Hits {
            ret += string(*hit.Source)
            if cnt < res.Hits.TotalHits {
                ret += ","
            }
            cnt += 1
//            fmt.Println(ret)
//            return ret, nil
        }
    }
    ret += "]"
    return ret, err
}
