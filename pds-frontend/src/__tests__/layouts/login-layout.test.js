import { LoginLayout } from "../../components";
import { render, screen } from "@testing-library/react";

describe("LoginLayout", () => {
  it("should render login layout", () => {
    render(<LoginLayout>test</LoginLayout>);
    expect(screen.getByText("test")).toBeTruthy();
  });
});
