package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ellieasager/helloGoWeather/weather"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// write `Hello` using `io.WriteString` function
	io.WriteString(res, "Hello")
	// write `World` using `fmt.Fprint` function
	fmt.Fprint(res, " World! ")
	// write `❤️` using simple `Write` call
	res.Write([]byte("❤️"))
}

func query(city string) (weather.WeatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=<YOUR_API_KEY>&q=" + city)
	if err != nil {
		return weather.WeatherData{}, err
	}
	defer resp.Body.Close()
	var d weather.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weather.WeatherData{}, err
	}
	return d, nil
}

func main() {

	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World! ❤️"))
	})

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application / json; charset = utf-8")
		json.NewEncoder(w).Encode(data)
	})

	// listen and serve using `http.DefaultServeMux`
	http.ListenAndServe(":9000", nil)
}
