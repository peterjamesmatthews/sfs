import { createBrowserRouter } from "react-router-dom";
import Node from "../components/Node";
import Path from "../components/Path";

export default createBrowserRouter([
  {
    path: "*",
    element: (
      <>
        <Path />
        <Node />
      </>
    ),
  },
]);
