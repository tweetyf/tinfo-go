<template>
    <div class="w3-container w3-center w3-card-4 w3-margin">
        <h4>Search Country Code By IP address </h4>
        <div class="w3-bar w3-margin ">
            <form @submit="onSubmit">
                <input class="w3-input w3-bar-item  w3-border" id="ipaddr" type="text" placeholder="IP address"
                    v-model="ipAddr">
                <button class="w3-btn w3-blue  w3-bar-item" @click="onSubmit">Search</button>
            </form>
        </div>

        <div id="countryCode" class="w3-panel w3-pale-green w3-border">
            <p>IP: {{ ipLoc.ip }} Country Code: {{ ipLoc.country }}</p>
        </div>
    </div>
</template>

<script>
export default {
    Name: "IPlocation",
    data() {
        return {
            ipAddr: "",
            ipLoc: JSON.parse('{"ip": "","country": "", "status": ""}'),
        }
    },
    methods: {
        async fetchData() {
            const url = "http://0.0.0.0:8080/iptool/country?ip=" + this.ipAddr;
            const ab1 = await fetch(url)
            const res = await (ab1).json();
            console.log(ab1)
            this.ipLoc = res;
            this.ipAddr ="";
        },
        onSubmit(e) {
            e.preventDefault();
            this.fetchData();
        },
    }
}
</script>