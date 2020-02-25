<template>
    <div>
        <v-container>
            <v-card>
                <v-card-title>
                    Accounts
                    <v-spacer></v-spacer>
                    <v-text-field
                            v-model="search"
                            append-icon="mdi-magnify"
                            label="Search"
                            single-line
                            hide-details


                    ></v-text-field>
                </v-card-title>
                <v-data-table
                        :headers="headers"
                        :items="accounts"
                        :search="search"
                        :single-select=true
                        v-model="selected"
                        show-select
                ></v-data-table>
                <v-btn @click="deleteAccount()" v-if="selected[0]">Remove selected Account</v-btn>
                <v-btn @click="editAlias()" v-if="selected[0]">Edit Account</v-btn>
                <v-btn to="/account/new">New Account</v-btn><br><br>
            </v-card>


        </v-container>


    </div>


</template>

<script>

    import Client from "../service/Client";

    export default {
        name: 'Account',
        methods: {
            getAccounts: function () {
                Client.getAccounts().then((res) => {
                    this.accounts = res.data;
                });
            },
            editAlias: function () {
                this.$router.push("/account/"+this.selected[0].id)
            },
            deleteAccount: function () {
                this.$swal({
                    'title': 'Delete Account',
                    'icon': "warning",
                    showCancelButton: true,
                }).then((res) => {
                    if(res.value) {
                        Client.deleteAccount(this.selected[0].id).then(() => {
                            this.getAccounts();
                        })
                    }
                });
            }

        },
        mounted: function() {
            this.getAccounts();

        },
        components: {

        },
        data: () => ({
            'headers': [
                {
                    text: '#',
                    sortable: true,
                    value: 'id'
                },
                {
                    text: 'Source Username',
                    sortable: true,
                    value: 'username'
                },
                {
                    text: 'Source Domain',
                    sortable: true,
                    value: 'domain'
                },
                {
                    text: 'Quota (MB)',
                    sortable: true,
                    value: 'quota'
                },
                {
                    text: 'Send Only',
                    sortable: true,
                    value: 'sendonly'
                },
                {
                    text: 'Enabled',
                    sortable: true,
                    value: 'enabled'
                }
            ],
            'search': '',
            'accounts': [],
            'selected': [],

        }),
    }
</script>
