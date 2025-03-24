package mesh

import "testing"

func TestMeshNetworkStart(t *testing.T) {
    meshNetwork := NewMeshNetwork()
    err := meshNetwork.Start()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
}

