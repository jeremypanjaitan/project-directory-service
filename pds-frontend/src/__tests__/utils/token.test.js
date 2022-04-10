import { deleteToken, setToken } from "../../utils";
describe("setToken", () => {
  it("should set token", () => {
    const result = setToken();
    expect(result).toBeUndefined();
  });
});

describe("deleteToken", () => {
  it("should set token", () => {
    const result = deleteToken();
    expect(result).toBeUndefined();
  });
});
