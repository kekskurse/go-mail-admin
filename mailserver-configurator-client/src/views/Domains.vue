<template>
    <div>
        <v-container>
            <v-card>
                <v-card-title>
                    Add Domain

                </v-card-title>
                <v-card-text>
                    <v-text-field
                            v-model="newDomain"
                            placeholder="example.com"
                    ></v-text-field>
                    <v-btn @click="addDomain()">Add Domain</v-btn><br><br>

                </v-card-text>
            </v-card>


        </v-container>


        <v-container>
            <v-card>
                <v-card-title>
                    Domains
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
                        :items="domains"
                        :search="search"
                        :single-select=true
                        v-model="selected"
                        show-select
                >
                 <template v-slot:item.details.MXRecordCheck="{ item }">
                        <v-chip color="green" v-if="item.enabled">Yes</v-chip>
                        <v-chip color="red" v-if="!item.enabled">No</v-chip>
                    </template>
                </v-data-table>
                <v-btn @click="removeDomain()" v-if="selected[0]">Remove selected Domain</v-btn><br><br>
            </v-card>


        </v-container>


    </div>


</template>

<script>

    import Client from "../service/Client";

    export default {
        name: 'Domain',
        methods: {
            getDomains: function () {
                Client.getDomains().then((res) => {
                   this.domains = res.data;
                });
            },
            removeDomain: function() {
                this.$swal({
                    'title': 'Delete Domain',
                    'text': "Do you want to delete the Domain "+this.selected[0].domain+"?",
                    'icon': "warning",
                    showCancelButton: true,
                }).then((res) => {
                    if(res.value) {
                        Client.removeDomain(this.selected[0].domain).then(() => {
                            this.getDomains();
                        })
                    }
                });
            },
            addDomain: function () {
                Client.addDomain(this.newDomain).then(() => {
                    this.getDomains();
                    this.newDomain = "";

                }).catch(() => {
                    alert("Something go wrong");
                });
            }
        },
        mounted: function() {
          this.getDomains();

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
                }
            ],
            'search': '',
            'domains': [],
            'selected': [],
            'newDomain': ''

        }),
    }
</script>
