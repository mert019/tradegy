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

    async cancelOrder(orderId) {

        var myHeaders = new Headers();
        myHeaders = this.appendAuthTokenToHeader(myHeaders);
        
        myHeaders.append("Content-Type", "text/plain");
        
        var raw = orderId;
        
        var requestOptions = {
          method: 'POST',
          headers: myHeaders,
          body: raw,
          redirect: 'follow'
        };

        const url = this.getBaseApiUrl() + "/api/v1/order/cancel";

        return await this.sendApiRequest(url, requestOptions);
    }

}

export default OrderService;
