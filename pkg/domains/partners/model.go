package partners

type Partner struct {
	ID           string       `json:"id" bson:"_id,omitempty"`
	TradingName  string       `json:"tradingName" bson:"tradingName"`
	OwnerName    string       `json:"ownerName" bson:"ownerName"`
	Document     string       `json:"document" bson:"document"`
	CoverageArea CoverageArea `json:"coverageArea" bson:"coverageArea"`
	Address      Address      `json:"address" bson:"address"`
}

type Address struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type CoverageArea struct {
	Type        string `json:"type" bson:"type"`
	Coordinates `json:"coordinates" bson:"coordinates"`
}

type Point struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Coordinates [][][][]float64
