<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { defineProps, ref, computed, watch, defineEmits } from "vue";
import { type CountOption } from "@/types";

const { t } = useI18n();

// 定义 Props 和 Emits
const props = defineProps<{
  dataSize: { page: number; pageSize: number }; // 双向绑定分页数据
  pageCount: number; // 双向绑定总页数
  animated: boolean; // 双向绑定动画状态
  updateData: () => void; // 更新数据的方法
}>();

const emit = defineEmits([
  "update:data-size", // 更新 dataSize 的事件
  "update:page-count", // 更新 pageCount 的事件
  "update:animated", // 更新 animated 的事件
]);

// 定义分页数据
const dataSize = ref({ page: props.dataSize.page, pageSize: props.dataSize.pageSize });

// 动态分页选项，支持国际化
const dataCountOptions = computed<CountOption[]>(() => [
  { label: t("pagination.perPage10"), value: 10 },
  { label: t("pagination.perPage20"), value: 20 },
  { label: t("pagination.perPage50"), value: 50 },
  { label: t("pagination.perPage100"), value: 100 },
]);

// 监听 props.dataSize，同步到 dataSize
watch(
    () => props.dataSize,
    (newValue) => {
      dataSize.value = { ...newValue };
    },
    { immediate: true, deep: true }
);

// 同步分页数据到父组件
watch(
    dataSize,
    (newValue) => {
      emit("update:data-size", newValue);
    },
    { deep: true }
);
</script>

<template>
  <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
    <!-- 分页 -->
    <n-pagination
        size="medium"
        v-model:page="dataSize.page"
        :page-count="props.pageCount"
        @update:page="(newPage: number) => {
        dataSize.page = newPage;
        emit('update:data-size', { ...dataSize });
        emit('update:animated', false);
        props.updateData();
      }"
    />
    <!-- 每页数量选择 -->
    <n-select
        style="width: 260px; margin-left: 20px;"
        v-model:value="dataSize.pageSize"
        size="small"
        :options="dataCountOptions"
        @update:value="(newSize: number) => {
        dataSize.pageSize = newSize;
        dataSize.page = 1;
        emit('update:data-size', { ...dataSize });
        emit('update:animated', false);
        props.updateData();
      }"
    />
  </div>
</template>