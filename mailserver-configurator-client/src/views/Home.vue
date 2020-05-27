<template>
  <div class="container">
    <v-row no-gutters>
      <v-col cols="12" sm="6">
        <v-card style="margin-right: 10px; height: 100%;">
          <v-card-title>Go Mail Admin</v-card-title>
          <v-card-text>
            Mail Server Admin GUI<br><br>
            <b>Fast Access</b><br>
            <v-list-item-group v-model="item" color="primary">
              <v-list-item to="/account/new">
                <v-list-item-content>
                  <v-list-item-title>New Account</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-list-item to="/alias/new">
                <v-list-item-content>
                  <v-list-item-title>New Aliases</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-list-item href="https://github.com/kekskurse/go-mail-admin/issues" target="_blank">
                <v-list-item-content>
                  <v-list-item-title>Report an Issue</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list-item-group>

          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="2">
        <v-card style="margin-right: 10px;">
          <v-card-title>Domains</v-card-title>
          <v-card-text>
            <DomainChart></DomainChart>
            </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="2">
        <v-card style="margin-right: 10px;">
          <v-card-title>Aliases</v-card-title>
          <v-card-text>
            <AliasChart></AliasChart>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="2">
        <v-card style="margin-right: 10px;">
          <v-card-title>Accounts</v-card-title>
          <v-card-text>
            <AccountChart></AccountChart>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-card style="margin-top: 10px;">
      <v-card-title>Icons</v-card-title>
      <v-card-text>
        <v-simple-table>
          <template v-slot:default>
            <thead>
            <tr>
              <th>Icon</th>
              <th>Description</th>
            </tr>
            </thead>
            <tbody>
            <tr>
              <td><v-icon>mdi-plus-circle-outline</v-icon></td>
              <td>Add a new entry to the current list (e.g. on the Account Page)</td>
            </tr>
            <tr>
              <td><v-icon>mdi-circle-edit-outline</v-icon></td>
              <td>Edit the current selected entry</td>
            </tr>
            <tr>
              <td><v-icon>mdi-close-circle-outline</v-icon></td>
              <td>Remove the current selected entry</td>
            </tr>
            <tr>
              <td><v-icon>mdi-dns</v-icon></td>
              <td>Domain List</td>
            </tr>
            <tr>
              <td><v-icon>mdi-forwardburger</v-icon></td>
              <td>Aliases</td>
            </tr>
            <tr>
              <td><v-icon>mdi-account</v-icon></td>
              <td>Accounts</td>
            </tr>
            <tr>
              <td><v-icon>mdi-security</v-icon></td>
              <td>TLS Policys</td>
            </tr>
            <tr>
              <td><v-icon>mdi-view-dashboard-variant</v-icon></td>
              <td>Dashboard</td>
            </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
// @ is an alias to /src
import DomainChart from "../components/DomainChart";
import AliasChart from "../components/AliasChart";
import AccountChart from "../components/AccountChart";
export default {
  name: 'Home',
  components: {
    DomainChart,
    AliasChart,
    AccountChart
  }
}
</script>
