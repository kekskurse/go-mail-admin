<template>
    <div>
        <v-container>
            <v-card style="padding-bottom: 10px;">
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
                <span style="background-color:#BBDEFB; margin-left: 10px; border-radius: 5px; padding-top: 10px;padding-bottom:8px;">
                    <v-btn to="/account/new" icon><v-icon>mdi-plus-circle-outline</v-icon></v-btn>
                    <v-btn @click="deleteAccount()" v-if="selected[0]" icon><v-icon>mdi-close-circle-outline</v-icon></v-btn>
                    <v-btn @click="editAlias()" v-if="selected[0]" icon><v-icon>mdi-circle-edit-outline</v-icon></v-btn>
                </span>
                <v-data-table
                        :headers="headers"
                        :items="accounts"
                        :search="search"
                        :single-select=true
                        v-model="selected"
                        show-select
                ></v-data-table>
                <span style="background-color:#BBDEFB; margin-left: 10px; border-radius: 5px; padding-top: 10px;padding-bottom:8px;">
                    <v-btn to="/account/new" icon><v-icon>mdi-plus-circle-outline</v-icon></v-btn>
                    <v-btn @click="deleteAccount()" v-if="selected[0]" icon><v-icon>mdi-close-circle-outline</v-icon></v-btn>
                    <v-btn @click="editAlias()" v-if="selected[0]" icon><v-icon>mdi-circle-edit-outline</v-icon></v-btn>
                </span>
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
