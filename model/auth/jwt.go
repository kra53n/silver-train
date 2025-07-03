package authModel

import (
	"time"

	"gorm.io/gorm"

	// "silver-train/types"
	"github.com/google/uuid"
)

type RefreshToken struct {
	gorm.Model
	ID        string `primaryKey;size:36"` // for sqlite
	// ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"` // for postgre
	UserID  string `gorm:"size:36;index"`
	TokenHash string `gorm:"type:text"`
	UserAgent string `gorm:"size:255"`
	IPAddress string `gorm:"size:45"`
	ExpiresAt time.Time
	Revoked   bool
}

func (rt *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	if rt.ID == "" {
		rt.ID = uuid.New().String()
	}
	return nil
}
