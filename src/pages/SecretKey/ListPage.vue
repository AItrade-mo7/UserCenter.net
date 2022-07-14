<script setup lang="ts">
import { GetOkxKeyList, DelOkxkey } from '@/api/okxKey';
import { DateFormat } from '@/utils/filters';
import { getUrlParams, cloneDeep } from '@/utils/tools';
import { defineAsyncComponent } from 'vue';
import AuthModal from '@/lib/AuthModal';

const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));

let list = $ref([]);

const GetList = async () => {
  const res = await GetOkxKeyList();
  if (res.Code > 0) {
    list = res.Data || [];
  }
};
const UrlParam = getUrlParams();

GetList();

const delKey = (item) => {
  const data = cloneDeep(item);
  AuthModal({
    IsPassword: true,
    async OkBack(param) {
      return DelOkxkey({
        OkxKeyID: data.OkxKeyID,
        Password: param.Password,
      }).then((res) => {
        if (res.Code > 0) {
          window.location.reload();
        }
      });
    },
  });
};
</script>

<template>
  <PageTitle>SecretKey</PageTitle>
  <div class="PageWrapper SecretKey">
    <div v-if="list.length" class="SecretKey_wrapper">
      <n-card
        v-for="item in list"
        :key="item.OkxKeyID"
        embedded
        hoverable
        size="small"
        :class="{ Active: UrlParam.id == item.OkxKeyID }"
      >
        <n-button class="SecretKey_del" strong secondary circle size="small" type="error" @click="delKey(item)">
          <template #icon>
            <XIcon name="DeleteOutlined" />
          </template>
        </n-button>
        <div class="SecretKey__item">
          <span class="SecretKey__label"> Name </span>
          <span class="SecretKey__val">
            {{ item.Name }}
          </span>
        </div>
        <div class="SecretKey__item">
          <span class="SecretKey__label"> ApiKey </span>
          <span class="SecretKey__val">
            {{ item.ApiKey }}
          </span>
        </div>
        <div class="SecretKey__item">
          <span class="SecretKey__label"> SecretKey </span>
          <span class="SecretKey__val">
            {{ item.SecretKey }}
          </span>
        </div>
        <div class="SecretKey__item">
          <span class="SecretKey__label"> Passphrase </span>
          <span class="SecretKey__val">
            {{ item.Passphrase }}
          </span>
        </div>
        <div class="SecretKey__item">
          <span class="SecretKey__label"> IP </span>
          <span class="SecretKey__val">
            {{ item.IP }}
          </span>
        </div>
        <div class="SecretKey__item">
          <span class="SecretKey__label"> 创建时间 </span>
          <span class="SecretKey__val">
            {{ DateFormat(item.CreateTime) }}
          </span>
        </div>
        <div class="SecretKey__item">
          <span class="SecretKey__label"> 备注 </span>
          <span class="SecretKey__val">
            {{ item.Note }}
          </span>
        </div>
      </n-card>
    </div>

    <n-empty description="你什么也找不到" v-else> </n-empty>

    <div class="AddBnt__wrapper">
      <RouterLink to="/secret_key/add">
        <n-button size="small" type="primary"> 新增一个秘钥 </n-button>
      </RouterLink>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

@media screen and (min-width: 648px) {
  .SecretKey_wrapper {
    .n-card {
      max-width: 290px;
      display: inline-block;
    }
  }
}

.AddBnt__wrapper {
  text-align: center;
  margin-top: 24px;
}
.SecretKey_wrapper {
  .n-card {
    margin: 0 4px;
    margin-bottom: 8px;
  }
  .n-card.Active {
    border-color: @mainColor;
  }
  .n-card__content {
    position: relative;
    padding: 12px;
    padding-top: 12px;
    .SecretKey_del {
      position: absolute;
      right: 10px;
      top: 10px;
    }
  }
}
.SecretKey__item {
  display: flex;
  font-size: 12px;
  padding: 8px 0;
  align-items: center;
  word-break: break-all;
}
.SecretKey__label {
  width: 74px;
  flex-grow: 0;
  flex-shrink: 0;
  color: #333;
  &::after {
    content: ':';
  }
}

.SecretKey__val {
  color: #999;
}
</style>
