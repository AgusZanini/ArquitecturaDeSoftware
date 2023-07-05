import React from "react";
import Cookies from "universal-cookie";
export const PropertyItems = (
    { id,
        tittle,
        size,
        bathrooms,
        service,
        city,
        state,
        country,
        street,
        price,
        rooms,
        image,
        description,
        iduser
    }) => {
      const cookies = new Cookies();
      const id_user = cookies.get("user").split(",")[0];
  return (
    <div class="Property">
      <div>
        <div class="Property_img">
          <img class="image" src={image} alt="" />
        </div>
      </div>
      <div class="Property_footer">
        <h1>{tittle}</h1>
        <h2>{description}</h2>
        {id_user==iduser ?<div>
          <p class="Datos-3">DUEÑO</p>
          <p class="price">U$S {price}</p>
          </div>:
          <div>
        <p class="price">U$S {price}</p>
        <p class="Datos">{size} m2</p>
        <p class="Datos">{rooms} Ambientes</p>
        <p class="Datos">{bathrooms} Baños</p>
        <p class="Datos">{service}</p>
        <p class="Datos-2">Direccion: {street},{city},{state}</p>
        <p class="Datos-2">{country}</p>
          </div>
          }
      </div>
    </div>
  );

}
