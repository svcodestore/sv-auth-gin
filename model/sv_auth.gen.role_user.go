package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _RoleUserMgr struct {
	*_BaseMgr
}

// RoleUserMgr open func
func RoleUserMgr(db *gorm.DB) *_RoleUserMgr {
	if db == nil {
		panic(fmt.Errorf("RoleUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RoleUserMgr{_BaseMgr: &_BaseMgr{DB: db.Model(RoleUser{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RoleUserMgr) GetTableName() string {
	return "role_user"
}

// Get 获取
func (obj *_RoleUserMgr) Get() (result RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("roles").Where("id = ?", result.RoleID).Find(&result.Roles).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_RoleUserMgr) Gets() (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
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
func (obj *_RoleUserMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithApplicationID application_id获取
func (obj *_RoleUserMgr) WithApplicationID(applicationID string) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithRoleID role_id获取
func (obj *_RoleUserMgr) WithRoleID(roleID string) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithUserID user_id获取
func (obj *_RoleUserMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithStatus status获取
func (obj *_RoleUserMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_RoleUserMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_RoleUserMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_RoleUserMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_RoleUserMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_RoleUserMgr) GetByOption(opts ...Option) (result RoleUser, err error) {
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
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RoleUserMgr) GetByOptions(opts ...Option) (results []*RoleUser, err error) {
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
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_RoleUserMgr) GetFromID(id string) (result RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("roles").Where("id = ?", result.RoleID).Find(&result.Roles).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_RoleUserMgr) GetBatchFromID(ids []string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_RoleUserMgr) GetFromApplicationID(applicationID string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicationID 批量查找
func (obj *_RoleUserMgr) GetBatchFromApplicationID(applicationIDs []string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` IN (?)", applicationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_RoleUserMgr) GetFromRoleID(roleID string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_id` = ?", roleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRoleID 批量查找
func (obj *_RoleUserMgr) GetBatchFromRoleID(roleIDs []string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_id` IN (?)", roleIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_RoleUserMgr) GetFromUserID(userID string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量查找
func (obj *_RoleUserMgr) GetBatchFromUserID(userIDs []string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容
func (obj *_RoleUserMgr) GetFromStatus(status uint8) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_RoleUserMgr) GetBatchFromStatus(statuss []uint8) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_RoleUserMgr) GetFromCreatedAt(createdAt time.Time) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_RoleUserMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_RoleUserMgr) GetFromCreatedBy(createdBy string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_RoleUserMgr) GetBatchFromCreatedBy(createdBys []string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_RoleUserMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_RoleUserMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_RoleUserMgr) GetFromUpdatedBy(updatedBy string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_RoleUserMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
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
func (obj *_RoleUserMgr) FetchByPrimaryKey(id string) (result RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("roles").Where("id = ?", result.RoleID).Find(&result.Roles).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByRoleUserUniqueIndex primary or index 获取唯一内容
func (obj *_RoleUserMgr) FetchUniqueIndexByRoleUserUniqueIndex(applicationID string, roleID string, userID string) (result RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `role_id` = ? AND `user_id` = ?", applicationID, roleID, userID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("roles").Where("id = ?", result.RoleID).Find(&result.Roles).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByRoleUserRoleIDIndex  获取多个内容
func (obj *_RoleUserMgr) FetchIndexByRoleUserRoleIDIndex(roleID string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_id` = ?", roleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoleUserUserIDIndex  获取多个内容
func (obj *_RoleUserMgr) FetchIndexByRoleUserUserIDIndex(userID string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoleUserFkCreatedBy  获取多个内容
func (obj *_RoleUserMgr) FetchIndexByRoleUserFkCreatedBy(createdBy string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoleUserFkUpdatedBy  获取多个内容
func (obj *_RoleUserMgr) FetchIndexByRoleUserFkUpdatedBy(updatedBy string) (results []*RoleUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("roles").Where("id = ?", results[i].RoleID).Find(&results[i].Roles).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
