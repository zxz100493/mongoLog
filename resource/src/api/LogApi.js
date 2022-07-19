import axios from '@/plugins/axios'

// interface LogInfoParams {
//     LogId:String
// }
// export function apiGetLogInfo(params: LogInfoParams) {
//     let that = this
//     this.axios.get("/api/db.json")
//     .then(function(res){
//       console.log(res.data.companies)
//       that.customers=res.data.users
//     })
// }
const API = {}

// 登录
API.login = function (params) {
  return axios.get('/test', params)
}

export default API
