import goodsList from "../static/goodsList.json";
import { createStore } from "vuex";

const state = {
  userID: window.sessionStorage.getItem("userID"),
  token: window.sessionStorage.getItem("token"),
  goodsList,
  is_login: false,
};
const mutations = {
  SET_TOKEN: (state, data) => {
    state.token = data;
    localStorage.setItem("token", data);
    state.is_login = true;
  },
  GET_USER: (state, data) => {
    state.userID = data;
    localStorage.setItem("userID", data);
  },
  LOGOUT: (state) => {
    state.token = null;
    state.user = null;
    localStorage.removeItem("token");
    localStorage.removeItem("userID");
  },
};

const actions = {};
export default createStore({
  state,
  mutations,
  actions,
});
