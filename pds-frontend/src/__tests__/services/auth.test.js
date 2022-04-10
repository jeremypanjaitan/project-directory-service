import { render, screen } from "@testing-library/react";
import { renderHook } from "@testing-library/react-hooks";
import { axios } from "../../config";
import { useFetchCall } from "../../hooks";
import { Navigate } from "react-router-dom";
import {
  useAuth,
  useCheckAuth,
  useLogin,
  useLogout,
  CheckAuth,
  RequireAuth,
  AlreadyAuth,
  AuthProvider,
} from "../../services";
import { setToken, deleteToken } from "../../utils";

jest.mock("../../hooks", () => {
  return {
    useFetchCall: jest.fn(),
  };
});
jest.mock("../../config", () => {
  const originalModule = jest.requireActual("../../config");
  return {
    ...originalModule,
    axios: {
      get: jest.fn(),
      post: jest.fn(),
    },
  };
});
jest.mock("../../utils", () => {
  const originalModule = jest.requireActual("../../utils");
  return {
    ...originalModule,
    setToken: jest.fn(),
    deleteToken: jest.fn(),
  };
});

describe("useAuth", () => {
  it("should return auth context", () => {
    const { result } = renderHook(useAuth);
    expect(result.current).toBe(null);
  });
});

describe("useCheckAuth", () => {
  it("should return success auth", async () => {
    useFetchCall.mockImplementation((login) => {
      return login();
    });

    axios.get.mockResolvedValue({ data: "success" });
    const { result } = renderHook(useCheckAuth);
    const res = await result.current.then((res) => res);
    expect(res).toBe("success");
  });

  it("should return failed auth", async () => {
    useFetchCall.mockImplementation((login) => {
      return login();
    });

    axios.get.mockRejectedValue({ data: "failed" });
    const { result } = renderHook(useCheckAuth);
    const res = await result.current.catch((err) => {
      return err;
    });
    expect(res.data).toBe("failed");
  });
});

describe("useLogin", () => {
  it("should return success", async () => {
    useFetchCall.mockImplementation((login) => {
      return login();
    });
    const resolvedValue = {
      data: { data: { tokenData: { accessToken: "token" } } },
    };
    axios.post.mockResolvedValue(resolvedValue);
    const { result } = renderHook(useLogin);
    const res = await result.current.then((res) => res);
    expect(res.data.tokenData.accessToken).toBe(
      resolvedValue.data.data.tokenData.accessToken
    );
    expect(setToken).toHaveBeenCalledWith(
      resolvedValue.data.data.tokenData.accessToken
    );
  });

  it("should return failed", async () => {
    useFetchCall.mockImplementation((login) => {
      return login();
    });
    const rejectedValue = { data: "failed" };
    axios.post.mockRejectedValue(rejectedValue);
    const { result } = renderHook(useLogin);
    const res = await result.current.catch((err) => err);
    expect(res.data).toBe(rejectedValue.data);
  });
});

describe("useLogout", () => {
  it("should return success", async () => {
    useFetchCall.mockImplementation((login) => {
      return login();
    });
    const resolvedValue = { data: { token: "token" } };
    axios.post.mockResolvedValue(resolvedValue);
    const { result } = renderHook(useLogout);
    const res = await result.current.then((res) => res);
    expect(res.token).toBe(resolvedValue.data.token);
    expect(deleteToken).toBeCalledTimes(1);
  });

  it("should return failed", async () => {
    useFetchCall.mockImplementation((login) => {
      return login();
    });
    const rejectedValue = { data: "failed" };
    axios.post.mockRejectedValue(rejectedValue);
    const { result } = renderHook(useLogout);
    const res = await result.current.catch((err) => err);
    expect(res.data).toBe(rejectedValue.data);
  });
});

describe("Check Auth", () => {
  it("should render children", () => {
    const startMock = jest.fn();
    const setUserMock = jest.fn();
    const useAuthMock = jest.fn().mockReturnValue({
      setUser: setUserMock,
    });
    const useCheckAuthMock = jest.fn().mockReturnValue({
      isSuccess: true,
      start: startMock,
      data: { data: { name: "jeremy" } },
    });

    render(
      <CheckAuth _useAuth={useAuthMock} _useCheckAuth={useCheckAuthMock}>
        <div>Test</div>
      </CheckAuth>
    );
    expect(screen.getByText("Test")).toBeTruthy();
    expect(startMock).toHaveBeenCalledTimes(1);
    expect(setUserMock).toHaveBeenCalledTimes(1);
  });
});

describe("Require Auth", () => {
  it("should render children", () => {
    const children = <div>Test</div>;
    const useAuthMock = jest.fn().mockReturnValue({
      user: "jeremy",
    });
    const useLocationMock = jest.fn().mockReturnValue("/");

    render(
      <RequireAuth _useAuth={useAuthMock} _useLocation={useLocationMock}>
        {children}
      </RequireAuth>
    );
    expect(screen.getByText("Test")).toBeTruthy();
  });
  it("should navigate to login", () => {
    const children = <div>Test</div>;
    const useAuthMock = jest.fn().mockReturnValue({});
    const useLocationMock = jest.fn().mockReturnValue("/");
    const screen = RequireAuth({
      children,
      _useAuth: useAuthMock,
      _useLocation: useLocationMock,
    });
    expect(screen).toMatchObject(
      <Navigate to={"/login"} state={{ from: "/" }} replace />
    );
  });
});

describe("Already Auth", () => {
  it("should render children", () => {
    const children = <div>Test</div>;
    const useAuthMock = jest.fn().mockReturnValue({});
    const useLocationMock = jest.fn().mockReturnValue("/");
    render(
      <AlreadyAuth _useAuth={useAuthMock} _useLocation={useLocationMock}>
        {children}
      </AlreadyAuth>
    );
    expect(screen.getByText("Test")).toBeTruthy();
  });
  it("should navigate to user previous page", () => {
    const children = <div>Test</div>;
    const useAuthMock = jest.fn().mockReturnValue({ user: "jeremy" });
    const useLocationMock = jest.fn().mockReturnValue("/");
    const screen = AlreadyAuth({
      children,
      _useAuth: useAuthMock,
      _useLocation: useLocationMock,
    });
    expect(screen).toMatchObject(
      <Navigate to={"/"} state={{ from: "/" }} replace />
    );
  });
});

describe("AuthProvider", () => {
  it("should return auth context provide", () => {
    render(<AuthProvider>children</AuthProvider>);
  });
});
