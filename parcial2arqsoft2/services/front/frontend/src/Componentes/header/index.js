import React from "react";
import swal from "sweetalert2";
import './header.css';
import {Link} from "react-router-dom";
import Cookies from "universal-cookie";

function LogOut(){
    //var resultado = window.confirm('Estas seguro?');
    swal.fire({
        text: "Estas seguro?",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Yes'
      }).then((result) => {
        if (result.isConfirmed) {
            Cookie.set("user", "undefined,undefined", { path: "/" });
            window.location.replace("/");
        }})
  }
    
  const Cookie = new Cookies();
  let cookie = Cookie.get("user")
  let id_user;
  const propertyId = Cookie.get("propertyId");
  if(cookie!=undefined){
  let array = cookie.split(",")
   id_user = array[0]
  }
  else{
       id_user = "undefined"
  }
export const Header = () => {
return(
<header className="global" id="header-global">
{id_user!="undefined"?
<ul className="ul_header">
    <li>
       <Link to="/home" className="botones_menu">
       Inicio
       </Link> 
    </li>
    <li>
       <Link to="/properties" className="botones_menu">
       Publica tu propiedad
       </Link> 
    </li>
    <button className="btn" onClick={()=>LogOut()}>
        Salir
    </button>
</ul>
    :
       <Link to="/" className="botones_menu">
        </Link> 
     }
</header>
)}