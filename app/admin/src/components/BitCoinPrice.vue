<script>
export default {
    Name: "Bitcoin Price Index (BPI) in real-time",
    data() {
        return {
            btcPrice: JSON.parse('{"time":{"updated":"Jun 22, 2022 08:13:00 UTC","updatedISO":"2022-06-22T08:13:00+00:00","updateduk":"Jun 22, 2022 at 09:13 BST"},"disclaimer":"This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org","chartName":"Bitcoin","bpi":{"USD":{"code":"USD","symbol":"&#36;","rate":"20,188.1738","description":"United States Dollar","rate_float":20188.1738},"GBP":{"code":"GBP","symbol":"&pound;","rate":"16,391.1619","description":"British Pound Sterling","rate_float":16391.1619},"EUR":{"code":"EUR","symbol":"&euro;","rate":"19,190.2724","description":"Euro","rate_float":19190.2724}}}'),
        }
    },
    created() {
        this.fetchData();
        setInterval(this.fetchData, 30000);
    },
    methods: {
        async fetchData() {
            const url = "https://api.coindesk.com/v1/bpi/currentprice.json";
            const res = await (await fetch(url)).json();
            console.log(res)
            this.btcPrice = res;
        },
    },
}
</script>

<template>
    <div class="w3-card-4 w3-center w3-margin w3-padding">
        <h4>Bitcoin Price Index (BPI) in real-time -auto refresh every 30 seconds</h4>
        <p>{{ btcPrice.chartName }} : {{ btcPrice.time.updated }} </p>
        <p>{{ btcPrice.bpi.USD.code }}: {{ btcPrice.bpi.USD.rate }}</p>
        <p>{{ btcPrice.bpi.EUR.code }}: {{ btcPrice.bpi.EUR.rate }}</p>
        <p>{{ btcPrice.bpi.GBP.code }}: {{ btcPrice.bpi.GBP.rate }}</p>
    </div>
</template>