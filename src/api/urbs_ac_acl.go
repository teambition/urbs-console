package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/tpl"
)

// UrbsAcAcl ..
type UrbsAcAcl struct {
	blls *bll.Blls
}

// Add 添加权限
func (a *UrbsAcAcl) Add(ctx *gear.Context) error {
	req := tpl.UrbsAcAclURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	body := tpl.UrbsAcAclAddReq{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.AddByReq(ctx, &req, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(struct{}{})
}
