package main

import (
	"encoding/json"
	"fmt"
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
	// router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/spots", getAllSpots).Methods("GET")
	// router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type SpotList struct {
	Records []Record `json:"records"`
	Offset  string   `json:"offset"` 
}

type Record struct {
	ID          string `json:"id"`         
	Fields      Fields `json:"fields"`     
	CreatedTime string `json:"createdTime"`
}

type Fields struct {
	SurfBreak               []string `json:"Surf Break"`               
	DifficultyLevel         int64    `json:"Difficulty Level"`         
	Destination             string   `json:"Destination"`              
	Geocode                 string   `json:"Geocode"`                  
	Influencers             []string `json:"Influencers"`              
	MagicSeaweedLink        string   `json:"Magic Seaweed Link"`       
	Photos                  []Photo  `json:"Photos"`                   
	PeakSurfSeasonBegins    string   `json:"Peak Surf Season Begins"`  
	DestinationStateCountry string   `json:"Destination State/Country"`
	PeakSurfSeasonEnds      string   `json:"Peak Surf Season Ends"`    
	Address                 string   `json:"Address"`                  
}

type Photo struct {
	ID         string     `json:"id"`        
	URL        string     `json:"url"`       
	Filename   string     `json:"filename"`  
	Size       int64      `json:"size"`      
	Type       string     `json:"type"`      
	Thumbnails Thumbnails `json:"thumbnails"`
}

type Thumbnails struct {
	Small Full `json:"small"`
	Large Full `json:"large"`
	Full  Full `json:"full"` 
}

type Full struct {
	URL    string `json:"url"`   
	Width  int64  `json:"width"` 
	Height int64  `json:"height"`
}


// func createEvent(w http.ResponseWriter, r *http.Request) {
// 	var newEvent event
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}

// 	json.Unmarshal(reqBody, &newEvent)
// 	events = append(events, newEvent)
// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(newEvent)
// }

func getAllSpots(w http.ResponseWriter, r *http.Request) {
	myJsonString := `{
		"records": [
			{
				"id": "rec5aF9TjMjBicXCK",
				"fields": {
					"Surf Break": [
						"Reef Break"
					],
					"Difficulty Level": 4,
					"Destination": "Pipeline",
					"Geocode": "🔵 eyJpIjoiUGlwZWxpbmUsIE9haHUsIEhhd2FpaSIsIm8iOnsic3RhdHVzIjoiT0siLCJmb3JtYXR0ZWRBZGRyZXNzIjoiRWh1a2FpIEJlYWNoIFBhcmssIEhhbGVpd2EsIEhJIDk2NzEyLCBVbml...",
					"Influencers": [
						"recD1zp1pQYc8O7l2",
						"rec1ptbRPxhS8rRun"
					],
					"Magic Seaweed Link": "https://magicseaweed.com/Pipeline-Backdoor-Surf-Report/616/",
					"Photos": [
						{
							"id": "attf6yu03NAtCuv5L",
							"url": "https://dl.airtable.com/ZuXJZ2NnTF40kCdBfTld_thomas-ashlock-64485-unsplash.jpg",
							"filename": "thomas-ashlock-64485-unsplash.jpg",
							"size": 688397,
							"type": "image/jpeg",
							"thumbnails": {
								"small": {
									"url": "https://dl.airtable.com/yfKxR9ZQqiT7drKxpjdF_small_thomas-ashlock-64485-unsplash.jpg",
									"width": 52,
									"height": 36
								},
								"large": {
									"url": "https://dl.airtable.com/cFfMuU8NQjaEskeC3B2h_large_thomas-ashlock-64485-unsplash.jpg",
									"width": 744,
									"height": 512
								},
								"full": {
									"url": "https://dl.airtable.com/psynuQNmSvOTe3BWa0Fw_full_thomas-ashlock-64485-unsplash.jpg",
									"width": 2233,
									"height": 1536
								}
							}
						}
					],
					"Peak Surf Season Begins": "2018-07-22",
					"Destination State/Country": "Oahu, Hawaii",
					"Peak Surf Season Ends": "2018-08-31",
					"Address": "Pipeline, Oahu, Hawaii"
				},
				"createdTime": "2018-05-31T00:16:16.000Z"
			},
			{
				"id": "recT98Z2El7YYwmc4",
				"fields": {
					"Surf Break": [
						"Point Break"
					],
					"Difficulty Level": 5,
					"Destination": "Skeleton Bay",
					"Geocode": "🔵 eyJpIjoiU2tlbGV0b24gQmF5LCBOYW1pYmlhIiwibyI6eyJzdGF0dXMiOiJPSyIsImZvcm1hdHRlZEFkZHJlc3MiOiJOYW1pYmlhIiwibGF0IjotMjUuOTE0NDkxOSwibG5nIjoxNC45MDY4NTk...",
					"Influencers": [
						"recD1zp1pQYc8O7l2",
						"rec1ptbRPxhS8rRun"
					],
					"Magic Seaweed Link": "https://magicseaweed.com/Skeleton-Bay-Surf-Report/4591/",
					"Photos": [
						{
							"id": "attKinKKZvgdS5A5U",
							"url": "https://dl.airtable.com/YzqA020RRLaTyAZAta9g_brandon-compagne-308937-unsplash.jpg",
							"filename": "brandon-compagne-308937-unsplash.jpg",
							"size": 1494974,
							"type": "image/jpeg",
							"thumbnails": {
								"small": {
									"url": "https://dl.airtable.com/Y8970kyrQHWfL6AMkxZQ_small_brandon-compagne-308937-unsplash.jpg",
									"width": 54,
									"height": 36
								},
								"large": {
									"url": "https://dl.airtable.com/dkQKXoUnTGiofIvg5TJR_large_brandon-compagne-308937-unsplash.jpg",
									"width": 768,
									"height": 512
								},
								"full": {
									"url": "https://dl.airtable.com/pexuxaQ6D2VV61pyhUwn_full_brandon-compagne-308937-unsplash.jpg",
									"width": 3000,
									"height": 2000
								}
							}
						}
					],
					"Peak Surf Season Begins": "2018-09-01",
					"Destination State/Country": "Namibia ",
					"Peak Surf Season Ends": "2018-11-30",
					"Address": "Skeleton Bay, Namibia "
				},
				"createdTime": "2018-05-31T00:16:16.000Z"
			},
			{
				"id": "recH2ennHFNOtB1Wt",
				"fields": {
					"Surf Break": [
						"Point Break"
					],
					"Difficulty Level": 4,
					"Destination": "Superbank",
					"Geocode": "🔵 eyJpIjoiU3VwZXJiYW5rLCBHb2xkIENvYXN0LCBBdXN0cmFsaWEiLCJvIjp7InN0YXR1cyI6Ik9LIiwiZm9ybWF0dGVkQWRkcmVzcyI6IlNuYXBwZXIgUm9ja3MgUmQsIENvb2xhbmdhdHRhIFF...",
					"Influencers": [
						"recSkJ4HuvzAUBrdd"
					],
					"Magic Seaweed Link": "https://magicseaweed.com/Surfers-Paradise-Gold-Coast-Surf-Report/1012/",
					"Photos": [
						{
							"id": "attmtbEOAQteRjz2p",
							"url": "https://dl.airtable.com/I4E4xZeQbO2g814udQDm_jeremy-bishop-80371-unsplash.jpg",
							"filename": "jeremy-bishop-80371-unsplash.jpg",
							"size": 1524876,
							"type": "image/jpeg",
							"thumbnails": {
								"small": {
									"url": "https://dl.airtable.com/IWP9RPvvSM2pX1sHeigV_small_jeremy-bishop-80371-unsplash.jpg",
									"width": 48,
									"height": 36
								},
								"large": {
									"url": "https://dl.airtable.com/aBnINo8qQqDvER2f2wGg_large_jeremy-bishop-80371-unsplash.jpg",
									"width": 683,
									"height": 512
								},
								"full": {
									"url": "https://dl.airtable.com/eZxgLD8hRI2Y27J39LNl_full_jeremy-bishop-80371-unsplash.jpg",
									"width": 3000,
									"height": 2250
								}
							}
						}
					],
					"Peak Surf Season Begins": "2018-11-28",
					"Destination State/Country": "Gold Coast, Australia",
					"Peak Surf Season Ends": "2019-02-01",
					"Address": "Superbank, Gold Coast, Australia"
				},
				"createdTime": "2018-05-31T00:16:16.000Z"
			}
		],
		"offset": "recH2ennHFNOtB1Wt"
	}
	`
	var spots SpotList
	json.Unmarshal([]byte(myJsonString), &spots)
	json.NewEncoder(w).Encode(spots)
}

// func getOneEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]

// 	for _, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			json.NewEncoder(w).Encode(singleEvent)
// 		}
// 	}
// }
