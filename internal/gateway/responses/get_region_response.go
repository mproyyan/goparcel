package responses

type RegionResponse struct {
	ZipCode     string `json:"zip_code"`
	Province    string `json:"province_name"`
	City        string `json:"city_name"`
	District    string `json:"district_name"`
	Subdistrict string `json:"subdistrict_name"`
}
