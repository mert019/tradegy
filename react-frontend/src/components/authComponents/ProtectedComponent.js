import { useSelector } from 'react-redux';

const ProtectedComponent = (props) => {

  const token = useSelector(state => state.token);

  return (token.length > 0 ?
    <>
      {props.children}
    </> :
    <></>)
}

export default ProtectedComponent;
