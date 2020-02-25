<script>
    import { Doughnut } from "vue-chartjs";
    import Client from "../service/Client";

    export default {
        extends: Doughnut,
        mounted() {
            Client.getAlias().then((res) => {
                let enabled = 0;
                let disabled = 0;

                for(var i = 0; i < res.data.length; i++) {
                    if(res.data[i].enabled) {
                        enabled++;
                    } else {
                        disabled++;
                    }

                }

                this.renderChart(
                    {
                        labels: ["Einabled", "Disabled"],
                        datasets: [
                            {
                                backgroundColor: ["#4CAF50", "#F44336"],
                                data: [enabled, disabled]
                            }
                        ]
                    },
                    { responsive: true, maintainAspectRatio: true, legend: {display: false} }
                );
            })
        }
    };
</script>