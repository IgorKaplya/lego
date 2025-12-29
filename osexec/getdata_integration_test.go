package osexec_test

import (
	"testing"

	"github.com/IgorKaplya/lego/osexec"
)

func TestGetDataIntegration(t *testing.T) {
	got := osexec.GetData(osexec.GetReader())
	assertGetData(t, got, "WOOP WOOP")
}
