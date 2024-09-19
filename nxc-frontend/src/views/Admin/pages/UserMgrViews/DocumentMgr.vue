<script setup lang="ts">
import {ref} from "vue";
import {MdEditor} from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import type {DrawerPlacement, FormInst, useMessage} from 'naive-ui'
import instance from "@/axios/index";

const apiAddrStore = useApiAddrStore();
const placement = ref<DrawerPlacement>('right')
const active = ref(false)
const activate = (place: DrawerPlacement) => {
  active.value = true
  placement.value = place
}
const themeStore = useThemeStore();
const text = ref('# Hello Editor');
// 表单
const formRef = ref<FormInst | null>(null)
const message = useMessage()
let formValue = ref({
  doc: {
    title: '',
    category: '',
    language: '',
    md_text: ''
  }
})

let rules = {
  doc: {
    title: {
      required: true,
      message: '文档标题是必填的',
      trigger: 'blur'
    },
    category: {
      required: false,
      trigger: 'blur',
    },
    language: {
      required: false,
      trigger: 'blur'
    }
  }
}

let cancelAdd = () => {
  active.value = false
  formValue.value.doc.title = ''
  formValue.value.doc.category = ''
  formValue.value.doc.language = ''
  formValue.value.doc.md_text = ''

}

let submitDoc = async () => {
  if (formValue.value.doc.title !== '') {
    active.value = false
    console.log(formValue.value)
    let {data} = await instance.post(apiAddrStore.apiAddr.admin.addNewDocument, {
      title: formValue.value.doc.title,
      category: formValue.value.doc.category,
      language: formValue.value.doc.language,
      md_text: formValue.value.doc.md_text,
    })
    if (data.code === 200) {
      console.log(data)
      message.success('添加新的文档成功')
    } else {
      message.error('添加失败')

    }
  } else {
    message.error('文档标题不能为空')
  }
}


onMounted
(() => {
  themeStore.menuSelected = 'doc-manager'
  themeStore.contentPath = '/admin/dashboard/document';
})

</script>

<script lang="ts">
export default {
  name: 'DocumentMgr',
}
</script>

<template>
  <div class="root">
    <n-card class="body-card" hoverable :embedded="true" title="知识库管理">

      <n-button type="primary" @click="activate('right')">添加文档</n-button>

      <!--      <MdEditor v-model="text" :theme="themeStore.enableDarkMode?'dark':'light'"/>-->
    </n-card>

    <n-drawer v-model:show="active" width="80%" :placement="placement">
      <n-drawer-content title="新增文档">
        <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
        >
          <n-form-item label="标题" path="doc.title">
            <n-input v-model:value="formValue.doc.title" placeholder="输入文档标题"/>
          </n-form-item>
          <n-form-item label="分类" path="doc.category">
            <n-input v-model:value="formValue.doc.category" placeholder="输入文档分类"/>
          </n-form-item>
          <n-form-item label="语言" path="doc.language">
            <n-input v-model:value="formValue.doc.language" placeholder="选择文档语言"/>
          </n-form-item>
          <MdEditor style="background-color: rgba(0,0,0,0.0)" v-model="formValue.doc.md_text" :theme="themeStore.enableDarkMode?'dark':'light'"/>
        </n-form>
        <div style="display: flex; flex-direction: row; justify-content: right; margin-top: 20px">
          <n-button style="margin-right: 15px; width: 80px" @click="cancelAdd" secondary type="primary">取消</n-button>
          <n-button style="width: 80px" type="primary" @click="submitDoc">提交</n-button>
        </div>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
}
</style>