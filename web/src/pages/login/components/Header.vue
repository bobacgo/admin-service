<template>
  <header class="login-header">
    <logo-full-icon class="logo" />
    <div class="operations-container">
      <t-button theme="default" shape="square" variant="text" @click="navToGitHub">
        <t-icon name="logo-github" class="icon" />
      </t-button>
      <t-button theme="default" shape="square" variant="text" @click="navToHelper">
        <t-icon name="help-circle" class="icon" />
      </t-button>
      <t-button theme="default" shape="square" variant="text" @click="toggleSettingPanel">
        <t-icon name="setting" class="icon" />
      </t-button>
    </div>
  </header>
</template>
<script setup lang="ts">
import LogoFullIcon from '@/assets/assets-logo-full.svg?component';
import { useSettingStore } from '@/store';

const settingStore = useSettingStore();
const toggleSettingPanel = () => {
  settingStore.updateConfig({
    showSettingPanel: true,
  });
};

const navToGitHub = () => {
  window.open('https://github.com/tencent/tdesign-vue-next-starter');
};

const navToHelper = () => {
  window.open('http://tdesign.tencent.com/starter/docs/get-started');
};
</script>
<style lang="less" scoped>
.login-header {
  padding: 0 var(--td-comp-paddingLR-xl);
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: var(--td-comp-size-xxxl);
  z-index: 4;

  .logo {
    width: 178px;
    height: var(--td-comp-size-s);
    filter: drop-shadow(0 8px 28px rgba(155,80,255,0.18));
  }

  .operations-container {
    display: flex;
    align-items: center;

    .t-button {
      margin-left: var(--td-comp-margin-l);
      background: rgba(255,255,255,0.02);
      border-radius: 8px;
      border: 1px solid rgba(155,80,255,0.06);
      color: rgba(220,230,255,0.9);
      transition: transform .18s ease, box-shadow .18s ease;

      .icon {
        color: rgba(180,210,255,0.9);
        filter: drop-shadow(0 6px 14px rgba(99,123,255,0.12));
      }

      &:hover {
        transform: translateY(-3px);
        box-shadow: 0 8px 26px rgba(99,123,255,0.12);
      }
    }
  }
}
</style>
