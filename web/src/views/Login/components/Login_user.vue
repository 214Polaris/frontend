<template>
  <div class="MainContainer">
    <div class="login-container">
      <div class="login-form" v-show="!isRegisterFormVisible">
        <h2>用户登录</h2>
        <form>
          <div class="form-group">
            <label>用户名</label>
            <input
              type="text"
              class="form-control"
              v-model.trim="username"
              placeholder="请输入用户名"
              autocomplete="off"
            />
            <div v-if="!isUsernameValid" class="invalid-feedback">
              用户名长度应在5到12个字符之间。
            </div>
          </div>
          <div class="form-group">
            <label>密码</label>
            <input
              type="password"
              class="form-control"
              v-model.trim="password"
              placeholder="请输入密码"
              autocomplete="off"
            />
          </div>
          <el-button
            type="primary"
            class="btn btn-primary"
            @click="login"
            :loading="isSending"
          >
            登录
          </el-button>
          <div class="form-group mt-3">
            没有账号？<a href="/#/login" @click="toggle">点击注册</a>
          </div>
        </form>
      </div>
      <div class="register-form" v-show="isRegisterFormVisible">
        <h2>用户注册</h2>
        <form>
          <div class="form-group">
            <label>用户名</label>
            <input
              type="text"
              class="form-control"
              v-model.trim="registerUsername"
              placeholder="请输入用户名"
              autocomplete="off"
            />
            <div v-if="!isRegisterUsernameValid" class="invalid-feedback">
              用户名长度应在5到12个字符之间。
            </div>
          </div>
          <div class="form-group">
            <label>密码</label>
            <input
              type="password"
              class="form-control"
              v-model.trim="registerPassword"
              placeholder="请输入密码"
              autocomplete="off"
            />
          </div>
          <div class="form-group">
            <label>确认密码</label>
            <input
              type="password"
              class="form-control"
              v-model.trim="confirmPassword"
              placeholder="请确认密码"
              autocomplete="off"
            />
            <div v-if="!isConfirmPasswordValid" class="invalid-feedback">
              两次输入的密码不一致。
            </div>
          </div>
          <el-button
            type="primary"
            class="btn btn-primary"
            @click="register"
            :loading="isSending"
          >
            注册
          </el-button>
          <div class="form-group mt-3">
            已有账号？<a href="/#/login" @click="toggle">点击登录</a>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { ElMessageBox } from "element-plus";
export default {
  data() {
    return {
      isSending: false,
      username: "",
      password: "",
      isUsernameValid: true,
      isButtonDisabled: false,
      isRegisterFormVisible: false,
      registerUsername: "",
      registerPassword: "",
      confirmPassword: "",
      isRegisterUsernameValid: true,
      isConfirmPasswordValid: true,
      isRegisterButtonDisabled: false,
    };
  },
  methods: {
    back() {
      this.$router.replace({ path: "/" });
    },
    toggle() {
      this.isRegisterFormVisible = !this.isRegisterFormVisible;
      this.username = "";
      this.password = "";
      this.isUsernameValid = true;
      this.isButtonDisabled = false;
      this.registerUsername = "";
      this.registerPassword = "";
      this.confirmPassword = "";
      this.isRegisterUsernameValid = true;
      this.isConfirmPasswordValid = true;
      this.isRegisterButtonDisabled = false;
    },
    login() {
      if (this.username.length < 5 || this.username.length > 12) {
        this.isUsernameValid = false;
      } else {
        this.isUsernameValid = true;
      }
      if (this.isUsernameValid && this.password.length > 0) {
        this.isSending = true;
        // 发送登录请求
        axios({
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          method: "post",
          url: "/api/Login",
          data: {
            user: this.username,
            pass: this.password,
          },
        })
          .then((response) => {
            console.log(response);
            if (response.data.code === 200) {
              this.$store.commit("SET_TOKEN", response.data.token);
              this.$store.commit("GET_USER", response.data.user);
              console.log("登录成功！");
              location.reload();
              this.$router.replace({ path: "/" });
            }
          })
          .catch(function () {
            ElMessageBox.alert("您输入的账号或密码错误").then(() => {
              location.reload();
              this.$router.go(0);
            });
          });
      }
    },
    register() {
      if (
        this.registerUsername.length < 5 ||
        this.registerUsername.length > 12
      ) {
        this.isRegisterUsernameValid = false;
      } else {
        this.isRegisterUsernameValid = true;
      }
      if (this.registerPassword !== this.confirmPassword) {
        this.isConfirmPasswordValid = false;
      } else {
        this.isConfirmPasswordValid = true;
      }
      if (this.isRegisterUsernameValid && this.isConfirmPasswordValid) {
        this.isSending = true;
        // 发送注册请求
        axios({
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          method: "post",
          url: "/api/Register",
          data: {
            user: this.registerUsername,
            pass: this.registerPassword,
          },
        })
          .then((response) => {
            if (response.data.code === 200) {
              this.$store.commit("SET_TOKEN", response.data.token);
              this.$store.commit("GET_USER", response.data.user);
              this.$message({
                message: "注册成功",
                type: "success",
              });
              location.reload();
              this.$router.replace({ path: "/" });
            }
          })
          .catch(function () {
            ElMessageBox.alert("用户已存在！").then(() => {
              location.reload();
              this.$router.go(0);
            });
          });
      }
    },
  },
};
</script>

<style lang="scss">
.login-container {
  display: relative;
  height: 450px;
  background-color: #f5f5f5;
  background: url(@/assets/LoginBackground.png) no-repeat center / cover;
}

.login-form {
  position: relative;
  text-align: center;
  left: 900px;
  top: 20px;
  height: 350px;
  width: 250px;
  padding: 32px;
  border-radius: 4px;
  background-color: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}

.login-form h2 {
  margin-bottom: 24px;
  text-align: center;
}

.form-group {
  margin-bottom: 16px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
}

.form-control {
  display: block;
  width: 100%;
  padding: 8px;
  font-size: 16px;
  border-radius: 4px;
  border: 1px solid #ccc;
}

.form-control:focus {
  outline: none;
  border-color: #6c757d;
  box-shadow: 0 0 5px rgba(108, 117, 125, 0.5);
}

.invalid-feedback {
  display: block;
  color: #dc3545;
  font-size: 14px;
  margin-top: 4px;
}

.btn-primary {
  display: block;
  width: 100%;
  padding: 8px;
  font-size: 16px;
  border-radius: 4px;
  border: none;
  background-color: #007bff;
  cursor: pointer;
  color: #fff;
}

.btn-primary:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.form-group mt-3 {
  text-align: center;
}

.form-group a {
  color: #007bff;
  cursor: pointer;
}

.register-form {
  position: relative;
  text-align: center;
  left: 900px;
  top: 20px;
  height: 350px;
  width: 250px;
  padding: 32px;
  border-radius: 4px;
  background-color: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}
</style>
