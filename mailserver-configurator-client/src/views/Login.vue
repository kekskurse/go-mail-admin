<template>
    <div class="container">
        <p v-if="loginFailed">Login failed</p>
        <v-text-field
                v-model="username"
                placeholder="Username"
                label="Username"
        ></v-text-field>
        <v-text-field
                v-model="password"
                placeholder="Password"
                label="Password"
                type="password"
        ></v-text-field>
        <v-btn @click="login" >Login</v-btn>
    </div>
</template>

<script>
    // @ is an alias to /src
    import Client from "../service/Client";
    export default {
        name: 'Home',
        components: {
        },
        data: () => ({
            "username": "",
            "password": "",
            "loginFailed": false
        }),
        methods: {
            login: function () {
                Client.login(this.username, this.password).then((res) => {
                    if(res.data.login==false) {
                        this.loginFailed = true;
                    } else {
                        localStorage.setItem("token", res.data.token)
                        this.$router.push("/home")
                    }
                });
            }
        }
    }
</script>
