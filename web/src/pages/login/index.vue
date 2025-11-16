<template>
  <div class="login-wrapper">
    <div class="stars" />
    <login-header />

    <div class="login-container">
      <div class="pattern-grid" />
      <div class="title-container">
        <h1 class="title margin-no">{{ t('pages.login.loginTitle') }}</h1>
        <h1 class="title">TDesign Starter</h1>
        <div class="sub-title">
          <p class="tip">{{ type === 'register' ? t('pages.login.existAccount') : t('pages.login.noAccount') }}</p>
          <p class="tip" @click="switchType(type === 'register' ? 'login' : 'register')">
            {{ type === 'register' ? t('pages.login.signIn') : t('pages.login.createAccount') }}
          </p>
          <p class="time-greeting">{{ greeting }}</p>
        </div>
      </div>

      <login v-if="type === 'login'" />
      <register v-else @register-success="switchType('login')" />
      <tdesign-setting />
    </div>

    <footer class="copyright">Copyright @ 2021-2025 Tencent. All Rights Reserved</footer>
  </div>
</template>
<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue';

import TdesignSetting from '@/layouts/setting.vue';
import { t } from '@/locales';

import LoginHeader from './components/Header.vue';
import Login from './components/Login.vue';
import Register from './components/Register.vue';

defineOptions({
  name: 'LoginIndex',
});
const type = ref('login');
const switchType = (val: string) => {
  type.value = val;
};

// time-based friendly greeting
const hour = ref(new Date().getHours());
let hourTimer: number | undefined;
onMounted(() => {
  hourTimer = window.setInterval(() => {
    hour.value = new Date().getHours();
  }, 60 * 1000);
});
onUnmounted(() => {
  if (hourTimer) clearInterval(hourTimer);
});

const greeting = computed(() => {
  const h = hour.value;
  if (h >= 5 && h < 9) return '早安，愿你有个美好的一天。';
  if (h >= 9 && h < 12) return '上午好，工作顺利～';
  if (h >= 12 && h < 14) return '中午好，记得吃午饭，休息一下。';
  if (h >= 14 && h < 18) return '下午好，保持高效但别忘了喝水。';
  if (h >= 18 && h < 22) return '晚上好，别忘了放松和陪伴家人。';
  return '深夜了，不要太忙碌了，注意休息。';
});
</script>
<style lang="less" scoped>
@import './index.less';
</style>
