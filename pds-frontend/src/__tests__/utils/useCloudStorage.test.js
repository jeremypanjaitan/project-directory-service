import { renderHook } from "@testing-library/react-hooks";
import { useCloudStorage } from "../../utils";
describe("useCloudStorage", () => {
  it("should return object of cloud storage functionality", () => {
    const { result } = renderHook(useCloudStorage);
    result.current.uploadPicture();
    result.current.uploadFile();
    result.current.resetProgress();
    expect(result.current.progress).toBe(0);
  });
  it("should call upload function", async () => {
    const mockSetState = jest.fn();
    const mockUseState = jest.fn().mockReturnValue([null, mockSetState]);
    const mockUploadBytesResumable = jest.fn().mockReturnValue({
      on: jest
        .fn()
        .mockImplementation((firstParam, secondParam, thirdParam) => {
          const snapshot = {
            bytesTransferred: 20,
            totalBytes: 20,
          };
          secondParam(snapshot);
        }),
    });
    const mockRef = jest.fn();
    const { result } = renderHook(() =>
      useCloudStorage(mockUploadBytesResumable, mockRef, mockUseState)
    );
    result.current.uploadPicture({ name: "name" }).then((res) => res);
    expect(result.current.progress).toBe(null);
  });
  it("should reject upload function", async () => {
    const mockSetState = jest.fn();
    const mockUseState = jest.fn().mockReturnValue([null, mockSetState]);
    const mockUploadBytesResumable = jest.fn().mockReturnValue({
      on: jest
        .fn()
        .mockImplementation((firstParam, secondParam, thirdParam) => {
          thirdParam("error");
        }),
    });
    const mockRef = jest.fn();
    const { result } = renderHook(() =>
      useCloudStorage(mockUploadBytesResumable, mockRef, mockUseState)
    );
    result.current
      .uploadPicture({ name: "name" })
      .then((res) => res)
      .catch((err) => console.log(err));
    expect(result.current.progress).toBe(null);
  });
  it("should resolve upload function", async () => {
    const mockSetState = jest.fn();
    const mockUseState = jest.fn().mockReturnValue([null, mockSetState]);
    const mockUploadBytesResumable = jest.fn().mockReturnValue({
      on: jest
        .fn()
        .mockImplementation(
          (firstParam, secondParam, thirdParam, fourthParam) => {
            fourthParam("success");
          }
        ),
      snapshot: {
        ref: "ref",
      },
    });
    const mockRef = jest.fn();
    const { result } = renderHook(() =>
      useCloudStorage(mockUploadBytesResumable, mockRef, mockUseState)
    );
    result.current
      .uploadPicture({ name: "name" })
      .then((res) => res)
      .catch((err) => err);
    expect(result.current.progress).toBe(null);
  });
});
