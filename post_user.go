package main

import (
    "log"
    "time"
    "go-swagger/client/operations"
    "go-swagger/models"
    apiclient "go-swagger/client"
    httptransport "github.com/go-openapi/runtime/client"
    "github.com/go-openapi/strfmt"
)

func main() {
    p := &operations.PostUserParams{
        Body: &models.User{
            ID: 1,
            Name: "name_lob",
        },
    }
    p.SetTimeout(10 * time.Second)

    transport := httptransport.New("localhost:8000", "api", nil)
    client := apiclient.New(transport,strfmt.Default)

    res, _ := client.Operations.PostUser(p)

    log.Printf("%#v\n", res.Error())
}