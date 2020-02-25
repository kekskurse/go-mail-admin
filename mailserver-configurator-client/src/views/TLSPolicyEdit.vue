<template>
  <div class="container">
    <v-card>
      <v-card-title>TLS Policy</v-card-title>
      <v-card-text>
        <v-text-field type="text" v-model="domain" label="Domain"></v-text-field>
        <label>Policy</label>
        <v-select2 v-model="policy" data-app="true" :options="policyOptions" label="Destination-Domain"></v-select2>
        <v-text-field type="text" v-model="params" label="Params"></v-text-field>
        <v-btn @click="savePolicy()">Save Policy</v-btn>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
// @ is an alias to /src
import Client from "../service/Client";

export default {
  name: 'Home',
  components: {
  },
  methods: {
    'savePolicy': function () {
      if(this.$route.params.id) {
        Client.saveTLSPolicy({"id": this.$route.params.id,"domain": this.domain, "policy": this.policy, "params": this.params}).then(() => {
          this.$swal("Policy saved");
          this.getPolicy();
        }).catch((res) => {
          console.error(res)
          alert("Oups, something go wrong")
        });
      } else {
        Client.createTLSPolicy({"domain": this.domain, "policy": this.policy, "params": this.params}).then(() => {
          this.$swal("Policy saved");
        }).catch((res) => {
          console.error(res)
          alert("Oups, something go wrong")
        });
      }

    },
    getPolicy: function () {
      Client.getTLSPolicys().then((res) => {
        for(var i = 0; i < res.data.length; i++) {
          if(res.data[i].id == this.$route.params.id) {
            this.domain = res.data[i].domain
            this.policy = res.data[i].policy
            this.params = res.data[i].params
          }
        }
      });
    }
  },
  mounted: function () {
    this.getPolicy();
  },
  data: () => ({
    domain: "",
    policy: "",
    params: "",
    policyOptions: ['none', 'may', 'encrypt', 'dane', 'dane-only', 'fingerprint', 'verify', 'secure']
  })
}
</script>
