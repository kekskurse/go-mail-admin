<template>
    <div class="container">
      Logout...
    </div>
</template>

<script>
    // @ is an alias to /src
    import Client from "../service/Client";
    export default {
        name: 'Home',
        components: {
        },
        mounted() {
            console.log("Logout");
            Client.logout();
        }

    }
</script>
