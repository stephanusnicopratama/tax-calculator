package services

import "tax-calculator/models"

type ITaxServices interface {
	FindAll() []models.ITax
	Load(id string) models.ITax
	Create(payload models.ITax) bool
}