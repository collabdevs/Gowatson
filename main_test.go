package main

  import (
    "testing"
  )


func TestNaoDeveRetornarErro(t *testing.T) {
    if true == false {
        t.Fatalf("Ferrou %v:\n\t nao eh  %v", true, false)
    }

}