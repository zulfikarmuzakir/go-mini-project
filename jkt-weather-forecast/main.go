package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type WeatherForecast struct {
	List []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			Temp float32 `json:"temp"`
		}
	}
}

func main() {
	response, err := http.Get("https://api.openweathermap.org/data/2.5/forecast?lat=-6.2146&lon=106.8451&units=metric&appid=0368687e851bdf3c05fbd03f7ad6a45b&lang=id")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weather WeatherForecast
	if err := json.Unmarshal(responseData, &weather); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 5; i++ {
		forecast := weather.List[i*8]
		date := time.Unix(forecast.Dt, 0).Format("Monday, 02 January 2006")
		temp := forecast.Main.Temp

		fmt.Printf("%s: %.2f°C\n", date, temp)
	}
}
