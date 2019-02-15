package dataobjects

//InitiateRequest dataobject to model initate request
type InitiateRequest struct {
	Subject      string `json:"subject"`
	Questions    uint   `json:"questions"`
	Duration     uint   `json:"duration"`
	UserID       uint
	TestDetailID uint
}
