<template>
  <v-container>
    <v-row no-gutters justify="center" class="mt-2">
      <v-col
        v-if="game"
        cols="auto"
        class="title font-weight-light text-center"
      >{{ game.title }}</v-col>
      <v-col v-else>
        No game loaded :(
      </v-col>
    </v-row>
    <v-row no-gutters justify="center">
      <v-col cols="auto">
        <v-switch
          v-model="mode"
          color="grey lighten-1"
        >
          <template v-slot:prepend>
            <div :class="softcoreStyle">Normal</div>
          </template>
          <template v-slot:append>
            <div :class="hardcoreStyle">Hardcore</div>
          </template>
        </v-switch>
      </v-col>
    </v-row>
    <v-row
      no-gutters 
      justify="center"
      class="mt-4"
    >
      <v-col
        cols="auto"
        class="font-weight-light"
      >Achievement Order (Drag to Change)</v-col>
    </v-row>

    <v-list style="display: table-row;">
      <draggable
        class="list-group"
        v-model="list"
        @change="checkChange"
      >
        <v-list-item
          v-for="(i, index) in list"
          :key="i"
          style="float: left; padding: 0px;"
        >
          <v-list-item-avatar size="48" tile style="margin: 6px; cursor: move">
            <label style="width: 20px; height: 20px; position: absolute; top: 0px; left: 0px; background-color: grey; border-radius: 2px;"><b>{{ index+1 }}</b></label>
            <img :src="`http://i.retroachievements.org.s3.amazonaws.com/Badge/${game.achievements[i].BadgeName}.png`">
          </v-list-item-avatar>
        </v-list-item>
      </draggable>
    </v-list>
  </v-container>
</template>

<script>
import draggable from 'vuedraggable';
import { mapState, mapActions } from 'vuex';

export default {
  components: {
    draggable,
  },

  computed: {
    ...mapState('game', ['game', 'gameMode']),

    list: {
      get() {
        if (this.game === null) return [];
        return this.game.order
          .filter(x => !this.game.achievements[x][this.gameMode]);
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
    hardcoreStyle() {
      return {
        'mt-1': true,
        'font-weight-bold': this.mode,
        'white--text': this.mode,
        'grey--text': !this.mode,
      };
    },
    softcoreStyle() {
      return {
        'mt-1': true,
        'font-weight-bold': !this.mode,
        'white--text': !this.mode,
        'grey--text': this.mode,
      };
    },
  },

  methods: {
    ...mapActions('game', ['setGameOrder', 'setGameMode']),

    checkChange(data) {
      if (!data.moved) return;
      // determine where in the master list the element should go
      let goalIndex = 0;
      if (data.moved.newIndex > 0) {
        const destinationIndex = this.list[data.moved.newIndex - 1];
        goalIndex = this.game.order.indexOf(destinationIndex) + 1;
      }
      // remove the entry and reinsert it correctly
      this.game.order.splice(this.game.order.indexOf(data.moved.element), 1);
      this.game.order.splice(goalIndex, 0, data.moved.element);

      this.setGameOrder({ game: this.game.id, order: this.game.order });
    },
  },
};
</script>

<style scoped>
.sortable-chosen {
  background-color: #1e1e1e;
}
</style>
