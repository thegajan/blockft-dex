package tools

type Payload struct {
  Status    int         `json:"status"`
  Message   string      `json:"message"`
  Data      interface{} `json:"data"`
}

type KeyPair struct {
  PK      string    `json:"public_key"`
  SK      string    `json:"private_key"`
}

type NewAsset struct {
  Asset         string  `json:"asset" binding:"required"`
  Amount        string  `json:"amount" binding:"required"`
  Issuer        KeyPair `json:"issuer" binding:"required"`
  Distribution  KeyPair `json:"distribution" binding:"required"`
}
