import ApiService from "./ApiService";


class OrderService extends ApiService {

    constructor(token) {
        super(token);
    }


    async orderCreateInfo() {

        var myHeaders = new Headers();
        myHeaders = this.appendAuthTokenToHeader(myHeaders);

        var requestOptions = {
            method: 'GET',
            headers: myHeaders,
            redirect: 'follow'
        };

        const url = this.getBaseApiUrl() + "/api/v1/order/createinfo";

        return await this.sendApiRequest(url, requestOptions);
    }


    async createOrder(buyAssetId, sellAssetId, orderTypeId, amount, limit) {

        var myHeaders = new Headers();
        myHeaders = this.appendAuthTokenToHeader(myHeaders);
        myHeaders.append("Content-Type", "application/json");

        var raw = JSON.stringify({
            "amount": parseFloat(amount),
            "order_type_id": parseInt(orderTypeId),
            "buy_asset_id": parseInt(buyAssetId),
            "sell_asset_id": parseInt(sellAssetId),
            "limit": parseFloat(limit)
        });

        var requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };

        const url = this.getBaseApiUrl() + "/api/v1/order/create";

        return await this.sendApiRequest(url, requestOptions);
    }


    async getOrderHistory() {

        var myHeaders = new Headers();
        myHeaders = this.appendAuthTokenToHeader(myHeaders);

        var requestOptions = {
            method: 'GET',
            headers: myHeaders,
            redirect: 'follow'
        };

        const url = this.getBaseApiUrl() + "/api/v1/order/allhistory";

        return await this.sendApiRequest(url, requestOptions);
    }

}

// const orderCreateInfo = async (token) => {

//     var myHeaders = new Headers();
//     myHeaders.append("Authorization", "Bearer " + token);

//     var requestOptions = {
//         method: 'GET',
//         headers: myHeaders,
//         redirect: 'follow'
//     };

//     try {
//         const url = process.env.REACT_APP_API_URL + "/api/v1/order/createinfo";
//         const response = await fetch(url, requestOptions);
//         return await response.json();
//     } catch (error) {
//         return null;
//     }
// }

// const createOrder = async (token, buyAssetId, sellAssetId, orderTypeId, amount, limit) => {

//     var myHeaders = new Headers();
//     myHeaders.append("Authorization", "Bearer " + token);
//     myHeaders.append("Content-Type", "application/json");

//     var raw = JSON.stringify({
//         "amount": parseFloat(amount),
//         "order_type_id": parseInt(orderTypeId),
//         "buy_asset_id": parseInt(buyAssetId),
//         "sell_asset_id": parseInt(sellAssetId),
//         "limit": parseFloat(limit)
//     });

//     var requestOptions = {
//         method: 'POST',
//         headers: myHeaders,
//         body: raw,
//         redirect: 'follow'
//     };

//     try {
//         const url = process.env.REACT_APP_API_URL + "/api/v1/order/create";
//         const response = await fetch(url, requestOptions);
//         return await response.json();
//     } catch (error) {
//         return null;
//     }
// }

// const getOrderHistory = async (token) => {
//     var myHeaders = new Headers();
//     myHeaders.append("Authorization", "Bearer " + token);

//     var requestOptions = {
//         method: 'GET',
//         headers: myHeaders,
//         redirect: 'follow'
//     };

//     try {
//         const url = process.env.REACT_APP_API_URL + "/api/v1/order/allhistory";
//         const response = await fetch(url, requestOptions);
//         return await response.json();
//     } catch (error) {
//         return null;
//     }
// }

// const OrderService = {
//     orderCreateInfo: orderCreateInfo,
//     createOrder: createOrder,
//     orderHistory: getOrderHistory,
// }

export default OrderService;
