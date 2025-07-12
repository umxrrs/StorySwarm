package web

import (
    "html/template"
    "net/http"
    "github.com/umxrrs/StorySwarm/internal/db"
    "github.com/gorilla/mux"
)

func Start(port string) {
    r := mux.NewRouter()
    r.HandleFunc("/stories/{channelID}", storyHandler)
    http.ListenAndServe(":"+port, r)
}

func storyHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    channelID := vars["channelID"]
    story, err := db.GetActiveStory(channelID)
    if err != nil || story == nil {
        http.Error(w, "Story not found", http.StatusNotFound)
        return
    }
    contributions, _ := db.GetContributions(channelID)
    tmpl, _ := template.ParseFiles("internal/web/templates/story.html")
    data := struct {
        Theme         string
        Tone          string
        Contributions []db.Contribution
    }{story.Theme, story.Tone, contributions}
    tmpl.Execute(w, data)
}
