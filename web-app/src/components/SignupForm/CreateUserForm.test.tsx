import { fireEvent, render } from "@testing-library/react";
import { CreateUserForm } from "./CreateUserForm";

const renderComponent = () => {
  return render(<CreateUserForm />);
};

describe("SignupForm", () => {
  it("should render form", async () => {
    const component = renderComponent();
    const emailInput = await component.findByPlaceholderText("Email");
    const passwordInput = await component.findByPlaceholderText("Password");
    const nameInput = await component.findByPlaceholderText("Name");
    const lastNameInput = await component.findByPlaceholderText("Last Name");
    expect(emailInput).toBeInTheDocument();
    expect(passwordInput).toBeInTheDocument();
    expect(nameInput).toBeInTheDocument();
    expect(lastNameInput).toBeInTheDocument();
  });

  it("should show go to login ,  'Already have an account? ", async () => {
    const component = renderComponent();

    const goToLoginText = await component.findByText(
      "Already have an account?"
    );

    expect(goToLoginText).toBeInTheDocument();
  });

  it("should fill form with bad entries", async () => {
    const component = renderComponent();
    const emailInput = await component.findByPlaceholderText("Email");
    const passwordInput = await component.findByPlaceholderText("Password");
    const nameInput = await component.findByPlaceholderText("Name");
    const lastNameInput = await component.findByPlaceholderText("Last Name");

    fireEvent.change(emailInput, { target: { value: "badEmail" } });
    fireEvent.change(passwordInput, { target: { value: "badpass" } });
    fireEvent.change(nameInput, { target: { value: "" } });
    fireEvent.change(lastNameInput, { target: { value: "" } });

    const submittButton = await component.findByText("Create Account");

    fireEvent.click(submittButton);

    const errorMessage = await component.findByText("Name is required");

    expect(errorMessage).toBeInTheDocument();
  });
});
