package connector

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"app/matchingAppMatchingService/common/dataStructures"
)

var token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzE4MzU1NDUsInN1YiI6MiwidXNlciI6Mn0.sTVRUgPZl04VMSbcGvjRSacucCiVOQ4iYU_Nx4a_IE3jy6JtrXXCOOMjeLaxAWbNho5DFLxjKDFf05JVgTB9Mo8lkeGeogDHfumcz3yBnRv0cOXfTjuATGULF8vyM8sjTkkD3O9hYiK568UBJEFE8geY2q_k-3ONTZLv2ysZt2Y"

func GetProfileById(id int) (*dataStructures.User, error) {
	var user dataStructures.User
	query := "http://0.0.0.0:8080/profile/" + strconv.Itoa(id)
	restClient := http.Client{
		Timeout: time.Second * 40,
	}

	request, errReq := http.NewRequest(http.MethodGet, query, nil)
	if errReq != nil {
		log.Println("Could not query user!")
		return nil, errReq
	}

	request.Header.Set("Authorization", "Bearer "+token)

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
	query := "http://0.0.0.0:8080/profile/"
	restClient := http.Client{
		Timeout: time.Second * 40,
	}

	request, errReq := http.NewRequest(http.MethodGet, query, nil)
	if errReq != nil {
		log.Println("Could not query all users!")
		return nil, errReq
	}

	request.Header.Set("Authorization", "Bearer "+token)

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
	query := "http://0.0.0.0:8080/skill/" + strconv.Itoa(skillId) + "/users"
	restClient := http.Client{
		Timeout: time.Second * 40,
	}

	request, errReq := http.NewRequest(http.MethodGet, query, nil)
	if errReq != nil {
		log.Println("Could not query searched users!")
		return nil, errReq
	}

	request.Header.Set("Authorization", "Bearer "+token)

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
