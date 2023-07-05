import Cookies from "universal-cookie";

const Cookie= new Cookies()

export function CookieUser(id){
  
    let cookie = Cookie.get("property");
   
    if(cookie == undefined){
      Cookie.set("property", id + ";", {path: "/"});
      return
    }
    return
  }