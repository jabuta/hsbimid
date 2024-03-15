package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Client struct {
}

type apiConfig struct {
	hsapi        string
	hsapiversion string
}

func main() {
	godotenv.Load()
	cfg := apiConfig{
		hsapi:        os.Getenv("HS-API-KEY"),
		hsapiversion: "v3",
	}
	endpoint := fmt.Sprintf("https://api.hubapi.com/crm/%s/objects/contacts", cfg.hsapiversion)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatalf("Could not create request, error: %s", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.hsapi))
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Could not make request, error: %s", err)
	}
	defer resp.Body.Close()

	var body Client
	json.NewDecoder(resp.Body).Decode(&body)

	fmt.Print(body)

}
