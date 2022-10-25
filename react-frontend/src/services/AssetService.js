import ApiService from "./ApiService";


class AssetService extends ApiService {


    constructor(token) {
        super(token);
    }


    async getWealthInformation() {

        var myHeaders = new Headers();
        myHeaders = this.appendAuthTokenToHeader(myHeaders);

        var requestOptions = {
            method: 'GET',
            headers: myHeaders,
            redirect: 'follow'
        };

        const url = this.getBaseApiUrl() + "/api/v1/asset/wealthinfo";

        return await this.sendApiRequest(url, requestOptions);
    }


    async getExchangeRate(buyAssetId, sellAssetId) {

        var myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");

        var raw = JSON.stringify({
            "buy_asset_id": parseInt(buyAssetId),
            "sell_asset_id": parseInt(sellAssetId)
        });

        var requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };

        const url = this.getBaseApiUrl() + "/api/v1/asset/exchangerate";

        return await this.sendApiRequest(url, requestOptions);
    }

}


export default AssetService;
