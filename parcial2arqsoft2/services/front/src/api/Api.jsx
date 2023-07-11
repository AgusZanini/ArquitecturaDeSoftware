import React from 'react'

const fetchhome = async() => {
    const url = "http://localhost:8081/search= " 

    try {
        const response = await fetch(url)
        const responsejson = await response.json()
    } catch (error) {
        console.error("error fetching items: ", error)
    }

    return responsejson
}

const fetchuser = async(requestOptions) => {
    fetch('http://localhost:9000/login',requestOptions)
      .then(response => {if (response.status == 403) {
         swal.fire({
          text: "Datos incorrectos",
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
            window.location.replace("/")
            return response.json()
          }})
      }
      return response.json()})
      .then(response => {
          Cookie.set("user", response.token, {path: "/"})
  })
}