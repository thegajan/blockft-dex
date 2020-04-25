package main

import (
  "log"
  "github.com/gin-gonic/gin"
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

type NewAsset struct {
  Asset         string  `json:"asset" binding:"required"`
  Amount        string  `json:"amount" binding:"required"`
  Issuer        KeyPair `json:"issuer" binding:"required"`
  Distribution  KeyPair `json:"distribution" binding:"required"`
}

func ferror(c *gin.Context, err error, status int, message string) {
  errorResponse(c, status, message)
  log.Println(err)
}

func response(c *gin.Context, status int, message string, data interface{}) {
  payload := Payload{
    Status: status,
    Message: message,
    Data: data,
  }
  c.SecureJSON(status, payload)
}

func errorResponse(c *gin.Context, status int, message string) {
  payload := Payload{
    Status: status,
    Message: message,
  }
  c.SecureJSON(status, payload)
}
