package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) handleAccess(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handlePostAccess(w, r)
	} else if r.Method == "GET" {
		return s.handleGetAccess(w, r)
		//} else if r.Method == "DELETE" {
		//	return s.handleDeleteEmployee(w, r)
		//} else if r.Method == "PUT" {
		//	return s.handleUpdateEmployee(w, r)
	}
	return fmt.Errorf("method doesn't exist: %s", r.Method)
}

func (s *Server) handlePostAccess(w http.ResponseWriter, r *http.Request) error {
	var accessResponse AccessResponse
	var accessRequest = new(AccessRequest)
	if err := json.NewDecoder(r.Body).Decode(&accessRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	profileIds, err := getProfileIdsArray(w, accessRequest.ProfileIds)
	if err != nil {
		return err
	}
	for _, prof := range profileIds {
		log.Println(prof)
		checkAccessOfProfile(prof)
	}
	return WriteJson(w, http.StatusOK, accessResponse)
}

func getProfileIdsArray(w http.ResponseWriter, reqProfileIds interface{}) ([]int64, error) {
	var profileIds []int64
	typeOfProfileIds := getType(reqProfileIds)
	if typeOfProfileIds == "string" {
		profileId, err := strconv.ParseInt(reqProfileIds.(string), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return profileIds, err
		}
		profileIds = append(profileIds, profileId)
	} else if typeOfProfileIds == "[]float64" {
		for _, prof := range reqProfileIds.([]interface{}) {
			profileIds = append(profileIds, int64(prof.(float64)))
		}
	} else if typeOfProfileIds == "[]string" {
		for _, prof := range reqProfileIds.([]interface{}) {
			profileId, err := strconv.ParseInt(prof.(string), 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return profileIds, err
			}
			profileIds = append(profileIds, profileId)
		}
	} else if typeOfProfileIds == "float64" {
		profileIds = append(profileIds, int64(reqProfileIds.(float64)))
	}
	return profileIds, nil
}

func getType(v interface{}) string {
	switch x := v.(type) {
	case string:
		return "string"
	case float64:
		return "float64"
	case []interface{}:
		if len(x) > 0 {
			switch x[0].(type) {
			case string:
				return "[]string"
			case float64:
				return "[]float64"
			default:
				return "not found"
			}
		}
	default:
		return "not found"
	}
	return "not found"
}

func (s *Server) handleGetAccess(w http.ResponseWriter, r *http.Request) error {
	var accessResponse AccessResponse
	prof := mux.Vars(r)["prof_id"]
	profileId, err := strconv.ParseInt(prof, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	checkAccessOfProfile(profileId)
	return WriteJson(w, http.StatusOK, accessResponse)
}

func checkAccessOfProfile(_ int64) bool {
	return false
}
