import DataGrid, { Column, Selection, Sorting } from 'devextreme-react/data-grid';



const UserAssetList = ({ wealthInfo }) => {

    return (
        <DataGrid dataSource={wealthInfo}>

            <Selection mode="single" showCheckBoxesMode="none" />
            <Sorting mode="none"/>
            
            <Column dataField={"image_source"} caption="#" alignment="center" width={70} cellRender={(val) => { return val.value.length > 0 ? <><img className="small-asset-img" src={val.value} /></> : <span className="small-asset-img">$</span> }} />
            <Column dataField={"name"} alignment={"center"} />
            <Column dataField="amount" cellRender={(val) => { return <>{val.value.toFixed(6).toLocaleString("en-US")}</> }} />
            <Column dataField="usd_amount" caption="USD Amount" cellRender={(val) => { return <>$ {val.value.toFixed(2).toLocaleString("en-US")}</> }} />

        </DataGrid>
    )
}


export default UserAssetList
