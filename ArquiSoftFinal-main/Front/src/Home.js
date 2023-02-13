import React, { useState } from "react";
import "./css/Home.css";
import logo from "./images/logo.png"
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function getUserById(id){
    return await fetch('http://127.0.0.1:8080/user/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())

}

async function getCategories(){
  return await fetch('http://127.0.0.1:8080/categories', {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getProducts(){
  return await fetch('http://127.0.0.1:8080/products', {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}
async function getProductsByCategoryId(id){
  return await fetch('http://127.0.0.1:8080/products/' + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

function goto(path){
  window.location = window.location.origin + path
}


function gotologin(){
  goto("/login")
}


function retry() {
  goto("/")
}

function addToCart(id, setCartItems){
  let cookie = Cookie.get("cart");

  if(cookie == undefined){
    Cookie.set("cart", id + ",1;", {path: "/"});
    setCartItems(1)
    return
  }
  let newCookie = ""
  let isNewItem = true
  let toCompare = cookie.split(";")
  let total = 0;
  toCompare.forEach((item) => {
    if(item != ""){
      let array = item.split(",")
      let item_id = array[0]
      let item_quantity = array[1]
      if(id == item_id){
        item_quantity = Number(item_quantity) + 1
        isNewItem = false
      }
      newCookie += item_id + "," + item_quantity + ";"
      total += Number(item_quantity);
    }
  });
  if(isNewItem){
    newCookie += id + ",1;"
    total += 1;
  }
  cookie = newCookie
  Cookie.set("cart", cookie, {path: "/"})
  Cookie.set("cartItems", total, {path: "/"})
  setCartItems(total)
  return
}

function showProducts(products, setCartItems){
  return products.map((product) =>
    <div className="div-products">
      <div className="caja" key={product.product_id} >
        <div >
        <img className="pictur" src={product.picture_url}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
        </div>
        <div>
          <span><a className="name">{product.name}</a></span>
          <p>Cantidad disponible: {product.stock}</p>
          <p>Precio: ${product.base_price}</p>
        </div>
        <div>
          <button className="button1" onClick={() => addToCart(product.product_id, setCartItems)}>Agregar al carrito 🛒</button>
        </div>
      </div>
    </div>
 )
}

function logout(){
  Cookie.set("user_id", -1, {path: "/"})
  document.location.reload()
}

function search(){
  let input, filter, a, i;
  input = document.getElementById("search");
  filter = input.value.toUpperCase();
  a = document.getElementsByClassName("product");
  for (i = 0; i < a.length; i++) {
    let txtValue = a[i].children[1].textContent || a[i].children[1].innerText;
    if (txtValue.toUpperCase().indexOf(filter) > -1) {
      a[i].style.display = "inherit";
    } else {
      a[i].style.display = "none";
    }
  }
  if(input.value.toUpperCase().length <= 0){
    for(i = 0; i < a.length; i++){
      a[i].style.display = "inherit";
    }
  }

}

function deleteCategory(){
  Cookie.set("category", 0, {path: "/"})
  goto("/")
}

function gotocart(){
  goto("/cart")
}

function gotocompras(){
  goto("/compras")
}

async function getProductBySearch(query){
  return fetch("http://localhost:8080/products/search="+query, {
    method: "GET",
    header: "Content-Type: application/json"
  }).then(response=>response.json())
}


function Home() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})
  const [products, setProducts] = useState([])
  const [cartItems, setCartItems] = useState("")
  const [failedSearch, setFailedSearch] = useState(false)


  if (Cookie.get("user_id") > -1 && !isLogged){
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
  }

  if (!(Cookie.get("user_id") > -1) && isLogged){
    setIsLogged(false)
  }

  if (products.length <= 0){
    getProducts().then(response => {setProducts(response)})
  }

  if (!cartItems && Cookie.get("cartItems")){
    setCartItems(Cookie.get("cartItems"))
  }

  async function searchQuery(query){

    await getProductBySearch(query).then(response=>{
      console.log(query)
      if(response != null){
        if(response.length > 0){
          setProducts(response)
          setFailedSearch(false)
        }else{
          setProducts([])
          setFailedSearch(true)
        }
      }
      else{
        setFailedSearch(false)
        getProducts().then(response=>setProducts(response))
      }
    })


  }

  async function categories(id){

    await getProductsByCategoryId(id).then(response=>{
      console.log(id)
      if(response != null){
        if(response.length > 0){
          setProducts(response)
          setFailedSearch(false)
        }else{
          setProducts([])
          setFailedSearch(true)
        }
      }
      else{
        setFailedSearch(false)
        getProducts().then(response=>setProducts(response))
      }
    })


  }

  const login = (
    <ul >
      <li><a onClick={gotocart} className="carro" ><i class="material-icons black-text">shopping_cart</i></a></li>
      <li><p className="numero">{cartItems > 0 ? cartItems : 0}</p></li>
      <li><a onClick={gotocompras} className="Sesion">Mis Compras</a></li>
      <li><a onClick={logout} className="Sesion">Cerrar Sesion</a></li>
     
    </ul>
  )

  const renderFailedSearch = (
    <a>NO SE ENCUENTRA PRODUCTO</a>
  )
  
  return (
    <div className="home">
      <nav class=" purple">
        <div class="nav-wrapper">
            <form className="col ">
              <div class="input-field">
              <input className="barra" placeholder="Buscar productos" onKeyDown={(e) => e.keyCode === 13 ? searchQuery(e.target.value) : void(0)} />
                <input  className="hidden" type="search" required onChange={search}/>
                <label class="label-icon" for="search" className="Lupa"><i class="material-icons black-text">search</i></label>
                <i class="material-icons black-text" onClick={() => categories(0)}>close</i>
              </div>
            </form>
          <a href="/" class="brand-logo center blue-text text-darken-2"><img src={logo} width="60px" height="60px" /></a> 
          <ul class="right hide-on-med-and-down">
            <li>{isLogged ? login : <a onClick={gotologin} className="Sesion">Iniciar Sesion</a>}</li>
          </ul>
        
        </div>
      </nav>
      <div className="categories content">
      <button className="botoncat" onClick={() => categories(1)}>
            AUTOS
          </button>
      <button className="botoncat" onClick={() => categories(2)}>
            CELULARES
          </button>
          <button className="botoncat" onClick={() => categories(3)}>
            FERRETERIA
          </button>
          <button className="botoncat" onClick={() => categories(4)}>
            FITNESS
          </button>
          <button className="botoncat" onClick={() => categories(5)}>
            ROPA
          </button>
          <button className="botoncat" onClick={() => categories(6)}>
            JUGUETES
          </button>
          <button className="botoncat" onClick={() => categories(7)}>
            COMIDA
          </button>
          <button className="botoncat" onClick={() => categories(8)}>
            COMPUTACION
          </button>
          <div></div>
       
      </div>
     
         
      <div class="row" id="main">
        {products.length > 0 || failedSearch ? showProducts(products, setCartItems) : <a>CARGANDO...</a>}
      </div>

      
    </div>
  );
}

export default Home;
