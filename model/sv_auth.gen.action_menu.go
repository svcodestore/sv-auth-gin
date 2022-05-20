package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ActionMenuMgr struct {
	*_BaseMgr
}

// ActionMenuMgr open func
func ActionMenuMgr(db *gorm.DB) *_ActionMenuMgr {
	if db == nil {
		panic(fmt.Errorf("ActionMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ActionMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Model(ActionMenu{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ActionMenuMgr) GetTableName() string {
	return "action_menu"
}

// Get 获取
func (obj *_ActionMenuMgr) Get() (result ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) Gets() (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithApplicationID application_id获取
func (obj *_ActionMenuMgr) WithApplicationID(applicationID string) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithActionID action_id获取
func (obj *_ActionMenuMgr) WithActionID(actionID string) Option {
	return optionFunc(func(o *options) { o.query["action_id"] = actionID })
}

// WithMenuID menu_id获取
func (obj *_ActionMenuMgr) WithMenuID(menuID string) Option {
	return optionFunc(func(o *options) { o.query["menu_id"] = menuID })
}

// WithStatus status获取
func (obj *_ActionMenuMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_ActionMenuMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_ActionMenuMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_ActionMenuMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_ActionMenuMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_ActionMenuMgr) GetByOption(opts ...Option) (result ActionMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetByOptions(opts ...Option) (results []*ActionMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetFromID(id string) (result ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetBatchFromID(ids []string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetFromApplicationID(applicationID string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetBatchFromApplicationID(applicationIDs []string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` IN (?)", applicationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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

// GetFromActionID 通过action_id获取内容
func (obj *_ActionMenuMgr) GetFromActionID(actionID string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`action_id` = ?", actionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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

// GetBatchFromActionID 批量查找
func (obj *_ActionMenuMgr) GetBatchFromActionID(actionIDs []string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`action_id` IN (?)", actionIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetFromMenuID(menuID string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`menu_id` = ?", menuID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetBatchFromMenuID(menuIDs []string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`menu_id` IN (?)", menuIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetFromStatus(status uint8) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetBatchFromStatus(statuss []uint8) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetFromCreatedAt(createdAt time.Time) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetFromCreatedBy(createdBy string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetBatchFromCreatedBy(createdBys []string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetFromUpdatedBy(updatedBy string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_ActionMenuMgr) FetchByPrimaryKey(id string) (result ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
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

// FetchUniqueIndexByActionMenuUniqueIndex primary or index 获取唯一内容
func (obj *_ActionMenuMgr) FetchUniqueIndexByActionMenuUniqueIndex(applicationID string, actionID string, menuID string) (result ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `action_id` = ? AND `menu_id` = ?", applicationID, actionID, menuID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
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

// FetchIndexByActionMenuActionIDIndex  获取多个内容
func (obj *_ActionMenuMgr) FetchIndexByActionMenuActionIDIndex(actionID string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`action_id` = ?", actionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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

// FetchIndexByActionMenuMenuIDIndex  获取多个内容
func (obj *_ActionMenuMgr) FetchIndexByActionMenuMenuIDIndex(menuID string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`menu_id` = ?", menuID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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

// FetchIndexByActionMenuFkCreatedBy  获取多个内容
func (obj *_ActionMenuMgr) FetchIndexByActionMenuFkCreatedBy(createdBy string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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

// FetchIndexByActionMenuFkUpdatedBy  获取多个内容
func (obj *_ActionMenuMgr) FetchIndexByActionMenuFkUpdatedBy(updatedBy string) (results []*ActionMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
