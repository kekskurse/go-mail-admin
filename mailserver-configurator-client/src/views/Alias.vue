<template>
    <div>
        <v-container>
            <v-card style="padding-bottom: 10px;">
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
                <span style="background-color:#BBDEFB; margin-left: 10px; border-radius: 5px; padding-top: 10px;padding-bottom:8px;">
                    <v-btn to="/alias/new" icon><v-icon>mdi-plus-circle-outline</v-icon></v-btn>
                    <v-btn @click="removeAlias()" v-if="selected[0]" icon><v-icon>mdi-close-circle-outline</v-icon></v-btn>
                    <v-btn @click="editAlias()" v-if="selected[0]" icon><v-icon>mdi-circle-edit-outline</v-icon></v-btn>
                </span>

                <v-data-table
                        :headers="headers"
                        :items="aliases"
                        :search="search"
                        :single-select=true
                        v-model="selected"
                        show-select
                ></v-data-table>
                <span style="background-color:#BBDEFB; margin-left: 10px; border-radius: 5px; padding-top: 10px;padding-bottom:8px;">
                    <v-btn to="/alias/new" icon><v-icon>mdi-plus-circle-outline</v-icon></v-btn>
                    <v-btn @click="removeAlias()" v-if="selected[0]" icon><v-icon>mdi-close-circle-outline</v-icon></v-btn>
                    <v-btn @click="editAlias()" v-if="selected[0]" icon><v-icon>mdi-circle-edit-outline</v-icon></v-btn>
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
