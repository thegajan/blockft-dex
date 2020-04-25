package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/stellar/go/keypair"
  "github.com/thegajan/blockft-dex/src/transaction"
  "github.com/thegajan/blockft-dex/src/account"
)

type Payment struct {
  SK      string  `json:"key" binding:"required"`
  DEST    string  `json:"destination" binding:"required"`
  Amount  string  `json:"amount" binding:"required"`
  Asset   string  `json:"asset" binding:"required"`
  Issuer  string  `json:"issuer"`
}

func payment(c *gin.Context) {
  var p Payment
  c.BindJSON(&p)
  key := p.SK
  destination := p.DEST
  amount := p.Amount
  asset := p.Asset
  issuer := p.Issuer

  source_kp, err := keypair.Parse(key)
  sourceAccount, err := account.RequestAccountDetails(source_kp.Address())

  assetInterface, err := transaction.FindAsset(asset, issuer)
  if err != nil {
    ferror(c, err, http.StatusInternalServerError, "Payment failed.")
  }

  err = transaction.Payment(key, sourceAccount, destination, amount, assetInterface)

  if err == nil {
    response(c, http.StatusOK, "Payment successful", nil)
  } else {
    ferror(c, err, http.StatusInternalServerError, "Payment failed.")
  }
}
