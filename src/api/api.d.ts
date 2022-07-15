export interface LoginParam {
  Email: string;
  Password: string;
}

export interface AddCoinServerParam {
  OkxKeyID: string;
  Host: string;
  Port: string;
  Note: string;
  Password: string;
}

export interface DelCoinServerParam {
  CoinServeID: string;
  Password: string;
}

export interface AddOkxKeyParam {
  ApiKey: string;
  SecretKey: string;
  Passphrase: string;
  IP: string;
  Name: string;
  Note: string;
  Password: string;
}

export interface DelOkxKeyParam {
  OkxKeyID: string;
  Password: string;
}

export interface RegisterParam {
  Email: string;
  Code: string;
}

export interface SendCodeParam {
  Email: string;
  Action: string;
}

export interface ChangePasswordParam {
  Email: string;
  Code: string;
  Password: string;
  AgainPassword: string;
}

export interface EditProfileParam {
  OldEmailCode: string;
  NewEmail: string;
  NewEmailCode: string;
  Avatar: string;
  NickName: string;
  SecurityCode: string;
  Password: string;
}
