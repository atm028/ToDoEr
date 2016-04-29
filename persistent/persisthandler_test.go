package persistent

import (
    "encoding/json"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestPersistHandlerTestSet(t *testing.T) {
    w, err := NewPersistHandler()
    if(err != nil) {
        panic(err)
    }
    doc := `{"firstname": "John", "lastname": "Doe", "email":"john.doe@mail.com"}`
    var actUser User
    err = json.Unmarshal([]byte(doc), &actUser)
    if(err != nil) {
        panic(err)
    }

    Convey("PersistHandler testplan", t, func() {
        Convey("Create user", func() {
            w.persistHandlerDelete(doc)
            rsp, code := w.persistHandlerCreate(doc)
            So(rsp, ShouldNotEqual, "")
            So(code, ShouldEqual, 201)
        })

        Convey("Delete user", func() {
            w.persistHandlerDelete(doc)
            res, code := w.persistHandlerCreate(doc)
            So(code, ShouldEqual, 201)
            res, code = w.persistHandlerDelete(doc)
            So(res, ShouldEqual, "")
            So(code, ShouldEqual, 200)
        })

        Convey("Find user", func() {
            w.persistHandlerDelete(doc)
            _, code := w.persistHandlerCreate(doc)
            So(code, ShouldEqual, 201)
            _, code = w.persistHandlerRead(doc)
            var resUser User
            err = json.Unmarshal([]byte(doc), &resUser)

            So(actUser.FirstName, ShouldEqual, resUser.FirstName)
            So(actUser.LastName, ShouldEqual, resUser.LastName)
            So(actUser.EMail, ShouldEqual, resUser.EMail)
            So(code, ShouldEqual, 200)
        })
    })
}
