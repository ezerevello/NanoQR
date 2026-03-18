package model

type JsonRequest struct {
    Input string `json:"input"`
    Size  int `json:"size"`
    RecoverLevel string `json:"recoverLevel"`
}