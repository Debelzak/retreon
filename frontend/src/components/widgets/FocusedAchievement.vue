<template>
  <v-container pb-0 v-if="game">
    <div v-if="mastered === true">
      <div class="overline mb-1">
        No Remaining Goals
      </div>
      <v-img
        :src="`http://i.retroachievements.org.s3.amazonaws.com${this.game.icon}`"
        max-width="96"
        max-height="96"
        class="elevation-8 float-left mr-4"
      ></v-img>
      <div>
        <div class="subtitle-1 font-weight-bold text-center">SET MASTERED!</div>
        <v-divider></v-divider>
        <div style="overflow:hidden" class="my-3">
          <div
            :class="{
              'font-weight-light': true,
              'body-2': true,
              'text-no-wrap': settings.marquee,
            }"
            :id="settings.marquee ? 'marquee' : ''"
          >Congratulations!<br>You achieved every challenge of this game.<br></div>
        </div>
      </div>
    </div>
    <div v-else>
      <div class="overline mb-1">
        Current Goal - {{ focused.Points }} ({{ focused.TrueRatio }}) points
      </div>
      <v-img
        :src="`http://i.retroachievements.org.s3.amazonaws.com/Badge/${focused.BadgeName}.png`"
        max-width="96"
        max-height="96"
        class="elevation-8 float-left mr-4"
      ></v-img>
      <div>
        <div class="subtitle-1 font-weight-bold text-center">
          {{ focused.Title }}
        </div>
        <v-divider></v-divider>
        <div style="overflow:hidden" class="my-3">
          <div
            :class="{
              'font-weight-light': true,
              'body-2': true,
              'text-no-wrap': settings.marquee,
              'text-center': true,
            }"
            :id="settings.marquee ? 'marquee' : ''"
          >{{ focused.Description }}</div>
        </div>
      </div>
      <div
        v-if="settings.softPercent"
        class="text-center caption"
      >
        Won by {{ softcorePercent }}% of possible players. [Normal]
      </div>
      <div
        v-if="settings.hardPercent"
        class="text-center caption"
      >
        Won by {{ hardcorePercent }}% of possible players. [Hardcore]
      </div>
    </div>
  </v-container>
</template>

<script>
import { mapState } from 'vuex';

export default {
  props: ['settings'],

  data: () => ({
    mastered: false,
  }),

  computed: {
    ...mapState('game', ['game', 'gameMode']),

    focused() {
      if (this.game === null) return {};
      const candidates = this.game.order
        .filter(x => !this.game.achievements[x][this.gameMode]);
      
      // If Mastered achievement
      this.setMastered(candidates.length === 0);
      if (candidates.length === 0) {
        return {};
      }

      return this.game.achievements[candidates[0]];
    },
    softcorePercent() {
      if (this.game === null) return 0;
      return Math.floor(Number(this.focused.NumAwarded) / Number(this.game.softcore) * 10000) / 100;
    },
    hardcorePercent() {
      if (this.game === null) return 0;
      return Math.floor(Number(this.focused.NumAwardedHardcore) / Number(this.game.softcore) * 10000) / 100;
    },
  },

  methods: {
    setMastered(value) {
      this.mastered = value;
    }
  }
};
</script>

<style scoped>
#marquee {
  animation: scroll 10s linear 0s infinite;
  display: inline-block;
}
@keyframes scroll {
    from {
        margin-left: 100%;
        transform: translate(0%,0)
    }

    to {
        margin-left: 0;
        transform: translate(-100%,0)
    }
}
</style>