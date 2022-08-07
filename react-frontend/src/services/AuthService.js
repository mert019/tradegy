const login = async (username, password) => {
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
        const url = process.env.REACT_APP_API_URL + "/api/v1/auth/login";
        const response = await fetch(url, requestOptions);
        return await response.json();
    } catch (error) {
        return null;
    }
}

const AuthService = {
    login: login,
}

export default AuthService;
