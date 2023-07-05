import React, { useContext, useEffect, useState} from "react";
import Cookies from "universal-cookie";
import swal from "sweetalert2";
import '../Main/main.css';
import NewMessage from "./newMessage"
import { PropertyItems } from "./PropertyItems";
import { MessageItems } from "./MessageItems";

export const InsertMessage = () => {

    const [properties, setProperties] = useState([]);
    const [messages, setMessages] = useState([]);
    const cookies = new Cookies();
    const propertyId = cookies.get("propertyId");
    const id_user = cookies.get("user").split(",")[0];
    
    const [form, setForm] = useState({
        'userid': id_user,
        'propertyid': propertyId,
        'body': "",
      });
      const handleChange = (event) => {
        const { name, value } = event.target;
        setForm({
          ...form,
          [name]: value,
        });
    }
        const handleSubmit = async event => {
            event.preventDefault();
            if(form.body !=""){
                const requestOptions = {
              method: "POST",
              headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Access-Control-Allow-Origin': '*'
              },
              body: JSON.stringify({
                ...form,
                userid:Number(form.userid)
              }),
            };
            try {
              const response = await fetch('http://localhost:9001/message', requestOptions);
              if (response.ok) {
                const data = await response.json();
                swal.fire({
                  icon: 'success',
                  text: "Mensaje Enviado",
              }).then((result) => {
                if (result.isConfirmed) {
                    window.location.reload();
                }});
                // Aquí podrías hacer algo con la respuesta, si es necesario
              } else {
                throw new Error("No se pudo enviar el mensaje");
              }
            } catch (error) {
              console.error("Error:", error);
              swal.fire({
                icon: 'error',
                text: "No se pudo enviar el mensaje",
              })
            }
            console.log(form)}
            else{
        swal.fire({
          icon: 'error',
          text: "No se pudo enviar el mensaje",
        }) 
      }
      }; 

      const fetchApiProperty = async () => {
        const response = await fetch('http://localhost:8081/search/'+ propertyId)
          .then((response) => response.json());
          setProperties(response)
          console.log(response); 
      };
      useEffect(() => {
        fetchApiProperty();
      }, []);

      const fetchApiMessages = async () => {
        const response = await fetch('http://localhost:8081/properties/'+ propertyId +'/messages')
          .then((response) => response.json());
          setMessages(response)
          console.log(response);
      };
      useEffect(() => {
        fetchApiMessages();
      }, []);
    
    return (
        <div>
<div class="Property-dad">
            {properties.map((property) => (
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
        {messages && messages.length > 0 ? messages.map((message) => (
  <MessageItems key={message.id}
    id={message.id}
    userid={message.userid}
    propertyid={message.propertyid}
    body={message.body}
    created_at={message.created_at}
  />
)) : <p class="consulta">Nadie ha realizado consultas</p>}
 
    <NewMessage
        form={form}
    onChange={handleChange}
    onSubmit={handleSubmit}
    />
        </div>
    );
    
};
