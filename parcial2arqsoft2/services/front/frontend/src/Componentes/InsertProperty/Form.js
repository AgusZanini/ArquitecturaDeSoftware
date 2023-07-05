import React from 'react'
import './insert.css'

function imagen(id) {
    var x = document.getElementById("image").value;
    if(x == ""){
    document.getElementById("div-imagen-form").style.display = "none";
    }
    else{
        document.getElementById("div-imagen-form").style.display = "block";
        document.getElementById(id).src = x;
    } 
}

const Form = ({ onChange, onSubmit, form }) => (
    <div className="container" id="container">
        <h1 className='publicar'>
            Publicar propiedad
        </h1>
        <form className='form-publicar'
            onSubmit={onSubmit}
        >
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Titulo"
                    name="title"
                    onChange={onChange}
                    value={form.title}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Imagen"
                    name="image"
                    id="image"
                    onChange={onChange}
                    onBlur={() => imagen("image-2")}
                    value={form.image}
                />
            </div>
            {/*
            <div className='div-imagen-form' id="div-imagen-form">
            <img className="imagen" id="imagen-2" src='' />
            </div>*/}
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputpb"
                    placeholder="Moneda"
                    name="currency"
                    onChange={onChange}
                    value={form.currency}
                />
            </div>
            <div className="form-group">
                <input
                    type="number"
                    className="form-control, inputpb"
                    placeholder="Precio"
                    name="price"
                    onChange={onChange}
                    value={form.price}
                />
            </div>
            <div className="form-group">
                <input
                    type="number"
                    className="form-control, inputp"
                    placeholder="Estado"
                    name="state"
                    onChange={onChange}
                    value={form.state}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Condicion"
                    name="condition"
                    onChange={onChange}
                    value={form.condition}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Direccion"
                    name="address"
                    onChange={onChange}
                    value={form.address}
                />
            </div>
            <div>
                <button
                    type="submit"
                    className="btn-summit"
                >
                    Publicar
                </button>
            </div>
        </form>
    </div>
)



export default Form