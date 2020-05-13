
import React, { useEffect, useState, useMemo, useCallback } from 'react';
import { Button, Input, Modal, Icon } from 'antd';
import { connect } from 'dva';
import { PublishRecord, ContentDetail, ContentTabs, UserGroup, Users, GrayscaleTagModifyModal, PublishTagModal } from '../';
import { SettingDetailComponentProps, TagTabsKey, PaginationParameters, FieldsValue, UserPercentRule, DEFAULT_PAGE_SIZE } from '../../declare';
import styles from '../style/TagDetailModal.less';
import { formatTableTime } from '../../utils/format';

const SettingDetailModal: React.FC<SettingDetailComponentProps> = (props) => {
  const {
    product,
    onSettingEdit,
    onCancel,
    visible,
    title,
    settingInfo,
    dispatch,
    onGotoGroups,
    onGotoUsers,
    settingLogsList,
    settingGroupsList,
    settingGroupsNextPageToken,
    settingGroupsPrePageToken,
    settingGroupsPageTotal,
    settingUsersList,
    settingUsersNextPageToken,
    settingUsersPrePageToken,
    settingUsersPageTotal,
  } = props;
  const [settingGroupPageSize, changeSettingGroupPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [settingUserPageSize, changeSettingUserPageSize] = useState(DEFAULT_PAGE_SIZE);

  const [userPercentRule, changeUserPercentRule] = useState<UserPercentRule>();
  const [tabsActiveKey, setTabsActiveKey] = useState(String(TagTabsKey.Publish));
  const [tabsSearchWord, setTabsSearchWord] = useState('');
  const [publishTagModalVisible, setPublishTagModalVisible] = useState(false);
  const [grayscaleTagModalVisible, setGrayscaleTagModalVisible] = useState(false);
  const fetchSettingLogs = useCallback(() => {
    dispatch({
      type: 'products/getSettingLogs',
      payload: {
        params: {
          pageSize: 100
        },
        product,
        module: settingInfo?.module,
        setting: settingInfo?.name,
      },
    });
  }, [dispatch, product, settingInfo]);
  const fetchSettingGroups = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getSettingGroups',
      payload: {
        type,
        params,
        product,
        module: settingInfo?.module,
        setting: settingInfo?.name,
      },
    });
  }, [dispatch, product, settingInfo]);
  const fetchSettingUsers = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getSettingUsers',
      payload: {
        type,
        params,
        product,
        module: settingInfo?.module,
        setting: settingInfo?.name,
      },
    });
  }, [dispatch, product, settingInfo]);
  useEffect(() => {
    fetchSettingLogs();
    fetchSettingGroups({
      pageSize: settingGroupPageSize,
    });
    fetchSettingUsers({
      pageSize: settingUserPageSize,
    });
  }, [fetchSettingLogs, fetchSettingGroups, fetchSettingUsers, settingGroupPageSize, settingUserPageSize]);
  const handleTabsActiveKeyChange = (activeKey: string) => {
    setTabsActiveKey(activeKey);
    setTabsSearchWord('');
    switch(activeKey) {
      case TagTabsKey.Publish:
        fetchSettingLogs();
        break;
      case TagTabsKey.Group:
        fetchSettingGroups({
          pageSize: settingGroupPageSize,
        }, 'del');
        break;
      case TagTabsKey.User:
        fetchSettingUsers({
          pageSize: settingUserPageSize,
        }, 'del');
        break;
      default:
        break;
    }
  };
  const handleTabsSearch = (searchWord: string) => {
    switch(tabsActiveKey) {
      case TagTabsKey.Group:
        fetchSettingGroups({
          pageSize: settingGroupPageSize,
          q: searchWord,
        });
        break;
      case TagTabsKey.User:
        fetchSettingUsers({
          pageSize: settingUserPageSize,
          q: searchWord,
        });
        break;
      default:
        break;
    }
  };
  const handleTabsSearchWordChange = (e: React.ChangeEvent) => {
    const nativeEvent = e.nativeEvent;
    const target = nativeEvent.target || nativeEvent.srcElement;
    setTabsSearchWord((target as any).value);
  };
  const changePublishTagModalVisible = (visible: boolean) => {
    setPublishTagModalVisible(visible);
  };
  const handleOpenPublishTagModal = () => {
    dispatch({
      type: 'products/getPublishSettingRules',
      payload: {
        product,
        module: settingInfo?.module,
        setting: settingInfo?.name,
        cb: (rule?: UserPercentRule) => {
          changeUserPercentRule(rule);
          changePublishTagModalVisible(true);
        }
      }
    });
  };
  const handleOpenPublishTagModalOk = (values: FieldsValue) => {
    if (userPercentRule && values.kind === 'userPercent') {
      dispatch({
        type: 'products/updateProductSettingRule',
        payload: {
          product,
          module: settingInfo?.module,
          setting: settingInfo?.name,
          rule: userPercentRule.hid,
          params: values,
          cb: () => {
            fetchSettingLogs();
            changePublishTagModalVisible(false);
          }
        },
      });
    } else {
      dispatch({
        type: 'products/publishProductSettings',
        payload: {
          product,
          module: settingInfo?.module,
          setting: settingInfo?.name,
          params: values,
          cb: () => {
            fetchSettingLogs();
            changePublishTagModalVisible(false);
          }
        },
      });
    }
  };
  const handleOpenPublishTagModalCancel = () => {
    changePublishTagModalVisible(false);
  };
  const renderModalTitle = () => {
    return (
      <div className={ styles['tag-modal-title'] }>
        <div>{ title }</div>
        <div>
          <Icon type="setting" onClick={ onSettingEdit }></Icon>
        </div>
      </div>
    )
  };
  const tagTabsConfig = [{
    key: TagTabsKey.Publish,
    title: '发布记录',
    content: (
      <PublishRecord
        publishRecordList={ settingLogsList }
      />
    ),
    action: (
      <Button
        type="link"
        icon="plus"
        block
        onClick={ handleOpenPublishTagModal }
      >
        添加灰度发布
      </Button>
    ),
  }, {
    key: TagTabsKey.Group,
    title: '群组',
    content: (
      <UserGroup
        dataSource={ settingGroupsList }
        hideColumns={ ['syncAt'] }
        paginationProps={
          {
            total: settingGroupsPageTotal,
            nextPageToken: settingGroupsNextPageToken,
            prePageToken: settingGroupsPrePageToken,
            pageSize: settingGroupPageSize,
            pageSizeOptions: [10, 20, 30, 40],
            onPageSizeChange: (size) => {
              changeSettingGroupPageSize(size);
              fetchSettingGroups({
                pageSize: size,
              }, 'del');
            },
            onTokenChange: (type, token) => {
              fetchSettingGroups({
                pageSize: settingGroupPageSize,
                pageToken: token,
              }, type);
            }
          }
        }
      />
    ),
    action: (
      <Input.Search
        value={ tabsSearchWord }
        placeholder="请输入搜索关键字"
        onChange={ handleTabsSearchWordChange }
        onSearch={ handleTabsSearch }
      />
    ),
  }, {
    key: TagTabsKey.User,
    title: '用户',
    content: (
      <Users
        dataSource={ settingUsersList }
        hideColumns={ ['syncAt'] }
        paginationProps={
          {
            total: settingUsersPageTotal,
            nextPageToken: settingUsersNextPageToken,
            prePageToken: settingUsersPrePageToken,
            pageSize: settingUserPageSize,
            pageSizeOptions: [10, 20, 30, 40],
            onPageSizeChange: (size) => {
              changeSettingUserPageSize(size);
              fetchSettingUsers({
                pageSize: size,
              }, 'del');
            },
            onTokenChange: (type, token) => {
              fetchSettingUsers({
                pageSize: settingGroupPageSize,
                pageToken: token,
              }, type);
            }
          }
        }
      />
    ),
    action: (
      <Input.Search
        value={ tabsSearchWord }
        placeholder="请输入搜索关键字"
        onChange={ handleTabsSearchWordChange }
        onSearch={ handleTabsSearch }
      />
    ),
  }];
  const labelContentDetail = useMemo(() => {
    return settingInfo && ([
      {
        title: '名称',
        content: settingInfo.name,
      },
      {
        title: '所属产品',
        content: product,
      },
      {
        title: '所属模块',
        content: settingInfo.module,
      },
      {
        title: '负责人',
        content: Array.isArray(settingInfo.users) && settingInfo.users.length ? settingInfo.users.map(item => item.name).join(',') : '',
      },
      {
        title: '描述',
        content: settingInfo.desc,
      },
      {
        title: '版本通道',
        content: Array.isArray(settingInfo.channels) && settingInfo.channels.length ? settingInfo.channels.join(',') : '空',
      },
      {
        title: '端类型',
        content: Array.isArray(settingInfo.clients) && settingInfo.clients.length ? settingInfo.clients.join(',') : '空',
      },
      {
        title: '可选值',
        content: Array.isArray(settingInfo.values) && settingInfo.values.length ? settingInfo.values.join(',') : '空',
      },
      {
        title: '发布次数',
        content: settingInfo.release,
      },
      {
        title: '灰度进度',
        content: settingInfo.status,
      },
      {
        title: '创建时间',
        content: formatTableTime(settingInfo.createdAt),
      },
      {
        title: '更新时间',
        content: formatTableTime(settingInfo.updatedAt),
      },
    ]);
  }, [settingInfo, product]);
  return (
    <Modal title={ renderModalTitle() } visible={ visible } onCancel={ onCancel } footer={ null }>
      <ContentDetail content={ labelContentDetail }></ContentDetail>
      <ContentTabs
        activeKey={ tabsActiveKey }
        handleActiveKeyChange={ handleTabsActiveKeyChange }
        tabs={ tagTabsConfig }
      />
      {/* 弹窗 */}
      {
        publishTagModalVisible && <PublishTagModal
          title="发布配置项"
          visible={ publishTagModalVisible }
          onCancel={ handleOpenPublishTagModalCancel }
          onOk={ handleOpenPublishTagModalOk }
          label={ settingInfo?.name }
          module={ settingInfo?.module }
          product={ product }
          grayscale={ settingInfo?.values }
          onGotoGroups={ onGotoGroups }
          onGotoUsers={ onGotoUsers }
          defauleRule={ userPercentRule }
        ></PublishTagModal>
      }
      <GrayscaleTagModifyModal
        visible={ grayscaleTagModalVisible }
        isEdit={ true }
        onCancel={ () => setGrayscaleTagModalVisible(false)}
        onOk={ () => setGrayscaleTagModalVisible(false)}
      ></GrayscaleTagModifyModal>
    </Modal>
  );
};

export default connect((state) => {
  return {
    ...(state as any).products
  };
})(SettingDetailModal);