package partners

type Partner struct {
	ID           string       `json:"id"`
	TradingName  string       `json:"tradingName"`
	OwnerName    string       `json:"ownerName"`
	Document     string       `json:"document"`
	CoverageArea CoverageArea `json:"coverageArea"`
	Address      Address      `json:"address"`
}

type Address struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type CoverageArea struct {
	Type        string `json:"type"`
	Coordinates `json:"coordinates"`
}

type Point struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Pdvs struct {
	Pdvs []Partner `json:"pdvs"`
}

type Coordinates [][][][]int64
