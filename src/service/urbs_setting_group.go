package service

import (
	"context"
	"fmt"

	"github.com/mushroomsir/request"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
)

// GroupListLables ...
func (a *UrbsSetting) GroupListLables(ctx context.Context, args *tpl.GroupPaginationURL) (*tpl.MyLabelsRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/labels?skip=%d&pageSize=%d&pageToken=%s&q=%s&kind=%s", conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize, args.PageToken, args.Q, args.Kind)

	result := new(tpl.MyLabelsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupListSettings ...
func (a *UrbsSetting) GroupListSettings(ctx context.Context, args *tpl.MySettingsQueryURL) (*tpl.MySettingsRes, error) {
	path := "%s/v1/groups/%s/settings?skip=%d&pageSize=%d&pageToken=%s&q=%s"
	path += "&product=%s"
	path += "&module=%s"
	path += "&setting=%s"
	path += "&client=%s"
	path += "&channel=%s"
	path += "&kind=%s"
	url := fmt.Sprintf(path, conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize, args.PageToken, args.Q, args.Product, args.Module, args.Setting, args.Client, args.Channel, args.Kind)

	result := new(tpl.MySettingsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupList ...
func (a *UrbsSetting) GroupList(ctx context.Context, args *tpl.GroupsURL) (*tpl.GroupsRes, error) {
	url := fmt.Sprintf("%s/v1/groups?kind=%s&skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Kind, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.GroupsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupCheckExists ...
func (a *UrbsSetting) GroupCheckExists(ctx context.Context, kind, uid string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s:exists?kind=%s", conf.Config.UrbsSetting.Addr, uid, kind)

	result := new(tpl.BoolRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupBatchAdd ...
func (a *UrbsSetting) GroupBatchAdd(ctx context.Context, groups []tpl.GroupBody) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups:batch", conf.Config.UrbsSetting.Addr)

	body := new(tpl.GroupsBody)
	body.Groups = groups

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupUpdate ...
func (a *UrbsSetting) GroupUpdate(ctx context.Context, kind, uid string, body *urbssetting.GroupUpdateBody) (*tpl.GroupRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s?kind=%s", conf.Config.UrbsSetting.Addr, uid, kind)

	result := new(tpl.GroupRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupDelete ...
func (a *UrbsSetting) GroupDelete(ctx context.Context, kind, uid string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s?kind=%s", conf.Config.UrbsSetting.Addr, uid, kind)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupListMembers ...
func (a *UrbsSetting) GroupListMembers(ctx context.Context, args *tpl.GroupPaginationURL) (*tpl.GroupMembersRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/members?skip=%d&pageSize=%d&pageToken=%s&q=%s&kind=%s", conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize, args.PageToken, args.Q, args.Kind)

	result := new(tpl.GroupMembersRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupBatchAddMembers ...
func (a *UrbsSetting) GroupBatchAddMembers(ctx context.Context, kind, uid string, users []string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/members:batch?kind=%s", conf.Config.UrbsSetting.Addr, uid, kind)

	body := new(tpl.UsersBody)
	body.Users = users

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupRemoveMembers ...
func (a *UrbsSetting) GroupRemoveMembers(ctx context.Context, args *tpl.GroupMembersURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/members?syncLt=%d&user=%s&kind=%s", conf.Config.UrbsSetting.Addr, args.UID, args.SyncLt, args.User, args.Kind)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
