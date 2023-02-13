import React, { useState, List, Checkbox} from "react";
import "./css/Compras.css";
import logo from "./images/logo.png"
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function getUserById(id){
    return await fetch('http://localhost:8080/user/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())

}

async function getOrdersByUserId(id){
    return await fetch('http://localhost:8080/orders/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())

}




function logout(){
  Cookie.set("user_id", -1, {path: "/"})
  document.location.reload()
}

function showDetails(id){
  let div = document.getElementsByClassName("o" + id)
  for(let detail of div){
    let show = true;
    detail.classList.value.split(" ").forEach(c=>{
      if (c == "hidden"){
        show = true;
      }
    })
    if(show){
      detail.classList.remove("hidden")
    }else{
      detail.classList.add("hidden")
    }
  }


}

function getDetails(details){
  return details.map((detail) =>
      <tbody>
        <tr className="detail">
          <td className="id">{detail.product_id}</td>
          <td className="name">{detail.name}</td>
          <td className="quantity">{detail.quantity}</td>
          <td className="price">{"$" + detail.price}</td>
        </tr>
      </tbody>
  )
}

function showUserOrders(orders){
  return orders.map((order) =>

  <div obj={order} key={order.order_id} className="order" onClick={()=>showDetails(order.order_id)}>
    <a className="ordern">Orden: {order.order_id}       </a>
    <a className="date">Fecha:     {order.date.split("T")[0]}       </a>
    <a className="total">Total: <span>{"$" + order.total}  </span></a>

    <div className={"details" + " o" + order.order_id + " hidden"}>
      <table>
        <thead>
          <tr>
            <td className="name">ID</td>
            <td className="name">Producto</td>
            <td className="quantity">Cantidad</td>
            <td className="price">Precio</td>
          </tr>
        </thead>

      {getDetails(order.details)}
      </table>
    </div>

  </div>
 )
}



function User() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})
  const [cartItems, setCartItems] = useState("")
  const [orders, setOrders] = useState([])


  if (Cookie.get("user_id") > -1 && !isLogged){
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
    if (!cartItems && Cookie.get("cartItems")){
      setCartItems(Cookie.get("cartItems"))
    }

    if(Cookie.get("orders") == undefined){
      getOrdersByUserId(Cookie.get("user_id")).then(response => {setOrders(response)})
    }


  }

  if (!(Cookie.get("user_id") > -1) && isLogged){
    setIsLogged(false)
  }

  const login = (

    <span>
    <a id="logout" onClick={logout}> <span>{user.first_name} </span> </a>
    </span>
  )

  const showUserInfo = (
    <div>
      <div className="Nombre">Compras de {user.first_name} {user.last_name} </div>
      <div className="Username"> Username: {user.username} </div>
      <div >{orders.length > 0 ? showUserOrders(orders) : <a className="norders"> NO HAY COMPRAS</a>}</div>
    </div>
  )

  const pleaseLogin = (
    <a href="/login"><button className="inciesesion">INICIE SESION PARA VER SUS COMPRAS </button> </a>
  )

  return (
    <div className="home">
<nav class=" purple ">
        <div class="nav-wrapper">
          <a href="/" class="brand-logo center blue-text text-darken-2"><img src={logo} width="60px" height="60px"/> </a>
          <ul id="nav-mobile" class="right hide-on-med-and-down">
          <li><a className="Sesion" onClick={logout}>Cerrar Sesion</a></li>
          </ul>
        </div>
      </nav>
      <div id="main">
        {isLogged ? showUserInfo : pleaseLogin}
      </div>
    </div>
  );
}

export default User;