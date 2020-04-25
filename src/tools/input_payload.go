package tools

type Asset struct {
  SK      string  `json:"key" binding:"required"`
  Asset   string  `json:"asset" binding:"required"`
  Amount  string  `json:"amount" binding:"required"`
}
