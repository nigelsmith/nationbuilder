package nationbuilder

// The address resource represents an address within many nationbuilder types such as people
type Address struct {
	FirstLine   string  `json:"address1,omitempty"`
	SecondLine  string  `json:"address2,omitempty"`
	ThirdLine   string  `json:"address3,omitempty"`
	City        string  `json:"city,omitempty"`
	State       string  `json:"state,omitempty"`
	ZIPCode     string  `json:"zip,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	Latitude    float32 `json:"lat,omitempty"` // TODO: format of lat/lng?
	Longtitude  float32 `json:"lng,omitempty"`
}
