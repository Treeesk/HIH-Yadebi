<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-content">
      <button class="close-button" @click="$emit('close')">×</button>

      <h2>Регистрация</h2>

      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
      </div>

      <form @submit.prevent="handleRegister">
        <div class="input-group">
          <label>Имя и фамилия</label>
          <input
              v-model="registerData.name"
              type="text"
              required
              placeholder="Введите имя и фамилию"
          >
        </div>

        <div class="input-group">
          <label>E-mail</label>
          <input
              v-model="registerData.email"
              type="email"
              required
              placeholder="email@mail.ru"
          >
        </div>

        <div class="input-group">
          <label>Пароль</label>
          <input
              v-model="registerData.password"
              type="password"
              required
              placeholder="Введите пароль"
          >
        </div>

        <button type="submit" class="submit-button" :disabled="loading">
          {{ loading ? 'Регистрация...' : 'Зарегистрироваться' }}
        </button>
      </form>

      <p class="switch-auth">
        Уже есть аккаунт?
        <a href="#" @click.prevent="$emit('switch-to-login')">Войти</a>
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'RegisterModal',
  data() {
    return {
      registerData: {
        name: '',
        email: '',
        password: ''
      },
      loading: false,
      errorMessage: ''
    }
  },
  methods: {
    async handleRegister() {
      this.loading = true
      this.errorMessage = ''

      try {
        const response = await fetch('/api/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.registerData)
        })

        if (response.ok) {
          const data = await response.json()
          this.$emit('close')
          this.$emit('register-success', this.registerData.email)

          // ДОБАВЛЯЕМ ПРОВЕРКУ НАЛИЧИЯ ROUTER И РЕДИРЕКТИМ
          if (this.$router) {
            this.$router.push('/')
          } else {
            // Если router недоступен, используем window.location
            window.location.href = '/'
          }
        } else {
          switch (response.status) {
            case 400:
              this.errorMessage = 'Некорректный запрос. Проверьте введенные данные.'
              break
            case 409:
              this.errorMessage = 'Пользователь с таким email уже существует.'
              break
            case 500:
              this.errorMessage = 'Ошибка сервера. Попробуйте позже.'
              break
            default:
              this.errorMessage = 'Произошла ошибка. Попробуйте снова.'
          }
        }
      } catch (error) {
        console.error('Register error:', error)
        this.errorMessage = 'Ошибка сети. Проверьте подключение к интернету.'
      } finally {
        this.loading = false
      }
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