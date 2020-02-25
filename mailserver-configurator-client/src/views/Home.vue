<template>
  <div class="container">
    <v-card>
      <v-card-title>Go Mail Admin</v-card-title>
      <v-card-text>Mail Server Admin GUI</v-card-text>
    </v-card>
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


export default {
  name: 'Home',
  components: {

  }
}
</script>
