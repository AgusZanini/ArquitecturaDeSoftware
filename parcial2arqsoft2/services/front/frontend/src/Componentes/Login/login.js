
import React,{ useState} from "react";
import "./login.css";
import Cookies from "universal-cookie";
import Libertador from "../../images/logo.jpeg";
import swal from "sweetalert2";
const Cookie = new Cookies();


export default function Login(){

  window.onload = function(){
    document.getElementById("header-global").style.display = "none"
  }
  
  const[user,setUser]= useState("");
  const[password,setPassword] = useState("");

  const onChangeUser =  (user)=>{
      setUser(user.target.value);
      
  }
  
  const onChangePas = (password)=>{
  setPassword(password.target.value)};

  
  const requestOptions={
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      
      body: JSON.stringify({username : user, password : password })
  };

  const login = async()=>{
      fetch('http://localhost:9000/login',requestOptions)
      .then(response => {if (response.status == 400) {
         swal.fire({
          text: "Usuario incorrectos",
          icon: 'error',
         }).then((result) => {
          if (result.isConfirmed) {
              window.location.reload();
              return response.json()
          }})
      }
      if(response.status==200){
        swal.fire({icon: 'success'}
        ).then((result) => {
          if (result.isConfirmed) {
            window.location.replace("/home")
            return response.json()
          }})
      }
      return response.json()})
      .then(response => {
        //alert(response.userid)
          Cookie.set("user", response.userid + "," + response.token, {path: "/"})
  })
 
  };
 
  const handleSubmit= (event)=>{
      event.preventDefault();
      login();

  };

  return(
    <div className="app">
      <div className="login-form">
      <form onSubmit={handleSubmit} >
      <div className="ul">
      <div class="logo-div-2">
        <img class="logo-2" src={Libertador} />
        </div>
      
      <h1 className="login">Ingreso</h1>
      
        <div>
      <input id="user" className="user" type={"text"} placeholder="Ingrese su usuario" onChange={onChangeUser} value ={user} required></input>
        </div>
      <div>
      <input id="password" type={"password"} placeholder="Ingrese su contraseña" onChange={onChangePas} value={password} required></input>
      </div>
      <input type="submit" value="Ingresar"></input>
      </div>
      
      </form>
      </div>
      </div>

  );
}
