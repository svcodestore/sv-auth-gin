package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _RoleMenuMgr struct {
	*_BaseMgr
}

// RoleMenuMgr open func
func RoleMenuMgr(db *gorm.DB) *_RoleMenuMgr {
	if db == nil {
		panic(fmt.Errorf("RoleMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RoleMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Model(RoleMenu{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RoleMenuMgr) GetTableName() string {
	return "role_menu"
}

// Get 获取
func (obj *_RoleMenuMgr) Get() (result RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("roles").Where("id = ?", result.RoleID).Find(&result.Roles).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("menus").Where("id = ?", result.MenuID).Find(&result.Menus).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_RoleMenuMgr) Gets() (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_RoleMenuMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithApplicationID application_id获取
func (obj *_RoleMenuMgr) WithApplicationID(applicationID string) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithRoleID role_id获取
func (obj *_RoleMenuMgr) WithRoleID(roleID string) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithMenuID menu_id获取
func (obj *_RoleMenuMgr) WithMenuID(menuID string) Option {
	return optionFunc(func(o *options) { o.query["menu_id"] = menuID })
}

// WithStatus status获取
func (obj *_RoleMenuMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_RoleMenuMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_RoleMenuMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_RoleMenuMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_RoleMenuMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_RoleMenuMgr) GetByOption(opts ...Option) (result RoleMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("roles").Where("id = ?", result.RoleID).Find(&result.Roles).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("menus").Where("id = ?", result.MenuID).Find(&result.Menus).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RoleMenuMgr) GetByOptions(opts ...Option) (results []*RoleMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_RoleMenuMgr) GetFromID(id string) (result RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("roles").Where("id = ?", result.RoleID).Find(&result.Roles).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("menus").Where("id = ?", result.MenuID).Find(&result.Menus).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_RoleMenuMgr) GetBatchFromID(ids []string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_RoleMenuMgr) GetFromApplicationID(applicationID string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicationID 批量查找
func (obj *_RoleMenuMgr) GetBatchFromApplicationID(applicationIDs []string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` IN (?)", applicationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_RoleMenuMgr) GetFromRoleID(roleID string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_id` = ?", roleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRoleID 批量查找
func (obj *_RoleMenuMgr) GetBatchFromRoleID(roleIDs []string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_id` IN (?)", roleIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMenuID 通过menu_id获取内容
func (obj *_RoleMenuMgr) GetFromMenuID(menuID string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`menu_id` = ?", menuID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMenuID 批量查找
func (obj *_RoleMenuMgr) GetBatchFromMenuID(menuIDs []string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`menu_id` IN (?)", menuIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容
func (obj *_RoleMenuMgr) GetFromStatus(status uint8) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_RoleMenuMgr) GetBatchFromStatus(statuss []uint8) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_RoleMenuMgr) GetFromCreatedAt(createdAt time.Time) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_RoleMenuMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_RoleMenuMgr) GetFromCreatedBy(createdBy string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_RoleMenuMgr) GetBatchFromCreatedBy(createdBys []string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_RoleMenuMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_RoleMenuMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_RoleMenuMgr) GetFromUpdatedBy(updatedBy string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_RoleMenuMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RoleMenuMgr) FetchByPrimaryKey(id string) (result RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("roles").Where("id = ?", result.RoleID).Find(&result.Roles).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("menus").Where("id = ?", result.MenuID).Find(&result.Menus).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByRoleMenuUniqueIndex primary or index 获取唯一内容
func (obj *_RoleMenuMgr) FetchUniqueIndexByRoleMenuUniqueIndex(applicationID string, roleID string, menuID string) (result RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `role_id` = ? AND `menu_id` = ?", applicationID, roleID, menuID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("roles").Where("id = ?", result.RoleID).Find(&result.Roles).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("menus").Where("id = ?", result.MenuID).Find(&result.Menus).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByRoleMenuRoleIDIndex  获取多个内容
func (obj *_RoleMenuMgr) FetchIndexByRoleMenuRoleIDIndex(roleID string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_id` = ?", roleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoleMenuMenuIDIndex  获取多个内容
func (obj *_RoleMenuMgr) FetchIndexByRoleMenuMenuIDIndex(menuID string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`menu_id` = ?", menuID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoleMenuFkCreatedBy  获取多个内容
func (obj *_RoleMenuMgr) FetchIndexByRoleMenuFkCreatedBy(createdBy string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoleMenuFkUpdatedBy  获取多个内容
func (obj *_RoleMenuMgr) FetchIndexByRoleMenuFkUpdatedBy(updatedBy string) (results []*RoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("menus").Where("id = ?", results[i].MenuID).Find(&results[i].Menus).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
