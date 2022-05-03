package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func authenticate(r *http.Request) bool {
	passwd := []byte(r.URL.Query().Get("ddns-web-passwd"))
	hasher := sha256.New()
	hasher.Write(passwd)
	hashed_passwd := hex.EncodeToString(hasher.Sum(nil))
	real_passwd_hash, err := os.ReadFile("passwd")
	if err != nil {
		log.Fatalln("Failed to read passwd: ", err)
	}
	return hashed_passwd == string(real_passwd_hash)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	auth := authenticate(r)
	new_ip := r.RemoteAddr
	if !auth {
		log.Println(new_ip, "not authenticated")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	log.Println("Change IP on DDNS to ", new_ip)
	cmd := exec.Command("sh", "ddns-update", new_ip)
	err := cmd.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("Couldn't update server: ", err)
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	logfile, _ := os.OpenFile("logs", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Println("Starting Server")

	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServeTLS("127.0.0.1:8143", "/secrets/server.cer", "/secrets/server.key", nil))
}
