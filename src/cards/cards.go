package cards

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type card struct {
	ScryfallID      uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	FlavorName      string     `json:"flavor_name"`
	CardDescription string     `json:"oracle_text"`
	ManaCost        string     `json:"mana_cost"`
	Legalities      Legalities `json:"legalities"`
}

func GetCardHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cardName := params["name"]

	url := fmt.Sprintf("https://api.scryfall.com/cards/named?fuzzy=%s", cardName)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error retrieving card: %s", err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s", err)
		return
	}

	var card card
	err = json.Unmarshal(body, &card)
	if err != nil {
		fmt.Printf("Error unmarshalling card: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err != json.NewEncoder(w).Encode(card) {
		fmt.Printf("Error encoding JSON: %s", err)
		return
	}
}
