import React from 'react'
import Message from './form'

const NewMessage = ({form, onChange, onSubmit}) => (
    <div className="NewMessage_Lateral_Spaces row">

        <div className="col-sm NewMessage_Form_Space">
            <Message
                onChange={onChange}
                onSubmit={onSubmit}
                form={form}
            />            
        </div>
    </div>
)

export default NewMessage