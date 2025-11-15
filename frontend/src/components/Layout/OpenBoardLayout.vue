<template>
  <div class="open-board-layout">
    <!-- Логотип -->
    <div class="logo-section">
      <img
          src="@/assets/logos/rustore-logo.svg"
          alt="RuStore"
          class="rustore-logo"
      >
      <h1 class="welcome-title">Добро пожаловать в RuStore!</h1>
      <p class="welcome-subtitle">Откройте для себя мир российских приложений</p>
    </div>

    <!-- Кнопки действий -->
    <div class="actions-section">
      <button
          class="auth-button primary"
          @click="showLoginModal = true"
      >
        войти/зарегистрироваться
      </button>
      <button
          class="auth-button secondary"
          @click="goToHome"
      >
        Пропустить
      </button>
    </div>

    <!-- Модальные окна -->
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
  components: {
    LoginModal,
    RegisterModal,
    ConfirmEmailModal
  },
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
      this.$router.push('/');
    },

    switchToRegister() {
      this.showLoginModal = false;
      this.showRegisterModal = true;
    },

    switchToLogin() {
      this.showRegisterModal = false;
      this.showLoginModal = true;
    },

    handleRegisterSuccess(email) {
      this.userEmail = email;
      this.showConfirmModal = true;
    },

    handleConfirmSuccess(tokenData) {
      console.log('User confirmed email and got token:', tokenData);
      this.$router.push('/');
    }
  },

  // 🔥 Отключаем скролл только на этой странице
  mounted() {
    document.documentElement.style.overflow = 'hidden';
    document.body.style.overflow = 'hidden';
    document.body.style.position = 'fixed';
    document.body.style.width = '100%';
    document.body.style.height = '100%';
  },

  // 🔥 Возвращаем назад при уходе со страницы
  beforeUnmount() {
    document.documentElement.style.overflow = '';
    document.body.style.overflow = '';
    document.body.style.position = '';
    document.body.style.width = '';
    document.body.style.height = '';
  }
}
</script>

<style scoped>
.open-board-layout {
  min-height: 100vh;
  height: 100vh;
  background: #FFFFFF;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
  color: #000000;
  text-align: center;
  overflow: hidden;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  width: 100vw;
}

.logo-section {
  text-align: center;
  margin-bottom: 60px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 1;
  max-height: 70vh;
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

.auth-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15);
}

.auth-button.primary:hover {
  background: #3B82F6;
  border-color: #3B82F6;
}

.auth-button.secondary:hover {
  background: #1E3A8A;
  color: #FFFFFF;
}

@media (max-width: 480px) {
  .open-board-layout {
    padding: 15px;
  }

  .rustore-logo {
    width: 150px;
    height: 150px;
  }

  .welcome-title {
    font-size: 24px;
    margin-bottom: 20px;
  }

  .welcome-subtitle {
    font-size: 18px;
    margin-top: 8px;
  }

  .actions-section {
    flex-direction: column;
    max-width: 300px;
  }

  .auth-button {
    max-width: none;
  }
}
</style>
