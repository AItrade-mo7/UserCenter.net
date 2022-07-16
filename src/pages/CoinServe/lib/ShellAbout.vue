<script setup lang="ts">
import { CopyText } from '@/utils/tools';
import { GetCoinFundConfig } from '@/api/CoinFundServe';

const props = defineProps({
  Src: String,
  CoinServeID: String,
});

let Port = '';
let Host = '';
if (props.CoinServeID) {
  const host_arr = props.CoinServeID.split(':');
  Port = host_arr[1];
  Host = host_arr[0];
}

const copy = () => {
  CopyText(wgetSh);
};

const wgetSh = `wget -qO- ${window.location.protocol}${props.Src} | sudo bash`;

const getConfig = () => {
  GetCoinFundConfig({
    CoinServeID: props.CoinServeID,
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
    <h3>CoinFundServe.net 部署文档</h3>
    <div class="ShellAbout_hint">系统已为您生成了专属的部署指令:</div>
    <div class="ShellAbout__urlBox">
      <n-code :code="wgetSh" word-wrap> </n-code>
      <n-button size="tiny" type="primary" @click="copy"> 复制 </n-button>
    </div>

    <div class="ShellAbout_desc">
      复制该指令，并在 ip 为
      <div class="ShellAbout_desc-ip">
        <a :href="`http://${props.CoinServeID}`" target="_blank">
          <span class="lineHight"> {{ Host }} </span>:{{ Port }}
        </a>
      </div>
      的主机上执行。
      <br />
      请开放该主机的 <span class="lineHight">{{ Port }}</span> 端口。
      <br />
      <br />
      系统要求：
      <span class="lineHight">aarch64</span>或<span class="lineHight">x86_64</span>的
      <span class="lineHight">Linux</span>
      系统
      <br />
      <br />
      系统版本：
      <span class="lineHight">Ubuntu 20.04</span> 或以上版本
      <br />
      <br />
      硬件配置：
      <span class="lineHight">1GB</span> 及以上内存 <span class="lineHight">20GB</span> 及以上存储。
      <span class="desc"> 该程序采用 Go 语言编写, 基于内存进行超高频计算,运行峰值内存消耗为50M~60M </span>
      <br />
      <br />
      位置要求：
      <br />
      优先推荐 AWS
      <span class="lineHight">美国西部 (加利福尼亚北部)</span> 的云主机
      <br />
      <span class="lineHight">日本</span>或<span class="lineHight">香港</span>等地的 海外 主机均可
      <br />
      <br />
      <span class="desc">
        同一台主机可以部署多个程序,为防止IP限制同一主机下部署不应超过<span class="lineHight">8</span>个
      </span>
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

.lineHight {
  color: @mainColor;
}
.desc {
  font-size: 12px;
  color: #666;
}
</style>
