package persist

import (
	"gorm.io/gorm"
	"time"
)

type Crop struct {
	gorm.Model
	Name      string
	Quantity  int64
	HarvestAt time.Time
}

type CropStatistics struct {
	Name  string
	Total int64
}
