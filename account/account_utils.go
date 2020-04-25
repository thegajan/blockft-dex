package account

import (
  "github.com/stellar/go/keypair"
  "github.com/thegajan/blockft-dex/tools"
  "github.com/stellar/go/clients/horizonclient"
  "github.com/stellar/go/protocols/horizon"
  "github.com/stellar/go/txnbuild"
)

func requestAccountDetails(account string) (*horizon.Account, error) {
  accountRequest := horizonclient.AccountRequest{AccountID: account}
  hAccount0, err := tools.CLIENT.AccountDetail(accountRequest)
  return &hAccount0, err
}

func CreateAccount(source keypair.KP) (*keypair.Full, error) {
  kp, err := keypair.Random()
  if err != nil {
    return kp, err
  }

  sourceAccount, err := requestAccountDetails(source.Address())
  if err != nil {
    return kp, err
  }

  createAccountOp := txnbuild.CreateAccount{
      Destination: kp.Address(),
      Amount:      "100",
  }

  err = tools.Transaction(source, sourceAccount, &createAccountOp)

  return kp, err
}
