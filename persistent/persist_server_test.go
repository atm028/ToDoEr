package persistent

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
        req, err := http.NewRequest("DELETE", server.URL+"/persist/users?id="+resourse, nil)
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
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "email": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/users", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            body, err := ioutil.ReadAll(rsp.Body)
            So(body, ShouldNotEqual, "")
            So(err, ShouldEqual, nil)
            So(rsp.StatusCode, ShouldEqual, 201)
            res, err := clean(string(body), server, client)
            So(res, ShouldEqual, 200)
        })

        Convey("Find existing user", func() {
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "email": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/users", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 201)

            req, err = http.NewRequest("GET", server.URL+"/persist/users?email=john.doe@mail.com", nil)
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

            res, err := clean(doc.ID, server, client)
            So(res, ShouldEqual, 200)
        })

        Convey("Inserting same user only once", func() {

            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "email": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/users", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            id1, err := ioutil.ReadAll(rsp.Body)
            So(rsp.StatusCode, ShouldEqual, 201)

            jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "email": "john.doe@mail.com"}`)
            req, err = http.NewRequest("PUT", server.URL+"/persist/users", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            id2, err := ioutil.ReadAll(rsp.Body)
            So(rsp.StatusCode, ShouldEqual, 201)

            res, err := clean(string(id1), server, client)
            So(res, ShouldEqual, 200)
            res, err = clean(string(id2), server, client)
            So(res, ShouldEqual, 501)
        })

        Convey("Find several existing user with same First name but different ilast name and emails", func() {
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "email": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/users", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            id1, err := ioutil.ReadAll(rsp.Body)
            So(rsp.StatusCode, ShouldEqual, 201)

            jsonUser = []byte(`{"firstname": "John", "lastname": "Moe", "email": "john.moe@mail.com"}`)
            req, err = http.NewRequest("PUT", server.URL+"/persist/users", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            id2, err := ioutil.ReadAll(rsp.Body)
            So(rsp.StatusCode, ShouldEqual, 201)

            req, err = http.NewRequest("GET", server.URL+"/persist/users?firstname=John", nil)
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            So(rsp.StatusCode, ShouldEqual, 200)
            defer rsp.Body.Close()
//            body, err := ioutil.ReadAll(rsp.Body)
//            So(err, ShouldEqual, nil)
//            type jsonitem []string
            //TODO: implement when multiple search will be implemented
//            var datas jsonitem
//            json.Unmarshal(body, &datas)
//            So(len(datas), ShouldEqual, 2)
//            for ind := range datas {
//                var user User
//                json.Unmarshal([]byte(datas[ind]), &user)
//                So(user.FirstName, ShouldEqual, "John")
//            }

            res, err := clean(string(id1), server, client)
            So(res, ShouldEqual, 200)
            res, err = clean(string(id2), server, client)
            So(res, ShouldEqual, 200)
        })

        Convey("Update existing user", func() {
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "email": "john.doe@mail.com"}`)
            req, err := http.NewRequest("PUT", server.URL+"/persist/users", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            id, err := ioutil.ReadAll(rsp.Body)
            So(rsp.StatusCode, ShouldEqual, 201)

            jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "email": "john.doe_work@mail.com"}`)
            req, err = http.NewRequest("POST", server.URL+"/persist/users?id="+string(id), bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err = client.Do(req)
            So(err, ShouldEqual, nil)
            defer rsp.Body.Close()
            So(rsp.StatusCode, ShouldEqual, 200)

            body, err := ioutil.ReadAll(rsp.Body)
            So(err, ShouldEqual, nil)
            var doc User
            json.Unmarshal(body, &doc)
            //TODO: uncomment when update will be implemented
            //So(doc.EMail, ShouldEqual, "john.doe_work@mail.com")

            res, err := clean(string(id), server, client)
            So(res, ShouldEqual, 200)
        })

        Convey("Find non-existing user", func() {
            req, err := http.NewRequest("GET", server.URL+"/persist/users?firstname=John,lastname=Doe", nil)
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            //TODO: change to expected 404 when multiple search will be implemented
            So(rsp.StatusCode, ShouldEqual, 200)
        })

        Convey("Update non-existing user", func() {
            var jsonUser = []byte(`{"firstname": "John", "lastname": "Doe", "email": "john.doe_work@mail.com"}`)
            req, err := http.NewRequest("POST", server.URL+"/persist/users", bytes.NewBuffer(jsonUser))
            So(err, ShouldEqual, nil)
            rsp, err := client.Do(req)
            So(err, ShouldEqual, nil)
            //TODO: change expected value to 412 when updte will be implemented
            So(rsp.StatusCode, ShouldEqual, 200)
        })

    })
}
