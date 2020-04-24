package main

import (
  // "log"
  "net/http"
  // "encoding/json"
  "github.com/gin-gonic/gin"
  // "github.com/stellar/go/clients/horizonclient"
  "github.com/stellar/go/keypair"
  // "github.com/stellar/go/network"
  // "github.com/stellar/go/txnbuild"
)

type Payload struct {
  Status    int         `json:"status"`
  Message   string      `json:"message"`
  Data      interface{} `json:"data"`
}

type KeyPair struct {
  PK      string    `json:"public_key"`
  SK      string    `json:"private_key"`
}

func main() {
  r := gin.Default()
  api :=  r.Group("/dex/api/v1")
  {
    api.GET("/", index)
    api.GET("/newAccount", newAccount)
    // api.PATCH("/fundAccount", fundAccount)
    // api.GET("/getAccount", getAccount)
    // api.DELETE("/deleteAccount", deleteAccount)
  }
  r.Run()
}

func index (c *gin.Context) {
  payload := Payload{
    Status: http.StatusOK,
    Message: "Welcome! You have reached the BlockFT Distribusted Exchange API. Refer to documentation for API endpoints.",
  }
  c.JSON(http.StatusOK, payload)
}

func newAccount (c *gin.Context) {
  kp, err := keypair.Random()
  if err != nil {
    payload := Payload{
      Status: http.StatusInternalServerError,
      Message: "Failed to create new account.",
    }
    c.JSON(http.StatusInternalServerError, payload)
  }
  _ = kp
  // payload := Payload{Status: http.StatusOK, Message: "Account Created."}
  d := KeyPair{
    PK: kp.Address(),
    SK: kp.Seed(),
  }
  payload := Payload{
    Status: http.StatusOK,
    Message: "New account created.",
    Data: d,
  }
  c.JSON(http.StatusOK, payload)
}
