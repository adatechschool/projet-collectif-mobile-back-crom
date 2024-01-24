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
	router.HandleFunc("/api/spots", getSpots).Methods("GET")
    router.HandleFunc("/api/spots", putFavorites).Methods("PUT")
	// router.HandleFunc("/partialspots", getPartialSpots).Methods("GET")
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


func putFavorites(w http.ResponseWriter, r *http.Request) {
    spotID := mux.Vars(r)["id"]

    // Exécutez une requête SELECT pour obtenir la valeur actuelle de "favorites" pour le spot avec l'ID spécifié.
    row := db.QueryRow("SELECT favorites FROM spot WHERE id = ?", spotID)

    var currentFavorites bool
    err := row.Scan(&currentFavorites)
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Erreur lors de la récupération des favoris du spot", http.StatusInternalServerError)
        return
    }

    // Inversez la valeur de currentFavorites (basculer entre true et false)
    newFavorites := !currentFavorites

    // Exécutez une requête UPDATE pour mettre à jour la valeur de "favorites" pour le spot avec la nouvelle valeur calculée.
    _, err = db.Exec("UPDATE spot SET favorites = ? WHERE id = ?", newFavorites, spotID)
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Erreur lors de la mise à jour des favoris du spot", http.StatusInternalServerError)
        return
    }

    // Répondez avec un statut de succès si tout s'est bien passé.
    w.WriteHeader(http.StatusOK)
}

