import { formatDate } from "../../utils";

describe("format date", () => {
  it("should return format date", () => {
    const mockReturnValue = "test";
    const mockMoment = jest
      .fn()
      .mockReturnValue({ format: jest.fn().mockReturnValue(mockReturnValue) });
    const result = formatDate(mockReturnValue, mockMoment);
    expect(result).toBe(mockReturnValue);
  });
});
