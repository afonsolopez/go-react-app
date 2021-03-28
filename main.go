package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/gobuffalo/packr"
	"github.com/webview/webview"
)

// Message : struct for message
type Message struct {
	Text string `json:"text"`
}

func main() {
	// Bind folder path for packaging with Packr
	folder := packr.NewBox("./reactjs/build")

	// Handle to ./static/build folder on root path
	http.Handle("/", http.FileServer(folder))

	// Handle to showMessage func on /hello path
	http.HandleFunc("/hello", showMessage)

	// Run server at port 8000 as goroutine
	// for non-block working
	go http.ListenAndServe(":8000", nil)

	// Let's open window app with:
	//  - name: Golang App
	//  - address: http://localhost:8000
	//  - sizes: 800x600 px
	debug := true
	w := webview.New(debug)
	//	defer w.Destroy()
	w.SetTitle("Golang App")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:8000")
	w.Run()
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	// Create Message JSON data
	message := Message{runtime.Version()}

	// Return JSON encoding to output
	output, err := json.Marshal(message)

	// Catch error, if it happens
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Write output
	w.Write(output)
}
