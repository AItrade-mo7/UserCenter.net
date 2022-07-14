import type { resType, resDataType } from './utils.d';
import { logout } from '@/api/Account';

export const res_dispose = (response: resType): resDataType => {
  const data = response.data;

  // 账号尚未注册
  if (data.Code == -7) {
    window.$message.warning(data.Msg, {
      onLeave() {
        window.location.replace('/register');
      },
    });
    return data;
  }

  // Token 错误
  if (data.Code == -3) {
    window.$message.warning(data.Msg, {
      onLeave() {
        logout();
      },
    });
    return data;
  }

  if (data.Code < 0) {
    window.$message.error(data.Msg);
    return data;
  }

  return data;
};
