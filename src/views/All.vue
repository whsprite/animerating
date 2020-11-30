<template>
<v-main>
    <v-container>
        <v-row class="mt-9">
            <div class="welcome mx-auto text-caption text-h5 text-lg-h3 text-sm-h4 font-weight-bold mb-5">Done your first comment at banping</div>
        </v-row>
        <v-row cols="auto">
            <v-col :sm="4" :lg="2" :md="3" cols="6" v-for="item in animes" :key="item.Id">
                <v-hover>
                    <template v-slot:default="{ hover }">
                        <v-card color="transparer" elevation="12" :loading="false">
                            <v-img :src="item.Url" :spect-ratio="3/4" max-height="340px" lazy-src cover class="white--text align-end"></v-img>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title class="mx-auto text-lg-h6 text-sm-body-2 text-caption text-center">{{item.Name}}</v-list-item-title>
                                    <v-list-item-subtitle>
                                        <v-chip class="ma-2 text-caption" color="secondary" v-for="(cate,id) in item.Category.split('/')" :key="id">{{cate}}</v-chip>
                                    </v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>

                            <v-row class="mx-auto align-center">
                                <v-rating class="mx-auto" v-model="item.Rating" background-color="pink lighten-3" color="pink lighten-2" readonly dense :size="30"></v-rating>
                                <v-list-item>
                                    <v-list-item-content>
                                        <v-list-item-subtitle class="font-weight-bold">
                                            rate:{{item.Rating}}
                                            <v-spacer></v-spacer>
                                            reviewers:{{item.RateNum}}
                                        </v-list-item-subtitle>
                                    </v-list-item-content>
                                </v-list-item>
                            </v-row>
                            <v-fade-transition>
                                <v-overlay v-if="hover" absolute color="black lighten-3">
                                    <v-rating class="mx-auto" v-model="rate" background-color="black" color="yellow accent-4" :x-small="$vuetify.breakpoint.mobile ? true : false"></v-rating>
                                    <v-row>
                                        <v-btn class="mx-auto mt-5 black--text" @click="doRate(item.Id)" color="white">评分</v-btn>
                                    </v-row>
                                    <v-row>
                                        <v-btn class="mx-auto mt-5 black--text" @click="getCommet(item.Id)" color="white">评论</v-btn>
                                    </v-row>
                                </v-overlay>
                            </v-fade-transition>
                        </v-card>
                    </template>
                </v-hover>
            </v-col>
        </v-row>
        <v-snackbar v-model="snackbar" :timeout="2000">
            {{ text }}
            <template v-slot:action="{ attrs }">
                <v-btn color="pink lighten-1" text v-bind="attrs" @click="snackbar = false">Close</v-btn>
            </template>
        </v-snackbar>
    </v-container>
    <v-dialog v-model="dialog" max-width="1000px">
        <v-card>
            <v-toolbar flat color="pink" dark>
                <v-toolbar-title>评论</v-toolbar-title>
            </v-toolbar>

            <v-card-text>
                <v-textarea class="mt-5" filled v-model="comment" value></v-textarea>
            </v-card-text>
            <v-rating class="mx-auto ml-12" v-model="rate" background-color="pink lighten-3" color="yellow accent-4" dense :size="30"></v-rating>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="pink lighten-4" depressed @click="doComment">Post</v-btn>
            </v-card-actions>
            <v-divider></v-divider>

            <template v-for="item in comments">
                <v-textarea solo readonly class="pr-10 pl-10 pt-10" :key="item.ID" :value="item.CommentString" auto-grow></v-textarea>
                <v-row :key="item.Id">
                    <v-rating class="ml-12" dense v-model="item.Rate" background-color="black" color="yellow accent-4" readonly :x-small="$vuetify.breakpoint.mobile ? true : false"></v-rating>
                </v-row>
            </template>
        </v-card>
    </v-dialog>
</v-main>
</template>

<script>
import {
    reactive,
    toRefs,
    onMounted
} from "@vue/composition-api";
import LoginVue from "./Login.vue";
import axios from "axios";
axios.defaults.baseURL = "/api/v1/";
export default {
    setup() {
        const state = reactive({
            animes: [],
            rate: 5,
            snackbar: false,
            text: "",
            dialog: false,
            comments: [],
            comment: "",
            id: "",
        });
        const getAnime = async () => {
            const {
                data: res
            } = await axios.get("/anime_rate");
            state.animes = res.data;
        }
        const doRate = async id => {
            const {
                data: res
            } = await axios.put("/do_rate", JSON.stringify({
                id: id,
                rate: state.rate
            }))
            if (res.code !== 1000) {
                state.snackbar = true;
                state.text = "评分失败"
                return
            }
            state.snackbar = true;
            state.text = "评分成功";
            getAnime();
            return
        }
        const getCommet = async id => {
            state.id = id;
            const {
                data: res
            } = await axios.get(`/comment/${id}`)
            state.comments = res.data;
            state.dialog = true;
            console.log(res)
            return
        }
        const doComment = async () => {
            const {
                data: res
            } = await axios.post(`/comment/${state.id}`, JSON.stringify({

                comment: state.comment,
                rate: state.rate,

            }))
            if (res.code !== 1000) {
                state.snackbar = true;
                state.text = "评论失败"
                return
            }
            state.snackbar = true;
            state.text = "评论成功";
            state.comment = "";
            getCommet(state.id);
            return
        }
        onMounted(() => {
            getAnime();
        });
        return {
            ...toRefs(state),
            getAnime,
            doRate,
            getCommet,
            doComment,
        };
    }
};
</script>

<style scoped>
.welcome {
    background: linear-gradient(to right,
            rgba(249, 152, 226, 1),
            rgba(42, 121, 215, 0.849645390070922));
    -webkit-background-clip: text;
    color: transparent;
}
</style>
