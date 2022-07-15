<script setup lang="ts">
import { GetServerList, DelServer } from '@/api/AIFundServe';
import { GetOkxKeyList } from '@/api/OkxKey';
import { DateFormat } from '@/utils/filters';
import { CoinFundPing } from '@/api/CoinFundServe';
import { $lcg, mStorage, cloneDeep } from '@/utils/tools';
import { useRouter } from 'vue-router';
import AuthModal from '@/lib/AuthModal';

import { defineAsyncComponent } from 'vue';
import { flatMap } from 'lodash';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));

const $router = useRouter();

let KeyList = $ref([]);

const GetKeyList = async () => {
  const res = await GetOkxKeyList();
  if (res.Code > 0) {
    KeyList = res.Data || [];
  }
};

let ServerList = $ref([]);

const AIFundNetPing = (id: string) => {
  if (id) {
    CoinFundPing({
      CoinServeID: id,
    }).then((res) => {
      const label = $lcg(res, 'Data.AppInfo.name', '');
      if (label == 'coin-fund') {
        for (let i = 0; i < ServerList.length; i++) {
          const item = ServerList[i];
          if (item.CoinServeID == id) {
            ServerList[i].Status = true;
            break;
          }
        }
      }
    });
  }
};

const GetList = async () => {
  const res = await GetServerList();
  if (res.Code > 0) {
    ServerList = res.Data || [];

    for (const item of ServerList) {
      AIFundNetPing(item.CoinServeID);
    }
  }
};

GetKeyList();
GetList();

function GetKey(KeyId: string): any {
  let OkxKey = {};

  for (const item of KeyList) {
    if (item.OkxKeyID == KeyId) {
      OkxKey = item;
      break;
    }
  }

  return OkxKey;
}

const OpenServerDetail = (id: string) => {
  mStorage.set('CoinServeID', id);
  $router.push('/CoinServe/Info');
};

const delServer = (item) => {
  const data = cloneDeep(item);
  AuthModal({
    IsPassword: true,
    async OkBack(param) {
      return DelServer({
        CoinServeID: data.CoinServeID,
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
  <PageTitle>AIFundServer</PageTitle>
  <div class="PageWrapper AIFundServerList">
    <div v-if="ServerList.length > 0 && KeyList.length > 0" class="ListWrapper">
      <n-card
        v-for="item in ServerList"
        :key="item.CoinServeID"
        :title="item.CoinServeID"
        embedded
        hoverable
        size="small"
      >
        <n-button class="Server_del" strong secondary circle size="small" type="error" @click="delServer(item)">
          <template #icon>
            <XIcon name="DeleteOutlined" />
          </template>
        </n-button>
        <div class="Server__item">
          <span class="Server__label"> Host </span>
          <span class="Server__val">
            {{ item.Host }}
          </span>
        </div>

        <div class="Server__item">
          <span class="Server__label"> Port </span>
          <span class="Server__val">
            {{ item.Port }}
          </span>
        </div>

        <div class="Server__item">
          <span class="Server__label"> 密钥 </span>
          <span class="Server__val">
            {{ GetKey(item.OkxKeyID).Name }}
            <RouterLink :to="'/OkxKey?id=' + item.OkxKeyID" class="Server__link"> 【查看】 </RouterLink>
          </span>
        </div>

        <div class="Server__item">
          <span class="Server__label"> 创建时间 </span>
          <span class="Server__val"> {{ DateFormat(item.CreateTime) }} </span>
        </div>

        <div class="Server__item">
          <span class="Server__label"> 备注 </span>
          <span class="Server__val">
            {{ item.Note }}
          </span>
        </div>
        <template #footer>
          <div class="card_footer">
            <n-button size="small" v-if="item.Status" type="success" @click="OpenServerDetail(item.CoinServeID)">
              服务正在运行【查看】
            </n-button>
            <n-button size="small" v-else type="error" @click="OpenServerDetail(item.CoinServeID)">
              服务尚未运行【去部署】
            </n-button>
          </div>
        </template>
      </n-card>
    </div>

    <n-empty description="你什么也找不到" v-else> </n-empty>

    <div class="AddBnt__wrapper">
      <RouterLink to="/CoinServe/Add" class="item_btn" v-if="KeyList.length > 0">
        <n-button size="small" type="primary"> 创建一个新的服务 </n-button>
      </RouterLink>
      <RouterLink to="/OkxKey/Add" class="item_btn" v-else>
        <n-button size="small" type="primary"> 创建一个Key </n-button>
      </RouterLink>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
.AddBnt__wrapper {
  text-align: center;
  margin-top: 24px;
}
.ListWrapper {
  .n-card {
    margin: 0 4px;
    margin-bottom: 8px;
  }
  .n-card__content {
    padding: 12px;
    padding-top: 12px;
    .Server_del {
      position: absolute;
      right: 10px;
      top: 10px;
    }
  }
}

.Server__item {
  word-break: break-all;
  display: flex;
  font-size: 12px;
  padding: 8px 0;
  align-items: center;
}
.Server__label {
  width: 74px;
  flex-grow: 0;
  flex-shrink: 0;
  color: #333;
  &::after {
    content: ':';
  }
}

.Server__val {
  color: #999;
}

.Server__link {
  color: @mainColor;
}

.card_footer {
  text-align: right;
}
</style>
