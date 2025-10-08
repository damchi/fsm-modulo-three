package domain

// BinaryRequest defines the input payload for /check
type BinaryRequest struct {
	Binary string `json:"binary" binding:"required" example:"110"`
	Mod    int    `json:"mod" binding:"required,min=1" example:"3"`
}

// BinaryResponse defines the API output
type BinaryResponse struct {
	Binary      string `json:"binary"`
	Remainder   int    `json:"remainder"`
	IsDivisible bool   `json:"is_divisible"`
}
