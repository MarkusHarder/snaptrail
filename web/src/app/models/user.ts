export interface User {
  username: string;
  password: string;
}

export interface PasswordChange {
  username: string;
  oldPassword: string;
  newPassword: string;
}
