import axios from 'axios'

export default() => {
    return axios.create({
        baseURL: `/`,
        withCredentials: false,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    })
}
