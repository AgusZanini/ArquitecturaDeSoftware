import React, { useState, List, Checkbox} from "react";
import "./css/Cart.css";
import "./css/materialize.css"
import logo from "./images/logo.png"
import Cookies from "universal-cookie";
import RemoveShoppingCartOutlinedIcon from '@mui/icons-material/RemoveShoppingCartOutlined';

const Cookie = new Cookies();

async function getUserById(id){
  return fetch("http://127.0.0.1:8080/user/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getProductById(id){
  return fetch("http://127.0.0.1:8080/product/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function postOrder(products) {
  let details = []
  products.forEach((item) => {
    let detail = {
      product_id: item.product_id,
      quantity: Number(item.quantity),
      price: Number(item.base_price),
      name: item.name
    }
    details.push(detail)
  });

  return fetch("http://127.0.0.1:8080/order", {
    method:"POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      user_id: Number(Cookie.get("user_id")),
      details: details
    })
  })
}

function goto(path){
  window.location = window.location.origin + path
}

function logout(){
  Cookie.set("user_id", -1, {path:"/"})
  window.location.reload();
}

async function getCartProducts(){
  let items = []
  let a = Cookie.get("cart").split(";")

  for (let i = 0; i < a.length; i++){
    let item = a[i];
    if(item != ""){
      let array = item.split(",")
      let id = array[0]
      let quantity = array[1]
      let product = await getProductById(id)
      product.quantity = quantity;
      items.push(product)
    }
  }
  return items
}

function getOptions(n){
  let options = []
  for(let i=1; i <= n; i++){
    options.push(i)
  }
  return options.map((option) =>
    <option value={option}> {option} </option>
)
}

function remove(n, p_id){
  let cookie = Cookie.get("cart");
  let newCookie = ""
  let toCompare = cookie.split(";")
  let isEmpty = false
  toCompare.forEach((item) => {
    if(item != ""){
      let array = item.split(",")
      let item_id = array[0]
      let item_quantity = array[1]
      if(p_id == item_id){
        item_quantity = Number(item_quantity) - n
        if(item_quantity == 0){
          isEmpty = true
        }
      }
      if(p_id == item_id && !isEmpty){
        newCookie += item_id + "," + item_quantity + ";"
      }
      else if (p_id != item_id){
        newCookie += item_id + "," + item_quantity + ";"
      }
    }
  });
  cookie = newCookie
  Cookie.set("cart", cookie, {path: "/"})
  goto("/cart")
  return
}

function showProducts(products){
  return products.map((product) =>
  
  <div className="productt">
    <hr></hr>
    <div obj={product} key={product.product_id}>
      <div >
        <img className="picture" src={product.picture_url}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
      </div>
      <div >
      <h4 className="name">{product.name}</h4>
      <h5 className="price">Precio por unidad: {"$" + product.base_price}</h5>
      <h6 className="detalles"> Cantidad: {product.quantity}</h6>
      <h6 className="detalles"> Subtotal: ${product.quantity * product.base_price} </h6>
      </div>
    </div>
    <div className="detalles-carrito">
      <select id={"removeSelect" + product.product_id}>
        {getOptions(product.quantity)}
      </select>
      
      <div className="subtotal">
      <button  className="remove" onClick={() => remove(document.getElementById("removeSelect" + product.product_id).value, product.product_id)}>Eliminar Uno  <i className="material-icons">clear</i></button>
     
      </div>
    </div>
</div>
 )

}

async function setCart(setter, setterTotal){
  let total = 0;
  await getCartProducts().then(response => {
    setter(response)
    response.forEach((item) => {
      total += item.base_price * item.quantity;
    });
    setterTotal(total)
  })
}

async function buy(products){
  await postOrder(products).then(response => {
    if(response.status == 409){
      response.json().then(response => Cookie.set("orderError", response, {path: "/order"}))
      goto("/order/error")
    }
    if(response.status == 201){
      Cookie.set("order", Cookie.get("cart"), {path: "/order"})
      Cookie.set("cart", "", {path: "/"})
      Cookie.set("cartItems", "", {path: "/"})
      goto("/order/complete")
    }
  })
}

function Cart(){
  const [user, setUser] = useState({});
  const [isLogged, setIsLogged] = useState(false);
  const [cartProducts, setCartProducts] = useState([]);
  const [total, setTotal] = useState(0);


  if (cartProducts.length <= 0 && Cookie.get("user_id") > -1){
    setCart(setCartProducts, setTotal)
  }

  const login = (

    <span>
    <a id="logout" onClick={logout}> <span>{user.first_name} </span> </a>
    </span>
  )

  if (Cookie.get("user_id") > -1 && !isLogged) {
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
  }

  const renderEmptyCart = (
    <a className="boton-vacio" href="/login">NECESITA INICIAR SESION PARA PODER VER SUS PRODUCTOS DEL CARRITO</a>
  )

  const renderOrderButton = (
    <footer >
      <div >
        <div >
          <ul>
            <li><button  onClick={() => buy(cartProducts)} className="Finalizar">FINALIZAR COMPRA</button></li>
          </ul>
        </div>
        <div >
          <ul>
            <li><h3 className="Total"> Total: ${total} </h3></li>
          </ul>
        </div>
      </div>
  </footer>
  )

  return (
    <div className="cart">
      <nav class=" purple ">
        <div class="nav-wrapper">
          <a href="/" class="brand-logo center blue-text text-darken-2" ><img  src={logo} width="60px" height="60px"/> </a>
          <ul id="nav-mobile" class="right hide-on-med-and-down">
          <li><a className="Sesion" onClick={logout}>Cerrar Sesion</a></li>
          </ul>
        </div>
      </nav>



      <div id="main">
        {Cookie.get("cart") ? (Cookie.get("user_id") > -1 ? showProducts(cartProducts) : renderEmptyCart) : <a className="boton-vacio" href="/"> CARRITO VACIO, PRESIONE PARA AGREGAR UN PRODUCTO</a>}
        {Cookie.get("cart") && Cookie.get("user_id") > -1 ? renderOrderButton : <span/>}


      </div>
    </div>
  );
}

export default Cart;
