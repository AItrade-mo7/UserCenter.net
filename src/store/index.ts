import { reactive } from 'vue';
import { Logo } from '@/config/constant';
import { cloneDeep } from '@/utils/tools';

// ===================== TopBar ============================
export const TopBarStore = reactive({
  status: false,
  close() {
    this.status = false;
  },
  open() {
    this.status = true;
  },
});

// ===================== UserInfo ============================
export interface UserInfoType {
  Avatar: string;
  CreateTime: number;
  Email: string;
  NickName: string;
  UserID: string;
  UpdateTime: number;
  SecurityCode: string;
}

const UserInfo: UserInfoType = {
  Avatar: Logo,
  CreateTime: 0,
  Email: '',
  NickName: '',
  UserID: '',
  UpdateTime: 0,
  SecurityCode: '',
};
export const UserInfoStore = reactive({
  value: UserInfo,
  update(data: UserInfoType) {
    this.value = data;
  },
  get() {
    return cloneDeep(this.value);
  },
});

// ===================== PingData ============================

export interface AppInfo {
  Port: number;
  name: string;
  version: string;
}

export interface PingDataType {
  AppInfo: AppInfo;
  ContentType: string;
  FullPath: string;
  Method: 'POST' | 'GET' | '';
  ResParam: any;
  Token: string;
  UserAgent: string;
}

const PingData: PingDataType = {
  AppInfo: {
    Port: 0,
    name: '',
    version: '',
  },
  ContentType: '',
  FullPath: '',
  Method: '',
  ResParam: {},
  Token: '',
  UserAgent: '',
};
export const PingDataStore = reactive({
  value: PingData,
  update(data: PingDataType) {
    this.value = data;
  },
});

// ===================== Loading ============================
export const LoadingStore = reactive({
  status: false,
  info: '',
  open(text?: string) {
    this.status = true;
    if (text) {
      this.info = text;
    }
  },
  close() {
    if (this.status) {
      this.status = false;
      this.info = '';
    }
  },
});
