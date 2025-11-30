<template>
  <div class="role-management">
    <t-card class="list-card-container" :bordered="false">
      <t-row justify="space-between" align="center" class="operation-row">
        <div class="left-operation-container">
          <t-button theme="primary" @click="handleAdd">
            <template #icon><add-icon /></template>
            添加角色
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
            placeholder="搜索角色编码或描述"
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

      <t-table
        :data="roleList"
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
        <template #status="{ row }">
          <t-tag :theme="row.status === 1 ? 'success' : 'danger'" variant="light">
            {{ row.status === 1 ? '启用' : '禁用' }}
          </t-tag>
        </template>

        <template #created_at="{ row }">
          {{ formatTimestamp(row.created_at) }}
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
    </t-card>

    <t-dialog v-model:visible="dialogVisible" :header="dialogType==='add'?'添加角色':'编辑角色'" :width="560" @confirm="handleDialogConfirm" @close="handleDialogClose">
      <t-form ref="formRef" :data="formData" :rules="formRules" label-width="100px" @submit="handleDialogConfirm">
        <t-form-item label="编码" name="code">
          <t-input v-model="formData.code" placeholder="例如 admin" />
        </t-form-item>
        <t-form-item label="描述" name="description">
          <t-input v-model="formData.description" placeholder="角色描述" />
        </t-form-item>
        <t-form-item label="状态" name="status">
          <t-select v-model="formData.status" placeholder="请选择状态">
            <t-option :value="1" label="启用" />
            <t-option :value="0" label="禁用" />
          </t-select>
        </t-form-item>
      </t-form>
    </t-dialog>

    <t-dialog v-model:visible="confirmVisible" header="确认删除" :body="confirmBody" @confirm="onConfirmDelete" @close="onCancel" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { SearchIcon, AddIcon, EditIcon, DeleteIcon } from 'tdesign-icons-vue-next';
import { MessagePlugin, type FormInstanceFunctions, type PrimaryTableCol, type FormRules } from 'tdesign-vue-next';
import dayjs from 'dayjs';
import { getRoleList, addRole, updateRole, deleteRole, type Role, type RoleCreateReq, type RoleUpdateReq } from '@/api/mgr/role';

const roleList = ref<Role[]>([]);
const dataLoading = ref(false);
const selectedRowKeys = ref<(string | number)[]>([]);
const searchValue = ref('');

const pagination = ref({ defaultPageSize: 10, total: 0, current: 1, showPageSize: true, showJumper: true, pageSizeOptions: [10,20,50,100] });

const dialogVisible = ref(false);
const dialogType = ref<'add'|'edit'>('add');
const confirmVisible = ref(false);
const deleteIdx = ref<number | string | null>(null);

const formRef = ref<FormInstanceFunctions>();
const formData = ref<RoleCreateReq | RoleUpdateReq>({ id: 0 as unknown as number, code: '', description: '', status: 1 });

const columns: PrimaryTableCol[] = [
  { colKey: 'row-select', type: 'multiple', width: 64, fixed: 'left' },
  { title: '编码', colKey: 'code', width: 180 },
  { title: '描述', colKey: 'description', width: 240 },
  { title: '状态', colKey: 'status', width: 100, align: 'center' },
  { title: '创建时间', colKey: 'created_at', width: 180, align: 'center' },
  { title: '操作', colKey: 'op', width: 140, fixed: 'right', align: 'center' }
];

const formRules: FormRules = { code: [{ required: true, message: '编码不能为空' }], description: [] };

const formatTimestamp = (timestamp: number): string => {
  return timestamp ? dayjs.unix(timestamp).format('YYYY-MM-DD HH:mm:ss') : '-';
};

const fetchRoleList = async () => {
  dataLoading.value = true;
  try {
    const resp = await getRoleList({ page: pagination.value.current, page_size: pagination.value.defaultPageSize, keyword: searchValue.value });
    roleList.value = resp.list || [];
    pagination.value.total = resp.total || 0;
  } catch (e) { MessagePlugin.error('获取角色列表失败'); console.error(e); } finally { dataLoading.value = false; }
};

const handleSelectChange = (v: (string|number)[]) => selectedRowKeys.value = v;
const handlePageChange = (pageInfo: { current: number; pageSize: number }) => { pagination.value.current = pageInfo.current; pagination.value.defaultPageSize = pageInfo.pageSize; fetchRoleList(); };
const handleSearch = () => { pagination.value.current = 1; fetchRoleList(); };
const handleAdd = () => { dialogType.value = 'add'; formData.value = { id:0, code:'', description:'', status:1 }; dialogVisible.value = true; };
const handleEdit = (row: Role) => { dialogType.value='edit'; formData.value = { ...(row as RoleUpdateReq), id: row.id }; dialogVisible.value = true; };
const handleDelete = (row: Role) => { deleteIdx.value = row.id; confirmVisible.value = true; };
const handleBatchDelete = () => { if(!selectedRowKeys.value.length){ MessagePlugin.warning('请选择要删除的角色'); return; } deleteIdx.value = null; confirmVisible.value = true; };

const handleDialogConfirm = async () => {
  try {
    const valid = await formRef.value?.validate();
    if (!valid) return;
    if (dialogType.value === 'add') { await addRole(formData.value); MessagePlugin.success('添加角色成功'); }
    else { await updateRole(formData.value); MessagePlugin.success('编辑角色成功'); }
    dialogVisible.value = false; fetchRoleList();
  } catch (e:any) { const msg = e.response?.data?.message || '操作失败'; MessagePlugin.error(msg); }
};

const handleDialogClose = () => formRef.value?.clearValidate();

const confirmBody = computed(() => { if (deleteIdx.value === null) return `确定要删除选中的 ${selectedRowKeys.value.length} 个角色吗？`; const r = roleList.value.find(i=>i.id===deleteIdx.value); return r?`确定要删除角色 "${r.code}" 吗？` : ''; });

const onConfirmDelete = async () => {
  try {
    if (deleteIdx.value === null) { const ids = selectedRowKeys.value.map(id=>Number(id)); await deleteRole(ids); selectedRowKeys.value = []; MessagePlugin.success('批量删除成功'); }
    else { const ids=[Number(deleteIdx.value)]; await deleteRole(ids); MessagePlugin.success('删除成功'); }
    fetchRoleList();
  } catch (e:any) { const msg = e.response?.data?.message || '删除失败'; MessagePlugin.error(msg); } finally { confirmVisible.value=false; deleteIdx.value=null; }
};

const onCancel = () => { confirmVisible.value=false; deleteIdx.value=null; };

onMounted(()=>{ fetchRoleList(); });
</script>

<style scoped lang="less">
.role-management { padding:24px; background-color: var(--td-bg-color-container); min-height:100%; }
.list-card-container { :deep(.t-card__body){ padding:24px; } }
.operation-row { margin-bottom:16px; .left-operation-container{ display:flex; align-items:center; gap:12px; .selected-count{ color: var(--td-text-color-secondary); font-size:14px; } } .search-input{ width:360px; } }
:deep(.t-table){ .t-table__content{ border-radius: var(--td-radius-medium); } }
@media screen and (max-width:768px){ .operation-row{ flex-direction:column; align-items:flex-start; gap:12px; .search-input{ width:100%; } } }
</style>
