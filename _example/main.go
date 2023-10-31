package main

import (
	"fmt"

	"github.com/candy12t/go-json"
)

func main() {
	{
		input := `
{
  "Image": {
    "Width": 800,
    "Height": 600,
    "Title": "View from 15th Floor",
    "Thumbnail": {
      "Url": "http://www.example.com/image/481989943",
      "Height": 125,
      "Width": 100
    },
    "Animated": false,
    "IDs": [
      116,
      943,
      234,
      38793
    ]
  }
}
`
		result := json.NewParser(input).Parse()
		fmt.Println(result)
	}

	{
		input := `
[
  {
    "precision": "zip",
    "Latitude": 37.7668,
    "Longitude": -122.3959,
    "Address": "",
    "City": "SAN FRANCISCO",
    "State": "CA",
    "Zip": "94107",
    "Country": "US"
  },
  {
    "precision": "zip",
    "Latitude": 37.371991,
    "Longitude": -122.02602,
    "Address": "",
    "City": "SUNNYVALE",
    "State": "CA",
    "Zip": "94085",
    "Country": "US"
  }
]
`
		result := json.NewParser(input).Parse()
		fmt.Println(result)
	}
}
