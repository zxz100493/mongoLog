'use strict'

// import Vue from 'vue';
import axios from 'axios'

// Full config:  https://github.com/axios/axios#request-config
// axios.defaults.baseURL = process.env.baseURL || process.env.apiUrl || '';
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

const config = {
  // baseURL: process.env.baseURL || process.env.apiUrl || ""
  // timeout: 60 * 1000, // Timeout
  // withCredentials: true, // Check cross-site Access-Control
}

const _axios = axios.create(config)

_axios.interceptors.request.use(
  function (config: any) {
    // Do something before request is sent
    return config
  },
  function (error: any) {
    // Do something with request error
    return Promise.reject(error)
  }
)

// Add a response interceptor
_axios.interceptors.response.use(
  function (response: any) {
    // Do something with response data
    return response
  },
  function (error: any) {
    // Do something with response error
    return Promise.reject(error)
  }
)

// Plugin.install = function (Vue, options) {
//   Vue.axios = _axios
//   window.axios = _axios
//   Object.defineProperties(Vue.prototype, {
//     axios: {
//       get () {
//         return _axios
//       }
//     },
//     $axios: {
//       get () {
//         return _axios
//       }
//     }
//   })
// }

// Vue.use(Plugin)

// export default Plugin
export default {
  install: function (app: any, options: any) {
    console.log(options)
    // add global method
    app.config.globalProperties.axios = _axios
    // add global method
    app.config.globalProperties.$translate = (key: any) => {
      return key
    }
  }
}
