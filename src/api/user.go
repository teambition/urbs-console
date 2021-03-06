package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util"
)

// User ..
type User struct {
	blls *bll.Blls
}

// List ..
func (a *User) List(ctx *gear.Context) error {
	req := tpl.Pagination{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	res, err := a.blls.User.List(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// RefreshCachedLables 强制更新 user 的 labels 缓存
func (a *User) RefreshCachedLables(ctx *gear.Context) error {
	req := tpl.UIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckSuperAdmin(ctx)
	if err != nil {
		return err
	}
	res, err := a.blls.User.RefreshCachedLables(ctx, req.UID)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// ListLables 返回 user 的 labels，按照 label 指派时间正序，支持分页
func (a *User) ListLables(ctx *gear.Context) error {
	req := new(tpl.UIDPaginationURL)
	if err := ctx.ParseURL(req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckSuperAdmin(ctx)
	if err != nil {
		return err
	}
	res, err := a.blls.User.ListLables(ctx, req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// ListSettings 返回 user 的 settings，按照 setting 设置时间正序，支持分页
func (a *User) ListSettings(ctx *gear.Context) error {
	req := tpl.UIDPaginationURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.User.ListSettings(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// ListSettingsUnionAll 返回 user 的 settings，按照 setting 设置时间反序，支持分页
// 包含了 user 从属的 group 的 settings
func (a *User) ListSettingsUnionAll(ctx *gear.Context) error {
	req := tpl.MySettingsQueryURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckSuperAdmin(ctx)
	if err != nil {
		return err
	}
	res, err := a.blls.User.ListSettingsUnionAll(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// CheckExists ..
func (a *User) CheckExists(ctx *gear.Context) error {
	req := tpl.UIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckSuperAdmin(ctx)
	if err != nil {
		return err
	}
	res, err := a.blls.User.CheckExists(ctx, req.UID)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// BatchAdd ..
func (a *User) BatchAdd(ctx *gear.Context) error {
	req := tpl.UsersBody{}
	if err := ctx.ParseBody(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckSuperAdmin(ctx)
	if err != nil {
		return err
	}
	res, err := a.blls.User.BatchAdd(ctx, req.Users)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// ListSettingsUnionAllClient 返回 user 的 settings，按照 setting 设置时间反序，支持分页
// 包含了 user 从属的 group 的 settings
func (a *User) ListSettingsUnionAllClient(ctx *gear.Context) error {
	req := tpl.MySettingsQueryURL{
		UID: util.GetUid(ctx),
	}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.User.ListSettingsUnionAll(ctx, &req)
	if err != nil {
		return err
	}
	for _, item := range res.Result {
		item.Desc = ""
		item.Release = 0
		item.LastValue = ""
		item.UpdatedAt = item.AssignedAt
	}
	return ctx.OkJSON(res)
}

// ListLablesForClient 返回 user 的 labels，按照 label 指派时间正序，支持分页
func (a *User) ListLablesForClient(ctx *gear.Context) error {
	req := new(tpl.UIDPaginationURL)
	if err := ctx.ParseURL(req); err != nil {
		return err
	}
	req.UID = util.GetUid(ctx)
	res, err := a.blls.User.ListLables(ctx, req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// ApplyRules ..
func (a *User) ApplyRules(ctx *gear.Context) error {
	req := &tpl.ProductURL{}
	if err := ctx.ParseURL(req); err != nil {
		return err
	}
	body := &tpl.ApplyRulesBody{}
	if err := ctx.ParseBody(body); err != nil {
		return err
	}
	res, err := a.blls.User.ApplyRules(ctx, req.Product, body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// ListSettingsUnionAllBackend 返回 user 的 settings，按照 setting 设置时间反序，支持分页
// 包含了 user 从属的 group 的 settings
func (a *User) ListSettingsUnionAllBackend(ctx *gear.Context) error {
	req := tpl.MySettingsQueryURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.User.ListSettingsUnionAll(ctx, &req)
	if err != nil {
		return err
	}
	for _, item := range res.Result {
		item.Desc = ""
		item.Release = 0
		item.LastValue = ""
		item.UpdatedAt = item.AssignedAt
	}
	return ctx.OkJSON(res)
}
