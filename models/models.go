package models

type FuzzTestCase struct {
	Description string      `json:"description"`
	Name        interface{} `json:"name"`
	Age         interface{} `json:"age"`
}
