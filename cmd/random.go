/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
<<<<<<< HEAD
	"math/rand"
	"net/http"
	"time"
=======
	"net/http"
>>>>>>> 64c08232ad36ad5867ce3b2e91eedcc83e9b6cae

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Gets a random dad joke",
	Long:  `Fetches a random joke from the icanhazdadjoke API`,
	Run: func(cmd *cobra.Command, args []string) {
<<<<<<< HEAD
		searchTerm, _ := cmd.Flags().GetString("term")

		if searchTerm != "" {
			getRandomJokeWithTerm(searchTerm)
		} else {
			getRandomJoke()
		}
=======
		getRandomJoke()
>>>>>>> 64c08232ad36ad5867ce3b2e91eedcc83e9b6cae
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
<<<<<<< HEAD

	randomCmd.PersistentFlags().String("term", "", "Search for dad jokes related to given term")
=======
>>>>>>> 64c08232ad36ad5867ce3b2e91eedcc83e9b6cae
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

<<<<<<< HEAD
type SearchTermResponse struct {
	Results    json.RawMessage `json:"results"`
	SearchTerm string          `json:"search_term"`
	Status     int             `json:"status"`
	TotalJokes int             `json:"total_jokes"`
}

=======
>>>>>>> 64c08232ad36ad5867ce3b2e91eedcc83e9b6cae
func getRandomJoke() {
	requestUrl := "https://icanhazdadjoke.com/"
	responseBytes := requestJoke(requestUrl)
	joke := Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
<<<<<<< HEAD
		log.Fatalf("Failed to unmarshal API response - %v", err)
=======
		log.Printf("Failed to unmarshal API response - %v", err)
>>>>>>> 64c08232ad36ad5867ce3b2e91eedcc83e9b6cae
	}

	fmt.Println(string(joke.Joke))
}

<<<<<<< HEAD
func getRandomJokeWithTerm(searchTerm string) {
	rand.Seed(time.Now().Unix())

	totalJokes, jokes := requestJokeWithTerm(searchTerm)
	min, max := 0, totalJokes-1

	if totalJokes <= 0 {
		fmt.Println("No jokes could be found with the search term")
	} else {
		randomIdx := min + rand.Intn(max-min)
		fmt.Println(jokes[randomIdx].Joke)
	}
}

func requestJokeWithTerm(searchTerm string) (totalJokes int, jokeList []Joke) {
	requestUrl := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", searchTerm)
	responseBytes := requestJoke(requestUrl)
	jokeListRaw := SearchTermResponse{}

	if err := json.Unmarshal(responseBytes, &jokeListRaw); err != nil {
		log.Fatalf("Failed to unmarshal API response - %v", err)
	}

	jokes := []Joke{}
	if err := json.Unmarshal(jokeListRaw.Results, &jokes); err != nil {
		log.Fatalf("Failed to unmarshal jokeListRaw results - %v", err)
	}

	return jokeListRaw.TotalJokes, jokes

}

=======
>>>>>>> 64c08232ad36ad5867ce3b2e91eedcc83e9b6cae
func requestJoke(requestUrl string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		requestUrl,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to create a request for joke- %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI Tool (github.com/KennethSC/dadjoke)")
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatalf("Request to get joke failed - %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Failed to read response body- %v", err)
	}

	return responseBytes

}
