import axios from 'axios'

export default() => {
    let url;
    url = process.env.VUE_APP_API_URL
    console.log(process.env)
    if(process.env.VUE_APP_DYNAMIC_URL == "true") {
        url = window.location.href.substr(0,  window.location.href.indexOf("#"))
        console.log("Use dynamic API URL: " + url)
    } else {
        console.log("Use static API URL: " + url)
    }
    var header = {};
    header["Accept"] = 'application/json';
    header["Content-Type"] = 'application/json';
    if(localStorage.getItem("token") != null) {
        header["X-APITOKEN"] = localStorage.getItem("token");
    }

    let a;
     a = axios.create({
        baseURL: url,
        withCredentials: false,
        headers: header
    });
    a.interceptors.response.use(response => {
        return response;
    }, error => {
        if (error.response.status === 401) {
            location.href = "/#/login";
        }
        return error;
    });
    return a;
}
