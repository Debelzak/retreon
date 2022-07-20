<template>
  <v-container v-if="locked.length > 0" pb-0>
    <div class="overline mb-1">
      Next {{ locked.length }} Locked Achievement{{ locked.length === 1 ? '' : 's' }}
    </div>
    <v-row no-gutters justify="center">
      <v-tooltip
        v-for="(achievement, i) in locked"
        :key="i"
        top
      >
        <template v-slot:activator="{ on, attrs }">
          <v-img
            :src="`${base}/Badge/${achievement.BadgeName}_lock.png`"
            :max-width="settings.iconHeight"
            :width="settings.iconHeight"
            class="mr-1 mt-1 elevation-4"
            v-bind="attrs"
            v-on="settings.showTooltips ? on : null"
          ></v-img>
        </template>
        <div v-if="settings.showTooltips">
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

    locked() {
      if (this.game === null) return [];
      return this.game.order
        .filter(x => !this.game.achievements[x][this.gameMode])
        .slice(0, this.settings.count)
        .map(x => this.game.achievements[x]);
    },
  },
}
</script>