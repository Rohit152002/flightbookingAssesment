<template>
    <div class="max-w-2xl mx-auto mt-5">
        <h2 class="mb-6 text-3xl font-bold text-gray-900">Payment Details</h2>
        <form @submit.prevent="submitPayment">
            <div class="mb-5">
                <label for="ccNumber" class="block text-sm font-medium text-gray-800">Credit Card Number</label>
                <input type="text" id="ccNumber"
                    class="mt-1 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:ring-4 focus:ring-blue-600 focus:outline-none"
                    v-model="payment.ccNumber" @input="validateCCNumber" :class="{ 'border-red-500': errors.ccNumber }"
                    placeholder="Enter your credit card number" maxlength="16" required>
                <div v-if="errors.ccNumber" class="mt-2 text-xs text-red-600">{{ errors.ccNumber }}</div>
            </div>
            <div class="mb-5">
                <label for="cvv" class="block text-sm font-medium text-gray-800">CVV</label>
                <input type="text" id="cvv"
                    class="mt-1 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:ring-4 focus:ring-blue-600 focus:outline-none"
                    v-model="payment.cvv" @input="validateCVV" :class="{ 'border-red-500': errors.cvv }"
                    placeholder="Enter CVV" maxlength="3" required>
                <div v-if="errors.cvv" class="mt-2 text-xs text-red-600">{{ errors.cvv }}</div>
            </div>
            <div class="mb-5">
                <label for="expiryDate" class="block text-sm font-medium text-gray-800">Expiry Date (MM/YY)</label>
                <input type="text" id="expiryDate"
                    class="mt-1 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:ring-4 focus:ring-blue-600 focus:outline-none"
                    v-model="payment.expiryDate" @input="validateExpiryDate"
                    :class="{ 'border-red-500': errors.expiryDate }" placeholder="MM/YY" maxlength="5" required>
                <div v-if="errors.expiryDate" class="mt-2 text-xs text-red-600">{{ errors.expiryDate }}</div>
            </div>
            <div class="flex space-x-4 mt-6">
                <button type="submit"
                    class="w-full py-3 px-6 bg-blue-600 text-white rounded-lg shadow-lg transform transition duration-200 hover:scale-105 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50">
                    Submit Payment
                </button>
                <button type="button" @click="goBack"
                    class="w-full py-3 px-6 bg-gray-500 text-white rounded-lg shadow-lg transform transition duration-200 hover:scale-105 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-opacity-50">
                    Cancel
                </button>
            </div>

        </form>

        <div v-if="bookingReference"
            class="mt-6 p-5 bg-green-100 text-green-800 border border-green-300 rounded-lg shadow-md">
            <h5 class="text-xl font-semibold">Booking Successful!</h5>
            <p>Your booking reference number is: <strong>{{ bookingReference }}</strong></p>
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
            bookingReference: null,
            currentYear: new Date().getFullYear(),
            errors: {}
        };
    },
    methods: {
        validateCCNumber() {
            const regex = /^\d{16}$/;
            this.errors.ccNumber = regex.test(this.payment.ccNumber) ? '' : 'Invalid credit card number';
        },
        validateCVV() {
            const regex = /^\d{3}$/;
            this.errors.cvv = regex.test(this.payment.cvv) ? '' : 'Invalid CVV';
        },
        validateExpiryDate() {
            const regex = /^(0[1-9]|1[0-2])\/\d{2}$/; // MM/YY format
            if (!regex.test(this.payment.expiryDate)) {
                this.errors.expiryDate = 'Invalid expiry date format (MM/YY)';
            } else {
                const [month, year] = this.payment.expiryDate.split('/');
                const expiryYear = parseInt(`20${year}`, 10); // Assume YY is 20XX
                const expiryMonth = parseInt(month, 10);
                const currentDate = new Date();
                const currentMonth = currentDate.getMonth() + 1; // Month is zero-indexed
                const currentYear = currentDate.getFullYear();

                if (expiryYear < currentYear || (expiryYear === currentYear && expiryMonth < currentMonth)) {
                    this.errors.expiryDate = 'Expiry date must be in the future';
                } else {
                    this.errors.expiryDate = '';
                }
            }
        },
        goBack() {
            // Go back to the previous page in the browser's history
            this.$router.back()
        },
        async submitPayment() {
            this.validateCCNumber();
            this.validateCVV();
            this.validateExpiryDate();

            // Only submit if no errors
            if (!Object.values(this.errors).some(error => error)) {
                // Simulate API submission
                this.bookingReference = '123ABC456'; // Example booking reference
            }
        }
    }
};
</script>

<style scoped>
/* Additional custom styles if necessary */
</style>
