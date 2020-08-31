package domain

type Data struct {
	Coord 		Coordinate 	`json:"coord"`
	Weather 	[]Weather 	`json:"weather"`
	Base 		string 		`json:"base"`
	Main 		Main 		`json:"main"`
	Visibility 	int 		`json:"visibility"`
	Wind		Wind		`json:"wind"`
	Clouds 		Cloud		`json:"clouds"`
	Dt  		int 		`json:"dt"`
	Sys 		Sys 		`json:"sys"`
	ID   		int    		`json:"id"`
	Name 		string 		`json:"name"`
	Cod  		int    		`json:"cod"`
}

type Coordinate struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Cloud struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}