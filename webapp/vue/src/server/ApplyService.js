import axios from "axios"

// 发送用户提交的报修请求至后端,返回Promise,其匿名函数参数为状态码
export const postApply = (apply) => {
    return axios.post('/apply', apply)
        .then((response) => {
            return response.status
        })
}