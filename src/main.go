package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

/* GLOBALS */

var credentials map[string]string
var db map[string]string
var changes = 0
var dataDir = ".onion-kv"
var dbPath = fmt.Sprintf("%s/database", dataDir)
var crendentialsPath = fmt.Sprintf("%s/credentials.json", dataDir)

func createNewCredentials() {
	fmt.Println("Creating new credentials file")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("select root username: (default: root)")
	username, _ := reader.ReadString('\n')
	if username == "\n" {
		username = "root"
	} else {
		username = username[:len(username)-1]
	}
	fmt.Print("select root password: (default: dbadmin)")
	password, _ := reader.ReadString('\n')
	if password == "\n" {
		password = "dbadmin"
	} else {
		password = password[:len(password)-1]
	}

	fmt.Println("{\"" + username + "\":\"" + password + "\"}")
	os.WriteFile(crendentialsPath, []byte("{\""+username+"\":\""+password+"\"}"), 0644)
}

func initCredentials() {
	if info, err := os.Stat(crendentialsPath); os.IsNotExist(err) || info.IsDir() {
		createNewCredentials()
	}

	data, err := os.ReadFile(crendentialsPath)
	if err != nil {
		log.Fatal("Error reading credentials file:", err)
	}

	credentials = map[string]string{}
	err = json.Unmarshal(data, &credentials)
	if err != nil {
		log.Fatal("Error unmarshalling credentials:", err)
	}

	go heartBeat()
}

func heartBeat() {
	for range time.Tick(time.Second * 5) {
		if changes > 0 {
			go saveDb(db)
			changes = 0
		}
	}
}

func main() {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.Mkdir(dataDir, 0755)
	}

	initCredentials()

	if !existsDb() {
		log.Println("Creating new db")
		db = make(map[string]string)
		saveDb(db)
	} else {
		db = loadDb()
	}
	log.Println("Db loaded")
	openServer()
}
