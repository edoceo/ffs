/**
	HTTP Handlers for FFS
*/

package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"html/template"
	"path/filepath"
	"github.com/go-ini/ini"
)

var app_cfg *ini.File

func HTTP_init(cfg *ini.File) {

	fmt.Printf("HTTP_init()\n")

	app_cfg = cfg

	bind := cfg.Section("HTTP").Key("bind").String()
	fmt.Printf("Binding: %s\n", bind)

	root := cfg.Section("App").Key("root").String()
	fmt.Printf("Webroot: %s\n", root)

	// http.HandleFunc("/", http_page_index)
	http.HandleFunc("/api/v2016/flags", HTTP_api_flag_select);
	http.HandleFunc("/api/v2016/flag/", HTTP_api_flag_single);

	http.HandleFunc("/api/v2016/users", HTTP_api_user_select);
	http.HandleFunc("/api/v2016/user/", HTTP_api_user_single);

	//http.HandleFunc("/flag/", HTTP_flag_single);
	//http.HandleFunc("/user/", HTTP_user_single);

	// I'm not sure why I have to handle this way
	fs := http.FileServer(http.Dir(filepath.Join("webroot", "tag")))
	http.Handle("/tag/", http.StripPrefix("/tag/", fs))

	//http.Handle("/js", http.FileServer(http.Dir(root)))
	http.HandleFunc("/", HTTP_file_handler);
	//http.Handle("/", fs)

	if err := http.ListenAndServe(bind, nil); err != nil {
		fmt.Printf("http.ListendAndServe() failed with %s\n", err)
	}

}

//func serveFileFromDir(w http.ResponseWriter, r *http.Request, dir, fileName string) {
//	if redirectIfFoundMatching(w, r, dir, fileName) {
//		return
//	}
//	filePath := filepath.Join(dir, fileName)
//	if u.PathExists(filePath) {
//		//logger.Noticef("serveFileFromDir(): %q", filePath)
//		http.ServeFile(w, r, filePath)
//	} else {
//		logger.Noticef("serveFileFromDir() file %q doesn't exist, referer: %q", fileName, getReferer(r))
//		http.NotFound(w, r)
//	}
//}

func HTTP_auth(r *http.Request) bool {

	req_user, req_pass, _ := r.BasicAuth()

	fmt.Printf("HTTP_auth(%s, %s)\n", req_user, req_pass)

	app_user := app_cfg.Section("HTTP").Key("username").String();
	app_pass := app_cfg.Section("HTTP").Key("password").String();

	if app_user == req_user {
		if app_pass == req_pass {
			return true;
		}
	}

	return false;

}

func HTTP_file_handler(w http.ResponseWriter, r *http.Request) {

	// Layout Path
	lp := filepath.Join("layout", "html.html")

	// View Path
	fp := filepath.Clean(r.URL.Path)
	if "/" == fp {
		fp = "/index.html"
	}

	// Local Static
	fp0 := filepath.Join("webroot", fp)

	full_path, err := filepath.Abs(fp0)
	log.Print(full_path)
	log.Print(err)
	if err != nil {
		log.Print(err)
		return;
	}

	chk, err := os.Stat(full_path)
	log.Print(chk);
	if err == nil {
		fmt.Printf("HTTP_file_handler(%s)\n", full_path)
		http.ServeFile(w, r, full_path)
		return
	}

	//fp = r.URL.Path
	fp = filepath.Join("view", fp)

	fmt.Printf("HTTP_file_handler(%s)\n", fp)

	//tmpl, err := template.New("pages").ParseFiles(root)
	//check(err)

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", nil)

}
