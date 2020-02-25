<script>
    import { Doughnut } from "vue-chartjs";
    import Client from "../service/Client";

    export default {
        extends: Doughnut,
        mounted() {
            Client.getDomains().then((res) => {
                this.renderChart(
                    {
                        labels: ["Domains"],
                        datasets: [
                            {
                                backgroundColor: ["#4CAF50"],
                                data: [res.data.length]
                            }
                        ]
                    },
                    { responsive: true, maintainAspectRatio: true, legend: {display: false} }
                );
            })

        }
    };
</script>