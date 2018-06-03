/**
	HTTP API for Users
*/

package main

import (
	"log"
	"regexp"
	"strings"
	"encoding/json"
	"net/http"
)

type User_Flag struct {
	Stub string `json:"stub"`
	Name string `json:"name"`
	Data string `json:"data"`
}

type User_Data struct {
	User User_Record
	Flag_List []User_Flag
}

type User_JSON struct {
	id string
	Hash string `json:"hash"`
	Name string `json:"name"`
}

// Return a list of Users in JSON format
func HTTP_api_user_select(w http.ResponseWriter, r *http.Request) {

	log.Println("HTTP_api_users_handler(%s)", r.URL.Path)

	//HTTP_auth(r)

	res0 := DB_user_list()

	user_list := []User_JSON{}

	for res0.Next() {

		u := User_JSON{}

		err := res0.Scan(&u.id, &u.Hash, &u.Name)
		if err != nil {
			log.Print("Error:", err)
			panic("Failed to Scan the Flags")
		}

		user_list = append(user_list, u)
	}

	// FOreach THese to JSON
	res1, err := json.Marshal(user_list)
		if err != nil {
			log.Print("Error:", err)
			panic("Failed to Marshal user_list")
		}
	//log.Print("Error:", err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res1)
	//log.Println(res0)

}

// Return a User
func HTTP_api_user_single(w http.ResponseWriter, r *http.Request) {

	log.Println("HTTP_api_user_handler(%s)", r.URL.Path)

	if !HTTP_auth(r) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("{ \"status\": \"failure\", \"detail\": \"Not Authorized\" }"))
		return
	}

	// Handle Errors
	defer func() {
		if r := recover(); r != nil {
			// w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{ \"status\": \"failure\", \"detail\": \"Trap\" }"))
		}
	}()


	path := strings.ToLower(r.URL.Path)

	re, _ := regexp.Compile("/user/([0-9a-f]+)(/?([0-9a-f]+)?)")
	match := re.FindStringSubmatch(path)
	//log.Println(match)

	var user string
	var flag string

	switch len(match) {
	case 2:
		user = match[1]
		flag = ""
	case 4:
		user = match[1]
		flag = match[3]
	default:
		panic("Everything Sucks")
	}

	if 0 == len(user) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{ \"status\": \"failure\", \"detail\": \"No User\" }"))
		return
	}

	// Lookup in Database
	userO := DB_user_pick(user)
	//log.Println(userO)

	retO := &User_Data{
		User: userO,
	}


	if 0 == len(flag) {

		resF := DB_user_flag_list(userO.Id)

		// Foreach these to JSON
		for resF.Next() {


			F := User_Flag{}
			err := resF.Scan(&F.Stub, &F.Name, &F.Data)
			if err != nil {
				log.Print("Error:", err)
				panic("Failed to Scan the Flags")
			}
			log.Println(F)
			retO.Flag_List = append(retO.Flag_List, F)
		}
	}

	// if (len(flag) > 0) {
	//	F := User_Flag{}
	//	F.stub = flag
	//	F.data = "zyxw"
	//	resD.Flag_List = append(resD.Flag_List, F)
	//}

	ret1, _ := json.Marshal(retO)

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret1)

}
