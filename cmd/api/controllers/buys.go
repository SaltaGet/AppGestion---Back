package controllers

type EntityRequest struct {
	Name  string `json:"name" example:"Sample Name"`
	Value int    `json:"value" example:"42"`
}