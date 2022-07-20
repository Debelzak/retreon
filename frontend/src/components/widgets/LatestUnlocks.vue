<template>
  <v-container pb-0>
    <div class="overline mb-1">
      Latest {{ latest.length }} Achievement{{ latest.length === 1 ? '' : 's'}}
    </div>
    <v-row no-gutters justify="center">
      <v-tooltip
        v-for="(achievement, i) in latest"
        :key="i"
        top
      >
        <template v-slot:activator="{ on, attrs }">
          <v-img
            :src="`${base}/Badge/${achievement.BadgeName}.png`"
            :max-width="settings.iconHeight"
            :width="settings.iconHeight"
            class="mr-1 mt-1 elevation-4"
            v-bind="attrs"
            v-on="settings.showTooltips ? on : null"
          ></v-img>
        </template>
        <div>
          <div class="subtitle-2">{{ achievement.Title }}</div>
          <div class="caption">{{ achievement.Description }}</div>
        </div>
      </v-tooltip>
    </v-row>
  </v-container>
</template>

<script>
import { mapState } from 'vuex';

export default {
  props: ['settings'],

  data: () => ({
    base: 'http://i.retroachievements.org.s3.amazonaws.com',
  }),

  computed: {
    ...mapState('game', ['game', 'gameMode']),

    latest() {
      if (this.game === null) return [];

      return Object.values(this.game.achievements)
        .filter(x => x[this.gameMode])
        .sort((a,b) => b[this.gameMode].localeCompare(a[this.gameMode]))
        .slice(0, this.settings.count);
    },
  },
};
</script>
