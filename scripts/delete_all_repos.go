package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Participant struct {
	IntraLogin string `json:"intra_login"`
}

type ShortConfig struct {
	Participants []Participant `json:"participants"`
}

func deleteRepo(repoID string) bool {
	org := os.Getenv("GITHUB_ORGANISATION")
	token := os.Getenv("GITHUB_TOKEN")
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", org, repoID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		log.Printf("Successfully deleted repo %s", repoID)
		return true
	} else {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Failed to delete repo %s: %d %s", repoID, resp.StatusCode, string(bodyBytes))
		return false
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	file, err := os.Open("shortconfig/shortconfig.json")
	if err != nil {
		log.Fatalf("Error opening shortconfig.json: %v", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var shortConfig ShortConfig
	if err := json.Unmarshal(byteValue, &shortConfig); err != nil {
		log.Fatalf("Error parsing shortconfig.json: %v", err)
	}

	for _, participant := range shortConfig.Participants {
		deleteRepo(fmt.Sprintf("%s-00", participant.IntraLogin))
	}
}
