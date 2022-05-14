import React, {useCallback, useContext, useEffect, useState} from "react";
import PropTypes from 'prop-types'
import {AppContext} from "../App";

const Toast = ({error}) => {
    const {setRootState} = useContext(AppContext)
    const [internalError, setInternalError] = useState(() => error)
    const [opacity, setOpacity] = useState(0)

    const close = useCallback(() => {
        setOpacity(0)
        setTimeout(() =>  setRootState({apiError: ""}), 750)
    }, [setRootState])

    useEffect(() => {
        if(error) {
            setInternalError(error)
            return
        }
        setTimeout(setInternalError, 0)
    }, [error])

    useEffect(() => {
        if(error) {
            setOpacity(1)
        }
    }, [internalError])

    if (!internalError) return null

    return (
        <div className="position-fixed bottom-0 end-0 p-1 container d-flex flex-row-reverse"
             style={{
                 zIndex: 99,
                 opacity: opacity,
                 transition: "opacity ease-in-out 0.75s"
             }}
        >
            <div className="d-flex flex-column rounded shadow col-6" role="alert" aria-live="assertive"
                 aria-atomic="true">
                <div className="toast-header">
                    <strong className="me-auto">🤑 Biller</strong>
                    <small>now</small>
                    <button
                        type="button"
                        className="btn-close"
                        data-bs-dismiss="toast"
                        aria-label="Close"
                        onClick={close}
                    />
                </div>

                <div className="toast-body">
                    {internalError}
                </div>
            </div>
        </div>
    )
}

Toast.propTypes = {
    error: PropTypes.string
}

export default React.memo(Toast)
