<template>
  <div class="menu-management">
    <div class="mgr-layout">
      <div class="left-tree">
        <t-card :bordered="false" class="tree-card">
          <div class="tree-header">
            <t-input v-model="treeFilter" placeholder="搜索菜单" clearable @press-enter="handleTreeFilter" @clear="handleTreeFilter">
              <template #suffix-icon>
                <search-icon size="14px" @click="handleTreeFilter" />
              </template>
            </t-input>
          </div>
          <div class="tree-body">
            <t-tree
              :data="menuTree"
              activatable
              :actived="activedKeys"
              expand-all
              @click="handleTreeClick"
            />
          </div>
        </t-card>
      </div>

      <div class="right-content">
        <t-card class="list-card-container" :bordered="false">
      <t-row justify="space-between" align="center" class="operation-row">
        <div class="left-operation-container">
          <t-button theme="primary" @click="handleAdd">
            <template #icon><add-icon /></template>
            添加菜单
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
            placeholder="搜索菜单名称或路径"
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
        :data="menuList"
        :columns="columns"
        row-key="id"
        vertical-align="middle"
        :hover="true"
        :pagination="pagination"
        :selected-row-keys="selectedRowKeys"
        :loading="dataLoading"
        :row-class-name="rowClassName"
        @select-change="handleSelectChange"
        @page-change="handlePageChange"
      >
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
      </div>
    </div>

    <!-- 添加/编辑菜单对话框 -->
    <t-dialog
      v-model:visible="dialogVisible"
      :header="dialogType === 'add' ? '添加菜单' : '编辑菜单'"
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
        <t-form-item label="父ID" name="parent_id">
          <t-input v-model="formData.parent_id" type="number" placeholder="父菜单ID，根为0" />
        </t-form-item>
        <t-form-item label="名称" name="name">
          <t-input v-model="formData.name" placeholder="请输入菜单名称" />
        </t-form-item>
        <t-form-item label="路径" name="path">
          <t-input v-model="formData.path" placeholder="例如 /mgr/user" />
        </t-form-item>
        <t-form-item label="组件" name="component">
          <t-input v-model="formData.component" placeholder="组件路径或 LAYOUT" />
        </t-form-item>
        <t-form-item label="图标" name="icon">
          <t-input v-model="formData.icon" placeholder="图标标识" />
        </t-form-item>
        <t-form-item label="排序" name="sort">
          <t-input v-model="formData.sort" type="number" placeholder="排序，数字" />
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
import { ref, onMounted, computed } from 'vue';
import { SearchIcon, AddIcon, EditIcon, DeleteIcon } from 'tdesign-icons-vue-next';
import { MessagePlugin, type FormInstanceFunctions, type PrimaryTableCol } from 'tdesign-vue-next';
import dayjs from 'dayjs';
import { getMenuList, addMenu, updateMenu, deleteMenu } from '@/api/menu';
import type { MenuItem, MenuCreateReq, MenuUpdateReq } from '@/api/model/menuModel';

const menuList = ref<MenuItem[]>([]);
const dataLoading = ref(false);
const selectedRowKeys = ref<(string | number)[]>([]);
const searchValue = ref('');

// tree related
const menuTree = ref<any[]>([]);
const treeFilter = ref('');
const activeNodeId = ref<number | null>(null);
const activedKeys = computed(() => (activeNodeId.value ? [activeNodeId.value] : []));
const selectedParentId = ref<number | null>(null);

const pagination = ref({
  defaultPageSize: 10,
  total: 0,
  current: 1,
  showPageSize: true,
  showJumper: true,
  pageSizeOptions: [10, 20, 50, 100]
});

const dialogVisible = ref(false);
const dialogType = ref<'add' | 'edit'>('add');
const confirmVisible = ref(false);
const deleteIdx = ref<number | string | null>(null);

const formRef = ref<FormInstanceFunctions>();
const formData = ref<any>({ id: 0, parent_id: 0, name: '', path: '', component: '', icon: '', sort: 0 });

const columns: PrimaryTableCol[] = [
  { colKey: 'row-select', type: 'multiple', width: 64, fixed: 'left' },
  { title: '名称', colKey: 'name', width: 180 },
  { title: '路径', colKey: 'path', width: 160 },
  { title: '组件', colKey: 'component', width: 160 },
  { title: '图标', colKey: 'icon', width: 100, align: 'center' },
  { title: '排序', colKey: 'sort', width: 80, align: 'center' },
  { title: '创建时间', colKey: 'created_at', width: 180, align: 'center' },
  { title: '操作', colKey: 'op', width: 140, fixed: 'right', align: 'center' }
];

const formRules = {
  name: [ { required: true, message: '名称不能为空' } ],
  path: [ { required: true, message: '路径不能为空' } ],
};

const formatTimestamp = (timestamp: number) => (timestamp ? dayjs.unix(timestamp).format('YYYY-MM-DD HH:mm:ss') : '-');

const rowClassName = ({ row }: { row: MenuItem }) => {
  if (selectedParentId.value && row.id === selectedParentId.value) {
    return 'menu-parent-row';
  }
  return '';
};

const fetchMenuList = async () => {
  dataLoading.value = true;
  try {
    const isTreeSelected = selectedParentId.value !== null;
    const resp = await getMenuList({
      page: isTreeSelected ? 1 : pagination.value.current,
      page_size: isTreeSelected ? 1000 : pagination.value.defaultPageSize,
      name: searchValue.value || undefined,
    });
    let list = resp.list || [];
    if (isTreeSelected) {
      const parent = list.find((i: MenuItem) => i.id === selectedParentId.value);
      const children = list.filter((i: MenuItem) => Number(i.parent_id || 0) === Number(selectedParentId.value));
      list = parent ? [parent, ...children] : children;
    }
    menuList.value = list;
    pagination.value.total = isTreeSelected ? list.length : (resp.total || 0);
  } catch (e) {
    console.error('获取菜单列表失败', e);
    MessagePlugin.error('获取菜单列表失败');
  } finally {
    dataLoading.value = false;
  }
};

const buildTree = (items: MenuItem[]) => {
  const map = new Map<number, any>();
  const roots: any[] = [];
  items.forEach((it) => {
    map.set(it.id, { ...it, value: it.id, label: it.name, children: [] });
  });
  map.forEach((node) => {
    const parentId = Number(node.parent_id || 0);
    if (parentId && map.has(parentId)) {
      map.get(parentId).children.push(node);
    } else {
      roots.push(node);
    }
  });
  return roots;
};

const fetchTree = async () => {
  try {
    const resp = await getMenuList({ page: 1, page_size: 1000 });
    const list = resp.list || [];
    const filtered = treeFilter.value ? list.filter((i: MenuItem) => i.name.includes(treeFilter.value) || i.path.includes(treeFilter.value)) : list;
    menuTree.value = buildTree(filtered as MenuItem[]);
  } catch (e) {
    console.error('获取菜单树失败', e);
  }
};

const handleTreeFilter = () => { fetchTree(); };

const handleTreeClick = ({ node }: { node: any }) => {
  const val = node.value;
  if (activeNodeId.value === val) {
    activeNodeId.value = null;
  } else {
    activeNodeId.value = val;
  }
  selectedParentId.value = activeNodeId.value;
  pagination.value.current = 1;
  fetchMenuList();
};

const handleSelectChange = (value: (string | number)[]) => { selectedRowKeys.value = value; };

const handlePageChange = (pageInfo: { current: number; pageSize: number }) => {
  pagination.value.current = pageInfo.current;
  pagination.value.defaultPageSize = pageInfo.pageSize;
  fetchMenuList();
};

const handleSearch = () => { pagination.value.current = 1; fetchMenuList(); };

const handleAdd = () => {
  dialogType.value = 'add';
  formData.value = { id: 0, parent_id: selectedParentId.value || 0, name: '', path: '', component: '', icon: '', sort: 0 };
  dialogVisible.value = true;
};

const handleEdit = (row: MenuItem) => {
  dialogType.value = 'edit';
  formData.value = { ...row };
  dialogVisible.value = true;
};

const handleDelete = (row: MenuItem) => { deleteIdx.value = row.id; confirmVisible.value = true; };

const handleBatchDelete = () => {
  if (!selectedRowKeys.value.length) { MessagePlugin.warning('请选择要删除的菜单'); return; }
  deleteIdx.value = null; confirmVisible.value = true;
};

const handleDialogConfirm = async () => {
  try {
    const valid = await formRef.value?.validate();
    if (!valid) return;
    if (dialogType.value === 'add') {
      await addMenu(formData.value);
      MessagePlugin.success('添加菜单成功');
    } else {
      await updateMenu(formData.value);
      MessagePlugin.success('编辑菜单成功');
    }
    dialogVisible.value = false;
    fetchMenuList();
    fetchTree();
  } catch (err: any) {
    const msg = err?.response?.data?.message || '操作失败';
    MessagePlugin.error(msg);
  }
};

const handleDialogClose = () => { formRef.value?.clearValidate(); };

const confirmBody = computed(() => {
  if (deleteIdx.value === null) {
    return `确定要删除选中的 ${selectedRowKeys.value.length} 个菜单吗？`;
  }
  const m = menuList.value.find((i) => i.id === deleteIdx.value);
  return m ? `确定要删除菜单 "${m.name}" 吗？` : '';
});

const onConfirmDelete = async () => {
  try {
    let ids: number[] = [];
    if (deleteIdx.value === null) {
      ids = selectedRowKeys.value.map(id => Number(id));
    } else {
      ids = [Number(deleteIdx.value)];
    }

    await deleteMenu(ids);

    if (selectedParentId.value && ids.includes(Number(selectedParentId.value))) {
      selectedParentId.value = null;
      activeNodeId.value = null;
    }

    if (deleteIdx.value === null) {
      selectedRowKeys.value = [];
      MessagePlugin.success('批量删除成功');
    } else {
      MessagePlugin.success('删除成功');
    }
    fetchMenuList();
    fetchTree();
  } catch (e: any) {
    const msg = e.response?.data?.message || '删除失败';
    MessagePlugin.error(msg);
  } finally {
    confirmVisible.value = false; deleteIdx.value = null;
  }
};

const onCancel = () => { confirmVisible.value = false; deleteIdx.value = null; };

onMounted(() => { fetchTree(); fetchMenuList(); });
</script>

<style scoped lang="less">
.menu-management {
  padding: 24px;
  background-color: var(--td-bg-color-container);
  min-height: 100%;
}

.mgr-layout {
  display: flex;
  gap: 24px;
  align-items: flex-start;
}

.left-tree {
  width: 280px;
  min-width: 220px;
  max-width: 34%;
}

.tree-card {
  height: 100%;
  :deep(.t-card__body) {
    padding: 12px;
  }
}

.tree-header { padding-bottom: 8px; }
.tree-body {
  max-height: calc(100vh - 240px);
  overflow: auto;
  padding-right: 6px;
}

.right-content { flex: 1 1 0; min-width: 320px; }
.list-card-container { :deep(.t-card__body) { padding: 20px; } }

.operation-row {
  margin-bottom: 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .left-operation-container {
    display:flex;
    align-items:center;
    gap:12px;
    .selected-count{ color: var(--td-text-color-secondary); font-size:14px;}
  }
  .search-input{ width:360px; }
}

:deep(.t-table) { .t-table__content { border-radius: var(--td-radius-medium); } }

@media screen and (max-width: 1000px) {
  .left-tree { width: 220px; min-width: 180px; }
  .mgr-layout { gap: 16px; }
  .operation-row .search-input { width: 240px; }
}

@media screen and (max-width: 760px) {
  .mgr-layout { flex-direction: column; gap: 16px; }
  .left-tree { width: 100%; min-width: auto; }
  .tree-card { order: 1 }
  .right-content { order: 2 }
  .operation-row { flex-direction: column; align-items: stretch; gap: 12px; }
  .operation-row .search-input { width: 100%; }
  .list-card-container { :deep(.t-card__body) { padding: 12px; } }
}

:deep(.menu-parent-row) {
  background-color: var(--td-bg-color-secondarycontainer);
  font-weight: bold;
  td {
    border-bottom: 2px solid var(--td-component-stroke) !important;
  }
}
</style>
