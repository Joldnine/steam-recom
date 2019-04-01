<template>
  <div class="hello">
    <el-container>
      <el-header>Steam Games Recommender</el-header>
      <el-main>
        <el-breadcrumb  separator-class="el-icon-arrow-right">
          <el-breadcrumb-item :to="{ path: '/' }">Steam Games</el-breadcrumb-item>
          <el-breadcrumb-item>Recommender</el-breadcrumb-item>
        </el-breadcrumb>
        <el-card  class="c-card" v-loading="loading">
          <div slot="header" class="clearfix">
            <span>About Me</span>
          </div>
          <el-row :gutter="8">
            <el-col :xs="24" :sm="24" :md="20" :lg="20" :xl="20">
              <el-input v-model="user_id" placeholder="User ID Here"></el-input>
            </el-col>
            <el-col :xs="24" :sm="24" :md="4" :lg="4" :xl="4">
                <el-button type="primary" @click="onGo">Go</el-button>
            </el-col>
          </el-row>
          <el-card :hidden="favorateGames.length === 0">
            <div slot="header" class="clearfix">
              <span>My favorites</span>
            </div>
            <el-carousel :interval="4000" height="215px" type="card" >
              <el-carousel-item v-for="item in favorateGames" :key="item.ID" :label="item.Name">
                <img :src="getImg(item.ID)"/>
              </el-carousel-item>
            </el-carousel>
          </el-card>
        </el-card>
        <el-card v-loading="loading" :hidden="recomGames.length === 0" class="c-card">
          <div slot="header" class="clearfix">
            <span>Recommend for you</span>
          </div>
          <el-carousel :interval="4000" height="215px" type="card" >
            <el-carousel-item v-for="item in recomGames" :key="item.ID" :label="item.Name">
              <img :src="getImg(item.ID)"/>
            </el-carousel-item>
          </el-carousel>
        </el-card>
      </el-main>
    </el-container>
  </div>
</template>

<script>
const paramsToUrl = (params) => {
  var esc = encodeURIComponent
  return '?' + Object.keys(params)
    .map(k => `${esc(k)}=${esc(params[k])}`)
    .join('&')
}

function getFavorite(user_id) {
  let API_ROOT = "http://localhost:8081/"
  const params = {
    user_id
  }
  var query = API_ROOT + 'favorite' + paramsToUrl(params)
  return fetch(query)
}

function getRecommend(user_id) {
  let API_ROOT = "http://localhost:8081/"
  const params = {
    user_id
  }
  var query = API_ROOT + 'recommend' + paramsToUrl(params)
  return fetch(query)
}

export default {
  data() {
    return {
      loading: false,
      user_id: "",
      favorateGames: [
      ],
      recomGames: [
      ]
    }
  },
  methods: {
    onGo() {
      this.loading = true
      getFavorite(this.user_id).then((response) => {
        if (!response.ok) {
          response.json().then((error) => {
            throw Error(error.Msg);
          }).catch(error => {
            this.$message.error(error.message)
            this.loading = false
          })
        } else {
          response.json().then((response) => {
            this.favorateGames = response
            this.loading = false
          })
        }
      })
      getRecommend(this.user_id).then((response) => {
        if (!response.ok) {
          response.json().then((error) => {
            throw Error(error.Msg);
          }).catch(error => {
            this.$message.error(error.message)
            this.loading = false
          })
        } else {
          response.json().then((response) => {
            this.recomGames = response
            this.loading = false
          })
        }
      })
    },
    getUrl(game_id) {
      return "https://store.steampowered.com/app/" + game_id
    },
    getImg(game_id) {
      return "https://steamcdn-a.akamaihd.net/steam/apps/" + game_id + "/header.jpg"
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.el-header, .el-footer {
  background-color: #183545;
  color: white;
  text-align: left;
  line-height: 60px;
}
.el-main {
  /* background-color: #E9EEF3; */
  color: #333;
  text-align: left;
  /* height: 90vh; */
  /* line-height: 160px; */
  /* margin: 0 100px; */
  margin: auto;
  font-family: tahoma,arial,'PingFang SC','Hiragino Sans GB','WenQuanYi Micro Hei','Microsoft YaHei',sans-serif;
}
.el-card {
  margin-top: 16px;
}
.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}

.c-card {
  width: 960px
}
</style>
