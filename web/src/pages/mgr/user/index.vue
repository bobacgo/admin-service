<template>
  <div class="user-management">
    <t-card class="list-card-container" :bordered="false">
      <t-row justify="space-between" align="center" class="operation-row">
        <div class="left-operation-container">
          <t-button theme="primary" @click="handleAdd">
            <template #icon><add-icon /></template>
            添加用户
          </t-button>
          <t-button 
            variant="base" 
            theme="danger" 
            :disabled="!selectedRowKeys.length" 
            @click="handleBatchDelete"
          >
            <template #icon><delete-icon /></template>
            批量删除
          </t-button>
          <span v-if="selectedRowKeys.length > 0" class="selected-count">
            已选择 {{ selectedRowKeys.length }} 项
          </span>
        </div>
        <div class="search-input">
          <t-input
            v-model="searchValue"
            placeholder="搜索用户账号、手机号或邮箱"
            clearable
            @clear="handleSearch"
            @press-enter="handleSearch"
          >
            <template #suffix-icon>
              <search-icon size="16px" @click="handleSearch" />
            </template>
          </t-input>
        </div>
      </t-row>

      <div class="table-wrapper">
        <t-table
          :data="userList"
          :columns="columns"
          row-key="id"
          vertical-align="middle"
          :hover="true"
          :pagination="pagination"
          :selected-row-keys="selectedRowKeys"
          :loading="dataLoading"
          @select-change="handleSelectChange"
          @page-change="handlePageChange"
        >
          <template #role_codes="{ row }">
            <t-space size="4px" break-line v-if="row.role_codes">
              <t-tag 
                v-for="code in row.role_codes.split(',')" 
                :key="code"
                theme="primary" 
                variant="light"
              >
                {{ code.trim() }}
              </t-tag>
            </t-space>
            <span v-else class="text-secondary">-</span>
          </template>
          
          <template #status="{ row }">
            <t-tag :theme="row.status === 1 ? 'success' : 'danger'" variant="light">
              {{ formatStatus(row.status) }}
            </t-tag>
          </template>
          
          <template #register_at="{ row }">
            {{ formatTimestamp(row.register_at) }}
          </template>
          
          <template #login_at="{ row }">
            {{ formatTimestamp(row.login_at) }}
          </template>
          
          <template #op="{ row }">
            <t-space>
              <t-link theme="primary" hover="color" @click="handleEdit(row)">
                <template #icon><edit-icon /></template>
                编辑
              </t-link>
              <t-link theme="danger" hover="color" @click="handleDelete(row)">
                <template #icon><delete-icon /></template>
                删除
              </t-link>
            </t-space>
          </template>
        </t-table>
      </div>
    </t-card>

    <!-- 添加/编辑用户对话框 -->
    <t-dialog
      v-model:visible="dialogVisible"
      :header="dialogType === 'add' ? '添加用户' : '编辑用户'"
      :width="600"
      :confirm-on-enter="false"
      @confirm="handleDialogConfirm"
      @close="handleDialogClose"
    >
      <t-form
        ref="formRef"
        :data="formData"
        :rules="formRules"
        label-width="100px"
        @submit="handleDialogConfirm"
      >
        <t-form-item label="账号" name="account">
          <t-input v-model="formData.account" placeholder="请输入账号" />
        </t-form-item>
        <t-form-item label="手机号" name="phone">
          <t-input v-model="formData.phone" placeholder="请输入手机号" />
        </t-form-item>
        <t-form-item label="邮箱" name="email">
          <t-input v-model="formData.email" placeholder="请输入邮箱" />
        </t-form-item>
        <t-form-item label="状态" name="status">
          <t-select v-model="formData.status" placeholder="请选择状态">
            <t-option :value="1" label="启用" />
            <t-option :value="0" label="禁用" />
          </t-select>
        </t-form-item>
        <t-form-item label="角色" name="roleIds">
          <t-select 
            v-model="formData.roleIds" 
            multiple
            placeholder="请选择角色"
            clearable
          >
            <t-option 
              v-for="role in roleList" 
              :key="role.id" 
              :value="role.id" 
              :label="role.code"
            />
          </t-select>
        </t-form-item>
        <t-form-item v-if="dialogType === 'add'" label="密码" name="password">
          <t-input v-model="formData.password" type="password" placeholder="请输入密码" />
        </t-form-item>
      </t-form>
    </t-dialog>

    <!-- 删除确认对话框 -->
    <t-dialog
      v-model:visible="confirmVisible"
      header="确认删除"
      :body="confirmBody"
      @confirm="onConfirmDelete"
      @close="onCancel"
    />
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { 
  SearchIcon, 
  AddIcon, 
  EditIcon, 
  DeleteIcon 
} from 'tdesign-icons-vue-next';
import { 
  MessagePlugin, 
  FormInstanceFunctions, 
  type PrimaryTableCol 
} from 'tdesign-vue-next';
import dayjs from 'dayjs';
import { getUserList, addUser, updateUser, deleteUser, type User, type UserAddReq, type UserUpdateReq } from '@/api/mgr/user'
import { getRoleList, type Role } from '@/api/mgr/role'
import type { IdsReq } from '@/api/model';

// 响应式数据
const userList = ref<User[]>([]);
const roleList = ref<Role[]>([]);
const dataLoading = ref(false);
const selectedRowKeys = ref<(string | number)[]>([]);
const searchValue = ref('');

// 分页配置
const pagination = ref({
  defaultPageSize: 10,
  total: 0,
  current: 1,
  showPageSize: true,
  showJumper: true,
  pageSizeOptions: [10, 20, 50, 100]
});

// 对话框控制
const dialogVisible = ref(false);
const dialogType = ref<'add' | 'edit'>('add');
const confirmVisible = ref(false);
const deleteIdx = ref<number | string | null>(null);

// 表单相关
const formRef = ref<FormInstanceFunctions>();
const formData = ref({
  id: 0,
  account: '',
  phone: '',
  email: '',
  status: 1,
  password: '',
  roleIds: [] as number[]
});

// 表格列定义
const columns: PrimaryTableCol[] = [
  { colKey: 'row-select', type: 'multiple', width: 64, fixed: 'left' },
  {
    title: '账号',
    colKey: 'account',
    width: 120,
    fixed: 'left',
    ellipsis: true
  },
  {
    title: '手机号',
    colKey: 'phone',
    width: 120,
    ellipsis: true
  },
  {
    title: '邮箱',
    colKey: 'email',
    width: 180,
    ellipsis: true
  },
  {
    title: '角色编码',
    colKey: 'role_codes',
    width: 120,
    ellipsis: true
  },
  {
    title: '状态',
    colKey: 'status',
    width: 80,
    align: 'center'
  },
  {
    title: '注册时间',
    colKey: 'register_at',
    width: 160,
    align: 'center'
  },
  {
    title: '注册IP',
    colKey: 'register_ip',
    width: 120,
    ellipsis: true
  },
  {
    title: '最后登录时间',
    colKey: 'login_at',
    width: 160,
    align: 'center'
  },
  {
    title: '登录IP',
    colKey: 'login_ip',
    width: 120,
    ellipsis: true
  },
  {
    title: '操作',
    colKey: 'op',
    width: 120,
    fixed: 'right',
    align: 'center'
  }
];

// 表单验证规则
const formRules = {
  account: [
    { required: true, message: '账号不能为空' },
    { min: 3, max: 20, message: '账号长度必须在3-20个字符之间' }
  ],
  phone: [
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号格式' }
  ],
  password: [
    { required: true, message: '密码不能为空' },
    { min: 6, max: 20, message: '密码长度必须在6-20个字符之间' }
  ]
};

// 格式化函数
const formatTimestamp = (timestamp: number): string => {
  return timestamp ? dayjs.unix(timestamp).format('YYYY-MM-DD HH:mm:ss') : '-';
};

const formatStatus = (status: number): string => {
  return status === 1 ? '启用' : '禁用';
};

// 获取用户列表 - 使用新的API接口
const fetchUserList = async () => {
  dataLoading.value = true;
  try {
    const response = await getUserList({
      page: pagination.value.current,
      page_size: pagination.value.defaultPageSize,
      keyword: searchValue.value
    });
    
    const { list, total } = response;
    userList.value = list || [];
    pagination.value.total = total || 0;
  } catch (error) {
    console.error('获取用户列表失败:', error);
    MessagePlugin.error('获取用户列表失败');
  } finally {
    dataLoading.value = false;
  }
};

// 获取角色列表
const fetchRoleList = async () => {
  try {
    const response = await getRoleList({
      page: 1,
      page_size: 100
    });
    roleList.value = response.list || [];
  } catch (error) {
    console.error('获取角色列表失败:', error);
  }
};

// 事件处理函数
const handleSelectChange = (value: (string | number)[]) => {
  selectedRowKeys.value = value;
};

const handlePageChange = (pageInfo: any) => {
  pagination.value.current = pageInfo.current;
  pagination.value.defaultPageSize = pageInfo.pageSize;
  fetchUserList();
};

const handleSearch = () => {
  pagination.value.current = 1;
  fetchUserList();
};

// 添加用户
const handleAdd = () => {
  dialogType.value = 'add';
  formData.value = {
    id: 0,
    account: '',
    phone: '',
    email: '',
    status: 1,
    password: '',
    roleIds: []
  };
  dialogVisible.value = true;
};

// 编辑用户
const handleEdit = (row: User) => {
  dialogType.value = 'edit';
  // 将 role_codes 字符串转换为 roleIds 数组（需要通过 API 获取或映射）
  formData.value = {
    id: row.id,
    account: row.account,
    phone: row.phone,
    email: row.email,
    status: row.status,
    password: '',
    roleIds: []
  };
  dialogVisible.value = true;
};

// 删除用户
const handleDelete = (row: User) => {
  deleteIdx.value = row.id;
  confirmVisible.value = true;
};

// 批量删除
const handleBatchDelete = () => {
  if (selectedRowKeys.value.length === 0) {
    MessagePlugin.warning('请选择要删除的用户');
    return;
  }
  deleteIdx.value = null; // 标记为批量删除
  confirmVisible.value = true;
};

// 对话框确认
const handleDialogConfirm = async () => {
  try {
    const valid = await formRef.value?.validate();
    if (!valid) return;

    if (dialogType.value === 'add') {
      // 添加用户
      const addData: UserAddReq = {
        account: formData.value.account,
        password: formData.value.password,
        email: formData.value.email,
        phone: formData.value.phone,
        status: formData.value.status,
        roleIds: formData.value.roleIds.length > 0 ? formData.value.roleIds : undefined,
      };
      await addUser(addData);
      MessagePlugin.success('添加用户成功');
    } else {
      // 编辑用户
      const updateData: UserUpdateReq = {
        id: formData.value.id,
        account: formData.value.account,
        email: formData.value.email,
        phone: formData.value.phone,
        status: formData.value.status,
        roleIds: formData.value.roleIds.length > 0 ? formData.value.roleIds : undefined,
      };
      await updateUser(updateData);
      MessagePlugin.success('编辑用户成功');
    }

    dialogVisible.value = false;
    fetchUserList();
  } catch (error: any) {
    const message = error.response?.data?.message || '操作失败';
    MessagePlugin.error(message);
  }
};

const handleDialogClose = () => {
  formRef.value?.clearValidate();
};

// 删除确认
const confirmBody = computed(() => {
  if (deleteIdx.value === null) {
    return `确定要删除选中的 ${selectedRowKeys.value.length} 个用户吗？`;
  }
  const user = userList.value.find(u => u.id === deleteIdx.value);
  return user ? `确定要删除用户 "${user.account}" 吗？删除后该用户的所有信息将被清空且无法恢复。` : '';
});

const onConfirmDelete = async () => {
  try {
    if (deleteIdx.value === null) {
      // 批量删除
      const ids = selectedRowKeys.value.map(id => Number(id))
      await deleteUser(ids);
      selectedRowKeys.value = [];
      MessagePlugin.success('批量删除成功');
    } else {
      // 单个删除
      const ids = [Number(deleteIdx.value)]
      await deleteUser(ids);
      MessagePlugin.success('删除成功');
    }
    fetchUserList();
  } catch (error: any) {
    const message = error.response?.data?.message || '删除失败';
    MessagePlugin.error(message);
  } finally {
    confirmVisible.value = false;
    deleteIdx.value = null;
  }
};

const onCancel = () => {
  confirmVisible.value = false;
  deleteIdx.value = null;
};

// 生命周期
onMounted(() => {
  fetchRoleList();
  fetchUserList();
});
</script>

<style scoped lang="less">
.user-management {
  padding: 24px;
  background-color: var(--td-bg-color-container);
  min-height: 100%;
}

.list-card-container {
  :deep(.t-card__body) {
    padding: 24px;
  }
}

.operation-row {
  margin-bottom: 16px;

  .left-operation-container {
    display: flex;
    align-items: center;
    gap: 12px;

    .selected-count {
      color: var(--td-text-color-secondary);
      font-size: 14px;
    }
  }

  .search-input {
    width: 360px;
  }
}

.table-wrapper {
  overflow-x: auto;
  overflow-y: hidden;
  border-radius: var(--td-radius-medium);
  
  :deep(.t-table) {
    min-width: 100%;
    
    .t-table__content {
      border-radius: var(--td-radius-medium);
    }
  }
}

.text-secondary {
  color: var(--td-text-color-secondary);
}

:deep(.t-table) {
  .t-table__content {
    border-radius: var(--td-radius-medium);
  }
}

@media screen and (max-width: 768px) {
  .operation-row {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;

    .search-input {
      width: 100%;
    }
  }
}
</style>