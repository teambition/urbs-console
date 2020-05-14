
import React, { useState, useEffect, useCallback } from 'react';
import { connect } from 'dva';
import { TableTitle, Setting as SettingTable, SettingModifyModal, SettingDetailModal } from '../../components'
import styleNames from './index.less';
import { SettingComponentProps, DEFAULT_PAGE_SIZE, PaginationParameters, Setting, FieldsValue } from '../../declare';

const Settings: React.FC<SettingComponentProps> = (props) => {
  const {
    dispatch,
    match,
    productSettingsList,
    settingNextPageToken,
    settingPrePageToken,
    settingPageTotal,
    history,
  } = props as any;
  const { params } = match;
  const productName = params.name;
  const [curentSetting, setCurentSetting] = useState<Setting>();
  const [pageSize, setPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [settingModifyVisible, changeSettingModifyVisible] = useState(false);
  const [settingDetailVisible, changeSettingDetailVisible] = useState(false);

  const fetchSettingList = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getProductSettings',
      payload: {
        productName,
        params,
        type,
      }
    })
  }, [dispatch, productName]);
  const fetchSettingLogList = useCallback((curentSetting: Setting) => {
    dispatch({
      type: 'products/getSettingLogs',
      payload: {
        product: productName,
        module: curentSetting?.module,
        setting: curentSetting?.name,
        params: {
          pageSize: 1000
        },
      }
    });
  }, [dispatch, productName]);
  useEffect(() => {
    fetchSettingList({
      pageSize,
    });
  }, [fetchSettingList, pageSize]);

  const handleOnRow = (record: Setting) => {
    return {
      onDoubleClick: () => {
        setCurentSetting(record);
        changeSettingDetailVisible(true);
        fetchSettingLogList(record);
      }
    };
  };
  const handleSettingModifyOk = (values: FieldsValue) => {
    if (curentSetting) {
      dispatch({
        type: 'products/updateProductSettings',
        payload: {
          params: values,
          productName,
          cb: (record: Setting) => {
            setCurentSetting(record);
            fetchSettingList({
              pageSize,
            }, 'del');
            changeSettingModifyVisible(false);
          },
        },
      });
    } else {
      dispatch({
        type: 'products/addProductSettings',
        payload: {
          params: values,
          productName,
          cb: () => {
            fetchSettingList({
              pageSize,
            }, 'del');
            changeSettingModifyVisible(false);
          },
        },
      });
    }
  };
  const handleOfflineSetting = () => {
    dispatch({
      type: 'products/offlineProductSettings',
      payload: {
        productName,
        module: curentSetting?.module,
        setting: curentSetting?.name,
        cb: () => {
          fetchSettingList({
            pageSize,
          }, 'del');
          changeSettingModifyVisible(false);
        },
      },
    });
  };
  const handlePlusClick = () => {
    setCurentSetting(undefined);
    changeSettingModifyVisible(true);
  };
  return (
    <div className={ styleNames.normal }>
      <TableTitle
        plusTitle="添加配置项"
        handlePlusClick={ handlePlusClick }
        handleSearch={ (value: string) => { fetchSettingList({pageSize, q: value}, 'del') } }
      />
      <SettingTable
        onRow={ handleOnRow }
        dataSource={ productSettingsList }
        paginationProps={
          {
            pageSize,
            total: settingPageTotal,
            pageSizeOptions: [10, 20, 30, 40],
            nextPageToken: settingNextPageToken,
            prePageToken: settingPrePageToken,
            onTokenChange: (type: string, token?: string) => {
              fetchSettingList({
                pageSize,
                pageToken: token,
              }, type);
            },
            onPageSizeChange: (size: number) => {
              setPageSize(size);
              fetchSettingList({
                pageSize: size,
              }, 'del');
            }
          }
        }
      />
      {/* 弹窗 */}
      {
        settingModifyVisible && <SettingModifyModal
          visible={ settingModifyVisible }
          isEdit={ !!curentSetting }
          defaultValue={ curentSetting }
          onOffline={ handleOfflineSetting }
          onOk={ handleSettingModifyOk }
          onCancel={ () => changeSettingModifyVisible(false) }
        />
      }
      {
        settingDetailVisible && (
          <SettingDetailModal
            visible={ settingDetailVisible }
            settingInfo={ curentSetting }
            title="配置项"
            product={ productName }
            onSettingEdit={ () => changeSettingModifyVisible(true) }
            onCancel={ () => changeSettingDetailVisible(false) }
            onGotoGroups={ () => history.push('/group') }
            onGotoUsers={ () => history.push('/user') }
          />
        )
      }
    </div>
  );
}

export default connect((state) => {
  return { ...(state as any).products };
})(Settings);
