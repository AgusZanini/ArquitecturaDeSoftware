import React, { useContext, useEffect, useState} from "react";
import Cookies from "universal-cookie";
import swal from "sweetalert2";

export const MessageItems = (
    { id,
        userid,
        usermame,
        propertyid,
        body,
        created_at
    }) => {
const [messages, setMessages] = useState([]);
const [users, setUsers] = useState('');
const cookies = new Cookies();
const id_user = cookies.get("user").split(",")[0];
const fetchApiUser = async () => {
  const response = await fetch('http://localhost:9000/user/'+ userid)
    .then((response) => response.json());
    setUsers(response)
    console.log(response); 
};
useEffect(() => {
  fetchApiUser();
}, []);

const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost:9001/messages/`+ id, {
        method: "DELETE",
      });
  
      if (response.ok) {
        // Actualizar el estado de la lista de elementos eliminando el elemento correspondiente
        // Puedes utilizar el método filter() para filtrar el elemento con el id correspondiente
        setMessages(messages.filter((message) => message.id !== id));
  
        // Mostrar mensaje de éxito
        swal.fire({
          icon: "success",
          text: "Mensaje eliminado",
        }).then((result) => {
            if (result.isConfirmed) {
                window.location.reload();
            }});
      } else {
        throw new Error("No se pudo eliminar el mensaje");
      }
    } catch (error) {
      console.error("Error:", error);
  
      // Mostrar mensaje de error
      swal.fire({
        icon: "error",
        text: "No se pudo eliminar el mensaje",
      });
    }
  };

  return (
    <div class="Message-padre">
          <div class="Message">
      <div class="Message_footer">
        <p class="date">{created_at}</p>
        <p class="body">{body}</p>
        {id_user == userid?
            <button class="eliminar" type="submit" onClick={() => handleDelete(id)}>Eliminar</button>
            :null}
      </div>
    </div>
    </div>
  );

}