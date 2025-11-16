<template>
  <div v-if="visible" class="tutorial-overlay">

    <div class="overlay-backdrop"></div>

    <div class="highlight" :style="highlightStyle"></div>

    <img
        src="@/assets/tutorial-arrow.svg"
        class="arrow"
        :style="arrowStyle"
    />

    <div class="bubble" :style="bubbleStyle">
      <p>{{ steps[current].text }}</p>
      <button @click="next">Далее</button>
    </div>

  </div>
</template>

<script>
export default {
  name: "TutorialOverlay",

  props: {
    steps: { type: Array, required: true }
  },

  emits: ["finish"],

  data() {
    return {
      visible: true,
      current: 0,
      rect: null
    }
  },

  mounted() {
    this.calculate()
    window.addEventListener("resize", this.calculate)
  },

  beforeUnmount() {
    window.removeEventListener("resize", this.calculate)
  },

  methods: {
    next() {
      if (this.current === this.steps.length - 1) {
        this.visible = false
        this.$emit("finish")
        return
      }

      this.current++
      this.calculate()
    },

    calculate() {
      this.$nextTick(() => {
        setTimeout(() => {
          const el = document.querySelector(this.steps[this.current].el)
          if (!el) return

          const r = el.getBoundingClientRect()

          this.rect = {
            top: r.top + window.scrollY - 8,
            left: r.left + window.scrollX - 8,
            width: r.width + 16,
            height: r.height + 16
          }
        }, 50)
      })
    }
  },

  computed: {
    highlightStyle() {
      if (!this.rect) return {}
      return {
        top: this.rect.top + "px",
        left: this.rect.left + "px",
        width: this.rect.width + "px",
        height: this.rect.height + "px"
      }
    },

    arrowStyle() {
      if (!this.rect) return {}
      return {
        top: this.rect.top - 40 + "px",
        left: this.rect.left + this.rect.width / 2 + "px"
      }
    },

    bubbleStyle() {
      if (!this.rect) return {}
      return {
        top: this.rect.bottom + 16 + "px",
        left: this.rect.left + "px"
      }
    }
  }
}
</script>

<style scoped>
.tutorial-overlay {
  position: fixed;
  inset: 0;
  z-index: 9999;
}

.overlay-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.65);
  backdrop-filter: blur(2px);
}

.highlight {
  position: absolute;
  border-radius: 12px;
  box-shadow: 0 0 0 9999px rgba(0,0,0,0.65);
}

.arrow {
  position: absolute;
  width: 32px;
  height: 32px;
  transform: translate(-50%, -100%);
}

.bubble {
  position: absolute;
  background: white;
  padding: 14px 16px;
  border-radius: 12px;
  max-width: 220px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.25);
}

.bubble p {
  margin: 0 0 8px;
  color: #111;
  font-size: 14px;
}

.bubble button {
  margin-top: 10px;
  padding: 8px 14px;
  background: #3b82f6 !important;
  color: #fff !important;
  border: none;
  font-weight: 500;
  border-radius: 8px;
  cursor: pointer;
}
</style>