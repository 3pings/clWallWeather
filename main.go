package main

import (
	"database/sql"
	"github.com/3pings/clWallWeather/config"
	"github.com/3pings/clWallWeather/weather"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"log/syslog"
	"os"
	"time"
)

type weatherData struct {
	Description string
	Icon        string
	Temp        float64
	TempMin     float64
	TempMax     float64
	Timestamp   int
	Humidity    int
	City        string
}

func main() {
	apiKey := os.Getenv("owApiKey")
	cityCode := "Barcelona,es"
	units := "metric"
	//Get Weather Data

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "weather")
	if e == nil {
		log.SetOutput(logwriter)
	}
	for {
		//Get Weather information from Openweathermap
		w := weather.GetWeather(apiKey, cityCode, units)

		//Parse fields for specific information
		wd := weatherData{}
		wd.Description = w.WeatherInfo[0].Description
		wd.Icon = w.WeatherInfo[0].Icon
		wd.Temp = w.Core.Temp
		wd.TempMin = w.Core.TempMin
		wd.TempMax = w.Core.TempMax
		wd.Timestamp = w.Dt
		wd.Humidity = w.Core.Humidity
		wd.City = w.Name

		//Run Function to insert Data into DB
		id := insertData(config.DB, wd)
		if id != nil {
			log.Fatalln(id)
		}
		time.Sleep(600 * time.Second)
	}
}

//noinspection ALL
func insertData(s *sql.DB, w weatherData) error {

	////Delete data from DB
	//_, err := s.Exec("delete from weather")
	//Insert Data into Database

	_, err := s.Exec("INSERT weather(Description, Icon, Temp, Temp_min, Temp_max, Timestamp, Humidity, City) VALUES(?,?,?,?,?,?,?,?)", w.Description, w.Icon, w.Temp, w.TempMin, w.TempMax, w.Timestamp, w.Humidity, w.City)
	log.Print("Successfully created DB record for weather info")

	return err

}
