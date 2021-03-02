package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
)

const VARNAME = "API_KEY"

func main() {
        _, apiKeySet := os.LookupEnv(VARNAME)
        if !apiKeySet {
                log.Fatal("API_KEY is not set")
        }
        fmt.Println("Ambassador running on :8080")

        http.HandleFunc("/movies", TheMovieDBServer)
        http.ListenAndServe(":8080", nil)
}

func TheMovieDBServer(w http.ResponseWriter, r *http.Request) {
        apiKey := os.Getenv(VARNAME)
        resp, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?api_key=%s", apiKey))
        if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(err.Error()))
        }
        defer resp.Body.Close()

        responseData, err := ioutil.ReadAll(resp.Body)

        responseString := string(responseData)
        fmt.Fprint(w, responseString)
}
