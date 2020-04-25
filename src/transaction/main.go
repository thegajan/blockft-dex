package transaction

import (
  "github.com/stellar/go/keypair"
  "github.com/stellar/go/txnbuild"
  "github.com/stellar/go/clients/horizonclient"
  "github.com/thegajan/blockft-dex/src/tools"
  "github.com/thegajan/blockft-dex/src/account"
)

func FindAsset(asset string, issuer string) (txnbuild.Asset, error) {
  var a txnbuild.Asset

  if asset != "XLM" {
    assetRequest := horizonclient.AssetRequest{ForAssetCode: asset, ForAssetIssuer: issuer}
    _, err := tools.CLIENT.Assets(assetRequest)
    if err != nil {
      return a, err
    }
    a = txnbuild.CreditAsset{Code: asset, Issuer: issuer}
  } else {
    a = txnbuild.NativeAsset{}
  }
  return a, nil
}

//TODO: validate account exists
//      check asset is valid
//      currently only works on native assets
//      need to trust non native assets
func Payment(souce_sk string, dest_pk string, amount string, asset txnbuild.Asset) error {
  source_kp, _ := keypair.Parse(souce_sk)
  sourceAccount, err := account.RequestAccountDetails(source_kp.Address())
  if err != nil {
    return err
  }

  paymentOp := txnbuild.Payment{
    Destination:    dest_pk,
    Amount:         amount,
    Asset:          asset,
  }

  err = tools.Transaction(source_kp, sourceAccount, &paymentOp)

  return err
}
