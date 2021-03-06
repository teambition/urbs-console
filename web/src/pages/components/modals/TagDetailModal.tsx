
import React, { useEffect, useState, useMemo, useCallback } from 'react';
import { Button, Input, Modal, Icon, message } from 'antd';
import { connect } from 'dva';
import { PublishRecord, ContentDetail, ContentTabs, UserGroup, Users, GrayscaleTagModifyModal, PublishTagModal } from '../';
import { DEFAULT_MODAL_WIDTH, TagDetailComponentProps, TagTabsKey, PaginationParameters, FieldsValue, UserPercentRule, DEFAULT_PAGE_SIZE } from '../../declare';
import styles from '../style/TagDetailModal.less';
import { formatTableTime } from '../../utils/format';

const DEFAULT_TITLE = '环境标签';

const TagDetailModal: React.FC<TagDetailComponentProps> = (props) => {
  const {
    product,
    onSettingEdit,
    onCancel,
    visible,
    title = DEFAULT_TITLE,
    labelInfo,
    dispatch,
    onGotoGroups,
    onGotoUsers,
    labelLogsList,
    labelGroupsList,
    labelGroupsNextPageToken,
    labelGroupsPrePageToken,
    labelGroupsPageTotal,
    labelUsersList,
    labelUsersNextPageToken,
    labelUsersPrePageToken,
    labelUsersPageTotal,
  } = props as any;
  const [labelGroupPageSize, changeLabelGroupPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [labelUserPageSize, changeLabelUserPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [userPercentRule, changeUserPercentRule] = useState<UserPercentRule>();
  const [newUserPercentRule, changeNewUserPercentRule] = useState<UserPercentRule>();
  const [childLabelUserPercentRule, changeChildLabelUserPercentRule] = useState<UserPercentRule>();
  const [tabsActiveKey, setTabsActiveKey] = useState(String(TagTabsKey.Publish));
  const [tabsSearchWord, setTabsSearchWord] = useState('');
  const [publishTagModalVisible, setPublishTagModalVisible] = useState(false);
  const [grayscaleTagModalVisible, setGrayscaleTagModalVisible] = useState(false);
  const [grayscaleTagCanEdit, setGrayscaleTagCanEdit] = useState(false);
  const fetchLabelLogs = useCallback(() => {
    dispatch({
      type: 'products/getLabelLogs',
      payload: {
        params: {
          pageSize: 100
        },
        product,
        label: labelInfo?.name,
      },
    });
  }, [dispatch, labelInfo, product]);
  const isChildLabel = labelInfo?.name.indexOf('-') > 0
  const fetchLabelGroups = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getLabelGroups',
      payload: {
        type,
        params,
        product,
        label: labelInfo?.name,
      },
    });
  }, [dispatch, labelInfo, product]);
  const fetchLabelUsers = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getLabelUsers',
      payload: {
        type,
        params,
        product,
        label: labelInfo?.name,
      },
    });
  }, [dispatch, labelInfo, product]);
  useEffect(() => {
    fetchLabelLogs();
    fetchLabelGroups({
      pageSize: labelGroupPageSize,
      q: tabsSearchWord,
    });
    fetchLabelUsers({
      pageSize: labelUserPageSize,
      q: tabsSearchWord,
    });
  }, [fetchLabelLogs, fetchLabelGroups, fetchLabelUsers, labelGroupPageSize, labelUserPageSize, tabsSearchWord]);
  useEffect(() => {
    dispatch({
      type: 'products/getPermission',
      payload: {
        cb: (canEdit: boolean) => {
          setGrayscaleTagCanEdit(!!canEdit);
        },
        params: {
          product,
          label: labelInfo?.name,
        }
      },
    })
  }, [dispatch, labelInfo, product]);
  const handleTabsActiveKeyChange = (activeKey: string) => {
    setTabsActiveKey(activeKey);
    setTabsSearchWord('');
    switch (activeKey) {
      case TagTabsKey.Publish:
        fetchLabelLogs();
        break;
      case TagTabsKey.Group:
        fetchLabelGroups({
          pageSize: labelGroupPageSize,
        }, 'del');
        break;
      case TagTabsKey.User:
        fetchLabelUsers({
          pageSize: labelUserPageSize,
        }, 'del');
        break;
      default:
        break;
    }
  };
  const handleTabsSearch = (searchWord: string) => {
    switch (tabsActiveKey) {
      case TagTabsKey.Group:
        fetchLabelGroups({
          pageSize: labelGroupPageSize,
          q: searchWord,
        });
        break;
      case TagTabsKey.User:
        fetchLabelUsers({
          pageSize: labelUserPageSize,
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
      type: 'products/getPublishRules',
      payload: {
        product,
        label: labelInfo?.name,
        cb: (rules?: UserPercentRule[]) => {
          if (!Array.isArray(rules)) return;
          for (const rule of rules) {
            if (rule?.kind === 'userPercent') {
              changeUserPercentRule(rule)
            } else if (rule?.kind === 'newUserPercent') {
              changeNewUserPercentRule(rule)
            } else if (rule?.kind === 'childLabelUserPercent') {
              changeChildLabelUserPercentRule(rule)
            }
          }
          changePublishTagModalVisible(true);
        }
      }
    });
  };

  const updateProductTagRule = (rule: UserPercentRule, values: FieldsValue) => {
    dispatch({
      type: 'products/updateProductTagRule',
      payload: {
        product,
        label: labelInfo?.name,
        rule: rule.hid,
        params: values,
        cb: () => {
          fetchLabelLogs();
          changePublishTagModalVisible(false);
        }
      },
    });
  }
  const handleOpenPublishTagModalOk = (values: FieldsValue) => {
    if (newUserPercentRule && values.kind === 'newUserPercent') {
      updateProductTagRule(newUserPercentRule, values)
    }
    if (childLabelUserPercentRule && values.kind === 'childLabelUserPercent') {
      updateProductTagRule(childLabelUserPercentRule, values)
    }
    if (userPercentRule && values.kind === 'userPercent') {
      updateProductTagRule(userPercentRule, values)
    }
    if (values.kind === 'batch' || (!newUserPercentRule && values.kind === 'newUserPercent') || (!userPercentRule && values.kind === 'userPercent') || (!childLabelUserPercentRule && values.kind === 'childLabelUserPercent')) {
      dispatch({
        type: 'products/publishProductTags',
        payload: {
          product,
          label: labelInfo?.name,
          params: values,
          cb: () => {
            fetchLabelLogs();
            changePublishTagModalVisible(false);
          }
        },
      });
    }
  };
  const handleOpenPublishTagModalCancel = () => {
    changePublishTagModalVisible(false);
  };
  const handleLabelLogReback = (hid: string) => {
    dispatch({
      type: 'products/recallLabelLogs',
      payload: {
        product,
        label: labelInfo?.name,
        hid: hid,
        cb: () => {
          fetchLabelLogs();
          message.success('撤回成功');
        },
      },
    });
  };
  const handleDeleteGroup = (uid: string) => {
    Modal.confirm({
      title: '操作不可逆，请再次确认',
      content: '确认删除该群组？',
      onOk: () => {
        dispatch({
          type: 'products/deleteLabelGroup',
          payload: {
            product,
            label: labelInfo?.name,
            uid,
            cb: () => {
              message.success('删除群组成功');
              fetchLabelGroups({
                pageSize: labelGroupPageSize,
                q: tabsSearchWord,
              }, 'del');
            },
          },
        });
      },
    });
  };
  const handleCleanUp = () => {
    Modal.confirm({
      title: '操作不可逆，请再次确认',
      content: '确认清空全部用户？',
      onOk: () => {
        dispatch({
          type: 'products/cleanUpLabel',
          payload: {
            product,
            label: labelInfo?.name,
            cb: () => {
              message.success('删除成功');
            },
          },
        });
      },
    });
  };

  const handleDeleteUser = (uid: string) => {
    Modal.confirm({
      title: '操作不可逆，请再次确认',
      content: '确认删除该用户？',
      onOk: () => {
        dispatch({
          type: 'products/deleteLabeUser',
          payload: {
            product,
            label: labelInfo?.name,
            uid,
            cb: () => {
              message.success('删除用户成功');
              fetchLabelUsers({
                pageSize: labelUserPageSize,
                q: tabsSearchWord,
              }, 'del');
            },
          },
        });
      },
    });
  };
  const renderModalTitle = () => {
    return (
      <div className={styles['tag-modal-title']}>
        <div>{title}</div>
        {
          grayscaleTagCanEdit && <div className={styles['tag-icon']}>
            <Icon type="setting" onClick={onSettingEdit}></Icon>
          </div>
        }
      </div>
    )
  };
  const tagTabsConfig = [{
    key: TagTabsKey.Publish,
    title: '发布记录',
    content: (
      <PublishRecord
        publishRecordList={labelLogsList}
        onReback={handleLabelLogReback}
        canRecall={grayscaleTagCanEdit}
      />
    ),
    action: grayscaleTagCanEdit && (
      <div>
        <Button
          type="link"
          icon="plus"
          block
          onClick={handleOpenPublishTagModal}
        >
          添加灰度发布
      </Button>
        <Button
          type="link"
          icon="delete"
          block
          onClick={handleCleanUp}
        >
          清空全部用户
        </Button>
      </div>
    ),
  }, {
    key: TagTabsKey.Group,
    title: '群组',
    content: (
      <UserGroup
        dataSource={labelGroupsList}
        hideColumns={['uid', 'desc']}
        paginationProps={
          {
            total: labelGroupsPageTotal,
            nextPageToken: labelGroupsNextPageToken,
            prePageToken: labelGroupsPrePageToken,
            pageSize: labelGroupPageSize,
            pageSizeOptions: [10, 20, 50, 100],
            onPageSizeChange: (size) => {
              changeLabelGroupPageSize(size);
              fetchLabelGroups({
                pageSize: size,
                q: tabsSearchWord,
              }, 'del');
            },
            onTokenChange: (type, token) => {
              fetchLabelGroups({
                pageSize: labelGroupPageSize,
                pageToken: token,
                q: tabsSearchWord,
              }, type);
            }
          }
        }
        onAction={(record) => {
          return {
            onDelete: () => {
              handleDeleteGroup(record.uid || record.group);
            },
          }
        }}
      />
    ),
    action: (
      <Input.Search
        value={tabsSearchWord}
        placeholder="请输入搜索关键字"
        onChange={handleTabsSearchWordChange}
        onSearch={handleTabsSearch}
        allowClear
      />
    ),
  }, {
    key: TagTabsKey.User,
    title: '用户',
    content: (
      <Users
        dataSource={labelUsersList}
        paginationProps={
          {
            total: labelUsersPageTotal,
            nextPageToken: labelUsersNextPageToken,
            prePageToken: labelUsersPrePageToken,
            pageSize: labelUserPageSize,
            pageSizeOptions: [10, 20, 50, 100],
            onPageSizeChange: (size) => {
              changeLabelUserPageSize(size);
              fetchLabelUsers({
                pageSize: size,
                q: tabsSearchWord,
              }, 'del');
            },
            onTokenChange: (type, token) => {
              fetchLabelUsers({
                pageSize: labelGroupPageSize,
                pageToken: token,
                q: tabsSearchWord,
              }, type);
            }
          }
        }
        onAction={(record) => {
          return {
            onDelete: () => {
              handleDeleteUser(record.user);
            },
          }
        }}
      />
    ),
    action: (
      <Input.Search
        value={tabsSearchWord}
        placeholder="请输入搜索关键字"
        onChange={handleTabsSearchWordChange}
        onSearch={handleTabsSearch}
        allowClear
      />
    ),
  }];
  const labelContentDetail = useMemo(() => {
    return labelInfo && ([
      {
        title: '名称',
        content: labelInfo.name,
      },
      {
        title: '所属产品',
        content: product,
      },
      {
        title: '负责人',
        content: Array.isArray(labelInfo.users) ? labelInfo.users.map(item => item.name).join(',') : '',
      },
      {
        title: '描述',
        content: labelInfo.desc,
      },
      {
        title: '版本通道',
        content: Array.isArray(labelInfo.channels) ? labelInfo.channels.join(',') : '空',
      },
      {
        title: '端类型',
        content: Array.isArray(labelInfo.clients) ? labelInfo.clients.join(',') : '空',
      },
      {
        title: '发布次数',
        content: labelInfo.release,
      },
      {
        title: '灰度进度',
        content: labelInfo.status,
      },
      {
        title: '创建时间',
        content: formatTableTime(labelInfo.createdAt),
      },
      {
        title: '更新时间',
        content: formatTableTime(labelInfo.updatedAt),
      },
    ]);
  }, [labelInfo, product]);
  return (
    <Modal width={DEFAULT_MODAL_WIDTH} title={renderModalTitle()} visible={visible} onCancel={onCancel} footer={null} >
      <ContentDetail content={labelContentDetail}></ContentDetail>
      <ContentTabs
        activeKey={tabsActiveKey}
        handleActiveKeyChange={handleTabsActiveKeyChange}
        tabs={tagTabsConfig}
      />
      {/* 弹窗 */}
      {
        publishTagModalVisible && <PublishTagModal
          title="发布环境标签"
          visible={publishTagModalVisible}
          onCancel={handleOpenPublishTagModalCancel}
          onOk={handleOpenPublishTagModalOk}
          label={labelInfo?.name}
          product={product}
          onGotoGroups={onGotoGroups}
          onGotoUsers={onGotoUsers}
          newUserPercentRule={newUserPercentRule}
          childLabelUserPercentRule={childLabelUserPercentRule}
          isChildLabel={isChildLabel}
          defauleRule={userPercentRule}
        ></PublishTagModal>
      }
      <GrayscaleTagModifyModal
        visible={grayscaleTagModalVisible}
        isEdit={true}
        onCancel={() => setGrayscaleTagModalVisible(false)}
        onOk={() => setGrayscaleTagModalVisible(false)}
      ></GrayscaleTagModifyModal>
    </Modal >
  );
};

export default connect((state) => {
  return {
    ...(state as any).products
  };
})(TagDetailModal);
