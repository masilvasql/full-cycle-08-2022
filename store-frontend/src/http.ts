import axios from "axios";

const http = axios.create({
    baseURL:'http://localhost:3000/api'//QDO for para o docker, o nome do container é app e a porta é 3000
})

export default http;