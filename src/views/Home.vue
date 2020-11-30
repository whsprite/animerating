<template>
<v-main>
    <v-row>
        <v-col>
            <v-app-bar absolute color="pink lighten-3" dark>
                <v-tabs slider-color="blue lighten-3" color="purple lighten-2">
                    <v-tab class="font-weight-bold text-body-1" to="all">一览</v-tab>
                    <v-tab class="font-weight-bold text-body-1" to="ranking">ランキング</v-tab>
                </v-tabs>
            </v-app-bar>
        </v-col>
    </v-row>
    <router-view></router-view>
    <v-footer dark padless>
        <v-card class="pink lighten-3 white--text text-center" min-width="100%">
            <v-card-text class="white--text">banping -- a anime rating website</v-card-text>

            <v-card-text>Created by W.h & powerd by Vue & Go</v-card-text>
        </v-card>
    </v-footer>
</v-main>
</template>

<script>

</script>

<style scoped>

</style>
