package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
type weatherInfo []struct {
	ID          int    `json:"id"`
	Style       string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type core struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}
type wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}
type clouds struct {
	All int `json:"all"`
}
type sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}
type data struct {
	Coord       coord
	WeatherInfo weatherInfo `json:"weather"`
	Base        string      `json:"base"`
	Core        core        `json:"main"`
	Visibility  int         `json:"visibility"`
	Wind        wind
	Clouds      clouds
	Dt          int `json:"dt"`
	Sys         sys
	Timezone    int    `json:"timezone"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	//Cod         string `json:"cod"`
}

//Get Weather Data and place it into a Struct
func GetWeather(key, city, measurement string) (d data) {
	//Set Variables
	var w data
	baseUrl := "https://api.openweathermap.org/data/2.5/weather?"
	url := baseUrl + "q=" + city + "&units=" + measurement + "&appid=" + key
	//Build New Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("error with GET Request", err)
	}
	//Get Response from Request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("error with GET Response", err)
	}
	//fmt.Println(url)
	defer res.Body.Close()
	//Unmarshal Json into data Struct
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}
	return w

}
