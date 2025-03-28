import * as config from "../config";

async function getAccount() {
  const res = await fetch(`${config.casdoorUrl}/api/get-account`, {
    method: "GET",
    credentials: "include",
  });
  console.log(res);
  return await res.json();
}

export default {
  getAccount,
};
