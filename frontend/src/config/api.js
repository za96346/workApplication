import axios from "axios";
import { urlPath, proto, ip, port } from "./config";
const axiosObj = (method = 'POST', url, data, token = null) => {
    console.log(proto+ '://' + ip + ':' + port + url)
    if (!token) {
        //login
        return(
            {
                method: method,
                url: proto+ '://' + ip + ':' + port + url,
                data: JSON.stringify({
                    ...data
                }),
                headers: {
                    'Content-Type': 'application/json',
                    'Access-Control-Allow-Origin': '*'
                }
            }
        )
    }
    else {
        return(
            {
                method: method,
                url: proto+ '://' + ip + ':' + port + url,
                data: JSON.stringify({
                    'data': data
                }),
                headers: {
                    'token': token,
                    'Content-Type': 'application/json',
                    'Access-Control-Allow-Origin': '*'
                }
            }
        )
    }

}

export async function login(data){
    return await axios(axiosObj('POST',urlPath.login, data = data)).then((response) => {
        console.log('login response',response.data)
    })
}