<script>
    import { Doughnut } from "vue-chartjs";
    import Client from "../service/Client";

    export default {
        extends: Doughnut,
        mounted() {
            Client.getAccounts().then((res) => {
                let enabled_sendonly = 0;
                let enabled_receive= 0;
                let disabled_sendonly = 0;
                let disabled_receive = 0

                for(var i = 0; i < res.data.length; i++) {
                    if(res.data[i].enabled) {
                        if(res.data[i].sendonly) {
                            enabled_sendonly++;
                        } else {
                            enabled_receive++;
                        }
                    } else {
                        if(res.data[i].sendonly) {
                            disabled_sendonly++;
                        } else {
                            disabled_receive++;
                        }
                    }

                }

                this.renderChart(
                    {
                        labels: ["Enabled", "Enabled (Sendonly)", "Disabled", "Disabled (Sendonly)"],
                        datasets: [
                            {
                                backgroundColor: ["#4CAF50", "#8BC34A", "#F44336", "#EF5350"],
                                data: [enabled_receive, enabled_sendonly, disabled_receive, disabled_sendonly]
                            }
                        ]
                    },
                    { responsive: true, maintainAspectRatio: true, legend: {display: false} }
                );
            })
        }
    };
</script>