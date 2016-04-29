package persistent

import (
    "encoding/json"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestMongoWrapperTestSet(t *testing.T) {
    w, err := NewMongoWrapper("127.0.0.1", "test", "users")
    if(err != nil) {
        panic(err)
    }

    Convey("Mongo Wrapper testplan", t, func() {
        Convey("Create document", func() {
            w.Remove("john.doe@mail.com")
            res, err := w.Create("John", "Doe", "john.doe@mail.com")
            So(res, ShouldNotEqual, "")
            So(err, ShouldEqual, nil)
        })

        Convey("Delete document", func() {
            w.Remove("john.doe@mail.com")
            _, err := w.Create("John", "Doe", "john.doe@mail.com")
            So(err, ShouldEqual, nil)
            err = w.Remove("john.doe@mail.com")
            So(err, ShouldEqual, nil)
        })

        Convey("Find document", func() {
            w.Remove("john.doe@mail.com")
            _, err := w.Create("John", "Doe", "john.doe@mail.com")
            So(err, ShouldEqual, nil)
            res, err := w.Find("john.doe@mail.com")
            So(err, ShouldEqual, nil)
            var user User
            json.Unmarshal([]byte(res), &user)
            So(user.FirstName, ShouldEqual, "John")
            So(user.LastName, ShouldEqual, "Doe")
            So(user.EMail, ShouldEqual, "john.doe@mail.com")
        })

        Convey("Document is unique", func() {})
    })
}
