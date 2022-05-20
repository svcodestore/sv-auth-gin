package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _RolesMgr struct {
	*_BaseMgr
}

// RolesMgr open func
func RolesMgr(db *gorm.DB) *_RolesMgr {
	if db == nil {
		panic(fmt.Errorf("RolesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RolesMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Roles{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RolesMgr) GetTableName() string {
	return "roles"
}

// Get 获取
func (obj *_RolesMgr) Get() (result Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RolesMgr) Gets() (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_RolesMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPid pid获取
func (obj *_RolesMgr) WithPid(pid string) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithApplicationID application_id获取
func (obj *_RolesMgr) WithApplicationID(applicationID string) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithCode code获取
func (obj *_RolesMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取
func (obj *_RolesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIsGroup is_group获取
func (obj *_RolesMgr) WithIsGroup(isGroup uint8) Option {
	return optionFunc(func(o *options) { o.query["is_group"] = isGroup })
}

// WithStatus status获取
func (obj *_RolesMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_RolesMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_RolesMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_RolesMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_RolesMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_RolesMgr) GetByOption(opts ...Option) (result Roles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RolesMgr) GetByOptions(opts ...Option) (results []*Roles, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_RolesMgr) GetFromID(id string) (result Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_RolesMgr) GetBatchFromID(ids []string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容
func (obj *_RolesMgr) GetFromPid(pid string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`pid` = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量查找
func (obj *_RolesMgr) GetBatchFromPid(pids []string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`pid` IN (?)", pids).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_RolesMgr) GetFromApplicationID(applicationID string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量查找
func (obj *_RolesMgr) GetBatchFromApplicationID(applicationIDs []string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_RolesMgr) GetFromCode(code string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_RolesMgr) GetBatchFromCode(codes []string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_RolesMgr) GetFromName(name string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_RolesMgr) GetBatchFromName(names []string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromIsGroup 通过is_group获取内容
func (obj *_RolesMgr) GetFromIsGroup(isGroup uint8) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`is_group` = ?", isGroup).Find(&results).Error

	return
}

// GetBatchFromIsGroup 批量查找
func (obj *_RolesMgr) GetBatchFromIsGroup(isGroups []uint8) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`is_group` IN (?)", isGroups).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_RolesMgr) GetFromStatus(status uint8) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_RolesMgr) GetBatchFromStatus(statuss []uint8) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_RolesMgr) GetFromCreatedAt(createdAt time.Time) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_RolesMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_RolesMgr) GetFromCreatedBy(createdBy string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_RolesMgr) GetBatchFromCreatedBy(createdBys []string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_RolesMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_RolesMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_RolesMgr) GetFromUpdatedBy(updatedBy string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_RolesMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RolesMgr) FetchByPrimaryKey(id string) (result Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByRolesUniqueIndex primary or index 获取唯一内容
func (obj *_RolesMgr) FetchUniqueIndexByRolesUniqueIndex(pid string, applicationID string, code string) (result Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`pid` = ? AND `application_id` = ? AND `code` = ?", pid, applicationID, code).Find(&result).Error

	return
}

// FetchIndexByRolesFkCreatedBy  获取多个内容
func (obj *_RolesMgr) FetchIndexByRolesFkCreatedBy(createdBy string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// FetchIndexByRolesFkUpdatedBy  获取多个内容
func (obj *_RolesMgr) FetchIndexByRolesFkUpdatedBy(updatedBy string) (results []*Roles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}
