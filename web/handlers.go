package web

import (
    "html/template"
    "net/http"
    "path/filepath"
)

// StartWebServer initializes and starts the web server
func StartWebServer() {
    http.HandleFunc("/", dashboardHandler)
    http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Prometheus metrics are already handled in main.go
        http.Redirect(w, r, "/metrics", http.StatusMovedPermanently)
    }))

    http.ListenAndServe(":8080", nil)
}

// dashboardHandler serves the dashboard HTML
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles(filepath.Join("web", "templates", "dashboard.html"))
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}
