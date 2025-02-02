package account

import (
  "github.com/stellar/go/keypair"
  "github.com/thegajan/blockft-dex/src/tools"
  "github.com/stellar/go/clients/horizonclient"
  "github.com/stellar/go/protocols/horizon"
  "github.com/stellar/go/txnbuild"
)

func RequestAccountDetails(account string) (*horizon.Account, error) {
  accountRequest := horizonclient.AccountRequest{AccountID: account}
  hAccount0, err := tools.CLIENT.AccountDetail(accountRequest)
  return &hAccount0, err
}

func CreateAccount(source keypair.KP, sourceAccount *horizon.Account, amount string) (*keypair.Full, error) {
  kp, err := keypair.Random()
  if err != nil {
    return kp, err
  }

  createAccountOp := txnbuild.CreateAccount{
      Destination: kp.Address(),
      Amount:      amount,
  }
  s := []txnbuild.Operation{&createAccountOp}

  err = tools.Transaction(source, sourceAccount, s)

  return kp, err
}
