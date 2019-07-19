package orm

import (
	"time"
)

//数据库基础model with tenant
type TenantModel struct {
	ID        int64      `json:"id" gorm:"primary_key;size:19"`
	Tenant    int64      `json:"tenant,omitempty" gorm:"not null;size:19"`
	CreatedAt time.Time  `json:"-"`
	CreatedBy int64      `json:"-" gorm:"size:19"`
	UpdatedAt time.Time  `json:",omitempty"`
	UpdatedBy int64      `json:"-" gorm:"size:19"`
	DeletedAt *time.Time `json:",omitempty"`
	DeletedBy int64      `json:"-" gorm:"size:19"`
}

//数据库基础model
type Model struct {
	ID        int64      `json:"id" gorm:"primary_key;size:19"`
	CreatedAt time.Time  `json:"-"`
	CreatedBy int64      `json:"-" gorm:"size:19"`
	UpdatedAt time.Time  `json:",omitempty"`
	UpdatedBy int64      `json:"-" gorm:"size:19"`
	DeletedAt *time.Time `json:",omitempty"`
	DeletedBy int64      `json:"-" gorm:"size:19"`
}
