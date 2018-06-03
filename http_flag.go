/**
	HTTP API for Flags
*/

package main

import (
	"log"
	"regexp"
	"strings"
	"encoding/json"
	"net/http"
)

type flag_list struct {
	// ?
	Flag []Flag_Record
}

// Return a list of Flags in JSON format
func HTTP_api_flag_select(w http.ResponseWriter, r *http.Request) {

	log.Printf("HTTP_api_flag_select(%s)", r.URL.Path)

	if !HTTP_auth(r) {
		//HTTP_bail(w)
	}

	res0 := DB_flag_list()

	res1 := flag_list{}

	// Foreach these to JSON
	for res0.Next() {

		flag := Flag_Record{}
		err := res0.Scan(&flag.Id, &flag.Stub, &flag.Name, &flag.Hash)
		if err != nil {
			log.Print("Error:", err)
			panic("Failed to Scan the Flags")
		}
		res1.Flag = append(res1.Flag, flag)
	}

	json, err := json.Marshal(res1)
	if err != nil {
		log.Print("Error:", err)
		panic("Failed to Get the Flags")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}

// Return a Flag
func HTTP_api_flag_single(w http.ResponseWriter, r *http.Request) {

	log.Println("HTTP_api_flag_handler(%s)", r.URL.Path)

	HTTP_auth(r)

	path := strings.ToLower(r.URL.Path)

	//re, _ := regexp.Compile("/flag/([0-9a-z][\\w\\-]+)(.*)?)")
	re, _ := regexp.Compile("/flag/([\\w\\-]+)(.*)?")
	match := re.FindStringSubmatch(path)

	var flag string

	switch len(match) {
	case 0:
		//w.Write([]byte("{ status: \"failure\" }"))
		// Resets my Header
		//http.Error(w, "{ status: \"failure\" }", http.StatusNotFound)
		//w.Header().Set("Content-Type", "application/json")
		return
		break
	case 1:

		flag = match[1]

		// Lookup in Database
		flag0 := DB_find_flag(flag)
		log.Println(flag0)

		resD := Flag_Record{}

		resB, _ := json.Marshal(resD)

		w.Header().Set("Content-Type", "application/json")
		w.Write(resB)

		break

	case 2:
	case 3:
		flag = match[1]
		// opts = match[1
		break
	default:
		log.Println(len(match));
		panic("Everything Sucks")
	}

	if 0 == len(flag) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{ status: \"failure\" }"))
		return
	}

	// Lookup in Database
	flag0 := DB_find_flag(flag)
	//log.Println(userO)

	//flagD := &Flag_Data{
	//	User: flagO,
	//	Flag_List: []string{ "manifest-v2" },
	//}

	//if (len(flag) > 0) {
	//	resD.Flag_List = append(resD.Flag_List, flag)
	//}

	resB, _ := json.Marshal(flag0)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resB)

	// Determine the URI User ID
	// Lookup in Database
	// Examine Blacklist
	// Examine Whitelist
	// Examine Greenlist

}
