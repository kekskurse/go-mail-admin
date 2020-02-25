<template>
    <div>
        <v-container>
            <v-card>
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
                <v-data-table
                        :headers="headers"
                        :items="tlspolicys"
                        :search="search"
                        :single-select=true
                        v-model="selected"
                        show-select
                ></v-data-table>
                <v-btn @click="removePolicy()" v-if="selected[0]">Remove selected Policy</v-btn>
                <v-btn @click="editPolicy()" v-if="selected[0]">Edit TLS Policy</v-btn>
                <v-btn to="/tls/new">New TLS-Policy</v-btn><br><br>
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
                Client.deleteTLSPolicy(this.selected[0].id).then(() => {
                   this.getPolicys();
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
