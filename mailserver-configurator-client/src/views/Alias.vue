<template>
    <div>
        <v-container>
            <v-card>
                <v-card-title>
                    Aliases
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
                        :items="aliases"
                        :search="search"
                        :single-select=true
                        v-model="selected"
                        show-select
                ></v-data-table>
                <v-btn @click="removeAlias()" v-if="selected[0]">Remove selected Alias</v-btn>
                <v-btn @click="editAlias()" v-if="selected[0]">Edit Alias</v-btn>
                <v-btn to="/alias/new">New Alias</v-btn><br><br>
            </v-card>


        </v-container>


    </div>


</template>

<script>

    import Client from "../service/Client";

    export default {
        name: 'Domain',
        methods: {
            getAliases: function () {
                Client.getAlias().then((res) => {
                    this.aliases = res.data;
                });
            },
            editAlias: function () {
                this.$router.push("/alias/"+this.selected[0].id)
            },
            removeAlias: function () {
                this.$swal({
                    'title': 'Delete Alias',
                    'icon': "warning",
                    showCancelButton: true,
                }).then((res) => {
                    if(res.value) {
                        Client.removeAlias(this.selected[0].id).then(() => {
                            this.getAliases();
                        })
                    }
                });
            }

        },
        mounted: function() {
            this.getAliases();

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
                    value: 'source_username'
                },
                {
                    text: 'Source Domain',
                    sortable: true,
                    value: 'source_domain'
                },
                {
                    text: 'Destination Username',
                    sortable: true,
                    value: 'destination_username'
                },
                {
                    text: 'Destination Domain',
                    sortable: true,
                    value: 'destination_domain'
                },
                {
                    text: 'Enabled',
                    sortable: true,
                    value: 'enabled'
                }
            ],
            'search': '',
            'aliases': [],
            'selected': [],

        }),
    }
</script>
