import { useState, useEffect } from 'react';
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

// REDUX ACTIONS
import { deleteToken } from "../states/actions/token"

// CONSTANTS
import orderTypes from "../constants/orderTypes";

// SERVICES
import OrderService from '../services/OrderService';
import AssetService from '../services/AssetService';

// ALERT
import { useAlert } from '@blaumaus/react-alert'



const CreateOrderPage = () => {

    const token = useSelector(state => state.token);

    const dispatch = useDispatch();

    const navigate = useNavigate();

    const alert = useAlert();


    const [orderType, setOrderType] = useState(orderTypes[0].id + "");

    const [buyAssets, setBuyAssets] = useState([]);
    const [buyAsset, setBuyAsset] = useState(null);
    
    const [sellAssets, setSellAssets] = useState([]);
    const [sellAsset, setSellAsset] = useState(null);
    const [sellAmount, setSellAmount] = useState("");

    const [availableAmount, setAvailableAmount] = useState(0);

    const [limitAmount, setLimitAmount] = useState("");

    const [exchangeRate, setExchangeRate] = useState(0);


    const orderService = new OrderService(token);
    const assetService = new AssetService(token);
    

    async function init(){

        let response = await orderService.orderCreateInfo();

        if (response === null) {
            alert.error("Ooops something went wrong.")
        }
        else if (200 <= response.Status && response.Status <= 299) {
            setBuyAssets(response.Payload.buy_assets);
            setSellAssets(response.Payload.sell_assets);
        }
        else if (response.Status == 401) {
            dispatch(deleteToken());
            navigate("/login");
        }
        else {
            alert.error(response.Message);
        }
    }


    useEffect(() => {
        init();
    }, [])


    useEffect(() => {
      getExchangeRate();
    }, [buyAsset, sellAsset])
    


    const createOrderHandler = async (e) => {
        e.preventDefault();

        let response = await orderService.createOrder(buyAsset, sellAsset, orderType, sellAmount, limitAmount);
        if (response === null) {
            alert.error("Ooops something went wrong.")
        }
        else if (200 <= response.Status && response.Status <= 299) {
            alert.success(response.Message);
            navigate("/profile");
        }
        else if (response.Status == 401) {
            dispatch(deleteToken());
            navigate("/login");
        }
        else {
            alert.error(response.Message);
        }
    }


    const getExchangeRate = async () => {

        let response = await assetService.getExchangeRate(buyAsset, sellAsset);
        if (200 <= response.Status && response.Status <= 299) {
            setExchangeRate(1 / response.Payload);
        } else {
            setExchangeRate(0);
        }

    }


    const limitAmountChangeHandler = (e) => {
        if (e.target.value >= 0) {
            setLimitAmount(e.target.value);
        }
        else {
            setLimitAmount("");
        }
    }


    const orderTypeChangeHandler = (e) => {
        setOrderType(e.target.value);
    }


    const buyAssetChangeHandler = (e) => {
        setBuyAsset(e.target.value);
        getExchangeRate();
    }


    const sellAmountChangeHandler = (e) => {

        if (e.target.value >= 0 && e.target.value <= availableAmount) {
            setSellAmount(e.target.value);
        }
        else {
            setSellAmount("");
        }
    }


    const sellAssetChangeHandler = async (e) => {
        
        let asset = parseInt(e.target.value);
        setSellAsset(asset);

        if(asset > 0){
            setAvailableAmount(sellAssets.filter((elem)=> elem.asset_id === asset)[0].available_amount);
        }
        else {
            setAvailableAmount(0);
        }

        setSellAmount("");
    }


    return (
        <div id="order" className="mb-5">
            <div className="container mt-5">
                <div id="order-row" className="row justify-content-center align-items-center">
                    <div id="order-column" className="col-md-6">
                        <div id="order-box" className="col-md-12 border border-success rounded">
                            <form id="order-form" className="form" onSubmit={createOrderHandler}>

                                <h5 className="text-center mt-3">Create Order</h5>

                                {/* ORDER TYPE */}
                                <div className="form-group">
                                    <label>Order Type:</label>
                                    <select className="custom-select" value={orderType} onChange={orderTypeChangeHandler}>
                                        {orderTypes.map((elem) =>
                                            <option value={elem.id} key={elem.id}>{elem.name}</option>
                                        )}
                                    </select>
                                </div>

                                <div className="form-group">
                                    {exchangeRate > 0 &&  <label>Exchange Rate: {exchangeRate}</label>}
                                </div>

                                {/* ASSETS */}
                                <div className="d-flex flex-row">

                                    {/* BUY ASSET */}
                                    <div className="p-2 flex-fill">
                                        <div className="text-center">Buy</div>
                                        <select className="custom-select" value={buyAsset} onChange={buyAssetChangeHandler}>
                                            <option selected>Select Asset</option>
                                            {buyAssets.map((elem) => 
                                                <option key={"buyAsset" + elem.ID} value={elem.ID}>{elem.code}</option>
                                            )}
                                        </select>
                                        <div className="mt-3 align-middle p-2">
                                            {exchangeRate > 0 && "~" + (sellAmount / exchangeRate).toFixed(6).toLocaleString("en-US")}
                                        </div>
                                    </div>

                                    {/* SELL ASSET */}
                                    <div className="p-2 flex-fill">
                                        <div className="text-center">Sell</div>
                                        <select className="custom-select" value={sellAsset} onChange={sellAssetChangeHandler}>
                                            <option selected>Select Asset</option>
                                            {sellAssets.map((elem) => 
                                                <option key={"sellAsset" + elem.asset_id} value={elem.asset_id}>{elem.code}</option>
                                            )}
                                        </select>
                                        <div className="mt-3">
                                            <input type="number" className="form-control" value={sellAmount} onChange={sellAmountChangeHandler}/>
                                        </div>
                                        <small>
                                            {availableAmount !== 0 && "Available: " + availableAmount}
                                        </small>
                                    </div>
                                </div>

                                {/* LIMIT */}
                                {!["10001"].includes(orderType) &&
                                    <div className="form-group mt-2 p-2">
                                        <label htmlFor="limit">Limit:</label><br />
                                        <input type="number" name="limit" id="limit" className="form-control" value={limitAmount} onChange={limitAmountChangeHandler}/>
                                    </div>
                                }

                                {/* SUBMIT */}
                                <div className="form-group text-center mt-2">
                                    <input type="submit" name="submit" className="btn btn-success btn-md" value="Create" />
                                </div>

                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}


export default CreateOrderPage;
