package web

import (
    "html/template"
    "log"
    "net/http"
    "dns-over-mesh/dns"
)

type WebServer struct {
    dnsServer *dns.DNSServer
    templates *template.Template
}

func NewWebServer(dnsServer *dns.DNSServer) *WebServer {
    return &WebServer{
        dnsServer: dnsServer,
        templates: template.Must(template.ParseGlob("web/templates/*.html")),
    }
}

func (ws *WebServer) Start() error {
    http.HandleFunc("/", ws.handleRoot)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
    log.Println("Web server started on :8080")
    return http.ListenAndServe(":8080", nil)
}

func (ws *WebServer) handleRoot(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
        "Status":         "active",
        "QueriesHandled": 1234,
        "Uptime":         "24h",
    }
    ws.templates.ExecuteTemplate(w, "index.html", data)
}

