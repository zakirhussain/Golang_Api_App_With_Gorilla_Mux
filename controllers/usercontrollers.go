package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

type user struct {
	Id   string
	Name string
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	u := []user{}
	path := "/mnt/hgfs/workspace/user-management/controllers/db.txt"
	buf, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	for snl.Scan() {
		fmt.Println(snl.Text())
		u1 := strings.Split(snl.Text(), ",")
		u2 := user{}
		u2.Id = u1[0]
		u2.Name = u1[1]
		u = append(u, u2)
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(u)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	u := user{}
	path := "/mnt/hgfs/workspace/user-management/controllers/db.txt"
	buf, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// get url values
	URL_Values := mux.Vars(r)
	id := URL_Values["id"]

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	for snl.Scan() {
		fmt.Println(snl.Text())
		u1 := strings.Split(snl.Text(), ",")
		if id == u1[0] {
			u.Id = u1[0]
			u.Name = u1[1]
			break
		}
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(u)
}
