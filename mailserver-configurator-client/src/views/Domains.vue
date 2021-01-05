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

                <template v-slot:item.mx="{ item }">
                    <span v-if="item.detail.RecordChecked">
                        <v-chip color="green" v-if="item.detail.MXRecordCheck" style="width:100px;">Yes</v-chip>
                        <v-chip color="red" v-if="!item.detail.MXRecordCheck" style="width:100px;">No</v-chip>
                    </span>
                    <v-chip v-if="!item.detail.RecordChecked" style="width:100px;">Unknown</v-chip>
                </template>

                <template v-slot:item.spf="{ item }">
                    <span v-if="item.detail.RecordChecked">
                        <v-chip color="green" v-if="item.detail.SPFRecordCheck" style="width:100px;">Yes</v-chip>
                        <v-chip color="red" v-if="!item.detail.SPFRecordCheck" style="width:100px;">No</v-chip>
                    </span>
                    <v-chip v-if="!item.detail.RecordChecked" style="width:100px;">Unknown</v-chip>
                </template>

                <template v-slot:item.dmarc="{ item }">
                    <span v-if="item.detail.RecordChecked">
                        <v-chip color="green" v-if="item.detail.DMARCRecordCheck" style="width:100px;">Yes</v-chip>
                        <v-chip color="red" v-if="!item.detail.DMARCRecordCheck" style="width:100px;">No</v-chip>
                    </span>
                    <v-chip v-if="!item.detail.RecordChecked" style="width:100px;">Unknown</v-chip>
                </template>
                
                <template v-slot:item.dkmi="{ item }">
                    <span v-if="item.detail.RecordChecked">
                        <v-chip color="green" v-if="item.detail.DKIMCheck" style="width:100px;">Yes</v-chip>
                        <v-chip color="red" v-if="!item.detail.DKIMCheck" style="width:100px;">No</v-chip>
                    </span>
                    <v-chip v-if="!item.detail.RecordChecked" style="width:100px;">Unknown</v-chip>
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
            getDomainDetails: function () {
                for(var i = 0; i < this.domains.length; i++) {
                    Client.getDomainDetails(this.domains[i].domain).then((res) => {
                        for(var i = 0; i < this.domains.length; i++) {
                            if(this.domains[i].domain == res.data.domain_name) {
                                this.domains[i].detail = res.data;
                                console.log("Found Domain Details")
                            }
                        }
                    });
                }
            },
            getFetaturesToggle: function () {
                Client.featureToggles().then((res) => {
                    console.log(res.data); 
                    if(res.data.showDomainDetails) {
                        console.log("Show Domain Details");
                        this.headers.push({"text": "MX-Record", "sortable": false, "value": "mx"});
                        this.headers.push({"text": "SPF-Record", "sortable": false, "value": "spf"});
                        this.headers.push({"text": "DMARC-Record", "sortable": false, "value": "dmarc"});
                        this.headers.push({"text": "DKMI-Record", "sortable": false, "value": "dkmi"});
                    }
                });
            },
            getDomains: function () {
                Client.getDomains().then((res) => {
                   this.domains = res.data;
                   this.getDomainDetails();
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
            this.getFetaturesToggle();
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
                },
                
            ],
            'search': '',
            'domains': [],
            'selected': [],
            'newDomain': ''

        }),
    }
</script>
