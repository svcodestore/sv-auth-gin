package client

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type UsersWithoutModInfo struct {
	ID        string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	UUID      uuid.UUID `gorm:"unique;index:user_uuid_index;column:uuid;type:binary(16);not null" json:"uuid"`
	LoginID   string    `gorm:"unique;column:login_id;type:varchar(16);not null" json:"loginId"`
	Password  string    `gorm:"column:password;type:varchar(1024)" json:"-"`
	Name      string    `gorm:"column:name;type:varchar(32)" json:"name"`
	Alias     string    `gorm:"column:alias;type:varchar(32)" json:"alias"`
	Phone     string    `gorm:"column:phone;type:varchar(16)" json:"phone"`
	Email     string    `gorm:"column:email;type:varchar(1024)" json:"email"`
	Lang      string    `gorm:"column:lang;type:char(5);default:zh_CN" json:"lang"`
	Status    bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"-"`
	CreatedBy string    `gorm:"column:created_by;type:bigint;not null" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"-"`
	UpdatedBy string    `gorm:"column:updated_by;type:bigint;not null" json:"-"`
}

type Users struct {
	ID            string              `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	UUID          uuid.UUID           `gorm:"unique;index:user_uuid_index;column:uuid;type:binary(16);not null" json:"uuid"`
	LoginID       string              `gorm:"unique;column:login_id;type:varchar(16);not null" json:"loginId"`
	Password      string              `gorm:"column:password;type:varchar(1024)" json:"password"`
	Name          string              `gorm:"column:name;type:varchar(32)" json:"name"`
	Alias         string              `gorm:"column:alias;type:varchar(32)" json:"alias"`
	Phone         string              `gorm:"column:phone;type:varchar(16)" json:"phone"`
	Email         string              `gorm:"column:email;type:varchar(1024)" json:"email"`
	Lang          string              `gorm:"column:lang;type:char(5);default:zh_CN" json:"lang"`
	Status        bool                `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt     time.Time           `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     string              `gorm:"column:created_by;type:bigint;not null" json:"-"`
	CreatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:CreatedBy;reference:CreatedBy" json:"createdByUser"`
	UpdatedAt     time.Time           `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     string              `gorm:"column:updated_by;type:bigint;not null" json:"-"`
	UpdatedByUser UsersWithoutModInfo `gorm:"joinForeignKey:created_by;foreignKey:UpdatedBy;reference:UpdatedBy" json:"updatedByUser"`
}
