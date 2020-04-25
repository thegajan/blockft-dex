package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/thegajan/blockft-dex/src/tools"
  "github.com/thegajan/blockft-dex/src/account"
)

type Account struct {
  PK  string  `json:"account" binding:"required"`
}

func viewAccountRoute(c *gin.Context) {
  var a Account
  c.BindJSON(&a)
  acnt := a.PK

  hAccount0, err := account.RequestAccountDetails(acnt)
  if err == nil {
    balance := hAccount0.Balances
    response(c, http.StatusOK, "Account balance found.", balance)
  } else {
    ferror(c, err, http.StatusInternalServerError, "Failed to find account information.")
  }
}

func createAccountRoute(c *gin.Context) {
  sourceAccount, err := account.RequestAccountDetails(tools.ROOT_ACCOUNT_SEED_KP.Address())

  kp, err := account.CreateAccount(tools.ROOT_ACCOUNT_SEED_KP, sourceAccount, "100")

  if err == nil {
    d := KeyPair{
      PK: kp.Address(),
      SK: kp.Seed(),
    }
    response(c, http.StatusOK, "New account created.", &d)
  } else {
    ferror(c, err, http.StatusInternalServerError, "Failed to create account.")
  }
}
