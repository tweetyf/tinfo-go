<script>
export default {
    Name: "Home",
    data() {
        return {
            sResults: '',
            qSearch: '',
            qKW: '',
        }
    },
    created() {
        //this.fetchData("golang");
    },
    methods: {
        async fetchData(kw) {
            if (!kw) {
                alert('Please put content for searching')
                return
            }
            const url = "https://api.duckduckgo.com/?q=" + kw + "&format=json&pretty=1";
            const res = await (await fetch(url)).json();
            console.log(res);
            this.sResults = res;
        },

        onSubmit(e) {
            e.preventDefault();
            console.log("submitting..." + this.qSearch);
            this.fetchData(this.qSearch);
            this.qKW = this.qSearch;
            this.qSearch = "";

        }
    },
}
</script>

<template>
    <div class="w3-container">
        <div class="w3-center w3-margin">
            <form class="w3-container w3-bar w3-padding" @submit="onSubmit">
                <input class="w3-input w3-border w3-light-grey w3-bar-item" v-model="qSearch" type="text"
                    placeholder="Search The Web">
                <button class="w3-btn w3-blue-grey w3-bar-item" @click="onSubmit">Search</button>
            </form>
        </div>

        <div class="w3-container w3-bar">
            examples:<br>
            <a class="w3-margin-left" @click="fetchData('Valley Forge')">Valley Forge
                National Historical
                Park</a>
            <button class="w3-btn w3-blue w3-margin-left" @click="fetchData('The Simpsons characters')">The
                Simpsons
                characters</button>
            <button class="w3-btn w3-blue w3-margin-left" @click="fetchData('Apple')">Apple</button>
            <button class="w3-btn w3-blue w3-margin-left" @click="fetchData('IBM')">IBM</button>

        </div>
        <div class="w3-container search-result">
            <p>Quick Answer for {{ qKW }}:</p>
            <p>{{ sResults.Abstract }}</p><br>
            <p>{{ sResults.AbstractSource }}: {{ sResults.AbstractURL }}</p>
        </div>
    </div>

</template>