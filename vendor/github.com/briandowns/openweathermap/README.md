# OpenWeatherMap Go API

[![GoDoc](https://godoc.org/github.com/briandowns/openweathermap?status.svg)](https://godoc.org/github.com/briandowns/openweathermap) [![Build Status](https://travis-ci.org/briandowns/openweathermap.svg?branch=master)](https://travis-ci.org/briandowns/openweathermap) [![Coverage Status](https://coveralls.io/repos/github/briandowns/openweathermap/badge.svg?branch=master)](https://coveralls.io/github/briandowns/openweathermap?branch=master)

Go (golang) package for use with openweathermap.org's API.

For more detail about the library and its features, reference your local godoc once installed.

[Website](https://briandowns.github.io/openweathermap)!

To use the OpenweatherMap API, you need to obtain an API key.  Sign up [here](http://home.openweathermap.org/users/sign_up).  Once you have your key, create an environment variable called `OWM_API_KEY`.  Start coding!

[Slack Channel](https://openweathermapgolang.slack.com/messages/general)

Contributions welcome!

## Features

### Current Weather Conditions

- By City
- By City,St (State)
- By City,Co (Country)
- By City ID
- By Zip,Co (Country)
- By Longitude and Latitude

## Forecast

Get the weather conditions for a given number of days.

- By City
- By City,St (State)
- By City,Co (Country)
- By City ID
- By Longitude and Latitude

### Access to Condition Codes and Icons

Gain access to OpenWeatherMap icons and condition codes.

- Thunderstorms
- Drizzle
- Rain
- Snow
- Atmosphere
- Clouds
- Extreme
- Additional

### Data Available in Multiple Measurement Systems

- Fahrenheit (OpenWeatherMap API - imperial)
- Celsius (OpenWeatherMap API - metric)
- Kelvin (OpenWeatherMap API - internal)

### UV Index Data

- Current
- Historical

### Pollution Data

- Current

## Historical Conditions

- By Name
- By ID
- By Coordinates

## Supported Languages

English - en, Russian - ru, Italian - it, Spanish - es (or sp), Ukrainian - uk (or ua), German - de, Portuguese - pt, Romanian - ro, Polish - pl, Finnish - fi, Dutch - nl, French - fr, Bulgarian - bg, Swedish - sv (or se), Chinese Traditional - zh_tw, Chinese Simplified - zh (or zh_cn), Turkish - tr, Croatian - hr, Catalan - ca

## Installation

```bash
go get github.com/briandowns/openweathermap
```

## Examples

There are a few full examples in the examples directory that can be referenced.  1 is a command line application and 1 is a simple web application.

```Go
package main

import (
	"log"
	"fmt"
	"os"

	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OWM_API_KEY")

func main() {
	w, err := owm.NewCurrent("F", "ru", apiKey) // fahrenheit (imperial) with Russian output
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Phoenix")
	fmt.Println(w)
}

```

### Current Conditions by location name

```Go
func main() {
    w, err := owm.NewCurrent("K", "EN", apiKey) // (internal - OpenWeatherMap reference for kelvin) with English output
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByName("Phoenix,AZ")
    fmt.Println(w)
}
```

### Forecast Conditions in imperial (fahrenheit) by coordinates

```Go
func main() {
    w, err := owm.NewForecast("5", "F", "FI", apiKey) // valid options for first parameter are "5" and "16"
    if err != nil {
        log.Fatalln(err)
    }

    w.DailyByCoordinates(
        &owm.Coordinates{
                Longitude: -112.07,
                Latitude: 33.45,
        },
        5 // five days forecast
    )
    fmt.Println(w)
}
```

### Current conditions in metric (celsius) by location ID

```Go
func main() {
    w, err := owm.NewCurrent("C", "PL", apiKey)
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByID(2172797)
    fmt.Println(w)
}
```

### Current conditions by zip code. 2 character country code required

```Go
func main() {
    w, err := owm.NewCurrent("F", "EN", apiKey)
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByZip(19125, "US")
    fmt.Println(w)
}
```

### Configure http client

```Go
func main() {
    client := &http.Client{}
    w, err := owm.NewCurrent("F", "EN", apiKey, owm.WithHttpClient(client))
    if err != nil {
        log.Fatalln(err)
    }
}
```

### Current UV conditions

```Go
func main() {
    uv, err := owm.NewUV(apiKey)
    if err != nil {
        log.Fatalln(err)
    }

    coord := &owm.Coordinates{
        Longitude: 53.343497,
        Latitude:  -6.288379,
    }

    if err := uv.Current(coord); err != nil {
        log.Fatalln(err)
    }
    
    fmt.Println(coord)
}
```

### Historical UV conditions

```Go
func main() {
    uv, err := owm.NewUV(apiKey)
    if err != nil {
        log.Fatalln(err)
    }

    coord := &owm.Coordinates{
        Longitude: 54.995656,
        Latitude:  -7.326834,
    }

    end := time.Now().UTC()
    start := time.Now().UTC().Add(-time.Hour * time.Duration(24))

    if err := uv.Historical(coord, start, end); err != nil {
        log.Fatalln(err)
    }
}
```

### UV Information

```Go
func main() {
    uv, err := owm.NewUV(apiKey)
    if err != nil {
        log.Fatalln(err)
    }
    
    coord := &owm.Coordinates{
    	Longitude: 53.343497,
    	Latitude:  -6.288379,
    }
    
    if err := uv.Current(coord); err != nil {
    	log.Fatalln(err)
    }

    info, err := uv.UVInformation()
    if err != nil {
        log.Fatalln(err)
    }
    
    fmt.Println(info)
}
```

### Pollution Information

```Go
func main() {
    pollution, err := owm.NewPollution(apiKey)
    if err != nil {
        log.Fatalln(err)
    }

    params := &owm.PollutionParameters{
        Location: owm.Coordinates{
            Latitude:  0.0,
            Longitude: 10.0,
        },
        Datetime: "current",
    }

    if err := pollution.PollutionByParams(params); err != nil {
        log.Fatalln(err)
    }
}
```
