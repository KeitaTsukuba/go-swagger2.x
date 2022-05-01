package main

import (
    "log"
    "time"
    "go-swagger/client/operations"
    apiclient"go-swagger/client"
    httptransport "github.com/go-openapi/runtime/client"
    "github.com/go-openapi/strfmt"
)

func main()  {
    p := operations.GetUserByIDParams{
        ID: "1",
    }
    p.SetTimeout(10 * time.Second)

    transport := httptransport.New("localhost:8000", "api", nil)
    client := apiclient.New(transport,strfmt.Default)

    res, _ := client.Operations.GetUserByID(&p)

    log.Printf("%+v\n", res)
}