<script setup lang="ts">
import {defineProps, ref} from "vue";

const props = defineProps(['updateData', 'pageCount', 'size', 'page'])
// let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: props.size,
  page: props.page,
})

let dataCountOptions = [
  {
    label: '10条数据/页',
    value: 10,
  },
  {
    label: '20条数据/页',
    value: 20,
  },
  {
    label: '50条数据/页',
    value: 50,
  },
  {
    label: '100条数据/页',
    value: 100,
  },
]
</script>

<script lang="ts">
export default {
  name: 'FormSuffix'
}
</script>

<template>
  <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
    <n-pagination
        size="medium"
        v-model:page.number="dataSize.page"
        :page-count="props.pageCount"
        @update:page="props.updateData(dataSize.page, dataSize.pageSize)"
    />
    <n-select
        style="width: 160px; margin-left: 20px"
        v-model:value.number="dataSize.pageSize"
        size="small"
        :options="dataCountOptions"
        :remote="true"
        @update:value="dataSize.page = 1; props.updateData(dataSize.page, dataSize.pageSize)"
    />
  </div>
</template>

<style scoped lang="less">

</style>