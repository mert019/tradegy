const register =  async function RegisterService(username, password) {

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

    try {
        const url = process.env.REACT_APP_API_URL + "/api/v1/user/register";
        const response = await fetch(url, requestOptions);
        return await response.json();
    } catch (error) {
        return null;
    }
}

const RegisterService = {
    register: register,
}

export default RegisterService;
