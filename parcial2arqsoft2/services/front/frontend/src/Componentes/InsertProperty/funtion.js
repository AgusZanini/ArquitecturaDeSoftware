import React, { useContext, useEffect, useState} from "react";
import NewProperty from './newProperty'
import Cookies from "universal-cookie";
import swal from "sweetalert2";

export const InsertProperty = () => {
    const cookies = new Cookies();
    const id_user = cookies.get("user").split(",")[0];
    const [form, setForm] = useState([{
      "title": "",
      "userid": Number(id_user),
      "image": "",
      "currency": "",
      "price": "",
      "state": "",
      "condition": "",
      "address": ""
    }]);
  
    const handleChange = (event) => {
      const { name, value } = event.target;
      setForm({
        ...form,
        [name]: name === "userid"|| name === "price" || name === "state"? Number(value) : value,
      });
    };


    const handleSubmit = async event => {
      event.preventDefault();
      if(Number(form.price) > 0 && form.title != "" ){
        const requestOptions = {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
          'Access-Control-Allow-Origin': '*'
        },
        body: JSON.stringify({
          ...form,
          price: Number(form.price), // Convertir a float
          userid:Number(form.userid),
          state:Number(form.state)
        }),
      };
      try {
        const response = await fetch('http://localhost:8090/items', requestOptions);
        if (response.status == 201) {
          const data = await response.json();
          console.log(data);
          swal.fire({
            icon: 'success',
            text: "Propiedad Agregada",
        }).then((result) => {
          if (result.isConfirmed) {
            window.location.href = "http://localhost:3000/home"
          }});
          // Aquí podrías hacer algo con la respuesta, si es necesario
        } else {
          throw new Error("No se pudo agregar la propiedad1");
        }
      } catch (error) {
        console.error("Error:", error);
        swal.fire({
          icon: 'error',
          text: "No se pudo agregar la propiedad2",
        })
      }
      console.log(form)}
      else{
        swal.fire({
          icon: 'error',
          text: "No se pudo agregar la propiedad3",
        }) 
      }
    };

    return <NewProperty 
    form={form}
    onChange={handleChange}
    onSubmit={handleSubmit}
/>


};
