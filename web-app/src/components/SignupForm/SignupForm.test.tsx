import { render } from "@testing-library/react";
import { SignupForm } from "./SignupForm";

const renderComponent = () => {
    return render(<SignupForm/>)
}

describe('SignupForm', () => { 
    it("should render form", async () => {
        const component = renderComponent()
        const emailInput = await component.findByPlaceholderText("Email")
        const passwordInput = await component.findByPlaceholderText("Password")
        const nameInput = await component.findByPlaceholderText("Name")
        const lastNameInput = await component.findByPlaceholderText("Last Name")
        expect(emailInput).toBeInTheDocument()
        expect(passwordInput).toBeInTheDocument()
        expect(nameInput).toBeInTheDocument()
        expect(lastNameInput).toBeInTheDocument()
    })
 })