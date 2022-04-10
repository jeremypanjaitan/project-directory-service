import {useState} from "react";
const useFetchCall = (fetchCall, _useState = useState) => {
  const [fetchState, setFetchState] = _useState({
    isError: false,
    isSuccess: false,
    error: null,
    data: null,
  });

  const start = data => {
    setFetchState({...fetchState, isLoading: true});
    fetchCall(data)
      .then(res => {
        setFetchState({
          ...fetchState,
          data: res,
          isSuccess: true,
          isLoading: false,
        });
      })
      .catch(e => {
        setFetchState({
          ...fetchState,
          error: e,
          isError: true,
          isLoading: false,
        });
      });
  };
  const reset = () => {
    setFetchState({
      isError: false,
      isSuccess: false,
      error: null,
      data: null,
    });
  };
  return {
    ...fetchState,
    start,
    reset,
  };
};

export default useFetchCall;
