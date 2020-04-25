package account

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/thegajan/blockft-dex/tools"
)

type Account struct {
  PK  string  `json:"account" binding:"required"`
}

func ViewAccountRoute(c *gin.Context) {
  var a Account
  c.BindJSON(&a)
  account := a.PK

  hAccount0, err := requestAccountDetails(account)
  if err == nil {
    balance := hAccount0.Balances
    tools.Response(c, http.StatusOK, "Account balance found.", balance)
  } else {
    tools.Error(c, err, http.StatusInternalServerError, "Failed to find account information.")
  }
}

func CreateAccountRoute(c *gin.Context) {
  kp, err := CreateAccount(tools.ROOT_ACCOUNT_SEED_KP)

  if err == nil {
    d := tools.KeyPair{
      PK: kp.Address(),
      SK: kp.Seed(),
    }
    tools.Response(c, http.StatusOK, "New account created.", &d)
  } else {
    tools.Error(c, err, http.StatusInternalServerError, "Failed to create account.")
  }
}
