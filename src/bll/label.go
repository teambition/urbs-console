package bll

import (
	"context"

	"github.com/mushroomsir/tcc"
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util"
)

// Label ...
type Label struct {
	services *service.Services
	daos     *dao.Daos

	operationLog *OperationLog
	urbsAcAcl    *UrbsAcAcl
	group        *Group
}

// ListGroups ...
func (a *Label) ListGroups(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelGroupsInfoRes, error) {
	return a.services.UrbsSetting.LabelListGroups(ctx, args)
}

// DeleteGroup ...
func (a *Label) DeleteGroup(ctx context.Context, args *tpl.ProductLabelUIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.LabelDeleteGroup(ctx, args)
}

// ListUsers ...
func (a *Label) ListUsers(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelUsersInfoRes, error) {
	return a.services.UrbsSetting.LabelListUsers(ctx, args)
}

// DeleteUser ...
func (a *Label) DeleteUser(ctx context.Context, args *tpl.ProductLabelUIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.LabelDeleteUser(ctx, args)
}

// Create ...
func (a *Label) Create(ctx context.Context, product string, args *tpl.LabelBody) (*tpl.LabelInfoRes, error) {
	aclObject := product + args.Name
	err := a.urbsAcAcl.Update(ctx, &tpl.UidsBody{Uids: args.Uids}, aclObject)
	if err != nil {
		return nil, err
	}
	res, err := a.services.UrbsSetting.LabelCreate(ctx, product, args)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = a.urbsAcAcl.FindUsersByObject(ctx, aclObject)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// List 返回产品下的标签列表
func (a *Label) List(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.LabelsInfoRes, error) {
	labels, err := a.services.UrbsSetting.LabelList(ctx, args)
	if err != nil {
		return nil, err
	}
	objects := make([]string, len(labels.Result))
	for i, label := range labels.Result {
		objects[i] = args.Product + label.Name
	}
	subjects, err := a.urbsAcAcl.FindUsersByObjects(ctx, objects)
	if err != nil {
		return nil, err
	}
	for _, label := range labels.Result {
		label.Users = subjects[args.Product+label.Name]
	}
	return labels, nil
}

// Update ...
func (a *Label) Update(ctx context.Context, product, label string, body *tpl.LabelUpdateBody) (*tpl.LabelInfoRes, error) {
	aclObject := product + label
	err := a.urbsAcAcl.Update(ctx, body.UidsBody, product+label)
	if err != nil {
		return nil, err
	}
	res, err := a.services.UrbsSetting.LabelUpdate(ctx, product, label, body)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = a.urbsAcAcl.FindUsersByObject(ctx, aclObject)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Offline 下线标签
func (a *Label) Offline(ctx context.Context, product, label string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.LabelOffline(ctx, product, label)
}

// Assign 把标签批量分配给用户或群组
func (a *Label) Assign(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.UsersGroupsBody) (*tpl.LabelReleaseInfoRes, error) {
	a.group.AddUserAndOrg(ctx, body.Users, body.Groups)
	groupBody := &urbssetting.UsersGroupsBody{
		Users:  body.Users,
		Groups: parseGroupUIDs(body.Groups),
		Value:  body.Value,
	}
	res, err := a.services.UrbsSetting.LabelAssign(ctx, args.Product, args.Label, groupBody)
	if err != nil {
		return nil, err
	}
	object := args.Product + args.Label
	logContent := &dto.OperationLogContent{
		Users:   body.Users,
		Groups:  body.Groups,
		Desc:    body.Desc,
		Value:   body.Value,
		Release: res.Result.Release,
	}
	err = a.operationLog.Add(ctx, object, actionCreate, logContent)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete 物理删除标签
func (a *Label) Delete(ctx context.Context, product, label string) (*tpl.BoolRes, error) {
	res, err := a.services.UrbsSetting.LabelDelete(ctx, product, label)
	if err != nil {
		return nil, err
	}
	err = a.daos.UrbsAcAcl.DeleteByObject(ctx, product+label)
	if err != nil {
		logger.Err(ctx, err.Error())
	}
	return res, nil
}

// Recall 批量撤销对用户或群组设置的产品环境标签
func (a *Label) Recall(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.RecallBody) (*tpl.BoolRes, error) {
	logID := service.HIDToID(body.HID, "log")
	log, err := a.daos.OperationLog.FindOneByID(ctx, logID)
	if err != nil {
		return nil, err
	}
	release := getRelease(log.Content)
	if release < 1 {
		return nil, gear.ErrBadRequest.WithMsgf("invalid release %d", release)
	}
	body.Release = release

	value := &labelRecallReq{
		Args: args,
		Body: body,
	}
	tx := a.services.TCC.NewTransaction(TccSettingRecall)
	msgSql := tx.TryPlan(tcc.ObjToJSON(value))

	err = a.daos.OperationLog.TxDelete(ctx, logID, msgSql)
	if err != nil {
		return nil, err
	}

	recallRes, err := a.services.UrbsSetting.LabelRecall(ctx, args, body)
	if err != nil {
		tx.Confirm()
		return nil, err
	}
	tx.Cancel()

	logger.Info(ctx, "labelRecall", "operator", util.GetUid(ctx), "log", log.String)
	return recallRes, nil
}

// CreateRule ...
func (a *Label) CreateRule(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.LabelRuleBody) (*tpl.LabelRuleInfoRes, error) {
	res, err := a.services.UrbsSetting.LabelCreateRule(ctx, args, body)
	if err != nil {
		return nil, err
	}

	object := args.Product + args.Label
	logContent := &dto.OperationLogContent{
		Desc:    body.Desc,
		Percent: &body.Rule.Value,
		Kind:    body.Kind,
	}
	err = a.operationLog.Add(ctx, object, actionCreate, logContent)
	if err != nil {
		logger.Err(ctx, "labelCreateRuleLog", "logContent", logContent, "object", object)
	}

	hc := &dto.HookRule{
		Product:  args.Product,
		Label:    args.Label,
		Kind:     body.Kind,
		Percent:  &body.Rule.Value,
		Desc:     body.Desc,
		Operator: util.GetUid(ctx),
	}
	content := &thrid.HookSendReq{
		Event:   service.EventRuleCreate,
		Content: hc.Marshal(),
	}
	a.services.Hook.SendAsync(ctx, content)

	return res, nil
}

// ListRules ...
func (a *Label) ListRules(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelRulesInfoRes, error) {
	return a.services.UrbsSetting.LabelListRule(ctx, args)
}

// UpdateRule ...
func (a *Label) UpdateRule(ctx context.Context, args *tpl.ProductLabelHIDURL, body *tpl.LabelRuleBody) (*tpl.LabelRuleInfoRes, error) {
	res, err := a.services.UrbsSetting.LabelUpdateRule(ctx, args, body)
	if err != nil {
		return nil, err
	}

	object := args.Product + args.Label
	logContent := &dto.OperationLogContent{
		Desc:    body.Desc,
		Percent: &body.Rule.Value,
		Kind:    body.Kind,
	}
	err = a.operationLog.Add(ctx, object, actionUpdate, logContent)
	if err != nil {
		logger.Err(ctx, "labelUpdateRuleLog", "logContent", logContent, "object", object)
	}

	hc := &dto.HookRule{
		Product:  args.Product,
		Label:    args.Label,
		Kind:     body.Kind,
		Percent:  &body.Rule.Value,
		Desc:     body.Desc,
		Operator: util.GetUid(ctx),
	}
	content := &thrid.HookSendReq{
		Event:   service.EventRuleUpdate,
		Content: hc.Marshal(),
	}
	a.services.Hook.SendAsync(ctx, content)

	return res, nil
}

// DeleteRule ...
func (a *Label) DeleteRule(ctx context.Context, args *tpl.ProductLabelHIDURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.LabelDeleteRule(ctx, args)
}

// CleanUp ...
func (a *Label) CleanUp(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.BoolRes, error) {
	res, err := a.services.UrbsSetting.LabelCleanUp(ctx, args)
	if err != nil {
		return nil, err
	}

	object := args.Product + args.Label
	logContent := &dto.OperationLogContent{
		Desc: actionCleanup,
	}
	err = a.operationLog.Add(ctx, object, actionCleanup, logContent)
	if err != nil {
		logger.Err(ctx, "labelCleanUpLog", "logContent", logContent, "object", object)
	}

	hc := &dto.HookCleanup{
		Product:  args.Product,
		Label:    args.Label,
		Operator: util.GetUid(ctx),
	}
	content := &thrid.HookSendReq{
		Event:   service.EventCleanup,
		Content: hc.Marshal(),
	}
	a.services.Hook.SendAsync(ctx, content)

	return res, nil
}
