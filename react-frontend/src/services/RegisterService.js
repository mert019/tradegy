import ApiService from "./ApiService";


class RegisterService extends ApiService {

    constructor() {
        super();
    }


    async register(username, password) {

        var myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");

        var raw = JSON.stringify({
            "username": username,
            "password": password
        });

        var requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };

        const url = this.getBaseApiUrl() + "/api/v1/user/register";
        
        return await this.sendApiRequest(url, requestOptions);
    }
}


export default RegisterService;
