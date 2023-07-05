import React from 'react'
import '../InsertProperty/insert.css'


const Message = ({ onChange, onSubmit, form }) => (
    <div className="mensaje_container">
        <h1 className='publicar_mensaje'>
            Enviar Mensaje
        </h1>
        <form className='form_publicar_mensaje'
            onSubmit={onSubmit}
        >
            <div className="form_group_mensaje">
                <input
                    type="text"
                    className="form-control, input_mensaje"
                    placeholder="Escriba un Mensaje para el Dueño de la Propiedad"
                    name="body"
                    onChange={onChange}
                    value={form.body}
                />
            </div>
            <div>
                <button
                    type="submit"
                    className="btn_summit_mensaje"
                >
                    Enviar
                </button>
            </div>
        </form>
    </div>
)



export default Message