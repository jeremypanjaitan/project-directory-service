import { useState } from "react";
const useFetchCall = (fetchCall, _useState = useState) => {
  const [isLoading, setIsLoading] = _useState(false);
  const [isError, setIsError] = _useState(false);
  const [isSuccess, setIsSuccess] = _useState(false);
  const [error, setError] = _useState();
  const [data, setData] = _useState();

  const start = (data) => {
    setIsLoading(true);
    fetchCall(data)
      .then((res) => {
        setData(res);
        setIsSuccess(true);
        setIsLoading(false);
      })
      .catch((e) => {
        setError(e);
        setIsError(true);
        setIsLoading(false);
      });
  };
  const reset = () => {
    setIsLoading(false);
    setIsError(false);
    setIsSuccess(false);
    setError();
    setData();
  };
  return {
    isLoading,
    isError,
    isSuccess,
    error,
    data,
    start,
    reset,
  };
};

export default useFetchCall;
