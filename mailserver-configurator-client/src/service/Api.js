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
    return axios.create({
        baseURL: url,
        withCredentials: false,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    })
}
