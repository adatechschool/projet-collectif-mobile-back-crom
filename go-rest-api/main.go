package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type spot struct {
		ID     string `json:"id"`
		Fields struct {
			SurfBreak        []string `json:"Surf Break"`
			DifficultyLevel  int      `json:"Difficulty Level"`
			Destination      string   `json:"Destination"`
			Geocode          string   `json:"Geocode"`
			Influencers      []string `json:"Influencers"`
			MagicSeaweedLink string   `json:"Magic Seaweed Link"`
			Photos           []struct {
				ID         string `json:"id"`
				URL        string `json:"url"`
				Filename   string `json:"filename"`
				Size       int    `json:"size"`
				Type       string `json:"type"`
				Thumbnails struct {
					Small struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"small"`
					Large struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"large"`
					Full struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"full"`
				} `json:"thumbnails"`
			} `json:"Photos"`
			PeakSurfSeasonBegins    string `json:"Peak Surf Season Begins"`
			DestinationStateCountry string `json:"Destination State/Country"`
			PeakSurfSeasonEnds      string `json:"Peak Surf Season Ends"`
			Address                 string `json:"Address"`
		} `json:"fields"`
}

type allSpots []spot

var spots = allSpots{
	{
			{
				id: "rec5aF9TjMjBicXCK",
				fields: {
					SurfBreak: [
						"Reef Break"
					],
					DifficultyLevel: 4,
					Destination: "Pipeline",
					Geocode: "🔵 eyJpIjoiUGlwZWxpbmUsIE9haHUsIEhhd2FpaSIsIm8iOnsic3RhdHVzIjoiT0siLCJmb3JtYXR0ZWRBZGRyZXNzIjoiRWh1a2FpIEJlYWNoIFBhcmssIEhhbGVpd2EsIEhJIDk2NzEyLCBVbml...",
					Influencers: [
						"recD1zp1pQYc8O7l2",
						"rec1ptbRPxhS8rRun"
					],
					Magic Seaweed Link: "https://magicseaweed.com/Pipeline-Backdoor-Surf-Report/616/",
					Photos: [
						{
							id: "attf6yu03NAtCuv5L",
							url: "https://dl.airtable.com/ZuXJZ2NnTF40kCdBfTld_thomas-ashlock-64485-unsplash.jpg",
							filename: "thomas-ashlock-64485-unsplash.jpg",
							size: 688397,
							type: "image/jpeg",
							thumbnails: {
								small: {
									url: "https://dl.airtable.com/yfKxR9ZQqiT7drKxpjdF_small_thomas-ashlock-64485-unsplash.jpg",
									width: 52,
									height: 36
								},
								large: {
									url: "https://dl.airtable.com/cFfMuU8NQjaEskeC3B2h_large_thomas-ashlock-64485-unsplash.jpg",
									width: 744,
									height: 512
								},
								full: {
									url: "https://dl.airtable.com/psynuQNmSvOTe3BWa0Fw_full_thomas-ashlock-64485-unsplash.jpg",
									width: 2233,
									height: 1536
								}
							}
						}
					],
					Peak Surf Season Begins: "2018-07-22",
					Destination State/Country: "Oahu, Hawaii",
					Peak Surf Season Ends: "2018-08-31",
					Address: "Pipeline, Oahu, Hawaii"
				},
			}
	}	
}

// type event struct {
// 	ID          string `json:"ID"`
// 	Title       string `json:"Title"`
// 	Description string `json:"Description"`
// }

// type allEvents []event

// var events = allEvents{
// 	{
// 		ID:          "1",
// 		Title:       "Introduction to Golang",
// 		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
// 	},
// }

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}
