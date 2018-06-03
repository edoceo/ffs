/**
	Main
*/

package main

import (
	"os"
//	"encoding/hex"
//	"encoding/json"
//	"flag"
	"fmt"
//	"io"
//	"io/ioutil"
//	"log"
//	"math/rand"
//	"net/http"
//	"net/url"
//	"path/filepath"
//	"sort"
//	"strconv"
//	"strings"
//	"time"
//	"unicode/utf8"
	"github.com/go-ini/ini"
//	"github.com/gorilla/securecookie"
	// "github.com/garyburd/go-oauth/oauth"
	//"github.com/kjk/u"
)

func main() {

	// Load Config
	Config, _ := ini.Load("ffs.ini")
	//fmt.Print(Config)
	//fmt.Printf("%+v\n", Config)

	setRootPath(Config)

	HTTP_init(Config)
}

func setRootPath(Config *ini.File) {

	// Set the root
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Config.Section("App").NewKey("root", pwd + "/webroot")

	// Get Root Directory
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    //if err != nil {
	//	log.Fatal(err)
    //}
    //fmt.Println(dir)

}
