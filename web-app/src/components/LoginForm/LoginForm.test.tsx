import {render} from "@testing-library/react"
import { LoginForm } from "./LoginForm"

const renderComponent = () => {
    return render(<LoginForm/>)
}

describe("Login Form", ()=>{
    it("should render Login Form", ()=>{
        const {container} = renderComponent()
        expect(container).toMatchSnapshot()
    })
})