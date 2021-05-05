package iemail

type IEmail interface {
	EmailRegisterUserConfirmation(email, firstName, confirmationLink string) error
}
