package responsemodels

type WealthInformationResponse struct {
	AssetId     uint    `json:"assetid"`
	Amount      float64 `json:"amount"`
	UsdAmount   float64 `json:"usd_amount"`
	Name        string  `json:"name"`
	ImageSource string  `json:"image_source"`
}
