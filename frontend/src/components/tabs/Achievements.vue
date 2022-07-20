<template>
  <v-container>
    <v-row no-gutters justify="center">
      <v-col
        v-if="game"
        cols="auto"
        class="title font-weight-light text-center"
      >Achievement List</v-col>
      <v-col v-else>
        No game loaded :(
      </v-col>
    </v-row>
    <v-row
      no-gutters 
      justify="center"
    >
    </v-row>
    <v-list
      style="max-height: calc(100vh - 105px);"
      class="overflow-y-auto"
      dense
    >
      <v-list-item
      v-for="i in list"
      :key="i"
      class="px-1"
      :style="earnedStyle(i)"
      >
          <v-list-item-icon
          style="margin: auto 10px auto auto;"
          >
            <!--v-icon v-if="game.achievements[i][gameMode]" color="yellow darken-2">mdi-trophy</v-icon>
            <v-icon v-else color="grey darken-2">mdi-trophy</v-icon-->
            <v-icon v-if="game.achievements[i].Points < 10" color="brown">mdi-trophy-award</v-icon>
            <v-icon v-else-if="game.achievements[i].Points < 25" color="grey">mdi-trophy-award</v-icon>
            <v-icon v-else-if="game.achievements[i].Points < 50" color="yellow">mdi-trophy-award</v-icon>
            <v-icon v-else-if="game.achievements[i].Points < 100" color="cyan">mdi-compass-rose</v-icon>
            <v-icon v-else-if="game.achievements[i].Points >= 100" color="red">mdi-compass-rose</v-icon>
            <v-icon v-else></v-icon>
          </v-list-item-icon>
          <v-list-item-avatar
              tile
              size="45"
              style="margin: 0px 15px 0px 0px;"
          >
              <img :src="`http://i.retroachievements.org.s3.amazonaws.com/Badge/${badgeName(i)}.png`">
          </v-list-item-avatar>
          <v-list-item-content>
              <v-list-item-title>
                <label>{{ game.achievements[i].Title }}</label>
              </v-list-item-title>
              <v-list-item-subtitle>
                <v-tooltip top>
                  <template v-slot:activator="{ on, attrs }">
                    <label
                      v-bind="attrs"
                      v-on="on"
                    >
                      {{ game.achievements[i].Description }}
                    </label>
                  </template>
                  <span>{{ game.achievements[i].Description }}</span>
                </v-tooltip>
              </v-list-item-subtitle>
          </v-list-item-content>
          <v-tooltip top>
            <template v-slot:activator="{ on, attrs }">
              <div
                style="width: 55px; text-align: center; font-size: 0.550rem; font-weight: 500; text-transform: uppercase;"
                v-bind="attrs"
                v-on="on"
              >
                <label v-if="getAchievedPercent(i) < 5" style="color: darkorange;">Ultra Rare</label>
                <label v-else-if="getAchievedPercent(i) < 15" style="color: darkorchid;">Very Rare</label>
                <label v-else-if="getAchievedPercent(i) < 50" style="color: darkcyan;">Rare</label>
                <label v-else style="color: green;">Common</label>
              </div>
            </template>
            <span>Won by <strong>{{ getAchievedPercent(i) }}</strong>% of possible players</span>
          </v-tooltip>
      </v-list-item>
    </v-list>
  </v-container>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  computed: {
    ...mapState('game', ['game', 'gameMode']),

    list: {
      get() {
        if (this.game === null) return [];

        return this.game.defaultOrder;
      },
      set() {},
    },
    mode: {
      get() {
        return this.gameMode === 'DateEarnedHardcore';
      },
      set(val) {
        this.setGameMode(val ? 'DateEarnedHardcore' : 'DateEarned');
      },
    },
  },

  methods: {
    ...mapActions('game', ['setGameOrder', 'setGameMode']),

    earnedStyle(i) {
      if(this.game.achievements[i][this.gameMode]) {
        return {
            'opacity' : 1,
        }
      } else {
        return {
            'opacity' : 0.5,
        }
      }
    },

    getAchievedPercent(i) {
      if (this.game === null) return 0;
      var currentMode = (this.gameMode === 'DateEarnedHardcore') ? this.game.achievements[i].NumAwardedHardcore : this.game.achievements[i].NumAwarded;
      return Math.floor(Number(currentMode) / Number(this.game.softcore) * 10000) / 100;
    },

    badgeName(i) {
      if(this.game.achievements[i][this.gameMode]) {
        return this.game.achievements[i].BadgeName;
      } else {
        return this.game.achievements[i].BadgeName + "_lock";
      }
    }
  },
};
</script>