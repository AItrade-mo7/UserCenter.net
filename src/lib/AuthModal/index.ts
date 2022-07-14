import CodeView from './CodeView.vue';
import PasswordView from './PasswordView.vue';
import AllView from './AllView.vue';
import { fetchSendCode } from '@/api/Account';
import { UserInfoStore } from '@/store';

export interface BackData {
  Email: string;
  Password: string;
  Code: string;
}

export interface Config {
  Title?: string;
  IsPassword?: boolean;
  EmailAction?: string;
  OkBack: (res: BackData) => void;
}
export default (opt: Config) => {
  const { Title = '需要验证' } = opt;

  let Modal = PasswordView.render;

  if (opt.EmailAction) {
    Modal = CodeView.render;
    fetchSendCode({
      Email: UserInfoStore.value.Email,
      Action: opt.EmailAction,
    }).then((res) => {
      if (res.Code > 0) {
        window.$message.success(res.Msg);
      }
    });
  }

  if (opt.EmailAction && opt.IsPassword) {
    Modal = AllView.render;
  }

  window.$dialog.info({
    title: Title,
    content: Modal,
    bordered: true,
    maskClosable: false,
    closable: false,
    positiveText: '确认',
    negativeText: '取消',
    onPositiveClick: (event) => {
      event.path[0].disabled = true;
      event.path[1].disabled = true;

      const CodeElm: any = window.document.getElementById('AuthModal_Code');
      const PasswordElm: any = window.document.getElementById('AuthModal_Password');

      return opt.OkBack({
        Email: UserInfoStore.value.Email,
        Password: PasswordElm?.value,
        Code: CodeElm?.value,
      });
    },
  });
};
