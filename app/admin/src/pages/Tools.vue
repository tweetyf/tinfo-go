<script>
import IPlocationVue from "../components/IPlocation.vue";
import RandomUserVue from "../components/RandomUser.vue";

export default{
    Name:"Tools",
    components: {
        IPlocationVue,
        RandomUserVue
    },
}
</script>
<template>
<div class="w3-container">
    <RandomUserVue></RandomUserVue>
    <IPlocationVue></IPlocationVue>

</div>
</template>
