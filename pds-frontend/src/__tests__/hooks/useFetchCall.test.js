import { renderHook } from "@testing-library/react-hooks";
import { useFetchCall } from "../../hooks";
describe("useFetchCall", () => {
  it("should render useFetchCall hooks", () => {
    const { result } = renderHook(useFetchCall);
    expect(result.current.isLoading).toBeFalsy();
    expect(result.current.isError).toBeFalsy();
    expect(result.current.isSuccess).toBeFalsy();
    expect(result.current.error).toBeUndefined();
    expect(result.current.data).toBeUndefined();
  });
  it("should call start with resolved value", async () => {
    const mockFetchCall = jest.fn().mockResolvedValue("success");
    const mockSetState = jest.fn();
    const mockUseState = jest.fn().mockReturnValue([null, mockSetState]);
    const { result } = renderHook(() =>
      useFetchCall(mockFetchCall, mockUseState)
    );
    result.current.start();
    expect(mockSetState).toHaveBeenCalledTimes(1);
    expect(mockUseState).toHaveBeenCalledTimes(5);
  });
  it("should call start rejected value", async () => {
    const mockFetchCall = jest.fn().mockRejectedValue("failed");
    const mockSetState = jest.fn();
    const mockUseState = jest.fn().mockReturnValue([null, mockSetState]);
    const { result } = renderHook(() =>
      useFetchCall(mockFetchCall, mockUseState)
    );
    result.current.start();
    expect(mockSetState).toHaveBeenCalledTimes(1);
    expect(mockUseState).toHaveBeenCalledTimes(5);
  });

  it("should call reset rejected value", async () => {
    const mockFetchCall = jest.fn().mockRejectedValue("failed");
    const mockSetState = jest.fn();
    const mockUseState = jest.fn().mockReturnValue([null, mockSetState]);
    const { result } = renderHook(() =>
      useFetchCall(mockFetchCall, mockUseState)
    );
    result.current.reset();
    expect(mockSetState).toHaveBeenCalledTimes(5);
    expect(mockUseState).toHaveBeenCalledTimes(5);
  });
});
