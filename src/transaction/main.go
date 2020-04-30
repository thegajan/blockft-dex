package transaction

import (
  "github.com/stellar/go/keypair"
  "github.com/stellar/go/txnbuild"
  "github.com/stellar/go/clients/horizonclient"
  "github.com/stellar/go/protocols/horizon"
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

func EstablishTrust(source_kp keypair.KP, sourceAccount *horizon.Account, asset txnbuild.CreditAsset) error {
  changeTrustOp := txnbuild.ChangeTrust{
    Line:           &asset,
  }
  s := []txnbuild.Operation{&changeTrustOp}
  err := tools.Transaction(source_kp, sourceAccount, s)
  return err
}

func CreateAsset(source_sk string, asset string, amount string) (*keypair.Full, *keypair.Full, error) {
  source_kp, _ := keypair.Parse(source_sk)
  sourceAccount, err := account.RequestAccountDetails(source_kp.Address())
  if err != nil {
    return nil, nil, err
  }

  issuer_kp, err := account.CreateAccount(source_kp, sourceAccount, "100")
  dist_kp, err := account.CreateAccount(source_kp, sourceAccount, "100")
  if err != nil {
    return nil, nil, err
  }

  issuerAccount, err := account.RequestAccountDetails(issuer_kp.Address())
  distAccount, err := account.RequestAccountDetails(dist_kp.Address())
  if err != nil {
    return nil, nil, err
  }

  ca := txnbuild.CreditAsset{Code: asset, Issuer: issuer_kp.Address()}
  err = EstablishTrust(dist_kp, distAccount, ca)
  if err != nil {
    return nil, nil, err
  }

  err = Payment(issuer_kp.Seed(), issuerAccount, dist_kp.Address(), amount, ca)
  if err != nil {
    return nil, nil, err
  }

  return issuer_kp, dist_kp, err
}

//TODO: validate account exists
//      currently only works on native assets
//      need to trust non native assets
func Payment(souce_sk string, sourceAccount *horizon.Account, dest_pk string, amount string, asset txnbuild.Asset) error {
  source_kp, _ := keypair.Parse(souce_sk)

  paymentOp := txnbuild.Payment{
    Destination:    dest_pk,
    Amount:         amount,
    Asset:          asset,
  }
  s := []txnbuild.Operation{&paymentOp}

  err := tools.Transaction(source_kp, sourceAccount, s)

  return err
}
