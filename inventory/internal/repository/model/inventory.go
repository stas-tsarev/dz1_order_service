package model

import "time"

// Value - интерфейс для метаданных
type Value interface {
	GetValue() interface{}
	GetType() string
}

// StringValue - строковое значение
type StringValue struct {
	Value string `json:"value"`
}

func (s StringValue) GetValue() interface{} {
	return s.Value
}

func (s StringValue) GetType() string {
	return "string"
}

// Int64Value - целочисленное значение
type Int64Value struct {
	Value int64 `json:"value"`
}

func (i Int64Value) GetValue() interface{} {
	return i.Value
}

func (i Int64Value) GetType() string {
	return "int64"
}

// DoubleValue - дробное значение
type DoubleValue struct {
	Value float64 `json:"value"`
}

func (d DoubleValue) GetValue() interface{} {
	return d.Value
}

func (d DoubleValue) GetType() string {
	return "double"
}

// BoolValue - логическое значение
type BoolValue struct {
	Value bool `json:"value"`
}

func (b BoolValue) GetValue() interface{} {
	return b.Value
}

func (b BoolValue) GetType() string {
	return "bool"
}

type Category int

const (
	CATEGORY_UNKNOWN_UNSPECIFIED Category = iota // Неизвестная категория
	CATEGORY_ENGINE                              // Двигатель
	CATEGORY_FUEL                                // Топливо
	CATEGORY_PORTHOLE                            // Иллюминатор
	CATEGORY_WING                                // Крыло                 // Деньги инвестора (внутренний метод)
)

func (i Category) String() string {
	switch i {
	case CATEGORY_UNKNOWN_UNSPECIFIED:
		return "CATEGORY_UNKNOWN_UNSPECIFIED"
	case CATEGORY_ENGINE:
		return "CATEGORY_ENGINE"
	case CATEGORY_FUEL:
		return "CATEGORY_FUEL"
	case CATEGORY_PORTHOLE:
		return "CATEGORY_PORTHOLE"
	case CATEGORY_WING:
		return "CATEGORY_WING"
	}
	return "CATEGORY_UNKNOWN_UNSPECIFIED"
}

type Dimensions struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

type Manufacturer struct {
	Name    string
	Country string
	WebSite string
}

type PartInfo struct {
	Name          string // Название детали
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	Metadata      map[string]Value
}

// Part содержит в себе информацию о детали
type Part struct {
	Uuid      string
	Info      PartInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PartsFilter struct {
	Uuids                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}
