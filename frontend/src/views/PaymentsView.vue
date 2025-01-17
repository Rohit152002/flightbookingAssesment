<template>
    <div class="min-h-screen bg-gray-50 py-12 px-4 sm:px-6">
      <div class="max-w-xl mx-auto bg-white rounded-xl shadow-sm p-8">
        <!-- Header -->
        <div class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-800">Payment Details</h2>
          <p class="mt-2 text-sm text-gray-500">Please enter your payment information securely</p>
        </div>

        <form @submit.prevent="submitPayment" class="space-y-6">
          <!-- Credit Card Number -->
          <div class="space-y-1">
            <label for="ccNumber" class="text-sm font-medium text-gray-700">
              Card Number
            </label>
            <input
              type="text"
              id="ccNumber"
              v-model="payment.ccNumber"
              @input="validateCCNumber"
              maxlength="16"
              required
              placeholder="1234 5678 9012 3456"
              :class="[
                'w-full px-4 py-3 rounded-lg border transition-colors duration-200',
                'placeholder-gray-400 text-gray-900',
                'focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 focus:outline-none',
                errors.ccNumber ? 'border-red-300 bg-red-50' : 'border-gray-200'
              ]"
            />
            <p v-if="errors.ccNumber" class="text-sm text-red-500 mt-1">{{ errors.ccNumber }}</p>
          </div>

          <!-- CVV and Expiry Date Grid -->
          <div class="grid grid-cols-2 gap-4">
            <!-- CVV -->
            <div class="space-y-1">
              <label for="cvv" class="text-sm font-medium text-gray-700">
                CVV
              </label>
              <input
                type="text"
                id="cvv"
                v-model="payment.cvv"
                @input="validateCVV"
                maxlength="3"
                required
                placeholder="123"
                :class="[
                  'w-full px-4 py-3 rounded-lg border transition-colors duration-200',
                  'placeholder-gray-400 text-gray-900',
                  'focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 focus:outline-none',
                  errors.cvv ? 'border-red-300 bg-red-50' : 'border-gray-200'
                ]"
              />
              <p v-if="errors.cvv" class="text-sm text-red-500 mt-1">{{ errors.cvv }}</p>
            </div>

            <!-- Expiry Date -->
            <div class="space-y-1">
              <label for="expiryDate" class="text-sm font-medium text-gray-700">
                Expiry Date
              </label>
              <input
                type="text"
                id="expiryDate"
                v-model="payment.expiryDate"
                @input="validateExpiryDate"
                maxlength="5"
                required
                placeholder="MM/YY"
                :class="[
                  'w-full px-4 py-3 rounded-lg border transition-colors duration-200',
                  'placeholder-gray-400 text-gray-900',
                  'focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 focus:outline-none',
                  errors.expiryDate ? 'border-red-300 bg-red-50' : 'border-gray-200'
                ]"
              />
              <p v-if="errors.expiryDate" class="text-sm text-red-500 mt-1">{{ errors.expiryDate }}</p>
            </div>
          </div>

          <!-- Action Buttons -->
          <div class="flex gap-4 pt-4">
            <button
            :disabled="!checkValidation"
              type="submit"
              class="flex-1 bg-blue-600 text-white py-3 px-6 rounded-lg font-medium
                     hover:bg-blue-700 transition-colors duration-200
                     focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"

            >
              Complete Payment
            </button>
            <button
              type="button"
              @click="goBack"
              class="flex-1 bg-gray-100 text-gray-700 py-3 px-6 rounded-lg font-medium
                     hover:bg-gray-200 transition-colors duration-200
                     focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            >
              Back
            </button>
          </div>
        </form>

        <!-- Success Message -->
        <div
          v-if="bookingReference"
          class="mt-8 bg-green-50 border border-green-100 rounded-lg p-6 space-y-2"
        >
          <div class="flex items-center gap-2">
            <div class="w-2 h-2 bg-green-500 rounded-full"></div>
            <h3 class="text-lg font-semibold text-green-800">Payment Successful</h3>
          </div>
          <p class="text-green-700">
            Your booking reference number is:
            <span class="font-mono font-medium">{{ bookingReference }}</span>
          </p>
        </div>
      </div>
    </div>
  </template>

  <script>
  export default {
    name: 'PaymentDetails',
    data() {
      return {
        payment: {
          ccNumber: '',
          cvv: '',
          expiryDate: ''
        },
        validate:false,
        bookingReference: null,
        currentYear: new Date().getFullYear(),
        errors: {}
      }
    },
    computed:{
        checkValidation()
        {
            if (this.errors==null){
                this.validate=true
                return true
            }
            return false
        }
    },
    methods: {
      validateCCNumber() {
        const regex = /^\d{16}$/
        this.errors.ccNumber = regex.test(this.payment.ccNumber)
          ? ''
          : 'Please enter a valid 16-digit card number'
      },
      validateCVV() {
        const regex = /^\d{3}$/
        this.errors.cvv = regex.test(this.payment.cvv)
          ? ''
          : 'Please enter a valid 3-digit CVV'
      },
      validateExpiryDate() {
        const regex = /^(0[1-9]|1[0-2])\/\d{2}$/
        if (!regex.test(this.payment.expiryDate)) {
          this.errors.expiryDate = 'Please use MM/YY format'
          return
        }

        const [month, year] = this.payment.expiryDate.split('/')
        const expiryYear = parseInt(`20${year}`, 10)
        const expiryMonth = parseInt(month, 10)
        const currentDate = new Date()
        const currentMonth = currentDate.getMonth() + 1
        const currentYear = currentDate.getFullYear()

        if (expiryYear < currentYear || (expiryYear === currentYear && expiryMonth < currentMonth)) {
          this.errors.expiryDate = 'Card has expired'
        } else {
          this.errors.expiryDate = ''
        }
      },
      goBack() {
        this.$router.back()
      },
      async submitPayment() {
        this.validateCCNumber()
        this.validateCVV()
        this.validateExpiryDate()

        if (!Object.values(this.errors).some(error => error)) {
          // Simulate API submission
          this.bookingReference = '123ABC456'
        }
      }
    }
  }
  </script>
