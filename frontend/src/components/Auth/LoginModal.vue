<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-content">
      <button class="close-button" @click="$emit('close')">×</button>

      <h2>Вход</h2>

      <!-- Блок для отображения ошибок -->
      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
      </div>

      <form @submit.prevent="handleLogin">
        <div class="input-group">
          <label>E-mail</label>
          <input
              v-model="loginData.email"
              type="email"
              required
              placeholder="email@mail.ru"
              :class="{ 'error': errorMessage }"
          >
        </div>

        <div class="input-group">
          <label>Пароль</label>
          <input
              v-model="loginData.password"
              type="password"
              required
              placeholder="Введите пароль"
              :class="{ 'error': errorMessage }"
          >
        </div>

        <button type="submit" class="submit-button" :disabled="loading">
          {{ loading ? 'Вход...' : 'Войти' }}
        </button>
      </form>

      <p class="switch-auth">
        Нет аккаунта?
        <a href="#" @click.prevent="$emit('switch-to-register')">Зарегистрироваться</a>
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LoginModal',
  data() {
    return {
      loginData: {
        email: '',
        password: ''
      },
      loading: false,
      errorMessage: ''
    }
  },
  methods: {
    async handleLogin() {
      this.loading = true
      this.errorMessage = ''

      try {
        const response = await fetch('/api/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.loginData)
        })

        if (response.ok) {
          const data = await response.json()
          localStorage.setItem('auth_token', data.auth_token)
          this.$emit('close')
          this.$emit('login-success', data)
        } else {
          // Обрабатываем ТОЛЬКО те статусы, которые есть в API
          switch (response.status) {
            case 400:
              this.errorMessage = 'Некорректный запрос. Проверьте введенные данные.'
              break
            case 401:
              this.errorMessage = 'Неверный email или пароль. Пользователь не найден.'
              break
            case 500:
              this.errorMessage = 'Ошибка сервера. Попробуйте позже.'
              break
            default:
              // Для любых других статусов используем общее сообщение
              this.errorMessage = 'Произошла ошибка. Попробуйте снова.'
          }
        }
      } catch (error) {
        console.error('Login error:', error)
        this.errorMessage = 'Ошибка сети. Проверьте подключение к интернету.'
      } finally {
        this.loading = false
      }
    }
  },
  watch: {
    'loginData.email'() {
      this.errorMessage = ''
    },
    'loginData.password'() {
      this.errorMessage = ''
    }
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  padding: 30px;
  border-radius: 15px;
  width: 90%;
  max-width: 400px;
  position: relative;
  color: #1E3A8A;
}

.close-button {
  position: absolute;
  top: 15px;
  right: 15px;
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #1E3A8A;
}

h2 {
  text-align: center;
  margin-bottom: 20px;
  color: #1E3A8A;
}

.error-message {
  background: #FEE2E2;
  color: #DC2626;
  padding: 12px;
  border-radius: 8px;
  margin-bottom: 20px;
  text-align: center;
  font-size: 14px;
  border: 1px solid #FECACA;
}

.input-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: 600;
}

input {
  width: 100%;
  padding: 12px;
  border: 2px solid #E5E7EB;
  border-radius: 8px;
  font-size: 16px;
  box-sizing: border-box;
  transition: border-color 0.3s ease;
}

input:focus {
  border-color: #3B82F6;
  outline: none;
}

input.error {
  border-color: #DC2626;
}

.submit-button {
  width: 100%;
  padding: 15px;
  background: #1E3A8A;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s ease;
}

.submit-button:disabled {
  background: #9CA3AF;
  cursor: not-allowed;
}

.submit-button:hover:not(:disabled) {
  background: #3B82F6;
}

.switch-auth {
  text-align: center;
  margin-top: 20px;
}

.switch-auth a {
  color: #3B82F6;
  text-decoration: none;
}
</style>