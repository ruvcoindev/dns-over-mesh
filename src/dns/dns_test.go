package dns

import (
    "dns-over-mesh/mesh"
    "testing"
)

func TestDNSServerStart(t *testing.T) {
    meshNetwork := mesh.NewMeshNetwork()
    dnsServer := NewDNSServer(meshNetwork)
    err := dnsServer.Start()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
}

