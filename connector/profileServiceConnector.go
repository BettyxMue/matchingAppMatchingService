package connector

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"app/matchingAppMatchingService/common/dataStructures"
)

func GetProfileById(id int) (*dataStructures.User, error) {
	var user dataStructures.User
	query := os.Getenv("PROFILE_SERVICE_HOST") + "/profile/" + strconv.Itoa(id)
	restClient := http.Client{
		Timeout: time.Second * 40,
	}

	request, errReq := http.NewRequest(http.MethodGet, query, nil)
	if errReq != nil {
		log.Println("Could not query user!")
		return nil, errReq
	}

	request.Header.Set("Authorization", "Bearer "+os.Getenv("JWT"))

	result, errRes := restClient.Do(request)
	if errRes != nil {
		log.Println("Could not query user!")
		return nil, errRes
	}
	if result.Body != nil {
		defer result.Body.Close()
	}
	body, errRead := ioutil.ReadAll(result.Body)
	if errRead != nil {
		log.Println("Could not read body")
		return nil, errRead
	}
	if errJson := json.Unmarshal(body, &user); errJson != nil {
		log.Println(errJson)
		return nil, errJson
	}
	return &user, nil
}

func GetAllProfiles() (*[]dataStructures.User, error) {
	var users []dataStructures.User
	query := os.Getenv("PROFILE_SERVICE_HOST") + "/profile/"
	restClient := http.Client{
		Timeout: time.Second * 40,
	}

	request, errReq := http.NewRequest(http.MethodGet, query, nil)
	if errReq != nil {
		log.Println("Could not query all users!")
		return nil, errReq
	}

	request.Header.Set("Authorization", "Bearer "+os.Getenv("JWT"))

	result, errRes := restClient.Do(request)
	if errRes != nil {
		log.Println("Could not query all users!")
		return nil, errRes
	}
	if result.Body != nil {
		defer result.Body.Close()
	}
	body, errRead := ioutil.ReadAll(result.Body)
	if errRead != nil {
		log.Println("Could not read body")
		return nil, errRead
	}
	if errJson := json.Unmarshal(body, &users); errJson != nil {
		log.Println(errJson)
		return nil, errJson
	}
	return &users, nil
}

func GetProfilesBySkill(skillId int) (*[]dataStructures.User, error) {
	var users []dataStructures.User
	query := os.Getenv("PROFILE_SERVICE_HOST") + "/skill/" + strconv.Itoa(skillId) + "/users"
	restClient := http.Client{
		Timeout: time.Second * 40,
	}

	request, errReq := http.NewRequest(http.MethodGet, query, nil)
	if errReq != nil {
		log.Println("Could not query searched users!")
		return nil, errReq
	}

	request.Header.Set("Authorization", "Bearer "+os.Getenv("JWT"))

	result, errRes := restClient.Do(request)
	if errRes != nil {
		log.Println("Could not query searched users!")
		return nil, errRes
	}
	if result.Body != nil {
		defer result.Body.Close()
	}
	body, errRead := ioutil.ReadAll(result.Body)
	if errRead != nil {
		log.Println("Could not read body")
		return nil, errRead
	}
	if errJson := json.Unmarshal(body, &users); errJson != nil {
		log.Println(errJson)
		return nil, errJson
	}
	return &users, nil
}
