package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/stellar/go/keypair"
  "github.com/stellar/go/txnbuild"
  "github.com/thegajan/blockft-dex/src/account"
  "github.com/thegajan/blockft-dex/src/transaction"
)

type Asset struct {
  SK      string  `json:"key" binding:"required"`
  Asset   string  `json:"asset" binding:"required"`
  Amount  string  `json:"amount" binding:"required"`
}

type NewAsset struct {
  Asset         string  `json:"asset" binding:"required"`
  Amount        string  `json:"amount" binding:"required"`
  Issuer        KeyPair `json:"issuer" binding:"required"`
  Distribution  KeyPair `json:"distribution" binding:"required"`
}

func createAsset(c *gin.Context) {
  var a Asset
  c.BindJSON(&a)
  source := a.SK
  asset := a.Asset
  amount := a.Amount

  issuer_kp, dist_kp, err := transaction.CreateAsset(source, asset, amount)

  if err == nil {
    na := NewAsset{
      Asset: asset,
      Amount: amount,
      Issuer: KeyPair{PK: issuer_kp.Address(), SK: issuer_kp.Seed()},
      Distribution: KeyPair{PK: dist_kp.Address(), SK: dist_kp.Seed()},
    }

    response(c, http.StatusOK, asset + " created and " + amount + " issued.", &na)
  } else {
    ferror(c, err, http.StatusInternalServerError, "Failed to create asset.")
  }
}

type Trust struct {
  PK     string  `json:"key" binding:"required"`
  Asset  string  `json:"asset" binding:"required"`
  Issuer string  `json:"issuer" binding:"required"`
}

func establishTrust(c *gin.Context) {
  var t Trust
  c.BindJSON(&t)
  source := t.PK
  asset := t.Asset
  issuer := t.Issuer

  source_kp, _ := keypair.Parse(source)
  sourceAccount, err := account.RequestAccountDetails(source_kp.Address())
  if err != nil {
    ferror(c, err, http.StatusInternalServerError, "Failed to establish trust.")
  }

  ca := txnbuild.CreditAsset{Code: asset, Issuer: issuer}


  e := transaction.EstablishTrust(source_kp, sourceAccount, ca)

  if e == nil {
    response(c, http.StatusOK, "Trustline established.", nil)
  } else {
    ferror(c, err, http.StatusInternalServerError, "Failed to establish trust.")
  }
}
