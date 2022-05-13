package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _MenusMgr struct {
	*_BaseMgr
}

// MenusMgr open func
func MenusMgr(db *gorm.DB) *_MenusMgr {
	if db == nil {
		panic(fmt.Errorf("MenusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MenusMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Menus{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MenusMgr) GetTableName() string {
	return "menus"
}

// Get 获取
func (obj *_MenusMgr) Get() (result Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MenusMgr) Gets() (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MenusMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPid pid获取
func (obj *_MenusMgr) WithPid(pid string) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithApplicationID application_id获取
func (obj *_MenusMgr) WithApplicationID(applicationID string) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithCode code获取
func (obj *_MenusMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取
func (obj *_MenusMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSortNo sort_no获取
func (obj *_MenusMgr) WithSortNo(sortNo uint16) Option {
	return optionFunc(func(o *options) { o.query["sort_no"] = sortNo })
}

// WithPath path获取
func (obj *_MenusMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithRedirect redirect获取
func (obj *_MenusMgr) WithRedirect(redirect string) Option {
	return optionFunc(func(o *options) { o.query["redirect"] = redirect })
}

// WithComponent component获取
func (obj *_MenusMgr) WithComponent(component string) Option {
	return optionFunc(func(o *options) { o.query["component"] = component })
}

// WithIcon icon获取
func (obj *_MenusMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithHide hide获取
func (obj *_MenusMgr) WithHide(hide bool) Option {
	return optionFunc(func(o *options) { o.query["hide"] = hide })
}

// WithStatus status获取
func (obj *_MenusMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_MenusMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_MenusMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_MenusMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_MenusMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_MenusMgr) GetByOption(opts ...Option) (result Menus, err error) {
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
func (obj *_MenusMgr) GetByOptions(opts ...Option) (results []*Menus, err error) {
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
func (obj *_MenusMgr) GetFromID(id string) (result Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MenusMgr) GetBatchFromID(ids []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容
func (obj *_MenusMgr) GetFromPid(pid string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`pid` = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量查找
func (obj *_MenusMgr) GetBatchFromPid(pids []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`pid` IN (?)", pids).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_MenusMgr) GetFromApplicationID(applicationID string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量查找
func (obj *_MenusMgr) GetBatchFromApplicationID(applicationIDs []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_MenusMgr) GetFromCode(code string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_MenusMgr) GetBatchFromCode(codes []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_MenusMgr) GetFromName(name string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_MenusMgr) GetBatchFromName(names []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSortNo 通过sort_no获取内容
func (obj *_MenusMgr) GetFromSortNo(sortNo uint16) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`sort_no` = ?", sortNo).Find(&results).Error

	return
}

// GetBatchFromSortNo 批量查找
func (obj *_MenusMgr) GetBatchFromSortNo(sortNos []uint16) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`sort_no` IN (?)", sortNos).Find(&results).Error

	return
}

// GetFromPath 通过path获取内容
func (obj *_MenusMgr) GetFromPath(path string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`path` = ?", path).Find(&results).Error

	return
}

// GetBatchFromPath 批量查找
func (obj *_MenusMgr) GetBatchFromPath(paths []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`path` IN (?)", paths).Find(&results).Error

	return
}

// GetFromRedirect 通过redirect获取内容
func (obj *_MenusMgr) GetFromRedirect(redirect string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`redirect` = ?", redirect).Find(&results).Error

	return
}

// GetBatchFromRedirect 批量查找
func (obj *_MenusMgr) GetBatchFromRedirect(redirects []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`redirect` IN (?)", redirects).Find(&results).Error

	return
}

// GetFromComponent 通过component获取内容
func (obj *_MenusMgr) GetFromComponent(component string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`component` = ?", component).Find(&results).Error

	return
}

// GetBatchFromComponent 批量查找
func (obj *_MenusMgr) GetBatchFromComponent(components []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`component` IN (?)", components).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容
func (obj *_MenusMgr) GetFromIcon(icon string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`icon` = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量查找
func (obj *_MenusMgr) GetBatchFromIcon(icons []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`icon` IN (?)", icons).Find(&results).Error

	return
}

// GetFromHide 通过hide获取内容
func (obj *_MenusMgr) GetFromHide(hide bool) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`hide` = ?", hide).Find(&results).Error

	return
}

// GetBatchFromHide 批量查找
func (obj *_MenusMgr) GetBatchFromHide(hides []bool) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`hide` IN (?)", hides).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_MenusMgr) GetFromStatus(status bool) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_MenusMgr) GetBatchFromStatus(statuss []bool) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_MenusMgr) GetFromCreatedAt(createdAt time.Time) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_MenusMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_MenusMgr) GetFromCreatedBy(createdBy string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_MenusMgr) GetBatchFromCreatedBy(createdBys []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_MenusMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_MenusMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_MenusMgr) GetFromUpdatedBy(updatedBy string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_MenusMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MenusMgr) FetchByPrimaryKey(id string) (result Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByMenusUnique4Index primary or index 获取唯一内容
func (obj *_MenusMgr) FetchUniqueIndexByMenusUnique4Index(pid string, applicationID string, path string) (result Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`pid` = ? AND `application_id` = ? AND `path` = ?", pid, applicationID, path).Find(&result).Error

	return
}

// FetchUniqueIndexByMenusUnique5Index primary or index 获取唯一内容
func (obj *_MenusMgr) FetchUniqueIndexByMenusUnique5Index(pid string, applicationID string, code string) (result Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`pid` = ? AND `application_id` = ? AND `code` = ?", pid, applicationID, code).Find(&result).Error

	return
}

// FetchUniqueIndexByMenusUnique1Index primary or index 获取唯一内容
func (obj *_MenusMgr) FetchUniqueIndexByMenusUnique1Index(applicationID string, code string, name string) (result Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `code` = ? AND `name` = ?", applicationID, code, name).Find(&result).Error

	return
}

// FetchUniqueIndexByMenusUnique2Index primary or index 获取唯一内容
func (obj *_MenusMgr) FetchUniqueIndexByMenusUnique2Index(applicationID string, code string) (result Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `code` = ?", applicationID, code).Find(&result).Error

	return
}

// FetchUniqueIndexByMenusUnique3Index primary or index 获取唯一内容
func (obj *_MenusMgr) FetchUniqueIndexByMenusUnique3Index(applicationID string, name string) (result Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`application_id` = ? AND `name` = ?", applicationID, name).Find(&result).Error

	return
}

// FetchIndexByMenusFkCreatedBy  获取多个内容
func (obj *_MenusMgr) FetchIndexByMenusFkCreatedBy(createdBy string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// FetchIndexByMenusFkUpdatedBy  获取多个内容
func (obj *_MenusMgr) FetchIndexByMenusFkUpdatedBy(updatedBy string) (results []*Menus, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}
