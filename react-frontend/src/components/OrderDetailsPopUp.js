// DEVEXTREME
import { Button } from 'devextreme-react/button';
import { Popup } from 'devextreme-react/popup';
import { confirm } from 'devextreme/ui/dialog';

// ALERT
import { useAlert } from '@blaumaus/react-alert'



const OrderDetailsPopUp = ({visible, setVisibility, orderData, updateData, orderService}) => {
    

    const alert = useAlert();

    const isOrderOpen = orderData !== null ? orderData.order_status === "Open" : false;
    const isOrderExecuted  = orderData !== null ? orderData.order_status === "Executed" : false;
    const isOrderUseLimit = orderData !== null ? !orderData.order_type.includes("MarketOrder") : false;

    const renderContent = () => {
        return (
            <>

                <div className="d-flex justify-content-center align-items-center">
                    {
                        orderData.buy_asset_image_source.length > 0 ?
                            <img className="small-asset-img" src={orderData.buy_asset_image_source} /> : <span className="small-asset-img">$</span>
                    }

                    <span className="m-1 h4">{orderData.buy_asset_code}</span> <span className="h4 m-1">/</span> <span className="m-1 h4">{orderData.sell_asset_code}</span>

                    {
                        orderData.sell_asset_image_source.length > 0 ?
                            <img className="small-asset-img" src={orderData.sell_asset_image_source} /> : <span className="small-asset-img">$</span>
                    }
                </div>

                <div className="mt-4">

                    <div className="mb-2">
                        <b>Order Status:</b> {orderData.order_status}
                    </div>

                    <div className="mb-2">
                        <b>Order Type:</b> {orderData.order_type}
                    </div>

                    {isOrderUseLimit &&
                        <div className="mb-2">
                            <b>Limit:</b> {orderData.limit}
                        </div>
                    }

                    {isOrderExecuted &&
                        <div className="mb-2">
                            <b>Buy Amount:</b> {orderData.buy_amount} {orderData.buy_asset_code}
                        </div>
                    }

                    <div className="mb-2">
                        <b>Sell Amount:</b> {orderData.sell_amount} {orderData.sell_asset_code}
                    </div>

                    <div className="mb-2">
                        <b>Created At:</b> {(new Date(orderData.created_at)).toLocaleString("en-US")}
                    </div>

                    {isOrderExecuted &&
                        <div className="mb-2">
                            <b>Executed At:</b> {(new Date(orderData.executed_at)).toLocaleString("en-US")}
                        </div>
                    }

                </div>

                {isOrderOpen &&
                    <div className="d-flex justify-content-center mt-4">
                        <Button text="Cancel"
                            type="danger"
                            onClick={() => cancelOrderBtnHandler(orderData.order_id)} />
                    </div>
                }
            </>
        );
    }


    const hidePopup = async () => {
        setVisibility(false);
    };


    const cancelOrder = async(orderId) => {

        let response = await orderService.cancelOrder(orderId);

        // Error
        if (response === null) {
            alert.error("Ooops something went wrong.");
        }
        // Success
        else if (200 <= response.Status && response.Status <= 299) {
            alert.success(response.Message);
        }
        // Error
        else {
            if(response.Message.length > 0){
              alert.error(response.Message);
            }
          }
    }


    const cancelOrderBtnHandler = async(orderId) => {
        confirm("The order will be cancelled. Are you sure?", "Confirm changes")
            .then((dialogResult) => {
                if (dialogResult === true) {
                    cancelOrder(orderId);
                    hidePopup();
                    updateData();
                }
            })
    }


    return (
        <Popup
            title="Order Details"
            visible={visible}
            contentRender={renderContent}
            onHiding={hidePopup}
            maxHeight={"360px"}
            maxWidth={"400px"}
            dragEnabled={false}
            showCloseButton={true}
        />
    );
}


export default OrderDetailsPopUp;
