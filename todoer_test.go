package todoer

import (
    "encoding/json"
    "io/ioutil"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "net/http"
    "net/http/httptest"
    "bytes"
)

func clean(resourse string, server *httptest.Server, client http.Client) (res int, err error) {
        req, err := http.NewRequest("DELETE", server.URL+"/persist/"+resourse, nil)
        if( err != nil) {
            return 0, err
        }
        rsp, err := client.Do(req)
        if(err != nil) {
            return 0, err
        }
        return rsp.StatusCode, nil
}

func TestToDoErAPITestSet(t *testing.T) {
    var server = httptest.NewServer(Handlers())
    var client = http.Client{}

    Convey("ToDoEr testplan", t, func() {
        Convey("Create new user and delete them", func() {
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 201)
            res, err := clean("user1", server, client)
            So(res, ShouldEqual, 200)
        })

        Convey("Find existing user", func() {
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 201)

            req, err = http.NewRequest("GET", server.URL+"/persist/user?firstname=John,lastname=Doe", nil)
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            So(rsp.StatusCode, ShouldEqual, 200)
            defer rsp.Body.Close()
            body, err := ioutil.ReadAll(rsp.Body)
            So(err, ShouldEqual, nil)
            var doc User
            json.Unmarshal(body, &doc)
            So(doc.FirstName, ShouldEqual, "John")
            So(doc.LastName, ShouldEqual, "Doe")
            So(doc.EMail, ShouldEqual, "john.doe@mail.com")

            res, err := clean("user1", server, client)
            So(res, ShouldEqual, 200)
        })

        Convey("Inserting same user only once", func() {

            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 201)

            jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe@mail.com"}`)
            req, err = http.NewRequest("PUT", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 409)

            res, err := clean("user1", server, client)
            So(res, ShouldEqual, 200)
            res, err = clean("user2", server, client)
            So(res, ShouldEqual, 200)
        })

        Convey("Find several existing user with same First name but different ilast name and emails", func() {
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 201)

            jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe@mail.com"}`)
            req, err = http.NewRequest("PUT", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 409)

            jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe@mail.com"}`)
            req, err = http.NewRequest("PUT", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 201)

            jsonUser = []byte(`{"firstname": "John", "lastname": "Moe", "mail": "john.moe@mail.com"}`)
            req, err = http.NewRequest("PUT", server.URL+"/persist/user2", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 201)

            req, err = http.NewRequest("GET", server.URL+"/persist/user?firstname=John", nil)
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            So(rsp.StatusCode, ShouldEqual, 200)
            defer rsp.Body.Close()
            body, err := ioutil.ReadAll(rsp.Body)
            So(err, ShouldEqual, nil)
            type jsonitem []string
            var datas jsonitem
            json.Unmarshal(body, &datas)
            So(len(datas), ShouldEqual, 2)
            for ind := range datas {
                var user User
                json.Unmarshal([]byte(datas[ind]), &user)
                So(user.FirstName, ShouldEqual, "John")
            }

            res, err := clean("user1", server, client)
            So(res, ShouldEqual, 200)
            res, err = clean("user2", server, client)
            So(res, ShouldEqual, 200)
        })

        Convey("Update existing user", func() {
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 201)

            jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe_work@mail.com"}`)
            req, err = http.NewRequest("POST", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 200)

            body, err := ioutil.ReadAll(rsp.Body)
            So(err, ShouldEqual, nil)
            var doc User
            json.Unmarshal(body, &doc)
            So(doc.EMail, ShouldEqual, "john.doe_work@mail.com")

            res, err := clean("user1", server, client)
            So(res, ShouldEqual, 200)
        })

        Convey("Find non-existing user", func() {
            req, err := http.NewRequest("GET", server.URL+"/persist/user?firstname=John,lastname=Doe", nil)
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            So(rsp.StatusCode, ShouldEqual, 404)
        })

        Convey("Update non-existing user", func() {
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "mail": "john.doe_work@mail.com"}`)
            req, err := http.NewRequest("POST", server.URL+"/persist/user1", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            So(rsp.StatusCode, ShouldEqual, 412)
        })

    })
}
