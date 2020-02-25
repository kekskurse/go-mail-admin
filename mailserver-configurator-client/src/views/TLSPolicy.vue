<template>
    <div>
        <v-container>
            <v-card style="padding-bottom: 10px;">
                <v-card-title>
                    TLSPolicy
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
                    <v-btn to="/tls/new" icon><v-icon>mdi-plus-circle-outline</v-icon></v-btn>
                    <v-btn @click="removePolicy()" v-if="selected[0]" icon><v-icon>mdi-close-circle-outline</v-icon></v-btn>
                    <v-btn @click="editPolicy()" v-if="selected[0]" icon><v-icon>mdi-circle-edit-outline</v-icon></v-btn>
                </span>
                <v-data-table
                        :headers="headers"
                        :items="tlspolicys"
                        :search="search"
                        :single-select=true
                        v-model="selected"
                        show-select
                ></v-data-table>
                <span style="background-color:#BBDEFB; margin-left: 10px; border-radius: 5px; padding-top: 10px;padding-bottom:8px;">
                    <v-btn to="/tls/new" icon><v-icon>mdi-plus-circle-outline</v-icon></v-btn>
                    <v-btn @click="removePolicy()" v-if="selected[0]" icon><v-icon>mdi-close-circle-outline</v-icon></v-btn>
                    <v-btn @click="editPolicy()" v-if="selected[0]" icon><v-icon>mdi-circle-edit-outline</v-icon></v-btn>
                </span>

            </v-card>


        </v-container>


    </div>


</template>

<script>

    import Client from "../service/Client";

    export default {
        name: 'Domain',
        methods: {
            getPolicys: function () {
                Client.getTLSPolicys().then((res) => {
                    this.tlspolicys = res.data;
                });
            },
            editPolicy: function () {
                this.$router.push("/tls/"+this.selected[0].id)
            },
            removePolicy: function () {
                this.$swal({
                    'title': 'Delete TLS Policy',
                    'icon': "warning",
                    showCancelButton: true,
                }).then((res) => {
                    if(res.value) {
                        Client.deleteTLSPolicy(this.selected[0].id).then(() => {
                            this.getPolicys();
                        });
                    }
                });

            }
        },
        mounted: function() {
            this.getPolicys();

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
                    text: 'Domain',
                    sortable: true,
                    value: 'domain'
                },
                {
                    text: 'Policy',
                    sortable: true,
                    value: 'policy'
                },
                {
                    text: 'Params',
                    sortable: true,
                    value: 'params'
                }
            ],
            'search': '',
            'tlspolicys': [],
            'selected': [],

        }),
    }
</script>
