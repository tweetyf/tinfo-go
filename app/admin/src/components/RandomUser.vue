<script>
export default {
    Name: "Random User",
    data() {
        return {
            avatar: '',
            firstName: "",
            title: "",
            lastName: "",
            email: "",
        }
    },
    created() {
        this.fetchData();
    },
    methods: {
        async fetchData() {
            const url = "https://randomuser.me/api/";
            const res = await (await fetch(url)).json();
            console.log(res)
            this.email = res.results[0].email
            this.avatar = res.results[0].picture.thumbnail
            this.title = res.results[0].name.title
            this.firstName = res.results[0].name.first
            this.lastName = res.results[0].name.last
        },
    },
}
</script>

<template>
    <div class="w3-card-4 w3-center w3-margin w3-padding">
        <h4>Get A Random User</h4>
        <div class="w3-bar">
            <img :src="avatar" alt="Avatar" class="w3-border w3-bar-item">
            <div class="w3-bar-item">
                <h5>{{ title }}. {{ firstName }} {{ lastName }}</h5>
                <p>{{ email }}</p>
            </div>
        </div><br>
        <button class="w3-button w3-blue" @click="fetchData()">Refresh</button>
    </div>
</template>