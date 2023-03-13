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
      style="max-height: calc(100vh - 130px);"
      class="overflow-y-auto"
      dense
    >
      <v-list-item
      v-for="i in list"
      :key="i"
      class="px-1 achievement"
      @mouseenter="mouseEnter"
      >
        <v-list-item-icon
        style="margin: auto 10px auto auto;"
        >
          <v-icon v-if="game.achievements[i].Points < 10" color="brown">mdi-trophy-award</v-icon>
          <v-icon v-else-if="game.achievements[i].Points < 25" color="grey">mdi-trophy-award</v-icon>
          <v-icon v-else-if="game.achievements[i].Points < 50" color="yellow">mdi-trophy-award</v-icon>
          <v-icon v-else-if="game.achievements[i].Points < 100" color="cyan">mdi-compass-rose</v-icon>
          <v-icon v-else-if="game.achievements[i].Points >= 100" color="red">mdi-compass-rose</v-icon>
          <v-icon v-else></v-icon>
        </v-list-item-icon>
        <div class="wrap">
          <div class="AchievementIconWrapper" v-bind:class="(getAchievedPercent(i) < 5) ? 'LegendaryAchievement' : ''"><div class="AchievementIconGlowContainerRoot"><div class="AchievementIconGlowContainer"><div class="AchievementIconGlow"></div></div></div>
              <div class="achievementIcon" style="width: 48px; height: 48px;">
                <img :src="`http://i.retroachievements.org.s3.amazonaws.com/Badge/${badgeName(i)}.png`">
              </div>
          </div>
        </div>
        <div class="tooltip">
          <v-card
            class="mx-auto"
            max-width="344"
            outlined
          >
            <v-list-item three-line>
              <div
                style="align-self: flex-start; margin: 16px 16px 16px 0px"
              >
              <div class="AchievementIconWrapper" v-bind:class="(getAchievedPercent(i) < 5) ? 'LegendaryAchievement' : ''"><div class="AchievementIconGlowContainerRoot"><div class="AchievementIconGlowContainer"><div class="AchievementIconGlow"></div></div></div>
                  <div class="achievementIcon" style="width: 64px; height: 64px;">
                    <img :src="`http://i.retroachievements.org.s3.amazonaws.com/Badge/${game.achievements[i].BadgeName}.png`">
                  </div>
              </div>
              </div>
              <v-list-item-content>
                <v-list-item-title class="text-subtitle-2 mb-1 mt-2">
                  {{ game.achievements[i].Title }}
                </v-list-item-title>
                <div class="text-caption mb-4">{{ game.achievements[i].Description }}</div>
                <div class="text-overline">
                  {{ game.achievements[i].Points }} ({{ game.achievements[i].TrueRatio }}) Points
                </div>
                <v-list-item-subtitle v-if="getAchievedDate(i) != null">Unlocked: {{ getAchievedDate(i) }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </div>
        <v-list-item-content :style="earnedStyle(i)">
          <a target="_blank" :href="`https://retroachievements.org/achievement/${game.achievements[i].ID}`" style="text-decoration: inherit; color: inherit; cursor: pointer">
              <v-list-item-title>
                {{ game.achievements[i].Title }}
              </v-list-item-title>
              <v-list-item-subtitle>
                {{ game.achievements[i].Description }}
              </v-list-item-subtitle>
          </a>
        </v-list-item-content>
        <v-tooltip top>
          <template v-slot:activator="{ on, attrs }">
            <div
              style="width: 55px; text-align: center; font-size: 0.600rem; font-weight: 500; text-transform: uppercase;"
              v-bind="attrs"
              v-on="on"
            >
              <v-progress-circular
                :rotate="90"
                :size="28"
                :width="4"
                :value="getAchievedPercent(i)"
                :color="(getAchievedPercent(i) < 5) ? 'orange' : (getAchievedPercent(i) < 15) ? 'purple' : (getAchievedPercent(i) < 50) ? 'cyan' : 'green'"
              >
                {{ Math.floor(getAchievedPercent(i)).toFixed(0) }}
              </v-progress-circular>
            </div>
          </template>
          <span>Won by <strong>{{ getAchievedPercent(i) }}</strong>% of possible players</span>
        </v-tooltip>
      </v-list-item>
    </v-list>
  </v-container>
</template>
<style>
.achievement:hover .tooltip {
    display: flex;
}

.tooltip {
    display: none;
    margin-left: 20px;
    margin-top: -50px;
    position: fixed;
    z-index: 1000;
    min-width:250px;
    max-width:400px;
    border-radius: 5px;
}
</style>
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

    mouseEnter(event) {
      this.$el.addEventListener('mousemove', this.mouseMove, false);
    },

    mouseMove(event){
      var tooltip = document.querySelectorAll('.tooltip');

      for (var i=tooltip.length; i--;) {
          tooltip[i].style.left = event.clientX + 'px';
          tooltip[i].style.top = event.clientY + 'px';
      }
    },

    mouseLeave(event){
      
    },

    earnedStyle(i) {
      if(this.game.achievements[i][this.gameMode]) {
        return {
            'opacity' : 1,
        }
      } else {
        return {
            'opacity' : 0.25,
        }
      }
    },

    getAchievedPercent(i) {
      if (this.game === null) return 0;
      var currentMode = (this.gameMode === 'DateEarnedHardcore') ? this.game.achievements[i].NumAwardedHardcore : this.game.achievements[i].NumAwarded;
      return Math.floor(Number(currentMode) / Number(this.game.softcore) * 10000) / 100;
    },

    getAchievedDate(i) {
      var dateEarned = (this.gameMode === 'DateEarnedHardcore') ? this.game.achievements[i].DateEarnedHardcore : this.game.achievements[i].DateEarned;
      return dateEarned;
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