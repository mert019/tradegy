import { transitions, positions } from '@blaumaus/react-alert'



const Alert = ({ style, options, message, close }) => {

    const getClassName = (type) => {
        switch (type) {
            case 'info': return 'info';
            case 'success': return 'success';
            case 'error': return 'danger';
            default: return '';
        }
    }

    return (
        <div style={style}>
            <div className={"d-flex flex-row alert alert-" + getClassName(options.type)} role="alert">
                <div className="p-2">
                    <span>{message}</span>
                </div>
                <div onClick={close} className="p-2">
                    <b>X</b>
                </div>
            </div>
        </div>);
};

// Alert options
export const alertOptions = {
    position: positions.TOP_CENTER,
    timeout: 5000,
    offset: '30px',
    transition: transitions.SCALE
}

export default Alert;
