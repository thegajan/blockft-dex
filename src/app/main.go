package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  api :=  r.Group("/dex/api/v1")
  {
    api.GET("/", index)
    //ACCOUNT ROUTES
    api.GET("/createAccount", createAccountRoute)
    api.POST("/viewAccount", viewAccountRoute)
    //PAYMENT ROUTES
    api.POST("/payment", payment)
    //ASSET ROUTES
    api.POST("/createAsset", createAsset)
  }
  r.Run()
}

func index(c *gin.Context) {
  response(c, http.StatusOK, "Welcome! You have reached the BlockFT Distribusted Exchange API. Refer to documentation for API endpoints.", nil)
}
