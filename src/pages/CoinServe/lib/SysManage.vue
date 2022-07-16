<script setup lang="ts">
import { cloneDeep } from '@/utils/tools';
import { ReStart, Remove } from '@/api/CoinAI/sys';
import AuthModal from '@/lib/AuthModal';
import { mStorage } from '@/utils/tools';

const props = defineProps({
  config: Object,
});

const nowVersion = cloneDeep(props.config.AppInfo.version);
const newVersion = cloneDeep(props.config.GithubInfo.version);
let isUpdate = false;

if (newVersion && nowVersion) {
  isUpdate = newVersion !== nowVersion;
}

let BtnStatus = $ref(0);
const stopServer = () => {
  BtnStatus = -1;
};

const restartServer = () => {
  BtnStatus = 1;
};

const SendFetch = (type: number) => {
  AuthModal({
    EmailAction: '修改资料',
    IsPassword: true,
    async OkBack(param) {
      const Info = {
        Password: param.Password,
        Code: param.Code,
      };
      if (type === 1) {
        return SendReStart(Info);
      }

      if (type === -1) {
        return SendStop(Info);
      }
    },
  });
};

const SendStop = async (Info) => {
  await Remove({
    ...Info,
    ServerInfo: {
      Host: props.config.AppEnv.CoinServeID,
    },
  });
  window.$message.success('删除指令已发送!', {
    onAfterLeave() {
      mStorage.remove('CoinServeID');
      window.location.replace('/CoinServe');
    },
  });
};

const SendReStart = async (Info) => {
  await ReStart({
    ...Info,
    ServerInfo: {
      Host: props.config.AppEnv.CoinServeID,
    },
  });

  window.$message.success('重启指令已发送!', {
    onAfterLeave() {
      mStorage.remove('CoinServeID');
      window.location.replace('/CoinServe');
    },
  });
};
</script>

<template>
  <div class="SysManage">
    <div className="SysManage__hint">当前版本：{{ nowVersion }}</div>
    <div className="SysManage__hint new">
      最新版本：{{ newVersion }}
      <a href="https://github.com/mo7static/AITrade.net" target="_blank" class="lineHeight">(GitHub)</a>
    </div>

    <div class="btn-wrapper">
      <n-button ghost type="error" @click="stopServer"> 删除服务 </n-button>
      <n-button ghost type="success" @click="restartServer"> 更新并重启 </n-button>
    </div>
    <div class="warn">
      <div v-if="BtnStatus == 0">
        <div v-if="isUpdate" class="SysManage__hint">
          建议执行
          <span className="lineHeight-success">更新并重启</span>
          操作或查看
          <RouterLink to="/About/ReleaseNotes" className="lineHeight">更新日志</RouterLink>
        </div>
        <div v-else>程序将 7*24 小时监听您的账户,并根据您的设置自动执行交易。</div>
      </div>
      <div v-if="BtnStatus == 1">
        将会从
        <a href="https://github.com/mo7static/AITrade.net" target="_blank" class="lineHeight">GitHub</a>
        下载最新代码,并完成
        <span class="lineHeight-warn">冷重启</span>
        ,建议检查并清理您当前的仓位
        <n-button size="tiny" type="success" @click="SendFetch(1)">执行升级</n-button>
      </div>
      <div v-if="BtnStatus == -1">
        这将彻底
        <span class="lineHeight-err">停止</span>
        并
        <span class="lineHeight-err">删除</span>
        您的程序和运行日志 , 但数据不会丢失,稍后可重新部署
        <n-button size="tiny" type="error" @click="SendFetch(-1)">我就是要删掉它</n-button>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.SysManage {
  font-size: 14px;
}

.SysManage__hint {
  color: #999;
  &.new {
    color: #0b8235;
  }
}
.lineHeight {
  color: #f0b90b;
}

.lineHeight-warn {
  color: #fcb040;
}
.lineHeight-err {
  color: #d03050;
}

.lineHeight-success {
  color: #18a058;
}

.btn-wrapper {
  display: flex;
  align-items: center;
  justify-content: space-around;
  margin-top: 20px;
}

.warn {
  font-size: 16px;
  margin-top: 20px;
}
</style>
