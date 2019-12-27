package models

import "testing"

func TestInitInterfaceModel(t *testing.T) {
	int := InitInterfaceModel(10, 5)
	int.Add()
}
