import { useState, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

// REDUX ACTIONS
import { deleteToken, setToken } from "../states/actions/token"

// SERVICES
import AssetService from "../services/AssetService";
import OrderService from "../services/OrderService";

// DEVEXTREME COMPONENTS
import Accordion, { Item } from 'devextreme-react/accordion';

// COMPONENTS
import WealthInformation from "../components/WealthInformation";
import UserAssetList from "../components/UserAssetList";
import UserOrderHistory from "../components/UserOrderHistory";

// ALERT
import { useAlert } from '@blaumaus/react-alert'



const ProfilePage = () => {


  const token = useSelector(state => state.token);

  const dispatch = useDispatch();

  const navigate = useNavigate();


  const [wealthInfo, setWealthInfo] = useState({});
  const [orderHistory, setOrderHistory] = useState({});


  const assetService = new AssetService(token);
  const orderService = new OrderService(token);


  const setWealthInformation = async () => {

    let response = await assetService.getWealthInformation();
    
    // Error
    if (response === null) {
      alert.error("Ooops something went wrong.")
    }
    
    // Unauthorized
    else if(response.Status === 401){
      dispatch(deleteToken());
      navigate("/login");
    }
    
    // Success
    else if(200 <= response.Status && response.Status <= 299){
      setWealthInfo(response.Payload);
    }

    // Error
    else {
      if(response.Message.length > 0){
        alert.error(response.Message)
      }
    }

  }


  const setOrderHistoryList = async () => {
    let response = await orderService.getOrderHistory(token);
    if (response === null) {
      alert.error("Error on order service. Please try a few minites later.");
      return;
    }
    // Response actions
    if (200 <= response.Status && response.Status <= 299) {
      setOrderHistory(response.Payload);
    } else if (response.Status == 401) {
      dispatch(deleteToken());
      navigate("/login");
    }
    else {
      // alert.error(response.Message);
    }
  }


  useEffect(() => {
    async function initPage() {
      await setWealthInformation();
      await setOrderHistoryList();
    }
    initPage();
  }, [])


  return (
    <div className='container pt-3 pb-3'>
      <Accordion>
        <Item title="Wealth Information">
          <WealthInformation wealthInfo={wealthInfo} />
        </Item>
        <Item title="Asset List">
          <UserAssetList wealthInfo={wealthInfo} />
        </Item>
        <Item title="Order History">
          <UserOrderHistory orderHistory={orderHistory}/>
        </Item>
      </Accordion>
    </div>
  )
}


export default ProfilePage;
