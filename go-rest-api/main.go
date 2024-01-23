package main

import (
	"database/sql"
	"encoding/json"
	"go-rest-api/main/models"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"fmt"
	"os"
)

var db *sql.DB
var err error

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Establish the database connection
	db, err = connectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Set up and start the HTTP server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/spots", getSpots).Methods("GET")
	router.HandleFunc("/partialspots", getPartialSpots).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func connectDB() (*sql.DB, error) {
	// Get database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the database connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open the database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Check if the connection is valid
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	fmt.Println("Connected to the database")
	return db, nil
}

func getSpots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var spots []models.Spot
	result, err := db.Query("SELECT * FROM spots")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var spot models.Spot
		err := result.Scan(&spot.ID, &spot.Destination, &spot.PhotoURL, &spot.Address, &spot.Favorites)
		if err != nil {
			panic(err.Error())
		}
		spots = append(spots, spot)
	}
	json.NewEncoder(w).Encode(spots)
}

func getPartialSpots(w http.ResponseWriter, r *http.Request) {
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
                    "Geocode": ":grand_cercle_bleu: eyJpIjoiUGlwZWxpbmUsIE9haHUsIEhhd2FpaSIsIm8iOnsic3RhdHVzIjoiT0siLCJmb3JtYXR0ZWRBZGRyZXNzIjoiRWh1a2FpIEJlYWNoIFBhcmssIEhhbGVpd2EsIEhJIDk2NzEyLCBVbml...",
                    "Influencers": [
                        "recD1zp1pQYc8O7l2",
                        "rec1ptbRPxhS8rRun"
                    ],
                    "Magic Seaweed Link": "https://magicseaweed.com/Pipeline-Backdoor-Surf-Report/616/",
                    "Photos": [
                        {
                            "id": "attf6yu03NAtCuv5L",
                            "url": "https://www.lonelyplanet.fr/sites/lonelyplanet/files/styles/article_main_image/public/media/article/image/zachary-shea-mafuz8nh7xq-unsplash_1.jpg?itok=Baex1h1y",
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
                    "Geocode": ":grand_cercle_bleu: eyJpIjoiU2tlbGV0b24gQmF5LCBOYW1pYmlhIiwibyI6eyJzdGF0dXMiOiJPSyIsImZvcm1hdHRlZEFkZHJlc3MiOiJOYW1pYmlhIiwibGF0IjotMjUuOTE0NDkxOSwibG5nIjoxNC45MDY4NTk...",
                    "Influencers": [
                        "recD1zp1pQYc8O7l2",
                        "rec1ptbRPxhS8rRun"
                    ],
                    "Magic Seaweed Link": "https://magicseaweed.com/Skeleton-Bay-Surf-Report/4591/",
                    "Photos": [
                        {
                            "id": "attKinKKZvgdS5A5U",
                            "url": "https://www.voyageursdumonde.fr/voyage-sur-mesure/magazine-voyage/ShowPhoto/1490/0",
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
                    "Geocode": ":grand_cercle_bleu: eyJpIjoiU3VwZXJiYW5rLCBHb2xkIENvYXN0LCBBdXN0cmFsaWEiLCJvIjp7InN0YXR1cyI6Ik9LIiwiZm9ybWF0dGVkQWRkcmVzcyI6IlNuYXBwZXIgUm9ja3MgUmQsIENvb2xhbmdhdHRhIFF...",
                    "Influencers": [
                        "recSkJ4HuvzAUBrdd"
                    ],
                    "Magic Seaweed Link": "https://magicseaweed.com/Surfers-Paradise-Gold-Coast-Surf-Report/1012/",
                    "Photos": [
                        {
                            "id": "attmtbEOAQteRjz2p",
                            "url": "https://www.hostelworld.com/blog/wp-content/uploads/2018/02/best-surfing-spots-in-the-world-@wypeoutchef-Hawaii.jpg",
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

	var completeSpotList models.CompleteSpotList
	json.Unmarshal([]byte(myJsonString), &completeSpotList)

	// Crée une nouvelle structure pour stocker les informations réduites
	var partialSpotList models.PartialSpotList
	partialSpotList.Records = make([]models.PartialRecord, len(completeSpotList.Records))

	// Remplit la nouvelle structure avec des informations réduites
	for i, completeSpot := range completeSpotList.Records {
		partialSpotList.Records[i].Destination = completeSpot.Fields.Destination

		// Vérifie si la liste des photos n'est pas vide
		if len(completeSpot.Fields.Photos) > 0 {
			// Remplit le champ PhotoURL de PartialRecord avec l'URL de la première photo
			partialSpotList.Records[i].PhotoURL = completeSpot.Fields.Photos[0].URL
		}

		partialSpotList.Records[i].Address = completeSpot.Fields.Address
		// Ajoutez d'autres champs que vous souhaitez inclure
	}

	// Encode la nouvelle structure en JSON et renvoie la réponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partialSpotList)
}
