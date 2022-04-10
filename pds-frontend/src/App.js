import React from "react";
import { RouterConfig } from "./router";
import { AuthProvider } from "./services";
import { LoadingProvider } from "./context";
import { CheckAuth } from "./services";

const App = () => {
  return (
    <>
      <AuthProvider>
        <CheckAuth>
          <LoadingProvider>
            <RouterConfig />
          </LoadingProvider>
        </CheckAuth>
      </AuthProvider>
    </>
  );
};

export default App;
