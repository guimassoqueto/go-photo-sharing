class Signup {
    private readonly email = document.getElementById("email") as HTMLInputElement;
    private readonly password = document.getElementById("password") as HTMLInputElement;
    private readonly passwordConfirm = document.getElementById("password-confirm") as HTMLInputElement;
    private readonly btnSubmitContainer = document.getElementById("btn-submit-container") as HTMLDivElement;
    private readonly emailMessageParagraph = document.getElementById("email-message") as HTMLParagraphElement;
    private readonly passwordMessageParagraph = document.getElementById("password-message") as HTMLParagraphElement;

    constructor() {
        this.validateFields();
    }

    private validateFields() {
        const emailValidated = this.isEmailValid();
        const passwordValidated = this.isPasswordValid();
        if (emailValidated && passwordValidated) {
            this.btnSubmitContainer.classList.remove("hidden");
        }
        else {
            this.btnSubmitContainer.classList.add("hidden");
        }
    }

    private isEmailValid(): boolean {
        if (!this.email.value) {
            this.emailMessageParagraph.innerText = "Email is required";
            return false;
        }
        this.emailMessageParagraph.innerText = "";
        return true
    }

    private isPasswordValid() {
        if (!this.password.value && !this.passwordConfirm.value) {
            this.passwordMessageParagraph.innerText = "Password is required\nPassword Confirmation is required"
            return false;
        }
        else if (!this.password.value) {
            this.passwordMessageParagraph.innerText = "Password is required";
            return false;
        }
        else if (!this.passwordConfirm.value) {
            this.passwordMessageParagraph.innerText = "Password Confirmation is required";
            return false;
        }
        else if (this.password.value !== this.passwordConfirm.value) {
            this.passwordMessageParagraph.innerText = "Password != password confirmation";
            return false;
        }
        this.passwordMessageParagraph.innerText = "";
        return true;
    }
}

document.addEventListener("DOMContentLoaded", () => {
    new Signup();
});
document.addEventListener("keyup", () => {
    new Signup();
});






