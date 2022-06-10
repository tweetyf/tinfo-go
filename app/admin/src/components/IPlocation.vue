<template>
    <div class="w3-container w3-center w3-card-4 w3-margin">
        <h4>Search Country Code By IP address </h4>
        <div class="w3-bar w3-margin ">
            <input class="w3-input w3-bar-item  w3-border" id="ipaddr" type="text" placeholder="IP address">
            <button class="w3-btn w3-blue  w3-bar-item" onclick="">Search</button>
        </div>

        <div id="countryCode" class="w3-panel w3-pale-green w3-border">
            <p> Country code: </p>
        </div>
    </div>
</template>
<script>
export default{
    Name: "IPlocation",
}
</script>