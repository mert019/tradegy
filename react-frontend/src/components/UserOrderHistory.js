import { useState } from 'react';

import DataGrid, { Scrolling, Paging, Column, HeaderFilter, Selection } from 'devextreme-react/data-grid';


import { Popup, Position, ToolbarItem } from 'devextreme-react/popup';
import { Button } from 'devextreme-react/button';

const UserOrderHistory = ({orderHistory}) => {


    const [popUpVisible, setPopUpVisible] = useState(false);

    const orderHistoryData = [
        {
            trade: "BTC/USD",
            order_type: "Market Order",
            status: "Executed",

        }
    ];

    const hidePopup = () => {
        setPopUpVisible(false);
      };
    
      const renderContent = () => {
        return (
            <>
            <p>Popup content</p>
            <h1>Cool</h1>
            <Button text='abc' onClick={() => setPopUpVisible(false)}/>
            </>
        );
    };
    
    return (
        <>
        <Popup
                    title="Order Details"
                    visible={popUpVisible}
                    contentRender={renderContent}
                    onHiding={hidePopup}
                />


            <DataGrid dataSource={orderHistory} onRowDblClick={(e) => setPopUpVisible(true)} height="300px">
                <Selection mode="single" showCheckBoxesMode="none" />
                {/* <Column dataField={"image_source"} caption="#" alignment="center" width={50} cellRender={(val) => { return val.value.length > 0 ? <><img className="small-asset-img" src={val.value} /></> : <span className="small-asset-img">$</span> }} /> */}
                {/* <Column dataField="amount" cellRender={(val) => { return <>{val.value.toFixed(6).toLocaleString("en-US")}</> }} /> */}
                {/* <Column dataField="usd_amount" caption="USD Amount" cellRender={(val) => { return <>$ {val.value.toFixed(2).toLocaleString("en-US")}</> }} /> */}
                {/* <Column dataField={"order_id"} /> */}
            </DataGrid>
        </>
    )
}

export default UserOrderHistory