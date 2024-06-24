package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	addr := flag.String("addr", ":8888", "http server address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/hash", handleHash)
	mux.HandleFunc("/verify", handleVerify)

	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func handleHash(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	raw := r.Form.Get("raw")
	cost := r.Form.Get("cost")
	if raw == "" || cost == "" {
		http.Error(w, "Missing params", 400)
		return
	}

	cost0, err := strconv.Atoi(cost)
	if err != nil {
		http.Error(w, "Invalid cost", 400)
		return
	}

	if cost0 > 16 || cost0 < 5 {
		http.Error(w, "Cost should be in the range from 5 to 16", 400)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(raw), cost0)
	if err != nil {
		http.Error(w, "Password encryption failed", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(hash)
}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	raw := r.Form.Get("raw")
	hash := r.Form.Get("hash")
	if raw == "" || hash == "" {
		http.Error(w, "Missing params", 400)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	if err != nil {
		http.Error(w, "Invalid password", 400)
		return
	}

	w.WriteHeader(200)
}
