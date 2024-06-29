package helpers

import (
	"fmt"
	"html/template"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

// var templates = template.Must(template.New("index").Parse(string(GenrateTemplate())))

// var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	scores := Scores{PlayerOne: 0, PlayerTwo: 0}
	renderTemplate(w, "index", scores)
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	formData := r.FormValue("choice")
	if len(formData) <= 0 {
		http.Error(w, "What do you think you're doing?", http.StatusMethodNotAllowed)
		return
	}

	data := strings.Split(strings.ToLower(formData), ",")

	newScore := decideWinner(w, r, &data)
	renderTemplate(w, "index", newScore)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func decideWinner(w http.ResponseWriter, r *http.Request, data *[]string) Scores {
	gameSheet := map[string]string{
		"paper":    "rock",
		"scissors": "paper",
		"rock":     "scissors",
	}
	options := make([]string, 0, len(gameSheet))
	for k := range gameSheet {
		options = append(options, k)
	}

	playerChoice := (*data)[0]
	scores := calculateScore(w, r, data)
	randomChoiceIndex := rand.IntN(2-0) + 0

	println("computer choce " + options[randomChoiceIndex])
	println("player win condition " + gameSheet[playerChoice])

	if gameSheet[playerChoice] == options[randomChoiceIndex] {
		scores.PlayerOne += 1
	} else if gameSheet[options[randomChoiceIndex]] == playerChoice {
		scores.PlayerTwo += 1
	}

	println("newPlayerScore" + fmt.Sprint(scores.PlayerOne))
	println("newComScore" + fmt.Sprint(scores.PlayerTwo))
	println("--------------------")

	return scores
}

func calculateScore(w http.ResponseWriter, r *http.Request, data *[]string) Scores {
	playerScore, err := strconv.Atoi((*data)[1])
	if err != nil {
		http.Redirect(w, r, err.Error(), http.StatusTeapot)
	}
	computerScore, err := strconv.Atoi((*data)[2])
	if err != nil {
		http.Redirect(w, r, err.Error(), http.StatusTeapot)
	}

	scores := Scores{
		PlayerOne: playerScore,
		PlayerTwo: computerScore,
	}

	return scores
}

// func parseJsonString(jsonString string) []byte {
// 	// jsonString := `{"test":"test","ajsdbssad":"j","asdkjkabshd":1,"bool":true,"aljsdbahsdba":[true,1,"asdahbds"]}`

// 	var data map[string]any

// 	// Unmarshal JSON string into the map
// 	err := json.Unmarshal([]byte(jsonString), &data)
// 	if err != nil {
// 		fmt.Println("Error parsing JSON:", err)
// 		panic(err)
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println("Error marshalling JSON:", err)
// 		panic(err)
// 	}

// 	return jsonData
// }

// func returnResponse(w http.ResponseWriter, contentType string, response []byte) {
// 	w.Header().Set("Content-Type", contentType)
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(response)
// }
