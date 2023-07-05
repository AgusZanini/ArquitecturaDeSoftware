import Cookies from "universal-cookie";

const Cookie= new Cookies()

export function CookieUser(userid,token){
  
    let cookie = Cookie.get("username");
   
    if(cookie == undefined){
      Cookie.set("username", userid + "," + token + ";", {path: "/"});
      return
    }
    return
  }