package nationbuilder

// Address Resource
// TODO: format of lat/lng?
type Address struct {
	FirstLine   string  `json:"address1,omitempty"`
	SecondLine  string  `json:"address2,omitempty"`
	ThirdLine   string  `json:"address3,omitempty"`
	City        string  `json:"city,omitempty"`
	State       string  `json:"state,omitempty"`
	ZIPCode     string  `json:"zip,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	Latitude    float32 `json:"lat,omitempty"`
	Longtitude  float32 `json:"lng,omitempty"`
}
