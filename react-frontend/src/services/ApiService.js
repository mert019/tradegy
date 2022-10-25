class ApiService {

    constructor(token) {

        this.token = token;

    }


    getBaseApiUrl() {
        return process.env.REACT_APP_API_URL;
    }


    appendAuthTokenToHeader(header) {
        header.append("Authorization", "Bearer " + this.token);
        return header;
    }


    async sendApiRequest(url, requestOptions) {
        try {
            const response = await fetch(url, requestOptions);
            return await response.json();
        }
        catch (error) {
            console.error("API_SERVICE:", error);
            return null;
        }
    }

}


export default ApiService;
