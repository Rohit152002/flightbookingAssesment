package services

import (
	"flight/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPaymentValidation(t *testing.T) {
	tests := []struct {
		name    string
		payment *models.PaymentModel
		want    bool
		wantErr bool
	}{
		{
			name: "Valid payment",
			payment: &models.PaymentModel{
				CardNumber: "4532015112830366",
				CVV:        "123",
				ExpiryDate: "12/25",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Invalid card number",
			payment: &models.PaymentModel{
				CardNumber: "1234567890123456",
				CVV:        "123",
				ExpiryDate: "12/25",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Invalid CVV",
			payment: &models.PaymentModel{
				CardNumber: "4532015112830366",
				CVV:        "12",
				ExpiryDate: "12/25",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Invalid expiry date format",
			payment: &models.PaymentModel{
				CardNumber: "4532015112830366",
				CVV:        "123",
				ExpiryDate: "1225",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Invalid expiry month",
			payment: &models.PaymentModel{
				CardNumber: "4532015112830366",
				CVV:        "123",
				ExpiryDate: "13/25",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Invalid expiry year",
			payment: &models.PaymentModel{
				CardNumber: "4532015112830366",
				CVV:        "123",
				ExpiryDate: "12/20",
			},
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckPaymentValidation(tt.payment)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
