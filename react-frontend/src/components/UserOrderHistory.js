import { useState } from 'react';

// DEVEXTREME
import DataGrid, { Scrolling, Paging, Column, HeaderFilter, Selection, FilterRow } from 'devextreme-react/data-grid';

// COMPONENTS
import OrderDetailsPopUp from './OrderDetailsPopUp';

// UTILS
import { getOrderListStatusBadgeStyle } from '../utils/StyleHelper';



const UserOrderHistory = ({orderHistory, updateData, orderService}) => {


    const [popUpVisible, setPopUpVisible] = useState(false);
    const [orderDetails, setOrderDetails] = useState(null);


    return (
        <>
    
            <OrderDetailsPopUp
                visible={popUpVisible}
                setVisibility={setPopUpVisible} 
                orderData={orderDetails} 
                updateData={updateData}
                orderService={orderService} />

            <DataGrid dataSource={orderHistory} onRowDblClick={(e) => { setOrderDetails(e.data); setPopUpVisible(true); }} height="300px">

                <HeaderFilter visible={true} />
                <Selection mode="single" showCheckBoxesMode="none" />
                <FilterRow visible={true} applyFilter={"onclick"} />

                <Column dataField={"order_id"} caption="ID" alignment={"center"} allowHeaderFiltering={false} allowFiltering={false} allowSorting={true} width={"100px"}/>

                <Column dataField={"buy_asset_code"} 
                    caption="Buy" 
                    alignment={"left"} 
                    cellRender={
                        (val) => {
                            return <>
                                {val.key.buy_asset_image_source.length > 0 ? <><img className="small-asset-img" src={val.key.buy_asset_image_source} /></> : <span className="small-asset-img">$</span>}
                                <span className="ml-3">{val.value}</span>
                            </>
                        }
                    }/>

                <Column dataField={"sell_asset_code"} 
                    caption="Sell" 
                    alignment={"left"} 
                    cellRender={
                        (val) => {
                            return <>
                                {val.key.sell_asset_image_source.length > 0 ? <><img className="small-asset-img" src={val.key.sell_asset_image_source} /></> : <span className="small-asset-img">$</span>}
                                <span className="ml-3">{val.value}</span>
                            </>
                        }
                    }/>

                <Column dataField={"order_status"} 
                    caption="Order Status"
                    alignment={"center"}
                    cellRender={
                        (val) => { return <span className={'status-badge ' + getOrderListStatusBadgeStyle(val.value)}>{val.value}</span> }
                    } />

                <Column dataField={"order_type"} caption="Order Type" alignment={"center"} />

            </DataGrid>
        </>
    )
}


export default UserOrderHistory
