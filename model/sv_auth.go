package model

import (
	"time"
)

// ActionMenu [...]
type ActionMenu struct {
	ID        string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	ActionID  string    `gorm:"index:action_menu_action_id_index;column:action_id;type:bigint;not null" json:"actionId"`
	Actions   Actions   `gorm:"joinForeignKey:action_id;foreignKey:ActionID" json:"actionsList"`
	MenuID    string    `gorm:"index:action_menu_menu_id_index;column:menu_id;type:bigint;not null" json:"menuId"`
	Menus     Menus     `gorm:"joinForeignKey:menu_id;foreignKey:MenuID" json:"menusList"`
	Status    bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy string    `gorm:"index:action_menu_fk_created_by;column:created_by;type:bigint;not null" json:"createdBy"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy string    `gorm:"index:action_menu_fk_updated_by;column:updated_by;type:bigint;not null" json:"updatedBy"`
}

// TableName get sql table name.获取数据库表名
func (m *ActionMenu) TableName() string {
	return "action_menu"
}

// ActionMenuColumns get sql column name.获取数据库列名
var ActionMenuColumns = struct {
	ID        string
	ActionID  string
	MenuID    string
	Status    string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	ActionID:  "action_id",
	MenuID:    "menu_id",
	Status:    "status",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// Actions [...]
type Actions struct {
	ID            string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	ApplicationID string    `gorm:"uniqueIndex:actions_unique1_index;uniqueIndex:actions_unique2_index;uniqueIndex:actions_unique3_index;column:application_id;type:bigint;not null" json:"applicationId"`
	Code          string    `gorm:"uniqueIndex:actions_unique1_index;uniqueIndex:actions_unique2_index;column:code;type:varchar(64);not null" json:"code"`
	Name          string    `gorm:"uniqueIndex:actions_unique1_index;uniqueIndex:actions_unique3_index;column:name;type:varchar(255);not null" json:"name"`
	Status        bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     string    `gorm:"index:actions_fk_created_by;column:created_by;type:bigint;not null" json:"createdBy"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     string    `gorm:"index:actions_fk_updated_by;column:updated_by;type:bigint;not null" json:"updatedBy"`
}

// TableName get sql table name.获取数据库表名
func (m *Actions) TableName() string {
	return "actions"
}

// ActionsColumns get sql column name.获取数据库列名
var ActionsColumns = struct {
	ID            string
	ApplicationID string
	Code          string
	Name          string
	Status        string
	CreatedAt     string
	CreatedBy     string
	UpdatedAt     string
	UpdatedBy     string
}{
	ID:            "id",
	ApplicationID: "application_id",
	Code:          "code",
	Name:          "name",
	Status:        "status",
	CreatedAt:     "created_at",
	CreatedBy:     "created_by",
	UpdatedAt:     "updated_at",
	UpdatedBy:     "updated_by",
}

// Menus [...]
type Menus struct {
	ID            string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	Pid           string    `gorm:"uniqueIndex:menus_unique4_index;uniqueIndex:menus_unique5_index;column:pid;type:bigint;not null" json:"pid"`
	ApplicationID string    `gorm:"uniqueIndex:menus_unique1_index;uniqueIndex:menus_unique2_index;uniqueIndex:menus_unique3_index;uniqueIndex:menus_unique4_index;uniqueIndex:menus_unique5_index;column:application_id;type:bigint;not null" json:"applicationId"`
	Code          string    `gorm:"uniqueIndex:menus_unique1_index;uniqueIndex:menus_unique2_index;uniqueIndex:menus_unique5_index;column:code;type:varchar(64);not null" json:"code"`
	Name          string    `gorm:"uniqueIndex:menus_unique1_index;uniqueIndex:menus_unique3_index;column:name;type:varchar(255);not null" json:"name"`
	SortNo        uint16    `gorm:"column:sort_no;type:smallint unsigned;not null" json:"sortNo"`
	Path          string    `gorm:"uniqueIndex:menus_unique4_index;column:path;type:varchar(255);not null" json:"path"`
	Redirect      string    `gorm:"column:redirect;type:varchar(255)" json:"redirect"`
	Component     string    `gorm:"column:component;type:varchar(255);not null" json:"component"`
	Icon          string    `gorm:"column:icon;type:varchar(255)" json:"icon"`
	Hide          bool      `gorm:"column:hide;type:tinyint(1);not null;default:0" json:"hide"`
	Status        bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     string    `gorm:"index:menus_fk_created_by;column:created_by;type:bigint;not null" json:"createdBy"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     string    `gorm:"index:menus_fk_updated_by;column:updated_by;type:bigint;not null" json:"updatedBy"`
}

// TableName get sql table name.获取数据库表名
func (m *Menus) TableName() string {
	return "menus"
}

// MenusColumns get sql column name.获取数据库列名
var MenusColumns = struct {
	ID            string
	Pid           string
	ApplicationID string
	Code          string
	Name          string
	SortNo        string
	Path          string
	Redirect      string
	Component     string
	Icon          string
	Hide          string
	Status        string
	CreatedAt     string
	CreatedBy     string
	UpdatedAt     string
	UpdatedBy     string
}{
	ID:            "id",
	Pid:           "pid",
	ApplicationID: "application_id",
	Code:          "code",
	Name:          "name",
	SortNo:        "sort_no",
	Path:          "path",
	Redirect:      "redirect",
	Component:     "component",
	Icon:          "icon",
	Hide:          "hide",
	Status:        "status",
	CreatedAt:     "created_at",
	CreatedBy:     "created_by",
	UpdatedAt:     "updated_at",
	UpdatedBy:     "updated_by",
}

// RoleMenu [...]
type RoleMenu struct {
	ID        string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	RoleID    string    `gorm:"index:role_menu_role_id_index;column:role_id;type:bigint;not null" json:"roleId"`
	Roles     Roles     `gorm:"joinForeignKey:role_id;foreignKey:RoleID" json:"rolesList"`
	MenuID    string    `gorm:"index:role_menu_menu_id_index;column:menu_id;type:bigint;not null" json:"menuId"`
	Menus     Menus     `gorm:"joinForeignKey:menu_id;foreignKey:MenuID" json:"menusList"`
	Status    bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy string    `gorm:"index:role_menu_fk_created_by;column:created_by;type:bigint;not null" json:"createdBy"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy string    `gorm:"index:role_menu_fk_updated_by;column:updated_by;type:bigint;not null" json:"updatedBy"`
}

// TableName get sql table name.获取数据库表名
func (m *RoleMenu) TableName() string {
	return "role_menu"
}

// RoleMenuColumns get sql column name.获取数据库列名
var RoleMenuColumns = struct {
	ID        string
	RoleID    string
	MenuID    string
	Status    string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	RoleID:    "role_id",
	MenuID:    "menu_id",
	Status:    "status",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// RoleUser [...]
type RoleUser struct {
	ID        string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	RoleID    string    `gorm:"uniqueIndex:role_user_unique_index;index:role_user_role_id_index;column:role_id;type:bigint;not null" json:"roleId"`
	Roles     Roles     `gorm:"joinForeignKey:role_id;foreignKey:id" json:"rolesList"`
	UserID    string    `gorm:"uniqueIndex:role_user_unique_index;index:role_user_user_id_index;column:user_id;type:bigint;not null" json:"userId"`
	Status    bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy string    `gorm:"index:role_user_fk_created_by;column:created_by;type:bigint;not null" json:"createdBy"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy string    `gorm:"index:role_user_fk_updated_by;column:updated_by;type:bigint;not null" json:"updatedBy"`
}

// TableName get sql table name.获取数据库表名
func (m *RoleUser) TableName() string {
	return "role_user"
}

// RoleUserColumns get sql column name.获取数据库列名
var RoleUserColumns = struct {
	ID        string
	RoleID    string
	UserID    string
	Status    string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	RoleID:    "role_id",
	UserID:    "user_id",
	Status:    "status",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// RoleUserAction [...]
type RoleUserAction struct {
	ID         string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	RoleUserID string    `gorm:"index:role_user__action_role_user_id_index;column:role_user_id;type:bigint;not null" json:"roleUserId"`
	RoleUser   RoleUser  `gorm:"joinForeignKey:role_user_id;foreignKey:id" json:"roleUserList"`
	ActionID   string    `gorm:"index:role_user__action_action_id_index;column:action_id;type:bigint;not null" json:"actionId"`
	Actions    Actions   `gorm:"joinForeignKey:action_id;foreignKey:id" json:"actionsList"`
	Status     bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy  string    `gorm:"index:role_user__action_fk_created_by;column:created_by;type:bigint;not null" json:"createdBy"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy  string    `gorm:"index:role_user__action_fk_updated_by;column:updated_by;type:bigint;not null" json:"updatedBy"`
}

// TableName get sql table name.获取数据库表名
func (m *RoleUserAction) TableName() string {
	return "role_user__action"
}

// RoleUserActionColumns get sql column name.获取数据库列名
var RoleUserActionColumns = struct {
	ID         string
	RoleUserID string
	ActionID   string
	Status     string
	CreatedAt  string
	CreatedBy  string
	UpdatedAt  string
	UpdatedBy  string
}{
	ID:         "id",
	RoleUserID: "role_user_id",
	ActionID:   "action_id",
	Status:     "status",
	CreatedAt:  "created_at",
	CreatedBy:  "created_by",
	UpdatedAt:  "updated_at",
	UpdatedBy:  "updated_by",
}

// Roles [...]
type Roles struct {
	ID            string    `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`
	Pid           string    `gorm:"column:pid;type:bigint;not null" json:"pid"`
	ApplicationID string    `gorm:"uniqueIndex:roles_unique1_index;uniqueIndex:roles_unique2_index;uniqueIndex:roles_unique3_index;column:application_id;type:bigint;not null" json:"applicationId"`
	Code          string    `gorm:"uniqueIndex:roles_unique1_index;uniqueIndex:roles_unique2_index;column:code;type:varchar(64);not null" json:"code"`
	Name          string    `gorm:"uniqueIndex:roles_unique1_index;uniqueIndex:roles_unique3_index;column:name;type:varchar(255);not null" json:"name"`
	IsGroup       bool      `gorm:"column:is_group;type:tinyint(1);not null;default:0" json:"isGroup"`
	Status        bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"createdAt"`
	CreatedBy     string    `gorm:"index:roles_fk_created_by;column:created_by;type:bigint;not null" json:"createdBy"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"updatedAt"`
	UpdatedBy     string    `gorm:"index:roles_fk_updated_by;column:updated_by;type:bigint;not null" json:"updatedBy"`
}

// TableName get sql table name.获取数据库表名
func (m *Roles) TableName() string {
	return "roles"
}

// RolesColumns get sql column name.获取数据库列名
var RolesColumns = struct {
	ID            string
	Pid           string
	ApplicationID string
	Code          string
	Name          string
	IsGroup       string
	Status        string
	CreatedAt     string
	CreatedBy     string
	UpdatedAt     string
	UpdatedBy     string
}{
	ID:            "id",
	Pid:           "pid",
	ApplicationID: "application_id",
	Code:          "code",
	Name:          "name",
	IsGroup:       "is_group",
	Status:        "status",
	CreatedAt:     "created_at",
	CreatedBy:     "created_by",
	UpdatedAt:     "updated_at",
	UpdatedBy:     "updated_by",
}
