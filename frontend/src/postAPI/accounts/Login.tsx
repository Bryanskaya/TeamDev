import axios from "axios";
import { backUrl } from "..";
import { Account } from "types/Account";

export const Login = async function (data: Account, setCookie) {
  const response = await axios
    .post(backUrl + `/accounts/login`, data, { withCredentials: true })
    .then((data) => data)
    .catch((error) => {
      return { status: error.response?.status, data: error.response.data };
    });

  if (response.status === 200) {
    setCookie("login", response.data.login, { path: "/" });
    setCookie("role", response.data.role, { path: "/" });
  }

  return {
    status: response?.status,
  };
};
