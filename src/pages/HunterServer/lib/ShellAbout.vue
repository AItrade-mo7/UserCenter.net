<script setup lang="ts">
import { CopyText } from '@/utils/tools';
import { mStorage } from '@/utils/tools';
import { GetHunterConfig } from '@/api/hunter_net';

const hunter_host = mStorage.get('hunter_host');

let Port = '';
if (hunter_host) {
  const host_arr = hunter_host.split(':');
  Port = host_arr[1];
}

const copy = () => {
  CopyText(wgetSh);
};
const props = defineProps({
  Src: String,
});
const wgetSh = `wget -qO- ${props.Src} | sudo bash`;

const getConfig = () => {
  GetHunterConfig({
    ServerInfo: {
      Host: hunter_host,
    },
  })
    .then((res) => {
      if (res.Code > 0) {
        window.$message.success(res.Msg);
        window.location.reload();
      } else {
        window.$message.warning('未获取到部署信息');
      }
    })
    .catch(() => {
      window.$message.warning('未获取到部署信息');
    });
};
</script>

<template>
  <div class="ShellAbout">
    <h3>Hunter.net 部署文档</h3>
    <div className="ShellAbout_hint">系统已为您生成了一键部署指令:</div>
    <div className="ShellAbout__urlBox">
      <n-code :code="wgetSh" word-wrap> </n-code>
      <n-button size="tiny" type="primary" @click="copy"> 复制 </n-button>
    </div>

    <div className="ShellAbout_desc">
      复制该指令，并在 ip 为
      <div className="ShellAbout_desc-ip">
        <code>
          <a :href="`http://${hunter_host}`" target="_blank"> {{ hunter_host }} </a>
        </code>
      </div>
      的主机上执行。
      <br />
      请开放该主机的 <span className="lineHight">{{ Port }}</span> 端口。
      <br />
      <br />
      主机硬件要求：
      <br />
      <span className="lineHight">64位(x86)</span>或<span className="lineHight">64位(ARM)</span>的
      <span className="lineHight">Linux</span>
      系统
      <br />
      <br />
      系统版本：
      <br />
      <span className="lineHight">Ubuntu 20.04 LTS</span> 或以上版本
      <br />
      <br />
      硬件配置：
      <br />
      <span className="lineHight">1GB</span> 以上内存 <span className="lineHight">15GB</span> 以上存储,
      <br />
      <br />
      位置要求：
      <br />
      优先推荐 AWS
      <span className="lineHight">美国西部 (加利福尼亚北部)</span> 的云主机
      <br />
      <span className="lineHight">日本</span>或<span className="lineHight">香港</span>等地的 海外 主机均可
    </div>
    <br />
    <br />
    <div>
      <n-button type="primary" size="small" @click="getConfig"> 我已部署完成 , 点击连接</n-button>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
.ShellAbout__urlBox {
  color: #24292f;
  background-color: #f6f8fa;
  margin: 16px 0;
  padding: 10px;
  position: relative;
  padding-bottom: 40px;
  .n-button {
    position: absolute;
    right: 0;
    bottom: 0;
    margin: 10px;
    float: right;
    font-size: 12px;
  }
  .n-code {
    font-size: 12px;
    color: #999;
    background-color: @mainColor;
  }
}
</style>
