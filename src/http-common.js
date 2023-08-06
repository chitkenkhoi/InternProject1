import axios from 'axios';

const HTTP = axios.create({
    baseURL: `http://localhost:8080/api/`,
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
})
export default HTTP