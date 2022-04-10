import { render, screen } from "@testing-library/react";
import { renderHook } from "@testing-library/react-hooks";
import { useLoading, LoadingProvider } from "../../context";

describe("useLoading", () => {
  it("should return context", () => {
    const { result } = renderHook(useLoading);
    expect(result.current).toBeNull();
  });
  it("should render Loading Provider", () => {
    const mockSetState = jest.fn();
    const mockUseState = jest.fn().mockReturnValue([null, mockSetState]);
    render(<LoadingProvider _useState={mockUseState}>test</LoadingProvider>);
    expect(screen.getByText("test")).toBeTruthy();
  });
});
