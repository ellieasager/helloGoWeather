package weather

// WeatherData models this structure:
/*
  "main": {
    "temp": 299.99
  },
  "name": "Tokyo"
*/

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}
