package main

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/stellar/go/clients/horizonclient"
  "github.com/stellar/go/keypair"
  // "github.com/stellar/go/network"
  "github.com/stellar/go/txnbuild"
)

func main()  {
  //connect to local horizon client
  client := horizonclient.Client{
    HorizonURL:     "http://localhost:8000",
    HTTP:           http.DefaultClient,
  }

  //root account seed
  master_kp, _ := keypair.Parse("SC5O7VZUXDJ6JBDSZ74DSERXL7W3Y5LTOAMRF7RQRL3TAGAPS7LUVG3L")


  // get master account details
	ar := horizonclient.AccountRequest{AccountID: master_kp.Address()}
	sourceAccount, err := client.AccountDetail(ar)
	if err != nil {
  		log.Println(err)
	}

  //print master balance
  account, err := json.Marshal(sourceAccount.Balances[0])
  if err != nil {
    log.Fatal(err)
    return
  }
  log.Println(string(account))

  //create a user account
  kp1, err := keypair.Random()
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Seed 1:", kp1.Seed())
  log.Println("Address 1:", kp1.Address())

  // //fund the new account
  createAccountOp := txnbuild.CreateAccount{
      Destination: kp1.Address(),
      Amount:      "100",
  }

  tx := txnbuild.Transaction{
      SourceAccount: &sourceAccount,
      Operations:    []txnbuild.Operation{&createAccountOp},
      Timebounds:    txnbuild.NewTimeout(300),
      Network:       "Standalone Network ; February 2017",
  }

  txeBase64, err := tx.BuildSignEncode(master_kp.(*keypair.Full))
  log.Println("Transaction base64: ", txeBase64)

  resp, err := client.SubmitTransactionXDR(txeBase64)
  if err != nil {
      hError := err.(*horizonclient.Error)
      // log.Fatal("Error submitting transaction: ", hError)
      // log.Fatal("Error Type: ", hError.Problem.Type)
      // log.Fatal("Error Details: ", hError.Problem.Detail)
      error, err := json.Marshal(hError.Problem)
      if err != nil {
        log.Fatal(err)
        return
      }
      log.Fatal(string(error))
  }

  log.Println("\nTransaction response: ", resp)
}
