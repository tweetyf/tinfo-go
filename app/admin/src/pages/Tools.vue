<script>
import IPlocationVue from "../components/IPlocation.vue";
import RandomUserVue from "../components/RandomUser.vue";
import BitCoinPriceVue from "../components/BitCoinPrice.vue";

export default{
    Name:"Tools",
    components: {
        IPlocationVue,
        RandomUserVue,
        BitCoinPriceVue
    },
}
</script>
<template>
<div class="w3-container">
    <BitCoinPriceVue></BitCoinPriceVue>
    <RandomUserVue></RandomUserVue>
    <IPlocationVue></IPlocationVue>

</div>
</template>
