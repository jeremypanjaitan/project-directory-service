import MainLayoutContainer from "../../components/layouts/main-layout/MainLayoutContainer";
import { render, screen } from "@testing-library/react";

describe("Main Layout Container", () => {
  it("should render main layout container", () => {
    const MainLayoutViewTest = ({ handleLogout, children }) => {
      handleLogout();
      return children;
    };
    const mockSetUser = jest.fn();
    const mockWindow = {
      confirm: jest.fn().mockReturnValue(true),
    };
    const mockUseLogout = jest
      .fn()
      .mockImplementation(() => ({ isSuccess: false, start: jest.fn() }));
    const mockUseAuth = jest
      .fn()
      .mockImplementation(() => ({ setUser: mockSetUser }));
    const mockSetIsLoading = jest.fn();
    const mockUseLoading = jest.fn().mockImplementation(() => ({
      isLoading: true,
      setIsLoading: mockSetIsLoading,
    }));
    render(
      <MainLayoutContainer
        _useLogout={mockUseLogout}
        _useAuth={mockUseAuth}
        _useLoading={mockUseLoading}
        _window={mockWindow}
        MainLayoutViewTest={MainLayoutViewTest}
      >
        test
      </MainLayoutContainer>
    );
    expect(screen.getByText("test")).toBeTruthy();
  });
  it("should render main layout container loading", () => {
    const MainLayoutViewTest = ({ handleLogout, children }) => {
      handleLogout();
      return children;
    };
    const mockSetUser = jest.fn();
    const mockWindow = {
      confirm: jest.fn().mockReturnValue(true),
    };
    const mockUseLogout = jest
      .fn()
      .mockImplementation(() => ({ isSuccess: true, start: jest.fn() }));
    const mockUseAuth = jest
      .fn()
      .mockImplementation(() => ({ setUser: mockSetUser }));
    const mockSetIsLoading = jest.fn();
    const mockUseLoading = jest.fn().mockImplementation(() => ({
      isLoading: true,
      setIsLoading: mockSetIsLoading,
    }));
    render(
      <MainLayoutContainer
        _useLogout={mockUseLogout}
        _useAuth={mockUseAuth}
        _useLoading={mockUseLoading}
        _window={mockWindow}
        MainLayoutViewTest={MainLayoutViewTest}
      >
        test
      </MainLayoutContainer>
    );
    expect(screen.getByText("test")).toBeTruthy();
  });
});
