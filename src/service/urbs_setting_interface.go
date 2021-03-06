package service

import (
	"context"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
)

// UrbsSettingInterface ....
type UrbsSettingInterface interface {
	// ***** product ******
	// 读取产品列表，支持条件筛选
	ProductList(ctx context.Context, args *tpl.Pagination) (*tpl.ProductsRes, error)
	// ProductStatistics ...
	ProductStatistics(ctx context.Context, product string) (*tpl.ProductStatisticsRes, error)
	// 创建产品
	ProductCreate(ctx context.Context, body *tpl.NameDescBody) (*tpl.ProductRes, error)
	// 更新产品
	ProductUpdate(ctx context.Context, product string, body *tpl.ProductUpdateBody) (*tpl.ProductRes, error)
	// 下线指定产品功能模块
	ProductOffline(ctx context.Context, product string) (*tpl.BoolRes, error)
	// 删除指定产品
	ProductDelete(ctx context.Context, product string) (*tpl.BoolRes, error)
	ProductApplyRule(ctx context.Context, product string, body *tpl.ApplyRulesBody) (*tpl.BoolRes, error)

	// ***** module ******
	// 读取指定产品的功能模块
	ModuleList(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.ModulesInfoRes, error)
	// 指定产品创建功能模块
	ModuleCreate(ctx context.Context, product string, body *tpl.NameDescBody) (*tpl.ModuleInfoRes, error)
	// 更新指定产品功能模块
	ModuleUpdate(ctx context.Context, product string, module string, body *tpl.ModuleUpdateBody) (*tpl.ModuleInfoRes, error)
	// 下线指定产品功能模块
	ModuleOffline(ctx context.Context, product string, module string) (*tpl.BoolRes, error)

	// ***** setting ******
	// 读取指定产品的配置项
	SettingListByProduct(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.SettingsInfoRes, error)
	// 读取指定产品功能模块的配置项列表
	SettingList(ctx context.Context, args *tpl.ProductModuleURL) (*tpl.SettingsInfoRes, error)
	//  返回产品下环境标签的用户列表
	SettingListUsers(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingUsersInfoRes, error)
	//  返回产品下功能配置项的群组列表
	SettingListGroups(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingGroupsInfoRes, error)
	// 创建指定产品功能模块配置项
	SettingCreate(ctx context.Context, args *tpl.ProductModuleURL, body *tpl.SettingBody) (*tpl.SettingInfoRes, error)
	// 创建指定产品功能模块配置项的灰度发布规则
	SettingCreateRule(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.SettingRuleBody) (*tpl.SettingRuleInfoRes, error)
	// 读取指定产品功能模块配置项的灰度发布规则列表
	SettingListRule(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingRulesInfoRes, error)
	// 更新指定产品功能模块配置项的指定灰度发布规则
	SettingUpdateRule(ctx context.Context, args *tpl.ProductModuleSettingHIDURL, body *tpl.SettingRuleBody) (*tpl.SettingRuleInfoRes, error)
	// 创建指定产品功能模块配置项的灰度发布规则
	SettingDeleteRule(ctx context.Context, args *tpl.ProductModuleSettingHIDURL) (*tpl.BoolRes, error)
	// 读取指定产品功能模块配置项
	SettingGet(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingInfoRes, error)
	// 更新指定产品功能模块配置项
	SettingUpdate(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.SettingUpdateBody) (*tpl.SettingInfoRes, error)
	// 下线指定产品功能模块配置项
	SettingOffline(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.BoolRes, error)
	// 批量为用户或群组设置产品功能模块配置项
	SettingAssign(ctx context.Context, args *tpl.ProductModuleSettingURL, body *urbssetting.UsersGroupsBody) (*tpl.SettingReleaseInfoRes, error)
	// Recall 撤销指定批次的用户或群组的配置项
	SettingRecall(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.RecallBody) (*tpl.BoolRes, error)
	// SettingCleanUp 清除指定产品功能模块配置项下所有的用户、群组和百分比规则
	SettingCleanUp(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.BoolRes, error)

	// 删除指定用户的配置项
	SettingDeleteUser(ctx context.Context, args *tpl.ProductModuleSettingUIDURL) (*tpl.BoolRes, error)
	// 回滚指定用户的配置项
	SettingRollbackUserSetting(ctx context.Context, args *tpl.ProductModuleSettingUIDURL) (*tpl.BoolRes, error)
	// 删除指定群组的配置项
	SettingDeleteGroup(ctx context.Context, args *tpl.ProductModuleSettingUIDURL) (*tpl.BoolRes, error)
	// 回滚指定群组的配置项
	SettingRollbackGroupSetting(ctx context.Context, args *tpl.ProductModuleSettingUIDURL) (*tpl.BoolRes, error)

	// ***** label ******
	// 读取指定产品环境标签
	LabelList(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.LabelsInfoRes, error)
	// 返回产品下环境标签的用户列表
	LabelListUsers(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelUsersInfoRes, error)
	// 返回产品下环境标签的群组列表
	LabelListGroups(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelGroupsInfoRes, error)
	// 给指定产品创建环境标签
	LabelCreate(ctx context.Context, product string, body *tpl.LabelBody) (*tpl.LabelInfoRes, error)
	// 创建指定产品功能模块配置项的灰度发布规则
	LabelCreateRule(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.LabelRuleBody) (*tpl.LabelRuleInfoRes, error)
	// 读取指定产品环境标签的灰度发布规则列表
	LabelListRule(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelRulesInfoRes, error)
	// 更新指定产品功能模块配置项的指定灰度发布规则
	LabelUpdateRule(ctx context.Context, args *tpl.ProductLabelHIDURL, body *tpl.LabelRuleBody) (*tpl.LabelRuleInfoRes, error)
	// 创建指定产品功能模块配置项的灰度发布规则
	LabelDeleteRule(ctx context.Context, args *tpl.ProductLabelHIDURL) (*tpl.BoolRes, error)
	// 更新指定产品环境标签
	LabelUpdate(ctx context.Context, product string, label string, body *tpl.LabelUpdateBody) (*tpl.LabelInfoRes, error)
	// 删除指定产品环境标签
	LabelDelete(ctx context.Context, product string, label string) (*tpl.BoolRes, error)
	// 下线指定产品环境标签
	LabelOffline(ctx context.Context, product string, label string) (*tpl.BoolRes, error)
	// 批量为用户或群组设置产品环境标签
	LabelAssign(ctx context.Context, product string, label string, body *urbssetting.UsersGroupsBody) (*tpl.LabelReleaseInfoRes, error)
	// 批量撤销对用户或群组设置的产品环境标签
	LabelRecall(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.RecallBody) (*tpl.BoolRes, error)
	// 删除指定用户的指定环境标签
	LabelDeleteUser(ctx context.Context, args *tpl.ProductLabelUIDURL) (*tpl.BoolRes, error)
	// 删除指定用户的配置项
	LabelDeleteGroup(ctx context.Context, args *tpl.ProductLabelUIDURL) (*tpl.BoolRes, error)
	// 清除产品环境标签下所有的用户、群组和百分比规则
	LabelCleanUp(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.BoolRes, error)

	// ***** user ******
	// 读取用户列表，支持条件筛选
	UserList(ctx context.Context, args *tpl.Pagination) (*tpl.UsersRes, error)
	// 读取指定用户的环境标签，支持条件筛选
	UserListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*tpl.MyLabelsRes, error)
	// 强制刷新指定用户的环境标签列表缓存
	UserRefreshCached(ctx context.Context, uid string) (*tpl.UserRes, error)
	// 读取指定用户的功能配置项，支持条件筛选
	UserListSettings(ctx context.Context, args *tpl.UIDPaginationURL) (*tpl.MySettingsRes, error)
	// 读取指定用户的功能配置项，支持条件筛选，数据用于客户端
	UserListSettingsUnionAll(ctx context.Context, args *tpl.MySettingsQueryURL) (*tpl.MySettingsRes, error)
	// 查询指定用户是否存在
	UserCheckExists(ctx context.Context, uid string) (*tpl.BoolRes, error)
	// 批量添加用户
	UserBatchAdd(ctx context.Context, users []string) (*tpl.BoolRes, error)
	// 读取指定用户的环境标签，包括继承自群组的标签，返回轻量级 labels，
	LabelsCache(ctx context.Context, product, uid string) (*tpl.CacheLabelsInfoRes, error)

	// ***** group ******
	// 读取指定群组的环境标签，支持条件筛选
	GroupListLables(ctx context.Context, args *tpl.GroupPaginationURL) (*tpl.MyLabelsRes, error)
	// 读取指定群组的功能配置项，支持条件筛选
	GroupListSettings(ctx context.Context, args *tpl.MySettingsQueryURL) (*tpl.MySettingsRes, error)
	// 读取群组列表，支持条件筛选
	GroupList(ctx context.Context, args *tpl.GroupsURL) (*tpl.GroupsRes, error)
	// 查询指定群组是否存在
	GroupCheckExists(ctx context.Context, kind, uid string) (*tpl.BoolRes, error)
	// 批量添加群组
	GroupBatchAdd(ctx context.Context, groups []tpl.GroupBody) (*tpl.BoolRes, error)
	// 更新指定群组
	GroupUpdate(ctx context.Context, kind, uid string, body *urbssetting.GroupUpdateBody) (*tpl.GroupRes, error)
	// 删除指定群组
	GroupDelete(ctx context.Context, kind, uid string) (*tpl.BoolRes, error)
	// 读取群组成员列表，支持条件筛选
	GroupListMembers(ctx context.Context, args *tpl.GroupPaginationURL) (*tpl.GroupMembersRes, error)
	// 指定群组批量添加成员
	GroupBatchAddMembers(ctx context.Context, kind, uid string, users []string) (*tpl.BoolRes, error)
	// 指定群组根据条件清理成员
	GroupRemoveMembers(ctx context.Context, args *tpl.GroupMembersURL) (*tpl.BoolRes, error)
}
