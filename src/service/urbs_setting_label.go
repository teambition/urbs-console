package service

import (
	"context"
	"fmt"

	"github.com/mushroomsir/request"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
)

// LabelList ...
func (a *UrbsSetting) LabelList(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.LabelsInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.LabelsInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelCreate ...
func (a *UrbsSetting) LabelCreate(ctx context.Context, product string, args *tpl.LabelBody) (*tpl.LabelInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels", conf.Config.UrbsSetting.Addr, product)

	result := new(tpl.LabelInfoRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(args).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelUpdate ...
func (a *UrbsSetting) LabelUpdate(ctx context.Context, product string, label string, body *tpl.LabelUpdateBody) (*tpl.LabelInfoRes, error) {

	url := fmt.Sprintf("%s/v1/products/%s/labels/%s", conf.Config.UrbsSetting.Addr, product, label)

	result := new(tpl.LabelInfoRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelDelete ...
func (a *UrbsSetting) LabelDelete(ctx context.Context, product string, label string) (*tpl.BoolRes, error) {

	url := fmt.Sprintf("%s/v1/products/%s/labels/%s", conf.Config.UrbsSetting.Addr, product, label)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelOffline ...
func (a *UrbsSetting) LabelOffline(ctx context.Context, product string, label string) (*tpl.BoolRes, error) {

	url := fmt.Sprintf("%s/v1/products/%s/labels/%s:offline", conf.Config.UrbsSetting.Addr, product, label)

	result := new(tpl.BoolRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelAssign ...
func (a *UrbsSetting) LabelAssign(ctx context.Context, product string, label string, body *urbssetting.UsersGroupsBody) (*tpl.LabelReleaseInfoRes, error) {

	url := fmt.Sprintf("%s/v2/products/%s/labels/%s:assign", conf.Config.UrbsSetting.Addr, product, label)

	result := new(tpl.LabelReleaseInfoRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelRecall ...
func (a *UrbsSetting) LabelRecall(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.RecallBody) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s:recall", conf.Config.UrbsSetting.Addr, args.Product, args.Label)

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelListUsers ...
func (a *UrbsSetting) LabelListUsers(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelUsersInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s/users?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Label, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.LabelUsersInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelListGroups ...
func (a *UrbsSetting) LabelListGroups(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelGroupsInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s/groups?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Label, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.LabelGroupsInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelCreateRule ...
func (a *UrbsSetting) LabelCreateRule(ctx context.Context, args *tpl.ProductLabelURL, body *tpl.LabelRuleBody) (*tpl.LabelRuleInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s/rules", conf.Config.UrbsSetting.Addr, args.Product, args.Label)

	result := new(tpl.LabelRuleInfoRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelUpdateRule ...
func (a *UrbsSetting) LabelUpdateRule(ctx context.Context, args *tpl.ProductLabelHIDURL, body *tpl.LabelRuleBody) (*tpl.LabelRuleInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s/rules/%s", conf.Config.UrbsSetting.Addr, args.Product, args.Label, args.HID)

	result := new(tpl.LabelRuleInfoRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelDeleteRule ...
func (a *UrbsSetting) LabelDeleteRule(ctx context.Context, args *tpl.ProductLabelHIDURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s/rules/%s", conf.Config.UrbsSetting.Addr, args.Product, args.Label, args.HID)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelListRule ...
func (a *UrbsSetting) LabelListRule(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.LabelRulesInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s/rules", conf.Config.UrbsSetting.Addr, args.Product, args.Label)

	result := new(tpl.LabelRulesInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelDeleteUser ...
func (a *UrbsSetting) LabelDeleteUser(ctx context.Context, args *tpl.ProductLabelUIDURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s/users/%s", conf.Config.UrbsSetting.Addr, args.Product, args.Label, args.UID)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelDeleteGroup ...
func (a *UrbsSetting) LabelDeleteGroup(ctx context.Context, args *tpl.ProductLabelUIDURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s/groups/%s?kind=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Label, args.UID, args.Kind)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelCleanUp ...
func (a *UrbsSetting) LabelCleanUp(ctx context.Context, args *tpl.ProductLabelURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels/%s:cleanup", conf.Config.UrbsSetting.Addr, args.Product, args.Label)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
