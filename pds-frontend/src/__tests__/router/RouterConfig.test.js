import { RouterConfig } from "../../router";

jest.mock("../../context");

describe("Router Config", () => {
  it("should return router config", () => {
    const routerConfig = RouterConfig();
    expect(routerConfig).toBeTruthy();
  });
});
