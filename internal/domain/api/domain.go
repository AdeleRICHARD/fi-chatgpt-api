package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/figarocms/fi-chatgpt-api/infrastucture/chat"
	"github.com/figarocms/fi-chatgpt-api/internal/domain/model"
	"github.com/figarocms/fi-chatgpt-api/tool"
)

func Question(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("../annonces_list.json")
	if err != nil {
		panic(err)
	}

	var classifiedList model.ClassifiedsList
	err = json.Unmarshal(data, &classifiedList)
	if err != nil {
		log.Fatal(err)
	}

	var classifiedDesc model.Classifieds
	var classifieds []model.Classifieds
	questionStart := "Extrait les atouts de l'annonce immobilère suivante avec les points les plus utiles dans une liste de points clés numérotée : "

	for i, hit := range classifiedList.Hits.Hits {
		if i == 10000 {
			break
		}
		err := json.Unmarshal(hit.Source, &classifiedDesc)
		if err != nil {
			log.Fatal(err)
		}

		// First classified = marseillan
		if i == 0 {
			continue
		} else {
			classifiedDesc.Property.Description = questionStart + classifiedDesc.Property.Description
			classifieds = append(classifieds, classifiedDesc)
		}
	}

	fmt.Print("Classified ", classifieds)
	var answers []string
	for _, classified := range classifieds {
		question := classified.Property.Description
		resp, err := chat.AskToChatGPT(question)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		answers = append(answers, strings.ReplaceAll(resp, "\n", ""))
	}

	assets := getAssets(answers)
	tool.Export_chatGPT_answer_csv(assets)
}

func getAssets(answers []string) []string {
	var assets []string
	for _, answer := range answers {
		re := regexp.MustCompile(`\d\.`)
		assets = append(assets, re.Split(answer, -1)...)
	}
	return assets
}
