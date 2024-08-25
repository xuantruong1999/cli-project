/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type Weather struct {
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		TempC    float64 `json:"temp"`
		FeelLike float64 `json:"feel_like"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
	} `json:"wind"`
	TimeZone int    `json:"timezone"`
	Name     string `json:"name"`
	Sys      struct {
		Country string `json:"country"`
	}
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get-weather",
	Short: "Get weather of HO CHI MINH city now",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Ho%20Chi%20Minh%20City,VN&APPID=538d02562e21a30b2cbd1ee5b9d0736b")
		if err != nil {
			panic("service is not available")
		}

		defer res.Body.Close()
		body, e := io.ReadAll(res.Body)
		if e != nil {
			panic(e)
		}

		var weather Weather
		e = json.Unmarshal(body, &weather)
		if e != nil {
			panic(e)
		}

		wea, main, wind, sys, name := weather.Weather[0], weather.Main, weather.Wind, weather.Sys, weather.Name
		fmt.Printf("Name: %s - %s, Main: %s, Description: %s, TempC: %.2f, Wind: %.2f \n",
			name, sys.Country, wea.Main, wea.Description, main.TempC, wind.Speed)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
