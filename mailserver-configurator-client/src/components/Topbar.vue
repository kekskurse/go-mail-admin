<template>
    <div>
        <v-app-bar
                color="#212121"

                dark
        >


            <v-toolbar-title>Go Mail Admin</v-toolbar-title>
            <v-icon :style="{ color: color}">mdi-heart</v-icon>

            <v-spacer></v-spacer>

            <v-btn to="/" icon>
                <v-icon>mdi-view-dashboard-variant</v-icon>
            </v-btn>
            <v-btn to="/domains" icon>
                <v-icon>mdi-dns</v-icon>
            </v-btn>
            <v-btn to="/alias" icon>
                <v-icon>mdi-forwardburger</v-icon>
            </v-btn>
            <v-btn to="/account" icon>
                <v-icon>mdi-account</v-icon>
            </v-btn>
            <v-btn to="/tls" icon>
                <v-icon>mdi-security</v-icon>
            </v-btn>
            <a href="https://github.com/kekskurse/go-mail-admin">
                <v-btn icon>
                    <v-icon>mdi-git</v-icon>
                </v-btn>
            </a>

            <v-menu
                    left
                    bottom
            >

            </v-menu>
        </v-app-bar>
    </div>
</template>
<script>
    import Client from "../service/Client";
    export default {
        name: 'Topbar',

        methods: {
            checkStatus: function() {
                Client.getStatus().then((res) => {
                    if(res.data == "Ok") {
                        this.color = "green";
                    } else {
                        this.color = "red";
                    }
                }).catch(() => {
                    this.color = "red";
                })
            }
        },
        mounted: function() {
          this.checkStatus();
        },

        data: () => ({
            color: "gray"
        }),
    }
</script>

