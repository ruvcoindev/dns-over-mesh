package main

import (
    "log"
    "dns-over-mesh/mesh"
    "dns-over-mesh/dns"
    "dns-over-mesh/web"
)

func main() {
    // Инициализация MESH-сети
    meshNetwork := mesh.NewMeshNetwork()
    if err := meshNetwork.Start(); err != nil {
        log.Fatalf("Failed to start mesh network: %v", err)
    }

    // Инициализация DNS-сервера
    dnsServer := dns.NewDNSServer(meshNetwork)
    if err := dnsServer.Start(); err != nil {
        log.Fatalf("Failed to start DNS server: %v", err)
    }

    // Запуск веб-сервера для мониторинга
    webServer := web.NewWebServer(dnsServer)
    if err := webServer.Start(); err != nil {
        log.Fatalf("Failed to start web server: %v", err)
    }
}

