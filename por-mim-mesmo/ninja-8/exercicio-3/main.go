package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	user1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service"},
	}

	user2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself."},
	}

	user3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?"},
	}

	users := []user{user1, user2, user3}

	//method chaining https://ttemporin.dev/como-fazer-encadeamento-de-metodos-chaining/
	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println("Erro na execução: ", err)
	}

}
