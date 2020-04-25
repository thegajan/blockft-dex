package tools

type Payment struct {
  SK      string  `json:"key" binding:"required"`
  DEST    string  `json:"destination" binding:"required"`
  Amount  string  `json:"amount" binding:"required"`
  Asset   string  `json:"asset" binding:"required"`
  Issuer  string  `json:"issuer"`
}

type Asset struct {
  SK      string  `json:"key" binding:"required"`
  Asset   string  `json:"asset" binding:"required"`
  Amount  string  `json:"amount" binding:"required"`
}
