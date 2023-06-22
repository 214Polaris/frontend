import axios from "axios";
axios.defaults.baseURL = "http://localhost:8081/";
/**
 * 请求拦截
 */
axios.interceptors.request.use(
  (config) => {
    // 在发送请求之前做些什么 验证token之类的
    const token = localStorage.getItem("token");
    const User = localStorage.getItem("userID");
    config.headers["token"] = token; //在请求头设置token
    return config;
  },
  (error) => {
    // 对请求错误做些什么
    return Promise.error(error);
  }
);

// 导出给 main.js 挂载
export default axios;
