<template>
    <div>
        <v-container>
            <v-card>
                <v-card-title>Alias {{alias.source_username}}@{{alias.source_domain}} to {{alias.destination_username}}@{{alias.destination_domain}}</v-card-title>
                <v-card-text>
                    <v-text-field
                        v-model="alias.source_username"
                        placeholder="Source Username"
                        label="Source Username"
                        ></v-text-field>
                    <v-select
                            :items="domainNames"
                            label="Source Domain"
                            v-model="alias.source_domain"
                    ></v-select>
                    <v-text-field
                            v-model="alias.destination_username"
                            placeholder="Source Username"
                            label="Source Username"
                    ></v-text-field>
                    <v-text-field
                            v-model="alias.destination_domain"
                            placeholder="Destination Domain"
                            label="Destination Domain"
                    ></v-text-field>
                    <v-checkbox v-model="alias.enabled" label="Enabled"></v-checkbox>
                    <span style="background-color:#BBDEFB; margin-left: 10px; border-radius: 5px; padding-top: 10px;padding-bottom:8px;">
                        <v-btn @click="saveAlias" icon><v-icon>mdi-content-save</v-icon></v-btn>
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
            getAliases: function () {
                Client.getAlias().then((res) => {
                    for(var i = 0; i < res.data.length; i++) {
                        if(res.data[i].id == this.$route.params.id) {
                            this.alias = res.data[i]
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
                if(this.alias.id) {
                    Client.saveAlias(this.alias).then(() => {
                        this.getAliases();
                        this.$swal("Alias saved");
                        this.$router.push("/alias")
                    })
                } else {
                    Client.createAlias(this.alias).then(() => {
                        this.getAliases();
                        this.$swal("Alias saved");
                    })
                }

            }
        },

        mounted: function() {
            this.getDomains();
            this.getAliases();

        },
        components: {

        },
        data: () => ({
            alias: {"enabled": true},
            domainNames: [],
            sample: ["abc", "asd", "sdf"]
        }),
    }
</script>
