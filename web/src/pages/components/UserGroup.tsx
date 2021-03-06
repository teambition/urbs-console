import React from 'react';
import { Table } from 'antd';
import { TableComponentProps, Group, ActionEventListeners } from '../declare';
import { formatTableTime } from '../utils/format';
import { Pagination } from './';

const UserGroup: React.FC<TableComponentProps<Group>> = (props) => {
  const { paginationProps, onAction, hideColumns } = props;
  const columns = [{
    title: '群组',
    dataIndex: 'group',
    key: 'group',
  }, {
    title: '群组',
    dataIndex: 'uid',
    key: 'uid',
  }, {
    title: '类型',
    dataIndex: 'kind',
    key: 'kind',
  }, {
    title: '描述',
    dataIndex: 'desc',
    key: 'desc',
  }, {
    title: '成员数量',
    dataIndex: 'status',
    key: 'status',
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    key: 'createdAt',
    render: (time: string) => {
      return formatTableTime(time);
    },
  }, {
    title: '分配时间',
    dataIndex: 'assignedAt',
    key: 'assignedAt',
    render: (time: string) => {
      return formatTableTime(time);
    },
  }, {
    dataIndex: 'action',
    key: 'action',
    width: 'auto',
    render: (_, record: Group) => {
      const { onDelete } = onAction ? onAction(record) : ({} as ActionEventListeners);
      return <a onClick={onDelete}>移除</a>
    }
  }];
  const generateTableColumns = () => {
    if (hideColumns) {
      return columns.filter(item => !hideColumns.includes(item.key));
    }
    return columns;
  };
  return (
    <div>
      <Table rowKey={(record) => (record.uid || record.group)} {...props} columns={generateTableColumns()} pagination={false}></Table>
      {
        paginationProps && (<Pagination {...paginationProps}></Pagination>)
      }
    </div>
  );
};

export default UserGroup;