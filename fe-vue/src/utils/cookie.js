import jsCookie from "js-cookie";

export const LoginTokenName = "ws-token";
export const roleCookieName = "ws-role";
export const LoginTokenExpires = 365;

export function setToken(
  value,
  tokenName = LoginTokenName,
  expires = LoginTokenExpires
) {
  jsCookie.set(tokenName, value, {
    path: "/",
    expires
  });
}

export function getToken(tokenName = LoginTokenName) {
  return jsCookie.get(tokenName);
}

export function removeToken(tokenName = LoginTokenName) {
  jsCookie.remove(tokenName, {
    expires: 0,
    path: "/"
  });
}
