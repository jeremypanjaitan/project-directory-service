import { fireEvent, render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import MainLayoutView from "../../components/layouts/main-layout/MainLayoutView";

describe("main layout view", () => {
  it("should render main layout view", () => {
    const menu = {
      home: "Home",
      accountSettings: "Account Settings",
      logout: "Logout",
    };
    render(
      <MemoryRouter>
        <MainLayoutView menu={menu}>Test</MainLayoutView>
      </MemoryRouter>
    );
    fireEvent.click(screen.getByText(menu.home));
    fireEvent.click(screen.getByText(menu.accountSettings));
    expect(screen.getByText("Test")).toBeTruthy();
  });
});
