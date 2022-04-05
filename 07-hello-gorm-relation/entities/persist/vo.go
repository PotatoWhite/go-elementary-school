package persist

import (
	"gorm.io/gorm"
	"time"
)

type Crop struct {
	gorm.Model
	Name      string
	Quantity  int64
	FarmID    uint
	Farm      Farm
	HarvestAt time.Time
}

type Farm struct {
	gorm.Model
	Name string

	OwnerId uint
	Person  Person `gorm:"foreignKey:OwnerId"`
}

type Person struct {
	gorm.Model
	Name     string
	Birthday time.Time
}
