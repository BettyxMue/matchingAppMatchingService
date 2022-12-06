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

	request.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzAzNDk5NTgsInN1YiI6MSwidXNlciI6MX0.EFSj4-Aj95t9O8kWzgtDz2odJ5OiA2zrvHuJsiJkw0_U5w8IqIF6z3z_mLeR2uKVqfHl8XtELs0BGs3JuaANRvoSi1nviwf58oKuF7AwyY2DXT0cdtGVmUiMzi0CWg9BumjRsyL0M42oJV25sGpzwgWctk34yvNz0ScS0hBzvrhx2rSVHW3rJtRDevMp_UG9kZRDMPTKX9ax2jv_43FCFRtdcLdPO-CYJMQHhgMAZKO5nwAVqtOOtWXohSDrUPnSPgqOkbB8mOls6uckHoEgLFAUIJsxTckyh5Xt6_enyZ68W3ggsHDGPu0irvRfOrTYrB4fbaTMOQFZXqi4IJC_Xg")

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

	request.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzAzNDk5NTgsInN1YiI6MSwidXNlciI6MX0.EFSj4-Aj95t9O8kWzgtDz2odJ5OiA2zrvHuJsiJkw0_U5w8IqIF6z3z_mLeR2uKVqfHl8XtELs0BGs3JuaANRvoSi1nviwf58oKuF7AwyY2DXT0cdtGVmUiMzi0CWg9BumjRsyL0M42oJV25sGpzwgWctk34yvNz0ScS0hBzvrhx2rSVHW3rJtRDevMp_UG9kZRDMPTKX9ax2jv_43FCFRtdcLdPO-CYJMQHhgMAZKO5nwAVqtOOtWXohSDrUPnSPgqOkbB8mOls6uckHoEgLFAUIJsxTckyh5Xt6_enyZ68W3ggsHDGPu0irvRfOrTYrB4fbaTMOQFZXqi4IJC_Xg")

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

	request.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzAzNDk5NTgsInN1YiI6MSwidXNlciI6MX0.EFSj4-Aj95t9O8kWzgtDz2odJ5OiA2zrvHuJsiJkw0_U5w8IqIF6z3z_mLeR2uKVqfHl8XtELs0BGs3JuaANRvoSi1nviwf58oKuF7AwyY2DXT0cdtGVmUiMzi0CWg9BumjRsyL0M42oJV25sGpzwgWctk34yvNz0ScS0hBzvrhx2rSVHW3rJtRDevMp_UG9kZRDMPTKX9ax2jv_43FCFRtdcLdPO-CYJMQHhgMAZKO5nwAVqtOOtWXohSDrUPnSPgqOkbB8mOls6uckHoEgLFAUIJsxTckyh5Xt6_enyZ68W3ggsHDGPu0irvRfOrTYrB4fbaTMOQFZXqi4IJC_Xg")

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
