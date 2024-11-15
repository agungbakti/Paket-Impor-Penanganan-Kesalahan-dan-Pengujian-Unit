package bank_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "Tugas/bank"
)

var _ = Describe("Account", func() {
    var account *bank.Account

    BeforeEach(func() {
        account = &bank.Account{}
    })

	Describe("Deposit", func() {
		//ketika amount kurang dari atau sama dengan 0
		Context("when amount is less than or equal to zero", func() {
			//return error jika amount adalah 0
			It("should return an error when amount is zero", func() {
				err := account.Deposit(0)
				Expect(err).To(MatchError("amount must be positive"))
			})

			//return error jika amount adalah negatif
			It("should return an error when amount is negative", func() {
				err := account.Deposit(-50)
				Expect(err).To(MatchError("amount must be positive"))
			})

		})

		//ketika amount lebih besar dari 0
		Context("when amount is greater than zero", func() {
			//menambahkan amount ke balance
			It("should add amount to balance", func() {
				account.Deposit(100)
				Expect(account.GetBalance()).To(Equal(100.0))
			})
			It("should add amount to balance", func() {
				account.Deposit(50)
				account.Deposit(50)
				Expect(account.GetBalance()).To(Equal(100.0))
			})
		})
	})

	Describe("Withdraw", func() {
		//ketika amount kurang dari atau sama dengan 0
		Context("when amount is less than or equal to zero", func() {
			It("should return an error when amount is zero", func() {
				err := account.Withdraw(0)
				Expect(err).To(MatchError("amount must be positive 0.0"))
			})

			It("should return an error when amount is negative", func() {
				err := account.Withdraw(-50)
				Expect(err).To(MatchError("amount must be positive -50.0"))
			})
		})

		//ketika amount lebih besar dari balance
		Context("when amount is greater than balance", func() {
			It("should return an error when amount is greater than balance", func() {
				err := account.Withdraw(100)
				Expect(err).To(MatchError("cannot withdraw 100.0 from balance of 0.0"))
			})

			It("should return an error when amount is greater than balance", func() {
				account.Deposit(100)
				err := account.Withdraw(200)
				Expect(err).To(MatchError("cannot withdraw 200.0 from balance of 100.0"))
			})
		})

		//ketika amount valid
		Context("when amount is valid", func() {
			It("should subtract amount from balance", func() {
				account.Deposit(100)
				account.Withdraw(50)
				Expect(account.GetBalance()).To(Equal(50.0))
			})
		})
	})


	Describe("GetBalance", func() {
		//mengembalikan balance
		It("should return balance", func() {
			account.Deposit(100)
			Expect(account.GetBalance()).To(Equal(100.0))
		})
	})
})