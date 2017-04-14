package main

  import (
    "testing"
  )


func TestDeveRetornarErro(t *testing.T) {
    if true != false {
        t.Fatalf("Ferrou %v:\n\t nao eh  %v", true, false)
    }

}