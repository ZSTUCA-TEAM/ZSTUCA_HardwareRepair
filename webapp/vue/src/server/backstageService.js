import { useBackstageStore } from '@/stores/backstage'
import axios from 'axios'

// 发送管理员登录请求至后端,返回Promise,其匿名函数参数为状态码
export const postLogin = (admin) => {
    return axios.post('http://localhost:25555/bs/login', admin)
        .then((response) => {
            useBackstageStore().username = admin.username
            useBackstageStore().jwt = response.data
            return response.status
        })
}

export const getApplyBy = (type) => {
    return axios.get(`http://localhost:25555/bs/apply/${type}`, {
        headers: {
            Authorization: `${useBackstageStore().jwt}`
        }
    })
        .then((response) => {
            return response
        })
}