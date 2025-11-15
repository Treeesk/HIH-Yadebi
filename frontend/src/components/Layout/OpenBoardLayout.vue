<template>
  <div class="open-board-layout">
    <div class="logo-section">
      <img
          src="@/assets/logos/rustore-logo.svg"
          alt="RuStore"
          class="rustore-logo"
      >
      <h1 class="welcome-title">Добро пожаловать в RuStore!</h1>
      <p class="welcome-subtitle">Откройте для себя мир российских приложений</p>
    </div>

    <div class="actions-section">
      <button class="auth-button primary" @click="showLoginModal = true">
        войти/зарегистрироваться
      </button>

      <button class="auth-button secondary" @click="goToHome">
        Пропустить
      </button>
    </div>

    <LoginModal
        v-if="showLoginModal"
        @close="showLoginModal = false"
        @switch-to-register="switchToRegister"
    />

    <RegisterModal
        v-if="showRegisterModal"
        @close="showRegisterModal = false"
        @switch-to-login="switchToLogin"
        @register-success="handleRegisterSuccess"
    />

    <ConfirmEmailModal
        v-if="showConfirmModal"
        :user-email="userEmail"
        @close="showConfirmModal = false"
        @confirm-success="handleConfirmSuccess"
    />
  </div>
</template>

<script>
import LoginModal from '@/components/Auth/LoginModal.vue'
import RegisterModal from '@/components/Auth/RegisterModal.vue'
import ConfirmEmailModal from '@/components/Auth/ConfirmEmailModal.vue'

export default {
  name: 'OpenBoardLayout',
  components: { LoginModal, RegisterModal, ConfirmEmailModal },

  data() {
    return {
      showLoginModal: false,
      showRegisterModal: false,
      showConfirmModal: false,
      userEmail: ''
    }
  },

  methods: {
    goToHome() {
      this.$router.push('/')
    },

    switchToRegister() {
      this.showLoginModal = false
      this.showRegisterModal = true
    },

    switchToLogin() {
      this.showRegisterModal = false
      this.showLoginModal = true
    },

    handleRegisterSuccess(email) {
      this.userEmail = email
      this.showConfirmModal = true
    },

    handleConfirmSuccess(token) {
      this.$router.push('/')
    }
  }
}
</script>

<style scoped>
.open-board-layout {
  min-height: 100vh;
  background: #FFFFFF;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
  color: #000000;
  text-align: center;
  /* ❗ убрали h=100vh, fixed, overflow-hidden */
}

.logo-section {
  text-align: center;
  margin-bottom: 60px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 1;
}

.rustore-logo {
  width: 200px;
  height: 200px;
  margin-bottom: 30px;
}

.welcome-title {
  font-size: 28px;
  font-weight: bold;
  margin-bottom: 25px;
  color: #000000;
  line-height: 1.3;
}

.welcome-subtitle {
  font-size: 20px;
  color: #333333;
  margin-bottom: 0;
  line-height: 1.5;
  margin-top: 10px;
}

.actions-section {
  display: flex;
  flex-direction: row;
  gap: 15px;
  width: 100%;
  max-width: 400px;
  justify-content: center;
  margin-top: auto;
  margin-bottom: 40px;
}

.auth-button {
  padding: 16px 25px;
  border: none;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  flex: 1;
  max-width: 200px;
}

.auth-button.primary {
  background: #1E3A8A;
  color: #FFFFFF;
  border: 2px solid #1E3A8A;
}

.auth-button.secondary {
  background: transparent;
  color: #1E3A8A;
  border: 2px solid #1E3A8A;
}
</style>