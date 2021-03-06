package bll

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/service/mock_service"
	"github.com/teambition/urbs-console/src/tpl"
)

func TestGroup(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usMock := mock_service.NewMockUrbsSettingInterface(ctrl)
	groupMock := mock_service.NewMockGroupMemberInterface(ctrl)

	group := &Group{services: service.NewServices(testDB), daos: testDaos}
	group.services.UrbsSetting = usMock
	group.services.GroupMember = groupMock

	t.Run("batchAddMember error", func(t *testing.T) {
		require := require.New(t)

		uid := tpl.RandUID()

		// 1
		now := time.Now().Unix()
		groupUpdateBody := new(urbssetting.GroupUpdateBody)
		groupUpdateBody.SyncAt = &now

		usMock.EXPECT().GroupUpdate(nil, dto.GroupOrgKind, uid, groupUpdateBody).Return(nil, errors.New("urbs-setting error"))

		err := group.BatchAddMember(nil, dto.GroupOrgKind, uid)
		require.Equal("urbs-setting error", err.Error())

		// 2
		mockReturn := new(tpl.GroupRes)
		usMock.EXPECT().GroupUpdate(nil, dto.GroupOrgKind, uid, groupUpdateBody).Return(mockReturn, nil)
		groupMock.EXPECT().List(nil, uid, "", 1000).Return(nil, errors.New("thrid group-list error"))

		err = group.BatchAddMember(nil, dto.GroupOrgKind, uid)
		require.Equal("thrid group-list error", err.Error())

		// 3
		usMock.EXPECT().GroupUpdate(nil, dto.GroupOrgKind, uid, groupUpdateBody).Return(mockReturn, nil)

		userID := tpl.RandUID()
		groupReturn := &thrid.ListGroupMembersResp{
			Members: []thrid.Member{{UID: userID}},
		}
		groupMock.EXPECT().List(nil, uid, "", 1000).Return(groupReturn, nil)

		usMock.EXPECT().GroupBatchAddMembers(nil, dto.GroupOrgKind, uid, []string{userID}).Return(nil, errors.New("urbs-setting BatchAddMember error"))

		err = group.BatchAddMember(nil, dto.GroupOrgKind, uid)
		require.Equal("urbs-setting BatchAddMember error", err.Error())
	})

	t.Run("batchAddMember", func(t *testing.T) {
		require := require.New(t)

		uid := tpl.RandUID()

		now := time.Now().Unix()
		groupUpdateBody := new(urbssetting.GroupUpdateBody)
		groupUpdateBody.SyncAt = &now
		mockReturn := new(tpl.GroupRes)
		usMock.EXPECT().GroupUpdate(nil, dto.GroupOrgKind, uid, groupUpdateBody).Return(mockReturn, nil)

		userID := tpl.RandUID()
		groupReturn := &thrid.ListGroupMembersResp{
			Members: []thrid.Member{{UID: userID}},
		}
		groupMock.EXPECT().List(nil, uid, "", 1000).Return(groupReturn, nil)

		usMock.EXPECT().GroupBatchAddMembers(nil, dto.GroupOrgKind, uid, []string{userID}).Return(nil, nil)

		args := new(tpl.GroupMembersURL)
		args.UID = uid
		args.SyncLt = now
		usMock.EXPECT().GroupRemoveMembers(nil, args).Return(nil, nil)

		err := group.BatchAddMember(nil, dto.GroupOrgKind, uid)
		require.Nil(err)
	})
}
