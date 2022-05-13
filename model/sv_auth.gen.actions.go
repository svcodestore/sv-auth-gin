package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ActionsMgr struct {
	*_BaseMgr
}

// ActionsMgr open func
func ActionsMgr(db *gorm.DB) *_ActionsMgr {
	if db == nil {
		panic(fmt.Errorf("ActionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ActionsMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Actions{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ActionsMgr) GetTableName() string {
	return "actions"
}

// Get 获取
func (obj *_ActionsMgr) Get() (result Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ActionsMgr) Gets() (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ActionsMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithApplicationID application_id获取
func (obj *_ActionsMgr) WithApplicationID(applicationID string) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithCode code获取
func (obj *_ActionsMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取
func (obj *_ActionsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithStatus status获取
func (obj *_ActionsMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_ActionsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_ActionsMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_ActionsMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_ActionsMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_ActionsMgr) GetByOption(opts ...Option) (result Actions, err error) {
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
func (obj *_ActionsMgr) GetByOptions(opts ...Option) (results []*Actions, err error) {
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
func (obj *_ActionsMgr) GetFromID(id string) (result Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ActionsMgr) GetBatchFromID(ids []string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_ActionsMgr) GetFromApplicationID(applicationID string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量查找
func (obj *_ActionsMgr) GetBatchFromApplicationID(applicationIDs []string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_ActionsMgr) GetFromCode(code string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_ActionsMgr) GetBatchFromCode(codes []string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ActionsMgr) GetFromName(name string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_ActionsMgr) GetBatchFromName(names []string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ActionsMgr) GetFromStatus(status bool) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_ActionsMgr) GetBatchFromStatus(statuss []bool) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ActionsMgr) GetFromCreatedAt(createdAt time.Time) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ActionsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_ActionsMgr) GetFromCreatedBy(createdBy string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_ActionsMgr) GetBatchFromCreatedBy(createdBys []string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_ActionsMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_ActionsMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_ActionsMgr) GetFromUpdatedBy(updatedBy string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_ActionsMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ActionsMgr) FetchByPrimaryKey(id string) (result Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByActionsUnique1Index primary or index 获取唯一内容
func (obj *_ActionsMgr) FetchUniqueIndexByActionsUnique1Index(applicationID string, code string, name string) (result Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `code` = ? AND `name` = ?", applicationID, code, name).Find(&result).Error

	return
}

// FetchUniqueIndexByActionsUnique2Index primary or index 获取唯一内容
func (obj *_ActionsMgr) FetchUniqueIndexByActionsUnique2Index(applicationID string, code string) (result Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `code` = ?", applicationID, code).Find(&result).Error

	return
}

// FetchUniqueIndexByActionsUnique3Index primary or index 获取唯一内容
func (obj *_ActionsMgr) FetchUniqueIndexByActionsUnique3Index(applicationID string, name string) (result Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `name` = ?", applicationID, name).Find(&result).Error

	return
}

// FetchIndexByActionsFkCreatedBy  获取多个内容
func (obj *_ActionsMgr) FetchIndexByActionsFkCreatedBy(createdBy string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// FetchIndexByActionsFkUpdatedBy  获取多个内容
func (obj *_ActionsMgr) FetchIndexByActionsFkUpdatedBy(updatedBy string) (results []*Actions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}
