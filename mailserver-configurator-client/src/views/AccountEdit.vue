<template>
    <div>
        <v-container>
            <v-card>
                <v-card-title>Account {{account.username}}@{{account.domain}} </v-card-title>
                <v-card-text>
                    <span v-if="!account.id">
                        <v-text-field v-model="account.username" label="Username" placeholder="Username"></v-text-field>
                         <v-select v-model="account.domain" data-app="true" :items="domainNames" label="Destination-Domain"></v-select>
                        <v-text-field v-model="account.password" label="Password" type="password" placeholder="Password"></v-text-field>
                    </span>
                    <v-text-field v-model="account.quota" label="Quota" placeholder="Quota"></v-text-field>
                    <v-checkbox v-model="account.enabled" label="Enabled"></v-checkbox>
                    <v-checkbox v-model="account.sendonly" label="Send Only"></v-checkbox>
                    <span style="background-color:#BBDEFB; margin-left: 10px; border-radius: 5px; padding-top: 10px;padding-bottom:8px;">
                        <v-btn @click="saveAlias" icon><v-icon>mdi-content-save</v-icon></v-btn>
                    </span>
                </v-card-text>
            </v-card>
        </v-container>
        <v-container v-if="account.id">
            <v-card>
                <v-card-title>Change Password</v-card-title>
                <v-card-text>
                    <v-text-field type="password" v-model="password" label="New Password" placeholder="New Password"></v-text-field>
                    <span style="background-color:#BBDEFB; margin-left: 10px; border-radius: 5px; padding-top: 10px;padding-bottom:8px;">
                        <v-btn @click="changePassword()" icon><v-icon>mdi-content-save</v-icon></v-btn>
                    </span>
                </v-card-text>
            </v-card>
        </v-container>
    </div>
</template>

<script>
    import Client from "../service/Client";
    export default {
        name: 'AliasEdit',
        methods: {
            getAccounts: function () {
                Client.getAccounts().then((res) => {
                    for(var i = 0; i < res.data.length; i++) {
                        if(res.data[i].id == this.$route.params.id) {
                            this.account = res.data[i]
                        }
                    }
                });
            },
            getDomains: function () {
                Client.getDomains().then((res) => {
                    for(var i = 0; i < res.data.length; i++) {
                        this.domainNames.push(res.data[i].domain)
                    }
                    console.log(this.domainNames)
                });

            },
            saveAlias: function () {
                if(this.account.id) {
                    Client.saveAccount(this.account).then(() => {
                        this.getAccounts();
                        this.$swal("Account saved");
                        this.$router.push("/account")
                    })
                } else {
                    Client.createAccount(this.account).then(() => {
                        this.getAccounts();
                        this.$swal("Account created");
                        this.$router.push("/account")
                    })
                }
            },
            changePassword: function () {
                Client.changePassword(this.account.id, this.password).then(()=> {
                    this.$swal("Password changed");
                }).catch(() => {
                    alert("Oups, something go wrong")
                })
            }
        },

        mounted: function() {
            this.getDomains();
            this.getAccounts();

        },
        components: {

        },
        data: () => ({
            account: {"quota": 1024, "enabled": true},
            password: '',
            domainNames: []
        }),
    }
</script>
