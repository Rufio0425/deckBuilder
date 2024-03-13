package handlers

import (
	"deckBuilder/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h Handler) GetCardHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cardName := params["name"]
	trueName := strings.ToLower(strings.Replace(cardName, "+", " ", -1))

	query := "SELECT * FROM cards.cards WHERE LOWER(name) = $1;"

	// Query for card first
	var card models.Card
	h.DB.QueryRow(query, trueName).Scan(
		&card.ID,
		&card.Name,
		&card.FlavorName,
		&card.CardDescription,
		&card.ManaCost,
		&card.ImageUris,
	)

	// If card is not found in DB, populate card from scryfall
	if card.ID == uuid.Nil {
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

		err = json.Unmarshal(body, &card)
		if err != nil {
			fmt.Printf("Error unmarshalling card: %s", err)
			return
		}

		// Save card to DB
		insertCard := "INSERT INTO cards.cards(" +
			"id," +
			"name," +
			"flavor_name," +
			"card_description," +
			"mana_cost" +
			") VALUES(" +
			"$1," +
			"$2," +
			"$3," +
			"$4," +
			"$5" +
			");"
		if _, err := h.DB.Exec(insertCard, &card.ID, &card.Name, &card.FlavorName, &card.CardDescription, &card.ManaCost); err != nil {
			w.WriteHeader(500)
			log.Fatalf("Error saving card: %s", err)
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(card)
}
