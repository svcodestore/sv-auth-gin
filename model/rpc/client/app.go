package client

import (

	"time"
)

type Applications struct {
	ID            string              `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	Code          string              `gorm:"unique;column:code;type:varchar(64);not null" json:"code"`
	Name          string              `gorm:"unique;column:name;type:varchar(255);not null" json:"name"`
	InternalURL   string              `gorm:"column:internal_url;type:varchar(255)" json:"internalUrl"`
	HomepageURL   string              `gorm:"column:homepage_url;type:varchar(255)" json:"homepageUrl"`
	Status        bool                `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	ClientID      string              `gorm:"unique;column:client_id;type:varchar(255);not null" json:"clientId"`
	ClientSecret  string              `gorm:"column:client_secret;type:varchar(255)" json:"clientSecret"`
	RedirectURIs  string              `gorm:"column:redirect_uris;type:varchar(255);not null" json:"redirectUris"`
	LoginURIs     string              `gorm:"column:login_uris;type:varchar(255);not null" json:"loginUris"`
	TokenFormat   string              `gorm:"column:token_format;type:varchar(100);default:JWT" json:"tokenFormat"`
	CreatedAt     time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     string              `gorm:"column:created_by;type:bigint;not null" json:"-"`
	CreatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt     time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     string              `gorm:"column:updated_by;type:bigint;not null" json:"-"`
	UpdatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:UpdatedBy;reference:UpdatedBy" json:"updatedByUser"`
}