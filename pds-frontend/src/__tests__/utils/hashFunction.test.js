import { hashFunction } from "../../utils";
describe("hashFunction", () => {
  it("should return hashed data", () => {
    const mockReturnValue = "hashed data";
    const mockMd5Function = jest.fn().mockReturnValue(mockReturnValue);
    const result = hashFunction(mockReturnValue, mockMd5Function);
    expect(result).toBe(mockReturnValue);
  });
});
