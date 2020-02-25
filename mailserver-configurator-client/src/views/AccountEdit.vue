<template>
    <div>
        <v-container>
            <v-card>
                <v-card-title>Account {{account.username}}@{{account.domain}} </v-card-title>
                <v-card-text>
                    <span v-if="!account.id">
                        <v-text-field v-model="account.username" label="Username" placeholder="Username"></v-text-field>
                        <label>Domain</label>
                         <v-select2 v-model="account.domain" data-app="true" :options="domainNames" label="Destination-Domain"></v-select2>
                        <v-text-field v-model="account.password" label="Password" type="password" placeholder="Password"></v-text-field>
                    </span>
                    <v-text-field v-model="account.quota" label="Quota" placeholder="Quota"></v-text-field>
                    <v-checkbox v-model="account.enabled" label="Enabled"></v-checkbox>
                    <v-checkbox v-model="account.sendonly" label="Send Only"></v-checkbox>
                    <v-btn @click="saveAlias">Save Account</v-btn>
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
                    })
                } else {
                    Client.createAccount(this.account).then(() => {
                        this.getAccounts();
                        this.$swal("Account saved");
                    })
                }

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
            domainNames: []
        }),
    }
</script>
