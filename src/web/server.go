package web

import (
    "log"
    "net/http"
    "dns-over-mesh/dns"
)

type WebServer struct {
    dnsServer *dns.DNSServer
    // Поля и методы для управления веб-сервером
}

func NewWebServer(dnsServer *dns.DNSServer) *WebServer {
    return &WebServer{dnsServer: dnsServer}
}

func (ws *WebServer) Start() error {
    http.HandleFunc("/", ws.handleRoot)
    log.Println("Web server started on :8080")
    return http.ListenAndServe(":8080", nil)
}

func (ws *WebServer) handleRoot(w http.ResponseWriter, r *http.Request) {
    // Логика обработки запросов
    w.Write([]byte("DNS over MESH Monitoring"))
}

