package customer

import (
	"github.com/google/uuid"
)

type Customer struct {
	UUID       uuid.UUID `json:"id"`
	FirebaseId string    `json:"firebaseId"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	Tel        string    `json:"tel"`
	Address    string    `json:"address"`
	Num        string    `json:"num"`
	Cap        string    `json:"cap"`
	City       string    `json:"city"`
	State      string    `json:"state"`
}
