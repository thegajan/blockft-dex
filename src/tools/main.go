package tools

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/stellar/go/clients/horizonclient"
  "github.com/stellar/go/protocols/horizon"
  "github.com/stellar/go/keypair"
  "github.com/stellar/go/txnbuild"
)

const ROOT_ACCOUNT_SEED = "SC5O7VZUXDJ6JBDSZ74DSERXL7W3Y5LTOAMRF7RQRL3TAGAPS7LUVG3L"

var ROOT_ACCOUNT_SEED_KP, _ = keypair.Parse(ROOT_ACCOUNT_SEED)


var CLIENT = horizonclient.Client{
  HorizonURL:     "http://localhost:8000",
  HTTP:           http.DefaultClient,
}

func Transaction(s keypair.KP, a *horizon.Account, op txnbuild.Operation) error {
  tx := txnbuild.Transaction{
      SourceAccount: a,
      Operations:    []txnbuild.Operation{op},
      Timebounds:    txnbuild.NewTimeout(300),
      Network:       "Standalone Network ; February 2017",
  }

  txeBase64, err := tx.BuildSignEncode(s.(*keypair.Full))

  _, err = CLIENT.SubmitTransactionXDR(txeBase64)


  if err != nil {
    // error(c, http.StatusInternalServerError, "Failed to fund account.")
    hError := err.(*horizonclient.Error)
    log.Println(json.Marshal(hError.Problem))
    log.Fatal("Error submitting transaction: ", hError)
  }

  return err
}
