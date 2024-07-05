import axios from "axios"

export const postApply = (apply) => {
    return axios.post('/apply', apply)
        .then((response) => {
            console.log(response.status + "1")
            return response.status
        })
}