package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _RoleUserActionMgr struct {
	*_BaseMgr
}

// RoleUserActionMgr open func
func RoleUserActionMgr(db *gorm.DB) *_RoleUserActionMgr {
	if db == nil {
		panic(fmt.Errorf("RoleUserActionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RoleUserActionMgr{_BaseMgr: &_BaseMgr{DB: db.Model(RoleUserAction{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RoleUserActionMgr) GetTableName() string {
	return "role_user__action"
}

// Get 获取
func (obj *_RoleUserActionMgr) Get() (result RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("role_user").Where("id = ?", result.RoleUserID).Find(&result.RoleUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_RoleUserActionMgr) Gets() (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_RoleUserActionMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithApplicationID application_id获取
func (obj *_RoleUserActionMgr) WithApplicationID(applicationID string) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithRoleUserID role_user_id获取
func (obj *_RoleUserActionMgr) WithRoleUserID(roleUserID string) Option {
	return optionFunc(func(o *options) { o.query["role_user_id"] = roleUserID })
}

// WithActionID action_id获取
func (obj *_RoleUserActionMgr) WithActionID(actionID string) Option {
	return optionFunc(func(o *options) { o.query["action_id"] = actionID })
}

// WithStatus status获取
func (obj *_RoleUserActionMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_RoleUserActionMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_RoleUserActionMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_RoleUserActionMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_RoleUserActionMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_RoleUserActionMgr) GetByOption(opts ...Option) (result RoleUserAction, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("role_user").Where("id = ?", result.RoleUserID).Find(&result.RoleUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RoleUserActionMgr) GetByOptions(opts ...Option) (results []*RoleUserAction, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_RoleUserActionMgr) GetFromID(id string) (result RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("role_user").Where("id = ?", result.RoleUserID).Find(&result.RoleUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_RoleUserActionMgr) GetBatchFromID(ids []string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_RoleUserActionMgr) GetFromApplicationID(applicationID string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicationID 批量查找
func (obj *_RoleUserActionMgr) GetBatchFromApplicationID(applicationIDs []string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` IN (?)", applicationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRoleUserID 通过role_user_id获取内容
func (obj *_RoleUserActionMgr) GetFromRoleUserID(roleUserID string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_user_id` = ?", roleUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRoleUserID 批量查找
func (obj *_RoleUserActionMgr) GetBatchFromRoleUserID(roleUserIDs []string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_user_id` IN (?)", roleUserIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromActionID 通过action_id获取内容
func (obj *_RoleUserActionMgr) GetFromActionID(actionID string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`action_id` = ?", actionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromActionID 批量查找
func (obj *_RoleUserActionMgr) GetBatchFromActionID(actionIDs []string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`action_id` IN (?)", actionIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容
func (obj *_RoleUserActionMgr) GetFromStatus(status uint8) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_RoleUserActionMgr) GetBatchFromStatus(statuss []uint8) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_RoleUserActionMgr) GetFromCreatedAt(createdAt time.Time) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_RoleUserActionMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_RoleUserActionMgr) GetFromCreatedBy(createdBy string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_RoleUserActionMgr) GetBatchFromCreatedBy(createdBys []string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_RoleUserActionMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_RoleUserActionMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_RoleUserActionMgr) GetFromUpdatedBy(updatedBy string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_RoleUserActionMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
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
func (obj *_RoleUserActionMgr) FetchByPrimaryKey(id string) (result RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("role_user").Where("id = ?", result.RoleUserID).Find(&result.RoleUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByRoleUserActionUniqueIndex primary or index 获取唯一内容
func (obj *_RoleUserActionMgr) FetchUniqueIndexByRoleUserActionUniqueIndex(applicationID string, roleUserID string, actionID string) (result RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `role_user_id` = ? AND `action_id` = ?", applicationID, roleUserID, actionID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("role_user").Where("id = ?", result.RoleUserID).Find(&result.RoleUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("actions").Where("id = ?", result.ActionID).Find(&result.Actions).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByRoleUserActionRoleUserIDIndex  获取多个内容
func (obj *_RoleUserActionMgr) FetchIndexByRoleUserActionRoleUserIDIndex(roleUserID string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_user_id` = ?", roleUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoleUserActionActionIDIndex  获取多个内容
func (obj *_RoleUserActionMgr) FetchIndexByRoleUserActionActionIDIndex(actionID string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`action_id` = ?", actionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoleUserActionFkCreatedBy  获取多个内容
func (obj *_RoleUserActionMgr) FetchIndexByRoleUserActionFkCreatedBy(createdBy string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoleUserActionFkUpdatedBy  获取多个内容
func (obj *_RoleUserActionMgr) FetchIndexByRoleUserActionFkUpdatedBy(updatedBy string) (results []*RoleUserAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("role_user").Where("id = ?", results[i].RoleUserID).Find(&results[i].RoleUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("actions").Where("id = ?", results[i].ActionID).Find(&results[i].Actions).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
