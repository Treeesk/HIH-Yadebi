<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-content">
      <button class="close-button" @click="$emit('close')">×</button>

      <h2>Подтверждение email</h2>

      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
      </div>

      <form @submit.prevent="handleConfirm">
        <div class="input-group">
          <label>Код из email</label>
          <input
              v-model="confirmData.code"
              type="number"
              required
              placeholder="123456"
              maxlength="6"
          >
        </div>

        <button type="submit" class="submit-button" :disabled="loading">
          {{ loading ? 'Подтверждение...' : 'Подтвердить' }}
        </button>
      </form>

      <p class="resend-code">
        Не получили код?
        <a href="#" @click.prevent="handleResend">Отправить повторно</a>
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ConfirmEmailModal',
  props: {
    userEmail: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      confirmData: {
        email: '',
        code: ''
      },
      loading: false,
      errorMessage: ''
    }
  },
  mounted() {
    this.confirmData.email = this.userEmail;
  },
  methods: {
    async handleConfirm() {
      this.loading = true;
      this.errorMessage = '';

      try {
        const response = await fetch('/api/register/confirmEmail', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.confirmData)
        });

        if (response.ok) {
          const data = await response.json();
          localStorage.setItem('auth_token', data.auth_token);
          this.$emit('close');
          this.$emit('confirm-success', data);
        } else {
          switch (response.status) {
            case 400:
              this.errorMessage = 'Некорректный запрос.';
              break;
            case 401:
              this.errorMessage = 'Неверный код подтверждения.';
              break;
            case 410:
              this.errorMessage = 'Код устарел. Запросите новый.';
              break;
            case 500:
              this.errorMessage = 'Ошибка сервера. Попробуйте позже.';
              break;
            default:
              this.errorMessage = 'Произошла ошибка. Попробуйте снова.';
          }
        }
      } catch (error) {
        this.errorMessage = 'Ошибка сети. Проверьте подключение.';
      } finally {
        this.loading = false;
      }
    },

    async handleResend() {
      // Логика повторной отправки кода
      // Можно вызвать регистрацию еще раз или отдельный endpoint
      this.errorMessage = 'Код отправлен повторно. Проверьте email.';
    }
  }
}
</script>

<style scoped>
/* Стили такие же как в LoginModal */
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

.resend-code {
  text-align: center;
  margin-top: 20px;
}

.resend-code a {
  color: #3B82F6;
  text-decoration: none;
}
</style>