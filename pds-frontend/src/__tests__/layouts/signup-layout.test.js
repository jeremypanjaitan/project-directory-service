import { SignupLayout } from "../../components";
import { render, screen } from "@testing-library/react";

describe("SignupLayout", () => {
  it("should render signup", () => {
    const mockUseLoading = jest
      .fn()
      .mockImplementation(() => ({ isLoading: false }));
    render(<SignupLayout _useLoading={mockUseLoading}>test</SignupLayout>);
    expect(screen.getByText("test")).toBeTruthy();
  });
});
