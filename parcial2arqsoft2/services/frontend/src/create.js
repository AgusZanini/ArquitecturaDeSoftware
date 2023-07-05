import React, { useState, useEffect } from 'react';
import { Box, Button, Container, TextField, Typography } from '@mui/material';
import Cookies from "universal-cookie";
import swal from "sweetalert2";
import jwtDecode from "jwt-decode";
const Cookie = new Cookies();

const Create = () => {
  const [token, setToken] = useState("");
  const [userId, setUserId] = useState("");
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState(0);
  const [currency, setCurrency] = useState("");
  const [image, setImage] = useState("");
  const [state, setState] = useState(0);
  const [address, setAddress] = useState("");
  const [condition, setCondition] = useState("");

  useEffect(() => {
    loadTokenFromCookie();
  }, []);

  const loadTokenFromCookie = () => {
    const userToken = Cookie.get("user");
    setToken(userToken);

    if (userToken) {
      const decodedToken = jwtDecode(userToken);
      const userid = decodedToken.userid;
      setUserId(userid);
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    const parsedPrice = parseFloat(price);
    const parsedState = parseInt(state);
    const requestOptions = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify([
        {
          title,
          parsedPrice,
          currency,
          image,
          condition,
          parsedState,
          address,
          userid: userId,
        }
      ])
    };

    try {
      const response = await fetch('http://localhost:8090/items', requestOptions);
      if (response.status === 201) {
        swal.fire({
          text: "Publicación creada exitosamente",
          icon: 'success'
        }).then(() => {
          setTitle("");
          setPrice();
          setCurrency("");
          setState();
          setImage("");
          setAddress("");
          setCondition("");
        });
      } else {
        swal.fire({
          text: "Error al crear la publicación",
          icon: 'error'
        });
      }
    } catch (error) {
      swal.fire({
        text: "Error en la solicitud",
        icon: 'error'
      });
    }
  };

  return (
    <Container maxWidth="sm">
      <Box
        sx={{
          marginTop: 4,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Typography component="h2" variant="h5">
          Crear Publicación
        </Typography>
        <Box component="form" sx={{ mt: 1 }}>

          <TextField label="Título" type="text" value={title} onChange={(event) => setTitle(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Price" type="number" value={price} onChange={(event) => setPrice(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Currency" type="text" value={currency} onChange={(event) => setCurrency(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="image (URL)" type="url" value={image} onChange={(event) => setImage(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="State" type="text" value={state} onChange={(event) => setState(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="address" type="text" value={address} onChange={(event) => setAddress(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="condition" type="text" value={condition} onChange={(event) => setCondition(event.target.value)} variant="outlined" margin="normal" required fullWidth />
   
          <Button type="button" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }} onClick={handleSubmit}>
            Crear Publicación
          </Button>

        </Box>
      </Box>
    </Container>
  );
};

export default Create;
