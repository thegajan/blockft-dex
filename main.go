package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/thegajan/blockft-dex/tools"
  "github.com/thegajan/blockft-dex/account"
)

func main() {
  r := gin.Default()
  api :=  r.Group("/dex/api/v1")
  {
    api.GET("/", index)
    //ACCOUNT ROUTES
    api.GET("/createAccount", account.CreateAccountRoute)
    api.POST("/viewAccount", account.ViewAccountRoute)
    //PAYMENT ROUTES
    // api.POST("/payment", payment)
    //ASSET ROUTES
    // api.POST("/createAsset", createAsset)
  }
  r.Run()
}

func index(c *gin.Context) {
  tools.Response(c, http.StatusOK, "Welcome! You have reached the BlockFT Distribusted Exchange API. Refer to documentation for API endpoints.", nil)
}

// //TODO: validate account exists
// //      check asset is valid
// //      currently only works on native assets
// func payment(c *gin.Context) {
//   var p Payment
//   c.BindJSON(&p)
//   key := p.SK
//   destination := p.DEST
//   amount := p.Amount
//   asset := p.Asset
//   issuer := p.Issuer
//
//   kp, _ := keypair.Parse(key)
//   ar := horizonclient.AccountRequest{AccountID: kp.Address()}
//   sourceAccount, err := CLIENT.AccountDetail(ar)
//   if err != nil {
//     error(c, http.StatusInternalServerError, "Payment failed.")
//     log.Fatal(err)
//   }
//
//   var a txnbuild.Asset
//
//   if asset != "XLM" {
//     assetRequest := horizonclient.AssetRequest{ForAssetCode: asset, ForAssetIssuer: issuer}
//     hAsset0, err := CLIENT.Assets(assetRequest)
//     if err != nil {
//       error(c, http.StatusInternalServerError, "Payment failed.")
//       log.Fatal(err)
//     }
//     log.Println(hAsset0)
//     a = txnbuild.CreditAsset{Code: asset, Issuer: issuer}
//   } else {
//     a = txnbuild.NativeAsset{}
//   }
//
//   op := txnbuild.Payment{
//     Destination:    destination,
//     Amount:         amount,
//     Asset:          a,
//   }
//
//   tx := txnbuild.Transaction{
//       SourceAccount: &sourceAccount,
//       Operations:    []txnbuild.Operation{&op},
//       Timebounds:    txnbuild.NewTimeout(300),
//       Network:       "Standalone Network ; February 2017",
//   }
//
//   txe, err := tx.BuildSignEncode(kp.(*keypair.Full))
//   if err != nil {
//     hError := err.(*horizonclient.Error)
//     error(c, http.StatusInternalServerError, "Payment failed.")
//     log.Fatal("Error submitting transaction: ", hError)
//   }
//   _, err = CLIENT.SubmitTransactionXDR(txe)
//   if err != nil {
//       hError := err.(*horizonclient.Error)
//       error(c, http.StatusInternalServerError, "Payment failed.")
//       log.Fatal("Error submitting transaction: ", hError)
//   }
//
//   _ = asset
//
//   response(c, http.StatusOK, "Payment successful.", nil)
//   }
//
// //TODO: abstract account creation to a separate fucntion
// //       abstract account details to a separate function
// func createAsset(c *gin.Context) {
//   var a Asset
//   c.BindJSON(&a)
//   source := a.SK
//   asset := a.Asset
//   amount := a.Amount
//
//   //get source account details
//   source_kp, _ := keypair.Parse(source)
//   ar := horizonclient.AccountRequest{AccountID: source_kp.Address()}
//   sourceAccount, err := CLIENT.AccountDetail(ar)
//   if err != nil {
//     error(c, http.StatusInternalServerError, "Failed to create asset.")
//     log.Fatal(err)
//   }
//
//   //create issuing account
//   issuer_kp, err := keypair.Random()
//   if err != nil {
//     log.Fatal(err)
//     error(c, http.StatusInternalServerError, "Failed to create asset.")
//   }
//
//   createAccountOp := txnbuild.CreateAccount{
//       Destination: issuer_kp.Address(),
//       Amount:      "100",
//   }
//
//   tx := txnbuild.Transaction{
//       SourceAccount: &sourceAccount,
//       Operations:    []txnbuild.Operation{&createAccountOp},
//       Timebounds:    txnbuild.NewTimeout(300),
//       Network:       "Standalone Network ; February 2017",
//   }
//
//   txeBase64, err := tx.BuildSignEncode(source_kp.(*keypair.Full))
//
//   _, err = CLIENT.SubmitTransactionXDR(txeBase64)
//   if err != nil {
//       hError := err.(*horizonclient.Error)
//       error(c, http.StatusInternalServerError, "Failed to create asset.")
//       log.Fatal("Error submitting transaction: ", hError)
//   }
//
//   //distribution issuing account
//   distribution_kp, err := keypair.Random()
//   if err != nil {
//     log.Fatal(err)
//     error(c, http.StatusInternalServerError, "Failed to create asset.")
//   }
//
//   createAccountOp = txnbuild.CreateAccount{
//       Destination: distribution_kp.Address(),
//       Amount:      "100",
//   }
//
//   tx = txnbuild.Transaction{
//       SourceAccount: &sourceAccount,
//       Operations:    []txnbuild.Operation{&createAccountOp},
//       Timebounds:    txnbuild.NewTimeout(300),
//       Network:       "Standalone Network ; February 2017",
//   }
//
//   txeBase64, err = tx.BuildSignEncode(source_kp.(*keypair.Full))
//
//   _, err = CLIENT.SubmitTransactionXDR(txeBase64)
//   if err != nil {
//       hError := err.(*horizonclient.Error)
//       error(c, http.StatusInternalServerError, "Failed to create asset.")
//       log.Fatal("Error submitting transaction: ", hError)
//   }
//
//   //create asset and establish trust
//   ir := horizonclient.AccountRequest{AccountID: issuer_kp.Address()}
//   issueAccount, err := CLIENT.AccountDetail(ir)
//   if err != nil {
//     error(c, http.StatusInternalServerError, "Failed to create asset.")
//     log.Fatal(err)
//   }
//   dr := horizonclient.AccountRequest{AccountID: distribution_kp.Address()}
//   distAccount, err := CLIENT.AccountDetail(dr)
//   if err != nil {
//     error(c, http.StatusInternalServerError, "Failed to create asset.")
//     log.Fatal(err)
//   }
//
//   ca := txnbuild.CreditAsset{Code: asset, Issuer: issuer_kp.Address()}
//
//   changeTrustOp := txnbuild.ChangeTrust{
//     Line:           &ca,
//   }
//
//   tx = txnbuild.Transaction{
//       SourceAccount: &distAccount,
//       Operations:    []txnbuild.Operation{&changeTrustOp},
//       Timebounds:    txnbuild.NewTimeout(300),
//       Network:       "Standalone Network ; February 2017",
//   }
//
//   // dist_kp_parse, _ := keypair.Parse(distribution_kp)
//   txe, err := tx.BuildSignEncode(distribution_kp)
//   if err != nil {
//     hError := err.(*horizonclient.Error)
//     error(c, http.StatusInternalServerError, "Payment failed.")
//     log.Fatal("Error submitting transaction: ", hError)
//   }
//   _, err = CLIENT.SubmitTransactionXDR(txe)
//   if err != nil {
//       hError := err.(*horizonclient.Error)
//       error(c, http.StatusInternalServerError, "Payment failed.")
//       log.Fatal("Error submitting transaction: ", hError)
//   }
//
//   log.Println("trust established")
//
//   op := txnbuild.Payment{
//     Destination:    distribution_kp.Address(),
//     Amount:         amount,
//     Asset:          &ca,
//   }
//
//   tx = txnbuild.Transaction{
//       SourceAccount: &issueAccount,
//       Operations:    []txnbuild.Operation{&op},
//       Timebounds:    txnbuild.NewTimeout(300),
//       Network:       "Standalone Network ; February 2017",
//   }
//
//   txe, err = tx.BuildSignEncode(issuer_kp)
//   if err != nil {
//     hError := err.(*horizonclient.Error)
//     error(c, http.StatusInternalServerError, "Payment failed.")
//     log.Fatal("Error submitting transaction: ", hError)
//   }
//   _, err = CLIENT.SubmitTransactionXDR(txe)
//   if err != nil {
//       hError := err.(*horizonclient.Error)
//       error(c, http.StatusInternalServerError, "Payment failed.")
//       log.Fatal("Error submitting transaction: ", hError)
//   }
//
//   na := NewAsset{
//     Asset: asset,
//     Amount: amount,
//     Issuer: KeyPair{PK: issuer_kp.Address(), SK: issuer_kp.Seed()},
//     Distribution: KeyPair{PK: distribution_kp.Address(), SK: distribution_kp.Seed()},
//   }
//
//   response(c, http.StatusOK, asset + " created and " + amount + " issued.", &na)
// }
