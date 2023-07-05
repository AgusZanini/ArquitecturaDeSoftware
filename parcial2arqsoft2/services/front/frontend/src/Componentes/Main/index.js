import { PropertyItems } from "../Main/PropertyItems";
import React, { useState } from "react";
import Libertador from "../../images/logo.jpeg";
import swal from "sweetalert2";
import '../Main/main.css';
import Cookies from "universal-cookie";
import { MDBFooter, MDBContainer, MDBRow, MDBCol, MDBIcon } from 'mdb-react-ui-kit';
import { Link } from "react-router-dom";



export const Main = () => {

const Cookie = new Cookies();
    let cookie = Cookie.get("user")
    let id_user;
    if(cookie!=undefined){
    let array = cookie.split(",")
     id_user = array[0]
    }
    else{
         id_user = "undefined"
    }
  const [properties, setProperties] = useState([]);
  
    const [busqueda, setBusqueda]= useState("");
    const fetchBuscador = async()=>{
      
        const response = await fetch('http://localhost:8081/search='+ busqueda)
       .then((response) => response.json())
       if(response===null){
         swal.fire({
           icon: 'error',
           text: "No se encontraron propiedades",
         }) 
         .then((result) => {
           if (result.isConfirmed) {
               window.location.reload();
           }})
     }else{
     setProperties(response)
     console.log(response); 
       }
        };

    const handleChange=e=>{
     setBusqueda(e.target.value);
      };

      const handleSubmit= (event)=>{
        event.preventDefault();
        fetchBuscador();
    };

  return (
    <header class="header">
      <div>
        <div class="logo-div">
        <img class="logo" src={Libertador} />
        </div>
        <div>
          
        </div>
      </div>
      <div class="buscar">
        Tu Inmobiliaria de Confianza
      </div>

      <div class="search-padre">
        <div>
          <div>
          <input
            class="form-control inputBuscar"
            value={busqueda}
            placeholder="Encuentra tu nuevo hogar"
            onChange={handleChange}
          
          />
          </div>
          <div>
            <button class="btn-search" type="button" onClick={handleSubmit}>
                Buscar
              </button>
          </div>
        </div>
      </div>

    <section className='footer-container'>
      <div className='footer-section'>
        <h6 className='text-uppercase fw-bold mb-4'>
          <MDBIcon icon="gem" className="me-3" />
          Leandro Garza
        <p>Corredor inmobiliario</p>
        </h6>

        <h6 className='text-uppercase fw-bold mb-4'>
         <MDBIcon icon="gem" className="me-3" />
         Agustin Zanini
        <p>Ingeniero de Sistemas</p>
       </h6>

       <h6 className='text-uppercase fw-bold mb-4'>
          <MDBIcon icon="gem" className="me-3" />
         Nicolas Bergami
       <p>Desarrollador Full Stack</p>
       </h6>
      </div>

      <div className='footer-section'>
        <h6 className='text-uppercase fw-bold mb-4'>Links Rápidos</h6>
        <p>
         <Link to='/'>Home</Link>
        </p>
        <p>
          <Link to='/'>Sobre Nosotros</Link>
        </p>
      </div>

    <div className='footer-section'>
      <h6 className='text-uppercase fw-bold mb-4'>Contacto</h6>
      <p>
        <MDBIcon icon="envelope" className="me-3" />
        nicolasbergami2013@gmail.com
      </p>
      <p>
        <MDBIcon icon="envelope" className="me-3" />
        leandrogarzap@gmail.com
      </p>
      <p>
        <MDBIcon icon="envelope" className="me-3" />
        aguszanini@gmail.com
      </p>
    </div>
  </section>
   
    
    

        
      <div class="Property-padre">
            {properties
        .map((property) => (
                  <PropertyItems key={property.id}
                    id={property.id}
                    tittle={property.tittle}
                    size={property.size}
                    bathrooms={property.bathrooms}
                    service={property.service}
                    city={property.city}
                    state={property.state}
                    country={property.country}
                    street={property.street}
                    price={property.price}
                    rooms={property.rooms}
                    image={property.image}
                    description={property.description}
                    iduser={property.userid}
                  />
            ))
        }
        </div>
    </header>
  )
}