/// <reference types="vite/client" />
/// <reference types="vite-plugin-pwa/client" />

import type { MessageApiInjection, NotificationApiInjection, DialogApiInjection } from 'naive-ui';

declare module '*.vue' {
  import { ComponentOptions } from 'vue';
  const componentOptions: ComponentOptions;
  export default componentOptions;
}

declare global {
  declare const $ref: any;
  declare const ViteConst: any;

  interface Window {
    mo7Encrypt: (msg: string) => string;
    mo7Md5: (msg: string) => string;
    mo7SecretKey: string;
    deferredPrompt: any;
    $message: MessageApiInjection;
    $dialog: DialogApiInjection;
    $notification: NotificationApiInjection;
    ViteConst: {
      AppVersion: string;
      AppName: string;
      ProxyUrl: string;
      rmAgin: string;
    };
  }
}

export {};
